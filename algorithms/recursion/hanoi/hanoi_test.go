package hanoi

import (
	"testing"
)

func TestMoveHanoi(t *testing.T) {

	t.Run("4", func(t *testing.T) {
		MoveHanoi(4, "A", "B", "C", true)
	})

}
