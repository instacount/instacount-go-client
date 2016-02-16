package instacount

type Errors struct {
	Errors    []Error `json:"errors"`
	StatuCode int `json:"statusCode"`
}

type Error struct {
	Message          string `json:"message"`
	DeveloperMessage string  `json:"developerMessage"`
	MoreInfo         int `json:"moreInfo"`
	StatusCode       string `json:"statusCode"`
}