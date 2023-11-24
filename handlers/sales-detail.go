package handlers

import (
	// "context"

	"fmt"
	"net/http"
	"strconv"

	dto "pengirimanbarang/dto/result"
	salesdetaildto "pengirimanbarang/dto/sales-detail"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"

	// "os"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerSalesDetail struct {
	SalesDetailRepository repositories.SalesDetailRepository
}

func HandlerSalesDetail(SalesDetailRepository repositories.SalesDetailRepository) *handlerSalesDetail {
	return &handlerSalesDetail{SalesDetailRepository}
}

func (h *handlerSalesDetail) FindSalesDetail(c echo.Context) error {
	salesdetail, err := h.SalesDetailRepository.FindSalesDetail()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   salesdetail})
}

func (h *handlerSalesDetail) GetSalesDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	salesdetail, err := h.SalesDetailRepository.GetSalesDetail(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}


	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   salesdetail})
}

func (h *handlerSalesDetail) CreateSalesDetail(c echo.Context) error {
	salesid, err := strconv.Atoi(c.FormValue("sale_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	productid, err := strconv.Atoi(c.FormValue("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	qty, err := strconv.Atoi(c.FormValue("qty"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid price value"})
	}

	amount, err := strconv.Atoi(c.FormValue("amount"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid product_category_id value"})
	}


	request := salesdetaildto.SalesDetailRequest{
		IDSales: salesid,
		IDProduct: productid,
		Qty: qty,
		Price: price,
		Amount: amount,
		Desc: c.FormValue("desc"),
		Status: c.FormValue("status"),
    }


    validation := validator.New()
    if errValid := validation.Struct(request); errValid != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: errValid.Error()})
    }

    salesdetail := models.SalesDetail{
		IDSales: request.IDSales,
		IDProduct: request.IDProduct,
		Qty: request.Qty,
		Price: request.Price,
		Amount: request.Amount,
		Desc: request.Desc,
		Status: request.Status,
    }

    // Call the CreateProductCategories method in the repository
    salesdetail, err = h.SalesDetailRepository.CreateSalesDetail(salesdetail)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    response := dto.SuccesResult{Status: "success", Data: salesdetail}
    return c.JSON(http.StatusOK, response)
}


func (h *handlerSalesDetail) UpdateSalesDetail(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}

	salesid, err := strconv.Atoi(c.FormValue("sale_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	productid, err := strconv.Atoi(c.FormValue("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	qty, err := strconv.Atoi(c.FormValue("qty"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid price value"})
	}

	amount, err := strconv.Atoi(c.FormValue("amount"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid product_category_id value"})
	}

	request := salesdetaildto.SalesDetailRequest{
		IDSales: salesid,
		IDProduct: productid,
		Qty: qty,
		Price: price,
		Amount: amount,
		Desc: c.FormValue("desc"),
		Status: c.FormValue("status"),
    }

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	salesdetail, err := h.SalesDetailRepository.GetSalesDetail(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	if request.IDSales != 0 {
		salesdetail.IDSales = request.IDSales
	}

	if request.IDProduct != 0 {
		salesdetail.IDProduct = request.IDProduct
	}

	if request.Qty != 0 {
		salesdetail.Qty = request.Qty
	}

	if request.Price != 0 {
		salesdetail.Price = request.Price
	}

	if request.Amount != 0 {
		salesdetail.Amount = request.Amount
	}

	if request.Desc != "" {
		salesdetail.Desc = request.Desc
	}

	if request.Status != "" {
		salesdetail.Status = request.Status
	}

	data, err := h.SalesDetailRepository.UpdateSalesDetail(salesdetail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: data})
}

// func (h *handlerSalesDetail) DeleteSalesDetail(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	salesdetail, err := h.SalesDetailRepository.GetSalesDetail(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
// 	}

// 	data, err := h.SalesDetailRepository.DeleteSalesDetail(salesdetail)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: data})
// }