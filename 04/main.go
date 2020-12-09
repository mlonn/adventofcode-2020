package main

import (
	"advent-of-code-2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)




// Part1 Part 1 of puzzle
func Part1(input string) string {
	passports := strings.Split(input, "\n\n")
	valid:= 0
	for _, passport := range passports {
		clean := strings.ReplaceAll(passport,"\n", " ")
		passportFields := strings.Split(clean, " ")
		fieldMap := make(map[string]string)
		for _, field := range passportFields {
			split := strings.Split(field,":")
			if split[0] != "cid" {
				fieldMap[split[0]] = split[1]
			}
		}
		if len(fieldMap) != 7 {
			continue
		}
		valid++
	}
	return "Answer " + strconv.Itoa(valid)
}

// Part2 Part2 of puzzle
func Part2(input string) string {
	passports := strings.Split(input, "\n\n")
	valid:= 0
	for _, passport := range passports {
		clean := strings.ReplaceAll(passport,"\n", " ")
		passportFields := strings.Split(clean, " ")
		fieldMap := make(map[string]string)
		for _, field := range passportFields {
			split := strings.Split(field,":")
			if split[0] != "cid" {
				fieldMap[split[0]] = split[1]
			}
		}
		if len(fieldMap) != 7 {
			continue
		}
		byr, _ := strconv.Atoi(fieldMap["byr"])
		if byr < 1920 || byr > 2002 {
			continue
		}
		iyr, _ := strconv.Atoi(fieldMap["iyr"])
		if iyr < 2010 || iyr > 2020 {
			continue
		}
		eyr, _ := strconv.Atoi(fieldMap["eyr"])
		if eyr < 2020 || eyr > 2030 {
			continue
		}
		hgt := fieldMap["hgt"]
		in := strings.HasSuffix(hgt,"in")
		cm := strings.HasSuffix(hgt,"cm")
		if !cm && !in{
			continue
		}
		if cm {
			cmI, _ := strconv.Atoi(strings.ReplaceAll(hgt,"cm",""))
			if cmI < 150 || cmI > 193 {
				continue
			}
		}
		if in {
			inI, _ := strconv.Atoi(strings.ReplaceAll(hgt,"in",""))
			if inI < 59 || inI > 76 {
				continue
			}
		}
		hcl := fieldMap["hcl"]
		match, _ := regexp.MatchString("^#[a-f0-9]{6}$", hcl)
		if !match{
			continue
		}
		ecl := fieldMap["ecl"]
		match, _ = regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth$", ecl)
		if !match{
			continue
		}
		pid := fieldMap["pid"]
		match, _ = regexp.MatchString("^[0-9]{9}$", pid)
		if !match {
			continue
		}
		valid++
	}
	return "Answer " + strconv.Itoa(valid)
}

func main() {
	file := utils.Input(2020,4)
	start := time.Now()
	fmt.Println("Part 1: " + Part1(file), "Time", time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: " + Part2(file),"Time", time.Since(start))
}
