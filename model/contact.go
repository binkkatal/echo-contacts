package model

// Contact model
type Contact struct {
	ID           int64  `json:"id" form:"id" query:"id"`
	FirstName    string `json:"first_name" form:"first_name" query:"first_name"`
	LastName     string `json:"last_name" form:"last_name" query:"last_name"`
	Organization string `json:"organization" form:"organization" query:"organization"`
	PhoneNumber  string `json:"phone_number" form:"phone_number" query:"phone_number"`
	Email        string `json:"email" form:"email" query:"email"`
	Website      string `json:"website" form:"website" query:"website"`
}
