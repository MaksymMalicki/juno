package trienode

import (
	"testing"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/core/trie2/utils"
	"github.com/stretchr/testify/require"
)

func TestNodeSet(t *testing.T) {
	t.Run("new node set", func(t *testing.T) {
		ns := NewNodeSet(felt.Zero)
		require.Equal(t, felt.Zero, ns.Owner)
		require.Empty(t, ns.Nodes)
		require.Zero(t, ns.updates)
		require.Zero(t, ns.deletes)
	})

	t.Run("add nodes", func(t *testing.T) {
		ns := NewNodeSet(felt.Zero)

		// Add a regular node
		key1 := utils.NewBitArray(8, 0xFF)
		node1 := NewNode(felt.Zero, []byte{1, 2, 3})
		ns.Add(key1, node1)
		require.Equal(t, 1, ns.updates)
		require.Equal(t, 0, ns.deletes)

		// Add a deleted node
		key2 := utils.NewBitArray(8, 0xAA)
		node2 := NewDeleted()
		ns.Add(key2, node2)
		require.Equal(t, 1, ns.updates)
		require.Equal(t, 1, ns.deletes)

		// Verify nodes are stored correctly
		require.Equal(t, node1, ns.Nodes[key1])
		require.Equal(t, node2, ns.Nodes[key2])
	})

	t.Run("merge sets", func(t *testing.T) {
		ns1 := NewNodeSet(felt.Zero)
		ns2 := NewNodeSet(felt.Zero)

		// Add nodes to first set
		key1 := utils.NewBitArray(8, 0xFF)
		node1 := NewNode(felt.Zero, []byte{1, 2, 3})
		ns1.Add(key1, node1)

		// Add nodes to second set
		key2 := utils.NewBitArray(8, 0xAA)
		node2 := NewDeleted()
		ns2.Add(key2, node2)

		// Merge sets
		err := ns1.MergeSet(ns2)
		require.NoError(t, err)

		// Verify merged state
		require.Equal(t, 2, len(ns1.Nodes))
		require.Equal(t, node1, ns1.Nodes[key1])
		require.Equal(t, node2, ns1.Nodes[key2])
		require.Equal(t, 1, ns1.updates)
		require.Equal(t, 1, ns1.deletes)
	})

	t.Run("merge with different owners", func(t *testing.T) {
		owner1 := new(felt.Felt).SetUint64(123)
		owner2 := new(felt.Felt).SetUint64(456)
		ns1 := NewNodeSet(*owner1)
		ns2 := NewNodeSet(*owner2)

		err := ns1.MergeSet(ns2)
		require.Error(t, err)
	})

	t.Run("merge map", func(t *testing.T) {
		owner := new(felt.Felt).SetUint64(123)
		ns := NewNodeSet(*owner)

		// Create a map to merge
		nodes := make(map[utils.BitArray]*Node)
		key1 := utils.NewBitArray(8, 0xFF)
		node1 := NewNode(felt.Zero, []byte{1, 2, 3})
		nodes[key1] = node1

		// Merge map
		err := ns.Merge(*owner, nodes)
		require.NoError(t, err)

		// Verify merged state
		require.Equal(t, 1, len(ns.Nodes))
		require.Equal(t, node1, ns.Nodes[key1])
		require.Equal(t, 1, ns.updates)
		require.Equal(t, 0, ns.deletes)
	})

	t.Run("foreach", func(t *testing.T) {
		ns := NewNodeSet(felt.Zero)

		// Add nodes in random order
		keys := []utils.BitArray{
			utils.NewBitArray(8, 0xFF),
			utils.NewBitArray(8, 0xAA),
			utils.NewBitArray(8, 0x55),
		}
		for _, key := range keys {
			ns.Add(key, NewNode(felt.Zero, []byte{1}))
		}

		t.Run("ascending order", func(t *testing.T) {
			var visited []utils.BitArray
			ns.ForEach(false, func(key utils.BitArray, node *Node) {
				visited = append(visited, key)
			})

			// Verify ascending order
			for i := 1; i < len(visited); i++ {
				require.True(t, visited[i-1].Cmp(&visited[i]) < 0)
			}
		})

		t.Run("descending order", func(t *testing.T) {
			var visited []utils.BitArray
			ns.ForEach(true, func(key utils.BitArray, node *Node) {
				visited = append(visited, key)
			})

			// Verify descending order
			for i := 1; i < len(visited); i++ {
				require.True(t, visited[i-1].Cmp(&visited[i]) > 0)
			}
		})
	})
}

func TestNode(t *testing.T) {
	t.Run("new node", func(t *testing.T) {
		hash := new(felt.Felt).SetUint64(123)
		blob := []byte{1, 2, 3}
		node := NewNode(*hash, blob)

		require.Equal(t, *hash, node.hash)
		require.Equal(t, blob, node.blob)
		require.False(t, node.IsDeleted())
	})

	t.Run("new deleted node", func(t *testing.T) {
		node := NewDeleted()
		require.True(t, node.IsDeleted())
		require.Equal(t, felt.Zero, node.hash)
		require.Nil(t, node.blob)
	})

	t.Run("is deleted", func(t *testing.T) {
		tests := []struct {
			name     string
			blob     []byte
			expected bool
		}{
			{
				name:     "nil blob",
				blob:     nil,
				expected: true,
			},
			{
				name:     "empty blob",
				blob:     []byte{},
				expected: true,
			},
			{
				name:     "non-empty blob",
				blob:     []byte{1, 2, 3},
				expected: false,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				node := NewNode(felt.Zero, test.blob)
				require.Equal(t, test.expected, node.IsDeleted())
			})
		}
	})
}
