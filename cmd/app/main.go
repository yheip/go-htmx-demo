package main

import (
	"context"
	"sse-demo/internal/application"
)

func main() {
	ctx := context.Background()
	app := application.New(ctx)
	app.Start(ctx)
}
