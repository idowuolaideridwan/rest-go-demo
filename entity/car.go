package entity

//Person object for REST(CRUD)
type Cars struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

//Person object for REST(CRUD)
type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Person object for REST(CRUD)
type States struct {
	ID        int    `json:"id"`
	CountryID string `json:"country_id"`
	Name      string `json:"name"`
}

//Person object for REST(CRUD)
type Payment struct {
	ID        int    `json:"id"`
	CountryID string `json:"country_id"`
	Name      string `json:"name"`
}

//Person object for REST(CRUD)
type History struct {
	ID        int    `json:"id"`
	CountryID string `json:"country_id"`
	Name      string `json:"name"`
}
