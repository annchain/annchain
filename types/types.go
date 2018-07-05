// Copyright 2017 Annchain Information Technology Services Co.,Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package types

import (
	"math/big"
)

var (
	big0 = big.NewInt(0)
)

var (
	TxTagApp = []byte{1}

	TxTagAppEvm       = []byte{1, 1}
	TxTagAppEvmCommon = []byte{1, 1, 1}

	TxTagAppEco              = []byte{1, 2}
	TxTagAppEcoShareTransfer = []byte{1, 2, 1}
	TxTagAppEcoGuarantee     = []byte{1, 2, 2}
	TxTagAppEcoRedeem        = []byte{1, 2, 3}
)

var (
	TxTagAngine = []byte{2}

	TxTagAngineEco        = []byte{2, 1}
	TxTagAngineEcoSuspect = []byte{2, 1, 1}

	TxTagAngineInit      = []byte{2, 2}
	TxTagAngineInitToken = []byte{2, 2, 1}
	TxTagAngineInitShare = []byte{2, 2, 2}

	TxTagAngineWorld     = []byte{2, 3}
	TxTagAngineWorldRand = []byte{2, 3, 1}
)

const (
	CODE_VAR_ENT = "ent_params"
	CODE_VAR_RET = "ret_params"
)

func BigInt0() *big.Int {
	return big0
}

const (
	QueryTypeContract = iota
	QueryTypeNonce
	QueryTypeBalance
	QueryTypeReceipt
	QueryTypeContractExistance
	QueryTypeShare
)

type QueryShareResult struct {
	ShareBalance  *big.Int
	ShareGuaranty *big.Int
	GHeight       uint64
}

type WorldRandVote struct {
	Pubkey []byte
	Sig    []byte
}
