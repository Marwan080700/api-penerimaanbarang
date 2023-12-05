package handlers

import (
	// "context"
	"fmt"
	"net/http"

	productdto "pengirimanbarang/dto/product"
	dto "pengirimanbarang/dto/result"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"

	// "os"
	"strconv"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)


type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(c echo.Context) error {
	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   products})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductRepository.GetProduct(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}


	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   product})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid price value"})
	}

	idCategoryProduct, err := strconv.Atoi(c.FormValue("product_category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid product_category_id value"})
	}

	request := productdto.ProductRequest{
		IdentityProduct:   c.FormValue("product_identity"),
		IDCategoryProduct: idCategoryProduct,
		NameProduct:       c.FormValue("product_name"),
		Unit:              c.FormValue("unit"),
		Price:             price,
		Desc:              c.FormValue("desc"),
		CreatedBy:         c.FormValue("created_by"),
		UpdatedBy:         c.FormValue("updated_by"),
	}

	validation := validator.New()
	if errValid := validation.Struct(request); errValid != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: errValid.Error()})
	}

	product := models.Product{
		IdentityProduct:   request.IdentityProduct,
		IDCategoryProduct: request.IDCategoryProduct,
		NameProduct:       request.NameProduct,
		Unit:              request.Unit,
		Price:             request.Price,
		Desc:              request.Desc,
		CreatedBy:         request.CreatedBy,
		UpdatedBy:         request.UpdatedBy,
	}

	var errProduct error
	product, errProduct = h.ProductRepository.CreateProduct(product)
	if errProduct != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: errProduct.Error()})
	}

	var errFetch error
	product, errFetch = h.ProductRepository.GetProduct(product.ID)
	if errFetch != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: errFetch.Error()})
	}

	response := dto.SuccesResult{Status: "success", Data: product}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}


	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid price value"})
	}

	idCategoryProduct, err := strconv.Atoi(c.FormValue("product_category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid product_category_id value"})
	}


	request := productdto.ProductRequest{
		IdentityProduct:   c.FormValue("product_identity"),
		IDCategoryProduct: idCategoryProduct,
		NameProduct:       c.FormValue("product_name"),
		Unit:              c.FormValue("unit"),
		Price:             price,
		Desc:              c.FormValue("desc"),
		CreatedBy:         c.FormValue("created_by"),
		UpdatedBy:         c.FormValue("updated_by"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	if request.IdentityProduct != "" {
		product.IdentityProduct = request.IdentityProduct
	}

	if request.IDCategoryProduct != 0 {
		product.IDCategoryProduct = request.IDCategoryProduct
	}

	if request.NameProduct != "" {
		product.NameProduct = request.NameProduct
	}

	if request.Unit != "" {
		product.Unit = request.Unit
	}
	if request.Price != 0 {
		product.Price = request.Price
	}
	if request.Desc != "" {
		product.Desc = request.Desc
	}

	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: data})
}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	data, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: data})
}

func (h *handlerProduct) GetProductsByCategoryID(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID"})
    }

    products, err := h.ProductRepository.GetProductsByCategoryID(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    return c.JSON(http.StatusOK, dto.SuccesResult{
        Status: "Success",
        Data:   products,
    })
}
