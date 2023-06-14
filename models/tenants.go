package models

// Tenant represents the request payload for creating a tenant
type Tenant struct {
	TenantID       string  `json:"tenantId" bson:"tenant_id"`
	Name           string  `json:"name" bson:"name"`
	Type           string  `json:"type" bson:"type"`
	State          string  `json:"state" bson:"state"`
	Address        Address `json:"address" bson:"address"`
	PrimaryContact string  `json:"primaryContact" bson:"primary_contact"`
	Admins         []Admin `json:"admins" bson:"admins"`
	// Licenses       []License    `json:"licenses" bson:"licenses"`
	Networks     []Network    `json:"networks" bson:"networks"`
	Sites        []Site       `json:"sites" bson:"sites"`
	DeviceCounts DeviceCount  `json:"deviceCounts" bson:"device_counts"`
	Organization Organization `json:"organization" bson:"organization"`
	CreatedAt    string       `json:"createdAt" bson:"createdAt"`
	CreatedBy    string       `json:"createdBy" bson:"createdBy"`
	UpdatedAt    string       `json:"updatedAt" bson:"updatedAt"`
	UpdatedBy    string       `json:"updatedBy" bson:"updatedBy"`
	DeletedAt    string       `json:"deletedAt" bson:"deletedAt"`
}

// Admin represents an administrator of a tenant
type Admin struct {
	UserId string `json:"userId" bson:"user_id"`
}

// License represents a license
// type License struct {
// 	LicenseId string `json:"licenseId" bson:"license_id"`
// }

// Network represents a network
type Network struct {
	NetworkId string `json:"networkId" bson:"network_id"`
}

// Site represents a site
type Site struct {
	SiteId string `json:"siteId" bson:"site_id"`
}

// DeviceCount represents a device count
type DeviceCount struct {
	GNB int64 `json:"gNB" bson:"gNB"`
	ENB int64 `json:"eNB" bson:"eNB"`
	CPE int64 `json:"CPE" bson:"CPE"`
}

// Organization represents the organization details of a tenant
type Organization struct {
	Name    string  `json:"name" bson:"name"`
	Address Address `json:"address" bson:"address"`
}

// Address represents an address
type Address struct {
	Country         string `json:"country" bson:"country"`
	StateOrProvince string `json:"stateOrProvince" bson:"state_or_province"`
	City            string `json:"city" bson:"city"`
	Address1        string `json:"address1" bson:"address1"`
	Address2        string `json:"address2" bson:"address2"`
	PostalCode      string `json:"postalCode" bson:"postal_code"`
}
