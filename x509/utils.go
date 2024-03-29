// Copyright 2022 s1ren@github.com/hxx258456.

package x509

/*
x509/utils.go 提供gmx509常用操作的公开函数:
ReadPrivateKeyFromPem : 将pem字节数组转为对应私钥
ReadPrivateKeyFromPemFile : 将pem文件转为对应私钥
WritePrivateKeyToPem : 将私钥转为pem字节数组
WritePrivateKeytoPemFile : 将私钥转为pem文件
ReadPublicKeyFromPem :  将pem字节数组转为对应公钥
ReadPublicKeyFromPemFile : 将pem文件转为对应公钥
WritePublicKeyToPem : 将公钥转为pem字节数组
WritePublicKeytoPemFile : 将公钥转为pem文件
ReadSm2PrivFromHex : 将hex字符串转为sm2私钥
WriteSm2PrivToHex : 将sm2私钥D转为hex字符串
ReadSm2PubFromHex : 将hex字符串转为sm2公钥
WriteSm2PubToHex : 将sm2公钥转为hex字符串
ReadCertificateRequestFromPem : 将pem字节数组转为证书申请
ReadCertificateRequestFromPemFile : 将pem文件转为证书申请
CreateCertificateRequestToPem : 创建证书申请并转为pem字节数组
CreateCertificateRequestToPemFile : 创建证书申请并转为pem文件
ReadCertificateFromPem : 将pem字节数组转为gmx509证书
ReadCertificateFromPemFile : 将pem文件转为gmx509证书
CreateCertificateToPem : 创建gmx509证书并转为pem字节数组
CreateCertificateToPemFile : 创建gmx509证书并转为pem文件
ParseGmx509DerToX509 : 将gmx509证书DER字节数组转为x509证书
CreateEllipticSKI : 根据椭圆曲线公钥参数生成其SKI值
GetRandBigInt : 随机生成序列号
*/

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"gitee.com/zhaochuninhefei/zcgolog/zclog"
	"github.com/hxx258456/ccgo/sm2"
	"github.com/hxx258456/ccgo/sm3"
)

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
// 私钥与pem相互转换

// ReadPrivateKeyFromPem 将pem字节数组转为对应私钥
//   - 私钥类型: *sm2.PrivateKey, *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey
//     @param privateKeyPem 私钥pem字节数组
//     @param pwd pem解密口令
//     @return interface{} 返回私钥
//     @return error
func ReadPrivateKeyFromPem(privateKeyPem []byte, pwd []byte) (interface{}, error) {
	var block *pem.Block
	block, _ = pem.Decode(privateKeyPem)
	if block == nil || !strings.HasSuffix(block.Type, "PRIVATE KEY") {
		return nil, errors.New("failed to decode private key")
	}
	var der []byte
	var err error
	if pwd != nil {
		der, err = DecryptPEMBlock(block, pwd)
		if err != nil {
			return nil, err
		}
	} else {
		der = block.Bytes
	}
	return ParsePKCS8PrivateKey(der)
}

// ReadPrivateKeyFromPemFile 将pem文件转为对应私钥
//   - 私钥类型: *sm2.PrivateKey, *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey
//     @param FileName pem文件路径
//     @param pwd pem解密口令
//     @return interface{} 返回私钥
//     @return error
func ReadPrivateKeyFromPemFile(FileName string, pwd []byte) (interface{}, error) {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		return nil, err
	}
	return ReadPrivateKeyFromPem(data, pwd)
}

// WritePrivateKeyToPem 将私钥转为pem字节数组
//   - 私钥类型: *sm2.PrivateKey, *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey
//     @param key 私钥
//     @param pwd pem加密口令
//     @return []byte 私钥pem字节数组
//     @return error
func WritePrivateKeyToPem(key interface{}, pwd []byte) ([]byte, error) {
	var block *pem.Block
	der, err := MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}
	var pemType string
	switch key.(type) {
	case *sm2.PrivateKey:
		pemType = "SM2 PRIVATE KEY"
	case *ecdsa.PrivateKey:
		pemType = "ECDSA PRIVATE KEY"
	case ed25519.PrivateKey:
		pemType = "ED25519 PRIVATE KEY"
	case *rsa.PrivateKey:
		pemType = "RSA PRIVATE KEY"
	default:
		return nil, fmt.Errorf("gmx509.WritePrivateKeyToPem : unsupported key: [%T]", key)
	}
	if pwd != nil {
		block, err = EncryptPEMBlock(rand.Reader, "ENCRYPTED "+pemType, der, pwd, PEMCipherSM4)
		if err != nil {
			return nil, err
		}
	} else {
		block = &pem.Block{
			Type:  pemType,
			Bytes: der,
		}
	}
	certPem := pem.EncodeToMemory(block)
	return certPem, nil
}

// WritePrivateKeytoPemFile 将私钥转为pem文件
//   - 私钥类型: *sm2.PrivateKey, *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey
//     @param FileName pem文件路径
//     @param key 私钥
//     @param pwd pem加密口令
//     @return bool 成功与否
//     @return error
func WritePrivateKeytoPemFile(FileName string, key interface{}, pwd []byte) (bool, error) {
	var block *pem.Block
	der, err := MarshalPKCS8PrivateKey(key)
	if err != nil {
		return false, err
	}
	var pemType string
	switch key.(type) {
	case *sm2.PrivateKey:
		pemType = "SM2 PRIVATE KEY"
	case *ecdsa.PrivateKey:
		pemType = "ECDSA PRIVATE KEY"
	case ed25519.PrivateKey:
		pemType = "ED25519 PRIVATE KEY"
	case *rsa.PrivateKey:
		pemType = "RSA PRIVATE KEY"
	default:
		return false, fmt.Errorf("gmx509.WritePrivateKeytoPemFile : unsupported key: [%T]", key)
	}
	if pwd != nil {
		block, err = EncryptPEMBlock(rand.Reader, "ENCRYPTED "+pemType, der, pwd, PEMCipherSM4)
		if err != nil {
			return false, err
		}
	} else {
		block = &pem.Block{
			Type:  pemType,
			Bytes: der,
		}
	}
	file, err := os.Create(FileName)
	if err != nil {
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zclog.Errorln(err)
		}
	}(file)
	err = pem.Encode(file, block)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 私钥与pem相互转换
// ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
// 公钥与pem相互转换

// ReadPublicKeyFromPem 将pem字节数组转为对应公钥
//   - 公钥类型: *sm2.PublicKey, *rsa.PublicKey, *dsa.PublicKey, *ecdsa.PublicKey, ed25519.PublicKey
//     @param publicKeyPem
//     @return interface{}
//     @return error
func ReadPublicKeyFromPem(publicKeyPem []byte) (interface{}, error) {
	block, _ := pem.Decode(publicKeyPem)
	if block == nil || !strings.HasSuffix(block.Type, "PUBLIC KEY") {
		return nil, errors.New("failed to decode public key")
	}
	return ParsePKIXPublicKey(block.Bytes)
}

// ReadPublicKeyFromPemFile 将pem文件转为对应公钥
//   - 公钥类型: *sm2.PublicKey, *rsa.PublicKey, *dsa.PublicKey, *ecdsa.PublicKey, ed25519.PublicKey
//     @param FileName
//     @return interface{}
//     @return error
func ReadPublicKeyFromPemFile(FileName string) (interface{}, error) {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		return nil, err
	}
	return ReadPublicKeyFromPem(data)
}

// WritePublicKeyToPem 将公钥转为pem字节数组
//
//   - 公钥类型: *sm2.PublicKey, *rsa.PublicKey, *ecdsa.PublicKey, ed25519.PublicKey
//
//     @param key
//     @return []byte
//     @return error
func WritePublicKeyToPem(key interface{}) ([]byte, error) {
	der, err := MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, err
	}
	var pemType string
	switch key.(type) {
	case *sm2.PublicKey:
		pemType = "SM2 PUBLIC KEY"
	case *ecdsa.PublicKey:
		pemType = "ECDSA PUBLIC KEY"
	case ed25519.PublicKey:
		pemType = "ED25519 PUBLIC KEY"
	case *rsa.PublicKey:
		pemType = "RSA PUBLIC KEY"
	default:
		return nil, fmt.Errorf("gmx509.WritePublicKeyToPem : unsupported key: [%T]", key)
	}
	block := &pem.Block{
		Type:  pemType,
		Bytes: der,
	}
	certPem := pem.EncodeToMemory(block)
	return certPem, nil
}

// WritePublicKeytoPemFile 将公钥转为pem文件
//
//   - 公钥类型: *sm2.PublicKey, *rsa.PublicKey, *ecdsa.PublicKey, ed25519.PublicKey
//
//     @param FileName
//     @param key
//     @return bool
//     @return error
func WritePublicKeytoPemFile(FileName string, key interface{}) (bool, error) {
	der, err := MarshalPKIXPublicKey(key)
	if err != nil {
		return false, err
	}
	var pemType string
	switch key.(type) {
	case *sm2.PublicKey:
		pemType = "SM2 PUBLIC KEY"
	case *ecdsa.PublicKey:
		pemType = "ECDSA PUBLIC KEY"
	case ed25519.PublicKey:
		pemType = "ED25519 PUBLIC KEY"
	case *rsa.PublicKey:
		pemType = "RSA PUBLIC KEY"
	default:
		return false, fmt.Errorf("gmx509.WritePublicKeytoPemFile : unsupported key: [%T]", key)
	}
	block := &pem.Block{
		Type:  pemType,
		Bytes: der,
	}
	file, err := os.Create(FileName)
	if err != nil {
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zclog.Errorln(err)
		}
	}(file)
	err = pem.Encode(file, block)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 公钥与pem相互转换
// ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

// // ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
// // SM2公私钥与hex相互转换

// // 将hex字符串转为sm2私钥
// // Dhex是16进制字符串，对应sm2.PrivateKey.D
// func ReadSm2PrivFromHex(Dhex string) (*sm2.PrivateKey, error) {
// 	c := sm2.P256Sm2()
// 	d, err := hex.DecodeString(Dhex)
// 	if err != nil {
// 		return nil, err
// 	}
// 	k := new(big.Int).SetBytes(d)
// 	params := c.Params()
// 	one := new(big.Int).SetInt64(1)
// 	n := new(big.Int).Sub(params.N, one)
// 	if k.Cmp(n) >= 0 {
// 		return nil, errors.New("privateKey's D is overflow")
// 	}
// 	priv := new(sm2.PrivateKey)
// 	priv.PublicKey.Curve = c
// 	priv.D = k
// 	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
// 	return priv, nil
// }

// // 将sm2私钥D转为hex字符串
// func WriteSm2PrivToHex(key *sm2.PrivateKey) string {
// 	return key.D.Text(16)
// }

// // 将hex字符串转为sm2公钥
// // Qhex是sm2公钥座标x,y的字节数组拼接后的hex转码字符串
// func ReadSm2PubFromHex(Qhex string) (*sm2.PublicKey, error) {
// 	q, err := hex.DecodeString(Qhex)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(q) == 65 && q[0] == byte(0x04) {
// 		q = q[1:]
// 	}
// 	if len(q) != 64 {
// 		return nil, errors.New("publicKey is not uncompressed")
// 	}
// 	pub := new(sm2.PublicKey)
// 	pub.Curve = sm2.P256Sm2()
// 	pub.X = new(big.Int).SetBytes(q[:32])
// 	pub.Y = new(big.Int).SetBytes(q[32:])
// 	return pub, nil
// }

// // 将sm2公钥转为hex字符串
// func WriteSm2PubToHex(key *sm2.PublicKey) string {
// 	x := key.X.Bytes()
// 	y := key.Y.Bytes()
// 	if n := len(x); n < 32 {
// 		x = append(utils.ZeroByteSlice()[:32-n], x...)
// 	}
// 	if n := len(y); n < 32 {
// 		y = append(utils.ZeroByteSlice()[:32-n], y...)
// 	}
// 	c := []byte{}
// 	c = append(c, x...)
// 	c = append(c, y...)
// 	c = append([]byte{0x04}, c...)
// 	return hex.EncodeToString(c)
// }

// // SM2公私钥与hex相互转换
// // ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
// 证书申请与pem相互转换

// ReadCertificateRequestFromPem 将pem字节数组转为证书申请
//
//	@param certPem
//	@return *CertificateRequest
//	@return error
func ReadCertificateRequestFromPem(certPem []byte) (*CertificateRequest, error) {
	block, _ := pem.Decode(certPem)
	if block == nil {
		return nil, errors.New("failed to decode certificate request")
	}
	return ParseCertificateRequest(block.Bytes)
}

// ReadCertificateRequestFromPemFile 将pem文件转为证书申请
//
//	@param FileName
//	@return *CertificateRequest
//	@return error
func ReadCertificateRequestFromPemFile(FileName string) (*CertificateRequest, error) {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		return nil, err
	}
	return ReadCertificateRequestFromPem(data)
}

// CreateCertificateRequestToPem 创建证书申请并转为pem字节数组
//
//	@param template
//	@param signer
//	@return []byte
//	@return error
func CreateCertificateRequestToPem(template *CertificateRequest, signer interface{}) ([]byte, error) {
	der, err := CreateCertificateRequest(rand.Reader, template, signer)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: der,
	}
	certPem := pem.EncodeToMemory(block)
	return certPem, nil
}

// CreateCertificateRequestToPemFile 创建证书申请并转为pem文件
//
//	@param FileName
//	@param template
//	@param signer
//	@return bool
//	@return error
func CreateCertificateRequestToPemFile(FileName string, template *CertificateRequest, signer interface{}) (bool, error) {
	der, err := CreateCertificateRequest(rand.Reader, template, signer)
	if err != nil {
		return false, err
	}
	block := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: der,
	}
	file, err := os.Create(FileName)
	if err != nil {
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zclog.Errorln(err)
		}
	}(file)
	err = pem.Encode(file, block)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 证书申请与pem相互转换
// ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
// gmx509证书与pem相互转换

// ReadCertificateFromPem 将pem字节数组转为gmx509证书
//
//	@param certPem
//	@return *Certificate
//	@return error
func ReadCertificateFromPem(certPem []byte) (*Certificate, error) {
	block, _ := pem.Decode(certPem)
	if block == nil {
		return nil, errors.New("failed to decode certificate request")
	}
	return ParseCertificate(block.Bytes)
}

// ReadCertificateFromPemFile 将pem文件转为gmx509证书
//
//	@param FileName
//	@return *Certificate
//	@return error
func ReadCertificateFromPemFile(FileName string) (*Certificate, error) {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		return nil, err
	}
	return ReadCertificateFromPem(data)
}

// CreateCertificateToPem 创建gmx509证书并转为pem字节数组
//
//	@param template
//	@param parent
//	@param pubKey
//	@param signer
//	@return []byte
//	@return error
func CreateCertificateToPem(template, parent *Certificate, pubKey, signer interface{}) ([]byte, error) {
	der, err := CreateCertificate(rand.Reader, template, parent, pubKey, signer)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: der,
	}
	certPem := pem.EncodeToMemory(block)
	return certPem, nil
}

// CreateCertificateToPemFile 创建gmx509证书并转为pem文件
//
//	@param FileName
//	@param template
//	@param parent
//	@param pubKey
//	@param privKey
//	@return bool
//	@return error
func CreateCertificateToPemFile(FileName string, template, parent *Certificate, pubKey, privKey interface{}) (bool, error) {
	der, err := CreateCertificate(rand.Reader, template, parent, pubKey, privKey)
	if err != nil {
		return false, err
	}
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: der,
	}
	file, err := os.Create(FileName)
	if err != nil {
		return false, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zclog.Errorln(err)
		}
	}(file)
	err = pem.Encode(file, block)
	if err != nil {
		return false, err
	}
	return true, nil
}

// gmx509证书与pem相互转换
// ↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑

// ParseGmx509DerToX509 将gmx509证书DER字节数组转为x509证书
//
//	@param asn1data
//	@return *x509.Certificate
//	@return error
//
//goland:noinspection GoUnusedExportedFunction
func ParseGmx509DerToX509(asn1data []byte) (*x509.Certificate, error) {
	sm2Cert, err := ParseCertificate(asn1data)
	if err != nil {
		return nil, err
	}
	return sm2Cert.ToX509Certificate(), nil
}

// CreateEllipticSKI 根据椭圆曲线公钥参数生成其SKI值
//
//	@param curve
//	@param x
//	@param y
//	@return []byte
func CreateEllipticSKI(curve elliptic.Curve, x, y *big.Int) []byte {
	if curve == nil {
		return nil
	}
	//Marshall the public key
	raw := elliptic.Marshal(curve, x, y)
	// Hash it 国密改造后改为sm3
	hash := sm3.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

// GetRandBigInt 随机生成序列号
//
//	@return *big.Int
func GetRandBigInt() *big.Int {
	sn, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		panic(err)
	}
	return sn
}
