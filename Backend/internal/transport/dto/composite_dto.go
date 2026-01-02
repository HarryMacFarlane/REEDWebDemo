package dto

// UserCompositeOutput composes multiple pieces of data into a single response.
// Extend this struct with additional fields from other models as needed.
type UserCompositeOutput struct {
	User  UserOutput  `json:"user"`
	Extra interface{} `json:"extra,omitempty"`
}
