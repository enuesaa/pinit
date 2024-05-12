package main

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
}
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
func (a *App) Greet2(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
