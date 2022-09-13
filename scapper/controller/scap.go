package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

// Course stores information about a coursera course
type Product struct {
	Title       string  `json:"title"`
	ImageUrl    string  `json:"imageurl"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	NoOfReviews string  `json:"totalreview"`
}

type request struct {
	Url string `json:"url"`
}

func ScapWeb(c *gin.Context) {

	req := request{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error":   err.Error(),
			"message": "given req is not same as desired",
			"data":    nil,
		})
		return
	}

	if req.Url == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "url is empty",
			"message": "please provide url",
			"data":    nil,
		})
		return
	}

	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:amazon\.)?([^:\/\n]+)`)
	if !re.MatchString(req.Url) {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error":   "url should only scap amazon",
			"message": "url should only scap amazon",
			"data":    nil,
		})
		return
	}

	scapper := colly.NewCollector()
	scapper.SetRequestTimeout(200 * time.Second)

	p := Product{}
	scapper.OnHTML("div#dp-container", func(e *colly.HTMLElement) {
		p.ImageUrl = e.ChildAttr("img", "src")
		p.Title = e.ChildText("span#productTitle")
		e.ForEach("ul.a-unordered-list.a-vertical.a-spacing-mini > li > span.a-list-item", func(i int, h *colly.HTMLElement) {
			p.Description += h.Text
		})
		p.NoOfReviews = strings.Split(e.ChildText("span#acrCustomerReviewText"), " ")[0]
		p.Price, err = strconv.ParseFloat(strings.Split(e.ChildText("span[aria-hidden='true'] > span.a-price-whole"), ".")[0], 64)
		if err != nil {
			fmt.Println("error while fetching prize, so setting it to zero", err)
		}
	})

	scapper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	scapper.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	scapper.Visit(req.Url)

	scapper.Wait()

	serviceReq := map[string]interface{}{
		"url":     req.Url,
		"product": p,
	}

	reqForStore, err := json.Marshal(serviceReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err.Error(),
			"message": "something went wrong",
			"data":    nil,
		})
		return
	}

	storeURL := fmt.Sprintf("http://%s:%s/%s", os.Getenv("SERVICE_URL"), os.Getenv("SERVICE_PORT"), os.Getenv("WRITE_SERVICE"))
	reqObj, err := http.NewRequest("POST", storeURL, bytes.NewBuffer(reqForStore))
	reqObj.Header.Set("content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(reqObj)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error":   nil,
			"message": "url added succefully!",
			"data":    p,
		})
		return
	} else if response.StatusCode == http.StatusAccepted {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error":   nil,
			"message": "url updated succefully!",
			"data":    p,
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "something went wrong with store",
			"message": "something went wrong with store",
			"data":    nil,
		})
		return
	}
}
