package models

type CartItem struct {
	ID         uint    // Unique identifier for the cart item
	UserID     uint    // Foreign key linking to the User who owns the cart
	ProductID  uint    // Foreign key linking to the Product in the cart
	Quantity   int     // Quantity of the product in the cart
	TotalPrice float64 // Total price for the cart item (Quantity * Product Price)
	// Add more cart item-related fields as needed
}
