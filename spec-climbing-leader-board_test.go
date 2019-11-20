package main

import (
	"fmt"
	"testing"
)

func TestClimbingLeaderBoardSample1(t *testing.T) {
	scores := []int32{100, 100, 50, 40, 40, 20, 10}
	alice := []int32{5, 25, 50, 120}

	rank := climbingLeaderboard(scores, alice)
	fmt.Println(rank)
	if rank[0] != 6 ||
		rank[1] != 4 ||
		rank[2] != 2 ||
		rank[3] != 1 {
		t.Errorf("got %+v instead of [6, 4, 2, 1]", rank)
	}
}

func TestClimbingLeaderBoardSample2(t *testing.T) {
	scores := []int32{100, 90, 90, 80, 75, 60}
	alice := []int32{50, 65, 77, 90, 102}

	rank := climbingLeaderboard(scores, alice)
	fmt.Println(rank)
	if rank[0] != 6 ||
		rank[1] != 5 ||
		rank[2] != 4 ||
		rank[3] != 2 ||
		rank[4] != 1 {
		t.Errorf("got %+v instead of [6, 5, 4, 2, 1]", rank)
	}
}

func TestClimbingLeaderBoardSample3(t *testing.T) {
	scores := []int32{100, 0}
	alice := []int32{50, 65, 77, 90, 102}

	rank := climbingLeaderboard(scores, alice)
	fmt.Println(rank)
	if rank[0] != 2 ||
		rank[1] != 2 ||
		rank[2] != 2 ||
		rank[3] != 2 ||
		rank[4] != 1 {
		t.Errorf("got %+v instead of [2, 2, 2, 2, 1]", rank)
	}
}
