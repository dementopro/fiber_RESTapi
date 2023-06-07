package tenants

import (
	"encoding/json"
	"github.com/toc/models"
	"github.com/toc/pkg/tenants"
	"github.com/toc/pkg/utils"
	"log"
	"net/http"
)

// CreateTenantHandler handles the creation of a new tenant
func CreateTenantHandler(tenantService tenants.TenantService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var createRequest models.Tenant
		err := json.NewDecoder(r.Body).Decode(&createRequest)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Call the TenantService to create the tenantID
		tenantID, err := tenantService.CreateTenant(createRequest)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, tenantID)
	}
}

// ListTenantsHandler handles the retrieval of all tenants
func ListTenantsHandler(tenantService tenants.TenantService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the TenantService to get the list of tenants
		tenantList, err := tenantService.GetAllTenants()
		if err != nil {
			log.Println("Failed to get tenants:", err)
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, tenantList)
	}
}

// ListTenantsHandler handles the retrieval of all tenants
func DeleteTenantByIDHandler(tenantService tenants.TenantService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
