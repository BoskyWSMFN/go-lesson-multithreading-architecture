package main

import (
	"fmt"
	"time"

	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/architecture/cqrs"
)

func main() {
	// Создание модели для хранения пользователей
	userModel := cqrs.NewUserModel()

	// Пример использования команды и запроса для добавления пользователя в текущей рутине
	addUserCmd := cqrs.AddUserCommand{Name: "Alice", Age: 30}
	userID := userModel.AddUser(addUserCmd)

	// Поскольку метод для обработки команды добавления пользователя был вызван в текущей рутине,
	// пользователь гарантированно будет найден, т.к. был записан до попытки чтения
	getUserQuery := cqrs.GetUserQuery{UserID: userID}
	user, found := userModel.GetUser(getUserQuery)

	if found {
		fmt.Printf("Пользователь: %+v\n", user)
	} else {
		fmt.Printf("Пользователь c ID %d не найден\n", userID)
	}

	// Пример использования команды и запроса для добавления пользователя с созданием новой рутины
	// и последующей записью в ней.
	addUserCmd = cqrs.AddUserCommand{Name: "Robert", Age: 29, DeferredAdd: true}
	userID = userModel.AddUser(addUserCmd)

	// Вызов метода добавления пользователя с флагом DeferredAdd не гарантирует успешного чтения, т.к.
	// момент записи зависит от планировщика Go.
	getUserQuery = cqrs.GetUserQuery{UserID: userID}
	user, found = userModel.GetUser(getUserQuery)

	if found {
		fmt.Printf("Пользователь: %+v\n", user)
	} else {
		fmt.Printf("Пользователь c ID %d не найден\n", userID)
	}

	time.Sleep(time.Millisecond * 10)

	// Однако исполнение кода в текущей рутине будет продолжено, и пользователя можно попытаться найти позднее.
	getUserQuery = cqrs.GetUserQuery{UserID: userID}
	user, found = userModel.GetUser(getUserQuery)

	if found {
		fmt.Printf("Пользователь: %+v\n", user)
	} else {
		fmt.Printf("Пользователь c ID %d не найден\n", userID)
	}
}
