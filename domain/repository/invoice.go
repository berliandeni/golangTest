package repository

import "golang_test/domain/model"

type InvoiceRepository interface {
	GetCustomerList(param string) ([]*model.Customer, error)
	GetItemList(param string) ([]*model.ItemDtl, error)
	GetInvoiceList(*model.InvoiceIdxReq) (*model.InvoiceIdxList, error)
	SaveInvoiceTrx(*model.InvoiceTrx) (*model.InvoiceTrx, error)
	GetInvoiceById(idInvoice int) (*model.InvoiceTrx, error)
	GetInvoiceDtlById(idInvoice int) ([]*model.ItemDtl, error)
	// SaveInvoiceTrxDtl([]*model.ItemDtl) ([]*model.ItemDtl, error)
	// DeleteItemDtl([]*model.ItemDtl) error
}
