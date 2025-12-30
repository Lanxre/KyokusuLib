package db

type Author struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Metier   string `json:"metier"`
	Picture  string `json:"picture"`
	Bio      string `json:"bio"`
}