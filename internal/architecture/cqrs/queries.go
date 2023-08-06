package cqrs

// Команда для добавления нового пользователя
type AddUserCommand struct {
	Name string
	Age  int

	// При положительном значении производит операцию записи в новой рутине.
	DeferredAdd bool
}

// Запрос для получения информации о пользователе
type GetUserQuery struct {
	UserID int64
}
