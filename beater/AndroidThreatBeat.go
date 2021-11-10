package beater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/atlantis0/androidthreatbeat/config"
)

// Threat - threat object
type Threat struct {
	Category string `json:"category"`
	Vector   string `json:"type"`
	Severity string `json:"severity"`
}

// androidthreatbeat configuration.
type androidthreatbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of androidthreatbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &androidthreatbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts androidthreatbeat.
func (bt *androidthreatbeat) Run(b *beat.Beat) error {
	logp.Info("androidthreatbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		//bt.config.ThreatFilePath
		// TODO - read json file
		threatFileData, err := ioutil.ReadFile(bt.config.ThreatFilePath)
		if err != nil {
			return err
		}
		var threat []Threat
		decodeErr := json.Unmarshal(threatFileData, &threat)
		if decodeErr != nil {
			return decodeErr
		}
		events := make([]beat.Event, len(threat))
		for i := 0; i < len(threat); i++ {
			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type":     b.Info.Name,
					"category": threat[i].Category,
					"severity": threat[i].Severity,
					"vector":   threat[i].Vector,
				},
			}
			events[i] = event
		}

		bt.client.PublishAll(events)
		logp.Info("Event sent")
		counter++
	}
}

// Stop stops androidthreatbeat.
func (bt *androidthreatbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
