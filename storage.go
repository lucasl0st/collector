package main

type Storage interface {
	Store(entity, id string, labels []string, data []byte) error
}
