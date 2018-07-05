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
	"fmt"

	"github.com/gogo/protobuf/proto"

	. "github.com/annchain/annchain/module/lib/go-common"
)

// TODO Make a new type "VoteType"

const (
	VoteTypePrevote   = byte(0x01)
	VoteTypePrecommit = byte(0x02)
)

func IsVoteTypeValid(type_ VoteType) bool {
	switch type_ {
	case VoteType_Prevote:
		return true
	case VoteType_Precommit:
		return true
	default:
		return false
	}
}

func (vote *Vote) Copy() *Vote {
	voteCopy := *vote
	return &voteCopy
}

func (vote *Vote) CString() string {
	if !vote.Exist() {
		return "nil-Vote"
	}
	vdata := vote.Data
	var typeString string
	switch vdata.Type {
	case VoteType_Prevote:
		typeString = "Prevote"
	case VoteType_Precommit:
		typeString = "Precommit"
	default:
		PanicSanity("Unknown vote type")
	}

	return fmt.Sprintf("Vote{%v:%X %v/%02d/%v(%v) %X %v}",
		vdata.ValidatorIndex, Fingerprint(vdata.ValidatorAddress),
		vdata.Height, vdata.Round, vdata.Type, typeString,
		Fingerprint(vdata.BlockID.Hash), vote.Signature)
}

func (vote *Vote) Exist() bool {
	return vote != nil && vote.Data != nil
}

///////////////////////////////////////////////////////////////////////////////////

func (vdata *VoteData) GetBytes(chainID string) (bys []byte, err error) {
	bys, err = proto.Marshal(vdata)
	if err != nil {
		return nil, err
	}
	st := SignableBase{
		ChainID: chainID,
		Data:    bys,
	}
	bys, err = proto.Marshal(&st)
	return bys, err
}
