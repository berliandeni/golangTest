package registry

import (
	"golang_test/adapter/controller"
	"golang_test/adapter/db"
	uc "golang_test/usecase/invoice_usecase"
)

func (r *registry) NewInvoiceController() controller.Invoice {
	u := uc.NewInvoiceUsecase(
		db.NewInvoiceDb(r.db),
	)

	return controller.NewInvoiceController(u)
}
