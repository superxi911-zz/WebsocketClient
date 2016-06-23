// WebsocketClient project WebsocketClient.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

var ws *websocket.Conn

//create_version__aliceUID__4857ca9c_ca0b_4c33_8a6c_010e50cf0491__2d4e5c1a_c996_4292_8f73_23c3cefb8528
//create_version__aliceUID__4857ca9c_ca0b_4c33_8a6c_010e50cf0491__2d4e5c1a_c996_4292_8f73_23c3cefb8528
//reate_version__c818874c_3ade_4525_a0ba_87b553fb7abd__57cb9182_4de5_4496_8485_dc3bd608c81f__789c0e5a_358b_4b81_af84_e5565ddb6add
const (
	UserID    = "superxi"
	ServiceID = "88ad3371-4a7b-47da-83a9-1b32fd800eaf"
	VersionID = "86300d4f-6589-4288-81bc-5233315de812"
)

func main() {
	fmt.Println("websocket client")

	if err := Dial(); nil != err {
		log.Fatal(err)
		return
	}
	defer ws.Close()

	go Communication()
	go Read()

	var cmd string
	for {
		fmt.Scanf("%s", cmd)
		//go Read()
		StartWatchLog()
		fmt.Scanf("%s", cmd)
		//go Read()
		StopWatchLog()
	}
}

func Dial() error {
	//origin := "http://118.193.143.243/"
	//url := "ws://118.193.143.243:8001/ws"
	//origin := "http://fornax-canary.caicloud.io/"
	//url := "ws://fornax-canary.caicloud.io:8001/ws"
	origin := "http://118.193.185.173/"
	url := "ws://118.193.185.173:8000/ws"
	var err error
	ws, err = websocket.Dial(url, "", origin)
	return err
}

func Read() {
	var msg = make([]byte, 1024)
	var n int
	var err error

	for {
		if n, err = ws.Read(msg); err != nil {
			fmt.Printf("read error %s\n", err.Error())
			if strings.Contains(err.Error(), "timeout") {
				fmt.Printf("time out !!!\n")
			}
			break
		}
		fmt.Printf("Received: %s\n", msg[:n])
		//printHexLog(msg[:n])
	}
	fmt.Println("end read !!!")
}

func printHexLog(msg []byte) {
	var mapData map[string]interface{}
	if err := json.Unmarshal(msg, &mapData); err != nil {
		panic(err)
	}

	if nil == mapData["log"] {
		return
	}

	slog := mapData["log"].(string)
	log := []byte(slog)
	for i := 0; i < len(log); i++ {
		fmt.Printf("%02x ", log[i])
	}
	fmt.Printf("\n")
}

func Write(msg []byte) {
	if _, err := ws.Write(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", msg[:])
}

func StartWatchLog() {
	msg := PacketWatchLog("create_version", UserID, ServiceID,
		VersionID, "start", uuid.NewV4().String())
	Write(msg)
}

func StopWatchLog() {
	msg := PacketWatchLog("create-version", UserID, ServiceID,
		VersionID, "stop", uuid.NewV4().String())
	Write(msg)
}

func HeartBeat() {
	msg := PacketHeartBeat(uuid.NewV4().String())
	Write(msg)
}

func Communication() {
	for {
		HeartBeat()
		time.Sleep(time.Second * 30)
	}
}
