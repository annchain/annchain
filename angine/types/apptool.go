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
)

const (
	databaseCache   = 128
	databaseHandles = 1024
)

var (
	ErrRevertFromBackup = errors.New("revert from backup,not find data")
	ErrBranchNameUsed   = errors.New("app:branch name has been used")
)

type BaseAppTool struct {
	BaseApplication
}

func (t *BaseAppTool) backupName(branchName string) []byte {
	return []byte(fmt.Sprintf("%s-%s", lastBlockKey, branchName))
}

func (t *BaseAppTool) RevertFromBackup(branchName string) error {
	preKeyName := t.backupName(branchName)
	bs := t.Database.Get(preKeyName)
	if len(bs) == 0 {
		return ErrRevertFromBackup
	}
	t.Database.Set(lastBlockKey, bs)
	return nil
}

func (t *BaseAppTool) DelBackup(branchName string) {
	t.Database.Delete(t.backupName(branchName))
}

func (t *BaseAppTool) BackupLastBlockData(branchName string, lastBlock interface{}) error {
	preKeyName := t.backupName(branchName)
	dataBs := t.Database.Get(preKeyName)
	if len(dataBs) > 0 {
		return ErrBranchNameUsed
	}
	t.SaveLastBlockByKey(preKeyName, lastBlock)
	return nil
}
