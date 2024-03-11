package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"time"
)

func GenerateTOTP(key []byte, timeStep int64) (string, error) {
	timeNow := time.Now().Unix() / timeStep
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, uint64(timeNow))

	h := hmac.New(sha1.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return "", err
	}

	hash := h.Sum(nil)
	offset := hash[len(hash)-1] & 0x0F
	truncatedHash := hash[offset : offset+4]

	value := binary.BigEndian.Uint32(truncatedHash)
	value &= 0x7FFFFFFF
	value %= 1000000

	return fmt.Sprintf("%06d", value), nil
}
