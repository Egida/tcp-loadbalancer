package config

import "flag"

type (
	Flags struct {
		Host string
		Port int
	}
)

var flags Flags

func init() {
	flg := Flags{}

	flag.StringVar(&flg.Host, "host", "localhost", "host to listen on")
	flag.IntVar(&flg.Port, "port", 8080, "port to listen on")

	flag.Parse()
}

func GetFlags() Flags {
	return flags
}

func GetAddress() string {
	return flags.Host + ":" + string(flags.Port)
}
