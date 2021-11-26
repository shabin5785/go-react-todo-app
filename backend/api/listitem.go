package api

type Item struct {
	ID      int    `json:"id"`
	Item    string `json:"item"`
	Done    bool   `json:"done"`
	OldItem string `json:"olditem"`
}
