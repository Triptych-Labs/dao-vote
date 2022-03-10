package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/authl2"
	"triptychlabs.io/dao/v2/src/generated/dao"
	"triptychlabs.io/dao/v2/src/keys"
	"triptychlabs.io/dao/v2/src/triptychdao"
)

var Operation string

func init() {
	dao.SetProgramID(solana.MustPublicKeyFromBase58("E18jUpqrxp8w4u556G4CcE1jHW7iAvt7i6JH7SkGjaD8"))
	// auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))
	keys.SetupProviders()
}

func main() {
	log.Println("Starting...")
	oracle := keys.GetProvider(0)

	var daoIndex, ballotIndex int64
	type funcOp func(solana.PrivateKey, uint64)
	flag.StringVar(&Operation, "operation", "", "Operation")
	flag.Int64Var(&daoIndex, "dao_id", -1, "dao_id")
	flag.Int64Var(&ballotIndex, "ballot_id", -1, "ballot_id")
	flag.Parse()

	funcs := make([]funcOp, 0)
	switch Operation {
	case "start_authl2":
		funcs = append(
			funcs,
			SetupCloseHandler,
			authl2.Subscribe,
		)
	case "create_dao":
		if daoIndex == -1 {
			panic("specify a dao_id")
		}
		triptychdao.CreateDao(uint64(daoIndex))
	case "create_proposal":
		if daoIndex == -1 {
			panic("specify a dao_id")
		}
		if ballotIndex == -1 {
			panic("specify a ballot_id")
		}
		triptychdao.Propose(uint64(daoIndex), uint64(ballotIndex))
	}

	var wg sync.WaitGroup
	for _, f := range funcs {
		wg.Add(1)
		go func(f funcOp, wg *sync.WaitGroup) {
			f(oracle, 0)
			wg.Done()
		}(f, &wg)
	}
	wg.Wait()
	fmt.Println("Done!!!")
}

func SetupCloseHandler(dontmindme solana.PrivateKey, dontintme uint64) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("!!!!!!!------------!!!!!!!")
		fmt.Println("Use the --recover option in the next command to resume operations!!!")
		os.Exit(0)
	}()
}
