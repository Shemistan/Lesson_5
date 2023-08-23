package service

import (
	"database/sql"
	"fmt"
	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/Shemistan/Lesson_5/internal/storage"
)

type IService interface {
	Add(user *models.User) (int, error)
	//Get(userId int) (*models.User, error)
	//Update(userId int, user *models.User) error
	//Delete(userID int) error
}

func New(repo storage.IStorage) IService {
	return &service{
		CounterAdd: 0,
		repo:       repo,
	}
}

type service struct {
	CounterAdd int
	repo       storage.IStorage
}

func (s *service) Add(user *models.User) (int, error) {
	res, err := s.repo.Add(user)
	if err == sql.ErrNoRows {
		fmt.Println("no rows")
	}

	if err != nil {
		return 0, err
	}

	s.CounterAdd++

	return res, nil
}
