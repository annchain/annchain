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
	"github.com/pkg/errors"
	pbtypes "github.com/annchain/annchain/angine/protos/types"
)

type (
	ErrWrongChainID        error
	ErrWrongHeight         error
	ErrWrongNonEmptyHeight error
	ErrWrongNumTxs         error
	ErrWrongBlockID        error
	ErrWrongLastCommitHash error
	ErrWrongDataHash       error
	ErrWrongAppHash        error
	ErrWrongReceiptsHash   error
	ErrWrongBasic          error
)

func NewErrWrongChainID(expected, got string) ErrWrongChainID {
	return ErrWrongChainID(errors.Errorf("Wrong Block.Header.ChainID. Expected %v, got %v", expected, got))
}

func NewErrWrongHeight(chainID string, expected, got int) ErrWrongHeight {
	return ErrWrongHeight(errors.Errorf("(%s) Wrong Block.Header.Height. Expected %v, got %v", chainID, expected, got))
}

func NewErrWrongNonEmptyHeight(chainID string, expected, got int) ErrWrongNonEmptyHeight {
	return ErrWrongNonEmptyHeight(errors.Errorf("(%s) Wrong Block.Header.LastNonEmptyHeight. Expected %v, got %v", chainID, expected, got))
}

func NewErrWrongNumTxs(chainID string, expected, got int) ErrWrongNumTxs {
	return ErrWrongNumTxs(errors.Errorf("(%s) Wrong Block.Header.NumTxs. Expected %v, got %v", chainID, expected, got))
}

func NewErrWrongBlockID(chainID string, expected, got pbtypes.BlockID) ErrWrongBlockID {
	return ErrWrongBlockID(errors.Errorf("(%s) Wrong Block.Header.LastBlockID.  Expected %v, got %v", chainID, expected, got))
}

func NewErrWrongLastCommitHash(chainID string, expected, got []byte) ErrWrongLastCommitHash {
	return ErrWrongLastCommitHash(errors.Errorf("(%s) Wrong Block.Header.LastCommitHash.  Expected %X, got %X", chainID, expected, got))
}

func NewErrWrongDataHash(chainID string, expected, got []byte) ErrWrongDataHash {
	return ErrWrongDataHash(errors.Errorf("(%s) Wrong Block.Header.DataHash.  Expected %X, got %X", chainID, expected, got))
}

func NewErrWrongAppHash(chainID string, expected, got []byte) ErrWrongAppHash {
	return ErrWrongAppHash(errors.Errorf("(%s) Wrong Block.Header.AppHash.  Expected %X, got %X", chainID, expected, got))
}

func NewErrWrongReceiptsHash(chainID string, expected, got []byte) ErrWrongReceiptsHash {
	return ErrWrongReceiptsHash(errors.Errorf("(%s) Wrong Block.Header.ReceiptsHash.  Expected %X, got %X", chainID, expected, got))
}

func NewErrWrongBasic(err error) ErrWrongBasic {
	return ErrWrongBasic(err)
}
