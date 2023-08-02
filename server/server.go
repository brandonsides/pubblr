package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/brandonsides/pubblr/logging"
	"github.com/brandonsides/pubblr/server/apiutil"
	"github.com/go-chi/chi"
)

type PubblrServer struct {
	http.Server
	logger apiutil.Logger
}

func NewPubblrServer(config PubblrRouterConfig, baseRouter chi.Router) (*PubblrServer, error) {
	logger := logging.NewStandardPubblrLogger(config.Logger)

	router, err := NewPubblrRouter(config, baseRouter, logger)
	if err != nil {
		return nil, fmt.Errorf("Failed to create pubblr router: %w", err)
	}
	return &PubblrServer{
		Server: http.Server{
			Addr:    config.Host + ":" + strconv.Itoa(config.Port),
			Handler: router,
		},
		logger: logger,
	}, nil
}

func (s *PubblrServer) ListenAndServe() error {
	s.logger.Debugf("Listening on %s", s.Addr)
	return s.Server.ListenAndServe()
}
