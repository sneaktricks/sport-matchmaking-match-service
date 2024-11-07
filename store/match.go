package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sneaktricks/sport-matchmaking-match-service/dal"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"gorm.io/gorm"
)

var (
	ErrMatchNotFound = errors.New("match not found")
)

type MatchStore interface {
	FindAll(ctx context.Context, page, limit uint, sportFilter []string, startsAfter time.Time) (matches []model.MatchDTO, err error)
	FindByID(ctx context.Context, id uuid.UUID) (match model.MatchDTO, err error)
	Create(ctx context.Context, createData model.MatchCreate, hostUserID uuid.UUID) (match model.MatchDTO, err error)
	Edit(ctx context.Context, id uuid.UUID, editData model.MatchEdit, hostUserID uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID, hostUserID uuid.UUID) error
}

type GormMatchStore struct {
	q *dal.Query
}

func NewGormMatchStore(q *dal.Query) *GormMatchStore {
	return &GormMatchStore{
		q: q,
	}
}

func (ms *GormMatchStore) FindAll(
	ctx context.Context,
	page,
	limit uint,
	sportFilter []string,
	startsAfter time.Time,
) (matches []model.MatchDTO, err error) {
	m := ms.q.Match

	// Build the query
	builder := m.WithContext(ctx)

	// Apply sport filter if not nil
	if sportFilter != nil {
		builder = builder.Where(m.Sport.In(sportFilter...))
	}

	// Filter by start time and sort
	builder = builder.
		Where(m.StartsAt.Gte(startsAfter)).
		Order(m.StartsAt.Asc())

	// Query with pagination
	offset := (page - 1) * limit
	dbMatches, _, err := builder.FindByPage(int(offset), int(limit))
	if err != nil {
		return nil, err
	}

	// Collect to a slice of DTOs
	matches = make([]model.MatchDTO, len(dbMatches))
	for i, match := range dbMatches {
		matches[i] = match.MatchDTO()
	}

	return matches, nil
}

func (ms *GormMatchStore) FindByID(ctx context.Context, id uuid.UUID) (match model.MatchDTO, err error) {
	m := ms.q.Match
	dbMatch, err := m.WithContext(ctx).Where(m.ID.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.MatchDTO{}, ErrMatchNotFound
		}
		return model.MatchDTO{}, err
	}
	match = dbMatch.MatchDTO()

	return match, nil
}

func (ms *GormMatchStore) Create(ctx context.Context, createData model.MatchCreate, hostUserID uuid.UUID) (match model.MatchDTO, err error) {
	m := ms.q.Match

	dbMatch := createData.Match()
	dbMatch.HostUserID = hostUserID

	if err := m.WithContext(ctx).Create(&dbMatch); err != nil {
		return model.MatchDTO{}, nil
	}

	match = dbMatch.MatchDTO()
	return match, nil
}

func (ms *GormMatchStore) Edit(ctx context.Context, id uuid.UUID, editData model.MatchEdit, hostUserID uuid.UUID) error {
	m := ms.q.Match

	dbMatch := editData.Match()
	result, err := m.WithContext(ctx).Where(m.ID.Eq(id), m.HostUserID.Eq(hostUserID)).Updates(dbMatch)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return ErrMatchNotFound
	}

	return nil
}

func (ms *GormMatchStore) Delete(ctx context.Context, id uuid.UUID, hostUserID uuid.UUID) error {
	m := ms.q.Match

	result, err := m.WithContext(ctx).Where(m.ID.Eq(id), m.HostUserID.Eq(hostUserID)).Delete()
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return ErrMatchNotFound
	}

	return nil
}
