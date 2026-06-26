package bexio

import "encoding/json"

type ValidationError struct {
	ErrorCode int             `json:"error_code"`
	Message   string          `json:"message"`
	Errors    json.RawMessage `json:"errors"`
}

func (ve ValidationError) Error() string {
	err := string(ve.Errors)
	return ve.Message + err
}

type Contact struct {
	Id                 int64  `json:"id"`
	Contact_type_id    int    `json:"contact_type_id,omitempty"`
	Nr                 string `json:"nr"`
	Name_1             string `json:"name_1"`
	Name_2             string `json:"name_2"`
	Salutation_id      int    `json:"salutation_id"`
	Salutation_form    int    `json:"salutation_form"`
	Title_id           int    `json:"title_id"`
	Birthday           string `json:"birthday"`
	Address            string `json:"address"`
	Postcode           string `json:"postcode"`
	City               string `json:"city"`
	Country_id         int    `json:"country_id"`
	Mail               string `json:"mail"`
	Mail_second        string `json:"mail_second"`
	Phone_fixed        string `json:"phone_fixed"`
	Phone_fixed_second string `json:"phone_fixed_second"`
	Phone_mobile       string `json:"phone_mobile"`
	Fax                string `json:"fax"`
	Url                string `json:"url"`
	Skype_name         string `json:"skype_name"`
	Remarks            string `json:"remarks"`
	Language_id        int    `json:"language_Id"`
	Is_lead            bool   `json:"is_lead"`
	Contact_group_ids  string `json:"contact_group_ids"`
	Contact_branch_ids string `json:"contact_branch_ids"`
	User_id            int    `json:"user_id"`
	Owner_id           int    `json:"owner_id"`
	Updated_at         string `json:"updated_at"`
}

type ContactGroup struct {
	Id   int64
	Name string
}

type Article struct {
	Id                    int64
	User_id               int64
	Article_type_id       int64
	Contact_id            uint8
	Deliverer_code        string
	Deliverer_name        string
	Deliverer_description string
	Intern_code           string
	Intern_name           string
	Intern_description    string
	Purchase_price        string
	Sale_price            string
	Purchase_total        string
	Sale_total            string
	Currency_id           int64
	Tax_income_id         int64
	Tax_id                int64
	Tax_expense_id        int64
	Unit_id               int64
	Is_stock              bool
	Stock_id              string
	Stock_place_id        int64
	Stock_nr              int64
	Stock_min_nr          int64
	Stock_reserved_nr     int64
	Stock_available_nr    int64
	Stock_picked_nr       int64
	Stock_disposed_nr     int64
	Stock_ordered_nr      int64
	Width                 string
	Height                string
	Weight                string
	Volume                string
	Html_text             string
	Remarks               string
	Delivery_price        string
	Article_group_id      string
}

type Invoice struct {
	Id                        int64             `json:"id,omitempty"`
	Document_nr               string            `json:"document_nr,omitempty"`
	Title                     string            `json:"title,omitempty"`
	ContactID                 int64             `json:"contact_id"`
	Contact_sub_id            int64             `json:"contact_sub_id,omitempty"`
	User_id                   int64             `json:"user_id"`
	Project_id                int64             `json:"project_id,omitempty"`
	Logopaper_id              int64             `json:"logopaper_id,omitempty"`
	Language_id               int64             `json:"language_id,omitempty"`
	Bank_account_id           int64             `json:"bank_account_id,omitempty"`
	Currency_id               int64             `json:"currency_id,omitempty"`
	Payment_type_id           int64             `json:"payment_type_id,omitempty"`
	Header                    string            `json:"header,omitempty"`
	Footer                    string            `json:"footer,omitempty"`
	Total_gross               string            `json:"total_gross,omitempty"`
	Total_net                 string            `json:"total_net,omitempty"`
	Total_taxes               string            `json:"total_taxes,omitempty"`
	Total_received_payments   string            `json:"total_received_payments,omitempty"`
	Total_credit_vouchers     string            `json:"total_credit_vouchers,omitempty"`
	Total_remaining_payments  string            `json:"total_remaining_payments,omitempty"`
	Total                     string            `json:"total,omitempty"`
	Total_rounding_difference float64           `json:"total_rounding_difference,omitempty"`
	Mwst_type                 int               `json:"mwst_type,omitempty"`
	Mwst_is_net               bool              `json:"mwst_is_net,omitempty"`
	Show_position_taxes       bool              `json:"show_position_taxes,omitempty"`
	Is_valid_from             string            `json:"is_valid_from,omitempty"`
	Is_valid_to               string            `json:"is_valid_to,omitempty"`
	Contact_address           string            `json:"contact_address,omitempty"`
	Kb_item_status_id         int               `json:"kb_item_status_id,omitempty"`
	Reference                 string            `json:"reference,omitempty"`
	Api_reference             string            `json:"api_reference,omitempty"`
	Viewed_by_client_at       string            `json:"viewed_by_client_at,omitempty"`
	Updated_at                string            `json:"updated_at,omitempty"`
	Esr_id                    int               `json:"esr_id,omitempty"`
	Qr_invoice_id             int               `json:"qr_invoice_id,omitempty"`
	Template_slug             string            `json:"template_slug,omitempty"`
	Taxs                      []Tax             `json:"taxs,omitempty"`
	Network_link              string            `json:"network_link,omitempty"`
	Positions                 []InvoicePosition `json:"positions,omitempty"`
}

type Tax struct {
	Percentage string
	Value      string
}

type InvoicePosition struct {
	Id                  int64  `json:"id,omitempty"`
	Amount              string `json:"amount,omitempty"`
	Unit_id             int64  `json:"unit_id,omitempty"`
	Account_id          int64  `json:"account_id,omitempty"`
	Tax_id              int64  `json:"tax_id,omitempty"`
	Text                string `json:"text,omitempty"`
	Unit_price          string `json:"unit_price,omitempty"`
	Discount_in_percent string `json:"discount_in_percent,omitempty"`
	Article_id          int64  `json:"article_id,omitempty"`
	Type                string `json:"type"`
	Parent_id           int64  `json:"parent_id,omitempty"`
}
