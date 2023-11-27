package handlers

import (
	// "context"
	"fmt"
	"net/http"
	"strconv"

	invoicedto "pengirimanbarang/dto/invoice"
	dto "pengirimanbarang/dto/result"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"

	"github.com/jung-kurt/gofpdf"

	// "os"

	// "github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerInvoice struct {
	InvoiceRepository repositories.InvoiceRepository
}

func HandlerInvoice(InvoiceRepository repositories.InvoiceRepository) *handlerInvoice {
	return &handlerInvoice{InvoiceRepository}
}

func (h *handlerInvoice) FindInvoices(c echo.Context) error {
	invoice, err := h.InvoiceRepository.FindInvoices()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   invoice})
}

func (h *handlerInvoice) GetInvoice(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := h.InvoiceRepository.GetInvoice(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Status:  "Error",
			Message: err.Error()})
	}


	return c.JSON(http.StatusOK, dto.SuccesResult{
		Status: "Success",
		Data:   invoice})
}

func (h *handlerInvoice) CreateInvoice(c echo.Context) error {
	saleID, err := strconv.Atoi(c.FormValue("sale_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid sale_id value"})
	}

	userID, err := strconv.Atoi(c.FormValue("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid user_id value"})
	}

	// Fetch sales data using the provided saleID
	salesData, err := h.InvoiceRepository.GetSales(saleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: "Failed to fetch sales data"})
	}

	// Calculate SubTotal based on salesData.AmountTotal
	subTotal := salesData.AmountTotal

	// Apply Discount if provided
	if discountStr := c.FormValue("discount"); discountStr != "" {
		discount, err := strconv.Atoi(discountStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid discount value"})
		}
		subTotal -= discount
	}

	// Automatically set PPN to 11%
	ppnPercentage := 11
	ppnAmount := (ppnPercentage * subTotal) / 100
	subTotal += ppnAmount

	invoice := models.Invoices{
		IDSales:                saleID,
		IDUser:                 userID,
		NumberInvoice:          c.FormValue("invoice_number"),
		DateInvoice:            c.FormValue("invoice_date"),
		DueDate:                c.FormValue("due_date"),
		SubTotal:               salesData.AmountTotal,
		Discount:               c.FormValue("discount"),
		PPN11:                  strconv.Itoa(ppnPercentage), // Set PPN to 11%
		TotalAmount:            subTotal,
		NoFakturPajak:          c.FormValue("no_faktur_pajak"),
		NoFakturPajakPengganti: c.FormValue("no_faktur_pajak_pengganti"),
		InvoiceStatus:          c.FormValue("invoice_status"),
		Approve1:               c.FormValue("approve_1"),
		Approve1Date:           c.FormValue("approve_1_date"),
		Approve1Desc:           c.FormValue("approve_1_desc"),
		Approve2:               c.FormValue("approve_2"),
		Approve2Date:           c.FormValue("approve_2_date"),
		Approve2Desc:           c.FormValue("approve_2_desc"),
	}

	// Call the CreateInvoice method in the repository
	invoice, err = h.InvoiceRepository.CreateInvoice(invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
	}

	response := dto.SuccesResult{Status: "success", Data: invoice}
	return c.JSON(http.StatusOK, response)
}

func (h *handlerInvoice) UpdateInvoice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}

	subtotal, err := strconv.Atoi(c.FormValue("sub_total"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	totalamount, err := strconv.Atoi(c.FormValue("total_amount"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid unit value"})
	}

	request := invoicedto.InvoiceRequest{
		NumberInvoice: c.FormValue("invoice_number"),
		DateInvoice: c.FormValue("invoice_date"),
		DueDate: c.FormValue("due_date"),
		SubTotal: subtotal,
		Discount: c.FormValue("discount"),
		PPN11: c.FormValue("ppn_11"),
		TotalAmount: totalamount,
		NoFakturPajak: c.FormValue("no_faktur_pajak"),
		NoFakturPajakPengganti: c.FormValue("no_faktur_pajak_pengganti"),
		InvoiceStatus: c.FormValue("invoice_status"),
		Approve1: c.FormValue("approve_1"),
		Approve1Date: c.FormValue("approve_1_date"),
		Approve1Desc: c.FormValue("approve_1_desc"),
		Approve2: c.FormValue("approve_2"),
		Approve2Date: c.FormValue("approve_2_date"),
		Approve2Desc: c.FormValue("approve_2_desc"),
    }

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	invoice, err := h.InvoiceRepository.GetInvoice(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	if request.NumberInvoice != "" {
		invoice.NumberInvoice = request.NumberInvoice
	}

	if request.DateInvoice != "" {
		invoice.DateInvoice = request.DateInvoice
	}

	if request.DueDate != "" {
		invoice.DueDate = request.DueDate
	}
	
	if request.SubTotal != 0 {
		invoice.SubTotal = request.SubTotal
	}

	if request.Discount != "" {
		invoice.Discount = request.Discount
	}

	if request.PPN11 != "" {
		invoice.PPN11 = request.PPN11
	}

	if request.TotalAmount != 0 {
		invoice.TotalAmount = request.TotalAmount
	}

	if request.NoFakturPajak != "" {
		invoice.NoFakturPajak = request.NoFakturPajak
	}

	if request.NoFakturPajakPengganti != "" {
		invoice.NoFakturPajakPengganti = request.NoFakturPajakPengganti
	}

	if request.InvoiceDesc != "" {
		invoice.InvoiceDesc = request.InvoiceDesc
	}

	if request.InvoiceStatus != "" {
		invoice.InvoiceStatus = request.InvoiceStatus
	}

	if request.Approve1 != "" {
		invoice.Approve1 = request.Approve1
	}

	if request.Approve1Date != "" {
		invoice.Approve1Date = request.Approve1Date
	}

	if request.Approve1Desc != "" {
		invoice.Approve1Desc = request.Approve1Desc
	}

	if request.Approve2 != "" {
		invoice.Approve2 = request.Approve2
	}

	if request.Approve2Date != "" {
		invoice.Approve2Date = request.Approve2Date
	}

	if request.Approve2Desc != "" {
		invoice.Approve2Desc = request.Approve2Desc
	}


	data, err := h.InvoiceRepository.UpdateInvoice(invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: data})
}

func (h *handlerInvoice) DeleteInvoice(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

    // Get the invoice details
    invoice, err := h.InvoiceRepository.GetInvoice(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }
    // Delete the sales
    _, err = h.InvoiceRepository.DeleteSale(invoice.Sales)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    // Delete the invoice
    _, err = h.InvoiceRepository.DeleteInvoice(invoice)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "success", Data: "Invoice, Sales, and SalesDetails deleted successfully"})
}

func (h *handlerInvoice) CancelInvoice(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as a number."})
    }

    // Get the invoice details
    invoice, err := h.InvoiceRepository.GetInvoice(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    // Call the CancelInvoice method in the repository
    updatedInvoice, err := h.InvoiceRepository.CancelInvoice(invoice)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
    }

    return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: updatedInvoice})
}



func (h *handlerInvoice) PrintInvoice(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as a number."})
    }

    invoice, err := h.InvoiceRepository.GetInvoice(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: err.Error()})
    }

    // Check if both approve_1 and approve_2 are okay
    if invoice.Approve1 == "ok" && invoice.Approve2 == "ok" {
        // Generate PDF
        marginX := 10.0
        marginY := 20.0
        pdf := gofpdf.New("P", "mm", "A4", "")
        pdf.SetMargins(marginX, marginY, marginX)
        pdf.AddPage()
        pdf.SetFont("Arial", "B", 14)

		pageWidth := 175.0 
		xCoordinate := pageWidth  / 2

		pdf.Ln(5)
		pdf.ImageOptions("asset/logo-arna.png", xCoordinate, 0, 35, 25, false, 
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
	

        // Center "Invoice" and "PT Arwana Ceramics" at the top
        pdf.CellFormat(0, 10, "Invoice", "", 0, "C", false, 0, "")
        pdf.Ln(5)
        pdf.CellFormat(0, 10, "PT Arwana Ceramics", "", 0, "C", false, 0, "")
        pdf.Ln(20)

        // Add invoice details to PDF with flexible layout
        cellWidth := 95.0 // Adjust the width based on your layout
        lineHeight := 5.0 // Adjust the line height based on your layout

        // Create a new line and add "Kepada yth" to the left
        // Add "Kepada yth" to the left
		pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("Kepada yth: %s", invoice.Sales.Customer.NameCustomer), "", 0, "L", false, 0, "")
		fmt.Println("Sales", invoice.Sales)
		fmt.Println("customer", invoice.Sales.Customer)


        // Add "Invoice Number" to the right in the same line
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("Invoice Number: %s", invoice.NumberInvoice), "", 0, "R", false, 0, "")
        pdf.Ln(lineHeight)

        // Add "Date" to the right in the same line
		pdf.CellFormat(cellWidth, lineHeight, "", "", 0, "L", false, 0, "")
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("Date: %s", invoice.DateInvoice), "", 0, "R", false, 0, "")
        pdf.Ln(lineHeight)

        // Add "No. Faktur Pajak" to the right in the same line
		pdf.CellFormat(cellWidth, lineHeight, "", "", 0, "L", false, 0, "")
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("No. Faktur Pajak: %s", invoice.NoFakturPajak), "", 0, "R", false, 0, "")
        pdf.Ln(lineHeight)

        // Add "No. Faktur Pajak Pengganti" to the right in the same line
		pdf.CellFormat(cellWidth, lineHeight, "", "", 0, "L", false, 0, "")
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("No.Faktur Pajak Pengganti: %s", invoice.NoFakturPajakPengganti), "", 0, "R", false, 0, "")
        pdf.Ln(2 * lineHeight) // Increased space after this line

        // Draw a line after the Date
        // pdf.Line(marginX, pdf.GetY(), 210-marginX, pdf.GetY())
        // pdf.Ln(lineHeight)

        // Add a table with column headers
        pdf.SetFont("Arial", "B", 12)

        // Set the width and height of each cell in the table
        cellWidth = 95.0
        cellHeight := 10.0

        pdf.CellFormat(cellWidth, cellHeight, "Delivery Order Number", "1", 0, "C", false, 0, "")
        pdf.CellFormat(cellWidth, cellHeight, "Quantity", "1", 0, "C", false, 0, "")
        pdf.CellFormat(cellWidth, cellHeight, "sales Description", "1", 0, "C", false, 0, "")
        // pdf.CellFormat(cellWidth, cellHeight, "Total Amount", "1", 0, "C", false, 0, "")
        pdf.Ln(cellHeight) // Move to the next line

        // Set font back to the original
        pdf.SetFont("Arial", "B", 14)

        // Add data to the table
		deliverynumber := fmt.Sprintf("%s", strconv.Itoa(invoice.Sales.DeliveryOrderNumber))
		pdf.CellFormat(cellWidth, cellHeight, deliverynumber, "1", 0, "C", false, 0, "")
        pdf.CellFormat(cellWidth, cellHeight, fmt.Sprintf("%s", invoice.Sales.DescriptionSale), "1", 0, "C", false, 0, "")
        // pdf.CellFormat(cellWidth, cellHeight, fmt.Sprintf("%s%%", invoice.PPN11), "1", 0, "C", false, 0, "")
		// totalAmount := fmt.Sprintf("%s", strconv.Itoa(invoice.TotalAmount))
		// pdf.CellFormat(cellWidth, cellHeight, totalAmount, "1", 0, "C", false, 0, "")
        // pdf.CellFormat(cellWidth, cellHeight, fmt.Sprintf("%s", invoice.TotalAmount), "1", 0, "C", false, 0, "")
        pdf.Ln(cellHeight)

        // Draw a line after the table
        pdf.Line(marginX, pdf.GetY(), 210-marginX, pdf.GetY())
        pdf.Ln(lineHeight)

        // Output the PDF as a response
        err = pdf.Output(c.Response().Writer)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: "Failed to generate PDF"})
        }

        c.Response().Header().Set("Content-Type", "application/pdf")
        c.Response().Header().Set("Content-Disposition", "inline; filename=invoice.pdf")

        fmt.Printf("Invoice %d PDF generated and shown successfully!\n", id)

        return nil
    }

    return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: "Invoice is not approved for showing PDF."})
}









