package handlers

import (
	// "context"

	"fmt"
	"net/http"
	"strconv"

	productcategoriesdto "pengirimanbarang/dto/product-categories"
	dto "pengirimanbarang/dto/result"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"

	// "os"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerProductCategories struct {
	ProductCategoriesRepository repositories.ProductCategoriesRepository
}

func HandlerProductCategories(ProductCategoriesRepository repositories.ProductCategoriesRepository) *handlerProductCategories {
	return &handlerProductCategories{ProductCategoriesRepository}
}

func (h *handlerProductCategories) FindProductCategories(c echo.Context) error {
	productCategories, err := h.ProductCategoriesRepository.FindProductCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   productCategories})
}

func (h *handlerProductCategories) GetProductCategories(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	productCategories, err := h.ProductCategoriesRepository.GetProductCategories(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}


	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   productCategories})
}

func (h *handlerProductCategories) CreateProductCategories(c echo.Context) error {
    request := productcategoriesdto.ProductCategoriesRequest{
        NameCategoryProduct: c.FormValue("product_category_name"),
        Desc:               c.FormValue("desc"),
    }

    validation := validator.New()
    if errValid := validation.Struct(request); errValid != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: errValid.Error()})
    }

    productcategories := models.ProductCategories{
        NameCategoryProduct: request.NameCategoryProduct,
        Desc:               request.Desc,
        CreatedBy:          c.FormValue("created_by"),
    }

    // Call the CreateProductCategories method in the repository
    productcategories, err := h.ProductCategoriesRepository.CreateProductCategories(productcategories)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    response := dto.SuccesResult{Status: "success", Data: productcategories}
    return c.JSON(http.StatusOK, response)
}

func (h *handlerProductCategories) UpdateProductCategories(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}

	request := productcategoriesdto.ProductCategoriesRequest{
		NameCategoryProduct: c.FormValue("product_category_name"),
		Desc: c.FormValue("desc"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	productcategories, err := h.ProductCategoriesRepository.GetProductCategories(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	if request.NameCategoryProduct != "" {
		productcategories.NameCategoryProduct = request.NameCategoryProduct
	}

	if request.Desc != "" {
		productcategories.Desc = request.Desc
	}

	data, err := h.ProductCategoriesRepository.UpdateProductCategories(productcategories)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: data})
}

func (h *handlerProductCategories) DeleteProductCategories(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    productcategories, err := h.ProductCategoriesRepository.GetProductCategories(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    // Delete associated products
    err = h.ProductCategoriesRepository.DeleteProductsByCategoryID(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    // Delete the product category
    data, err := h.ProductCategoriesRepository.DeleteProductCategories(productcategories)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: data})
}

func (h *handlerProductCategories) CancelProductCategories(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as a number."})
    }

    // Get the sale details
    productCategories, err := h.ProductCategoriesRepository.GetProductCategories(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    // Update the Status field to 1 (assuming 1 means canceled)
    productCategories.Status = 1

    // Call the UpdateSale method in the repository
    updatedProductCategories, err := h.ProductCategoriesRepository.UpdateProductCategories(productCategories)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: updatedProductCategories})
}