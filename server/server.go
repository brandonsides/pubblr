package server

import "net/http"

type ActivityPubServer http.Server

func NewActivityPubServer() *ActivityPubServer {
	return &ActivityPubServer{}
}
