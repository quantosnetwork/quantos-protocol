package hashid

import "sync"

type hashid interface {
	build(
		chaincode []byte,
		genesisHash []byte,
		version []byte,
		witness []byte,
		encryptedPayload []byte,
		link interface{},
		blockPayloadHash []byte,
		canDecode bool,
		containsTransactions bool,
		isFinalized bool,
		txDataHashed [][]byte,
		txCount int64,
		spaceLeft int64,
		checksum []byte,
	)
	ToProto()
	FromProto()
	Decode()
	Encode()
	Encrypt()
	Decrypt()
	GetAssetByLink(linkID []byte)
	Validate()
	ToHash()
	ToHex()
	ToString()
}

type HashData struct {
	chaincode            []byte
	genesisHash          []byte
	version              []byte
	witness              []byte
	encryptedPayload     []byte
	link                 interface{}
	blockPayloadHash     []byte
	timestamp            int64
	canDecode            bool
	containsTransactions bool
	isFinalized          bool
	txDataHashed         [][]byte
	txCount              int64
	spaceLeft            int64
	checksum             []byte
}

type HashID struct {
	mu sync.Mutex
	HashData
}

func (h *HashID) build(
	chaincode []byte,
	genesisHash []byte,
	version []byte,
	witness []byte,
	encryptedPayload []byte,
	link interface{},
	blockPayloadHash []byte,
	canDecode bool,
	containsTransactions bool,
	isFinalized bool,
	txDataHashed [][]byte,
	txCount int64,
	spaceLeft int64,
	checksum []byte,
) {
	h.chaincode = chaincode
	h.genesisHash = genesisHash
	h.version = version
	h.witness = witness
	h.encryptedPayload = encryptedPayload
	h.link = link
	h.blockPayloadHash = blockPayloadHash
	h.canDecode = canDecode
	h.containsTransactions = containsTransactions
	h.isFinalized = isFinalized
	h.txDataHashed = txDataHashed
	h.txCount = txCount
	h.spaceLeft = spaceLeft
	h.checksum = checksum
}
