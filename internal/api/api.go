package api

import (
	"net/http"

	"github.com/Dhairya3124/ReaderCLI/internal/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type Application struct {
	storage store.Store
	config  Config
	logger  *zap.SugaredLogger
}
type Config struct {
	Addr string
}

func NewServer(store store.Store, config Config, logger *zap.SugaredLogger) *Application {
	return &Application{
		storage: store,
		config:  config,
		logger:  logger,
	}

}
func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	// r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Route("/articles", func(r chi.Router) {
			r.Post("/", app.createArticleHandler)
		})

	})

	return r
}
func (app *Application) Run(mux http.Handler) error {

	srv := &http.Server{
		Addr:    app.config.Addr,
		Handler: mux,
	}
	app.logger.Infow("server has started", "addr", app.config.Addr)
	return srv.ListenAndServe()
}
