package http

import (
	"fmt"
	"net/http"

	"github.com/audricimanuel/awb-stock-allocation/src/config"
	"github.com/audricimanuel/awb-stock-allocation/src/internals/controller"
	"github.com/audricimanuel/awb-stock-allocation/src/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRouter(
	cfg config.Config,
	awbStockHandler controller.AWBStockController,
	orderHandler controller.OrderController,
	// register new controllers here
) chi.Router {
	r := chi.NewRouter()

	mid := middleware.InitMiddleware(cfg)

	setMiddlewareGlobal(mid, r)

	// Swagger
	r.Group(func(r chi.Router) {
		r.Use(mid.BasicAuth(cfg.SwaggerUsername, cfg.SwaggerPassword))
		r.Route("/swagger", func(r chi.Router) {
			r.Get("/*", httpSwagger.WrapHandler)
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
			})
		})
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		staticText := fmt.Sprintf("hello world: %s", cfg.Env)
		w.Write([]byte(staticText))
	})

	r.Get("/awb-stocks", awbStockHandler.GetAWBStock)
	r.Post("/orders", orderHandler.CreateOrder)
	r.Get("/orders/{id}", orderHandler.GetOrderById)
	r.Put("/orders/{id}/status", orderHandler.UpdateOrderStatus)

	return r
}

func setMiddlewareGlobal(mid middleware.GoMiddleware, r *chi.Mux) {
	// Logger
	r.Use(mid.LogRequest)

	// Cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Recovery
	r.Use(mid.RecoverPanic)
}
