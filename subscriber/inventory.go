package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	inventory "inventory-service/proto/inventory"
)

type Inventory struct{}

func (e *Inventory) Handle(ctx context.Context, msg *inventory.Supplier) error {
	log.Log("Handler Received Supplier: " + msg.nameEn)

	return nil
}

// This is the same above but only need to demo difrent invoking needs
func Handler(ctx context.Context, msg *inventory.Supplier) error {
	log.Log("Function Received Supplier: ")
	return nil
}
