package main

func main() {
	store := NewStore()
	service := NewService(store)

	// here we will add the code to create an order using the service
	service.CreateOrder(context.Background())
}