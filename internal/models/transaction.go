package models

import "time"

type Transaction struct {
	ID          int
	Date        time.Time
	Description string
	Amount      float64
	Category    string
	Account     string
}
