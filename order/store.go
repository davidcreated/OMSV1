package main

import "context"

type store struct {
	// add here our mongoDB
}

// Create implements [OrderStore].
func (s *store) Create(ctx context.Context) error {
	panic("unimplemented")
}

func NewStore() *store {
	return &store{}
}

func (s *store) create(ctx context.Context) error {
	// here we will add the code to create an order in the database
	return nil
}
