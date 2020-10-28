package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go-iris/services"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IProductService
}

func (c *ProductController) GetAll() mvc.View {
	productArray, _ := c.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productArray,
		},
	}
}
