package memory

import (
	"context"
	"sync"

	"cavea-movie.ge/metadata/internal/repository"
	"cavea-movie.ge/metadata/pkg/model"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.MetaData
}

func New() *Repository {
	return &Repository{data: map[string]*model.MetaData{}}
}

func (r *Repository) Get(_ context.Context, id string) (*model.MetaData, error) {
	r.Lock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

func (r *Repository) Put(_ context.Context, id string, metadata *model.MetaData) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
