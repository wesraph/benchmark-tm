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
	"encoding/base64"
	"fmt"

	"github.com/blockfint/benchmark-tm/abci/code"
	"github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
)

// app.ReturnDeliverTxLog return types.ResponseDeliverTx
func (app *DIDApplication) ReturnDeliverTxLog(code uint32, log string, extraData string) types.ResponseDeliverTx {
	var tags []cmn.KVPair
	if code == 0 {
		var tag cmn.KVPair
		tag.Key = []byte("success")
		tag.Value = []byte("true")
		tags = append(tags, tag)
	} else {
		var tag cmn.KVPair
		tag.Key = []byte("success")
		tag.Value = []byte("false")
		tags = append(tags, tag)
	}
	return types.ResponseDeliverTx{
		Code: code,
		Log:  fmt.Sprintf(log),
		Data: []byte(extraData),
		Tags: tags,
	}
}

func (app *DIDApplication) ReturnDeliverTxLogWitgTag(code uint32, log string, specialTag []cmn.KVPair) types.ResponseDeliverTx {
	var tags []cmn.KVPair
	if code == 0 {
		var tag cmn.KVPair
		tag.Key = []byte("success")
		tag.Value = []byte("true")
		tags = append(tags, tag)
	} else {
		var tag cmn.KVPair
		tag.Key = []byte("success")
		tag.Value = []byte("false")
		tags = append(tags, tag)
	}
	tags = append(tags, specialTag...)
	return types.ResponseDeliverTx{
		Code: code,
		Log:  fmt.Sprintf(log),
		Data: []byte(""),
		Tags: tags,
	}
}

// DeliverTxRouter is Pointer to function
func (app *DIDApplication) DeliverTxRouter(method string, param string, nonce []byte, signature []byte, nodeID string) types.ResponseDeliverTx {
	// ---- check authorization ----
	checkTxResult := app.CheckTxRouter(method, param, nonce, signature, nodeID)
	if checkTxResult.Code != code.OK {
		if checkTxResult.Log != "" {
			return app.ReturnDeliverTxLog(checkTxResult.Code, checkTxResult.Log, "")
		}
		return app.ReturnDeliverTxLog(checkTxResult.Code, "Unauthorized", "")
	}
	result := app.callDeliverTx(method, param, nodeID)

	// Set used nonce to stateDB
	emptyValue := make([]byte, 0)
	app.SetStateDB([]byte(nonce), emptyValue)
	nonceBase64 := base64.StdEncoding.EncodeToString(nonce)
	app.deliverTxNonceState[nonceBase64] = []byte(nil)
	return result
}

func (app *DIDApplication) callDeliverTx(name string, param string, nodeID string) types.ResponseDeliverTx {
	switch name {
	case "RegisterMasterNode":
		return app.RegisterMasterNode(param, nodeID)
	case "SetTx":
		return app.SetTx(param, nodeID)
	case "SetValidator":
		return app.setValidator(param, nodeID)
	default:
		return types.ResponseDeliverTx{Code: code.UnknownMethod, Log: "Unknown method name"}
	}
}
