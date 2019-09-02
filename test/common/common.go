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

package common

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"encoding/json"
	"fmt"

	client "github.com/tendermint/tendermint/rpc/client"
	"github.com/wesraph/benchmark-tm/abci/did/v1"
	"github.com/wesraph/benchmark-tm/test/utils"
)

const (
	_ = iota
	EcdsaPrivateKey
	RSAPrivateKey
)

func RegisterMasterNode(nodeID, privK string, param did.RegisterMasterNodeParam, keyType int) error {
	fmt.Println("test")
	var privKeyEcdsa *ecdsa.PrivateKey
	var privKeyRSA *rsa.PrivateKey
	var err error

	if keyType == EcdsaPrivateKey {
		privKeyEcdsa, err = utils.GetPrivateKeyFromStringEcdsa(privK)
	} else if keyType == RSAPrivateKey {
		privKeyRSA, err = utils.GetPrivateKeyFromString(privK)
	}

	if err != nil {
		return err
	}

	paramJSON, err := json.Marshal(param)
	if err != nil {
		return err
	}

	fnName := "RegisterMasterNode"
	var nonce string
	var signature []byte

	fmt.Println("Creating signature")
	if keyType == EcdsaPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonceEcdsa(fnName, paramJSON, privKeyEcdsa)
	} else if keyType == RSAPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonce(fnName, paramJSON, privKeyRSA)
	}

	_, err = utils.CreateTxn([]byte(fnName), paramJSON, []byte(nonce), signature, []byte(nodeID))
	if err != nil {
		panic(err)
	}

	//resultObj, _ := result.(utils.ResponseTx)
	//expected := "success"
	//pretty.Println(resultObj)
	//if actual := resultObj.Result.DeliverTx.Log; actual != expected {
	//fmt.Println(actual)
	//return fmt.Errorf("\n"+`CheckTx log: "%s"`, resultObj.Result.CheckTx.Log)
	//}

	return nil
}

func SetTx(nodeID, privK string, param did.SetTxParam, keyType int) error {
	var privKeyEcdsa *ecdsa.PrivateKey
	var privKeyRSA *rsa.PrivateKey
	var err error

	if keyType == EcdsaPrivateKey {
		privKeyEcdsa, err = utils.GetPrivateKeyFromStringEcdsa(privK)
	} else if keyType == RSAPrivateKey {
		privKeyRSA, err = utils.GetPrivateKeyFromString(privK)
	}

	if err != nil {
		return err
	}

	paramJSON, err := json.Marshal(param)
	if err != nil {
		fmt.Println("error:", err)
	}
	fnName := "SetTx"

	var nonce string
	var signature []byte
	if keyType == EcdsaPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonceEcdsa(fnName, paramJSON, privKeyEcdsa)
	} else if keyType == RSAPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonce(fnName, paramJSON, privKeyRSA)
	}

	_, err = utils.CreateTxn([]byte(fnName), paramJSON, []byte(nonce), signature, []byte(nodeID))

	return err
}

func SetTxWebSocket(nodeID, privK string, param did.SetTxParam, keyType int, ws *client.HTTP) error {
	var privKeyEcdsa *ecdsa.PrivateKey
	var privKeyRSA *rsa.PrivateKey
	var err error

	if keyType == EcdsaPrivateKey {
		privKeyEcdsa, err = utils.GetPrivateKeyFromStringEcdsa(privK)
	} else if keyType == RSAPrivateKey {
		privKeyRSA, err = utils.GetPrivateKeyFromString(privK)
	}

	if err != nil {
		return err
	}

	paramJSON, err := json.Marshal(param)
	if err != nil {
		fmt.Println("error:", err)
	}
	fnName := "SetTx"

	var nonce string
	var signature []byte
	if keyType == EcdsaPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonceEcdsa(fnName, paramJSON, privKeyEcdsa)
	} else if keyType == RSAPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonce(fnName, paramJSON, privKeyRSA)
	}

	res, err := utils.CreateTxnWebSocket([]byte(fnName), paramJSON, []byte(nonce), signature, []byte(nodeID), ws)

	if res == nil {
		return err
	}
	if res.Code != 0x0 {
		return fmt.Errorf(res.Log)
	}

	return err
}

func SetValidator(nodeID, privK string, param did.SetValidatorParam, keyType int) error {
	var privKeyEcdsa *ecdsa.PrivateKey
	var privKeyRSA *rsa.PrivateKey
	var err error

	if keyType == EcdsaPrivateKey {
		privKeyEcdsa, err = utils.GetPrivateKeyFromStringEcdsa(privK)
	} else if keyType == RSAPrivateKey {
		privKeyRSA, err = utils.GetPrivateKeyFromString(privK)
	}

	if err != nil {
		panic(err)
	}

	paramJSON, err := json.Marshal(param)
	if err != nil {
		fmt.Println("error:", err)
	}
	fnName := "SetValidator"

	var nonce string
	var signature []byte
	if keyType == EcdsaPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonceEcdsa(fnName, paramJSON, privKeyEcdsa)
	} else if keyType == RSAPrivateKey {
		nonce, signature = utils.CreateSignatureAndNonce(fnName, paramJSON, privKeyRSA)
	}

	_, err = utils.CreateTxn([]byte(fnName), paramJSON, []byte(nonce), signature, []byte(nodeID))

	return err
}
