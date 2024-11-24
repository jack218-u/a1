package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh" // 添加这一行
)

func main() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("无法生成RSA密钥对:", err)
		os.Exit(1)
	}

	// 将私钥编码为PEM格式
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// 将公钥编码为OpenSSH格式
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("无法生成OpenSSH公钥:", err)
		os.Exit(1)
	}

	// 将OpenSSH公钥转换为字符串
	publicKeyStr := string(ssh.MarshalAuthorizedKey(publicKey))

	// 输出OpenSSH公钥
	fmt.Println("OpenSSH公钥:")
	fmt.Println(publicKeyStr)

	// 将私钥PEM写入文件
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println("无法创建私钥文件:", err)
		os.Exit(1)
	}
	defer privateKeyFile.Close()

	pem.Encode(privateKeyFile, privateKeyPEM)
}
