package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
    inputs []string
    outputs []string
    width int
    height int
    placeholder bool
    idx int
}

var (
    borderStyle = lipgloss.NewStyle().Border(lipgloss.ThickBorder()).Padding(0, 1)
    inBorderStyle = borderStyle.BorderForeground(lipgloss.Color("#9BEDAB"))
    outBorderStyle = borderStyle.BorderForeground(lipgloss.Color("#9BA6ED"))
    placeholderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("240"))
)

func (m Model) Init() tea.Cmd {
    return tea.EnterAltScreen
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            killOllama()
            return m, tea.Quit
        case "backspace":
            if len(m.inputs[m.idx]) > 0 {
                m.inputs[m.idx] = m.inputs[m.idx][:len(m.inputs[m.idx])-1]
            }
            if len(m.inputs[m.idx]) == 0 {
                m.placeholder = true
            }
        case "enter":
            if len(m.inputs[m.idx]) > 0 {
                out, err := runOllama(m.inputs[m.idx], "deepseek-r1:1.5b")
                if err != nil {
                    m.outputs = append(m.outputs, err.Error())
                } else {
                    m.outputs = append(m.outputs, out)
                }
                m.inputs = append(m.inputs, "")
                m.idx++
                m.placeholder = true
            }
        default:
            if m.placeholder {
                m.inputs[m.idx] = ""
                m.placeholder = false
            }
            m.inputs[m.idx] += msg.String()
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }
    return m, nil
}

func (m Model) View() string {
    // create past entries
    var entries string
    for i := 0; i < len(m.outputs); i++ {
        entries += outBorderStyle.Width(m.width-4).Render("> " + m.inputs[i] + "\n" + m.outputs[i])
        entries += "\n"
    }
    // add current entries
    var inputText string
    if m.placeholder {
        inputText = placeholderStyle.Render("Enter your prompt")
    } else {
        inputText = m.inputs[m.idx]
    }
    inBox := inBorderStyle.Width(m.width-4).Render("> " + inputText)
    return entries + inBox
}

func startInterface() {
    p := tea.NewProgram(Model{
        inputs: []string{""},
        outputs: []string{},
        placeholder: true,
        idx: 0,
    }, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
