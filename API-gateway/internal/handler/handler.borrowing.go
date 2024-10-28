package handler

import "github.com/Sandhya-Pratama/Libary-API/API-gateway/config"

type Handler struct {
	config *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{config: cfg}
}
