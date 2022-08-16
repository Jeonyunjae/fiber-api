// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decimaltree

import (
	"github.com/jeonyunjae/fiber-api/models"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	NodeIndex, BrancdIndex int
	Node                   map[uint]Tree
	Value                  map[uint]models.PositionAddressInfo
}

func (tree *Tree) TreeInit(t *Tree, nodeIndex int, nodeTotalIndex int) *Tree {
	if nodeIndex == nodeTotalIndex {
		return nil
	}
	for i := 0; i < 10; i++ {
		t.Node = make(map[uint]Tree)
		t.NodeIndex = nodeIndex
		t.BrancdIndex = i
		t.TreeInit(t, nodeIndex+1, nodeTotalIndex)
	}
	return t
}

func (tree *Tree) Insert(t *Tree, PositionAddressInfo models.PositionAddressInfo) *Tree {

	// findvar := fmt.Sprintf("%f", PositionAddressInfo.Lat*1000000)
	// for i := 0; i < 10; i++ {
	// 	keys := make([]uint, 0, len(t.Node))
	// 	for k := range t.Node {
	// 		keys = append(keys, k)
	// 	}

	// }
	return tree
}

func findvalue(t *Tree, PositionAddressInfo models.PositionAddressInfo) {

}

// func (tree *Tree) Insert(t *Tree, models.PositionAddressInfo) *Tree {
// 	if nodeIndex == nodeTotalIndex {
// 		return nil
// 	}
// 	for i := 0; i < 10; i++ {
// 		tree.TreeInit(t, nodeIndex+1, nodeTotalIndex)
// 	}
// 	return t
// }

// // New returns a new, random binary tree holding the values k, 2k, ..., 10k.
// func New(datas []models.PositionAddressInfo) *Tree {
// 	var t *Tree
// 	for _, data := range datas {
// 		t = insert(t, data)
// 	}
// 	return t
// }

// func (t *Tree)insert(t *Tree, models.PositionAddressInfo) *Tree {
// 	if t == nil {
// 		return &Tree{nil, v, nil}
// 	}
// 	if v < t.Value {
// 		t.Left = insert(t.Left, v)
// 	} else {
// 		t.Right = insert(t.Right, v)
// 	}
// 	return t
// }

// func (t *Tree) String() string {
// 	if t == nil {
// 		return "()"
// 	}
// 	s := ""
// 	if t.Left != nil {
// 		s += t.Left.String() + " "
// 	}
// 	s += fmt.Sprint(t.Value)
// 	if t.Right != nil {
// 		s += " " + t.Right.String()
// 	}
// 	return "(" + s + ")"
// }
