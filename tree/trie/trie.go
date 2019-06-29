package trie

import (
	"errors"
)

type Node struct {
	children map[rune]*Node
	value    interface{}
}

func New() *Node {
	return &Node{children: map[rune]*Node{}}
}

func (n *Node) Find(key string) (value interface{}, found bool) {
	next := n
	rest := []rune(key)

	for {
		if len(rest) == 0 {
			return next.value, true
		}
		r := rest[0]
		rest = []rune(rest)[1:]
		next = next.children[r]
		if next == nil {
			return nil, false
		}
	}
}

func (n *Node) Add(key string, value interface{}) error {
	next := n
	rest := []rune(key)

	for {
		if len(rest) == 0 {
			if next.value != nil {
				return errors.New("cannot add")
			}
			next.value = value
			return nil
		}
		r := rest[0]
		rest = rest[1:]

		c, ok := next.children[r]
		if ok {
			next = c
			continue
		}
		next.children[r] = &Node{
			children: map[rune]*Node{},
		}
		next = next.children[r]
	}
}
