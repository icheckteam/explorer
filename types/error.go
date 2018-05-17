package types

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Msg string `json:"msg"`
}
