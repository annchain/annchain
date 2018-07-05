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
	"gopkg.in/urfave/cli.v1"
)

type AnntoolFlags struct {
	abif,
	callf,
	addr,
	pwd,
	payload,
	bytecode,
	privkey,
	nonce,
	abistr,
	callstr,
	value,
	fee,
	hash,
	accountPubkey,
	peerPubkey,
	validatorPubkey,
	power,
	isCA,
	rpc,
	chainid,
	to,
	codeHash cli.Flag
}

var anntoolFlags = AnntoolFlags{
	abif: cli.StringFlag{
		Name:  "abif",
		Usage: "abi definition file",
	},
	callf: cli.StringFlag{
		Name:  "callf",
		Usage: "params file defined in JSON",
	},
	addr: cli.StringFlag{
		Name: "address",
	},
	pwd: cli.StringFlag{
		Name: "passwd",
	},
	payload: cli.StringFlag{
		Name: "payload",
	},
	bytecode: cli.StringFlag{
		Name: "bytecode",
	},
	privkey: cli.StringFlag{
		Name: "privkey",
	},
	nonce: cli.Uint64Flag{
		Name: "nonce",
	},
	abistr: cli.StringFlag{
		Name: "abi",
	},
	callstr: cli.StringFlag{
		Name: "calljson",
	},
	to: cli.StringFlag{
		Name: "to",
	},
	value: cli.Int64Flag{
		Name: "value",
	},
	fee: cli.Int64Flag{
		Name: "fee",
	},
	hash: cli.StringFlag{
		Name: "hash",
	},
	accountPubkey: cli.StringFlag{
		Name: "account_pubkey",
	},
	peerPubkey: cli.StringFlag{
		Name: "peer_pubkey",
	},
	validatorPubkey: cli.StringFlag{
		Name: "validator_pubkey",
	},
	power: cli.IntFlag{
		Name: "power",
	},
	isCA: cli.BoolFlag{
		Name: "isCA",
	},
	rpc: cli.StringFlag{
		Name:  "rpc",
		Value: "tcp://0.0.0.0:46657",
	},
	chainid: cli.StringFlag{
		Name:  "chainid",
		Value: "",
	},
	codeHash: cli.StringFlag{
		Name:  "code_hash",
		Value: "",
	},
}
