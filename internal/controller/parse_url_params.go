package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func parseUserIDAndTaskID(r *http.Request) (userID, taskID int, err error) {
	userIDString := chi.URLParam(r, "userId")
	userID, err = strconv.Atoi(userIDString)
	if err != nil {
		return 0, 0, err
	}

	taskIDString := chi.URLParam(r, "taskId")
	taskID, err = strconv.Atoi(taskIDString)
	if err != nil {
		return 0, 0, err
	}

	return userID, taskID, nil
}
