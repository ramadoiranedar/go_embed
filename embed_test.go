package go_embed

import (
	_ "embed"
)

//go:embed version.txt
var version string
