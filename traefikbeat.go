package main

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
)

func main() {
	fmt.Printf("Hello, world.\n")
	beat.Run("traefikbeat", "0.0.0", New())
}

type Traefikbeat struct {
	isAlive bool
	client  publisher.Client
}

func New() *Traefikbeat {
	return &Traefikbeat{}
}

func (tb *Traefikbeat) Config(b *beat.Beat) error {
	logp.Info("%s", "Traefikbeat init")
	return nil
}

func (tb *Traefikbeat) Setup(b *beat.Beat) error {
	tb.isAlive = true
	tb.client = b.Publisher.Connect()
	return nil
}

func (tb *Traefikbeat) Run(b *beat.Beat) error {
	for tb.isAlive {
		time.Sleep(1000)
		var i int = 0
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       "traefik",
			"count":      i,
		}
		tb.client.PublishEvent(event)
		i++
	}
	return nil
}

func (tb *Traefikbeat) Cleanup(b *beat.Beat) error {
	logp.Info("Traefik cleanup")
	return nil
}

func (tb *Traefikbeat) Stop() {
	tb.isAlive = false
}
