package server

type Command struct {
	Version uint16
	ID      string
	Action  string
	Data    map[string]string
}
