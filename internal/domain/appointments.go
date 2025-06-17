package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Appointment struct {
	bun.BaseModel `bun:"table:appointments,alias:a"`

	Id            int64     `bun:",pk,autoincrement" json:"id"`
	ScheduledTime time.Time `bun:"type:timestamptz,notnull,nullzero," json:"scheduledTime"`
	ServiceId     uuid.UUID `bun:"type:uuid,notnull" json:"serviceId"`
	Service       *Service  `bun:"rel:has-one,join:service_id=id" json:"service"`
	NoShow        bool      `bun:"type:boolean,notnull,nullzero,default:false" json:"noShow"`
	UserId        uuid.UUID `bun:"type:uuid,notnull" json:"userId"`
	// PaidAtBooking
	// PaymentRecieved
	// CreatedAt
	// UpdatedAt
}
