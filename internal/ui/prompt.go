package ui

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

type Prompter struct {
	In  io.Reader
	Out io.Writer
}

func NewPrompter(in io.Reader, out io.Writer) *Prompter {
	return &Prompter{In: in, Out: out}
}

func (p *Prompter) Prompt(label, defaultValue string) (string, error) {
	if defaultValue != "" {
		fmt.Fprintf(p.Out, "%s [%s]: ", label, defaultValue)
	} else {
		fmt.Fprintf(p.Out, "%s: ", label)
	}
	reader := bufio.NewReader(p.In)
	value, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultValue, nil
	}
	return value, nil
}

func (p *Prompter) PromptSecret(label string) (string, error) {
	// Print the prompt first
	fmt.Fprintf(p.Out, "%s: ", label)

	// If the input is a file and a terminal, use term.ReadPassword to hide input.
	if file, ok := p.In.(*os.File); ok {
		if term.IsTerminal(int(file.Fd())) {
			// ReadPassword reads from the terminal connected to the given file descriptor.
			bytes, err := term.ReadPassword(int(file.Fd()))
			// After ReadPassword, echoing is restored and no newline is printed, so print one.
			fmt.Fprintln(p.Out)
			if err == nil {
				return strings.TrimSpace(string(bytes)), nil
			}
			// fallthrough to fallback Prompt on error
		}
	}

	// Fallback to normal prompt (visible input)
	return p.Prompt(label, "")
}

func (p *Prompter) Confirm(label string, defaultYes bool) (bool, error) {
	hint := "[y/N]"
	if defaultYes {
		hint = "[Y/n]"
	}
	answer, err := p.Prompt(label+" "+hint, "")
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(strings.ToLower(answer))
	if answer == "" {
		return defaultYes, nil
	}
	return strings.HasPrefix(answer, "y"), nil
}

// setTTYEcho removed: platform-specific stty usage replaced by golang.org/x/term
