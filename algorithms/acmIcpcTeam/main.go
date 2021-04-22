package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// utilities

func readLine(reader *bufio.Reader) string {
	text, _, _ := reader.ReadLine()
	return string(text)
}

func splitInput(input string) []int {
	var data []int

	for _, val := range strings.Split(input, " ") {
		intVal, _ := strconv.ParseInt(val, 10, 64)
		data = append(data, int(intVal))
	}

	return data
}

func calculateKnownTopicForTeam(a, b string) int {
	known := 0
	for i := 0; i < len(a); i++ {
		if string(a[i]) == "1" || string(b[i]) == "1" {
			known++
		}
	}
	return known
}

func getMax(arr []int) int {
	highest := 0

	for _, val := range arr {
		if val > highest {
			highest = val
		}
	}

	return highest
}

func getTeamCountWithScore(teams []int, score int) int {
	amountOfTeams := 0

	for _, team := range teams {
		if team == score {
			amountOfTeams++
		}
	}

	return amountOfTeams
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	data := splitInput(readLine(reader))
	attendeesCount := data[0]
	attendees := []string{}
	possibleTeams := []int{}
	handledCombos := make(map[string]bool)

	for i := 0; i < attendeesCount; i++ {
		attendees = append(attendees, readLine(reader))
	}

	for i := 0; i < attendeesCount; i++ {
		for j := 0; j < attendeesCount; j++ {
			slug := string(i) + "-" + string(j)
			if handledCombos[slug] || handledCombos[reverse(slug)] {
				continue
			}
			possibleTeams = append(possibleTeams, calculateKnownTopicForTeam(attendees[i], attendees[j]))
			handledCombos[slug] = true
		}
	}

	highestTopics := getMax(possibleTeams)
	fmt.Println(highestTopics)
	fmt.Println(getTeamCountWithScore(possibleTeams, highestTopics))
}
