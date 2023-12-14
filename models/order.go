package models

type Order struct {
	ID          uint    // Unique identifier for the order
	UserID      uint    // Foreign key linking to the User who placed the order
	ProductID   uint    // Foreign key linking to the Product in the order
	Quantity    int     // Quantity of the product in the order
	TotalAmount float64 // Total amount of the order
	// Add more order-related fields as needed
}
