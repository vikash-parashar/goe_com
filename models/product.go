package models

type Product struct {
	ID                uint
	Name              string
	Description       string
	Price             float64
	ImageURL          string
	AvailableQuantity int
	InventoryQuantity int
	NewQuantity       int
	OldQuantity       int
	CountryOfOrigin   string
	LaunchYear        int
	Category          string
}
