package company

import (
	"context"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

type Service interface {
	AddCompany(ctx context.Context, comp Company) (string, error)
	ModifyCompany(ctx context.Context, comp Company, id string) (string, error)
	GetCompany(ctx context.Context, companyId string) (*Company, error)
	RemoveCompany(ctx context.Context, companyId string) (string, error)
}

// NewService creates a new DNS service with an instance of a repository
func NewCompanyService(r Repository) Service {
	o := &service{
		repo: r,
	}
	return o
}

// CalculatePath gets the data bank location for a repository
func (s *service) AddCompany(ctx context.Context, comp Company) (string, error) {
	comp.ID = uuid.New().String()

	compId, err := s.repo.CreateCompany(ctx, comp)
	if err != nil {
		return "", err
	}
	return compId, nil

}

func (s *service) ModifyCompany(ctx context.Context, comp Company, id string) (string, error) {
	compId, err := s.repo.UpdateCompany(ctx, comp, id)
	if err != nil {
		return "", err
	}
	return compId, nil

}

func (s *service) GetCompany(ctx context.Context, companyId string) (*Company, error) {
	comp, err := s.repo.GetCompany(ctx, companyId)
	if err != nil {
		return nil, err
	}
	return comp, nil

}

func (s *service) RemoveCompany(ctx context.Context, companyId string) (string, error) {
	compId, err := s.repo.DeleteCompany(ctx, companyId)
	if err != nil {
		return "", err
	}
	return compId, nil

}
