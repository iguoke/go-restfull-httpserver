package main

import "time"

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `josn:"completed"`
	Due       time.Time `josn:"due"`
}

type Todos []Todo
