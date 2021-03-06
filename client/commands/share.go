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
	"encoding/hex"
	"fmt"
	"math/big"

	agtypes "github.com/annchain/annchain/angine/types"
	gcommon "github.com/annchain/annchain/module/lib/go-common"
	// "github.com/annchain/anth/common"
	"github.com/annchain/annchain/tools"
	"github.com/annchain/annchain/types"
	"github.com/annchain/anth/crypto"
	"gopkg.in/urfave/cli.v1"

	"github.com/annchain/annchain/client/commons"
	gcrypto "github.com/annchain/annchain/module/lib/go-crypto"
	cl "github.com/annchain/annchain/module/lib/go-rpc/client"
)

var (
	ShareCommands = cli.Command{
		Name:     "share",
		Usage:    "operations for share transaction",
		Category: "Share",
		Subcommands: []cli.Command{
			{
				Name:   "send",
				Usage:  "send a mount of share",
				Action: sendShare,
				Flags: []cli.Flag{
					anntoolFlags.nonce,
					anntoolFlags.to,
					anntoolFlags.value,
					cli.StringFlag{
						Name:  "nodeprivkey",
						Usage: "node account privkey",
					},
					cli.StringFlag{
						Name:  "evmprivkey",
						Usage: "evm account privkey",
					},
				},
			},
			{
				Name:   "guarantee",
				Usage:  "use share guarantee to participate election",
				Action: shareGuarantee,
				Flags: []cli.Flag{
					anntoolFlags.nonce,
					anntoolFlags.value,
					cli.StringFlag{
						Name:  "nodeprivkey",
						Usage: "node account privkey",
					},
					cli.StringFlag{
						Name:  "evmprivkey",
						Usage: "evm account privkey",
					},
				},
			},
			{
				Name:   "redeem",
				Usage:  "redeem share to exit election",
				Action: shareRedeem,
				Flags: []cli.Flag{
					anntoolFlags.nonce,
					anntoolFlags.value,
					cli.StringFlag{
						Name:  "nodeprivkey",
						Usage: "node account privkey",
					},
					cli.StringFlag{
						Name:  "evmprivkey",
						Usage: "evm account privkey",
					},
				},
			},
		},
	}
)

func sendShare(ctx *cli.Context) error {
	if !ctx.IsSet("to") {
		return cli.NewExitError("error: receiver account cannot be nil", 127)
	}

	//get node privkey
	nodepb, err := hex.DecodeString(gcommon.SanitizeHex(ctx.String("nodeprivkey")))
	if err != nil {
		return err
	}
	nodeprivkey := gcrypto.PrivKeyEd25519{}
	copy(nodeprivkey[:], nodepb)
	nodefrom := nodeprivkey.PubKey().(*gcrypto.PubKeyEd25519)

	tobs, err := agtypes.StringTo32byte(gcommon.SanitizeHex(ctx.String("to")))
	if err != nil {
		return err
	}
	to := gcrypto.PubKeyEd25519{}
	copy(to[:], tobs[:])

	bodyTx := &types.TxShareTransfer{
		ShareSrc: nodefrom[:],
		ShareDst: to[:],
		Amount:   big.NewInt(ctx.Int64("value")),
	}

	if bodyTx.ShareSig, err = tools.SignED25519(bodyTx, nodeprivkey[:]); err != nil {
		return err
	}
	bodybs, err := tools.ToBytes(bodyTx)
	if err != nil {
		return err
	}
	//construct blockTx
	skbs := ctx.String("evmprivkey")
	evmprivkey, err := crypto.HexToECDSA(skbs)
	if err != nil {
		panic(err)
	}
	nonce := ctx.Uint64("nonce")

	from := crypto.PubkeyToAddress(evmprivkey.PublicKey)
	fmt.Printf("%x\n", from)
	tx := types.NewBlockTx(big.NewInt(90000), big.NewInt(2), nonce, from[:], bodybs)

	if tx.Signature, err = tools.SignSecp256k1(tx, crypto.FromECDSA(evmprivkey)); err != nil {
		return cli.NewExitError(err.Error(), 127)
	}

	b, err := tools.ToBytes(tx)
	if err != nil {
		return cli.NewExitError(err.Error(), 127)
	}

	tmResult := new(agtypes.ResultBroadcastTxCommit)
	clientJSON := cl.NewClientJSONRPC(logger, commons.QueryServer)
	_, err = clientJSON.Call("broadcast_tx_commit", []interface{}{agtypes.WrapTx(types.TxTagAppEcoShareTransfer, b)}, tmResult)
	if err != nil {
		panic(err)
	}
	//res := (*tmResult).(*types.ResultBroadcastTxCommit)

	fmt.Printf("tx result: %x\n", tools.Hash(tx))

	return nil
}

func shareGuarantee(ctx *cli.Context) error {
	tx, b, err := constructEcoTx(ctx)

	tmResult := new(agtypes.ResultBroadcastTxCommit)
	clientJSON := cl.NewClientJSONRPC(logger, commons.QueryServer)
	_, err = clientJSON.Call("broadcast_tx_commit", []interface{}{agtypes.WrapTx(types.TxTagAppEcoGuarantee, b)}, tmResult)
	if err != nil {
		panic(err)
	}
	//res := (*tmResult).(*types.ResultBroadcastTxCommit)

	fmt.Printf("tx result: %x\n", tools.Hash(tx))
	return nil
}

func shareRedeem(ctx *cli.Context) error {
	tx, b, err := constructEcoTx(ctx)
	if err != nil {
		return err
	}

	tmResult := new(agtypes.ResultBroadcastTxCommit)
	clientJSON := cl.NewClientJSONRPC(logger, commons.QueryServer)
	_, err = clientJSON.Call("broadcast_tx_commit", []interface{}{agtypes.WrapTx(types.TxTagAppEcoRedeem, b)}, tmResult)
	if err != nil {
		panic(err)
	}
	//res := (*tmResult).(*types.ResultBroadcastTxCommit)

	fmt.Printf("tx result: %x\n", tools.Hash(tx))
	return nil
}

func constructEcoTx(ctx *cli.Context) (*types.BlockTx, []byte, error) {
	//get node privkey
	nodepb, err := hex.DecodeString(gcommon.SanitizeHex(ctx.String("nodeprivkey")))
	if err != nil {
		return nil, nil, err
	}
	nodeprivkey := gcrypto.PrivKeyEd25519{}
	copy(nodeprivkey[:], nodepb)
	nodefrom := nodeprivkey.PubKey().(*gcrypto.PubKeyEd25519)

	bodyTx := &types.TxShareEco{
		Source: nodefrom[:],
		Amount: big.NewInt(ctx.Int64("value")),
	}
	if bodyTx.Signature, err = tools.SignED25519(bodyTx, nodeprivkey[:]); err != nil {
		return nil, nil, err
	}
	bodybs, err := tools.ToBytes(bodyTx)
	if err != nil {
		return nil, nil, err
	}
	//construct blockTx
	skbs := ctx.String("evmprivkey")
	evmprivkey, err := crypto.HexToECDSA(skbs)
	if err != nil {
		panic(err)
	}
	nonce := ctx.Uint64("nonce")

	from := crypto.PubkeyToAddress(evmprivkey.PublicKey)
	fmt.Printf("%x\n", from)
	tx := types.NewBlockTx(big.NewInt(90000), big.NewInt(2), nonce, from[:], bodybs)

	if tx.Signature, err = tools.SignSecp256k1(tx, crypto.FromECDSA(evmprivkey)); err != nil {
		return nil, nil, cli.NewExitError(err.Error(), 127)
	}

	b, err := tools.ToBytes(tx)
	if err != nil {
		return nil, nil, cli.NewExitError(err.Error(), 127)
	}
	return tx, b, nil
}
