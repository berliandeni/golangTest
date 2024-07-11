package controller

import (
	"fmt"
	"golang_test/domain/model"
	uc "golang_test/usecase/invoice_usecase"
	"strconv"
)

type invoiceController struct {
	invoiceUsecase uc.Invoice
}

type Invoice interface {
	GetCustomerList(ctx Context) ([]*model.Customer, error)
	GetItemList(ctx Context) ([]*model.ItemDtl, error)
	GetInvoiceList(ctx Context) (*model.InvoiceIdxList, error)
	SaveInvoiceTrx(ctx Context) (*model.InvoiceTrx, error)
	GetInvoiceById(ctx Context) (*model.InvoiceTrx, error)
}

func NewInvoiceController(us uc.Invoice) Invoice {
	return &invoiceController{us}
}

func (p *invoiceController) GetCustomerList(ctx Context) ([]*model.Customer, error) {
	resp, err := p.invoiceUsecase.GetCustomerList(ctx.QueryParam("name"))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceController) GetItemList(ctx Context) ([]*model.ItemDtl, error) {
	resp, err := p.invoiceUsecase.GetItemList(ctx.QueryParam("name"))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceController) GetInvoiceList(ctx Context) (*model.InvoiceIdxList, error) {

	totalItem, err := strconv.Atoi(ctx.QueryParam("TotalItem"))
	if err != nil {
		return nil, err
	}
	page, err := strconv.Atoi(ctx.QueryParam("Page"))
	if err != nil {
		return nil, err
	}
	limit, err := strconv.Atoi(ctx.QueryParam("Limit"))
	if err != nil {
		return nil, err
	}

	inv := &model.InvoiceIdxReq{
		IssueDate:    ctx.QueryParam("IssueDate"),
		Subject:      ctx.QueryParam("Subject"),
		TotalItem:    totalItem,
		CustomerName: ctx.QueryParam("CustomerName"),
		DueDate:      ctx.QueryParam("DueDate"),
		Status:       ctx.QueryParam("Status"),
		Page:         page,
		Limit:        limit,
	}

	resp, err := p.invoiceUsecase.GetInvoiceList(inv)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceController) SaveInvoiceTrx(ctx Context) (*model.InvoiceTrx, error) {
	var params *model.InvoiceTrx
	if err := ctx.Bind(&params); err != nil {
		return nil, err
	}

	resp, err := p.invoiceUsecase.SaveInvoiceTrx(params)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceController) GetInvoiceById(ctx Context) (*model.InvoiceTrx, error) {
	fmt.Println(ctx.QueryParam("idInvoice"))
	id, err := strconv.Atoi(ctx.QueryParam("idInvoice"))
	if err != nil {
		return nil, err
	}
	resp, err := p.invoiceUsecase.GetInvoiceById(id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
