package utils

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/protocol"
)

type keyValue struct {
	Key         string
	Value       interface{}
	TotalLength int
	ExpiryTime  int64
}

// Function to read the RDB file and parse the key-value pairs
func ParseFile(path string) []keyValue {
	fileData, _ := os.ReadFile(path)
	keysLen := getKeysLen(fileData)

	var keyValues []keyValue
	currIdx := getPairsStartIdx(fileData)
	for i := 0; i < keysLen; i++ {
		kv := parseKeyValue(fileData, currIdx)
		keyValues = append(keyValues, kv)
		currIdx += kv.TotalLength
	}

	return keyValues
}

func parseKeyValue(fileData []byte, currIdx int) keyValue {
	firstByte := fileData[currIdx]
	var currIdxOffset int
	var expiryTime int64

	if firstByte == protocol.OpCodeExpireTime {
		expirationTime := binary.LittleEndian.Uint32(fileData[currIdx+1 : currIdx+5])
		currIdxOffset = 5
		expiryTime = int64(expirationTime * 1000) // convert to milliseconds
	} else if firstByte == protocol.OpCodeExpireTimeMs {
		expirationTime := binary.LittleEndian.Uint64(fileData[currIdx+1 : currIdx+9])
		currIdxOffset = 9
		expiryTime = int64(expirationTime)
	}

	kv := parseKeyValueInner(fileData, currIdx+currIdxOffset)
	kv.TotalLength += currIdxOffset
	if expiryTime > 0 {
		kv.ExpiryTime = expiryTime
	} else {
		kv.ExpiryTime = -1
	}

	return kv
}

func parseKeyValueInner(fileData []byte, currIdx int) keyValue {
	keyLen := int(fileData[currIdx+1])
	keyEndIdx := currIdx + 1 + keyLen + 1
	key := string(fileData[currIdx+2 : keyEndIdx])

	valueLen := int(fileData[keyEndIdx])
	value := string(fileData[keyEndIdx+1 : keyEndIdx+1+valueLen])

	return keyValue{key, value, len(key) + len(value) + 3, 0}
}

func getKeysLen(fileData []byte) int {
	idx := getOpCodeIdx(fileData, protocol.OpCodeResizeDB)
	if idx == -1 || idx+1 >= len(fileData) {
		return 0
	}
	return int(fileData[idx+1])
}

func getPairsStartIdx(fileData []byte) int {
	idx := getOpCodeIdx(fileData, protocol.OpCodeResizeDB)
	if idx == -1 || idx+3 >= len(fileData) {
		return len(fileData) // Return end of data if not enough length
	}
	return idx + 3
}

func getOpCodeIdx(fileData []byte, opCode byte) int {
	return bytes.IndexByte(fileData, opCode)
}
