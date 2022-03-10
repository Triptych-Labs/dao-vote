package events

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"sync"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"triptychlabs.io/dao/v2/src/authl2/typestructs"
	"triptychlabs.io/dao/v2/src/cluster"
	"triptychlabs.io/dao/v2/src/generated/auth"
)

type AdhocEvent struct {
	InstructionName string
	EventName       string
}

var AdhocEvents = append(
	make([]AdhocEvent, 0),
	AdhocEvent{
		InstructionName: "HolderRegister",
		EventName:       "RegisterHolderEvent",
	},
)

type EventCodex struct {
	Decoder         *bin.Decoder
	Bytes           []byte
	InstructionName string
}

func NewAdhocEventListener(adhocWg *sync.WaitGroup) {
	wsClient, err := ws.Connect(context.TODO(), cluster.WSEndpoint)
	if err != nil {
		log.Println("PANIC!!!", fmt.Errorf("unable to open WebSocket Client - %w", err))
	}

	auth.SetProgramID(solana.MustPublicKeyFromBase58("9LS5eDSs36coxwkfhvUyWn7ejvVYnKAebsHTuLmky4aT"))
	sub, err := wsClient.LogsSubscribeMentions(auth.ProgramID, rpc.CommitmentConfirmed)
	if err != nil {
		log.Println(fmt.Errorf("ad hoc ws create panic: %w", err))
		adhocWg.Done()
		return
	}

	adhocWg.Add(1)
	go func() {
		state := false
		for {
			if state {
				adhocWg.Done()
				state = false
				continue
			}
			event, err := sub.Recv()
			if err != nil {
				panic(fmt.Errorf("ad hoc event recv panic: %w", err))
			}
			state = true
			adhocWg.Add(1)
			go ProcessAdhocEvents(event.Value.Logs, adhocWg)

			continue
		}
	}()

	adhocWg.Done()

}

func getEventLogs(subEventLogs []string, eventName string) []EventCodex {
	eventLogs := make([]EventCodex, 0)

	var adhocEvent *AdhocEvent = nil
	for _, logger := range subEventLogs {
		if strings.Contains(logger, "Program log: ") {
			if strings.Contains(logger, "Instruction: ") {
				logs := strings.Split(logger, "Program log: Instruction: ")
				if len(logs) <= 1 {
					continue
				}
				log := logs[1]
				for _, e := range AdhocEvents {
					if e.InstructionName == log {
						adhocEvent = &e
						break
					}
				}
				if adhocEvent == nil {
					break
				}
				continue
			}
			if adhocEvent == nil {
				break
			}
			decoder, discriminator := func() (*bin.Decoder, []byte) {
				s := fmt.Sprint("event:", adhocEvent.EventName)
				h := sha256.New()
				h.Write([]byte(s))
				discriminatorBytes := h.Sum(nil)[:8]

				logger := strings.Split(logger, "Program log: ")[1]
				eventBytes, err := base64.StdEncoding.DecodeString(logger)
				if err != nil {
					panic(err)
				}

				decoder := bin.NewBorshDecoder(eventBytes)
				if err != nil {
					panic(err)
				}

				return decoder, discriminatorBytes

			}()
			eventLogs = append(eventLogs, EventCodex{decoder, discriminator, adhocEvent.EventName})
		} else {
			continue
		}
	}
	return eventLogs
}
func ProcessAdhocEvents(subEventLogs []string, buffer *sync.WaitGroup) error {
	buffer.Add(1)
	eventLogs := getEventLogs(subEventLogs, "canbeanythingunsupported")
	adhocEventNames := func() []string {
		names := make([]string, 0)
		for _, e := range AdhocEvents {
			names = append(names, e.EventName)
		}
		return names
	}()

	/*
	   Handle AdhocEvents
	*/
	for _, logger := range eventLogs {
		switch logger.InstructionName {
		case adhocEventNames[0]:
			event := typestructs.RegisterHolderEvent{}
			err := event.UnmarshalWithDecoder(logger.Decoder, logger.Bytes)
			if err != nil {
				fmt.Println(err)
				break
			}

			// set isScheduled
			buffer.Add(1)
			go registerHolder(&event)
		}
	}
	buffer.Done()
	return nil
}

