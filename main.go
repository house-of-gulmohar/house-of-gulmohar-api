package main

import (
	"house-of-gulmohar/internal/app"
)

func main() {
	c := app.Config{
		Name:    "House Of Gulmohar",
		Version: "v0.0.1",
	}
	c.Port = app.GetPort()
	c.PgConnStr = app.GetPgConfig()
	c.Db = app.PrepareDB(&c)

	s := app.NewServer(&c)
	s.Start()
}
