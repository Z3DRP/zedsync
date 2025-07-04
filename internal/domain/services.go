package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Service struct {
	bun.BaseModel `bun:"table:services,alias:s"`

	ID              uuid.UUID `bun:",pk,type:uuid,notnull," json:"id"`
	Name            string    `bun:"type:varchar(255),notnull,nullzero" json:"name" validate:"required,alphanum,min=1,max=255"`
	Description     string    `bun:"type:text,notnull,nullzero" json:"description" validate:"required,alphanumunicode"`
	Cost            int64     `bun:"type:integer,notnull,nullzero" json:"cost" validate:"required,numeric"`
	AverageDuration float64   `bun:"type:numeric(3,2),null,nullzero" json:"averageDuration" validate:"numeric"`
	Image           string    `bun:"type:varchar(255),null,nullzero" json:"image" validate:"url"`
	CreatedAt       time.Time `bun:"type:timestamptz,notnull,nullzero,default:current_timestamp" json:"createdAt"`
	UpdatedAt       time.Time `bun:"type:timestamptz,notnull,nullzero,default:current_timestamp" json:"updatedAt"`
}
