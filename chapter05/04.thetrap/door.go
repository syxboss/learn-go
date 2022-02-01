package main

type Door interface {
	unlock()
	Open()
	Close()
	lock()
}
