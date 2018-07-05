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


package utils

import (
	"sync"

	"github.com/annchain/annchain/module/xlib/def"
)

var data sync.Map

type AngineData struct {
	Height def.INT // LastBlockHeight of state
}

func UpdateHeight(chainID string, height def.INT) {
	d, _ := Load(chainID)
	d.Height = height
	Store(chainID, &d)
}

func LoadHeight(chainID string) def.INT {
	d, _ := Load(chainID)
	return d.Height
}

func Store(chainID string, d *AngineData) {
	data.Store(chainID, d)
}

func Load(chainID string) (AngineData, bool) {
	if pd, has := data.Load(chainID); has {
		return *(pd.(*AngineData)), true
	}
	return AngineData{}, false
}

func Delete(chainID string) {
	data.Delete(chainID)
}

func Has(chainID string) bool {
	_, has := data.Load(chainID)
	return has
}
