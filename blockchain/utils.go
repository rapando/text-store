package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

// ToHex : a utility that converts an integer to a hex byte
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panicf("unable to convert %d to hex because %v",num, err)
	}
	return buff.Bytes()
}
