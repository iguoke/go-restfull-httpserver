package main

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `josn:"completed"`
	Due       time.Time `josn:"due"`
}

type Todos []Todo
