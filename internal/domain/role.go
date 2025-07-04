package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"tabel:roles,alias:r"`
	ID            int64     `bun:",pk,autoincrement" json:"id"`
	Name          string    `bun:"type:varchar(255),notnull,unique" json:"name" validate:"required,alpha,min=1,max=255"`
	CreatedAt     time.Time `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"createdAt"`
	UpdatedAt     time.Time `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"updatedAt"`
}
