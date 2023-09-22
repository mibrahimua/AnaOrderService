package controller

import (
	"AnaOrderService/request"
	"AnaOrderService/response"
	"AnaOrderService/service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(userService *service.OrderService) *OrderController {
	return &OrderController{userService}
}

// @Summary		Get User By Id
// @Description	Get User By Id
// @Produce		json
// @Param product_name body request.OrderRequest false "product_name"
// @Success		200	{object} model.Product
// @Router			/order/checkout [post]
func (uc *OrderController) CheckoutItems(c *gin.Context) {
	request := request.OrderRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	data, err := uc.orderService.CheckoutItems(request)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Data:    data,
		Message: "success",
	})
}
