package telegram

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int    `json:"update_id"`
	Message  string `json:"message"`
}
