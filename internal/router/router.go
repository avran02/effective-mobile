package router

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/controller"
	"github.com/avran02/effective-mobile/internal/middleware"
	"github.com/go-chi/chi/v5"
	swagger "github.com/swaggo/http-swagger"
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
	r.Use(middleware.LoggingMiddleware)

	r.Get("/docs/openapi.yml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/openapi.yml")
	})
	r.Get("/swagger/*", swagger.Handler(
		swagger.URL("/docs/openapi.yml"),
	))
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})

	r.Route(apiPrefix, func(r chi.Router) {
		r.Post("/tasks", c.CreateTask)

		r.Route("/users", func(r chi.Router) {
			r.Get("/", c.GetUsers)
			r.Post("/", c.CreateUser)

			r.Route("/tasks/{taskId}", func(r chi.Router) {
				r.Post("/start", c.StartUserTask)
				r.Post("/stop", c.StopUserTask)
			})

			r.Route("/{userId}", func(r chi.Router) {
				r.Get("/tasks", c.GetUserTasks)
				r.Put("/", c.UpdateUserData)
				r.Delete("/", c.DeleteUser)
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
