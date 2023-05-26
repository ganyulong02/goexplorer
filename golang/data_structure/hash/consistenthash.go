package hash

import (
	"fmt"
	"github.com/ganyulong02/goexplorer/golang/data_structure/lang"
	"sort"
	"strconv"
	"sync"
)

// https://medium.com/codex/understanding-and-implementing-consistent-hash-algorithm-e53a35afa428
// https://github.com/zeromicro/go-zero/blob/master/core/hash/consistenthash.go

/**
Usage scenarios
1. Distributed caching. Can build a cache proxy on a storage system like redis cluster and control the routing freely.
For this routing rule, can use the consistent hash algorithm.
2. Service discovery
3. Distributed scheduling of tasks
*/

const (
	// TopWeight is the top weight that one entry might set.
	TopWeight = 100

	minReplicas = 100
	prime       = 16777619
)

type (
	// Func defines that hash method
	Func func(data []byte) uint64

	// A consistentHash is a ring hash implementation
	ConsistentHash struct {
		hashFunc Func                            // hash function
		replicas int                             // virtual node amplification factor
		keys     []uint64                        // store virtual node hash
		ring     map[uint64][]any                // virtual node to actual node correspondence
		nodes    map[string]lang.PlaceholderType // actual node storage [easy to find quickly, so use map]
		lock     sync.RWMutex                    // A reader/writer mutual exclusion lock
	}
)

// NewConsistentHash returns a ConsistentHash
func NewConsistentHash() *ConsistentHash {
	return NewCustomConsistentHash(minReplicas, Hash)
}

// NewCustomConsistentHash returns a ConsistentHash with replicas and hash func.
func NewCustomConsistentHash(replicas int, fn Func) *ConsistentHash {
	if replicas < minReplicas {
		replicas = minReplicas
	}

	if fn == nil {
		fn = Hash
	}

	return &ConsistentHash{
		hashFunc: fn,
		replicas: replicas,
		ring:     make(map[uint64][]any),
		nodes:    make(map[string]lang.PlaceholderType),
	}
}

func repr(node any) string {
	return lang.Repr(node)
}

func innerRepr(node any) string {
	return fmt.Sprintf("%d:%v", prime, node)
}

func (h *ConsistentHash) removeRingNode(hash uint64, nodeRepr string) {
	if nodes, ok := h.ring[hash]; ok {
		newNodes := nodes[:0]
		for _, x := range nodes {
			if repr(x) != nodeRepr {
				newNodes = append(newNodes, x)
			}
		}
		if len(newNodes) > 0 {
			h.ring[hash] = newNodes
		} else {
			delete(h.ring, hash)
		}
	}
}

// Get returns the corresponding node from h base on the given v.
func (h *ConsistentHash) Get(v any) (any, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	if len(h.ring) == 0 {
		return nil, false
	}

	// Calculate the hash of key
	hash := h.hashFunc([]byte(repr(v)))
	// Find the index of first matching virtual node and fetch the corresponding h.keys[index]:
	// virtual node hash value
	index := sort.Search(len(h.keys), func(i int) bool {
		return h.keys[i] >= hash
	}) % len(h.keys)

	// Go to this ring and find an actual node that matches it
	nodes := h.ring[h.keys[index]]
	switch len(nodes) {
	case 0:
		return nil, false
	case 1:
		return nodes[0], true
	default:
		innerIndex := h.hashFunc([]byte(innerRepr(v)))
		pos := int(innerIndex % uint64(len(nodes)))
		return nodes[pos], true
	}
}

// Add adds the node with the number of h.replicas,
// the later call will overwrite the replicas of the former calls
func (h *ConsistentHash) Add(node any) {
	h.AddWithReplicas(node, h.replicas)
}

func (h *ConsistentHash) addNode(nodeRepr string) {
	h.nodes[nodeRepr] = lang.Placeholder
}

func (h *ConsistentHash) containsNode(nodeRepr string) bool {
	_, ok := h.nodes[nodeRepr]
	return ok
}

func (h *ConsistentHash) removeNode(nodeRepr string) {
	delete(h.nodes, nodeRepr)
}

// AddWithReplicas adds the node with the number of replicas,
// replicas will be truncated to h.replicas if it's larger than h.replicas,
// the later call will overwrite the replicas of the former calls
func (h *ConsistentHash) AddWithReplicas(node any, replicas int) {
	h.Remove(node)

	if replicas > h.replicas {
		replicas = h.replicas
	}

	nodeRepr := repr(node)
	h.lock.Lock()
	defer h.lock.Unlock()
	h.addNode(nodeRepr)

	for i := 0; i < replicas; i++ {
		hash := h.hashFunc([]byte(nodeRepr + strconv.Itoa(i)))
		h.keys = append(h.keys, hash)
		h.ring[hash] = append(h.ring[hash], node)
	}

	sort.Slice(h.keys, func(i, j int) bool {
		return h.keys[i] < h.keys[j]
	})
}

// AddWithWeight adds the node with weight, the weight can be 1 to 100, indicates the percent,
// the later call will overwrite the replicas of the former calls
func (h *ConsistentHash) AddWithWeight(node any, weight int) {
	// don't need to make sure weight not larger than TopWeight,
	// because AddWithReplicas makes sure replicas cannot be larger than h.replicas
	replicas := h.replicas * weight / TopWeight
	h.AddWithReplicas(node, replicas)
}

// Remove removes the given node from h.
func (h *ConsistentHash) Remove(node any) {
	nodeRepr := repr(node)

	h.lock.Lock()
	defer h.lock.Unlock()

	if !h.containsNode(nodeRepr) {
		return
	}

	for i := 0; i < h.replicas; i++ {
		hash := h.hashFunc([]byte(nodeRepr + strconv.Itoa(i)))
		index := sort.Search(len(h.keys), func(i int) bool {
			return h.keys[i] >= hash
		})
		if index < len(h.keys) && h.keys[index] == hash {
			h.keys = append(h.keys[:index], h.keys[index+1:]...)
		}
		h.removeRingNode(hash, nodeRepr)
	}
}
