package repository

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/mapper"
	"github.com/avran02/effective-mobile/internal/models"
)

type Repository interface {
	GetUsers(page, pageSize int, filters map[string]string) ([]models.User, error)
	CreateUser(user *models.User) (int, error)
	UpdateUserData(user models.User) error
	DeleteUser(id int) error

	CreateTask(task models.Task) (int, error)
	GetUserTasks(userID int, startDate, endDate time.Time) ([]models.TaskTimeSpent, error)
	StartUserTask(taskID int) error
	StopUserTask(taskID int) error

	Close() error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetUsers(page, pageSize int, filters map[string]string) ([]models.User, error) {
	slog.Info("GetUsers repository")

	builder := strings.Builder{}
	builder.WriteString("SELECT * FROM users")

	var params []interface{}
	if len(filters) > 0 {
		builder.WriteString(" WHERE ")
		i := 1
		for filterName, filter := range filters {
			if i > 1 {
				builder.WriteString(" AND ")
			}
			builder.WriteString(fmt.Sprintf("%s = $%d", filterName, i))
			params = append(params, filter)
			i++
		}
	}

	builder.WriteString(" LIMIT $")
	builder.WriteString(fmt.Sprint(len(params) + 1))
	builder.WriteString(" OFFSET $")
	builder.WriteString(fmt.Sprint(len(params) + 2)) //nolint:mnd

	query := builder.String()
	slog.Info(query)

	limit := pageSize
	offset := (page - 1) * limit

	params = append(params, limit, offset)

	slog.Info(fmt.Sprintf("limit: %d, offset: %d", limit, offset))

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *repository) CreateUser(user *models.User) (int, error) {
	slog.Info("CreateUser repository")

	var id int
	query := `
	INSERT INTO users (
		passport_number, name, surname, patronymic, address
	) 
	VALUES 
		($1, $2, $3, $4, $5) 
	RETURNING id
	`

	if err := r.db.QueryRow(query, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) UpdateUserData(user models.User) error {
	slog.Info("UpdateUserData repository")

	builder := strings.Builder{}
	builder.WriteString("UPDATE users SET ")

	newUserData := mapper.ToDatabaseUpdateUserData(user)
	params := make([]interface{}, 0)
	i := 1

	for key, value := range newUserData {
		if len(params) > 0 {
			builder.WriteString(", ")
		}

		builder.WriteString(fmt.Sprintf("%s = $%d", key, i))
		params = append(params, value)
		i++
	}

	builder.WriteString(" WHERE id = $")
	builder.WriteString(fmt.Sprint(i))
	builder.WriteString(";")
	params = append(params, user.ID)

	query := builder.String()

	slog.Info(fmt.Sprintf("Update query: %s", query))
	slog.Info(fmt.Sprintf("Update params: %+v", params))

	_, err := r.db.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteUser(id int) error {
	slog.Info("DeleteUser repository")

	query := "DELETE FROM users WHERE id = $1"

	_, err := r.db.Exec(query, id)
	return err
}

func (r *repository) CreateTask(task models.Task) (int, error) {
	slog.Info("CreateTask repository")
	var id int

	query := `
	INSERT INTO tasks 
		(name, description, user_id) 
	VALUES 
		($1, $2, $3) 
	RETURNING id
	`

	err := r.db.QueryRow(query, task.Name, task.Description, task.UserID).Scan(&id)
	if err != nil {
		slog.Error(err.Error())
		return 0, err
	}
	return id, err
}

func (r *repository) GetUserTasks(userID int, startDate, endDate time.Time) ([]models.TaskTimeSpent, error) {
	slog.Info("GetUserTasks repository")
	// TODO: fix it
	query := `
	SELECT 
	    tasks.id,
	    tasks.name,
	    tasks.description,
	    SUM(EXTRACT(EPOCH FROM (time_logs.end_time - time_logs.start_time))) as total_time_seconds
	FROM 
	    tasks
	JOIN 
	    time_logs ON tasks.id = time_logs.task_id
	WHERE 
	    tasks.user_id = $1 AND
	    time_logs.start_time >= $2 AND 
	    time_logs.end_time <= $3
	GROUP BY 
	    tasks.id, tasks.name, tasks.description;
	`

	rows, err := r.db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taskTimeSpents []models.TaskTimeSpent
	for rows.Next() {
		var taskTimeSpent models.TaskTimeSpent
		if err := rows.Scan(&taskTimeSpent.TaskID, &taskTimeSpent.Name, &taskTimeSpent.Description, &taskTimeSpent.TimeSpentSeconds); err != nil {
			return nil, err
		}
		taskTimeSpents = append(taskTimeSpents, taskTimeSpent)
	}

	return taskTimeSpents, nil
}

func (r *repository) StartUserTask(taskID int) error {
	slog.Info("StartUserTask repository")

	query := "INSERT INTO time_logs (task_id, start_time) VALUES ($1, $2)"

	_, err := r.db.Exec(query, taskID, time.Now())
	return err
}

func (r *repository) StopUserTask(taskID int) error {
	slog.Info("StopUserTask repository")

	query := "UPDATE time_logs SET end_time = $1 WHERE task_id = $2 AND end_time IS NULL"
	_, err := r.db.Exec(query, time.Now(), taskID)
	return err
}

func (r *repository) Close() error {
	slog.Info("Close repository")
	return r.db.Close()
}

func New(conf config.DB) Repository {
	slog.Info(getDsn(conf))
	db, err := sql.Open("postgres", getDsn(conf))
	if err != nil {
		log.Fatal("can't open db conn:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("can't ping:", err)
	}

	return &repository{
		db: db,
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
