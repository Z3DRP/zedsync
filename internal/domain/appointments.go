package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Appointment struct {
	bun.BaseModel `bun:"table:appointments,alias:a"`

	ID            int64     `bun:",pk,autoincrement" json:"id"`
	ScheduledTime time.Time `bun:"type:timestamptz,notnull,nullzero," json:"scheduledTime" validate:"required"`
	ServiceID     uuid.UUID `bun:"type:uuid,notnull" json:"serviceId"`
	Service       *Service  `bun:"rel:has-one,join:service_id=id" json:"service"`
	NoShow        bool      `bun:"type:boolean,notnull,nullzero,default:false" json:"noShow" validate:"boolean"`
	UserID        uuid.UUID `bun:"type:uuid,notnull" json:"userId"`
	Paid          bool      `bun:"type:boolean,notnull,nullzero,default:false" json:"paid" validate:"boolean"`
	// PaidAtBooking
	// PaymentRecieved
	// CreatedAt
	// UpdatedAt
}
