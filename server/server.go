package server

import (
	"net/http"
	"strconv"

	"github.com/brandonsides/pubblr/database"
	"github.com/brandonsides/pubblr/logging"
	"github.com/brandonsides/pubblr/server/auth"
	"github.com/go-chi/chi"
)

type PubblrServer struct {
	http.Server
}

type PubblrServerConfig struct {
	MountPath string                        `json:"mountPath"`
	Database  database.PubblrDatabaseConfig `json:"database"`
	Logger    logging.PubblrLoggerConfig    `json:"logger"`
	Auth      auth.AuthConfig               `json:"auth"`
	Host      string                        `json:"host"`
	Port      int                           `json:"port"`
	PageSize  int                           `json:"pageSize"`
}

func NewPubblrServer(config PubblrServerConfig, baseRouter chi.Router) *http.Server {
	router, err := NewPubblrRouter(config, baseRouter)
	if err != nil {
		panic(err)
	}
	return &http.Server{
		Addr:    config.Host + ":" + strconv.Itoa(config.Port),
		Handler: router,
	}
}
