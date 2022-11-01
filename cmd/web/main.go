package main

import (
	"log"
	"strconv"

	"invoice-api/env"
	"invoice-api/web"
)

func main() {
	const PORT_DEFAULT_VALUE = 8000

	portStr := env.GetDefault("PORT", strconv.Itoa(PORT_DEFAULT_VALUE))

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf(`Specified port "%s" is invalid. Defaulting to %d`, portStr, PORT_DEFAULT_VALUE)
		port = PORT_DEFAULT_VALUE
	}

	server := web.NewServer(port)
	log.Fatalln(server.ListenAndServe())
}
