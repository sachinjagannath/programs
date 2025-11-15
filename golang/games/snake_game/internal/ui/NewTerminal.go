package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

type Terminal struct {
	oldState *term.State
	width    int
	height   int
	buffer   strings.Builder
}

func NewTerminal() (*Terminal, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("Failed to set raw mode %w", err)
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		term.Restore(int(os.Stdin.Fd()), oldState)
		return nil, fmt.Errorf("failed to get terminal size: %w", err)
	}

	t := &Terminal{
		oldState: oldState,
		width:    width,
		height:   height,
	}
	fmt.Print("\033[?25l")
	t.Clear()
	return t, nil
}
