package models

type Fact struct {
	Id          string `json:"id" db:"fact_id"`
	Body        string `json:"body" db:"body"`
	Info        string `json:"info" db:"info"`
	DateAdded   string `json:"date_added" db:"date_added"`
	DateUpdated string `json:"date_updated" db:"date_updated"`
}
