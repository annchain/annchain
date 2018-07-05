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
)

func (p *Proposal) CString() string {
	pdata := p.GetData()
	return fmt.Sprintf("Proposal{%v/%v %v (%v,%v) %X}", pdata.Height, pdata.Round,
		pdata.BlockPartsHeader, pdata.POLRound, pdata.POLBlockID.CString(), p.Signature)
}

func (pdata *ProposalData) GetBytes(chainID string) (bys []byte, err error) {
	bys, err = proto.Marshal(pdata)
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
