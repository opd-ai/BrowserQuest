package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/opd-ai/BrowserQuest/pkg/client/platform"
	"github.com/opd-ai/BrowserQuest/pkg/client/render"
	"github.com/opd-ai/BrowserQuest/pkg/gamemap"
)

const (
	screenWidth  = 960
	screenHeight = 540
)

type App struct {
	assets platform.AssetProvider
	world  *gamemap.Map
}

func New(provider platform.AssetProvider) (*App, error) {
	world, err := provider.LoadMap("maps/world_client.json")
	if err != nil {
		return nil, err
	}

	return &App{
		assets: provider,
		world:  world,
	}, nil
}

func (a *App) Update() error {
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	screen.Fill(render.BootstrapBackground)
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"BrowserQuest Go Client\nMilestone 0 bootstrap\nMap: %dx%d tiles @ %dpx\nEmbedded assets ready",
		a.world.Width,
		a.world.Height,
		a.world.TileSize,
	))
}

func (a *App) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func Run(title string) error {
	app, err := New(platform.NewEmbeddedAssets())
	if err != nil {
		return err
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	return ebiten.RunGame(app)
}
