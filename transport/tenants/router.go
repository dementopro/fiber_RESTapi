package tenants

import (
	"github.com/gorilla/mux"
	"github.com/toc/pkg/tenants"
	"net/http"
)

type Services struct {
	TenantService tenants.TenantService
}

// InitRoutes initializes the routes for the API
func InitRoutes(router *mux.Router, svc Services) {
	router.HandleFunc("/tenants", CreateTenantHandler(svc.TenantService)).Methods(http.MethodPost)
	router.HandleFunc("/tenants", ListTenantsHandler(svc.TenantService)).Methods(http.MethodGet)
	router.HandleFunc("/tenants/{id}", DeleteTenantByIDHandler(svc.TenantService)).Methods(http.MethodDelete)
}
