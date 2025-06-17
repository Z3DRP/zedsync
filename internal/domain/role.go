package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"tabel:roles,alias:r"`
	Id            int64     `bun:",pk,autoincrement" json:"id"`
	Name          string    `bun:"type:varchar(255),notnull" json:"name"`
	CreatedAt     time.Time `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"createdAt"`
	UpdatedAt     time.Time `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"updatedAt"`
}
