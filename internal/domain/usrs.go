// Package domain contains all the domain objects
package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:"column:id,pk,autoincrement" json:"-"`
	UID       uuid.UUID `bun:"type:uuid,notnull,unique" json:"uid"`
	Avatar    string    `bun:"type:varchar(255),null,nullzero" json:"avatar"`
	Email     string    `bun:"type:varchar(150),notnull,unique" json:"email"`
	Phone     string    `bun:"type:varchar(12),notnull" json:"phone"`
	Username  string    `bun:"type:varchar(75),notnull" json:"username"`
	Password  string    `bun:"type:varchar(255),notnull" json:"-"`
	FirstName string    `bun:"type:varchar(100),notnull" json:"firstName"`
	LastName  string    `bun:"type:varchar(100),notnull" json:"lastName"`
	RoleID    int64     `json:"roleId"`
	Role      *Role     `bun:"rel:belongs-to,join:role_id=id" json:"role"`
	// for the has many relationship for appointments
	Appointments []*Appointment `bun:"rel:has-many,join:uuid=user_id" json:"appointments"`
	CreatedAt    time.Time      `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"createdAt"`
	UpdatedAt    time.Time      `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"updatedAt"`
}
