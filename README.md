# QUANTOS BLOCKCHAIN PROTOCOL
### v 0.0.1
### &copy; Quantos Developers 2022

## 1.0 On-demand blockchain with hash table

Quantos rely on hash ids and hash tables to make the blockchain more secure and efficient.

### 1.0.1 Traditional blockchains

In traditional blockchains, a full node usually needs to synchronize the full history of transactions and blocks of the blockchain.
This is (depending on the bandwidth) usually a long and tedious process.

Since Ethereum, there are options for the nodes who wants to run a light node to only synchronice parts of the chain (ie.: headers, indexes).

However it provides a limited experience for those nodes as they cannot take advantage of the same services a full node can have / offer to help the chain grow.

They are usually limited to be able to do very basic things like send / receive transactions and participate in the sharing of the parts of the blockchain they have in sync.

### 1.0.2 Quantos blockchain synchronization process

Quantos is offering a hash table based synchronization process. The hash table contains validated hash ids of the full blockchain.

- Hash ids are secure and said to be quantum safe
- They provide a more secure way to communicate with the blockchain. Nobody can communicate directly with it and temper with blocks or txs.
- Its easy to scale and shard due to their very small size.
- Its fast, on a recent test we reached 10,000 hashid per seconds (equivalent to 10,000 blocks)

... to be continued

