package domain

import "github.com/uptrace/bun"

type Address struct {
	bun.BaseModel `bun:"table:addresses,alias=adrs"`
	ID            int64   `bun:"column:id,pk,autoincrement" json:"id"`
	Address       string  `bun:"type:varchar(255),notnull,nullzero" json:"address"`
	City          string  `bun:"type:varchar(150),notnull,nullzero" json:"city"`
	State         string  `bun:"type:varchar(75),notnull,nullzero" json:"state"`
	Country       string  `bun:"type:varchar(75),notnull,nullzero" json:"country"`
	Zipcode       string  `bun:"type:varchar(12),notnull,nullzero" json:"zipcode"`
	Lat           float64 `bun:"type:real,notnull,nullzero" json:"lat"`
	Lng           float64 `bun:"type:real,notnull,nullzero" json:"lng"`
}

func NewAddress(addrs, city, state, country, zip string) *Address {
	return &Address{

		Address: addrs,
		City:    city,
		State:   state,
		Country: country,
		Zipcode: zip,
	}
}
