package main

import (
	"context"
	"log"
	"strconv"

	"invoice-api/env"
	"invoice-api/pdf"
	"invoice-api/web"
)

const DEFAULT_PORT = 8000
const DEFAULT_CHROME_DEV_TOOLS_URL = "ws://127.0.0.1:9222"

func main() {
	chromeUrl := env.GetDefault("CHROME_DEV_TOOLS_URL", DEFAULT_CHROME_DEV_TOOLS_URL)
	chromeCtx := pdf.NewRemoteChromeDevToolsContext(context.Background(), chromeUrl)

	routes := web.MakeRoutes(chromeCtx)

	portStr := env.GetDefault("PORT", strconv.Itoa(DEFAULT_PORT))

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf(`Specified port "%s" is invalid. Defaulting to %d`, portStr, DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	server := web.NewServer(port, routes)
	log.Fatalln(server.ListenAndServe())
}
