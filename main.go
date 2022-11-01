package main

import (
	"github.com/frenchaindev/frenchain-faucet/cmd"
)

//go:generate npm run build
func main() {
	cmd.Execute()
}
