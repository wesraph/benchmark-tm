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

type RegisterMasterNodeParam struct {
	NodeID          string `json:"node_id"`
	PublicKey       string `json:"public_key"`
	MasterPublicKey string `json:"master_public_key"`
	NodeName        string `json:"node_name"`
}

type SetTxParam struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type SetValidatorParam struct {
	PublicKey string `json:"public_key"`
	Power     int64  `json:"power"`
}

