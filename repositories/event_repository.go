package repositories

import (
	"database/sql"
	"go-final-project/structs"
)

func GetEvent(db *sql.DB) (events []structs.Event, err error) {
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var event structs.Event
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.StartDate,
			&event.EndDate,
			&event.CategoryID,
			&event.CreatedAt,
			&event.UpdatedAt)
		if err != nil {
			panic(err)
		}
		events = append(events, event)
	}

	return events, nil
}

func CreateEvent(db *sql.DB, event structs.Event) (err error) {
	sqlStatement := `INSERT INTO events (id, title, description, location, start_date, end_date, category_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	errs := db.QueryRow(sqlStatement,
		event.ID,
		event.Title,
		event.Description,
		event.Location,
		event.StartDate,
		event.EndDate,
		event.CategoryID,
		event.CreatedAt,
		event.UpdatedAt)

	return errs.Err()
}

func UpdateEvent(db *sql.DB, id int, event structs.Event) error {
	_, err := db.Exec("UPDATE events SET title = $1, location = $2, start_date = $3, end_date = $4, category_id = $5, updated_at = $6 WHERE id = $7",
		event.Title, event.Location, event.StartDate, event.EndDate, event.CategoryID, event.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM events WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func FindEvent(db *sql.DB, id int) (structs.Event, error) {
	var event structs.Event

	err := db.QueryRow("SELECT * FROM events WHERE id = $1", id).
		Scan(&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.StartDate,
			&event.EndDate,
			&event.CategoryID,
			&event.CreatedAt,
			&event.UpdatedAt)
	if err != nil {
		return event, err
	}

	return event, nil
}
