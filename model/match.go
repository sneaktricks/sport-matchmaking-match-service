package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Sport             string
	MinParticipants   *int32
	MaxParticipants   *int32
	StartsAt          time.Time
	EndsAt            time.Time
	Location          string
	Description       string
	ParticipationFee  int64
	RequiredEquipment pq.StringArray `gorm:"type:text[]"`
	Level             string
	ChatLink          string
	HostUserID        uuid.UUID
}

type MatchDTO struct {
	ID                uuid.UUID `json:"id"`
	Sport             string    `json:"sport"`
	MinParticipants   *int32    `json:"minParticipants"`
	MaxParticipants   *int32    `json:"maxParticipants"`
	StartsAt          time.Time `json:"startsAt"`
	EndsAt            time.Time `json:"endsAt"`
	Location          string    `json:"location"`
	Description       string    `json:"description"`
	ParticipationFee  int64     `json:"participationFee"`
	RequiredEquipment []string  `json:"requiredEquipment"`
	Level             string    `json:"level"`
	ChatLink          string    `json:"chatLink"`
	HostUserID        uuid.UUID `json:"hostUserId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MatchCreate struct {
	Sport             string    `json:"sport" validate:"required,max=50"`
	MinParticipants   *int32    `json:"minParticipants" validate:"omitnil,min=2"`
	MaxParticipants   *int32    `json:"maxParticipants" validate:"omitnil,min=2"`
	StartsAt          time.Time `json:"startsAt" validate:"required"`
	EndsAt            time.Time `json:"endsAt" validate:"required,gte,gtfield=StartsAt"`
	Location          string    `json:"location" validate:"required,max=100"`
	Description       string    `json:"description" validate:"required,max=1000"`
	ParticipationFee  int64     `json:"participationFee" validate:"required,min=0"`
	RequiredEquipment []string  `json:"requiredEquipment"`
	Level             string    `json:"level" validate:"required,max=100"`
	ChatLink          string    `json:"chatLink" validate:"required,max=200"`
}

type MatchEdit MatchCreate
