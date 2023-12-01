package handlers

import (
	// "context"

	"fmt"
	"net/http"
	"strconv"

	customerdto "pengirimanbarang/dto/customer"
	dto "pengirimanbarang/dto/result"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"

	// "os"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCustomer struct {
	CustomerRepository repositories.CustomerRepository
}

func HandlerCustomerCategories(CustomerRepository repositories.CustomerRepository) *handlerCustomer {
	return &handlerCustomer{CustomerRepository}
}

func (h *handlerCustomer) FindCustomer(c echo.Context) error {
	customer, err := h.CustomerRepository.FindCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   customer})
}

func (h *handlerCustomer) GetCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := h.CustomerRepository.GetCustomer(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}


	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   customer})
}

func (h *handlerCustomer) CreateCustomer(c echo.Context) error {
    request := customerdto.CustomerRequest{
        IdentityCustomer: c.FormValue("customer_identity"),
        NameCustomer: c.FormValue("customer_name"),
        EmailCustomer: c.FormValue("customer_email"),
        HandponeCustomer: c.FormValue("customer_handpone"),
        NpwpCustomer: c.FormValue("customer_npwp"),
        AddressCustomer: c.FormValue("customer_address"),
    }

    validation := validator.New()
    if errValid := validation.Struct(request); errValid != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: errValid.Error()})
    }

    customer := models.Customer{
        IdentityCustomer: request.IdentityCustomer,
        NameCustomer: request.NameCustomer,
        EmailCustomer: request.EmailCustomer,
        HandponeCustomer: request.HandponeCustomer,
        NpwpCustomer: request.NpwpCustomer,
        AddressCustomer: request.AddressCustomer,
    }

    // Call the CreateProductCategories method in the repository
    customer, err := h.CustomerRepository.CreateCustomer(customer)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    response := dto.SuccesResult{Status: "success", Data: customer}
    return c.JSON(http.StatusOK, response)
}

func (h *handlerCustomer) UpdateCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}

	request := customerdto.CustomerRequest{
        IdentityCustomer: c.FormValue("customer_identity"),
        NameCustomer: c.FormValue("customer_name"),
        EmailCustomer: c.FormValue("customer_email"),
        HandponeCustomer: c.FormValue("customer_handpone"),
        NpwpCustomer: c.FormValue("customer_npwp"),
        AddressCustomer: c.FormValue("customer_address"),
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	customer, err := h.CustomerRepository.GetCustomer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	if request.IdentityCustomer != "" {
		customer.IdentityCustomer = request.IdentityCustomer
	}

	if request.NameCustomer != "" {
		customer.NameCustomer = request.NameCustomer
	}

	if request.EmailCustomer != "" {
		customer.EmailCustomer = request.EmailCustomer
	}

	if request.HandponeCustomer != "" {
		customer.HandponeCustomer = request.HandponeCustomer
	}

	if request.NpwpCustomer != "" {
		customer.NpwpCustomer = request.NpwpCustomer
	}

	if request.AddressCustomer != "" {
		customer.AddressCustomer = request.AddressCustomer
	}


	data, err := h.CustomerRepository.UpdateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: data})
}

func (h *handlerCustomer) DeleteCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := h.CustomerRepository.GetCustomer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	data, err := h.CustomerRepository.DeleteCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: data})
}

func (h *handlerCustomer) CancelCustomer(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as a number."})
    }

    // Get the sale details
    customer, err := h.CustomerRepository.GetCustomer(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    // Update the Status field to 1 (assuming 1 means canceled)
    customer.Status = 1

    // Call the UpdateSale method in the repository
    updatedCustomer, err := h.CustomerRepository.UpdateCustomer(customer)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: updatedCustomer})
}