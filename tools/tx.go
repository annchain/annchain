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
	"bytes"
	"fmt"

	tmcrypto "github.com/annchain/annchain/module/lib/go-crypto"
	ethcrypto "github.com/annchain/anth/crypto"
	"github.com/annchain/anth/rlp"
)

func ToBytes(tx interface{}) ([]byte, error) {
	return rlp.EncodeToBytes(tx)
}

func FromBytes(bs []byte, tx interface{}) error {
	return rlp.DecodeBytes(bs, tx)
}

func Hash(o interface{}) []byte {
	bs, err := ToBytes(o)
	if err != nil {
		panic(err)
	}

	return ethcrypto.Sha256(bs)
}

type CanSign interface {
	SigObject() interface{}
}

func SigHash(sigObj CanSign) ([]byte, error) {
	bs, err := ToBytes(sigObj.SigObject())
	if err != nil {
		return nil, err
	}

	return ethcrypto.Sha256(bs), nil
}

func SignSecp256k1(tx CanSign, privkey []byte) ([]byte, error) {
	h, err := SigHash(tx)
	if err != nil {
		return nil, err
	}

	sk := ethcrypto.ToECDSA(privkey)
	sig, err := ethcrypto.Sign(h, sk)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func VerifySecp256k1(tx CanSign, sender, sig []byte) error {
	if len(sig) == 0 {
		return fmt.Errorf("empty signature")
	}

	h, err := SigHash(tx)
	if err != nil {
		return err
	}

	pub, err := ethcrypto.Ecrecover(h, sig)
	if err != nil {
		return err
	}

	if addr := ethcrypto.Keccak256(pub[1:])[12:]; !bytes.Equal(sender, addr) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

func SignED25519(tx CanSign, privkey []byte) ([]byte, error) {
	h, err := SigHash(tx)
	if err != nil {
		return nil, err
	}

	sk := tmcrypto.PrivKeyEd25519{}
	copy(sk[:], privkey)

	return sk.Sign(h).Bytes(), nil
}

func VeirfyED25519(tx CanSign, sender, sig []byte) error {
	if len(sig) == 0 {
		return fmt.Errorf("empty signature")
	}

	h, err := SigHash(tx)
	if err != nil {
		return err
	}

	pk := &tmcrypto.PubKeyEd25519{}
	copy(pk[:], sender)

	s := &tmcrypto.SignatureEd25519{}
	copy(s[:], sig)

	if !pk.VerifyBytes(h, s) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
