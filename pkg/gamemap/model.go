package gamemap

import "fmt"

type Map struct {
	Width          int               `json:"width"`
	Height         int               `json:"height"`
	TileSize       int               `json:"tilesize"`
	Data           []int             `json:"data,omitempty"`
	High           []int             `json:"high,omitempty"`
	Collisions     []int             `json:"collisions,omitempty"`
	Blocking       []int             `json:"blocking,omitempty"`
	Plateau        []int             `json:"plateau,omitempty"`
	Doors          []Door            `json:"doors,omitempty"`
	Checkpoints    []Checkpoint      `json:"checkpoints,omitempty"`
	MusicAreas     []MusicArea       `json:"musicAreas,omitempty"`
	Animated       map[int]Animation `json:"animated,omitempty"`
	RoamingAreas   []RoamingArea     `json:"roamingAreas,omitempty"`
	ChestAreas     []ChestArea       `json:"chestAreas,omitempty"`
	StaticChests   []StaticChest     `json:"staticChests,omitempty"`
	StaticEntities map[int]string    `json:"staticEntities,omitempty"`
}

type Door struct {
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Portal        int    `json:"p,omitempty"`
	TargetCameraX int    `json:"tcx,omitempty"`
	TargetCameraY int    `json:"tcy,omitempty"`
	TargetX       int    `json:"tx,omitempty"`
	TargetY       int    `json:"ty,omitempty"`
	Orientation   string `json:"to,omitempty"`
}

type Checkpoint struct {
	ID     int `json:"id"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w"`
	Height int `json:"h"`
	Safe   int `json:"s,omitempty"`
}

type MusicArea struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"w"`
	Height int    `json:"h"`
	ID     string `json:"id"`
}

type Animation struct {
	Length int `json:"l,omitempty"`
	Delay  int `json:"d,omitempty"`
}

type RoamingArea struct {
	ID     int    `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
	Count  int    `json:"nb,omitempty"`
}

type ChestArea struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"w"`
	Height int    `json:"h"`
	Items  []int  `json:"i,omitempty"`
	TypeX  string `json:"tx,omitempty"`
	TypeY  string `json:"ty,omitempty"`
}

type StaticChest struct {
	X     int   `json:"x"`
	Y     int   `json:"y"`
	Items []int `json:"i,omitempty"`
}

func (m *Map) Validate() error {
	if m.Width <= 0 {
		return fmt.Errorf("invalid width %d", m.Width)
	}
	if m.Height <= 0 {
		return fmt.Errorf("invalid height %d", m.Height)
	}
	if m.TileSize <= 0 {
		return fmt.Errorf("invalid tile size %d", m.TileSize)
	}
	return nil
}
