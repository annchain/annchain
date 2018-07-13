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

package tools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/annchain/annchain/types"
)

func TestSig(t *testing.T) {
	s := "AQEB+HeDAV+QAoCU19qwXxIaYZnE2CoBQrc9atlVajmY15Qt52KZvonNe3IpHjXVYXRAK9aefoASuEEjiIgLBpDsqKzbD6WKeuR0LCU8jtebPpyF5fF6tWTJBTg3Us0yyexE6HkQuTrM45zbLZsEc7i2SUdqNdeUpQthAQ=="
	fmt.Println(len(s), len(s)/4*3)
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))

	b := make([]byte, len(s)/4*3)
	n, err := r.Read(b)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Printf("=====%x\n", b[3:5])
	fmt.Println(string(b))

	tx := &types.BlockTx{}
	err = FromBytes(b[3:], tx)
	if err != nil {
		panic(err)
	}

	fmt.Println(tx)
}

func TestFoo(t *testing.T) {
	// s := "\"AQEB+HeDAV+QAoCU19qwXxIaYZnE2CoBQrc9atlVajmY15Qt52KZvonNe3IpHjXVYXRAK9aefoASuEEjiIgLBpDsqKzbD6WKeuR0LCU8jtebPpyF5fF6tWTJBTg3Us0yyexE6HkQuTrM45zbLZsEc7i2SUdqNdeUpQthAQ==\""
	s := "\"AgIBeyJ0byI6IjE5cXdYeElhWVpuRTJDb0JRcmM5YXRsVmFqaz0iLCJhbW91bnQiOjEwMDAwMDAwMDAwMCwiZXh0cmEiOiJhR1ZzYkc4eE1URXhNUT09In0=\""

	var bs []byte
	err := json.Unmarshal([]byte(s), &bs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", bs)

	tx := &types.BlockTx{}
	err = FromBytes(bs[3:], tx)
	if err != nil {
		panic(err)
	}

	fmt.Println(tx)

}
