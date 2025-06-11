# chainsub – Minimal Smart Contract Event Listener (CLI + Go Library)

chainsub is a lightweight, developer-friendly CLI and Go package that lets you subscribe to smart contract events on any EVM-compatible blockchain (like BNB Chain, Ethereum, Polygon, etc.).

Use it to:

- 🧪 Debug contracts by watching emitted events
- 📤 Stream events to stdout, JSON, or custom outputs
- 🛠️ Embed in Go backends, bots, or indexers

⸻

## 🚀 Quick Start (CLI)

Install

```bash
go install github.com/HarisShariff/chainsub/cmd/chainsub@latest
```

Example Usage

```bash
chainsub listen \
  --rpc https://bsc-dataseed.binance.org \
  --contract 0x55d398326f99059ff775485246999027b3197955 \
  --abi ./contracts/erc20.json \
  --event Transfer \
  --output stdout
```

## 🎯 What It Does

- Connects to any EVM-compatible RPC
- Subscribes to a contract’s event logs
- Decodes events using a provided ABI
- Streams logs to:
- Terminal (human-readable)
- JSON file
- (Extendable: Kafka, DB, etc.)

## 🔧 CLI Flags

| Flag           | Description                                                | Required | Example                                  |
| -------------- | ---------------------------------------------------------- | -------- | ---------------------------------------- |
| `--rpc`        | RPC URL of the EVM-compatible node                         | ✅       | `--rpc https://bsc-dataseed.binance.org` |
| `--contract`   | Address of the smart contract to monitor                   | ✅       | `--contract 0x55d398...`                 |
| `--abi`        | Path to the ABI JSON file for decoding events              | ✅       | `--abi ./contracts/erc20.json`           |
| `--event`      | Name of the event to subscribe to (case-sensitive)         | ✅       | `--event Transfer`                       |
| `--from-block` | Block number to start listening from (`latest` by default) | ❌       | `--from-block 12345678`                  |
| `--output`     | Output type: `stdout` or `json`                            | ❌       | `--output json`                          |

## 🧰 Use Cases

- ✅ Test if contract emits expected logs
- ✅ Monitor token transfers or custom events
- ✅ Pipe on-chain activity into your own service
- ✅ Build fast log-based devtools or bots

📦 Using as a Go Package

```go
import "github.com/HarisShariff/chainsub"

cfg := chainsub.Config{
  RPCURL:          "https://bsc-dataseed.binance.org",
  ContractAddress: "0x...",
  ABIPath:         "./erc20.json",
  EventName:       "Transfer",
  FromBlock:       "latest",
  OutputTarget:    "stdout",
}

if err := chainsub.Listen(cfg); err != nil {
  log.Fatal(err)
}
```

🧪 Example Output (stdout)

```json
{
  "event": "Transfer",
  "blockNumber": 40129382,
  "txHash": "0xabc123...",
  "from": "0xdead...",
  "to": "0xbeef...",
  "value": "1000000000000000000"
}
```

## 🗂 Project Structure

```
chainsub/
├── cmd/               # CLI app (urfave/cli)
│   └── listen.go
├── chainsub/          # Reusable Go package
│   ├── listener.go
│   ├── config.go
│   ├── decoder.go
│   └── output/
│       ├── stdout.go
│       └── json.go
├── contracts/         # Example ABIs
│   └── erc20.json
├── go.mod
├── main.go
└── README.md
```

## 📦 Install From Source

```bash
git clone https://github.com/yourname/chainsub.git
cd chainsub
go build -o chainsub ./cmd/chainsub
./chainsub listen ...
```

## 🧠 Roadmap

- Kafka output writer
- Webhook or gRPC stream support
- Replay logs (batch from block range)
- Auto-ABI fetch from BscScan

## 📜 License

MIT

## 🔗 Related

- go-ethereum – For log subscription and ABI decoding
- urfave/cli – For CLI handling
- Compatible with: BNB Chain, Ethereum, Polygon, Arbitrum, Optimism, Avalanche, etc.
