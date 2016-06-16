// Package set provides ...
package set

import (
	"fmt"
	"testing"
)

func TestHashSet(t *testing.T) {
	set := NewHashSet()
	var other *HashSet

	set.Add(3)
	set.Add(4)

	// other.Add(3)
	is := set.Same(other)
	fmt.Println(is)

}
