package app

import (
	"context"
	"house-of-gulmohar/internal/data"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// Config holds the attributes of the Server
// that should appear in the server logs.
type Config struct {
	Name      string
	Version   string
	Port      string
	PgConnStr string
	Db        *pgxpool.Pool
}

// Server is the top level recruitment server application object.
type Server struct {
	*Config

	ProductRepo  data.ProductRepo
	CategoryRepo data.CategoryRepo
	BrandRepo    data.BrandRepo
}

func NewServer(c *Config) *Server {
	s := &Server{
		Config:       c,
		ProductRepo:  &data.ProductDb{Db: c.Db},
		CategoryRepo: &data.CategoryDb{Db: c.Db},
		BrandRepo:    &data.BrandDb{Db: c.Db},
	}
	return s
}

// Shutdown is called by gep, for graceful shutdown.
func (s *Server) Shutdown(c context.Context) {
	logrus.Warn("server shut down signalled")
}

// start server
func (s *Server) Start() {
	logrus.Info("starting server")
	s.Serve(s.Port, s.InitRouter())
}

func (s *Server) Serve(port string, r *chi.Mux) {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	go func() {
		logrus.Infof("listening on port %s", port)
		err := http.ListenAndServe(port, r)
		if err != http.ErrServerClosed {
			logrus.Fatal("failed to start server: ", err)
		}
	}()

	// support graceful shut down, listening for SIGTERM and then calling the
	// Shutdown hook on the service.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	logrus.Info("server shut down gracefully")
}
