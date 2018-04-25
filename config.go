package main

import (
	"io"
	"sync"
)

// Config struct
//
// Using for Worker configuration
type Config struct {
	BaseURI string
	TagName string
	MaxDeep int
	MaxLink int
}

// SeenMap struct
//
// Using for storing parsed/visited links
type SeenMap struct {
	sync.Mutex
	data map[string]bool
}

// ResultsMap struct
//
// Using for storing resulting information
type ResultsMap struct {
	sync.Mutex
	data map[string]int
}

// Link struct
//
// Using for storing Link, parsed data and links level
type Link struct {
	URI   string
	Data  io.ReadCloser
	Level int
}

// Add function
//
// Added new URI to map
//
// Params:
// - key {string}
func (m *SeenMap) Add(key string) {
	m.Lock()
	m.data[key] = true
	m.Unlock()
}

// Add function
//
// Added new result to map
//
// Params:
// - key 		{string}
// - value 	{string}
func (m *ResultsMap) Add(key string, value int) {
	m.Lock()
	m.data[key] = value
	m.Unlock()
}

// Exist function
//
// Check if URI already visited
//
// Params:
// - key {string}
//
// Result:
// - {bool}
func (m *SeenMap) Exist(key string) (res bool) {
	m.Lock()
	res = m.data[key]
	m.Unlock()
	return
}
