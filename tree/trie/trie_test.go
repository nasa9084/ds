package trie

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	root := New()
	if err := root.Add("foo", "hoge"); err != nil {
		t.Fatal(err)
	}
	if err := root.Add("bar", "fuga"); err != nil {
		t.Fatal(err)
	}
	if err := root.Add("baz", "piyo"); err != nil {
		t.Fatal(err)
	}
	want := &Node{
		children: map[rune]*Node{
			'f': &Node{
				children: map[rune]*Node{
					'o': &Node{
						children: map[rune]*Node{
							'o': &Node{
								children: map[rune]*Node{},
								value:    "hoge",
							},
						},
					},
				},
			},
			'b': &Node{
				children: map[rune]*Node{
					'a': &Node{
						children: map[rune]*Node{
							'r': &Node{
								children: map[rune]*Node{},
								value:    "fuga",
							},
							'z': &Node{
								children: map[rune]*Node{},
								value:    "piyo",
							},
						},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(root, want) {
		t.Error("error")
		print(t, root, 0)
		print(t, want, 0)
		return
	}
}

func print(t *testing.T, node *Node, depth int) {
	var indent string
	for i := 0; i < depth; i++ {
		indent += " "
	}
	t.Logf("%sNode{value: %+v, children: %+v}", indent, node.value, node.children)
	for _, child := range node.children {
		print(t, child, depth+1)
	}
}

func TestFind(t *testing.T) {
	root := &Node{
		children: map[rune]*Node{
			'f': &Node{
				children: map[rune]*Node{
					'o': &Node{
						children: map[rune]*Node{
							'o': &Node{
								children: map[rune]*Node{},
								value:    "hoge",
							},
						},
					},
				},
			},
			'b': &Node{
				children: map[rune]*Node{
					'a': &Node{
						children: map[rune]*Node{
							'r': &Node{
								children: map[rune]*Node{},
								value:    "fuga",
							},
							'z': &Node{
								children: map[rune]*Node{},
								value:    "piyo",
							},
						},
					},
				},
			},
		},
	}
	got, ok := root.Find("bar")
	if !ok {
		t.Fatal("not found")
	}
	want := "fuga"
	if got != want {
		t.Errorf("%s != %s", got, want)
	}
}
