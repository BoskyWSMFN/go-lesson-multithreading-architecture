package mvc

import (
	"html/template"
	"net/http"
)

func TasksHandler(w http.ResponseWriter, _ *http.Request) {
	// Пример данных Модели (в реальности они могут быть получены из базы данных или другого источника)
	tasks := []Task{
		{ID: 1, Title: "Купить продукты", Done: false},
		{ID: 2, Title: "Подготовить презентацию", Done: true},
		{ID: 3, Title: "Выполнить задание", Done: false},
	}

	// Загрузка шаблона представления
	tmpl, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	// Отображение шаблона с данными Модели
	err = tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
