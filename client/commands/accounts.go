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

	"github.com/annchain/annchain/eth/crypto"
	gcommon "github.com/annchain/annchain/module/lib/go-common"
	agcrypto "github.com/annchain/annchain/module/lib/go-crypto"
	"gopkg.in/urfave/cli.v1"
)

var (
	//AccountCommands defines a more git-like subcommand system
	AccountCommands = cli.Command{
		Name:     "account",
		Usage:    "operations for account",
		Category: "Account",
		Subcommands: []cli.Command{
			{
				Name:     "gen",
				Action:   generatePrivPubAddr,
				Usage:    "generate new private-pub key pair",
				Category: "Account",
			},
			{
				Name:     "cal",
				Action:   calculatePrivPubAddr,
				Usage:    "calculate public key and address from private key",
				Category: "Account",
				Flags: []cli.Flag{
					anntoolFlags.privkey,
				},
			},
			{
				Name:     "geneth",
				Action:   generateSecpPrivPubAddr,
				Usage:    "generate new private-pub key pair",
				Category: "Account",
			},
		},
	}
)

func generatePrivPubAddr(ctx *cli.Context) error {
	sk := agcrypto.GenPrivKeyEd25519()
	pk := sk.PubKey().(*agcrypto.PubKeyEd25519)

	fmt.Printf("privkey: %X\n", sk[:])
	fmt.Printf("pubkey: %X\n", pk[:])

	return nil
}

func calculatePrivPubAddr(ctx *cli.Context) error {
	if !ctx.IsSet("privkey") {
		return cli.NewExitError("private key is required", -1)
	}

	skBs, err := hex.DecodeString(gcommon.SanitizeHex(ctx.String("privkey")))
	if err != nil {
		return cli.NewExitError(err.Error(), -1)
	}

	var sk agcrypto.PrivKeyEd25519
	copy(sk[:], skBs)

	pk := sk.PubKey().(*agcrypto.PubKeyEd25519)
	addr := pk.Address()

	fmt.Printf("pubkey : %X\n", pk[:])
	fmt.Printf("address: %X\n", addr)

	return nil
}

func generateSecpPrivPubAddr(ctx *cli.Context) error {
	sk, _ := crypto.GenerateKey()
	pk := crypto.FromECDSAPub(&sk.PublicKey)
	addr := crypto.PubkeyToAddress(sk.PublicKey)

	fmt.Printf("privkey: %X\n", crypto.FromECDSA(sk))
	fmt.Printf("pubkey: %X\n", pk[:])
	fmt.Printf("address : %X\n", addr[:])

	return nil
}
