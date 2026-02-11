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
	p, _ := service.FindProfile("malub")
	// Resolução lógica travada do projeto
	sw, sh := float32(1280), float32(720)
	screenSize := basic.Size{W: sw, H: sh}

	// --- 1. SEÇÃO DO TÍTULO ---
	// Criamos um container invisível que ocupa a largura total da tela no topo
	titleSection := components.NewContainer(
		basic.Point{X: 0, Y: 40},
		basic.Size{W: sw, H: 80},
		0, nil, // Transparente/Invisível
		basic.Center, basic.Center, // Centraliza o texto dentro dos 1280px
		components.NewText(basic.Point{}, "PERFIL DE JOGADOR", colors.White, 40),
		nil,
	)

	// --- 2. SEÇÃO DE STATUS (PARTIDAS, VITÓRIAS, ETC) ---
	statusSection := components.NewContainer(
		basic.Point{X: 0, Y: 140},
		basic.Size{W: sw, H: 140},
		0, nil,
		basic.Center, basic.Center,
		components.NewRow(
			basic.Point{}, 40, basic.Size{W: 800, H: 140},
			basic.Center, basic.Center,
			[]components.Widget{
				createStatCard("Partidas", fmt.Sprintf("%d", p.GamesPlayed), 200, 110),
				createStatCard("Vitórias", "1", 200, 110),
				createStatCard("Taxa", "10%", 200, 110),
			},
		),
		nil,
	)

	// --- 3. SEÇÃO DO MURAL DE MEDALHAS ---
	// Título do Mural
	muralLabelSection := components.NewContainer(
		basic.Point{X: 0, Y: 320},
		basic.Size{W: sw, H: 40},
		0, nil,
		basic.Center, basic.Center,
		components.NewText(basic.Point{}, "MURAL DE MEDALHAS", colors.White, 22),
		nil,
	)

	// Balões das Medalhas
	medalsSection := components.NewContainer(
		basic.Point{X: 0, Y: 380},
		basic.Size{W: sw, H: 120},
		0, nil,
		basic.Center, basic.Center,
		components.NewRow(
			basic.Point{}, 30, basic.Size{W: 800, H: 120},
			basic.Center, basic.Center,
			[]components.Widget{
				createMedalCard("🥇", "VETERANO", "10+ Partidas", 350, 90),
				createMedalCard("🎯", "SNIPER", "90% Precisão", 350, 90),
			},
		),
		nil,
	)

	// --- 4. SEÇÃO DO BOTÃO RETORNAR ---
	buttonSection := components.NewContainer(
		basic.Point{X: 0, Y: 580},
		basic.Size{W: sw, H: 80},
		0, nil,
		basic.Center, basic.Center,
		components.NewButton(
			basic.Point{}, basic.Size{W: 220, H: 55},
			"RETORNAR", color.RGBA{45, 67, 103, 255}, colors.White,
			func(b *components.Button) { fmt.Println("Sair") },
		),
		nil,
	)

	// --- LAYOUT FINAL ---
	// Usamos uma Column apenas para agrupar as seções que já estão posicionadas
	mainLayout := components.NewColumn(
		basic.Point{}, 0, screenSize,
		basic.Start, basic.Start,
		[]components.Widget{titleSection, statusSection, muralLabelSection, medalsSection, buttonSection},
	)

	root := components.NewContainer(
		basic.Point{}, screenSize, 0,
		color.RGBA{10, 25, 40, 255},
		basic.Start, basic.Start,
		mainLayout, nil,
	)

	return &ProfileScene{state: s, profile: p, root: root}
}

// Helpers que usam o Container interno para alinhar o texto
func createStatCard(label, value string, w, h float32) *components.Container {
	content := components.NewColumn(
		basic.Point{}, 8, basic.Size{W: w, H: h},
		basic.Center, basic.Center,
		[]components.Widget{
			components.NewText(basic.Point{}, label, colors.Black, 14),
			components.NewText(basic.Point{}, value, colors.Black, 24),
		},
	)
	return components.NewContainer(basic.Point{}, basic.Size{W: w, H: h}, 15, colors.White, basic.Center, basic.Center, content, nil)
}

func createMedalCard(icon, title, desc string, w, h float32) *components.Container {
	textCol := components.NewColumn(
		basic.Point{}, 2, basic.Size{W: w * 0.7, H: h},
		basic.Center, basic.Start,
		[]components.Widget{
			components.NewText(basic.Point{}, title, colors.Black, 15),
			components.NewText(basic.Point{}, desc, color.RGBA{80, 80, 80, 255}, 10),
		},
	)
	content := components.NewRow(
		basic.Point{}, 15, basic.Size{W: w, H: h},
		basic.Start, basic.Center,
		[]components.Widget{components.NewText(basic.Point{}, icon, colors.Black, 22), textCol},
	)
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