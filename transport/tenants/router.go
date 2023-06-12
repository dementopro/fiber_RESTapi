package tenants

import (
	"github.com/gofiber/fiber/v2"
	"github.com/toc/pkg/tenants"
)

type Services struct {
	TenantService tenants.TenantService
}

// InitRoutes initializes the routes for the API
func InitRoutes(router fiber.Router, svc Services) {
	router.Post("/tenant/create", CreateTenantHandler(svc.TenantService))
	router.Get("/tenant/list", ListTenantsHandler(svc.TenantService))
	router.Delete("/tenant/delete", DeleteTenantByIDHandler(svc.TenantService))
	router.Get("/tenant/get/:id", GetTenantByIDHandler(svc.TenantService))
	router.Put("/tenant/update", UpdateTenantByIDHandler(svc.TenantService))
	// router.Post("/tenant/status", StatusTenantHandler(svc.TenantService))
}
