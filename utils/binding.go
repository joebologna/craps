package utils

import "fyne.io/fyne/v2/data/binding"

type BS struct{ binding.String }

func NewBS() BS { return BS{binding.NewString()} }

func (s BS) GetS() string { t, _ := s.Get(); return t }
