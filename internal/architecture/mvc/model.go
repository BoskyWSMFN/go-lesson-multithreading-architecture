package mvc

type Task struct {
	ID    int
	Title string
	Done  bool
}

// Методы для работы с Моделью

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) UpdateTitle(title string) {
	t.Title = title
}
