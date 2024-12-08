package server

type Websocket struct {
	port int
}

func (ws Websocket) Init() *Websocket{
	return &Websocket{
		port: 3000,
	}
}
