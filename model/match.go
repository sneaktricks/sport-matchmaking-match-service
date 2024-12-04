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
	HostUserID        string
	Participations    []Participation `gorm:"foreignKey:MatchID"`
}

func (m *Match) MatchDTO() MatchDTO {
	return MatchDTO{
		ID:                m.ID,
		Sport:             m.Sport,
		MinParticipants:   m.MinParticipants,
		MaxParticipants:   m.MaxParticipants,
		StartsAt:          m.StartsAt,
		EndsAt:            m.EndsAt,
		Location:          m.Location,
		Description:       m.Description,
		ParticipationFee:  m.ParticipationFee,
		RequiredEquipment: m.RequiredEquipment,
		Level:             m.Level,
		ChatLink:          m.ChatLink,
		HostUserID:        m.HostUserID,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}
}

func (m *Match) MatchWithParticipationsDTO() MatchWithParticipationsDTO {
	return MatchWithParticipationsDTO{
		MatchDTO:       m.MatchDTO(),
		Participations: Participations(m.Participations).ParticipationDTOs(),
	}
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
	HostUserID        string    `json:"hostUserId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MatchWithParticipationsDTO struct {
	MatchDTO
	Participations []ParticipationDTO `json:"participations"`
}

type MatchCreate struct {
	Sport             string    `json:"sport" validate:"required,max=50"`
	MinParticipants   *int32    `json:"minParticipants" validate:"omitnil,min=2"`
	MaxParticipants   *int32    `json:"maxParticipants" validate:"omitnil,min=2"`
	StartsAt          time.Time `json:"startsAt" validate:"required"`
	EndsAt            time.Time `json:"endsAt" validate:"required,gte,gtfield=StartsAt"`
	Location          string    `json:"location" validate:"required,max=100"`
	Description       string    `json:"description" validate:"required,max=1000"`
	ParticipationFee  int64     `json:"participationFee" validate:"min=0"`
	RequiredEquipment []string  `json:"requiredEquipment"`
	Level             string    `json:"level" validate:"required,max=100"`
	ChatLink          string    `json:"chatLink" validate:"max=200"`
}

func (mc *MatchCreate) Match() Match {
	return Match{
		Sport:             mc.Sport,
		MinParticipants:   mc.MinParticipants,
		MaxParticipants:   mc.MaxParticipants,
		StartsAt:          mc.StartsAt,
		EndsAt:            mc.EndsAt,
		Location:          mc.Location,
		Description:       mc.Description,
		ParticipationFee:  mc.ParticipationFee,
		RequiredEquipment: mc.RequiredEquipment,
		Level:             mc.Level,
		ChatLink:          mc.ChatLink,
	}
}

type MatchEdit MatchCreate

func (me *MatchEdit) Match() Match {
	return Match{
		Sport:             me.Sport,
		MinParticipants:   me.MinParticipants,
		MaxParticipants:   me.MaxParticipants,
		StartsAt:          me.StartsAt,
		EndsAt:            me.EndsAt,
		Location:          me.Location,
		Description:       me.Description,
		ParticipationFee:  me.ParticipationFee,
		RequiredEquipment: me.RequiredEquipment,
		Level:             me.Level,
		ChatLink:          me.ChatLink,
	}
}
