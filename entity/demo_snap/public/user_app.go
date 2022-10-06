package public

import (
	"time"

	"github.com/google/uuid"
	"github.com/randiapr/ottorm/entity"
	"gorm.io/gorm"
)

// orm for demo snap
type UserApp struct {
	entity.BaseEntity
	Username    string
	Password    string
	Dob         string
	LastLoginAt time.Time
}

func (UserApp) TableName() string {
	return "public.user_app"
}

func (e *UserApp) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	if e.CreatedBy == "" {
		e.CreatedBy = "System"
	}
	return
}

func (e *UserApp) BeforeUpdate(tx *gorm.DB) (err error) {
	if e.UpdatedBy == "" {
		e.UpdatedBy = "System"
	}
	return
}
