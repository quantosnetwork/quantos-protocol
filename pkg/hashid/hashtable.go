package hashid

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/sha3"
	"math/big"
)

const MaxHashMapRecordsAtTheSameTime = 10

type HashMap struct {
	Data []*Node
}

func NewDict() *HashMap {
	return &HashMap{Data: make([]*Node, MaxHashMapRecordsAtTheSameTime)}
}

type Node struct {
	key   string
	value interface{}
	next  *Node
}

var hashFunction = crypto.SHA3_256

func hash(key string) (hash uint64) {
	hash = 0
	buf := []byte(key)

	h := sha3.Sum512(buf)

	hx := hex.EncodeToString(h[:])
	fmt.Println("%x", hx)
	b := new(big.Int)
	b.SetString(hx, 10)
	hash = b.Uint64()
	return
}

func getIndex(key string) (index int) {
	return int(hash(key)) % MaxHashMapRecordsAtTheSameTime
}

func (h *HashMap) Insert(key string, value interface{}) {
	index := getIndex(key)
	spew.Dump(index)
	if h.Data[index] == nil {
		h.Data[index] = &Node{key: key, value: value}
	} else {
		starting_node := h.Data[index]
		for ; starting_node.next != nil; starting_node = starting_node.next {
			if starting_node.key == key {
				starting_node.value = value
				return
			}
		}
		starting_node.next = &Node{key: key, value: value}

	}
}

func (h *HashMap) Get(key string) (interface{}, bool) {
	index := getIndex(key)
	if h.Data[index] != nil {
		// key is on this index, but might be somewhere in linked list
		starting_node := h.Data[index]
		for ; ; starting_node = starting_node.next {
			if starting_node.key == key {
				// key matched
				return starting_node.value, true
			}

			if starting_node.next == nil {
				break
			}
		}
	}

	// key does not exists
	return "", false
}
