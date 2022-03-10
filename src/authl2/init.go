package authl2

import (
	"sync"

	"github.com/gagliardetto/solana-go"
	"triptychlabs.io/dao/v2/src/authl2/events"
)

var programAuthority solana.PublicKey
var duration = int64(60 * 60 * 24 * 3)

//Subscribe - Consume Anchor Events from `smart_wallet`
func Subscribe(OWNER solana.PrivateKey, dontinme uint64) {
	var wg sync.WaitGroup

	wg.Add(1)
	go events.NewAdhocEventListener(&wg)

	wg.Wait()

}

