/*
 * Copyright (c) 2019
 * InterDigital Communications, Inc.
 * All rights reserved.
 *
 * The information provided herein is the proprietary and confidential
 * information of InterDigital Communications, Inc.
 */

package bws

import (
	"errors"
	"time"
	"sync"
	"strconv"

	ceModel "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-ctrl-engine-model"
	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
	redis "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis"

)

const redisAddr string = "meep-redis-master:6379"
var BW_SHARING_CONTROLS_DB = 0

// BwSharing - 
type BwSharing struct {
	name        string
	isStarted   bool
	isReady     bool
	ticker      *time.Ticker
	rcCtrlEng *redis.Connector
	mutex sync.Mutex
	updateFilterCB func(string, string, float64)
	applyFilterCB func()
	config ConfigurationAttributes
}

type ConfigurationAttributes struct {
	Action                    string
	RecalculationPeriod       int
	LogVerbose                bool
	EnableTier1               bool
	EnableTier2		  bool
	EnableTier3		  bool
}

// NewBwSharing - Create, Initialize and connect
func NewBwSharing(name string, updateFilterRule func(string, string, float64), applyFilterRule func()) (bw *BwSharing, err error) {
	if name == "" {
		err = errors.New("Missing bwSharing name")
		log.Error(err)
		return nil, err
	}

        bw = new(BwSharing)
        bw.name = name
        bw.isStarted = false
        bw.isReady = false
        log.Debug("BwSharing created ", bw.name)

        bw.rcCtrlEng, err = redis.NewConnector(redisAddr, BW_SHARING_CONTROLS_DB)
        if err != nil {
                log.Error("Failed connection to redis DB. Error: ", err)
                return nil, err
        }
        log.Info("Connected to redis DB")

        // Subscribe to Pub-Sub events for MEEP Controller
        // NOTE: Current implementation is RedisDB Pub-Sub
        err = bw.rcCtrlEng.Subscribe(channelBwSharingControls)
        if err != nil {
                log.Error("Failed to subscribe to Pub/Sub events on channelBwSharingControls. Error: ", err)
                return nil, err
        }

        go bw.Run()
	
	bw.updateFilterCB = updateFilterRule
	bw.applyFilterCB = applyFilterRule
	//get values from the DB, or defaults
	initDefaultConfigAttributes()
	bw.UpdateControls()
	return bw, nil
}

// Run - Listening event
func (bw *BwSharing) Run() {
        // Listen for subscribed events. Provide event handler method.
        _ = bw.rcCtrlEng.Listen(bw.eventHandler)
}

func (bw *BwSharing) eventHandler(channel string, payload string) {
        // Handle Message according to Rx Channel
        switch channel {

	case channelBwSharingControls:
                log.Debug("Event received on channel: ", channelBwSharingControls)
		bw.UpdateControls()
        default:
                log.Warn("Unsupported channel")
        }
}

func (bw *BwSharing) ParseScenarioUpdate(scenario ceModel.Scenario) {
	if bw.isStarted {
	        // Parse scenario
	        bw.mutex.Lock()
	        bw.isReady = false
	        parseScenario(scenario)
	        bw.isReady = true
	        bw.mutex.Unlock()
	}
}

func (bw *BwSharing) updateFilter(dst string, src string, value float64) {
	bw.updateFilterCB(dst, src, value)
}

func (bw *BwSharing) applyFilter() {
        bw.applyFilterCB()
}

func (bw *BwSharing) UpdateControls() {

        var controls = make(map[string]interface{})

        keyName := bwSharingControls

	err := bw.rcCtrlEng.ForEachEntry(keyName, bw.getControlsEntryHandler, controls)
        if err != nil {
                log.Error("Failed to get entries: ", err)
                return
        }
}

func (bw *BwSharing) getControlsEntryHandler(key string, fields map[string]string, userData interface{}) (err error) {

	actionName := ""
	tickerPeriod := defaultTickerPeriod
	log := false
	enableTier1 := false
        enableTier2 := false
        enableTier3 := false

        for fieldName, fieldValue := range fields {
		switch(fieldName) {
		case "action":
			actionName = fieldValue
		case "recalculationPeriod":
			tickerPeriod, err = strconv.Atoi(fieldValue)
			if err != nil {
				tickerPeriod = defaultTickerPeriod
			} 
		case "logVerbose":
			if "yes" == fieldValue {
				log = true
			}
		default:
			updateDefaultConfigAttributes(fieldName, fieldValue)
		}
	}

	bw.config.Action = actionName
	bw.config.RecalculationPeriod = tickerPeriod
	bw.config.LogVerbose = log
	bw.config.EnableTier1 = enableTier1
	bw.config.EnableTier2 = enableTier2
	bw.config.EnableTier3 = enableTier3

	//for debug
	bw.ApplyAction()
	
        return nil
}

func (bw *BwSharing) ApplyAction() (err error) {
        switch(bw.config.Action) {
        case "start":
                if !bw.isStarted {
                        bw.Start()
                }
        case "stop":
                if bw.isStarted {
                        bw.Stop()
                }
        default:
        }

        return nil
}

// Start - starts bwSharing distribution calculations
func (bw *BwSharing) Start() (err error) {
	bw.isStarted = true
	bw.isReady = true
	bw.ticker = time.NewTicker(time.Duration(bw.config.RecalculationPeriod) * time.Millisecond)

	allocateBandwidthSharing()
	//bw.ParseScenarioUpdate()
        go func() {
                for range bw.ticker.C {

			//if it's not ready... we skip the whole ticker
                        if bw.isReady {
				bw.mutex.Lock()
				bw.isReady = false
				tickerFunction(bw.rcCtrlEng, bw.config.LogVerbose, bw.updateFilterCB, bw.applyFilterCB)
				bw.isReady = true
				bw.mutex.Unlock()
                        }
                }
        }()

	return nil
}

// Stop - stops bwSharing computation
func (bw *BwSharing) Stop() {
	if bw.isStarted {
		bw.ticker.Stop()
		log.Debug("BwSharing computation stopped ", bw.name)
		bw.isStarted = false
		bw.isReady = false
		cleanUp()
	}
}
