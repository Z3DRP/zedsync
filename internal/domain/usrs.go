// Package domain contains all the domain objects
package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	// TODO: check that making json for avatar be avatarUrl doesnt' cause problems
	ID         int64     `bun:"column:id,pk,autoincrement" json:"-"`
	UID        uuid.UUID `bun:"type:uuid,notnull,unique" json:"uid" validate:"uuid4"`
	Avatar     string    `bun:"type:varchar(255),null,nullzero" json:"avatarUrl" validate:"url"`
	Email      string    `bun:"type:varchar(150),notnull,unique" json:"email" validate:"asci"`
	Phone      string    `bun:"type:varchar(12),notnull" json:"phone" validate:"numeric"`
	Username   string    `bun:"type:varchar(75),notnull" json:"username" validate:"min=1,max=75"`
	Password   string    `bun:"type:varchar(255),notnull" json:"-" validate:"min=1,max=255"`
	FirstName  string    `bun:"type:varchar(100),notnull" json:"firstName" validate:"alpha,min=1,max=150"`
	LastName   string    `bun:"type:varchar(100),notnull" json:"lastName" validate:"alpha,min=1,max=150"`
	RoleID     int64     `json:"roleId"`
	Role       *Role     `bun:"rel:belongs-to,join:role_id=id" json:"role"`
	Status     string    `bun:"type:varchar(35),notnull,nullzero" json:"status" validate:"alph"`
	IsVerified bool      `bun:",notnull,nullzero,default:false" json:"isVerified" validate:"boolean"`
	// for the has many relationship for appointments
	Appointments []*Appointment `bun:"rel:has-many,join:uuid=user_id" json:"appointments"`
	Address      *Address       `bun:"rel:has-one,join:uid=userId" json:"address"`
	CreatedAt    time.Time      `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"createdAt"`
	UpdatedAt    time.Time      `bun:"type:timestamptz,notnull,nullzero,default=current_timestamp" json:"updatedAt"`
}

func NewUser(uid, avatar, email, phne, usrname, fname, lname string, role Role) (*User, error) {
	UID, err := uuid.Parse(uid)
	if err != nil {
		return nil, err
	}
	return &User{
		UID:       UID,
		Avatar:    avatar,
		Email:     email,
		Phone:     phne,
		Username:  usrname,
		FirstName: fname,
		LastName:  lname,
		RoleID:    role.ID,
	}, nil
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		u.UpdatedAt = time.Now()
	}
	return nil
}

func (u User) Info() string {
	return fmt.Sprintf("%#v\n", u)
}
