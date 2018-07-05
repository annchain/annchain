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


package log

import (
	"path"

	"go.uber.org/zap"

	cmn "github.com/annchain/annchain/module/lib/go-common"
)

func Initialize(env, output, errOutput string) *zap.Logger {
	var zapConf zap.Config
	var err error

	if env == "production" {
		zapConf = zap.NewProductionConfig()
	} else {
		zapConf = zap.NewDevelopmentConfig()
	}

	cmn.EnsureDir(path.Dir(output), 0775)
	cmn.EnsureDir(path.Dir(errOutput), 0775)

	zapConf.OutputPaths = []string{output}
	zapConf.ErrorOutputPaths = []string{errOutput}
	logger, err := zapConf.Build()
	if err != nil {
		panic(err.Error()) // which should never happen
	}

	logger.Debug("Starting zap! Have your fun!")

	return logger
}
