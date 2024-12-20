package models

type Campaign struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Status      string   `json:"status" bson:"status"`
	Leads       []string `json:"leads" bson:"leads"`
	AccountIDs  []string `json:"accountIDs" bson:"accountIDs"`
}
