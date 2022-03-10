package main

import (
	"syscall/js"

	"triptychlabs.io/dao/v2/wasm/integrations"
)

func main() {
	done := make(chan struct{})

	global := js.Global()

	registerFunc := js.FuncOf(integrations.RegisterAuth)
	defer registerFunc.Release()
	global.Set("register", registerFunc)

	getProgramIDFunc := js.FuncOf(integrations.GetProgramID)
	defer getProgramIDFunc.Release()
	global.Set("getProgramID", getProgramIDFunc)

	certifyFunc := js.FuncOf(integrations.Certify)
	defer certifyFunc.Release()
	global.Set("certify", certifyFunc)

	fetchProposalsFunc := js.FuncOf(integrations.FetchProposals)
	defer fetchProposalsFunc.Release()
	global.Set("fetchProposals", fetchProposalsFunc)

	<-done
}
