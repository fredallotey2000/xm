package company

import "context"

type Repository interface {
	CreateCompany(ctx context.Context, comp Company) (string, error)
	UpdateCompany(ctx context.Context, comp Company, compId string) (string, error)
	GetCompany(ctx context.Context, companyId string) (*Company, error)
	DeleteCompany(ctx context.Context, droneId string) (string, error)
}
