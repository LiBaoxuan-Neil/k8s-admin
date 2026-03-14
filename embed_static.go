package main

import "embed"

//go:embed all:web/dist
var EmbedFrontend embed.FS
