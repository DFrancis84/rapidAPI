package types

import "time"

type Result struct {
	Definitions []Definition `json:"list"`
}

type Definition struct {
	Definition  string    `json:"definition"`
	Permalink   string    `json:"permalink"`
	ThumbsUp    int       `json:"thumbs_up"`
	Author      string    `json:"author"`
	Word        string    `json:"word"`
	DefID       int       `json:"defid"`
	CurrentVote string    `json:"current_vote"`
	WrittenOn   time.Time `json:"written_on"`
	Example     string    `json:"example"`
	ThumbsDown  int       `json:"thumbs_down"`
}
