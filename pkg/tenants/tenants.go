package tenants

import (
	"github.com/toc/config"
	"github.com/toc/models"
	repository "github.com/toc/repo/tenants"
	"log"
)

// TenantService provides operations for managing tenants
type TenantService interface {
	CreateTenant(tenant models.Tenant) (*int64, error)
	GetAllTenants() ([]models.Tenant, error)
}

// tenantServiceImpl is the implementation of TenantService
type tenantServiceImpl struct {
	// Add any required dependencies or database connection here
	logger *log.Logger
	repo   repository.TenantRepository
}

// CreateTenant creates a new tenant
func (t *tenantServiceImpl) CreateTenant(tenant models.Tenant) (*int64, error) {
	// Implement the logic to create a tenant
	tenantID, err := t.repo.CreateTenant(&tenant)
	if err != nil {
		t.logger.Println("Failed to update client with error: %s", err.Error())
		return nil, err
	}
	return tenantID, nil
}

// GetAllTenants retrieves all tenants
func (t *tenantServiceImpl) GetAllTenants() ([]models.Tenant, error) {
	// Implement the logic to retrieve all tenants
	return t.repo.GetAllTenants()
}

func NewTenantService(logger *log.Logger, cfg config.Config, repos repository.TenantRepository) TenantService {
	return &tenantServiceImpl{
		logger: logger,
		repo:   repos,
	}
}
