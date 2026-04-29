package main 

import "context"



type OrderService interface {
	CreateOrder(ctx context.Context, )error
 }

 type OrderStore interface {
		Create(ctx context.Context) error
 }
 