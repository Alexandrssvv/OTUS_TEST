package main

import (
	"INTERFACE/02-practical/bl"
	"INTERFACE/02-practical/repo"
)

func main() {
	realRepo := repo.NewRealDbRepo()
	bl.DoBigOperation(realRepo)
}
