package models

type Phone struct {
	Number string `json:"number" db:"msisdn"`
}

type Phones []Phone
