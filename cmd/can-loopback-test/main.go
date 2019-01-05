package main

import (
	"flag"
	"github.com/brutella/can"
	"log"
	"time"
)

const frameID = 0x000FFFFE // any value, here is Volvo Vida diagnostic tool ID (leading extra zero-values in hex notation are higher priority in can bus)

var ifName string
var tickerDuration time.Duration

func main() {
	flag.StringVar(&ifName, "interface", "can0", "linux-can interface name, for example: can0")
	flag.DurationVar(&tickerDuration, "interval", time.Second, "Interval between debug frames send")
	flag.Parse()

	log.Printf("can-loopback-test, running on %s interface\n", ifName)

	bus, err := can.NewBusForInterfaceWithName(ifName)
	if err != nil {
		log.Panicln(err)
		return
	}

	bus.SubscribeFunc(func (frame can.Frame) {
		log.Printf("Received frame: %v\n", frame)
	})

	go func() {
		tick := time.NewTicker(tickerDuration)

		for {
			select {
			case <-tick.C:
				frame := can.Frame{
					ID:     frameID,
					Length: 1,
					Flags:  0,
					Res0:   0,
					Res1:   0,
					Data:   [8]uint8{0x05},
				}

				if err := bus.Publish(frame); err != nil {
					log.Panicln(err)
				}

				log.Printf("Sent frame: %v\n", frame)
			}
		}
	}()

	if err := bus.ConnectAndPublish(); err != nil {
		log.Panicln(err)
		return
	}
}