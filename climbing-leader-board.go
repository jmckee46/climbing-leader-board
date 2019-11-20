package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	leaderboard := stripDulicates(scores)
	aliceRank := make([]int32, 0)
	var rank int32
	for index, value := range alice {
		fmt.Println("alice index:", index)
		fmt.Println("alice value:", value)
		rank = calculateRank(value, leaderboard)
		fmt.Println("rank:", rank)
		aliceRank = append(aliceRank, rank)
	}
	return aliceRank
}

func stripDulicates(scores []int32) []int32 {
	for i := 0; i < len(scores)-2; i++ {
		if scores[i] == scores[i+1] {
			copy(scores[i:], scores[i+1:])
			scores[len(scores)-1] = 0
			scores = scores[:len(scores)-1]
		}
	}

	return scores
}

func calculateRank(score int32, leaderboard []int32) int32 {
	low := int32(0)
	high := int32(len(leaderboard) - 1)
	var median int32
	if score > leaderboard[low] {
		return int32(1)
	}
	if score < leaderboard[high] {
		return high + 2
	}

	for low <= high {
		median = (low + high) / 2
		fmt.Println("low:", low)
		fmt.Println("high:", high)
		fmt.Println("median:", median)
		fmt.Println("leaderboard[median]:", leaderboard[median])
		if leaderboard[median] == score {
			return median + 1
			// } else if scoreInbetweenRight(score, leaderboard, median) {
			// 	return median + 2
			// } else if scoreInbetweenLeft(score, leaderboard, median) {
			// 	return median
		} else if leaderboard[median] < score {
			high = median - 1
		} else {
			low = median + 1
		}
	}
	fmt.Println("ending low:", low)
	fmt.Println("ending high:", high)
	fmt.Println("ending median:", median)

	if leaderboard[median] > score && leaderboard[median+1] < score {
		fmt.Println("first part of if")
		return median + 2
	} else if leaderboard[median-1] > score && leaderboard[median] < score {
		fmt.Println("second part of if")
		return median + 1
	} else {
		panic("problem in calculate rank")
	}
}

func scoreInbetweenRight(score int32, leaderboard []int32, median int32) bool {
	if leaderboard[median] > score && leaderboard[median+1] < score {
		return true
	}

	return false
}

func scoreInbetweenLeft(score int32, leaderboard []int32, median int32) bool {
	if leaderboard[median-1] > score && leaderboard[median] < score {
		return true
	}

	return false
}

// ********************************************************************************
func main() {
	file, err := os.Open("test-case-6")
	checkError(err)

	reader := bufio.NewReaderSize(file, 1600*1600)

	stdout, err := os.Create("OUTPUT-test-case-6")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1600*1600)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	fmt.Println("scoresCount:", scoresCount)
	scoresTemp := strings.Split(readLine(reader), " ")
	fmt.Println("scoresTemp length:", len(scoresTemp))

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
