package models

// Data filter to get information from DB
type Data struct {
	ID               int64
	FirstName        string
	LastName         string
	Company          string
	City             string
	Country          string
	Phone1           string
	Phone2           string
	Email            string
	SubscriptionDate string
	Website          string

	IsOrEnabled bool
}
