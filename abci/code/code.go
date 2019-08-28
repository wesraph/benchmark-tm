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

package code

// Return codes for return result
const (
	OK                           uint32 = 0
	EncodingError                uint32 = 1
	DecodingError                uint32 = 2
	BadNonce                     uint32 = 3
	Unauthorized                 uint32 = 4
	UnmarshalError               uint32 = 5
	MarshalError                 uint32 = 6
	MasterNodeIsAlreadyExisted   uint32 = 7
	InvalidKeyFormat             uint32 = 8
	RSAKeyLengthTooShort         uint32 = 9
	UnsupportedKeyType           uint32 = 10
	UnknownKeyType               uint32 = 11
	CannotGetPublicKeyFromParam  uint32 = 12
	VerifySignatureError         uint32 = 13
	DuplicateNodeID              uint32 = 14
	DuplicateNonce               uint32 = 15
	MethodCanNotBeEmpty          uint32 = 16
	InvalidTransactionFormat     uint32 = 17
	CannotGetPublicKeyFromNodeID uint32 = 18
	UnknownMethod                uint32 = 998
	UnknownError                 uint32 = 999
)
