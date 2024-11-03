package store

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sneaktricks/sport-matchmaking-match-service/dal"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"gorm.io/gorm"
)

var (
	ErrMatchFull             = errors.New("match is already full")
	ErrParticipationNotFound = errors.New("participation not found")
	ErrAlreadyParticipated   = errors.New("user has already participated in the match")
)

type ParticipationStore interface {
	FindAllInMatch(ctx context.Context, matchID uuid.UUID, page, limit uint) (participations []model.ParticipationDTO, err error)
	Create(ctx context.Context, matchID, userID uuid.UUID) (participation model.ParticipationDTO, err error)
	Delete(ctx context.Context, matchID, userID uuid.UUID) error
}

type GormParticipationStore struct {
	q *dal.Query
}

func NewGormParticipationStore(q *dal.Query) *GormParticipationStore {
	return &GormParticipationStore{
		q: q,
	}
}

func (ps *GormParticipationStore) FindAllInMatch(ctx context.Context, matchID uuid.UUID, page uint, limit uint) (participations []model.ParticipationDTO, err error) {
	// Check if match exists
	m := ps.q.Match
	if _, err := m.WithContext(ctx).Where(m.ID.Eq(matchID)).First(); err != nil {
		// Replace with custom error if not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ErrMatchNotFound
		}
		return nil, err
	}

	p := ps.q.Participation

	// Query with pagination
	offset := (page - 1) * limit
	dbParticipations, _, err := p.WithContext(ctx).
		Where(p.MatchID.Eq(matchID)).
		FindByPage(int(offset), int(limit))
	if err != nil {
		return nil, err
	}

	// Collect to a slice of DTOs
	participations = make([]model.ParticipationDTO, len(dbParticipations))
	for i, participation := range dbParticipations {
		participations[i] = participation.ParticipationDTO()
	}

	return participations, nil
}

func (ps *GormParticipationStore) Create(ctx context.Context, matchID uuid.UUID, userID uuid.UUID) (participation model.ParticipationDTO, err error) {
	p := ps.q.Participation

	m := ps.q.Match
	match, err := m.WithContext(ctx).Preload(m.Participations).Where(m.ID.Eq(matchID)).First()
	// Check if match exists
	if err != nil {
		// Replace with custom error if not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ErrMatchNotFound
		}
		return model.ParticipationDTO{}, err
	}

	// Check if the match isn't at participant capacity
	if match.MaxParticipants != nil && len(match.Participations) >= int(*match.MaxParticipants) {
		return model.ParticipationDTO{}, ErrMatchFull
	}

	// Create participation
	dbParticipation := model.Participation{MatchID: matchID, UserID: userID}
	err = p.WithContext(ctx).Create(&dbParticipation)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			err = ErrAlreadyParticipated
		}

		return model.ParticipationDTO{}, err
	}

	return dbParticipation.ParticipationDTO(), nil
}

func (ps *GormParticipationStore) Delete(ctx context.Context, matchID, userID uuid.UUID) error {
	p := ps.q.Participation

	result, err := p.WithContext(ctx).Where(p.MatchID.Eq(matchID), p.UserID.Eq(userID)).Delete()
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return ErrParticipationNotFound
	}

	return nil
}
