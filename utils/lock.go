package utils

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

func EncryptTEAFromBytes(plainText, key []byte) []byte {
	const (
		TeaRounds int    = 8
		TeaDelta  uint32 = 0x9E3779B9
	)
	le := binary.LittleEndian
	// key小端分组 4*4 = 16
	k1 := le.Uint32(key[0:4])
	k2 := le.Uint32(key[4:8])
	k3 := le.Uint32(key[8:12])
	k4 := le.Uint32(key[12:16])
	// padding
	text := plainText[:]
	if pad := len(text) % 8; pad != 0 {
		text = append(text, make([]byte, 8-pad)...)
	}
	fmt.Println("Padding:" + hex.EncodeToString(text))
	// 分组加密
	var sum, v1, v2 uint32
	res := make([]byte, 0)
	for i := 0; i*4 < len(text); i += 2 {
		sum = 0
		v1 = le.Uint32(text[i*4 : (i+1)*4])
		v2 = le.Uint32(text[(i+1)*4 : (i+2)*4])
		for j := 0; j < TeaRounds; j++ {
			sum += TeaDelta
			v1 += ((v2 << 4) + k1) ^ (v2 + sum) ^ ((v2 >> 5) + k2)
			v2 += ((v1 << 4) + k3) ^ (v1 + sum) ^ ((v1 >> 5) + k4)
		}

		res = le.AppendUint32(res, v1)
		res = le.AppendUint32(res, v2)
	}
	return res
}

func getTimestamp() []byte {
	now := time.Now()
	year := byte(now.Year() % 100)
	month := byte(now.Month())
	day := byte(now.Day())
	hour := byte(now.Hour())
	minute := byte(now.Minute())
	second := byte(now.Second())
	return []byte{year, month, day, hour, minute, second}
}

func GenerateCommand(cmd, roll byte, mac, key []byte) []byte {
	var b []byte
	switch cmd {
	case 0xE0:
		fallthrough
	case 0x01:
		b = make([]byte, 17)
	default:
		return nil
	}
	// 基本信息, 不加密
	b[0] = cmd
	b[1] = roll
	b[2] = 0x00
	b[3] = byte(len(b) - 4)
	copy(b[4:10], mac)
	// 指令参数
	copy(b[10:16], getTimestamp())
	sum := byte(0x00)
	for _, a := range b[:16] {
		sum += a
	}
	b[16] = sum
	if b[0] != 0x01 {
		// 加密 b[4:]
		b = append(b[:4], EncryptTEAFromBytes(b[4:], key)...)
	}
	return b
}
