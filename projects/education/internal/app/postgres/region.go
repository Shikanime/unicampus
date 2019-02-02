package neo4j

type Region struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zipcode string `json:"zipCode"`
}
