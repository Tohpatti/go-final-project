package controllers

import (
	"database/sql"
	"go-final-project/databases"
	"go-final-project/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindRegisteredUserInEvent(db *sql.DB, id int) ([]structs.UserEventData, error) {
	var userEvents []structs.UserEventData

	rows, err := db.Query(`
	SELECT
		u.username,
		ue.regis_date,
		e.title,
		e.location,
		e.start_date,
		e.end_date
	FROM
		users u
	JOIN
		user_events ue ON u.id = ue.user_id
	JOIN
		events e ON e.id = ue.event_id
	WHERE
		ue.id = $1	
	`, id)
	if err != nil {
		return userEvents, err
	}
	defer rows.Close()

	for rows.Next() {
		var userEvent structs.UserEventData
		err = rows.Scan(
			&userEvent.Username,
			&userEvent.RegisDate,
			&userEvent.Title,
			&userEvent.Location,
			&userEvent.StartDate,
			&userEvent.EndDate)
		if err != nil {
			return userEvents, err
		}
		userEvents = append(userEvents, userEvent)
	}

	return userEvents, nil
}

func GetRegisteredEvents(ctx *gin.Context) {
	var userEvents []structs.UserEventData
	id, _ := strconv.Atoi(ctx.Param("id"))

	userEvents, err := FindRegisteredUserInEvent(databases.DbConn, id)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "User event not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": userEvents,
	})
}
