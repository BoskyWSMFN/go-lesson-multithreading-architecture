package cqrs

// Метод для обработки команды добавления пользователя
func (m *UserModel) AddUser(cmd AddUserCommand) (userID int64) {
	const idDelta = 1

	userID = m.idCounter.Add(idDelta)
	user := User{
		ID:   userID,
		Name: cmd.Name,
		Age:  cmd.Age,
	}

	if !cmd.DeferredAdd {
		// вызов метода в текущей рутине
		m.addUser(user)

		return
	}

	// создание новой рутины и вызов метода в ней
	go m.addUser(user)

	return
}

// Метод для выполнения запроса на получение информации о пользователе
func (m *UserModel) GetUser(query GetUserQuery) (user User, found bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, found = m.users[query.UserID]

	return
}
