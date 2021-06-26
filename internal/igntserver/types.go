package igntserver

type ignt struct {
	Seq  int    `json:"seq"`
	Name string `json:"name"`
}

type igntListResponse struct {
	Embedded struct {
		Ingredients []*ignt `json:"ingredientListResponseDtoList"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			HRef string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}
