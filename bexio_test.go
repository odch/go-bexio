package bexio

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	c := NewClient(srv.Client())
	c.BaseURL, _ = url.Parse(srv.URL + "/")
	return c
}

func TestCreateContact(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/contact" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"id":42,"name_1":"Acme Flying Club"}`))
	})

	got, err := c.Contacts.CreateContact(Contact{Name_1: "Acme Flying Club"})
	if err != nil {
		t.Fatalf("CreateContact: %v", err)
	}
	if got.Id != 42 {
		t.Errorf("id = %d, want 42", got.Id)
	}
}

func TestIssueInvoice(t *testing.T) {
	called := false
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		called = true
		if r.Method != http.MethodPost || r.URL.Path != "/kb_invoice/7/issue" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"success":true}`))
	})

	if err := c.Invoices.IssueInvoice(7); err != nil {
		t.Fatalf("IssueInvoice: %v", err)
	}
	if !called {
		t.Error("issue endpoint not called")
	}
}

func TestGetInvoicePDF(t *testing.T) {
	want := []byte("%PDF-1.4 fake pdf bytes")
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/kb_invoice/7/pdf" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		body := `{"name":"invoice.pdf","mime":"application/pdf","content":"` +
			base64.StdEncoding.EncodeToString(want) + `"}`
		_, _ = w.Write([]byte(body))
	})

	meta, data, err := c.Invoices.GetInvoicePDF(7)
	if err != nil {
		t.Fatalf("GetInvoicePDF: %v", err)
	}
	if meta.Mime != "application/pdf" {
		t.Errorf("mime = %q", meta.Mime)
	}
	if string(data) != string(want) {
		t.Errorf("pdf bytes = %q, want %q", data, want)
	}
}
