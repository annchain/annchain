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

	merkle "github.com/annchain/annchain/module/lib/go-merkle"
	"github.com/gogo/protobuf/proto"
)

func (blockID *BlockID) IsZero() bool {
	return len(blockID.Hash) == 0 && blockID.PartsHeader.IsZero()
}

func (blockID *BlockID) Equals(other *BlockID) bool {
	if blockID == other {
		return true
	}
	if blockID == nil || other == nil {
		return false
	}
	return bytes.Equal(blockID.Hash, other.Hash) &&
		blockID.PartsHeader.Equals(other.PartsHeader)
}

func (blockID *BlockID) Key() string {
	if blockID == nil {
		return ""
	}
	headerBys, _ := proto.Marshal(blockID.PartsHeader)
	return string(blockID.Hash) + string(headerBys)
}

func (blockID *BlockID) CString() string {
	if blockID == nil {
		return "nil"
	}
	return fmt.Sprintf(`%X:%v`, blockID.Hash, blockID.PartsHeader.CString())
}

/////////////////////////////////////////////////////////////////////////////

// NOTE: hash is nil if required fields are missing.
func (h *Header) Hash() []byte {
	if len(h.ValidatorsHash) == 0 {
		return nil
	}
	return merkle.SimpleHashFromMap(map[string]interface{}{
		"ChainID":            h.ChainID,
		"Height":             h.Height,
		"Time":               h.Time,
		"NumTxs":             h.NumTxs,
		"maker":              h.Maker,
		"LastBlockID":        h.LastBlockID,
		"LastCommit":         h.LastCommitHash,
		"Data":               h.DataHash,
		"Validators":         h.ValidatorsHash,
		"App":                h.AppHash,
		"Receipts":           h.ReceiptsHash,
		"LastNonEmptyHeight": h.LastNonEmptyHeight,
	})
}

func (h *Header) StringIndented(indent string) string {
	if h == nil {
		return "nil-Header"
	}
	return fmt.Sprintf(`Header{
%s  ChainID:        %v
%s  Height:         %v
%s  Time:           %v
%s  NumTxs:         %v
%s  Maker:          %X
%s  LastBlockID:    %v
%s  LastCommit:     %X
%s  Data:           %X
%s  Validators:     %X
%s  App:            %X
%s  Receipts:       %X
%s  LastNonEmptyHeight:       %v
%s}#%X`,
		indent, h.ChainID,
		indent, h.Height,
		indent, h.Time,
		indent, h.NumTxs,
		indent, h.Maker,
		indent, h.LastBlockID,
		indent, h.LastCommitHash,
		indent, h.DataHash,
		indent, h.ValidatorsHash,
		indent, h.AppHash,
		indent, h.ReceiptsHash,
		indent, h.LastNonEmptyHeight,
		indent, h.Hash(),
	)
}
