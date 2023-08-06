package inmemory

import (
	"fmt"

	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/architecture/cleanarch"
	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/architecture/cqrs"
)

func NewInMemoryStorage() InMemoryStorage {
	return InMemoryStorage{userModel: cqrs.NewUserModel()}
}

// Адаптер для записи пользвателя в память
type InMemoryStorage struct {
	userModel *cqrs.UserModel
}

func (i InMemoryStorage) Store(user cleanarch.User) (int64, error) {
	cmdAddUser := cqrs.AddUserCommand{
		Name: user.Name,
		Age:  user.Age,
	}

	userID := i.userModel.AddUser(cmdAddUser)

	return userID, nil
}

func (i InMemoryStorage) Lookup(id int64) (cleanarch.User, error) {
	cmdGetUser := cqrs.GetUserQuery{UserID: id}

	user, found := i.userModel.GetUser(cmdGetUser)
	if !found {
		return cleanarch.User(user), fmt.Errorf("user with id %d not found", id)
	}

	return cleanarch.User(user), nil
}

func NewInMemoryStorageDeferredFromInMemoryStorage(inMemoryStorage InMemoryStorage) InMemoryStorageDeferred {
	return InMemoryStorageDeferred{InMemoryStorage: inMemoryStorage}
}

// Адаптер для отложенной записи пользвателя в память с затеняемым методом Store.
type InMemoryStorageDeferred struct {
	InMemoryStorage
}

func (i InMemoryStorageDeferred) Store(user cleanarch.User) (int64, error) {
	cmdAddUser := cqrs.AddUserCommand{
		Name:        user.Name,
		Age:         user.Age,
		DeferredAdd: true,
	}

	userID := i.userModel.AddUser(cmdAddUser)

	return userID, nil
}
