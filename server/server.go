package server

import (
	"net/http"
	"strconv"

	"github.com/brandonsides/pubblr/database"
	"github.com/brandonsides/pubblr/logging"
	"github.com/go-chi/chi"
)

type PubblrServer struct {
	http.Server
}

type PubblrServerConfig struct {
	MountPath string                        `json:"mountPath"`
	Database  database.PubblrDatabaseConfig `json:"database"`
	Logger    logging.PubblrLoggerConfig    `json:"logger"`
	Host      string
	Port      int
}

func NewPubblrServer(config PubblrServerConfig, baseRouter chi.Router) *http.Server {
	return &http.Server{
		Addr:    config.Host + ":" + strconv.Itoa(config.Port),
		Handler: NewPubblrRouter(config, baseRouter),
	}
}
