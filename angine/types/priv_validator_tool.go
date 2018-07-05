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
	"errors"
	"fmt"
	"os"
	"path/filepath"

	crypto "github.com/annchain/annchain/module/lib/go-crypto"
	"github.com/annchain/annchain/module/xlib"
	"github.com/annchain/annchain/module/xlib/def"
)

var (
	ErrFileNotFound       = errors.New("priv_validator.json not found")
	ErrBranchIsUsed       = errors.New("priv_validator:branch name is used")
	ErrPVRevertFromBackup = errors.New("priv_validator:revert from backup, not find data")
)

const (
	PRIV_FILE_NAME = "priv_validator.json"
)

type PrivValidatorTool struct {
	dir string
	pv  *PrivValidator
}

func (pt *PrivValidatorTool) Init(dir string) error {
	pt.pv = LoadPrivValidator(nil, dir)
	if pt.pv == nil {
		return ErrFileNotFound
	}
	return nil
}

func (pt *PrivValidatorTool) backupName(branchName string) string {
	return fmt.Sprintf("%v/%v-%v.json", filepath.Dir(pt.pv.filePath), PRIV_FILE_NAME, branchName)
}

func (pt *PrivValidatorTool) BackupData(branchName string) error {
	bkName := pt.backupName(branchName)
	find, err := xlib.PathExists(bkName)
	if err != nil {
		return err
	}
	if find {
		return ErrBranchIsUsed
	}
	preDir := pt.pv.filePath
	pt.pv.SetFile(bkName)
	pt.pv.Save()
	pt.pv.SetFile(preDir)
	return nil
}

func (pt *PrivValidatorTool) RevertFromBackup(branchName string) error {
	bkName := pt.backupName(branchName)
	find, err := xlib.PathExists(bkName)
	if err != nil {
		return err
	}
	if !find {
		return ErrPVRevertFromBackup
	}
	xlib.CopyFile(pt.pv.filePath, bkName)
	return nil
}

func (pt *PrivValidatorTool) DelBackup(branchName string) {
	os.Remove(pt.backupName(branchName))
}

func (pt *PrivValidatorTool) SaveNewPrivV(toHeight def.INT) error {
	pt.pv.LastHeight = toHeight
	pt.pv.LastRound = 0
	pt.pv.LastStep = 0
	pt.pv.LastSignature = crypto.StSignature{nil}
	pt.pv.LastSignBytes = make([]byte, 0)
	pt.pv.Save()
	return nil
}
