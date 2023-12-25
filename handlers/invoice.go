package handlers

import (
	// "context"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	invoicedto "pengirimanbarang/dto/invoice"
	dto "pengirimanbarang/dto/result"
	"pengirimanbarang/models"
	"pengirimanbarang/repositories"

	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"

	// "github.com/signintech/gopdf"

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


func (h *handlerInvoice) GetSalesDetailBySale(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("sales idnya", id)


    salesDetails, err := h.InvoiceRepository.GetSalesDetailBySale(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "failed", Message: err.Error()})
    }

    return c.JSON(http.StatusOK, dto.SuccesResult{
        Status: "Success",
        Data:   salesDetails,
    })
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

	SubTotal, err := strconv.Atoi(c.FormValue("sub_total"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid user_id value"})
	}

	totalAmount, err := strconv.Atoi(c.FormValue("sub_total"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid user_id value"})
	}

	invoice := models.Invoices{
		IDSales:                saleID,
		IDUser:                 userID,
		NumberInvoice:          c.FormValue("invoice_number"),
		DateInvoice:            c.FormValue("invoice_date"),
		DueDate:                c.FormValue("due_date"),
		SubTotal:               SubTotal,
		Discount:               c.FormValue("discount"),
		PPN11:                  c.FormValue("ppn_11"),// Set PPN to 11%
		TotalAmount:            totalAmount,
		NoFakturPajak:          c.FormValue("no_faktur_pajak"),
		NoFakturPajakPengganti: c.FormValue("no_faktur_pajak_pengganti"),
		InvoiceDesc:			c.FormValue("invoice_description"),
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
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid sub total value"})
	}

	totalamount, err := strconv.Atoi(c.FormValue("total_amount"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid total amount value"})
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
		InvoiceDesc:			c.FormValue("invoice_description"),
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

func (h *handlerInvoice) UpdateApprove1(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}

	request := invoicedto.InvoiceRequest{
		Approve1:     c.FormValue("approve_1"),
		Approve1Date: c.FormValue("approve_1_date"),
		Approve1Desc: c.FormValue("approve_1_desc"),
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

	if request.Approve1 != "" {
		invoice.Approve1 = request.Approve1
	}

	if request.Approve1Date != "" {
		invoice.Approve1Date = request.Approve1Date
	} else {
		// If Approve1Date is not provided, set it to the current date
		now := time.Now()
		invoice.Approve1Date = now.Format("2006-01-02") // Format it as needed
	}

	if request.Approve1Desc != "" {
		invoice.Approve1Desc = request.Approve1Desc
	}

	data, err := h.InvoiceRepository.UpdateApprove1(invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Status: "Success", Data: data})
}

func (h *handlerInvoice) UpdateApprove2(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "failed", Message: "Invalid ID! Please input id as number."})
	}

	request := invoicedto.InvoiceRequest{
		Approve2:     c.FormValue("approve_2"),
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

	if request.Approve2 != "" {
		invoice.Approve2 = request.Approve2
	}

	if request.Approve2Date != "" {
		invoice.Approve2Date = request.Approve2Date
	} else {
		// If Approve2Date is not provided, set it to the current date
		now := time.Now()
		invoice.Approve2Date = now.Format("2006-01-02") // Format it as needed
	}

	if request.Approve2Desc != "" {
		invoice.Approve2Desc = request.Approve2Desc
	}

	data, err := h.InvoiceRepository.UpdateApprove2(invoice)
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

	salesDetails, err := h.InvoiceRepository.GetSalesDetailBySale(invoice.Sales.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: "Failed to fetch sales details"})
	}

    // Check if both approve_1 and approve_2 are okay
    if invoice.Approve1 == "ok" && invoice.Approve2 == "ok" {
		
        // Generate PDF
        marginX := 10.0
        marginY := 20.0
		var pdfBuffer bytes.Buffer
        pdf := gofpdf.New("P", "mm", "A4", "")
        pdf.SetMargins(marginX, marginY, marginX)
        pdf.AddPage()
        pdf.SetFont("Arial", "B", 12)

        pageWidth := 175.0
        xCoordinate := pageWidth / 2

        pdf.Ln(5)
        pdf.ImageOptions("asset/logo-arna.png", xCoordinate, 0, 35, 25, false,
            gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")

        // // Center "Invoice" and "PT Arwana Ceramics" at the top
        pdf.CellFormat(0, 10, "PT Arwana Ceramics", "", 0, "C", false, 0, "")
        pdf.Ln(10)
		pdf.SetFont("Arial", "B", 16) // Set font to Arial, bold, and size 16

		pdf.CellFormat(0, 10, "Invoice", "", 0, "C", false, 0, "")
		pdf.SetFont("Arial", "B", 12)
        pdf.Ln(20)

        // Add invoice details to PDF with flexible layout
        cellWidth := 95.0 // Adjust the width based on your layout
        lineHeight := 5.0 // Adjust the line height based on your layout

        // Create a new line and add "Kepada yth" to the left
        // Add "Kepada yth" to the left
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("Kepada yth: %s", invoice.Sales.Customer.NameCustomer), "", 0, "L", false, 0, "")

        // Add "Invoice Number" to the right in the same line
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("Invoice Number: %s", invoice.NumberInvoice), "", 0, "R", false, 0, "")
        pdf.Ln(lineHeight)
		
        // Add "Date" to the right in the same line
        pdf.CellFormat(cellWidth, lineHeight,fmt.Sprintf("     Alamat: %s", invoice.Sales.Customer.AddressCustomer),"", 0, "L", false, 0, "")
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("Date: %s", invoice.DateInvoice), "", 0, "R", false, 0, "")
        pdf.Ln(lineHeight)
		pdf.Ln(2 * lineHeight) 

        // Add "No. Faktur Pajak" to the right in the same line
        pdf.CellFormat(cellWidth, lineHeight, "", "", 0, "L", false, 0, "")
        pdf.CellFormat(cellWidth, lineHeight, fmt.Sprintf("No. Faktur Pajak: %s", invoice.NoFakturPajak), "", 0, "R", false, 0, "")
        pdf.Ln(lineHeight)

        pdf.Ln(2 * lineHeight) // Increased space after this line

		pdf.CellFormat(cellWidth, lineHeight, "Untuk pembayaraan barang - barang tersebut dibawah ini :", "", 0, "L", false, 0, "")
		pdf.Ln(2 * lineHeight)
		// pdf.CellFormat(0, 10, "Invoice", "", 0, "C", false, 0, "")
        // Draw a line after the Date
        // pdf.Line(marginX, pdf.GetY(), 210-marginX, pdf.GetY())
        // pdf.Ln(lineHeight)

		// Add a table with column headers
		pdf.SetFont("Arial", "B", 12)

		// Set the width and height of each cell in the table
		cellWidth = 50.0
		cellHeight := 10.0

		pdf.CellFormat(30.0, cellHeight, "No Order", "1", 0, "C", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, "Name Product", "1", 0, "C", false, 0, "")
		pdf.CellFormat(30.0, cellHeight, "Qty", "1", 0, "C", false, 0, "")
		pdf.CellFormat(30.0, cellHeight, "Price", "1", 0, "C", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, "Total", "1", 0, "C", false, 0, "")
		pdf.Ln(cellHeight) // Move to the next line

		for _, salesDetail := range salesDetails {
			// If there is not enough space for the next row, start a new page
			if pdf.GetY()+cellHeight > 280.0 { // Adjust this value based on your layout
				pdf.AddPage()
				pdf.Ln(cellHeight) // Move to the next line
			}

			// Add data to the table
			// deliveryNumber := fmt.Sprintf("%d", invoice.Sales.DeliveryOrderNumber)
			pdf.CellFormat(30.0, cellHeight,invoice.Sales.DeliveryOrderNumber, "1", 0, "C", false, 0, "")
			pdf.CellFormat(cellWidth, cellHeight, salesDetail.Product.NameProduct, "1", 0, "C", false, 0, "")
			quantity := fmt.Sprintf("%d", salesDetail.Qty)
			pdf.CellFormat(30.0, cellHeight, quantity, "1", 0, "C", false, 0, "")
			price := fmt.Sprintf("%d", salesDetail.Product.Price)
			pdf.CellFormat(30.0, cellHeight, price, "1", 0, "C", false, 0, "")
			Total := fmt.Sprintf("%d", invoice.Sales.AmountTotal)
			pdf.CellFormat(cellWidth, cellHeight, Total, "1", 0, "C", false, 0, "")
			pdf.Ln(cellHeight) // Move to the next line
		}
		pdf.Ln(cellHeight) 



        // pdf.CellFormat(cellWidth, cellHeight, fmt.Sprintf("%d", invoice.Sales.SalesDetail[0].Qty), "1", 0, "C", false, 0, "")
        pdf.Ln(cellHeight)
		// Move to the bottom of the page

		width := 95.0 // Adjust the width based on your layout
		height := 5.0 // Adjust the line height based on your layout

		pdf.CellFormat(width, height, "", "", 0, "L", false, 0, "")
		subtotal := fmt.Sprintf("%d", invoice.SubTotal)
		pdf.CellFormat(width, height, fmt.Sprintf("Subtotal: %s", subtotal), "", 0, "R", false, 0, "")
		pdf.Ln(lineHeight)

		pdf.CellFormat(width, height, "", "", 0, "L", false, 0, "")
		discount := fmt.Sprintf("%s", invoice.Discount) // Assuming Discount is already a percentage
		pdf.CellFormat(width, height, fmt.Sprintf("Discount: %s", discount), "", 0, "R", false, 0, "")
		pdf.Ln(lineHeight)

		pdf.CellFormat(width, height, "", "", 0, "L", false, 0, "")
		ppn := "11%" // Hardcoded as 11%
		pdf.CellFormat(width, height, fmt.Sprintf("Ppn: %s", ppn), "", 0, "R", false, 0, "")
		pdf.Ln(lineHeight)

		pdf.CellFormat(width, height, "", "", 0, "L", false, 0, "")
		grandtotal := fmt.Sprintf("%d", invoice.TotalAmount)
		pdf.CellFormat(width, height, fmt.Sprintf("Grand Total: %s", grandtotal), "", 0, "R", false, 0, "")
		pdf.Ln(2 * lineHeight)

		
		pdf.CellFormat(width, height, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(width, height, "Hormat Kami,", "", 0, "R", false, 0, "")
		pdf.Ln(2 * lineHeight)

		// Generate QR code
		qrData := fmt.Sprintf("No. Invoice: %s\nApprove MAnager %s", invoice.NumberInvoice, invoice.Approve2Date)
		qr, err := qrcode.New(qrData, qrcode.Medium)
		if err != nil {
			fmt.Println("QR Code Generation Error:", err)
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: "Failed to generate QR code"})
		}

		// Save QR code image to a file
		filename := "qrcode.png"
		err = qr.WriteFile(256, filename)  // Adjust the size as needed

		if err != nil {
			fmt.Println("QR Code WriteFile Error:", err)
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: "Failed to write QR code to file"})
		}


		// Add QR code image file to the PDF
		qrCodeX := xCoordinate + 80.0 // Adjust this value based on your layout
		qrCodeY := pdf.GetY()         // Align with the current Y-coordinate
		pdf.ImageOptions(filename, qrCodeX, qrCodeY, 25, 25, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")
		pdf.Ln(7 *lineHeight)


		pdf.CellFormat(width, height, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(width, height, "( Nama Manager )", "", 0, "R", false, 0, "")
		pdf.Ln(lineHeight)

		
		pdf.CellFormat(width, height, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(width, height, "Finance Manager", "", 0, "R", false, 0, "")
		pdf.Ln(lineHeight)

        // Output PDF to buffer
        if err := pdf.Output(&pdfBuffer); err != nil {
            fmt.Println("PDF Output Error:", err)
            return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: "Failed to generate PDF"})
        }

        // Set headers for download
        c.Response().Header().Set("Content-Type", "application/pdf")
        c.Response().Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")

        // Write the PDF content to the response
        _, err = c.Response().Writer.Write(pdfBuffer.Bytes())
        if err != nil {
            fmt.Println("Write to Response Error:", err)
            return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "Failed", Message: "Failed to write PDF to response"})
        }

        return nil
    }

    return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "Failed", Message: "Invoice is not approved for showing PDF."})
}













