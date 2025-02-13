
# TruthValidatorAgent

**Decentralized Information Verification System powered by AI Agent, Trusted Storage and Blockchain.**

## Project Overview

TruthValidatorAgent is an open-source system designed to combat misinformation by providing a transparent and efficient way to verify information. It combines **smart contracts**, **AI Agents (AI RAG-based)**, and **on-chain voting** to assess the authenticity of user-submitted proposals.

Users submit proposals to blockchain smart contracts through **robots**. The system then:

1. **Analyzes with AI:** An AI Agent uses Retrieval-Augmented Generation (RAG) to analyze the proposal and generate a preliminary judgment.
2. **Verifies by Community AI Agent:** The proposal is put to an **on-chain vote** for community consensus.
3. **Delivers Transparent Results:** Voting outcomes are recorded on the blockchain and communicated back to users via the robots.
4. **Stores Evidence on Filecoin:All voting evidence is decentralized and securely stored on the Filecoin network, ensuring data integrity and availability.** 

**Core Values:**

* **Enhanced Credibility:** Reduces misinformation spread.
* **Strengthened Trust:** Provides transparent verification.
* **Reduced Costs:** Automates verification using AI and community.
* **Data Integrity:** Ensures evidence is tamper-proof and accessible via **Filecoin**.
* **User Privacy:** Utilizes decentralized blockchain.
* **Community Governance:** Empowers community participation in verification.

**Key Applications:**

* Social Media Rumor Debunking
* News Verification
* E-commerce Anti-Counterfeiting
* Public Affairs Clarification
* Academic Research Validation

## System Architecture

![System Architecture](./docs/imgs/arch.png)
<div style="display: none;">
```mermaid
graph TD
    A[User Client] -->|1.Submit Proposal| B[Smart Contract]
    B -->|2.Trigger ProposalSubmitted Event| C[AI Agent & Vote Node]
    C -->|3.Call AI Model & RAG| D[AI Model & RAG]
    D -->|4.Return Judgment Result| C
    C -->|5.Generate Voting Evidence| E[Filecoin Storage Network]
    E -->|6.Store Evidence & Return CID| C
    C -->|7.Submit Vote with evidence to Smart Contract| B
    B -->|8.Trigger Finalized Event| F[Telegram Bot]
    F -->|9.Response Result| G[Telegram Bot]
    G -->|10.User Reply| A

    classDef contract fill:#f9f,stroke:#333;
    classDef agent fill:#ccf,stroke:#333;
    classDef model fill:#9f9,stroke:#333;
    classDef bot fill:#f96,stroke:#333;
    classDef storage fill:#aaf,stroke:#333;
    class B contract
    class C agent
    class D model
    class E storage
    class G bot
</div>

**Components:**

* **Smart Contract (Solidity):** Manages proposals and voting on EVM blockchains.
* **Filecoin storage and retrieval(Filecoin lassie & Storacha):** Decentralized storage for voting evidence, ensuring data integrity and accessibility.
* **AI Agent (Golang):** Uses RAG-based analysis for judgment and participates in voting.
* **AI Search Module (Golang):** Retrieves web information using SearxNG.
* **Telegram Bot (Golang):** Serves as the user interface for submissions and results.

## Technologies

* **AI:** RAG, LLM, Metasearch API
* **Filecoin:** Decentralized storage for voting evidence(Filecoin lassie & Storacha)
* **Blockchain:** EVM-compatible, Solidity
* **Telegram Bot API:** `telegram-bot` (Golang)

## Integration with Filecoin storage

* **Data Integrity:** Voting evidence is stored in a decentralized, tamper-proof manner.
* **Persistent Storage:** Ensures long-term availability of evidence, safeguarding against data loss.
* **Transparency:** Anyone can verify the authenticity of voting evidence via Filecoin's decentralized network.
* **Decentralized Governance:** Aligns with the system's philosophy of decentralization and community empowerment.
* **Enhanced Security:** Protects against censorship and single-point failures.


## 🚀Join us in building a more trustworthy information ecosystem!