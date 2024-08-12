package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"template-api-pg/internal/api/middleware"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

func NewServer(ctx context.Context, conn *sql.DB) *http.Server {
	// Routing
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeApplicationJsonMiddleware)
	addRouting(ctx, conn, r)

	// Middleware
	logger := negroni.NewLogger()
	n := negroni.New(
		logger,
		negroni.NewRecovery(),
	)
	if viper.GetBool("api_audit") {
		n.Use(middleware.NewAuditMiddleware(conn))
	}
	n.UseHandler(r)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("api_port")),
		Handler: n,
	}

	return server
}

func addRouting(ctx context.Context, conn *sql.DB, r *mux.Router) {
	r.HandleFunc("/_health", Health)

	// Controllers
	ExampleController(ctx, r.PathPrefix("/example").Subrouter(), conn)
}
