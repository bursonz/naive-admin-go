package utils

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"naive-admin-go/model"
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
	log.Println("指令补零:" + hex.EncodeToString(text))
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

func GenerateCommand(cmd, roll byte, mac, key, newKey []byte) []byte {
	var b []byte
	// 生成命令
	switch cmd {
	case 0x1F, 0xE0, 0x01:
		b = make([]byte, 17)
		// 基本信息, 不加密
		b[0] = cmd
		b[1] = roll
		b[2] = 0x00
		b[3] = 0x0d
		// 参数
		copy(b[4:10], mac)
		copy(b[10:16], getTimestamp()) // 时间戳
	case 0x10:
		b = make([]byte, 33)
		// 基本信息, 不加密
		b[0] = cmd
		b[1] = roll
		b[2] = 0x00
		b[3] = 0x1d
		// 参数
		copy(b[4:10], mac)
		copy(b[10:26], newKey) // 新密钥
		log.Println("新密钥：" + hex.EncodeToString(newKey))
		copy(b[26:32], getTimestamp()) // 时间戳
	default:
		return nil
	}
	// 校验和
	sum := byte(0x00)
	for _, a := range b[:len(b)-1] {
		sum += a
	}
	b[len(b)-1] = sum
	log.Println("明文命令：" + hex.EncodeToString(b) + " 长度：" + fmt.Sprintf("%d", len(b)) + " 字节")
	// 加密
	if b[0] != 0x01 {
		// 加密 b[4:]
		b = append(b[:4], EncryptTEAFromBytes(b[4:], key)...) // 不够8个字节的，会自动填充0
	}
	log.Println("加密命令：" + hex.EncodeToString(b) + " 长度：" + fmt.Sprintf("%d", len(b)) + " 字节")
	return b
}

func ParseCommand(cmd []byte, lock *model.Lock) error {
	if lock == nil {
		return errors.New("无法找到锁，请核对Id")
	}
	switch cmd[0] {
	case 0x01:
		lock.HardwareVersion = hex.EncodeToString(cmd[10:11])
		lock.SoftwareVersion = hex.EncodeToString(cmd[11:13])
		lock.FactoryId = hex.EncodeToString(cmd[13:17])
		lock.AlarmMode = hex.EncodeToString(cmd[17:18])
		lock.LockStatus = hex.EncodeToString(cmd[18:19])
		lock.BackupData = hex.EncodeToString(cmd[19:23])
		lock.NewLock = hex.EncodeToString(cmd[23:24]) //TODO 旧锁二次添加，目前不需要
		lock.UnlockRecord = hex.EncodeToString(cmd[24:26])
		lock.Power = hex.EncodeToString(cmd[26:27])
		lock.Muted = hex.EncodeToString(cmd[27:28])
		lock.Hibernate = hex.EncodeToString(cmd[28:29])
	case 0xE0:
		switch cmd[11] {
		case 0x01:
			// 开锁成功
			lock.LockStatus = hex.EncodeToString([]byte{0x01})
		case 0x05:
			return errors.New("开锁失败,MAC地址错误")
		}
		break
	default:
		return errors.New("未知指令")
	}
	return nil
}
