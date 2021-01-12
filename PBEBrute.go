package main

import (
	"bufio"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func getDerivedKey(password string, salt []byte, count int) ([]byte, []byte) {
	key := md5.Sum([]byte(password + string(salt)))
	for i := 0; i < count-1; i++ {
		key = md5.Sum(key[:])
	}
	return key[:8], key[8:]
}

func Encrypt(password string, obtenationIterations int, plainText string, salt []byte) (string, error) {
	padNum := byte(8 - len(plainText)%8)
	for i := byte(0); i < padNum; i++ {
		plainText += string(padNum)
	}

	dk, iv := getDerivedKey(password, salt, obtenationIterations)

	block, err := des.NewCipher(dk)

	if err != nil {
		return "", err
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(plainText))
	encrypter.CryptBlocks(encrypted, []byte(plainText))

	//return base64.StdEncoding.EncodeToString(encrypted), nil
	hexStr := fmt.Sprintf("%x", encrypted)
	return hexStr, nil

}

func Decrypt(password string, obtenationIterations int, cipherText string, salt []byte) (string, error) {
	msgBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	dk, iv := getDerivedKey(password, salt, obtenationIterations)
	block, err := des.NewCipher(dk)

	if err != nil {
		return "", err
	}

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(msgBytes))
	decrypter.CryptBlocks(decrypted, msgBytes)

	decryptedString := strings.TrimRight(string(decrypted), "\x01\x02\x03\x04\x05\x06\x07\x08")

	return decryptedString, nil
}

type Glimit struct {
	n int
	c chan struct{}
}

// initialization Glimit struct
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

func judgeCipher(password string, obtenationIterations int, plainText string, salt []byte, cipherText string) {
	res, _ := Encrypt(password, obtenationIterations, plainText, salt)
	if cipherText == res {
		fmt.Println("======Congratulations======")
		fmt.Println("user: ", plainText)
		fmt.Println("pass: ", password)
		Flag = 1
	}
}

var Flag = 0

func main() {
	banner := `

 ___  _         ___             _       
| . \| |_  ___ | . > _ _  _ _ _| |_ ___ 
|  _/| . \/ ._>| . \| '_>| | | | | / ._>
|_|  |___/\___.|___/|_|  |___| |_| \___.

	`
	fmt.Println(banner)
	var iterations int
	var userDict string
	var passDict string
	var limitGrouting int

	flag.IntVar(&iterations, "i", 1000, "PBE密钥生成循环次数, 默认1000")
	flag.StringVar(&userDict, "u", "user.txt", "含有: username,salt,password 的待爆破本 格式如：admin,cb362cfeefbf3d8d,RCGTeGiH 按行分割")
	flag.StringVar(&passDict, "f", "dict.txt", "密码文件，按行分割")
	flag.IntVar(&limitGrouting, "t", 10000, "爆破协程数，默认10000")
	flag.Parse()

	var wg = sync.WaitGroup{}

	//userFile, _ := os.Open("/Users/anyone/Documents/PenTest/hxb/admin.txt")
	userFile, err := os.Open(userDict)
	if err != nil {
		fmt.Println("[Error] Open ", userDict, " Error!")
		flag.Usage()
		return
	}
	defer userFile.Close()

	g := New(limitGrouting)

	userScanner := bufio.NewScanner(userFile)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " [INFO] Start Brute")

	for userScanner.Scan() {
		Flag = 0
		//fmt.Println(time.Now())
		lineText := userScanner.Text()
		//println(lineText)
		_tmp := strings.Split(lineText, ",")
		originalText := _tmp[0]
		salt := []byte(_tmp[1])
		cipherText := _tmp[2]
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " [INFO] Brute: ", originalText)
		passwordFile, err := os.Open(passDict)
		if err != nil {
			fmt.Println("[Error] Open ", passDict, " Error!")
			flag.Usage()
			return
		}
		defer passwordFile.Close()
		pwdScanner := bufio.NewScanner(passwordFile)

		for pwdScanner.Scan() {
			var password string
			password = pwdScanner.Text()
			wg.Add(1)
			goFunc := func() {
				// 做一些业务逻辑处理
				judgeCipher(password, iterations, originalText, salt, cipherText)
				time.Sleep(time.Second)
				wg.Done()
			}
			g.Run(goFunc)
			if Flag == 1 {
				break
			}
		}
		wg.Wait()
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " [INFO] END Brute")
}
