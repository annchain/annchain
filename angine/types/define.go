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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/annchain/annchain/module/xlib/def"
)

type Bytes []byte

type ValSetLoaderFunc func(height, round def.INT) *ValidatorSet

func (b *Bytes) MarshalJSON() ([]byte, error) {
	bys := strings.ToUpper(hex.EncodeToString(*b))
	return json.Marshal(bys)
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}
	var bys []byte
	bys, err = hex.DecodeString(str)
	if err != nil {
		return err
	}
	(*b) = Bytes(bys)
	return nil
}

func (b *Bytes) Bytes() []byte {
	return []byte(*b)
}

func (b *Bytes) String() string {
	ret, err := b.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("marshal err:%v", err)
	}
	return string(ret)
}

/////////////////////////////////////////////////////////////////

const (
	timeFormart = "2006-01-02 15:04:05"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(&t.Time)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	st := struct {
		time.Time
	}{}
	err := json.Unmarshal(data, &st)
	if err != nil {
		return err
	}
	t.Time = st.Time
	return nil
}

func (t *Time) String() string {
	return t.Format(timeFormart)
}
