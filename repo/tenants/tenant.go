package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/toc/config"
	"github.com/toc/models"
	mongo2 "github.com/toc/pkg/mongo-lib"
	"go.mongodb.org/mongo-driver/bson"
)

type TenantRepository interface {
	CreateTenant(tenant *models.Tenant) (*int64, error)
	GetAllTenants() ([]models.Tenant, error)
	DeleteTenantByID() ([]models.Tenant, error)
	GetTenantByID() ([]models.Tenant, error)
	UpdateTenant() ([]models.Tenant, error)
	StatusTenant() ([]models.Tenant, error)
}

// tenantRepositoryImpl represents the repository for managing tenants in MongoDB
type tenantRepositoryImpl struct {
	config  config.Config
	logger  *log.Logger
	mongoDB *mongo2.MongoDB
}

// NewTenantRepository creates a new instance of tenantRepositoryImpl
func NewTenantRepository(logger *log.Logger, cfg config.Config, mongoDB *mongo2.MongoDB) TenantRepository {
	return &tenantRepositoryImpl{
		mongoDB: mongoDB,
		logger:  logger,
		config:  cfg,
	}
}

// CreateTenant creates a new tenant in MongoDB
func (r *tenantRepositoryImpl) CreateTenant(tenant *models.Tenant) (*int64, error) {
	// TODO: Handle this during request validation
	if tenant.ID == 0 {
		return nil, fmt.Errorf("tenant id is required")
	}

	// Insert into mongo db
	insertOneDetails, err := r.mongoDB.Create(context.Background(), tenant)
	if err != nil {
		r.logger.Println("Failed to create tenant with error: %s", err.Error())
		return nil, err
	}

	// Return document id
	insertedID := insertOneDetails.InsertedID.(int64)
	return &insertedID, nil
}

// GetAllTenants retrieves all tenants from the database
func (r *tenantRepositoryImpl) GetAllTenants() ([]models.Tenant, error) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.config.MongoClient.ConnectTimeout)*time.Second)
	defer cancel()

	// Define an empty filter to retrieve all documents
	filter := bson.D{}

	// Create a slice to store the retrieved tenants
	var responseType models.Tenant
	var results []models.Tenant

	// Perform the find operation
	response, err := r.mongoDB.Read(ctx, filter, &responseType)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return []models.Tenant{}, nil
	}

	// Iterate over the cursor and decode each document into a Tenant struct
	for _, result := range response {
		var tenant *models.Tenant
		var ok bool
		if tenant, ok = result.(*models.Tenant); !ok {
			return nil, fmt.Errorf("failed to fetch tenants")
		}
		results = append(results, *tenant)
	}

	return results, nil
}

// DeleteTenentByID delete tenant by ID
func (r *tenantRepositoryImpl) DeleteTenantByID() ([]models.Tenant, error) {

}

// GetTenantByID  Get Individual Tenant By ID
func (r *tenantRepositoryImpl) GetTenantByID() ([]models.Tenant, error) {

}

// UpdateTenant
func (r *tenantRepositoryImpl) UpdateTenant() ([]models.Tenant, error) {

}

// StatusTenant
func (r *tenantRepositoryImpl) StatusTenant() ([]models.Tenant, error) {

}
