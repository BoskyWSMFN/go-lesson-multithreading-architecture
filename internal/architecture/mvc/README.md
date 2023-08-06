### Понятие

MVC (Model-View-Controller) является широко используемым архитектурным подходом, который разделяет приложение на три 
компонента: Модель, Представление и Контроллер. В Go также можно применить этот подход для создания хорошо 
структурированных приложений.

#### Модель (Model)

Модель представляет собой абстракцию данных и бизнес-логики в приложении. В Go Модель может быть представлена в виде 
структур, содержащих данные и методы для работы с ними. Рассмотрим пример структуры Модели для простого приложения 
управления списком задач: `model.go`

#### Представление (View)

Представление отвечает за отображение данных пользователю. В Go представление может быть реализовано с помощью 
HTML-шаблонов или других методов отображения данных. Рассмотрим пример HTML-шаблона для отображения списка 
задач: `view.html`

#### Контроллер (Controller)

Контроллер управляет взаимодействием между Моделью и Представлением. В Go контроллер может быть представлен в виде 
обработчиков HTTP-запросов или функций, которые обрабатывают пользовательский ввод и обновляют Модель. Рассмотрим 
пример контроллера для нашего приложения управления списком задач: `controller.go`   

Каждый из компонентов — Модель, Представление и Контроллер — имеет свою ответственность, что облегчает поддержку и 
расширение кода.