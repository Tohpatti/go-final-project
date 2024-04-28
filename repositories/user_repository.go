package repositories

import (
	"database/sql"
	"go-final-project/structs"
)

func GetUser(db *sql.DB) (users []structs.User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.User
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(db *sql.DB, user structs.User) (err error) {
	sqlStatement := `INSERT INTO users (id, username, email, password, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	errs := db.QueryRow(sqlStatement,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt)

	return errs.Err()
}

func UpdateUser(db *sql.DB, id int, user structs.User) error {
	_, err := db.Exec("UPDATE users SET username = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5",
		user.Username, user.Email, user.Password, user.UpdatedAt, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func FindUser(db *sql.DB, username string) (structs.User, error) {
	var user structs.User

	err := db.QueryRow("SELECT * FROM users WHERE username = $1", username).
		Scan(&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
