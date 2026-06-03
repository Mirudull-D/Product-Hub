package cart

import (
	"Product-Hub/db/generated"
	"Product-Hub/types"
	"context"
	"fmt"
	"strconv"
)

func GetCartitemsIds(items []types.CartItem) ([]int32, error) {
	productIds := make([]int32, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("quantity must be greater than 0")
		}
		productIds[i] = int32(item.ProductId)
	}
	return productIds, nil
}
func (h *Handler) createOrder(ctx context.Context,
	ps []generated.Product, items []types.CartItem,
	userID int64) (int64, float64, error) {

	fmt.Println("Creating order...")
	productMap := make(map[int]generated.Product)
	for _, product := range ps {
		productMap[int(product.ID)] = product
	}

	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}
	totalCost := calculateTotalCost(items, productMap)

	for _, item := range items {
		product := productMap[int(item.ProductId)]
		product.Quantity -= int32(item.Quantity)

		err := h.productStore.UpdateProduct(ctx, generated.UpdateProductParams{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Image:       product.Image,
			Description: product.Description,
			Quantity:    product.Quantity,
		})
		if err != nil {
			return 0, 0, err
		}
	}
	createdOrderId, err := h.store.CreateOrder(ctx, generated.CreateOrderParams{
		UserID:  int64(userID),
		Total:   strconv.FormatFloat(totalCost, 'f', 2, 64),
		Status:  "pending",
		Address: "A.dsdgfds",
	})
	fmt.Println("Created order:", createdOrderId)
	if err != nil {
		return 0, 0, err
	}
	for _, item := range items {
		err2 := h.store.CreateOrderItems(ctx, generated.CreateOrderItemsParams{
			OrderID:   createdOrderId,
			ProductID: int64(item.ProductId),
			Quantity:  int32(item.Quantity),
			Price:     productMap[item.ProductId].Price,
		})
		if err2 != nil {
			return 0, 0, err2
		}
		fmt.Println("Creating order item for product:", item.ProductId)
	}
	return createdOrderId, totalCost, nil

}

func checkIfCartIsInStock(cartitems []types.CartItem, products map[int]generated.Product) error {
	if len(cartitems) == 0 {
		return fmt.Errorf("no cartitems")
	}

	for _, item := range cartitems {
		product, ok := products[int(item.ProductId)]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductId)
		}
		if int(product.Quantity) < item.Quantity {
			return fmt.Errorf(
				"only %d items left in stock",
				product.Quantity,
			)
		}
	}
	return nil
}

func calculateTotalCost(items []types.CartItem, products map[int]generated.Product) float64 {
	var total float64
	for _, item := range items {
		product := products[int(item.ProductId)]
		price, err := strconv.ParseFloat(product.Price, 64)
		if err != nil {
			continue
		}
		total += price * float64(item.Quantity)
	}
	return total
}
