package db

import (
	"golang_test/domain/model"
	"golang_test/domain/repository"

	"github.com/jinzhu/gorm"
)

type invoiceDb struct {
	db *gorm.DB
}

func NewInvoiceDb(db *gorm.DB) repository.InvoiceRepository {
	return &invoiceDb{db}
}

func (p *invoiceDb) GetCustomerList(param string) ([]*model.Customer, error) {
	var resp []*model.Customer
	if param != "" {
		p.db = p.db.Where("name LIKE ?", "%"+param+"%")
	}
	err := p.db.Find(&resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceDb) GetItemList(param string) ([]*model.ItemDtl, error) {
	var resp []*model.ItemDtl
	if param != "" {
		p.db = p.db.Where("item_name LIKE ?", "%"+param+"%")
	}
	err := p.db.Table("item").Find(&resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *invoiceDb) GetInvoiceList(inv *model.InvoiceIdxReq) (*model.InvoiceIdxList, error) {
	var resp []*model.InvoiceIdx
	db := p.db.Table("invoice_trx")

	if inv.IssueDate != "" {
		db = db.Where("issue_date = ?", inv.IssueDate)
	}
	if inv.Subject != "" {
		db = db.Where("subject LIKE ?", "%"+inv.Subject+"%")
	}
	if inv.TotalItem != 0 {
		db = db.Where("total_item = ?", inv.TotalItem)
	}
	if inv.CustomerName != "" {
		db = db.Where("customer_name LIKE ?", "%"+inv.CustomerName+"%")
	}
	if inv.DueDate != "" {
		db = db.Where("due_date = ?", inv.DueDate)
	}
	if inv.Status != "" {
		db = db.Where("status = ?", inv.Status)
	}

	var count int
	err := db.Count(&count).Error
	if err != nil {
		return nil, err
	}

	err = db.Limit(inv.Limit).Offset(offset(inv.Page, inv.Limit)).Find(&resp).Error
	if err != nil {
		return nil, err
	}

	out := &model.InvoiceIdxList{
		Data:  resp,
		Page:  inv.Page,
		Limit: inv.Limit,
		Count: count,
	}

	return out, nil
}

func (p *invoiceDb) SaveInvoiceTrx(inv *model.InvoiceTrx) (*model.InvoiceTrx, error) {
	tx := p.db.Begin()
	inv.Status = "Unpaid"
	if err := tx.Save(&inv).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err := tx.Table("invoice_trx_dtl").Where("id_invoice = ?", inv.Id).Delete(&model.ItemDtl{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, v := range inv.ItemDetail {
		v.IdInvoice = inv.Id
		err := tx.Save(v).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return inv, nil
}

func (p *invoiceDb) GetInvoiceById(idInvoice int) (*model.InvoiceTrx, error) {
	resp := &model.InvoiceTrx{Id: idInvoice}
	err := p.db.First(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *invoiceDb) GetInvoiceDtlById(idInvoice int) ([]*model.ItemDtl, error) {
	var resp []*model.ItemDtl
	err := p.db.Where("id_invoice = ?", idInvoice).Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func offset(page, limit int) int {
	return (page - 1) * limit
}
