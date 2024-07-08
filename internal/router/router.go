package router

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/avran02/effective-mobile/internal/controller"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

type Router struct {
	c *controller.Controller
	*chi.Mux
}

func New(c controller.Controller) *Router {
	r := getRoutes(c)

	printRoutes(r)

	return &Router{
		Mux: r,
		c:   &c,
	}
}

func getRoutes(c controller.Controller) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/tasks", c.CreateTask)
	r.Route(viper.GetString("server.apiPathPrefix")+"/users", func(r chi.Router) {
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
