package server

type (
	server struct{}
	Server interface{}
)

func NewServer() Server {
	return &server{}
}

func main() {}
