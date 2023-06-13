package formats

import (
	"encoding/json"
	"fmt"
	"os"
)

type RequestContent struct {
	User    string //`json:"user"`
	Message string `json:"msg"`
}
type Request struct {
	Request RequestContent
	Author  string `json:"user"`
}

func LoadAndParseJson() {
	jsonData, err := os.ReadFile("exampls/example.json")
	fmt.Println(err)
	var request Request
	json.Unmarshal(jsonData, &request)
	fmt.Println(request)
}

// если хрен его знает что там лежит -> interface
func LoadAndParseRawMsgToMap() {
	jsonData, _ := os.ReadFile("exampls/example.json")
	var objmap map[string]interface{}
	json.Unmarshal(jsonData, &objmap)
	fmt.Println(objmap)
}

// какая то хитрая хрень
func LoadAndParseRawMsg() {
	jsonData, _ := os.ReadFile("exampls/example.json")
	var objmap map[string]json.RawMessage
	json.Unmarshal(jsonData, &objmap)
	fmt.Println(objmap)
	var internalMap map[string]string
	json.Unmarshal(objmap["request"], &internalMap)
	fmt.Println(internalMap)
}
