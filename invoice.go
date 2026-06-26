package bexio

import (
	"encoding/base64"
	"fmt"
)

type InvoiceService struct {
	client *Client
}

// InvoicePdf is the response of the invoice PDF endpoint; Content is base64.
type InvoicePdf struct {
	Name    string `json:"name"`
	Size    int    `json:"size"`
	Mime    string `json:"mime"`
	Content string `json:"content"`
}

func (c *InvoiceService) ListInvoices() ([]Invoice, error) {
	req, err := c.client.newRequest("GET", "kb_invoice", nil)
	if err != nil {
		return nil, err
	}
	var invoice []Invoice
	_, err = c.client.do(req, &invoice)
	return invoice, err
}

func (c *InvoiceService) GetInvoice(id int) (Invoice, error) {
	var invoice Invoice
	req, err := c.client.newRequest("GET", fmt.Sprintf("kb_invoice/%d", id), nil)
	if err != nil {
		return invoice, err
	}
	_, err = c.client.do(req, &invoice)
	return invoice, err
}

func (c *InvoiceService) CreateInvoice(invoice Invoice) (Invoice, error) {
	req, err := c.client.newRequest("POST", "kb_invoice", invoice)
	if err != nil {
		return invoice, err
	}
	_, err = c.client.do(req, &invoice)
	return invoice, err
}

// IssueInvoice finalizes a draft invoice (draft -> open), assigning the
// definitive document number. It does NOT send the invoice to the contact.
func (c *InvoiceService) IssueInvoice(id int) error {
	req, err := c.client.newRequest("POST", fmt.Sprintf("kb_invoice/%d/issue", id), nil)
	if err != nil {
		return err
	}
	var result struct {
		Success bool `json:"success"`
	}
	_, err = c.client.do(req, &result)
	return err
}

// GetInvoicePDF fetches the rendered invoice PDF. Returns the metadata and the
// decoded PDF bytes.
func (c *InvoiceService) GetInvoicePDF(id int) (InvoicePdf, []byte, error) {
	var pdf InvoicePdf
	req, err := c.client.newRequest("GET", fmt.Sprintf("kb_invoice/%d/pdf", id), nil)
	if err != nil {
		return pdf, nil, err
	}
	if _, err = c.client.do(req, &pdf); err != nil {
		return pdf, nil, err
	}
	data, err := base64.StdEncoding.DecodeString(pdf.Content)
	if err != nil {
		return pdf, nil, fmt.Errorf("decode invoice pdf: %w", err)
	}
	return pdf, data, nil
}
