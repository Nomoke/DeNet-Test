package models

type LeaderboardResp struct {
	Position int    `json:"position"`
	Nickname string `json:"nickname"`
	Points   int    `json:"points"`
}