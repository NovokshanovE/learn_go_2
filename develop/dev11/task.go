package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
 1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
 2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
 3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
 4. Реализовать middleware для логирования запросов

Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
 1. Реализовать все методы.
 2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
 3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
 4. Код должен проходить проверки go vet и golint.
*/

// Событие календаря
type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

// Calendar представляет календарь событий.
type Calendar struct {
	Events []*Event
}

var Calender = Calendar{
	Events: make([]*Event, 0),
}

// Шаблон для отправления ошибки
type JSONError struct {
	Error string `json:"error"`
}

// Шаблон для отправления ответа
type JSONResult struct {
	Result string `json:"result"`
}

// Отправляет ответ по http
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// парсит и валидирует параметры методов /create_event и /update_event и /delete_event
func ParseInputDataToEvent(r *http.Request) (*Event, error) {
	userID, err := strconv.Atoi(r.FormValue("user_id"))
	// log.Println(r.Form, " ",r.FormValue("user_id"))
	if err != nil {
		return nil, fmt.Errorf("invalid user_id")
	}

	dateStr := r.FormValue("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format")
	}

	title := r.FormValue("title")

	event := &Event{
		UserID: userID,
		Date:   date,
		Title:  title,
	}

	return event, nil
}

// POST /create_event.
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		JSONResponse(w, http.StatusInternalServerError, JSONError{Error: "Not a post request"})
		return
	}

	event, err := ParseInputDataToEvent(r)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	event.ID = len(Calender.Events)
	Calender.Events = append(Calender.Events, event)

	JSONResponse(w, http.StatusOK, JSONResult{Result: "Event created"})
}

// POST /update_event.
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		JSONResponse(w, http.StatusInternalServerError, JSONError{Error: "Not a post request"})
		return
	}

	event, err := ParseInputDataToEvent(r)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	iter := -1
	for i, val := range Calender.Events {
		if val.UserID == event.UserID && val.Date == event.Date {
			iter = i
			Calender.Events[iter] = event
			break
		}
	}

	if iter == -1 {
		JSONResponse(w, http.StatusServiceUnavailable, JSONError{Error: "No event found"})
		return
	}

	JSONResponse(w, http.StatusOK, JSONResult{Result: "Event updated"})
}

// POST /delete_event.
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		JSONResponse(w, http.StatusInternalServerError, JSONError{Error: "Not a post request"})
		return
	}

	event, err := ParseInputDataToEvent(r)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	iter := -1
	for i, val := range Calender.Events {
		if val.UserID == event.UserID && val.Date == event.Date && val.Title == event.Title {
			iter = i
			break
		}
	}

	if iter == -1 {
		JSONResponse(w, http.StatusServiceUnavailable, JSONError{Error: "No event found"})
		return
	}

	Calender.Events = append(Calender.Events[:iter], Calender.Events[iter+1:]...)
	JSONResponse(w, http.StatusOK, JSONResult{Result: "Event deleted"})
}

// GET /events_for_day.
func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONResponse(w, http.StatusInternalServerError, JSONError{Error: "Not a get request"})
		return
	}

	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	data := make([]*Event, 0)
	for _, val := range Calender.Events {
		if val.Date.Day() == date.Day() {
			data = append(data, val)
		}
	}

	res, err := json.Marshal(data)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, JSONResult{Result: "events for day " + string(res)})
}

// GET /events_for_week.
func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONResponse(w, http.StatusInternalServerError, JSONError{Error: "Not a get request"})
		return
	}

	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	data := make([]*Event, 0)
	for _, val := range Calender.Events {
		_, weekVal := val.Date.ISOWeek()
		_, WeekNeeded := date.ISOWeek()
		if weekVal == WeekNeeded {
			data = append(data, val)
		}
	}

	res, err := json.Marshal(data)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, JSONResult{Result: "Events for week " + string(res)})
}

// GET /events_for_month.
func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONResponse(w, http.StatusInternalServerError, JSONError{Error: "Not a get request"})
		return
	}

	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}

	data := make([]*Event, 0)
	for _, val := range Calender.Events {
		if val.Date.Month() == date.Month() {
			data = append(data, val)
		}
	}

	res, err := json.Marshal(data)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, JSONError{Error: err.Error()})
		return
	}
	JSONResponse(w, http.StatusOK, JSONResult{Result: "Events for month " + string(res)})
}

// Логирование запросов
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Создание маршрутов HTTP
	http.HandleFunc("/create_event", CreateEventHandler)
	http.HandleFunc("/update_event", UpdateEventHandler)
	http.HandleFunc("/delete_event", DeleteEventHandler)
	http.HandleFunc("/events_for_day", EventsForDayHandler)
	http.HandleFunc("/events_for_week", EventsForWeekHandler)
	http.HandleFunc("/events_for_month", EventsForMonthHandler)

	// Создание HTTP-сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      LoggingMiddleware(http.DefaultServeMux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Запуск HTTP-сервера
	log.Println("Starting server on port 8080...")
	log.Fatal(server.ListenAndServe())
}
