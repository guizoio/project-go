package entities

type Address struct {
	ID       string `json:"id"`
	Street   string `json:"street"`
	District string `json:"district"`
	Number   string `json:"number"`
	City     string `json:"city"`
	State    string `json:"state"`
	Enabled  bool   `json:"enabled"`
	UserId   string `json:"user_id"`
}
