package selector

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"

	"tm/internal/discovery"
	"tm/internal/worktree"
)

type Choice struct {
	Label   string
	Details string
	Value   string
}

func (c Choice) Title() string {
	return c.Label
}

func (c Choice) Description() string {
	return c.Details
}

func (c Choice) FilterValue() string {
	return strings.TrimSpace(c.Label + "\n" + c.Details)
}

type BubbleSelector struct {
	in  *os.File
	out io.Writer
}

func NewBubbleSelector(in *os.File, out io.Writer) *BubbleSelector {
	return &BubbleSelector{in: in, out: out}
}

func IsTTY() bool {
	stdinOK := term.IsTerminal(int(os.Stdin.Fd()))
	stdoutOK := term.IsTerminal(int(os.Stdout.Fd()))
	return stdinOK && stdoutOK
}

func RepoChoices(repos []discovery.Repo) []Choice {
	choices := make([]Choice, 0, len(repos))
	for _, repo := range repos {
		choices = append(choices, Choice{Label: repo.Name, Details: repo.Path, Value: repo.Path})
	}
	return choices
}

func TargetChoices(targets []worktree.Target) []Choice {
	choices := make([]Choice, 0, len(targets))
	for _, target := range targets {
		choices = append(choices, Choice{Label: target.Name, Details: target.Path, Value: target.Path})
	}
	return choices
}

func (s *BubbleSelector) Select(title string, items []Choice) (Choice, bool, error) {
	listItems := make([]list.Item, 0, len(items))
	for _, item := range items {
		listItems = append(listItems, item)
	}

	l := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.SetShowHelp(true)

	program := tea.NewProgram(selectionModel{list: l}, tea.WithInput(s.in), tea.WithOutput(s.out))
	result, err := program.Run()
	if err != nil {
		return Choice{}, false, fmt.Errorf("run selector: %w", err)
	}

	model, ok := result.(selectionModel)
	if !ok {
		return Choice{}, false, fmt.Errorf("unexpected selector result type %T", result)
	}
	if model.cancelled || model.choice.Value == "" {
		return Choice{}, false, nil
	}

	return model.choice, true, nil
}

type selectionModel struct {
	list      list.Model
	choice    Choice
	cancelled bool
}

func (m selectionModel) Init() tea.Cmd { return nil }

func (m selectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.cancelled = true
			return m, tea.Quit
		case "enter":
			selected, ok := m.list.SelectedItem().(Choice)
			if ok {
				m.choice = selected
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m selectionModel) View() string {
	return m.list.View()
}
