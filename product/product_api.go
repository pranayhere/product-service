package product

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductService ProductService
}

func ProvideProductAPI(p ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()
	c.JSON(http.StatusOK, gin.H{"products": ToProductDtos(products)})
}

func (p *ProductAPI) FindById(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindById(uint(id));

	c.JSON(http.StatusOK, gin.H{"product": ToProductDto(product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDto ProductDto
	err := c.BindJSON(&productDto)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	product := p.ProductService.Save(ToProduct(productDto))
	c.JSON(http.StatusCreated, gin.H{"product": ToProductDto(product)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDto ProductDto
	err := c.BindJSON(&productDto)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindById(uint(id))
	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDto.Code
	product.Price = productDto.Price
	p.ProductService.Save(product)

	c.Status(http.StatusOK)
}

func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindById(uint(id))
	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}