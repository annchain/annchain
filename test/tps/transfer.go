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

package main

import (
	"errors"
	"fmt"
	"math/big"

	agtypes "github.com/annchain/annchain/angine/types"
	ac "github.com/annchain/annchain/module/lib/go-common"
	cl "github.com/annchain/annchain/module/lib/go-rpc/client"
	"github.com/annchain/annchain/tools"
	"github.com/annchain/annchain/types"
	"github.com/annchain/anth/common"
	"github.com/annchain/anth/crypto"
)

func send(client *cl.ClientJSONRPC, privkey, toAddr string, value int64, nonce uint64) error {
	sk, err := crypto.HexToECDSA(ac.SanitizeHex(privkey))
	panicErr(err)

	btxbs, err := tools.ToBytes(&types.TxEvmCommon{
		To:     common.HexToAddress(toAddr).Bytes(),
		Amount: big.NewInt(value),
	})
	panicErr(err)

	tx := types.NewBlockTx(gasLimit, big.NewInt(0), nonce, crypto.PubkeyToAddress(sk.PublicKey).Bytes(), btxbs)
	tx.Signature, err = tools.SignSecp256k1(tx, crypto.FromECDSA(sk))
	panicErr(err)
	b, err := tools.ToBytes(tx)
	panicErr(err)

	res := new(agtypes.ResultBroadcastTx)
	if client == nil {
		client = cl.NewClientJSONRPC(logger, rpcTarget)
	}
	_, err = client.Call("broadcast_tx_sync", []interface{}{append(types.TxTagAppEvmCommon, b...)}, res)
	panicErr(err)

	if res.Code != 0 {
		fmt.Println(res.Code, string(res.Data), res.Log)
		return errors.New(string(res.Data))
	}

	return nil
}
