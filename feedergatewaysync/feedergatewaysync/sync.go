package feedergatewaysync

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type SyncManager struct {
	FeederGatewayURL string
	startBlock       int
	endBlock         int
	httpClient       *http.Client
	storage          *Storage
}

func NewSyncManager(feederGatewayURL string, startBlock int, endBlock int) *SyncManager {
	return &SyncManager{
		FeederGatewayURL: feederGatewayURL,
		startBlock:       startBlock,
		endBlock:         endBlock,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		storage: NewStorage("fedder_gateway_sync.db"),
	}
}

func (sm *SyncManager) SyncBlocks() {
	for i := sm.startBlock; i < sm.endBlock; i++ {
		url := fmt.Sprintf("%s/feeder_gateway/get_block?blockNumber=%d", sm.FeederGatewayURL, i)
		resp, err := sm.httpClient.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		block := &Block{}
		err = json.Unmarshal(body, block)
		if err != nil {
			log.Fatal(err)
		}
		sm.storage.Set([]byte(fmt.Sprintf("block_%d", i)), block)
		log.Printf("Block %d synced", i)
	}
}
