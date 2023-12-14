package models

type Cart struct {
	ID       uint       // Unique identifier for the cart
	UserID   uint       // Foreign key linking to the User who owns the cart
	Items    []CartItem // List of cart items
	Subtotal float64    // Total price of all items in the cart
	// Add more cart-related fields as needed
}
