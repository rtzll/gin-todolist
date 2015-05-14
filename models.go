package models

import "time"

type User struct {
	id            int
	username      string
	email         string
	password_hash string
	memberSince   time.Time
	lastSeen      time.Time
	isAdmin       bool
	todolists     []TodoList
}

type Todo struct {
	id          int
	description string
	createdAt   time.Time
	finishedAt  time.Time
	isFinished  bool
	creator     User
	todolistId  int
}

type TodoList struct {
	id        int
	title     string
	createdAt time.Time
	creator   User
	todos     []Todo
}
