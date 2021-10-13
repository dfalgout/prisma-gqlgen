package graph

import "github.com/dfalgout/prismagqlgen/prisma/db"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *db.PrismaClient
}

func NewResolver(client *db.PrismaClient) *Resolver {
	return &Resolver{
		client: client,
	}
}
