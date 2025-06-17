package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Company struct {
	bun.BaseModel    `bun:"table:companies,alias:c"`
	Id               uuid.UUID       `bun:",pk,type:uuid,notnull" json:"id"`
	Name             string          `bun:"type:varchar(255),notnull," json:"name"`
	OwnerId          int64           `json:"ownerId"`
	HoursOfOperation json.RawMessage `bun:"jsonb,json_use_number" json:"hoursOfOperation"`
	BusinessType     string          `bun:"type:varchar(255),notnull," json:"businessType"`
	CreatedAt        time.Time       `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"createdAt"`
	UpdatedAt        time.Time       `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"updatedAt"`
}
