package main

// Trie is used as a dictionary of without space and lower case words
import "fmt"

const alphabet = 26

type Node struct {
	children [alphabet]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	trie := &Trie{
		root: &Node{},
	}
	return trie
}

func (t *Trie) InsertWord(word string) {
	currNode := t.root
	for i := 0; i < len(word); i++ {
		letterIndex := word[i] - 'a'
		if currNode.children[letterIndex] == nil {
			currNode.children[letterIndex] = &Node{}
		}
		currNode = currNode.children[letterIndex]
	}
	currNode.isEnd = true
}

func (t *Trie) RemoveWord(word string) {
	currNode := t.root
	for i := 0; i < len(word); i++ {
		letterIndex := word[i] - 'a'
		if currNode.children[letterIndex] == nil {
			fmt.Println(word + " is not in the Trie")
			return
		}
		currNode = currNode.children[letterIndex]
	}
	if currNode.isEnd {
		currNode.isEnd = false
	}
}

func (t *Trie) SearchWord(word string) bool {
	currNode := t.root
	for i := 0; i < len(word); i++ {
		letterIndex := word[i] - 'a'
		if currNode.children[letterIndex] == nil {
			return false
		}
		currNode = currNode.children[letterIndex]
	}
	return currNode.isEnd
}

func main() {
	dict := InitTrie()
	dict.InsertWord("hello")
	dict.InsertWord("gopher")
	fmt.Println("Search for gopher returned", dict.SearchWord("gopher")) //true

	dict.InsertWord("go")
	dict.RemoveWord("gopher")
	fmt.Println("Search for gopher returned", dict.SearchWord("gopher")) //false
}
