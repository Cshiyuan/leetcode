package awesome

import (
	"testing"
)

func TestTrie_Constructor(t *testing.T) {

	 got := Trie_Constructor();
	 got.Insert("app")
	 got.Search("apple")
	 got.Search("app")
}
