package utils

import "encoding/hex"

type Locker struct{}

func NewLocker() *Locker {
	return &Locker{}
}

func (l *Locker) EncoderHex(b []byte) string {
	return hex.EncodeToString(b)
}
func (l *Locker) DecoderHex(str string) []byte {
	b, _ := hex.DecodeString(str)
	return b
}

func (l *Locker) Unlock() []byte {
	return nil
}
func (l *Locker) Lock() {

}
func (l *Locker) UpdateStatus(key string) {
	return
}
