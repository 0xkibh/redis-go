package main

import (
	"fmt"
	"strings"
)

func encodeSimpleString(str string) string{
	return fmt.Sprintf("+%s\r\n", str)
}

func encodeBulkString(strs []string) string{
	bulkString := strings.Join(strs, " ")
	return fmt.Sprintf("$%d\r\n%s\r\n", len(bulkString), bulkString)
}

func parseCommand(str string) []string{
	commandArray := strings.Split(str, CRLFTerminator)
	commands := []string{}
	for i:=2; i<len(commandArray); i+=2 {
		commands = append(commands, commandArray[i])
	}
	return commands
}