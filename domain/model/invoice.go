package model

type Customer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ItemDtl struct {
	Id        int    `gorm:"primary_key" json:"id"`
	ItemName  string `json:"item_name"`
	ItemType  string `json:"item_type"`
	Qty       int    `json:"qty"`
	UnitPrice int    `json:"unit_price"`
	Amount    int    `json:"amount"`
	IdInvoice int    `json:"id_invoice"`
}

type InvoiceTrx struct {
	Id              int    `gorm:"primary_key" json:"id"`
	IssueDate       string `json:"issue_date"`
	Subject         string `json:"subject"`
	TotalItem       int    `json:"total_item"`
	CustomerName    string `json:"customer_name"`
	DueDate         string `json:"due_date"`
	Status          string `json:"status"`
	CustomerAddress string `json:"customer_address"`
	SubTotal        int    `json:"sub_total"`
	Tax             int    `json:"tax"`
	GrandTotal      int    `json:"grand_total"`
	ItemDetail      []*ItemDtl
}

type InvoiceIdx struct {
	Id           int
	IssueDate    string
	Subject      string
	TotalItem    int
	CustomerName string
	DueDate      string
	Status       string
}

type InvoiceIdxReq struct {
	IssueDate    string
	Subject      string
	TotalItem    int
	CustomerName string
	DueDate      string
	Status       string
	Page         int
	Limit        int
}

type InvoiceIdxList struct {
	Data  []*InvoiceIdx
	Page  int
	Limit int
	Count int
}

func (Customer) TableName() string {
	return "customer"
}

func (ItemDtl) TableName() string {
	return "invoice_trx_dtl"
}

func (InvoiceTrx) TableName() string {
	return "invoice_trx"
}
