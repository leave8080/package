package MD5

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratePassword(pwd string, salt string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	p := hex.EncodeToString(h.Sum(nil))
	h.Reset()
	h.Write([]byte(p))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}
