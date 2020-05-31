package main

import (
	"strings"

	"github.com/ivan-bogach/utils"
)

// allowed - ...
type allowed struct {
	members []*Member
}

// GetMembers ...
func GetMembers() map[string]string {
	s, _ := utils.ReadFileLines("collegues.csv")
	m := utils.StringToMap(strings.Join(s, "\n"), "\n", ",")
	return m
}

// Member - ...
type Member struct {
	name string
	id   string
}

// str := "GAGEN,12345657899\nVIALO,12345469877"
// utils.WriteFile("collegues.csv", str)

// fmt.Println(GetMembers())
