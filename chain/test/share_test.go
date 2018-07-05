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

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/annchain/annchain/angine/types"
)

/*
  请将sender address 做为init_token，初始化给予一定balance
*/
var (
	shareReceiver = "0999D61A459023E2A04C3D560708EF4A5C6A5D2A8DCACF61EA3ADE6A8293A5EC"

	nodeprivkey = ""
	nodepub     = ""
)

func init() {
	//get priv_json frile
	privfile := path.Join(runtimePath, "priv_validator.json")
	fmt.Println("privfile path :", privfile)
	if _, err := os.Stat(privfile); err != nil {
		fmt.Println("cannot get priv_validator.json file")
		os.Exit(-1)
	}
	prvJsonBytes, err := ioutil.ReadFile(privfile)
	if err != nil {
		fmt.Println("err :", err)
		os.Exit(-1)
	}
	var privf types.PrivValidatorJSON
	err = json.Unmarshal(prvJsonBytes, &privf)
	if err != nil {
		fmt.Println("error :", err)
		os.Exit(-1)
	}
	nodeprivkey = privf.PrivKey.KeyString()
	nodepub = privf.PubKey.KeyString()
}

//test share transaction
func TestShareTransfer(t *testing.T) {
	msg := make(chan bool)
	nonce, err := getNonce(senderaddress)
	if err != nil {
		t.Error(err)
	}
	go func() {
		args := []string{"share", "send", "--nodeprivkey", nodeprivkey, "--evmprivkey", senderpriv, "--to", shareReceiver, "--value", "888", "--nonce", strconv.FormatUint(nonce, 10)}
		_, err := exec.Command(anntoolPath, args...).Output()
		if err != nil {
			t.Error(err)
		}
		close(msg)
	}()
	<-msg
	time.Sleep(time.Second * 1)
	args := []string{"query", "share", "--account_pubkey", shareReceiver}
	outs, err := exec.Command(anntoolPath, args...).Output()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(shareReceiver, " balance :", string(outs))
}

//test share guaranty
func TestGuaranty(t *testing.T) {
	msg := make(chan bool)
	nonce, err := getNonce(senderaddress)
	if err != nil {
		t.Error(err)
	}
	arg := []string{"query", "share", "--account_pubkey", nodepub}
	out1, err := exec.Command(anntoolPath, arg...).Output()
	if err != nil {
		t.Error(err)
	}
	go func() {
		args := []string{"share", "guarantee", "--nodeprivkey", nodeprivkey, "--evmprivkey", senderpriv, "--value", "123", "--nonce", strconv.FormatUint(nonce, 10)}
		_, err := exec.Command(anntoolPath, args...).Output()
		if err != nil {
			t.Error(err)
		}
		close(msg)
	}()
	<-msg
	time.Sleep(time.Second * 1)
	out2, err := exec.Command(anntoolPath, arg...).Output()
	if err != nil {
		t.Error(err)
	}
	if bytes.Equal(out1, out2) {
		t.Error("guarantee failed")
	}
}

//test share redeem
func TestRedeem(t *testing.T) {
	msg := make(chan bool)
	nonce, err := getNonce(senderaddress)
	if err != nil {
		t.Error(err)
	}
	arg := []string{"query", "share", "--account_pubkey", nodepub}
	out1, err := exec.Command(anntoolPath, arg...).Output()
	if err != nil {
		t.Error(err)
	}
	go func() {
		args := []string{"share", "redeem", "--nodeprivkey", nodeprivkey, "--evmprivkey", senderpriv, "--value", "123", "--nonce", strconv.FormatUint(nonce, 10)}
		_, err := exec.Command(anntoolPath, args...).Output()
		if err != nil {
			t.Error(err)
		}
		close(msg)
	}()
	<-msg
	time.Sleep(time.Second * 1)
	out2, err := exec.Command(anntoolPath, arg...).Output()
	if err != nil {
		t.Error(err)
	}
	if bytes.Equal(out1, out2) {
		t.Error("guarantee failed")
	}
}

func TestClean(t *testing.T) {
	node := <-nodeChan
	node.Process.Kill()
	exec.Command("rm", []string{"./node.*", "./client.*"}...)
}
