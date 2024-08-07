package service

import (
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"time"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/models"
	"github.com/avran02/effective-mobile/internal/repository"
)

type Service interface {
	GetUsers(limit, offset int, filters map[string]string) ([]models.User, error)
	CreateUser(passportNumber string) (int, error)
	UpdateUserData(user models.User) error
	DeleteUser(userID int) error

	CreateTask(task models.Task) (int, error)
	GetUserTasks(userID int, startDate, endDate string) ([]models.TaskTimeSpent, error)
	StartUserTask(taskID int) error
	StopUserTask(taskID int) error
}

type service struct {
	repo            repository.Repository
	externalAPIConf config.ExternalAPI
}

func (s *service) GetUsers(page, pageSize int, filters map[string]string) ([]models.User, error) {
	slog.Info("GetUsers service")
	return s.repo.GetUsers(page, pageSize, filters)
}

func (s *service) CreateUser(passportNumber string) (int, error) {
	slog.Info("CreateUser service")

	passportNumberParts := strings.Split(passportNumber, " ")

	slog.Info("Calling external API")
	user, err := enrichUserData(s.externalAPIConf.EnrichUserDataEndpoint, passportNumberParts[0], passportNumberParts[1])
	if err != nil {
		err = fmt.Errorf("can't create user: %w", err)
		slog.Error(err.Error())
		return 0, err
	}

	slog.Debug("External API response", "user", user)
	user.PassportNumber = passportNumber

	return s.repo.CreateUser(user)
}

func (s *service) UpdateUserData(user models.User) error {
	slog.Info("UpdateUserData service")
	return s.repo.UpdateUserData(user)
}

func (s *service) DeleteUser(userID int) error {
	slog.Info("DeleteUser service")
	return s.repo.DeleteUser(userID)
}

func (s *service) GetUserTasks(userID int, startDate, endDate string) ([]models.TaskTimeSpent, error) {
	slog.Info("GetUserTasks service")

	slog.Debug("Parse dates" + startDate + " " + endDate)
	startTime, err := time.Parse(config.ParceTimeLayout, startDate)
	if err != nil {
		err = fmt.Errorf("can't parse start date: %w", err)
		slog.Error(err.Error())
		return nil, nil
	}

	endTime, err := time.Parse(config.ParceTimeLayout, endDate)
	if err != nil {
		err = fmt.Errorf("can't parse end date: %w", err)
		slog.Error(err.Error())
		return nil, nil
	}

	tasksTimeSpent, err := s.repo.GetUserTasks(userID, startTime, endTime)
	if err != nil {
		err = fmt.Errorf("can't get user tasks from repo: %w", err)
		slog.Error(err.Error())
		return nil, nil
	}

	slog.Info("Sorting tasks")
	sort.Slice(tasksTimeSpent, func(i, j int) bool {
		return tasksTimeSpent[i].TimeSpentSeconds > tasksTimeSpent[j].TimeSpentSeconds
	})

	return tasksTimeSpent, nil
}

func (s *service) CreateTask(task models.Task) (int, error) {
	slog.Info("CreateTask service")
	return s.repo.CreateTask(task)
}

func (s *service) StartUserTask(taskID int) error {
	slog.Info("StartUserTask service")
	return s.repo.StartUserTask(taskID)
}

func (s *service) StopUserTask(taskID int) error {
	slog.Info("StopUserTask service")
	return s.repo.StopUserTask(taskID)
}

func New(repo repository.Repository, conf config.ExternalAPI) Service {
	return &service{
		repo: repo,

		externalAPIConf: conf,
	}
}
