# Genesis hash calculator

## Requirements

1. `go 1.19` installed

## Run

### Binaries

Please [download](https://github.com/mxmar/eth1genesis/releases/tag/v1.0.0) and use as a standard binary (no flags, just execute: `./eth1genesis-v1.0.0-<linux/darwin>-amd64`)

### Go run
1. Copy your `genesis.json` to working directory.
2. Execute: `go run .`
3. Check your terminal output for genesis block hash and stateRoot.

## Build from source

1. Execute: `go build .`

NOTE: On mac, it likely won't start right away, you then need to got to System Settings -> Privacy & Security and allow the `eth1genesis` binary.

### Run build

1. Execute: `./eth1genesis `
