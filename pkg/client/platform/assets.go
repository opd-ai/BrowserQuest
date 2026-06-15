package platform

import (
	"fmt"
	"io/fs"

	assetbundle "github.com/opd-ai/BrowserQuest/assets"
	"github.com/opd-ai/BrowserQuest/pkg/gamemap"
)

type AssetProvider interface {
	FS() fs.FS
	ReadFile(name string) ([]byte, error)
	LoadMap(name string) (*gamemap.Map, error)
}

type EmbeddedAssets struct {
	fsys fs.FS
}

func NewEmbeddedAssets() EmbeddedAssets {
	return EmbeddedAssets{fsys: assetbundle.FS()}
}

func (p EmbeddedAssets) FS() fs.FS {
	return p.fsys
}

func (p EmbeddedAssets) ReadFile(name string) ([]byte, error) {
	data, err := fs.ReadFile(p.fsys, name)
	if err != nil {
		return nil, fmt.Errorf("read asset %q: %w", name, err)
	}
	return data, nil
}

func (p EmbeddedAssets) LoadMap(name string) (*gamemap.Map, error) {
	return gamemap.Load(p.fsys, name)
}
