package main

import (
	"github.com/leslieleung/ptpt/cmd"
	"github.com/leslieleung/ptpt/internal/config"
	_ "github.com/leslieleung/ptpt/static"
)

func main() {
	config.Init()
	cmd.Execute()
}
