package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Participation struct {
	UserID  uuid.UUID `gorm:"type:uuid;primaryKey"`
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
	UserID  uuid.UUID `json:"userId"`
	MatchID uuid.UUID `json:"matchId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
