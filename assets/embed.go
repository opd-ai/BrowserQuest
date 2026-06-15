package assets

import (
	"embed"
	"io/fs"
)

//go:embed maps/** sprites/** audio/** img/** notices/**
var embedded embed.FS

func FS() fs.FS {
	return embedded
}
