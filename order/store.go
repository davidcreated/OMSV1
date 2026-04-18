package main 
type store struct {
	// add here our mongoDB
}

func NewStore() *store {
	return &store{}
}

func (s *store) create(ctx context.Context) error {
	// here we will add the code to create an order in the database
	return nil
}