package router

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/controller"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	c *controller.Controller
	*chi.Mux
	config *config.Server
}

func New(c controller.Controller, conf *config.Server) *Router {
	r := getRoutes(c, conf.APIPathPrefix)

	printRoutes(r)

	return &Router{
		Mux:    r,
		c:      &c,
		config: conf,
	}
}

func getRoutes(c controller.Controller, apiPrefix string) *chi.Mux {
	r := chi.NewRouter()

	r.Route(apiPrefix, func(r chi.Router) {
		r.Post("/tasks", c.CreateTask)
		r.Route("/users", func(r chi.Router) {
			r.Get("/", c.GetUsers)
			r.Post("/", c.CreateUser)

			r.Route("/{userId}", func(r chi.Router) {
				r.Put("/", c.UpdateUserData)
				r.Delete("/", c.DeleteUser)

				r.Route("/tasks", func(r chi.Router) {
					r.Get("/", c.GetUserTasks)
					r.Post("/{taskId}/start", c.StartUserTask)
					r.Post("/{taskId}/stop", c.StopUserTask)
				})
			})
		})
	})

	return r
}

func printRoutes(router chi.Routes) {
	slog.Info("Serving routes:")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		loggingStr := fmt.Sprintf("Method: %s, Route: %s", method, route)
		slog.Info(loggingStr)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Fatal(err)
	}
}
