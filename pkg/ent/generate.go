package ent

import (
	_ "github.com/olekukonko/tablewriter"
	_ "golang.org/x/tools/go/packages"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
