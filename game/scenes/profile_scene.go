package scenes

import (
	"fmt"
	"image/color"

	"github.com/allanjose001/go-battleship/game/components"
	"github.com/allanjose001/go-battleship/game/components/basic"
	"github.com/allanjose001/go-battleship/game/components/basic/colors"
	"github.com/allanjose001/go-battleship/game/state"
	"github.com/allanjose001/go-battleship/internal/service"
	"github.com/hajimehoshi/ebiten/v2"
)

type ProfileScene struct {
	state   *state.GameState
	profile *service.Profile
	root    *components.Container
}

func NewProfileScene(s *state.GameState) *ProfileScene {
	// 1. Carrega o perfil do usuário
	p, _ := service.FindProfile("malub")
	screenSize := basic.Size{W: 1280, H: 720} //

	// 2. Título 
	titleTxt := components.NewText(basic.Point{}, "PERFIL DE JOGADOR", colors.White, 38)

	// 3. Status (Row) com Size definido
	rowStats := components.NewRow(
		basic.Point{},
		30.0,
		basic.Size{W: screenSize.W * 0.8, H: 120},
		basic.Center,
		basic.Center,
		[]components.Widget{
			createStatCard("Partidas", fmt.Sprintf("%d", p.GamesPlayed), 180, 100),
			createStatCard("Vitórias", "1", 180, 100),
			createStatCard("Taxa", "10%", 180, 100),
		},
	)

	// 4. Mural de Medalhas (Título + Grid Dinâmico)
	muralTitle := components.NewText(basic.Point{}, "MURAL DE MEDALHAS", colors.White, 20)
	
	medals := []struct{ Icon, Title, Desc string }{
		{"🥇", "VETERANO", "10+ Partidas"},
		{"🎯", "SNIPER", "90% Precisão"},
	}
	medalGrid := createMedalGrid(medals, screenSize)

	// 5. Botão Retornar
	btnReturn := components.NewButton(
		basic.Point{}, 
		basic.Size{W: 200, H: 50}, 
		"RETORNAR",
		color.RGBA{R: 40, G: 50, B: 75, A: 255}, 
		colors.White,
		func(b *components.Button) { fmt.Println("Sair") },
	)

	// 6. Coluna Principal 
	mainColumn := components.NewColumn(
		basic.Point{X: 0, Y: 40}, // Distância do topo
		50,                        // Espaçamento entre os blocos 
		screenSize,
		basic.Start,               // Alinha do topo para baixo
		basic.Center,              // Centraliza horizontalmente
		append([]components.Widget{titleTxt, rowStats, muralTitle}, append(medalGrid, btnReturn)...),
	)

	// 7. Root Container
	root := components.NewContainer(
		basic.Point{}, 
		screenSize, 
		0, 
		color.RGBA{R: 10, G: 20, B: 30, A: 255}, 
		basic.Start, 
		basic.Start, 
		mainColumn, 
		nil,
	)

	return &ProfileScene{state: s, profile: p, root: root}
}

func createMedalGrid(data []struct{ Icon, Title, Desc string }, screen basic.Size) []components.Widget {
	var rows []components.Widget
	medalW, medalH := screen.W*0.25, screen.H*0.08
	for i := 0; i < len(data); i += 2 {
		var widgets []components.Widget
		for j := 0; j < 2 && (i+j) < len(data); j++ {
			m := data[i+j]
			widgets = append(widgets, createMedalCard(m.Icon, m.Title, m.Desc, medalW, medalH))
		}
		rows = append(rows, components.NewRow(basic.Point{}, 20, screen, basic.Center, basic.Center, widgets))
	}
	return rows
}

func createMedalCard(icon, title, desc string, w, h float32) *components.Container {
	t := components.NewText(basic.Point{}, title, colors.Black, 14)
	d := components.NewText(basic.Point{}, desc, color.RGBA{R: 80, G: 80, B: 80, A: 255}, 9)
	textCol := components.NewColumn(basic.Point{}, 2, basic.Size{W: w * 0.6, H: h}, basic.Center, basic.Start, []components.Widget{t, d})
	iconTxt := components.NewText(basic.Point{}, icon, colors.Black, 20)
	content := components.NewRow(basic.Point{}, 10, basic.Size{W: w, H: h}, basic.Start, basic.Center, []components.Widget{iconTxt, textCol})
	return components.NewContainer(basic.Point{}, basic.Size{W: w, H: h}, 10, colors.White, basic.Center, basic.Center, content, nil)
}

func createStatCard(label, value string, w, h float32) *components.Container {
	lbl := components.NewText(basic.Point{}, label, colors.Black, 14)
	val := components.NewText(basic.Point{}, value, colors.Black, 22)
	content := components.NewColumn(basic.Point{}, 5, basic.Size{W: w, H: h}, basic.Center, basic.Center, []components.Widget{lbl, val})
	return components.NewContainer(basic.Point{}, basic.Size{W: w, H: h}, 12, colors.White, basic.Center, basic.Center, content, nil)
}

func (s *ProfileScene) OnEnter(prev Scene, size basic.Size) {}
func (s *ProfileScene) OnExit(next Scene)                   {}
func (s *ProfileScene) Update() error {
	s.root.Update(basic.Point{X: 0, Y: 0})
	return nil
}
func (s *ProfileScene) Draw(screen *ebiten.Image) {
	s.root.Draw(screen)
}