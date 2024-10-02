// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Match struct {
	ID string `json:"id"`
	// Home team
	Home *Team `json:"home"`
	// Away team
	Away *Team `json:"away"`
	// When is kickoff? Timezone is UTC.
	// Encoded as RFC3339 (e.g. "2006-01-02T15:04:05Z07:00")
	Kickoff string `json:"kickoff"`
	// Which tournament is this match part of? (e.g. Premier Leage, Champions League, etc.)
	Tournament string `json:"tournament"`
	// golang internal time.Time representation of the event
	RawTime time.Time `json:"-"`
}

type Query struct {
}

type Team struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Shortcode string `json:"shortcode"`
	Crest     string `json:"crest"`
}

type Tournament struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}