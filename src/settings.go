package main

import (
	"github.com/caarlos0/env"
)

var Settings *SettingsModel

type SettingsModel struct {
	TmuxPath string `env:"TMUX_PATH" envDefault:"tmux"`
}

func LoadSettings() {
	Settings = &SettingsModel{}
	_ = env.Parse(Settings)
}
