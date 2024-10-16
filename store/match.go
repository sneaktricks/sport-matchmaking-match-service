package store

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sneaktricks/sport-matchmaking-match-service/dal"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
)

var (
	ErrMatchNotFound = errors.New("match not found")
)

type MatchStore interface {
	FindAll(ctx context.Context, page, limit uint) (matches []model.MatchDTO, err error)
	FindAllBySport(ctx context.Context, sport []string, page, limit uint) (matches []model.MatchDTO, err error)
	FindByID(ctx context.Context, id uuid.UUID) (match model.MatchDTO, err error)
	Create(ctx context.Context, createData model.MatchCreate, hostUserID uuid.UUID) (match model.MatchDTO, err error)
	Edit(ctx context.Context, id uuid.UUID, editData model.MatchEdit) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type GormMatchStore struct {
	q *dal.Query
}

func NewGormMatchStore(q *dal.Query) *GormMatchStore {
	return &GormMatchStore{
		q: q,
	}
}

func (ms *GormMatchStore) FindAll(ctx context.Context, page uint, limit uint) (matches []model.MatchDTO, err error) {
	panic("not implemented") // TODO: Implement
}

func (ms *GormMatchStore) FindAllBySport(ctx context.Context, sport []string, page uint, limit uint) (matches []model.MatchDTO, err error) {
	panic("not implemented") // TODO: Implement
}

func (ms *GormMatchStore) FindByID(ctx context.Context, id uuid.UUID) (match model.MatchDTO, err error) {
	panic("not implemented") // TODO: Implement
}

func (ms *GormMatchStore) Create(ctx context.Context, createData model.MatchCreate, hostUserID uuid.UUID) (match model.MatchDTO, err error) {
	panic("not implemented") // TODO: Implement
}

func (ms *GormMatchStore) Edit(ctx context.Context, id uuid.UUID, editData model.MatchEdit) error {
	panic("not implemented") // TODO: Implement
}

func (ms *GormMatchStore) Delete(ctx context.Context, id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
