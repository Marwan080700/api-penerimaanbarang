package handlers

import (
	// "context"

	"fmt"
	"net/http"
	"strconv"

	dto "pengirimanbarang/dto/result"
	salesdto "pengirimanbarang/dto/sales"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"

	// "os"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerSales struct {
	SalesRepository repositories.SalesRepository
}

func HandlerSales(SalesRepository repositories.SalesRepository) *handlerSales {
	return &handlerSales{SalesRepository}
}

func (h *handlerSales) FindSales(c echo.Context) error {
	sales, err := h.SalesRepository.FindSales()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   sales})
}

func (h *handlerSales) GetSale(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sales, err := h.SalesRepository.GetSale(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}


	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   sales})
}

func (h *handlerSales) CreateSale(c echo.Context) error {
	deliveryordernumber, err := strconv.Atoi(c.FormValue("delivery_order_number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	customerid, err := strconv.Atoi(c.FormValue("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	userid, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	totalamount, err := strconv.Atoi(c.FormValue("total_amount"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	request := salesdto.SalesRequest{
		DeliveryOrderNumber: deliveryordernumber,
		IDCustomer: customerid,
		IDUser: userid,
		DateSale: c.FormValue("sale_date"),
		DescriptionSale: c.FormValue("sale_description"),
		StatusSale: c.FormValue("sale_status"),
		AmountTotal: totalamount,
    }


    validation := validator.New()
    if errValid := validation.Struct(request); errValid != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: errValid.Error()})
    }

    sales := models.Sales{
		DeliveryOrderNumber: request.DeliveryOrderNumber,
		IDCustomer: request.IDCustomer,
		IDUser: request.IDUser,
		DateSale: request.DateSale,
		DescriptionSale: request.DescriptionSale,
		StatusSale: request.StatusSale,
		AmountTotal: request.AmountTotal,
    }

    // Call the CreateProductCategories method in the repository
    sales, err = h.SalesRepository.CreateSale(sales)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    response := dto.SuccesResult{Status: "success", Data: sales}
    return c.JSON(http.StatusOK, response)
}


func (h *handlerSales) UpdateSale(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}

	deliveryordernumber, err := strconv.Atoi(c.FormValue("delivery_order_number"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	customerid, err := strconv.Atoi(c.FormValue("customer_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	userid, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	totalamount, err := strconv.Atoi(c.FormValue("total_amount"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	request := salesdto.SalesRequest{
		DeliveryOrderNumber: deliveryordernumber,
		IDCustomer: customerid,
		IDUser: userid,
		DateSale: c.FormValue("sale_date"),
		DescriptionSale: c.FormValue("sale_description"),
		StatusSale: c.FormValue("sale_status"),
		AmountTotal: totalamount,
    }

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	sales, err := h.SalesRepository.GetSale(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	if request.DeliveryOrderNumber != 0 {
		sales.DeliveryOrderNumber = request.DeliveryOrderNumber
	}

	if request.IDCustomer != 0 {
		sales.IDCustomer = request.IDCustomer
	}

	if request.IDUser != 0 {
		sales.IDUser = request.IDUser
	}

	if request.DateSale != "" {
		sales.DateSale = request.DateSale
	}

	if request.DescriptionSale != "" {
		sales.DescriptionSale = request.DescriptionSale
	}
	
	if request.StatusSale != "" {
		sales.StatusSale = request.StatusSale
	}

	if request.AmountTotal != 0 {
		sales.AmountTotal = request.AmountTotal
	}


	data, err := h.SalesRepository.UpdateSale(sales)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: data})
}

func (h *handlerSales) DeleteSale(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sales, err := h.SalesRepository.GetSale(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	data, err := h.SalesRepository.DeleteSale(sales)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: data})
}

func (h *handlerSales) DeleteSaleWithAssociatedData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// Call repository method to delete sale and associated data
	err := h.SalesRepository.DeleteSaleAndAssociatedData(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: "Sale and associated data deleted"})
}

func (h *handlerSales) CancelSale(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as a number."})
    }

    // Get the sale details
    sale, err := h.SalesRepository.GetSale(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    // Update the status of the sale to 1 (assuming 1 means canceled)
    sale.StatusSale = "canceled"

    // Update the Status field to 1 (assuming 1 means canceled)
    sale.Status = 1

    // Call the UpdateSale method in the repository
    updatedSale, err := h.SalesRepository.UpdateSale(sale)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: updatedSale})
}
