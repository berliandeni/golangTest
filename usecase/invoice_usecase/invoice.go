package invoiceusecase

import (
	"golang_test/domain/model"
	"golang_test/domain/repository"
)

type invoiceUsecase struct {
	invoiceRepository repository.InvoiceRepository
}

type Invoice interface {
	GetCustomerList(param string) ([]*model.Customer, error)
	GetItemList(param string) ([]*model.ItemDtl, error)
	GetInvoiceList(*model.InvoiceIdxReq) (*model.InvoiceIdxList, error)
	SaveInvoiceTrx(*model.InvoiceTrx) (*model.InvoiceTrx, error)
	GetInvoiceById(idInvoice int) (*model.InvoiceTrx, error)
}

func NewInvoiceUsecase(p repository.InvoiceRepository) Invoice {
	return &invoiceUsecase{p}
}

func (p *invoiceUsecase) GetCustomerList(param string) ([]*model.Customer, error) {
	resp, err := p.invoiceRepository.GetCustomerList(param)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceUsecase) GetItemList(param string) ([]*model.ItemDtl, error) {
	resp, err := p.invoiceRepository.GetItemList(param)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceUsecase) GetInvoiceList(inv *model.InvoiceIdxReq) (*model.InvoiceIdxList, error) {
	resp, err := p.invoiceRepository.GetInvoiceList(inv)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceUsecase) SaveInvoiceTrx(inv *model.InvoiceTrx) (*model.InvoiceTrx, error) {
	resp, err := p.invoiceRepository.SaveInvoiceTrx(inv)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *invoiceUsecase) GetInvoiceById(idInvoice int) (*model.InvoiceTrx, error) {
	resp, err := p.invoiceRepository.GetInvoiceById(idInvoice)
	if err != nil {
		return nil, err
	}
	respDtl, err := p.invoiceRepository.GetInvoiceDtlById(idInvoice)
	if err != nil {
		return nil, err
	}
	resp.ItemDetail = respDtl
	return resp, nil
}
