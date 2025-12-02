package models

import "time"

// each vendor will have a Catalog
// Catalog represents a collection of products, often filtered by a Buyer's needs
// or a Vendor's offerings. This is more of a logical view container.
type Catalog struct {
	TotalProducts int       `json:"totalProducts"` // used for pagination and determine how to layout the UI
	FilterApplied string    `json:"filterApplied"` // used for pagination and determine how to layout the UI
	Products      []Product `json:"products"`
}

type Product struct { // (Represents a single SKU offered by a Vendor)
	// Must include inventory and specific compliance details.
	ID               string    `json:"id"`
	VendorID         string    `json:"vendorID"` // Foreign key to the Vendor document.
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Category         string    `json:"category"`
	SubCategory      string    `json:"subCategory"`
	IsMedical        bool	   `json:"isMedical"`
	PricePerUnit     float64   `json:"pricePerUnit"`
	AvailableUnits   int       `json:"availableUnits"`
	MinOrderQuantity int       `json:"minOrderQuantity"`
	MaxOrderQuantity int       `json:"maxOrderQuantity"`
	CoaLink          string    `json:"coaLink"`
	ComplianceTags   []string  `json:"complianceTags"` // ['THC: 25%', 'Sativa', 'Lab ID: 12345']
	UpdatedAt        time.Time `json:"updatedAt"`      // timestamp as a string
	// Stock            int      `json:"stock"`
}
