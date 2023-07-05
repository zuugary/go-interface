package ex04

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type Company struct {
	Name    string `json:"name"`
	Created string `json:"created"`
	Product string `json:"product"`
}

type Store interface {
	Get(name string) (*Company, error)
	Insert(c *Company) error
	Update(c *Company) (bool, error)
	Delete(name string) error
}

type SqliteStore struct {
	db *gorm.DB
}

func NewSqliteStore(db *gorm.DB) *SqliteStore {
	return &SqliteStore{db}
}

func (s *SqliteStore) Get(name string) (*Company, error) {
	var company Company

	if err := s.db.Where("name= ?", name).First(&company).Error; err != nil {
		return nil, err
	}

	return &company, nil
}

func (s *SqliteStore) Insert(c *Company) error {
	return s.db.Create(c).Error
}

func (s *SqliteStore) Update(c *Company) (bool, error) {
	if err := s.db.Save(c).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (s *SqliteStore) Delete(name string) error {
	var company Company
	if err := s.db.Where("name = ?", name).Delete(&company).Error; err != nil {
		return err
	}

	return nil
}

type InmemStore struct {
	data map[string]*Company
	mu   *sync.RWMutex
}

func NewInmemStore() *InmemStore {
	return &InmemStore{
		data: make(map[string]*Company),
		mu:   &sync.RWMutex{},
	}
}

func (s *InmemStore) Get(name string) (*Company, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[name]
	if !ok {
		return nil, errors.New("not found")
	}

	return val, nil
}

func (s *InmemStore) Insert(c *Company) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[c.Name]
	if ok {
		return errors.New("duplicate found")
	}

	s.data[c.Name] = c

	return nil
}

func (s *InmemStore) Update(c *Company) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[c.Name]
	if ok {
		return false, errors.New("not found")
	}

	s.data[c.Name] = c

	return true, nil
}

func (s *InmemStore) Delete(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[name]
	if ok {
		return errors.New("not found")
	}

	delete(s.data, name)

	return nil
}
