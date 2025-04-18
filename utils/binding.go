package utils

import "fyne.io/fyne/v2/data/binding"

type BS struct{ binding.String }

func NewBS() BS { return BS{binding.NewString()} }

func NewBSWithString(text string) BS {
	bs := BS{binding.NewString()}
	bs.Set(text)
	return bs
}

func (s BS) GetS() string { t, _ := s.Get(); return t }
