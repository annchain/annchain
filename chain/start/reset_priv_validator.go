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


package start

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/annchain/annchain/angine/types"
)

// NOTE: this is totally unsafe.
// it's only suitable for testnets.
func Reset_all(logger *zap.Logger, conf *viper.Viper) {
	Reset_priv_validator(logger, conf)
	os.RemoveAll(conf.GetString("db_dir"))
	os.RemoveAll(conf.GetString("cs_wal_dir"))
}

// NOTE: this is totally unsafe.
// it's only suitable for testnets.
func Reset_priv_validator(logger *zap.Logger, conf *viper.Viper) {
	// Get PrivValidator
	var privValidator *types.PrivValidator
	privValidatorFile := conf.GetString("priv_validator_file")
	if _, err := os.Stat(privValidatorFile); err == nil {
		privValidator = types.LoadPrivValidator(logger, privValidatorFile)
		privValidator.Reset()
		fmt.Println("Reset PrivValidator", "file", privValidatorFile)
	} else {
		privValidator = types.GenPrivValidator(logger)
		privValidator.SetFile(privValidatorFile)
		privValidator.Save()
		fmt.Println("Generated PrivValidator", "file", privValidatorFile)
	}
}
