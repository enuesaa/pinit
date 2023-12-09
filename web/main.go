package web

import (
	"embed"
)

//go:generate pnpm install
//go:generate pnpm build

//go:embed all:dist/*
var Dist embed.FS
