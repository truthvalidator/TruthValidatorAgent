# TruthValidator Smart Contract & Voting System

## ⛓️ The Decentralized Truth Consensus Layer

At the heart of TruthValidator lies this open-source smart contract system - a revolutionary approach to establishing truth through decentralized consensus. This component enables:

- **Community-Powered Verification:** On-chain voting for truth determination
- **AI-Assisted Judgments:** Smart contracts that integrate AI analysis
- **Tamper-Proof Records:** Immutable blockchain storage of all decisions
- **Transparent Governance:** Fully auditable verification processes

## 🗳️ How Decentralized Truth Verification Works

1. **Proposal Submission:** Users submit claims via Telegram bot
2. **AI Preliminary Analysis:** AI Agent provides initial assessment
3. **Community Voting:** Token holders vote on claim validity
4. **Consensus Recording:** Results permanently stored on blockchain
5. **Evidence Archiving:** All supporting materials stored on IPFS/Filecoin

## 🏗️ Technical Architecture

```mermaid
graph LR
    A[Telegram Bot] --> B[Smart Contract]
    B --> C[AI Analysis]
    C --> D[Community Voting]
    D --> E[IPFS Storage]
    E --> F[Result Publication]
```

**Core Components:**
- **EVM-Compatible Smart Contracts:** TruthValidatorSentientNet.sol (deployable on Filecoin FEVM, Ethereum, Polygon)
- **Golang AI Agent:** Integrates AI judgment with blockchain
- **IPFS/Filecoin Bridge:** For evidence storage
- **Telegram Bot Interface:** User interaction layer


## 🚀 Deployment Options

### Filecoin FEVM Deployment
```bash
# Deploy to Filecoin FEVM Calibration testnet
export PRIVATE_KEY="your_testnet_private_key_here"
./TruthValidatorAgent deploy --fevm

# For mainnet deployment use:
# export PRIVATE_KEY="your_mainnet_private_key_here"
# ./TruthValidatorAgent deploy --fevm --rpc https://api.node.glif.io/rpc/v1

# Run FEVM tests
make test-fevm
```

### Ethereum/Polygon Deployment
```bash
# Deploy to default network (configured in config)
./TruthValidatorAgent deploy
```

## 🌱 Contributing to Decentralized Governance

We welcome contributions to:
- Improve voting mechanisms
- Enhance smart contract security
- Develop new consensus models
- Build better integration with AI components

**Help us create a future where truth is determined by the people, for the people.**

🔗 [Smart Contract Documentation]()
📜 [Voting Protocol Specs]()
🛡️ [Security Guidelines]()

## 📜 License
Dual-licensed under:
- [Apache 2.0](../LICENSE-APACHE)
- [MIT](../LICENSE-MIT)
