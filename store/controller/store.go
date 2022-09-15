package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/store/model"
	"github.com/store/types"
)

func APIUpsertProduct(c *gin.Context) {

	req := types.StoreRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error":   err.Error(),
			"message": "given req is not same as desired",
			"data":    nil,
		})
		return
	}
	fmt.Println("req", req)
	updated, err := model.UpsertProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err.Error(),
			"message": "something went wrong",
			"data":    nil,
		})
		return
	}

	switch updated {
	default:
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "something went wrong",
			"message": "something went wrong",
			"data":    nil,
		})
		return
	case 1:
		c.JSON(http.StatusAccepted, map[string]interface{}{
			"error":   nil,
			"message": "updated",
			"data":    nil,
		})
		return
	case 0:
		c.JSON(http.StatusCreated, map[string]interface{}{
			"error":   nil,
			"message": "inserted",
			"data":    nil,
		})
		return
	}

}
