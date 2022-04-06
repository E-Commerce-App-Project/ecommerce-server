package repository

import "github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"

// Option anything any repo object needed
type Option struct {
	commons.Options
}

// Repository all repo object injected here
type Repository struct {
	// TODO: add all repo object here
	Auth  IAuthRepository
	Cache ICacheRepository
}
