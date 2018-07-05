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
	"math/big"
	"testing"

	"github.com/annchain/annchain/eth/crypto"
	"github.com/annchain/annchain/types"
)

func TestSig(t *testing.T) {
	privkey, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	tx := &types.BlockTx{
		GasLimit: big.NewInt(444),
		GasPrice: big.NewInt(444),
		Sender:   addr[:],
		Payload:  []byte("xxxx"),
	}

	if err := TxSign(tx, privkey); err != nil {
		t.Fatal(err)
	}

	bs, err := TxToBytes(tx)
	if err != nil {
		t.Fatal(err)
	}
	tx2 := &types.BlockTx{}
	if err = TxFromBytes(bs, tx2); err != nil {
		t.Fatal(err)
	}

	valid, err := TxVerifySignature(tx2)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatal("valid error")
	}
}
