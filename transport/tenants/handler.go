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
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid request payload")
		}

		// Call the TenantService to create the tenantID
		tenantID, err := tenantService.CreateTenant(createRequest)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
			// return
		}
		return c.Status(http.StatusOK).JSON(tenantID)
	}
}

// GetTenantByIDHandler handles the retrieval of a specific tenant
func GetTenantByIDHandler(tenantService tenants.TenantService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the request body
		tenantID := c.Params("id")
		// err := json.NewDecoder(r.Body).Decode(&createRequest)

		// Call the TenantService to create the tenantID
		tenant, err := tenantService.GetTenantByID(tenantID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
			// utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			// return
		}
		return c.Status(http.StatusOK).JSON(tenant)
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
		// Call the TennatService to to delete the tenant by ID.
		type DeleteRequest struct {
			ID string `json:"id"`
		}
		var deleteRequest DeleteRequest
		err := c.BodyParser(&deleteRequest)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid request payload")
		}
		err1 := tenantService.DeleteTenantByID(deleteRequest.ID)
		if err1 != nil {
			log.Println("Failed to get tenants:", err1)
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(http.StatusOK).SendString("success")
	}
}

// DeleteTenantByIDHandler handles the retrieval of all tenants
func UpdateTenantByIDHandler(tenantService tenants.TenantService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Call the TennatService to to delete the tenant by ID.
		var updateRequest models.Tenant
		err := c.BodyParser(&updateRequest)

		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid request payload")
		}
		err1 := tenantService.UpdateTenantByID(updateRequest.TenantID, updateRequest)
		if err1 != nil {
			log.Println("Failed to update tenants:", err1)
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return c.Status(http.StatusOK).SendString("success")
	}
}
