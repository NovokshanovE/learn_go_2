package main

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

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Event struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
}

func toJSON(e Event) ([]byte, error) {
	return json.Marshal(e)
}

func parseAndValidateParams(values url.Values) (Event, error) {
	user_id, err := strconv.Atoi(values.Get("user_id"))
	if err != nil {
		return Event{}, err
	}

	title := values.Get("title")
	date := values.Get("date")

	return Event{ID: len(events), UserID: user_id, Title: title, Date: date}, nil
}

// func createEvent(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	event, err := parseAndValidateParams(r.URL.Query())
// 	if err != nil {
// 		http.Error(w, "Invalid parameters", http.StatusBadRequest)
// 		return
// 	}

// 	// Здесь должна быть логика создания события

// 	jsonData, err := toJSON(event)
// 	if err != nil {
// 		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonData)
// }

// Аналогично можно реализовать остальные обработчики

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

var events = make(map[int]Event)

func createEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseAndValidateParams(r.URL.Query())
	// id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	// title := r.URL.Query().Get("title")
	// date := r.URL.Query().Get("date")

	if err != nil {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}
	// event := Event{
	// 	ID:    id,
	// 	Title: title,
	// 	Date:  date,
	// }
	log.Println(event, err)
	events[event.ID] = event

	jsonData, err := toJSON(event)
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	event, err := parseAndValidateParams(r.URL.Query())
	if err != nil {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	_, exists := events[event.ID]
	if !exists {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	events[event.ID] = event

	jsonData, err := toJSON(event)
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, exists := events[id]
	if !exists {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	delete(events, id)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "success"}`))
}

// Для методов events_for_day, events_for_week и events_for_month вам потребуется реализовать логику фильтрации событий по дате.
// Это может быть сложно, так как Go не имеет встроенной поддержки для работы с датами в формате "год-месяц-день".
// Вам потребуется использовать пакет "time" для парсинга дат и сравнения их.

func eventsForDay(w http.ResponseWriter, r *http.Request) {
	requestedDate, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, "Invalid date", http.StatusBadRequest)
		return
	}

	var eventsForDay []Event
	for _, event := range events {
		eventDate, err := time.Parse("2006-01-02", event.Date)
		if err != nil {
			continue // Если дата события не может быть разобрана, пропустите это событие
		}

		if eventDate.Year() == requestedDate.Year() && eventDate.YearDay() == requestedDate.YearDay() {
			eventsForDay = append(eventsForDay, event)
		}
	}

	jsonData, err := json.Marshal(eventsForDay)
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Аналогично можно реализовать функции eventsForWeek и eventsForMonth, сравнивая неделю и месяц года соответственно.

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", createEvent)
	mux.HandleFunc("/update_event", updateEvent)
	mux.HandleFunc("/delete_event", deleteEvent)
	mux.HandleFunc("/events_for_day", eventsForDay)
	// Добавьте обработчики для events_for_week и events_for_month здесь

	// Применение middleware
	loggedRouter := loggingMiddleware(mux)

	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
