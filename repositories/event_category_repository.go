package repositories

import (
	"database/sql"
	"go-final-project/structs"
)

func GetEventCategory(db *sql.DB) (event_categories []structs.Event_Categories, err error) {
	rows, err := db.Query("SELECT * FROM event_categories")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var event_category structs.Event_Categories
		err = rows.Scan(
			&event_category.ID,
			&event_category.Name,
			&event_category.CreatedAt,
			&event_category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		event_categories = append(event_categories, event_category)
	}

	return event_categories, nil
}

func CreateEventCategory(db *sql.DB, event_category structs.Event_Categories) (err error) {
	sqlStatement := `INSERT INTO event_categories (id, name, created_at, updated_at) 
		VALUES ($1, $2, $3, $4) RETURNING id`

	errs := db.QueryRow(sqlStatement,
		event_category.ID,
		event_category.Name,
		event_category.CreatedAt,
		event_category.UpdatedAt)

	return errs.Err()
}

func UpdateEventCategory(db *sql.DB, id int, event_category structs.Event_Categories) error {
	_, err := db.Exec("UPDATE event_categories SET name = $1, updated_at = $2 WHERE id = $3",
		event_category.Name, event_category.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEventCategory(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM event_categories WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func FindEventCategory(db *sql.DB, id int) (structs.Event_Categories, error) {
	var event_category structs.Event_Categories

	err := db.QueryRow("SELECT * FROM event_categories WHERE id = $1", event_category).
		Scan(&event_category.ID,
			&event_category.Name,
			&event_category.CreatedAt,
			&event_category.UpdatedAt)
	if err != nil {
		return event_category, err
	}

	return event_category, nil
}
