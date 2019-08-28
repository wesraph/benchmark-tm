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
	"github.com/blockfint/benchmark-tm/abci/code"
	"github.com/tendermint/tendermint/abci/types"
)

// ReturnQuery return types.ResponseQuery
func (app *DIDApplication) ReturnQuery(value []byte, log string, height int64) types.ResponseQuery {
	app.logger.Infof("Query result: %s", string(value))
	var res types.ResponseQuery
	res.Value = value
	res.Log = log
	res.Height = height
	return res
}

// QueryRouter is Pointer to function
func (app *DIDApplication) QueryRouter(method string, param string, height int64) types.ResponseQuery {
	result := app.callQuery(method, param, height)
	return result
}

func (app *DIDApplication) callQuery(name string, param string, height int64) types.ResponseQuery {
	switch name {
	default:
		return types.ResponseQuery{Code: code.UnknownMethod, Log: "Unknown method name"}
	}
}

