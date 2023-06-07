package models

// Tenant represents the request payload for creating a tenant
type Tenant struct {
	ID           int64        `json:"_id" bson:"_id"`
	Name         string       `json:"name" bson:"name"`
	Admins       []Admin      `json:"admins" bson:"admins"`
	Organization Organization `json:"organization" bson:"organization"`
}

// Admin represents an administrator of a tenant
type Admin struct {
	FirstName string `json:"firstName" bson:"first_name"`
	LastName  string `json:"lastName" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Phone     string `json:"phone" bson:"phone"`
}

// Organization represents the organization details of a tenant
type Organization struct {
	Name           string  `json:"name" bson:"name"`
	PrimaryContact string  `json:"primaryContact" bson:"primary_contact"`
	Address        Address `json:"address" bson:"address"`
}

// Address represents an address
type Address struct {
	Country    string `json:"country" bson:"country"`
	State      string `json:"state" bson:"state"`
	City       string `json:"city" bson:"city"`
	Address1   string `json:"address1" bson:"address1"`
	Address2   string `json:"address2" bson:"address2"`
	PostalCode string `json:"postalCode" bson:"postal_code"`
}
