package main

type TorrentInfo struct {
	Announce  string `json:"announce"`
	Info      struct {
		Length      int    `json:"length"`
		Name        string `json:"name"`
		PieceLength int    `json:"piece length"`
		Pieces      string `json:"pieces"`
	} `json:"info"`
}
