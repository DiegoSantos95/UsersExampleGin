package store

import (
	"context"

	"github.com/DiegoSantos95/UsersExampleGin/models"
)

type Store interface {
	SetUser(ctx context.Context, user models.User) models.User
	GetUserById(ctx context.Context, userId int64) models.User
}

type storeImpl struct {
	users         map[int64]models.User
	incrementalID int64
}

func New() Store {
	return &storeImpl{
		users:         make(map[int64]models.User),
		incrementalID: 1,
	}
}

func (s *storeImpl) SetUser(ctx context.Context, user models.User) models.User {
	user.ID = s.incrementalID
	s.users[s.incrementalID] = user
	s.incrementalID++

	return user
}

func (s *storeImpl) GetUserById(ctx context.Context, userId int64) models.User {
	return s.users[userId]
}
