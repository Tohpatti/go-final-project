package repositories

import (
	"database/sql"
	"go-final-project/structs"
)

func GetUserEvent(db *sql.DB) (user_events []structs.UserEvent, err error) {
	rows, err := db.Query("SELECT * FROM user_events")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user_event structs.UserEvent
		err = rows.Scan(
			&user_event.ID,
			&user_event.RegisDate,
			&user_event.UserID,
			&user_event.EventID,
			&user_event.CreatedAt,
			&user_event.UpdatedAt)
		if err != nil {
			panic(err)
		}
		user_events = append(user_events, user_event)
	}

	return user_events, nil
}

func CreateUserEvent(db *sql.DB, user_event structs.UserEvent) (err error) {
	sqlStatement := `INSERT INTO user_events (id, regis_date, user_id, event_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	errs := db.QueryRow(sqlStatement,
		user_event.ID,
		user_event.RegisDate,
		user_event.UserID,
		user_event.EventID,
		user_event.CreatedAt,
		user_event.UpdatedAt)

	return errs.Err()
}

func UpdateUserEvent(db *sql.DB, id int, user_event structs.UserEvent) error {
	_, err := db.Exec("UPDATE user_events SET user_id = $1, event_id = $2, updated_at = $3 WHERE id = $4",
		user_event.UserID, user_event.EventID, user_event.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserEvent(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM user_events WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func FindUserEvent(db *sql.DB, id int) (structs.UserEvent, error) {
	var user_event structs.UserEvent

	err := db.QueryRow("SELECT * FROM user_events WHERE id = $1", id).Scan(
		&user_event.ID,
		&user_event.RegisDate,
		&user_event.UserID,
		&user_event.EventID,
		&user_event.CreatedAt,
		&user_event.UpdatedAt)
	if err != nil {
		return user_event, err
	}

	return user_event, nil
}
