package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/avran02/effective-mobile/internal/dto"
	"github.com/avran02/effective-mobile/internal/mapper"
	"github.com/avran02/effective-mobile/internal/service"
	"github.com/go-chi/chi/v5"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Controller interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUserData(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)

	GetUserTasks(w http.ResponseWriter, r *http.Request)
	StartUserTask(w http.ResponseWriter, r *http.Request)
	StopUserTask(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.Service
}

func (c *controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	slog.Info("GetUsers controller")
	queryParams := r.URL.Query()
	pageString := queryParams.Get("page")
	pageSizeString := queryParams.Get("pageSize")
	passportNumber := queryParams.Get("passportNumber")
	surname := queryParams.Get("surname")
	name := queryParams.Get("name")
	patronymic := queryParams.Get("patronymic")
	address := queryParams.Get("address")

	filters := []string{passportNumber, surname, name, patronymic, address}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userModels, err := c.service.GetUsers(page, pageSize, filters)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := make([]dto.UserDTO, 0, len(userModels))
	for i, model := range userModels {
		user := mapper.FromUserModelToUserDTO(model)
		users[i] = user
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("CreateUser controller")
	req := dto.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := c.service.CreateUser(req.PassportNumber)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.CreateUserResponse{ID: userID}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *controller) UpdateUserData(w http.ResponseWriter, r *http.Request) {
	slog.Info("UpdateUserData controller")
	userIDString := chi.URLParam(r, "userId")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req dto.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userModel := mapper.FromUserDTOToUserModel(req)
	userModel.ID = userID

	if err := c.service.UpdateUserData(userModel); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("DeleteUser controller")
	userIDString := chi.URLParam(r, "userId")

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteUser(userID); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *controller) GetUserTasks(w http.ResponseWriter, r *http.Request) {
	slog.Info("GetUserTasks controller")
	userIDString := chi.URLParam(r, "userId")

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	startDate := queryParams.Get("startDate")
	endDate := queryParams.Get("endDate")

	tasks := c.service.GetUserTasks(userID, startDate, endDate)

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *controller) StartUserTask(w http.ResponseWriter, r *http.Request) {
	slog.Info("StartUserTask controller")
	userID, taskID, err := parseUserIDAndTaskID(r)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.StartUserTask(userID, taskID); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *controller) StopUserTask(w http.ResponseWriter, r *http.Request) {
	slog.Info("StopUserTask controller")
	userID, taskID, err := parseUserIDAndTaskID(r)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.StopUserTask(userID, taskID); err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func New(s service.Service) Controller {
	return &controller{
		service: s,
	}
}
