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
	for _, value := range alice {
		rank = calculateRank(value, leaderboard)
		aliceRank = append(aliceRank, rank)
	}
	return aliceRank
}

func stripDulicates(scores []int32) []int32 {
	uniqueMap := map[int32]bool{}
	unique := []int32{}

	for index := range scores {
		if uniqueMap[scores[index]] != true {
			uniqueMap[scores[index]] = true
			unique = append(unique, scores[index])
		}
	}

	return unique
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
		if leaderboard[median] == score {
			return median + 1
		} else if leaderboard[median] < score {
			high = median - 1
		} else {
			low = median + 1
		}
	}

	if leaderboard[median] > score && leaderboard[median+1] < score {
		return median + 2
	} else if leaderboard[median-1] > score && leaderboard[median] < score {
		return median + 1
	} else {
		panic("problem in calculate rank")
	}
}

// ********************************************************************************
func main() {
	file, err := os.Open("test-case-1")
	checkError(err)

	reader := bufio.NewReaderSize(file, 1600*1600)

	stdout, err := os.Create("OUTPUT-test-case-1")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1600*1600)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

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
