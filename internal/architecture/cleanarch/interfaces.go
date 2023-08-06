package cleanarch

// Интерфейс для хранилища пользователей
//
//go:generate moq -pkg mocks -out ../../mocks/userrepository.go -stub -rm -skip-ensure . UserRepository:UserRepositoryMock
type UserRepository interface {
	Store(user User) (int64, error)
	Lookup(id int64) (User, error)
}
