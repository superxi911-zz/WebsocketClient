package main

import (
	"encoding/json"
)

type WatchLogPacket struct {
	Action    string `json:"action"`
	Api       string `json:"api"`
	UserId    string `json:"user_id"`
	ServiceId string `json:"service_id"`
	VersionId string `json:"version_id"`
	Operation string `json:"operation"`
	Id        string `json:"id"`
}

type PushLogPacket struct {
	Action    string `json:"action"`
	Api       string `json:"api"`
	UserId    string `json:"user_id"`
	ServiceId string `json:"service_id"`
	VersionId string `json:"version_id"`
	Log       string `json:"log"`
	Id        string `json:"id"`
}

type HeartBeatPacket struct {
	Action string `json:"action"`
	Id     string `json:"id"`
}

type ResponsePacket struct {
	Action    string `json:"action"`
	Response  string `json:"response"`
	IdAck     string `json:"id_ack"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

const (
	Error_Code_Successful = 0
	Error_Code_Failure    = 4001
)

var ErrorMsgMap = map[int]string{
	0:    "successful",
	4001: "failure",
}

func PacketWatchLog(sApi string, sUserId string, sServiceId string,
	sVersion string, sOperation string, sId string) []byte {
	structData := &WatchLogPacket{
		Action:    "watch_log",
		Api:       sApi,
		UserId:    sUserId,
		ServiceId: sServiceId,
		VersionId: sVersion,
		Operation: sOperation,
		Id:        sId}
	jsonData, _ := json.Marshal(structData)
	return jsonData
}

func PacketPushLog(sApi string, sUserId string, sServiceId string,
	sVersion string, sLog string, sId string) []byte {
	structData := &PushLogPacket{
		Action:    "push_log",
		Api:       sApi,
		UserId:    sUserId,
		ServiceId: sServiceId,
		VersionId: sVersion,
		Log:       sLog,
		Id:        sId}
	jsonData, _ := json.Marshal(structData)
	return jsonData
}

func PacketHeartBeat(sId string) []byte {
	structData := &HeartBeatPacket{
		Action: "heart_beat",
		Id:     sId}
	jsonData, _ := json.Marshal(structData)
	return jsonData
}

func PacketResponse(sResponse string, sIdAck string, nErrorCode int) []byte {
	structData := &ResponsePacket{
		Action:    "response",
		Response:  sResponse,
		IdAck:     sIdAck,
		ErrorCode: nErrorCode,
		ErrorMsg:  ErrorMsgMap[nErrorCode]}
	jsonData, _ := json.Marshal(structData)
	return jsonData
}
