# bittorrent-go
![coverage](https://raw.githubusercontent.com/HoikanChan/my-bittorrent-go/badges/.badges/main/coverage.svg)
[![progress-banner](https://backend.codecrafters.io/progress/bittorrent/00204093-f0ca-4f92-9942-891d0c7155aa)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is a starting point for Go solutions to the
["Build Your Own BitTorrent" Challenge](https://app.codecrafters.io/courses/bittorrent/overview).

In this challenge, you’ll build a BitTorrent client that's capable of parsing a
.torrent file and downloading a file from a peer. Along the way, we’ll learn
about how torrent files are structured, HTTP trackers, BitTorrent’s Peer
Protocol, pipelining and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

# Stages
- [x] Introduction
- [x] Repository Setup
- [x] Decode bencoded strings
- [x] Decode bencoded integers
- [x] Decode bencoded lists
- [x] Decode bencoded dictionaries
- [x] Parse torrent file
- [ ] Calculate info hash
- [ ] Piece hashes
- [ ] Discover peers
- [ ] Peer handshake
- [ ] Download a piece
- [ ] Download the whole file
