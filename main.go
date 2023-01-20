// package main

// import (
// 	"cobb"
// 	"fmt"
// )

// func main() {
// 	archive := cobb.Archive{Camera: "Canon A1", Roll: "Portra 400", Date: "202211"}
// 	fmt.Println(archive.MakeArchiveName())
// }

package main

import (
	"cobb"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list    []list.Model
	cursor  int
	archive cobb.Archive
}

func (m model) Init() tea.Cmd {
	return nil
}

func ResizeWindowFrame(m model, msg tea.WindowSizeMsg) {
	h, v := docStyle.GetFrameSize()
	m.list[0].SetSize(msg.Width-h, msg.Height-v)
	m.list[1].SetSize(msg.Width-h, msg.Height-v)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == "enter" && m.cursor != len(m.list) {
			var cmd tea.Cmd
			m.list[m.cursor], cmd = m.list[m.cursor].Update(msg)
			if m.cursor == 0 {
				m.archive.Camera = m.list[0].SelectedItem().FilterValue()
			}
			if m.cursor == 1 {
				m.archive.Roll = m.list[1].SelectedItem().FilterValue()
			}
			m.cursor++
			return m, cmd
		}
	case tea.WindowSizeMsg:
		ResizeWindowFrame(m, msg)
	}

	if m.cursor != len(m.list) {
		var cmd tea.Cmd
		m.list[m.cursor], cmd = m.list[m.cursor].Update(msg)
		return m, cmd
	}

	return m, tea.Quit

}

func (m model) View() string {
	if m.cursor != len(m.list) {
		return m.list[m.cursor].View()
	}

	return fmt.Sprintf("Program end - %s ", m.archive.MakeArchiveName())
}

func main() {
	items := []list.Item{
		item{title: "Canon A1", desc: "Actually a great begineer camera"},
		item{title: "Mamiya 7", desc: "The analog Rolls-Royce"},
	}

	other_items := []list.Item{
		item{title: "Kodak Portra 400", desc: "Pro roll"},
		item{title: "Fujifilm XTRA 400", desc: "Great roll"},
	}

	var m model

	first_list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	second_list := list.New(other_items, list.NewDefaultDelegate(), 0, 0)
	m.list = []list.Model{first_list, second_list}
	m.cursor = 0
	m.list[0].Title = "Cameras"
	m.list[1].Title = "Film Rolls"
	m.archive.Date = "202311"

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
