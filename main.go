package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Header struct {
	CLIENT_LIST   []map[string]string `json:"clientList"`
	ROUTING_TABLE []map[string]string `json:"routingTable"`
	GLOBAL_STATS  string              `json:"globalStats"`
}

type Result struct {
	TITLE  string   `json:"title"`
	TIME   []string `json:"time"`
	HEADER Header   `json:"header"`
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Не указан путь к файлу лога")
		return
	}

	fileLog, err := os.Open(os.Args[1])

	if !check(err) {
		return
	}
	defer func() {
		if err = fileLog.Close(); err != nil {
			panic(err)
		}
	}()

	result := Result{
		TITLE: "",
		TIME:  []string{},
		HEADER: Header{
			CLIENT_LIST:   []map[string]string{},
			ROUTING_TABLE: []map[string]string{},
			GLOBAL_STATS:  "",
		},
	}

	var clientListHeader []string

	scanner := bufio.NewScanner(fileLog)
	for scanner.Scan() {
		frames := strings.Split(scanner.Text(), ",")
		if "TITLE" == frames[0] {
			result.TITLE = frames[1]
		}
		if "TIME" == frames[0] {
			result.TIME = frames[1:]
		}
		if "HEADER" == frames[0] && "CLIENT_LIST" == frames[1] {
			clientListHeader = frames[2:]
		}
		if "CLIENT_LIST" == frames[0] {
			client := map[string]string{}
			for index, value := range clientListHeader {
				client[value] = frames[index+1]
			}
			result.HEADER.CLIENT_LIST = append(result.HEADER.CLIENT_LIST, client)
		}
	}

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

func check(e error) bool {
	if e != nil {
		fmt.Println(e)
		return false
	}
	return true
}
