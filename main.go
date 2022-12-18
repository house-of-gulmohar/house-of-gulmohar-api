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
	c.Pool = app.PrepareDBPool(&c)

	s := app.NewServer(&c)
	s.Start()
}
