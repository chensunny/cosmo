// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type TimestampedString struct {
	// The value of the string.
	Value string `json:"value"`
	// The timestamp when the response was generated.
	UnixTime int `json:"unixTime"`
	// Sequence number
	Seq int `json:"seq"`
	// Total number of responses to be sent
	Total int `json:"total"`
}