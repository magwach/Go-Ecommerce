package schema

import (
	"time"

	"github.com/google/uuid"
)

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone" gorm:"unique;not null"`
	Password  string    `json:"password"`
	Code      string    `json:"code"`
	Expiry    time.Time `json:"expiry" gorm:"default:null"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:buyer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`

	BankAccount BankAccount `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;default:null"`
	Category    []Category  `gorm:"foreignKey:Owner;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Products    []Product   `gorm:"foreignKey:Owner;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
