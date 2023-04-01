package term

import (
	"strings"
	"time"

	"atomicgo.dev/cursor"
)

var (
	Frames1 = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	Frames2 = []string{"-", "\\", "|", "/"}
	Frames3 = []string{"◜", "◠", "◝", "◞", "◡", "◟"}
)

// spinner is a simple spinner that can be used to indicate progress.
type spinner struct {
	// Frames is the list of frames to use for the spinner.
	Frames []string
	// Index is the current index of the spinner.
	Index int
	// Suffix is the suffix to use for the spinner.
	Suffix string
	// Ticker is the ticker used for the spinner.
	Ticker *time.Ticker
}

// NewSpinner returns a new spinner.
func NewCustomSpinner(frames []string) *spinner {
	return &spinner{
		Frames:   frames,
	}
}

func NewSpinner() *spinner {
	return NewCustomSpinner(Frames1)
}

// Stop stops the spinner.
func (s *spinner) Stop() {
	if s.Ticker != nil {
		s.Ticker.Stop()
	}
	s.Ticker = nil
	cursor.ClearLine()
	cursor.StartOfLine()
}

// Start starts the spinner.
func (s *spinner) Start(interval time.Duration) error {
	s.Ticker = time.NewTicker(interval)
	go func() {
		for range s.Ticker.C {
			s.Index = (s.Index + 1) % len(s.Frames)
			cursor.StartOfLine()
			print(s.Frames[s.Index] + s.Suffix)
		}
	}()
	return nil
}

// SetString sets the suffix of the spinner.
// The suffix is trimmed and the first line is used.
// Because the spinner is always on the same line, the suffix should not contain "\n".
func (s *spinner) SetString(suffix string) {
	suffix = strings.TrimSpace(suffix)
	suffix = strings.Split(suffix, "\n")[0]
	s.Suffix = " " + suffix
}
