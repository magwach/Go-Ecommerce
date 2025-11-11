package schema

import (
	"time"

	"github.com/google/uuid"
)

type BankAccount struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;unique;not null"`
	User        *User     `json:"-" gorm:"foreignKey:UserID"`
	BankAccount string    `json:"bank_account" gorm:"index;unique;not null"`
	SwiftCode   string    `json:"swift_code" gorm:"index;unique;not null"`
	PaymentType string    `json:"payment_type" gorm:"type:payment_type;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
