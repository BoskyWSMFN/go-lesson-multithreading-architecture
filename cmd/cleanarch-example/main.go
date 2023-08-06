package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/architecture/cleanarch"
	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/architecture/cleanarch/adapters/inmemory"
)

func main() {
	// Создание хранилища
	storage := inmemory.NewInMemoryStorage()

	// Проверка на соблюдение контракта интерфейса во время компиляции
	userRepo := cleanarch.UserRepository(storage)

	// Запись в память
	addUser := cleanarch.NewAddUserUseCase(userRepo)

	userID, err := addUser.Execute("Alice", 30)
	if err != nil {
		log.Fatal(err)
	}

	// Чтение из памяти
	getUser := cleanarch.NewGetUserUseCase(userRepo)

	user, err := getUser.Execute(userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Пользователь: %+v\n", user)

	// ----------------------

	// Создание хранилища с отложенной записью в память
	storageDeferred := inmemory.NewInMemoryStorageDeferredFromInMemoryStorage(storage)

	// Проверка на соблюдение контракта интерфейса во время компиляции
	userRepo = cleanarch.UserRepository(storageDeferred)

	// Отложенная запись в память
	addUser = cleanarch.NewAddUserUseCase(userRepo)

	userID, err = addUser.Execute("Robert", 29)
	if err != nil {
		log.Fatal(err)
	}

	// Чтение из памяти
	getUser = cleanarch.NewGetUserUseCase(userRepo)

	user, err = getUser.Execute(userID)
	if err != nil {
		fmt.Println(err) // Ошибка, так как планировщик еще не запустил рутину с записью
	} else {
		fmt.Printf("Пользователь: %+v\n", user)
	}

	time.Sleep(time.Millisecond * 10)

	// Повторное чтение из памяти
	user, err = getUser.Execute(userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Пользователь: %+v\n", user)
}
