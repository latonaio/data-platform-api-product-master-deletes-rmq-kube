package requests

type BusinessPartner struct {
	Product              	string  `json:"Product"`
	BusinessPartner         int     `json:"BusinessPartner"`
	IsMarkedForDeletion		*bool   `json:"IsMarkedForDeletion"`
}
