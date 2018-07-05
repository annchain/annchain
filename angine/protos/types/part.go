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
	"bytes"
	"fmt"

	. "github.com/annchain/annchain/module/lib/go-common"
	"github.com/annchain/annchain/module/lib/go-merkle"
)

func (p *Part) MerkleProof() merkle.SimpleProof {
	return merkle.SimpleProof{
		Aunts: p.GetProof().GetBytes(),
	}
}

///////////////////////////////////////////////////////////////////////////////////

func (psh *PartSetHeader) CString() string {
	if psh == nil {
		return ""
	}
	return fmt.Sprintf("%v:%X", psh.Total, Fingerprint(psh.Hash))
}

func (psh *PartSetHeader) IsZero() bool {
	if psh == nil {
		return true
	}
	return psh.Total == 0
}

func (psh *PartSetHeader) Equals(other *PartSetHeader) bool {
	if psh == other {
		return true
	}
	if psh == nil || other == nil {
		return false
	}
	return psh.Total == other.Total && bytes.Equal(psh.Hash, other.Hash)
}
