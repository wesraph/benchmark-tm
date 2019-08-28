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

func (app *DIDApplication) SetStateDB(key, value []byte) {
	app.HashData = append(app.HashData, key...)
	app.HashData = append(app.HashData, value...)

	app.UncommittedState[string(key)] = value
}

func (app *DIDApplication) GetStateDB(key []byte) (err error, value []byte) {
	var existInUncommittedState bool
	value, existInUncommittedState = app.UncommittedState[string(key)]
	if !existInUncommittedState {
		value = app.state.db.Get(key)
	}

	return nil, value
}

func (app *DIDApplication) GetCommittedStateDB(key []byte) (err error, value []byte) {
	value = app.state.db.Get(key)
	return nil, value
}

func (app *DIDApplication) HasStateDB(key []byte) bool {
	_, existInUncommittedState := app.UncommittedState[string(key)]
	if existInUncommittedState {
		return true
	}
	return app.state.db.Has(key)
}

func (app *DIDApplication) HasVersionedStateDB(key []byte) bool {
	versionsKeyStr := string(key) + "|versions"
	versionsKey := []byte(versionsKeyStr)

	_, existInUncommittedState := app.UncommittedVersionsState[versionsKeyStr]
	if existInUncommittedState {
		return true
	}

	return app.state.db.Has(versionsKey)
}

func (app *DIDApplication) DeleteStateDB(key []byte) {
	if !app.HasStateDB(key) {
		return
	}
	app.HashData = append(app.HashData, key...)
	app.HashData = append(app.HashData, []byte("delete")...) // Remove or replace with something else?

	app.UncommittedState[string(key)] = nil
}

func (app *DIDApplication) SaveDBState() {
	batch := app.state.db.NewBatch()
	defer batch.Close()

	for key := range app.UncommittedState {
		value := app.UncommittedState[key]
		if value != nil {
			batch.Set([]byte(key), value)
		} else {
			batch.Delete([]byte(key))
		}
	}
	batch.WriteSync()
	app.UncommittedState = make(map[string][]byte)
	app.UncommittedVersionsState = make(map[string][]int64)
}
