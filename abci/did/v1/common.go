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
	"encoding/json"
	"fmt"

	"github.com/ndidplatform/smart-contract/abci/code"
	"github.com/ndidplatform/smart-contract/abci/utils"
	"github.com/ndidplatform/smart-contract/protos/data"
	"github.com/tendermint/tendermint/abci/types"
)

func (app *DIDApplication) RegisterMasterNode(param string, nodeID string) types.ResponseDeliverTx {
	app.logger.Infof("RegisterMasterNode, Parameter: %s", param)
	var funcParam RegisterMasterNodeParam
	err := json.Unmarshal([]byte(param), &funcParam)
	if err != nil {
		return app.ReturnDeliverTxLog(code.UnmarshalError, err.Error(), "")
	}
	key := "NodeID" + "|" + funcParam.NodeID
	// check Duplicate Node ID
	_, chkExists := app.GetStateDB([]byte(key))
	if chkExists != nil {
		return app.ReturnDeliverTxLog(code.DuplicateNodeID, "Duplicate Node ID", "")
	}
	var nodeDetail data.NodeDetail
	nodeDetail.PublicKey = funcParam.PublicKey
	nodeDetail.MasterPublicKey = funcParam.MasterPublicKey
	nodeDetail.NodeName = funcParam.NodeName
	nodeDetail.Active = true
	nodeDetailByte, err := utils.ProtoDeterministicMarshal(&nodeDetail)
	if err != nil {
		return app.ReturnDeliverTxLog(code.MarshalError, err.Error(), "")
	}
	masterNodeKey := "MasterNODE"
	nodeDetailKey := "NodeID" + "|" + nodeID
	app.SetStateDB([]byte(masterNodeKey), []byte(nodeID))
	app.SetStateDB([]byte(nodeDetailKey), []byte(nodeDetailByte))
	return app.ReturnDeliverTxLog(code.OK, "success", "")
}

func (app *DIDApplication) SetTx(param string, nodeID string) types.ResponseDeliverTx {
	app.logger.Infof("SetTx, Parameter: %s", param)
	var funcParam SetTxParam
	err := json.Unmarshal([]byte(param), &funcParam)
	if err != nil {
		return app.ReturnDeliverTxLog(code.UnmarshalError, err.Error(), "")
	}
	app.logger.Infof("Param size: %d\n", len(param))
	key := "Tx" + "|" + string(funcParam.From) + "|" + string(funcParam.To) + "|" + fmt.Sprintf("%f", funcParam.Price) + "|" + fmt.Sprintf("%f", funcParam.Amount)
	app.SetStateDB([]byte(key), []byte(param))
	return app.ReturnDeliverTxLog(code.OK, "success", "")
}
