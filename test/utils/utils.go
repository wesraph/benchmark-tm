package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	mathRand "math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	protoTm "github.com/wesraph/benchmark-tm/protos/tendermint"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/libs/common"
)

var tendermintAddr = GetEnv("TENDERMINT_ADDRESS", "http://localhost:45000")

func GetPrivateKeyFromStringEcdsa(privK string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privK))
	x509Encoded := block.Bytes
	return x509.ParseECPrivateKey(x509Encoded)
}

func GetPrivateKeyFromString(privK string) (*rsa.PrivateKey, error) {
	privK = strings.Replace(privK, "\t", "", -1)
	block, _ := pem.Decode([]byte(privK))
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func GeneratePublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	privBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubKeyBytes,
	}
	publicPEM := pem.EncodeToMemory(&privBlock)
	return publicPEM, nil
}

func CreateSignatureAndNonce(fnName string, paramJSON []byte, privKey *rsa.PrivateKey) (nonce string, signature []byte) {
	nonce = base64.StdEncoding.EncodeToString([]byte(common.RandStr(12)))
	tempPSSmessage := append([]byte(fnName), paramJSON...)
	tempPSSmessage = append(tempPSSmessage, []byte(nonce)...)
	PSSmessage := []byte(base64.StdEncoding.EncodeToString(tempPSSmessage))
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, newhash, hashed)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nonce, signature
}

func CreateSignatureAndNonceEcdsa(fnName string, paramJSON []byte, privKey *ecdsa.PrivateKey) (nonce string, signature []byte) {
	nonce = base64.StdEncoding.EncodeToString([]byte(common.RandStr(12)))

	// Generating message
	tempPSSmessage := append([]byte(fnName), paramJSON...)
	tempPSSmessage = append(tempPSSmessage, []byte(nonce)...)
	PSSmessage := []byte(base64.StdEncoding.EncodeToString(tempPSSmessage))

	// Hash the message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	signhash := pssh.Sum(nil)

	r := big.NewInt(0)
	s := big.NewInt(0)
	r, s, serr := ecdsa.Sign(rand.Reader, privKey, signhash)
	if serr != nil {
		panic(serr)
	}

	type SignValues struct {
		R BigInt
		S BigInt
	}

	var encBigInt SignValues
	encBigInt.R.Int = *r
	encBigInt.S.Int = *s

	signature, err := json.Marshal(encBigInt)
	if err != nil {
		panic(err)
	}

	return nonce, signature
}

func CreateTxn(fnName []byte, param []byte, nonce []byte, signature []byte, nodeID []byte) (interface{}, error) {
	var tx protoTm.Tx
	tx.Method = string(fnName)
	tx.Params = string(param)
	tx.Nonce = nonce
	tx.Signature = signature
	tx.NodeId = string(nodeID)
	txByte, err := proto.Marshal(&tx)
	if err != nil {
		log.Printf("err: %s", err.Error())
	}
	txEncoded := hex.EncodeToString(txByte)
	var URL *url.URL
	URL, err = url.Parse(tendermintAddr)
	if err != nil {
		panic("boom")
	}
	URL.Path += "/broadcast_tx_commit"
	parameters := url.Values{}
	parameters.Add("tx", `0x`+txEncoded)
	URL.RawQuery = parameters.Encode()
	encodedURL := URL.String()
	req, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body ResponseTx
	json.NewDecoder(resp.Body).Decode(&body)
	return body, nil
}

func Query(fnName []byte, param []byte) (interface{}, error) {
	var data protoTm.Query
	data.Method = string(fnName)
	data.Params = string(param)
	dataByte, err := proto.Marshal(&data)
	if err != nil {
		log.Printf("err: %s", err.Error())
	}
	dataEncoded := hex.EncodeToString(dataByte)
	var URL *url.URL
	URL, err = url.Parse(tendermintAddr)
	if err != nil {
		panic("boom")
	}
	URL.Path += "/abci_query"
	parameters := url.Values{}
	parameters.Add("data", `0x`+dataEncoded)
	URL.RawQuery = parameters.Encode()
	encodedURL := URL.String()
	req, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var body ResponseQuery
	json.NewDecoder(resp.Body).Decode(&body)
	return body, nil
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	mathRand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mathRand.Intn(len(letterRunes))]
	}
	return string(b)
}

type ResponseTx struct {
	Result struct {
		Height  int `json:"height"`
		CheckTx struct {
			Code int      `json:"code"`
			Log  string   `json:"log"`
			Fee  struct{} `json:"fee"`
		} `json:"check_tx"`
		DeliverTx struct {
			Log  string   `json:"log"`
			Fee  struct{} `json:"fee"`
			Tags []common.KVPair
		} `json:"deliver_tx"`
		Hash string `json:"hash"`
	} `json:"result"`
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
}

type ResponseQuery struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Response struct {
			Log    string `json:"log"`
			Value  string `json:"value"`
			Height string `json:"height"`
		} `json:"response"`
	} `json:"result"`
}

func EncodeEcdsa(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func DecodeEcdsa(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

type BigInt struct {
	big.Int
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *BigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Int
	_, ok := z.SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = z
	return nil
}
