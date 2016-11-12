package uuid

import (
	"fmt"
	"testing"
)

func TestUUIDv4(t *testing.T) {
	rand := NewV4().String()
	fmt.Println(rand)
}
