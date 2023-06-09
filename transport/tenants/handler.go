package tenants

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/toc/models"
	"github.com/toc/pkg/tenants"
)

// CreateTenantHandler handles the creation of a new tenant
func CreateTenantHandler(tenantService tenants.TenantService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body
		var createRequest models.Tenant
		err := c.BodyParser(&createRequest)
		// err := json.NewDecoder(r.Body).Decode(&createRequest)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid request payload")
			// http.Error(w, "Invalid request payload", http.StatusBadRequest)
			// return
		}

		// Call the TenantService to create the tenantID
		tenantID, err := tenantService.CreateTenant(createRequest)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
			// utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			// return
		}
		return c.Status(http.StatusOK).JSON(tenantID)
		// utils.RespondWithJSON(w, http.StatusOK, tenantID)
	}
}

// ListTenantsHandler handles the retrieval of all tenants
func ListTenantsHandler(tenantService tenants.TenantService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Call the TenantService to get the list of tenants
		tenantList, err := tenantService.GetAllTenants()
		if err != nil {
			log.Println("Failed to get tenants:", err)
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(http.StatusOK).JSON(tenantList)
	}
}

// DeleteTenantByIDHandler handles the retrieval of all tenants
func DeleteTenantByIDHandler(tenantService tenants.TenantService) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}
