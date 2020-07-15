package types

type Count []struct {
	Country        string `json:"country"`
	Date           string `json:"date"`
	ConfirmedCount int    `json:"confirmed,omitempty"`
	DeathCount     int    `json:"death,omitempty"`
	RecoveredCount int    `json:"recovered,omitempty"`
}
