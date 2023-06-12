package tenants

import (
	"github.com/sirupsen/logrus"
	"github.com/toc/config"
	"github.com/toc/models"
	repository "github.com/toc/repo/tenants"
)

// TenantService provides operations for managing tenants
type TenantService interface {
	CreateTenant(tenant models.Tenant) (*string, error)
	GetAllTenants() ([]models.Tenant, error)
	GetTenantByID(tenantID string) (models.Tenant, error)
	DeleteTenantByID(tenantID string) error
	UpdateTenantByID(tenantID string, tenant models.Tenant) error
	// StatusTenant() ([]models.Tenant, error)
}

// tenantServiceImpl is the implementation of TenantService
type tenantServiceImpl struct {
	// Add any required dependencies or database connection here
	logger *logrus.Logger
	repo   repository.TenantRepository
}

func NewTenantService(logger *logrus.Logger, cfg config.Config, repos repository.TenantRepository) TenantService {
	return &tenantServiceImpl{
		logger: logger,
		repo:   repos,
	}
}

// CreateTenant creates a new tenant
func (t *tenantServiceImpl) CreateTenant(tenant models.Tenant) (*string, error) {
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

// DeleteTenentByID delete tenant by ID
func (t *tenantServiceImpl) DeleteTenantByID(tenantID string) error {
	return t.repo.DeleteTenantByID(tenantID)
}

// GetTenantByID  Get Individual Tenant By ID
func (t *tenantServiceImpl) GetTenantByID(tenantID string) (models.Tenant, error) {
	return t.repo.GetTenantByID(tenantID)
}

// UpdateTenant
func (t *tenantServiceImpl) UpdateTenantByID(tenantID string, tenant models.Tenant) error {
	result := t.repo.UpdateTenantByID(tenantID, &tenant)
	return result
}

// // StatusTenant
// func (t *tenantServiceImpl) StatusTenant() ([]models.Tenant, error) {
// 	return t.repo.StatusTenant()
// }
