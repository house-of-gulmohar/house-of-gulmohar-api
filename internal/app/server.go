package app

import (
	"context"
	"database/sql"
	"house-of-gulmohar/internal/api/category"
	categoryQuery "house-of-gulmohar/internal/api/category/query"
	"house-of-gulmohar/internal/api/product"
	productQuery "house-of-gulmohar/internal/api/product/query"
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
	Db        *sql.DB
	Pool      *pgxpool.Pool
}

// Server is the top level recruitment server application object.
type Server struct {
	*Config
	Product  product.ProductHandler
	Category category.CategoryHandler
}

func NewServer(c *Config) *Server {
	s := &Server{
		Config: c,

		Product: product.ProductHandler{
			ProductService: &product.ProductService{
				ProductRepo: &product.ProductDb{Pool: c.Pool, Query: productQuery.ProductQuery{}},
			},
		},
		Category: category.CategoryHandler{
			CategoryService: &category.CategoryService{
				CategoryRepo: &category.CategoryDb{Pool: c.Pool, Query: categoryQuery.CategoryQuery{}},
			},
		},
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
