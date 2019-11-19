package climbingLeaderBoard

import (
	"testing"
)

func TestClimbingLeaderBoard(t *testing.T) {

	shifts := insertionSort(slice)
	if shifts != 0 {
		t.Errorf("got %d shifts instead of 0", shifts)
	}

}
