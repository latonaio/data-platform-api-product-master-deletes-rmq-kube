package requests

type General struct {
	Product              	string  `json:"Product"`
	IsMarkedForDeletion		*bool   `json:"IsMarkedForDeletion"`
}
