package main

import (
"bytes"
"encoding/json"
"fmt"
"github.com/ethereum/go-ethereum/core"
"github.com/ethereum/go-ethereum/core/types"
"github.com/protolambda/zrnt/eth2/configs"
"io/ioutil"
"os"
)

type Eth1GenesisCmd struct {
	configs.SpecOptions `ask:"."`
	Eth1Config          string `ask:"--eth1-config" help:"Path to config JSON for eth1. No transition yet if empty."`
}

func (c *Eth1GenesisCmd) Help() string {
	return "Get eth1 genesis block hash & stateRoot."
}

func (c *Eth1GenesisCmd) Default() {
	c.SpecOptions.Default()
	c.Eth1Config = "genesis.json"
}

func main() {
	fmt.Println("Starting!")
	cmd := &Eth1GenesisCmd{}
	cmd.Default()

	var eth1Block *types.Block
	var eth1Genesis *core.Genesis

	// Load the Eth1 block from the Eth1 genesis config
	eth1Genesis, err := loadEth1Genesis(cmd.Eth1Config)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("failed to load eth1 genesis: %v", err))
		os.Exit(1)
	}

	// Generate genesis block from the loaded config
	eth1Block = eth1Genesis.ToBlock()

	fmt.Printf("Your genesis block hash is: 0x%x\n", eth1Block.Hash())
	fmt.Printf("Your genesis block hash stateRoot is: 0x%x\n", eth1Block.Root())

	fmt.Println("Done!")
	os.Exit(0)
}

func loadEth1Genesis(configPath string) (*core.Genesis, error) {
	eth1ConfData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read eth1 config file: %v", err)
	}
	var eth1Genesis core.Genesis
	if err := json.NewDecoder(bytes.NewReader(eth1ConfData)).Decode(&eth1Genesis); err != nil {
		return nil, fmt.Errorf("failed to decode eth1 config file: %v", err)
	}
	return &eth1Genesis, nil
}

