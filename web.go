package main

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"log"
	"syscall/js"
)

var key = []byte("abe023f_933&#@fl")
var origData = []byte("Hello World") // 待加密的数据

 
 
// 加密
// -export aesen
func aesen(this js.Value, args []js.Value) interface{} {
	str := args[0].String()
	log.Println("参数", str)

	origData := []byte(str) // 待加密的数据
	encrypted := AesEncryptECB(origData, key)

	// // log.Println("------------------ ECB模式 --------------------")

	log.Println("密文(base64)：", base64.StdEncoding.EncodeToString(encrypted))

	// s2, _ := hex.DecodeString(s1)

	// decrypted := AesDecryptECB(s2, key)
	// log.Println("解密结果：", string(decrypted))

	return hex.EncodeToString(encrypted)
}

// 解密
// -export aesde
func aesde(this js.Value, args []js.Value) interface{} {

	str := args[0].String()

	s, _ := base64.StdEncoding.DecodeString(str)

	log.Println("(base64)：", s)

	// encrypted, err := hex.DecodeString(string(s))

	decrypted := AesDecryptECB(s, key)

	return string(decrypted)

}

func main() {

	p0 := make(chan struct{}, 0)
	js.Global().Get("util").Set("en", js.FuncOf(aesen))
	js.Global().Get("util").Set("de", js.FuncOf(aesde))
	<-p0

}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
