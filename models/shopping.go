package models

// Item is something on a List to purchase
type Item struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	ListID      int64  `gorm:"index" json:"list_id"`
}

// Status contains the state model for a List
type Status int

const (
	// StatusNew is a new shopping list
	StatusNew Status = iota

	// StatusDone is a completed shopping list
	StatusDone
)

// List is a list of stuff to buy
type List struct {
	Model
	Items  []Item `json:"items"`
	Status Status `json:"status"`
}
