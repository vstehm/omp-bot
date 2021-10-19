package user

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/user"
	"sync"
)

type UserService interface {
	Describe(userID uint64) (*user.User, error)
	List(cursor uint64, limit uint64) ([]user.User, error)
	Create(user.User) (uint64, error)
	Update(userID uint64, user user.User) error
	Remove(userID uint64) (bool, error)
}

type DummyUserService struct {
	lastIndex	uint64
	mu			sync.Mutex
}

func NewDummyUserService() UserService {
	return &DummyUserService{}
}

func (s *DummyUserService) Describe(userID uint64) (*user.User, error) {
	user, ok := allEntities[userID]
	if !ok {
		return nil, fmt.Errorf("USER WITH ID: %d NOT FOUND", userID)
	}
	return &user, nil
}

func (s *DummyUserService) List(cursor uint64, limit uint64) ([]user.User, error) {
	users := make([]user.User, 0, limit)
	for i := cursor; i <= s.lastIndex; i++ {
		if len(users) >= int(limit) {
			break
		}

		user, ok := allEntities[i]
		if !ok {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *DummyUserService) Create(user user.User) (uint64, error) {
	s.mu.Lock()
	s.lastIndex++
	id := s.lastIndex
	s.mu.Unlock()

	user.Id = id
	allEntities[user.Id] = user

	return id, nil
}
func (s *DummyUserService) Update(userID uint64, user user.User) error {
	user.Id = userID
	_, ok := allEntities[userID]
	if !ok {
		return fmt.Errorf("USER WITH ID: %d NOT FOUND", userID)
	}
	allEntities[userID] = user
	return nil
}
func (s *DummyUserService) Remove(userID uint64) (bool, error) {
	_, ok := allEntities[userID]
	if !ok {
		return false, fmt.Errorf("USER WITH ID: %d NOT FOUND", userID)
	}
	delete(allEntities, userID)
	return true, nil
}
