package cqrs

import (
	"sync"
	"sync/atomic"
)

func NewUserModel() *UserModel {
	return &UserModel{
		users: make(map[int64]User),
	}
}

// Модель для хранения информации о пользователях
type UserModel struct {
	idCounter atomic.Int64

	mu    sync.RWMutex
	users map[int64]User
}

// Метод для добавления пользователя, который может быть вызван как из текущей рутины, так и из новой
func (m *UserModel) addUser(user User) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.users[user.ID] = user
}

type User struct {
	ID   int64
	Name string
	Age  int
}
