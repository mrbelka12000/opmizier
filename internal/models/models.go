package models

// Request filter to get information from DB
type Request struct {
	ID               int64  `schema:"id,omitempty"`
	FirstName        string `schema:"first_name,omitempty"`
	LastName         string `schema:"last_name,omitempty"`
	Company          string `schema:"company,omitempty"`
	City             string `schema:"city,omitempty"`
	Country          string `schema:"country,omitempty"`
	Phone1           string `schema:"phone_1,omitempty"`
	Phone2           string `schema:"phone_2,omitempty"`
	Email            string `schema:"email,omitempty"`
	SubscriptionDate string `schema:"subscription_date,omitempty"`
	Website          string `schema:"website,omitempty"`

	IsOrEnabled bool `schema:"is_or_enabled,omitempty"`
}
