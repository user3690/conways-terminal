package terminal

import (
	"golang.org/x/term"
	"os"
)

type KeyboardReader struct {
	oldState *term.State
}

func NewKeyboard() (*KeyboardReader, error) {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return &KeyboardReader{}, nil
	}

	return &KeyboardReader{
		oldState: oldState,
	}, nil
}

func (reader *KeyboardReader) keyPress() (key []byte, err error) {
	key = make([]byte, 1)
	_, err = os.Stdin.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (reader *KeyboardReader) restore() error {
	return term.Restore(int(os.Stdin.Fd()), reader.oldState)
}
