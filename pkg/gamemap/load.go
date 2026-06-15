package gamemap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
)

func Load(fsys fs.FS, name string) (*Map, error) {
	raw, err := fs.ReadFile(fsys, name)
	if err != nil {
		return nil, fmt.Errorf("read map %q: %w", name, err)
	}
	return Parse(raw)
}

func Parse(raw []byte) (*Map, error) {
	normalized := bytes.TrimSpace(raw)
	normalized = bytes.TrimPrefix(normalized, []byte("var mapData ="))
	normalized = bytes.TrimSpace(bytes.TrimSuffix(normalized, []byte(";")))

	var out Map
	if err := json.Unmarshal(normalized, &out); err != nil {
		return nil, fmt.Errorf("decode map json: %w", err)
	}
	if out.Animated == nil {
		out.Animated = map[int]Animation{}
	}
	if out.StaticEntities == nil {
		out.StaticEntities = map[int]string{}
	}
	if err := out.Validate(); err != nil {
		return nil, err
	}
	return &out, nil
}

func (m *Map) Marshal(pretty bool) ([]byte, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	if pretty {
		return json.MarshalIndent(m, "", "  ")
	}
	return json.Marshal(m)
}
