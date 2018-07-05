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
	"reflect"
	"strings"

	agtypes "github.com/annchain/annchain/angine/types"
	"github.com/annchain/annchain/eth/accounts/abi"
	"github.com/annchain/annchain/eth/common"
	ethtypes "github.com/annchain/annchain/eth/core/types"
	"github.com/annchain/annchain/eth/crypto"
	"github.com/annchain/annchain/eth/rlp"
	ac "github.com/annchain/annchain/module/lib/go-common"
	cl "github.com/annchain/annchain/module/lib/go-rpc/client"
	"github.com/annchain/annchain/tools"
	"github.com/annchain/annchain/types"
)

func createContract(client *cl.ClientJSONRPC, privkey, bytecode string, nonce uint64) (string, error) {
	sk, err := crypto.HexToECDSA(ac.SanitizeHex(privkey))
	panicErr(err)

	btxbs, err := tools.ToBytes(&types.TxEvmCommon{
		Amount: new(big.Int),
		Load:   common.Hex2Bytes(bytecode),
	})
	panicErr(err)

	tx := types.NewBlockTx(gasLimit, big.NewInt(0), nonce, crypto.PubkeyToAddress(sk.PublicKey).Bytes(), btxbs)
	tx.Signature, err = tools.SignSecp256k1(tx, crypto.FromECDSA(sk))
	panicErr(err)
	b, err := tools.ToBytes(tx)
	panicErr(err)

	// tmResult := new(agtypes.TMResult)
	res := new(agtypes.ResultBroadcastTx)
	if client == nil {
		client = cl.NewClientJSONRPC(logger, rpcTarget)
	}
	_, err = client.Call("broadcast_tx_sync", []interface{}{append(types.TxTagAppEvmCommon, b...)}, res)
	panicErr(err)

	if res.Code != 0 {
		fmt.Println(res.Code, string(res.Data))
		return "", errors.New(string(res.Data))
	}

	fmt.Println(res.Code, string(res.Data))
	caller := crypto.PubkeyToAddress(crypto.ToECDSA(common.Hex2Bytes(privkey)).PublicKey)
	addr := crypto.CreateAddress(caller, nonce)
	fmt.Println("contract address:", addr.Hex())

	return addr.Hex(), nil
}

func executeContract(client *cl.ClientJSONRPC, privkey, contract, abijson, callfunc string, args []interface{}, nonce uint64) error {
	sk, err := crypto.HexToECDSA(ac.SanitizeHex(privkey))
	panicErr(err)

	aabbii, err := abi.JSON(strings.NewReader(abijson))
	panicErr(err)
	calldata, err := aabbii.Pack(callfunc, args...)
	panicErr(err)

	btxbs, err := tools.ToBytes(&types.TxEvmCommon{
		To:     common.HexToAddress(contract).Bytes(),
		Amount: new(big.Int),
		Load:   calldata,
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

func readContract(client *cl.ClientJSONRPC, privkey, contract, abijson, callfunc string, args []interface{}, nonce uint64) error {
	sk, err := crypto.HexToECDSA(privkey)
	panicErr(err)

	aabbii, err := abi.JSON(strings.NewReader(abijson))
	panicErr(err)
	data, err := aabbii.Pack(callfunc, args...)
	panicErr(err)

	// nonce := uint64(time.Now().UnixNano())
	to := common.HexToAddress(contract)
	privkey = ac.SanitizeHex(privkey)

	tx := ethtypes.NewTransaction(nonce, crypto.PubkeyToAddress(sk.PublicKey), to, big.NewInt(0), gasLimit, big.NewInt(0), data)
	b, err := rlp.EncodeToBytes(tx)
	panicErr(err)

	query := append([]byte{types.QueryTypeContract}, b...)
	res := new(agtypes.ResultQuery)
	if client == nil {
		client = cl.NewClientJSONRPC(logger, rpcTarget)
	}
	_, err = client.Call("query", []interface{}{query}, res)
	panicErr(err)

	fmt.Println("query result:", common.Bytes2Hex(res.Result.Data))
	parseResult, _ := unpackResult(callfunc, aabbii, string(res.Result.Data))
	fmt.Println("parse result:", reflect.TypeOf(parseResult), parseResult)

	return nil
}

func existContract(client *cl.ClientJSONRPC, privkey, contract, bytecode string) bool {
	sk, err := crypto.HexToECDSA(privkey)
	panicErr(err)

	if strings.Contains(bytecode, "f300") {
		bytecode = strings.Split(bytecode, "f300")[1]
	}

	data := common.Hex2Bytes(bytecode)
	privkey = ac.SanitizeHex(privkey)
	to := common.HexToAddress(ac.SanitizeHex(contract))

	tx := ethtypes.NewTransaction(0, crypto.PubkeyToAddress(sk.PublicKey), to, big.NewInt(0), gasLimit, big.NewInt(0), crypto.Keccak256(data))
	b, err := rlp.EncodeToBytes(tx)
	panicErr(err)

	query := append([]byte{types.QueryTypeContractExistance}, b...)
	res := new(agtypes.ResultQuery)
	if client == nil {
		client = cl.NewClientJSONRPC(logger, rpcTarget)
	}
	_, err = client.Call("query", []interface{}{query}, res)
	panicErr(err)

	hex := common.Bytes2Hex(res.Result.Data)
	if hex == "01" {
		return true
	}
	return false
}
