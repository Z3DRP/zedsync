package domain

import "github.com/uptrace/bun"

type Address struct {
	bun.BaseModel `bun:"table:addresses,alias=adrs"`
	ID            int64   `bun:"column:id,pk,autoincrement" json:"id"`
	Address       string  `bun:"type:varchar(255),notnull,nullzero" json:"address" validate:"required,alphanumunicode,min=1,max=255"`
	City          string  `bun:"type:varchar(150),notnull,nullzero" json:"city" validate:"required,alpha,min=1,max=150"`
	State         string  `bun:"type:varchar(75),notnull,nullzero" json:"state" validate:"required,alpha,min=1,max=75"`
	Country       string  `bun:"type:varchar(75),notnull,nullzero" json:"country" validate:"required,alpha,min=1,max=75"`
	Zipcode       string  `bun:"type:varchar(12),notnull,nullzero" json:"zipcode" validate:"required,min=12,max=12"`
	Lat           float64 `bun:"type:real,notnull,nullzero" json:"lat" validate:"numeric"`
	Lng           float64 `bun:"type:real,notnull,nullzero" json:"lng" validate:"numeric"`
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
