package structs

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Event_Categories struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     string    `json:"end_date"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserEvent struct {
	ID        int       `json:"id"`
	RegisDate time.Time `json:"regis_date"`
	UserID    int       `json:"user_id"`
	EventID   int       `json:"event_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserEventData struct {
	Username  string    `json:"username"`
	RegisDate time.Time `json:"regis_date"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	StartDate time.Time `json:"start_date"`
	EndDate   string    `json:"end_date"`
}
