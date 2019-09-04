/**
 * Copyright (c) 2018, 2019 National Digital ID COMPANY LIMITED
 *
 * This file is part of NDID software.
 *
 * NDID is the free software: you can redistribute it and/or modify it under
 * the terms of the Affero GNU General Public License as published by the
 * Free Software Foundation, either version 3 of the License, or any later
 * version.
 *
 * NDID is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the Affero GNU General Public License for more details.
 *
 * You should have received a copy of the Affero GNU General Public License
 * along with the NDID source code. If not, see https://www.gnu.org/licenses/agpl.txt.
 *
 * Please contact info@ndid.co.th for any further questions
 *
 */

package did

import (
	"crypto"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"encoding/asn1"
	"fmt"
	"math/big"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/wesraph/benchmark-tm/abci/code"
	"github.com/wesraph/benchmark-tm/protos/data"
)

var IsMethod = map[string]bool{
	"RegisterMasterNode": true,
	"SetTx":              true,
	"SetValidator":       true,
}

type BigInt struct {
	big.Int
}

type SignValues struct {
	R, S *big.Int
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

func (app *DIDApplication) checkCanRegisterMasterNode(param string, nodeID string) types.ResponseCheckTx {
	key := "MasterNODE"
	exist := app.HasStateDB([]byte(key))
	if exist {
		return ReturnCheckTx(code.MasterNodeIsAlreadyExisted, "Master node is already existed")
	}
	return ReturnCheckTx(code.OK, "")
}

func (app *DIDApplication) checkCanSetTx(param string, nodeID string) types.ResponseCheckTx {
	//Bypass
	return ReturnCheckTx(code.OK, "")
}

func verifySignature(param string, nonce []byte, signature []byte, publicKey string, method string) (result bool, err error) {
	// Get the certificate
	publicKey = strings.Replace(publicKey, "\t", "", -1)
	block, _ := pem.Decode([]byte(publicKey))

	// Get the key
	senderPublicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	//senderPublicKey := senderPublicKeyInterface.(*ecdsa.PublicKey)
	if err != nil {
		return false, err
	}

	// Build the message and encode to base64
	tempPSSmessage := append([]byte(method), []byte(param)...)
	tempPSSmessage = append(tempPSSmessage, []byte(nonce)...)
	PSSmessage := []byte(base64.StdEncoding.EncodeToString(tempPSSmessage))

	// Hash the message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	switch pubKey := senderPublicKeyInterface.(type) {
	case *ecdsa.PublicKey:
		var signVal SignValues
		_, err = asn1.Unmarshal(signature, &signVal)
		if err != nil {
			return false, nil
		}

		return ecdsa.Verify(pubKey, hashed, signVal.R, signVal.S), nil
	case *rsa.PublicKey:
		err = rsa.VerifyPKCS1v15(pubKey, newhash, hashed, signature)
		return err == nil, err
	}

	return false, fmt.Errorf("Unsupported keytype")
}

// ReturnCheckTx return types.ResponseDeliverTx
func ReturnCheckTx(code uint32, log string) types.ResponseCheckTx {
	return types.ResponseCheckTx{
		Code: code,
		Log:  fmt.Sprintf(log),
	}
}

func getPublicKeyRegisterMasterNode(param string) string {
	var funcParam RegisterMasterNodeParam
	err := json.Unmarshal([]byte(param), &funcParam)
	if err != nil {
		return ""
	}
	return funcParam.PublicKey
}

func (app *DIDApplication) getMasterPublicKeyFromNodeID(nodeID string) string {
	key := "NodeID" + "|" + nodeID
	_, value := app.GetStateDB([]byte(key))
	if value == nil {
		return ""
	}
	var nodeDetail data.NodeDetail
	err := proto.Unmarshal(value, &nodeDetail)
	if err != nil {
		return ""
	}
	return nodeDetail.MasterPublicKey
}

func (app *DIDApplication) getPublicKeyFromNodeID(nodeID string) string {
	key := "NodeID" + "|" + nodeID
	_, value := app.GetStateDB([]byte(key))
	if value == nil {
		return ""
	}
	var nodeDetail data.NodeDetail
	err := proto.Unmarshal(value, &nodeDetail)
	if err != nil {
		return ""
	}
	return nodeDetail.PublicKey
}

func checkPubKey(key string) (returnCode uint32, log string) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return code.InvalidKeyFormat, "Invalid key format. Cannot decode PEM."
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return code.InvalidKeyFormat, err.Error()
	}

	switch pubKey := pub.(type) {
	case *rsa.PublicKey:
		if pubKey.N.BitLen() < 2048 {
			return code.RSAKeyLengthTooShort, "RSA key length is too short. Must be at least 2048-bit."
		}
	case *ecdsa.PublicKey:
		return code.OK, ""
	case *dsa.PublicKey:
		return code.UnsupportedKeyType, "Unsupported key type. Only RSA is allowed."
	default:
		return code.UnknownKeyType, "Unknown key type. Only RSA is allowed."
	}
	return code.OK, ""
}

func checkNodePubKeys(param string) (returnCode uint32, log string) {
	var keys struct {
		MasterPublicKey string `json:"master_public_key"`
		PublicKey       string `json:"public_key"`
	}
	err := json.Unmarshal([]byte(param), &keys)
	if err != nil {
		return code.UnmarshalError, err.Error()
	}

	// Validate master public key format
	if keys.MasterPublicKey != "" {
		returnCode, log = checkPubKey(keys.MasterPublicKey)
		if returnCode != code.OK {
			return returnCode, log
		}
	}

	// Validate public key format
	if keys.PublicKey != "" {
		returnCode, log = checkPubKey(keys.PublicKey)
		if returnCode != code.OK {
			return returnCode, log
		}
	}
	return code.OK, ""
}

func checkAccessorPubKey(param string) (returnCode uint32, log string) {
	var key struct {
		AccessorPublicKey string `json:"accessor_public_key"`
	}
	err := json.Unmarshal([]byte(param), &key)
	if err != nil {
		return code.UnmarshalError, err.Error()
	}
	returnCode, log = checkPubKey(key.AccessorPublicKey)
	if returnCode != code.OK {
		return returnCode, log
	}
	return code.OK, ""
}

// CheckTxRouter is Pointer to function
func (app *DIDApplication) CheckTxRouter(method string, param string, nonce []byte, signature []byte, nodeID string) types.ResponseCheckTx {

	var publicKey string
	if method == "RegisterMasterNode" {
		publicKey = getPublicKeyRegisterMasterNode(param)
		if publicKey == "" {
			return ReturnCheckTx(code.CannotGetPublicKeyFromParam, "Can not get public key from parameter")
		}
	} else {
		publicKey = app.getPublicKeyFromNodeID(nodeID)
		if publicKey == "" {
			return ReturnCheckTx(code.CannotGetPublicKeyFromNodeID, "Can not get public key from node ID")
		}
	}

	// Check pub key
	if method == "RegisterMasterNode" {
		checkCode, log := checkNodePubKeys(param)
		if checkCode != code.OK {
			return ReturnCheckTx(checkCode, log)
		}
	}

	verifyResult, err := verifySignature(param, nonce, signature, publicKey, method)
	if err != nil || verifyResult == false {
		return ReturnCheckTx(code.VerifySignatureError, err.Error())
	}

	var result types.ResponseCheckTx
	result = app.callCheckTx(method, param, nodeID)
	return result
}

func (app *DIDApplication) callCheckTx(name string, param string, nodeID string) types.ResponseCheckTx {
	switch name {
	case "RegisterMasterNode":
		return app.checkCanRegisterMasterNode(param, nodeID)
	case "SetTx":
		return app.checkCanSetTx(param, nodeID)
	case "SetValidator":
		return app.checkCanSetTx(param, nodeID)
	default:
		return types.ResponseCheckTx{Code: code.UnknownMethod, Log: "Unknown method name"}
	}
}

func (app *DIDApplication) isDuplicateNonce(nonce []byte) bool {
	return app.HasStateDB(nonce)
}
