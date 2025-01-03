package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Participation struct {
	UserID  string    `gorm:"primaryKey"`
	MatchID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *Participation) ParticipationDTO() ParticipationDTO {
	return ParticipationDTO{
		UserID:    p.UserID,
		MatchID:   p.MatchID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

type ParticipationDTO struct {
	UserID  string    `json:"userId"`
	MatchID uuid.UUID `json:"matchId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Participations []Participation

func (p Participations) ParticipationDTOs() []ParticipationDTO {
	dtos := make([]ParticipationDTO, len(p))
	for i, participation := range p {
		dtos[i] = participation.ParticipationDTO()
	}

	return dtos
}
