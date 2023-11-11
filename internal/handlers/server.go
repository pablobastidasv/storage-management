package handlers

type Server interface {
	Start(addr string) error
}
