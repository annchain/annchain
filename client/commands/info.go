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

package commands

import (
	"encoding/json"
	"fmt"

	"github.com/annchain/annchain/angine/types"
	"github.com/annchain/annchain/client/commons"
	cl "github.com/annchain/annchain/module/lib/go-rpc/client"
	"gopkg.in/urfave/cli.v1"
)

var (
	InfoCommand = cli.Command{
		Name:  "info",
		Usage: "get ann info",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "last_block",
				Action: lastBlockInfo,
			},
			cli.Command{
				Name:   "num_unconfirmed_txs",
				Action: numUnconfirmedTxs,
			},
			cli.Command{
				Name:   "net",
				Action: netInfo,
			},
		},
	}
)

func lastBlockInfo(ctx *cli.Context) error {
	clientJSON := cl.NewClientJSONRPC(logger, commons.QueryServer)
	tmResult := new(types.RPCResult)
	_, err := clientJSON.Call("info", []interface{}{}, tmResult)
	if err != nil {
		return cli.NewExitError(err.Error(), 127)
	}
	res := (*tmResult).(*types.ResultInfo)
	var jsbytes []byte
	jsbytes, err = json.Marshal(res)
	if err != nil {
		return cli.NewExitError(err.Error(), 127)
	}
	fmt.Println(string(jsbytes))
	return nil
}

func numUnconfirmedTxs(ctx *cli.Context) error {
	clientJSON := cl.NewClientJSONRPC(logger, commons.QueryServer)
	tmResult := new(types.ResultUnconfirmedTxs)
	_, err := clientJSON.Call("num_unconfirmed_txs", []interface{}{}, tmResult)
	if err != nil {
		return cli.NewExitError(err.Error(), 127)
	}

	fmt.Println("num of unconfirmed txs: ", tmResult.N)
	return nil
}

func netInfo(ctx *cli.Context) error {
	clientJSON := cl.NewClientJSONRPC(logger, commons.QueryServer)
	res := new(types.ResultNetInfo)
	_, err := clientJSON.Call("net_info", []interface{}{}, res)
	if err != nil {
		panic(err)
	}
	fmt.Println("listening :", res.Listening)
	for _, l := range res.Listeners {
		fmt.Println("listener :", l)
	}
	for _, p := range res.Peers {
		fmt.Println("peer address :", p.RemoteAddr,
			" pub key :", p.PubKey,
			" send status :", p.ConnectionStatus.SendMonitor.Active,
			" recieve status :", p.ConnectionStatus.RecvMonitor.Active)
	}
	return nil
}
