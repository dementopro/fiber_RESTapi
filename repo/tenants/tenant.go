package repository

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/toc/config"
	"github.com/toc/models"
	mongo2 "github.com/toc/pkg/mongo-lib"
	"go.mongodb.org/mongo-driver/bson"
)

type TenantRepository interface {
	CreateTenant(tenant *models.Tenant) (*string, error)
	GetAllTenants() ([]models.Tenant, error)
	GetTenantByID(tenantID string) (models.Tenant, error)
	DeleteTenantByID(tenantID string) error
	UpdateTenantByID(tenantID string, tenant *models.Tenant) error
	// StatusTenant() ([]models.Tenant, error)
}

// tenantRepositoryImpl represents the repository for managing tenants in MongoDB
type tenantRepositoryImpl struct {
	config  config.Config
	logger  *logrus.Logger
	mongoDB *mongo2.MongoDB
}

// NewTenantRepository creates a new instance of tenantRepositoryImpl
func NewTenantRepository(logger *logrus.Logger, cfg config.Config, mongoDB *mongo2.MongoDB) TenantRepository {
	return &tenantRepositoryImpl{
		mongoDB: mongoDB,
		logger:  logger,
		config:  cfg,
	}
}

// CreateTenant creates a new tenant in MongoDB
func (r *tenantRepositoryImpl) CreateTenant(tenant *models.Tenant) (*string, error) {
	// TODO: Handle this during request validation

	// Check whether the tenant already exists
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.config.MongoClient.ConnectTimeout)*time.Second)
	defer cancel()

	// Create a slice to store the retrieved tenants
	var responseType models.Tenant
	response, err_read := r.mongoDB.ReadOne(ctx, tenant.TenantID, &responseType)
	if response != nil {
		r.logger.Error("Duplicate tenant")
		return nil, err_read
	}

	// Insert into mongo db

	now := time.Now().Format(time.RFC3339)
	tenant.CreatedAt = now
	tenant.UpdatedAt = now
	tenant.CreatedBy = tenant.Admins[0].UserId
	_, err := r.mongoDB.Create(context.Background(), r.config.MongoClient.Collection, tenant)

	if err != nil {
		err_create := "Failed to create tenant"
		r.logger.Error(err_create)
		return nil, errors.New(err_create)
	}

	// Return document id
	insertedID := tenant.TenantID
	return &insertedID, nil
}

// GetAllTenants retrieves all tenants from the database
func (r *tenantRepositoryImpl) GetAllTenants() ([]models.Tenant, error) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.config.MongoClient.ConnectTimeout)*time.Second)
	defer cancel()

	// Define an empty filter to retrieve all documents
	filter := bson.M{}

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
			err_fetch := "Failed to fetch tenants"
			r.logger.Error(err_fetch)
			return nil, errors.New(err_fetch)
		}
		results = append(results, *tenant)
	}
	return results, nil
}

// DeleteTenentByID delete tenant by ID
func (r *tenantRepositoryImpl) DeleteTenantByID(tenantID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.config.MongoClient.ConnectTimeout)*time.Second)
	defer cancel()

	_, err := r.mongoDB.DeleteOne(ctx, tenantID)

	if err != nil {
		err_del := "Failed to delete tenant"
		r.logger.Error(err_del)
		return errors.New(err_del)
	}
	return nil
}

// GetTenantByID  Get Individual Tenant By ID
func (r *tenantRepositoryImpl) GetTenantByID(tenantID string) (models.Tenant, error) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.config.MongoClient.ConnectTimeout)*time.Second)
	defer cancel()

	// Create a slice to store the retrieved tenants
	var responseType models.Tenant

	// Perform the find operation
	response, _ := r.mongoDB.ReadOne(ctx, tenantID, &responseType)
	if response == nil {
		return models.Tenant{}, nil
	}

	var tenant *models.Tenant
	var ok bool
	if tenant, ok = response.(*models.Tenant); !ok {
		err_fet := "Failed to fetch tenant"
		r.logger.Error("err_fet")
		return models.Tenant{}, errors.New(err_fet)
	}
	return *tenant, nil
}

// UpdateTenant
func (r *tenantRepositoryImpl) UpdateTenantByID(tenantID string, tenant *models.Tenant) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.config.MongoClient.ConnectTimeout)*time.Second)
	defer cancel()

	tenant.UpdatedBy = tenant.Admins[0].UserId

	//Set the updateAt field to the current time
	tenant.UpdatedAt = time.Now().Format(time.RFC3339)

	// Set the createdAt and createdBy fields if they are not already set
	if tenant.CreatedAt == "" {
		tenant.CreatedAt = time.Now().Format(time.RFC3339)
	}
	if tenant.CreatedBy == "" {
		tenant.CreatedBy = tenant.Admins[0].UserId
	}

	// Perform the find operation
	_, err := r.mongoDB.UpdateOne(ctx, tenantID, tenant)
	if err != nil {
		err_update := "Failed to update tenant"
		r.logger.Error(err_update)
		return errors.New(err_update)
	}
	return nil
}

// // StatusTenant
// func (r *tenantRepositoryImpl) StatusTenant() ([]models.Tenant, error) {

// }
