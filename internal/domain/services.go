package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Service struct {
	bun.BaseModel `bun:"table:services,alias:s"`

	Id              uuid.UUID `bun:",pk,type:uuid,notnull," json:"id"`
	Name            string    `bun:"type:varchar(255),notnull,nullzero" json:"name"`
	Description     string    `bun:"type:text,notnull,nullzero" json:"description"`
	Cost            int64     `bun:"type:integer,notnull,nullzero" json:"cost"`
	AverageDuration float64   `bun:"type:numeric(3,2),null,nullzero" json:"averageDuration"`
	Image           string    `bun:"type:varchar(255),null,nullzero" json:"image"`
	CreatedAt       time.Time `bun:"type:timestamptz,notnull,nullzero,default:current_timestamp" json:"createdAt"`
	UpdatedAt       time.Time `bun:"type:timestamptz,notnull,nullzero,default:current_timestamp" json:"updatedAt"`
}
