package main

type Trie struct {
	root *Node
}

type alphabet byte

type Node struct {
	isEnd    bool
	char     alphabet
	children []*Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{isEnd: true, children: nil}}
}
func (t *Trie) Insert(s string) {
	if t.root.isEnd == true {
		// the root is not the end now
		t.root.isEnd = false
		// insert char
		t.root.insert(s)
	}
	// trie exist
	// search alphabet
}
func (n *Node) insert(s string) {
	if len(s) == 0 {
		return
	}

	// check existance of s[0] in n.children
	if i, ok := n.getIndex(alphabet(s[0])); !ok {
		// s[0] not in n.children
		// create char target in n.children
		n.isEnd = false
		n.children = append(n.children, &Node{isEnd: true, char: alphabet(s[0]), children: nil})
		n.children[i].insert(s[1:])
	} else {
		// s[0] is in n.children
		n.children[i].insert(s[1:])
	}
}

// return the index of character in n.children,
// if character c is in n.children return index and true
// if not exist, return 0 and false
func (n *Node) getIndex(c alphabet) (index int, ok bool) {
	for i, v := range n.children {
		if c == v.char {
			// target found
			return i, true

		}
	}
	return 0, false
}
func (t *Trie) Search() {}

func main() {
	var t *Trie
	t = NewTrie()
	t.Insert("and")
	t.Insert("ant")
	t.Insert("ace")
	//t.Show()
	//found := t.Search("age")
	//fmt.Println(found)
}
