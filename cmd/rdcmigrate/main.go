package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wasawaz/go-blockchain/database"
)

func main() {
	state, err := database.NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	block0 := database.NewBlock(database.Hash{}, uint64(time.Now().Unix()), []database.Tx{
		database.NewTx("roodic", "roodic", 3, ""),
		database.NewTx("roodic", "roodic", 700, "reward"),
	})
	state.AddBlock(block0)
	block0hash, _ := state.Persist()

	block1 := database.NewBlock(block0hash, uint64(time.Now().Unix()), []database.Tx{
		database.NewTx("roodic", "babayaga", 2000, ""),
		database.NewTx("roodic", "roodic", 100, "reward"),
		database.NewTx("babayaga", "roodic", 1, ""),
		database.NewTx("babayaga", "caesar", 1000, ""),
		database.NewTx("babayaga", "roodic", 50, ""),
		database.NewTx("roodic", "roodic", 600, "reward"),
	})
	state.AddBlock(block1)
	state.Persist()

}
