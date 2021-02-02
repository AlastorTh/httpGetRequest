package models

// Response ...
type Response struct {
	Args    map[string]string
	Headers map[string]string
	Origin  string
	URL     string
}
