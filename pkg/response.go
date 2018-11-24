package root

type response interface {
}

type ResponseSlice struct {
	Message string   `json:"message"`
	Data    response `json:"data"`
	Err     error    `json:"err"`
}
