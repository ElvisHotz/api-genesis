package models

import "time"

type Exchanges struct {
	Id        int       `json:"id"`
	Amount    float64   `json:"amount"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Rate      float64   `json:"rate"`
	DtCreated time.Time `json:"dtCreated"`
	Sifra     string    `json:"-"`
	Value     float64   `json:"-"`
}
