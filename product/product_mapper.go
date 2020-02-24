package product

func ToProduct(productDto ProductDto) Product {
	return Product{Code: productDto.Code, Price: productDto.Price}
}

func ToProductDto(product Product) ProductDto {
	return ProductDto{ID: product.ID, Price: product.Price, Code: product.Code}
}

func ToProductDtos(products []Product) []ProductDto {
	productDtos := make([]ProductDto, len(products))

	for i, itm := range products {
		productDtos[i] = ToProductDto(itm)
	}

	return productDtos
}
