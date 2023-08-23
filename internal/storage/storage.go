package storage

import (
	"errors"
	"github.com/Shemistan/Lesson_5/internal/models"
	"log"
)

type IStorage interface {
	Add(user *models.User) (int, error)
	//Get(userId int) (*models.User, error)
	//Update(userId int, user *models.User) error
	//Delete(userID int) error
}

func New(host string, port, ttl int, conn *Conn) IStorage {
	return &storage{
		db:   make(map[int]*models.User),
		ids:  0,
		Host: host,
		Port: port,
		TLL:  ttl,
		conn: conn,
	}
}

type storage struct {
	db   map[int]*models.User
	ids  int
	Host string
	Port int
	TLL  int
	conn *Conn
}

func (s *storage) Add(user *models.User) (int, error) {
	if user == nil {
		return 0, errors.New("user is nil")
	}

	err := s.conn.Open()
	if err != nil {
		return 0, err
	}
	defer func() {
		errClose := s.conn.Close()
		if errClose != nil {
			log.Println(errClose)
		}
	}()

	s.ids++
	s.db[s.ids] = user

	log.Printf("user %v is added: %v", s.ids, user)

	return s.ids, nil
}

func NewConnect() *Conn {
	return &Conn{}
}

type Conn struct {
	val bool
}

func (c *Conn) Close() error {
	if !c.val {
		return errors.New("failed to close")
	}

	return nil
}

func (c *Conn) Open() error {
	if c.val {
		return errors.New("failed to open Conn")
	}

	return nil
}

func (c *Conn) IsClose() bool {
	if c.val {
		return false
	}

	return true
}
