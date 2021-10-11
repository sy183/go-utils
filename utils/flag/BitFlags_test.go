package flag

import (
	"fmt"
	"testing"
)

func TestNewBitFlag(t *testing.T) {
	bf := NewBitFlag(128)
	fmt.Println(bf)
}
