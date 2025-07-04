package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Company struct {
	bun.BaseModel    `bun:"table:companies,alias:c"`
	ID               uuid.UUID       `bun:",pk,type:uuid,notnull" json:"id" validate:"numeric"`
	Name             string          `bun:"type:varchar(255),notnull," json:"name" validate:"required,alpha,min=1,max=255"`
	OwnerID          int64           `json:"ownerId"`
	HoursOfOperation json.RawMessage `bun:"jsonb,json_use_number" json:"hoursOfOperation"`
	BusinessType     string          `bun:"type:varchar(255),notnull," json:"businessType" validate:"alpha,min=1,max=255"`
	CreatedAt        time.Time       `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"createdAt"`
	UpdatedAt        time.Time       `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"updatedAt"`
}
