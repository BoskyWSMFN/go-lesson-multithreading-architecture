package cleanarch

func NewAddUserUseCase(userRepository UserRepository) AddUserUseCase {
	return AddUserUseCase{userRepository: userRepository}
}

// Кейс для добавления нового пользователя
type AddUserUseCase struct {
	userRepository UserRepository
}

func (uc *AddUserUseCase) Execute(name string, age int) (int64, error) {
	user := User{Name: name, Age: age}
	return uc.userRepository.Store(user)
}

func NewGetUserUseCase(userRepository UserRepository) GetUserUseCase {
	return GetUserUseCase{userRepository: userRepository}
}

// Кейс для получения информации о пользователе
type GetUserUseCase struct {
	userRepository UserRepository
}

func (uc *GetUserUseCase) Execute(id int64) (User, error) {
	return uc.userRepository.Lookup(id)
}
