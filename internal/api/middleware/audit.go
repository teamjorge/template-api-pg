package middleware

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"
	"template-api-pg/internal/models"
	"time"

	"github.com/urfave/negroni"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type AuditMiddleware struct {
	conn *sql.DB
}

func NewAuditMiddleware(conn *sql.DB) *AuditMiddleware {
	return &AuditMiddleware{conn: conn}
}

func (a *AuditMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)
	res := rw.(negroni.ResponseWriter)

	audit := models.APIAudit{
		TS:          time.Now().UnixMicro(),
		IPAddress:   cleanRemoteAddr(r.RemoteAddr),
		Method:      r.Method,
		RequestPath: r.RequestURI,
		Status:      res.Status(),
		UserAgent:   null.StringFrom(r.UserAgent()),
	}

	if err := audit.Insert(context.Background(), a.conn, boil.Infer()); err != nil {
		log.Printf("failed to create audit log - %v\n", err)
	}
}

// cleanRemoteAddr will remove the port from an addr:port string
func cleanRemoteAddr(remoteAddr string) string {
	remoteAddrParts := strings.Split(remoteAddr, ":")
	if len(remoteAddrParts) == 1 {
		return remoteAddr
	}

	return strings.Join(remoteAddrParts[:len(remoteAddrParts)-1], ":")
}
