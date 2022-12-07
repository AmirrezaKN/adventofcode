package day07

import (
	"log"
	"strconv"
	"strings"
)

func Question07_1(lines []string) int {
	pwd := []string{"/"}
	dirSizes := make(map[string]int)

	for i := 1; i < len(lines); i++ {
		if string(lines[i][0]) == "$" {
			line := strings.Trim(lines[i], "$ ")
			commands := strings.Split(line, " ")
			if commands[0] == "ls" {
				j := i + 1
				for ; j < len(lines) && !strings.Contains(lines[j], "$"); j++ {
					parts := strings.Split(lines[j], " ")
					if parts[0] != "dir" {
						fileSize, err := strconv.Atoi(parts[0])
						if err != nil {
							log.Fatal(err, parts)
						}

						for _, dir := range pwd {
							dirSizes[dir] += fileSize
						}
					}
				}
				i = j - 1
			} else if commands[0] == "cd" {
				if commands[1] == ".." {
					pwd = pwd[:len(pwd)-1]
				} else if commands[1] == "/" {
					pwd = pwd[:1]
				} else {
					pwd = append(pwd, commands[1])
				}
			}
		}
	}

	result := 0

	for _, size := range dirSizes {
		if size <= 100000 {
			result += size
		}
	}

	return result
}

func Question07_2(str string) int {
	return 0
}
