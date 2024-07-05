package repository

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/models"
)

type Repository interface {
	GetUsers(limit, offset int, filterField string) ([]models.User, error)
	GetUserTasks(userID int) []models.Task
	CreateUser(user models.User) error
	UpdateUserData(user models.User) error
	DeleteUser(user models.User) error
	StartUserTask(task models.Task) error
	StopUserTask(task models.Task) error

	Close() error
}

type repository struct {
	*sql.DB
}

func (r *repository) GetUsers(limit, offset int, filterField string) ([]models.User, error) {
	slog.Info("GetUsers repository")
	return nil, nil
}

func (r *repository) GetUserTasks(userID int) []models.Task {
	slog.Info("GetUserTasks repository")
	return nil
}

func (r *repository) CreateUser(user models.User) error {
	slog.Info("CreateUser repository")
	return nil
}

func (r *repository) UpdateUserData(user models.User) error {
	slog.Info("UpdateUserData repository")
	return nil
}

func (r *repository) DeleteUser(user models.User) error {
	slog.Info("DeleteUser repository")
	return nil
}

func (r *repository) StartUserTask(task models.Task) error {
	slog.Info("StartUserTask repository")
	return nil
}

func (r *repository) StopUserTask(task models.Task) error {
	slog.Info("StopUserTask repository")
	return nil
}

func (r *repository) Close() error {
	slog.Info("Close repository")
	return r.DB.Close()
}

func New(conf config.DB) Repository {
	db, err := sql.Open("postgres", getDsn(conf))
	if err != nil {
		log.Fatal(err)
	}
	return &repository{
		DB: db,
	}
}

func getDsn(conf config.DB) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Database,
	)
}
