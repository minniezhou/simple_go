package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/minniezhou/jsonToolBox"
)

type Handler struct {
	router *chi.Mux
}

func (c *Config) NewHandler() *Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http//*", "https://"},
		AllowCredentials: false,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("ping"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := jsonToolBox.WriteJson(w, http.StatusAccepted, jsonToolBox.JsonResponse{Error: false, Message: "Ping GO Service"})
		if err != nil {
			fmt.Println("writing json error")
		}
	})
	r.Post("/submitform", c.HandleForm)
	return &Handler{
		router: r,
	}
}

func (c *Config) HandleForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit handling form function on go service")
}
