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
	"math/rand"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	infoOnly      struct{}
	infoWithDebug struct{}
	aboveWarn     struct{}
)

func (l infoOnly) Enabled(lv zapcore.Level) bool {
	return lv == zapcore.InfoLevel
}
func (l infoWithDebug) Enabled(lv zapcore.Level) bool {
	return lv == zapcore.InfoLevel || lv == zapcore.DebugLevel
}
func (l aboveWarn) Enabled(lv zapcore.Level) bool {
	return lv >= zapcore.WarnLevel
}

func makeErrorFilter() zapcore.LevelEnabler {
	return aboveWarn{}
}

func init() {
	rand.Seed(time.Now().UnixNano())

	var encoderCfg zapcore.EncoderConfig
	encoderCfg = zap.NewDevelopmentEncoderConfig()
	coreInfo := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.NewMultiWriteSyncer(os.Stdout),
		infoWithDebug{},
	)
	coreError := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.NewMultiWriteSyncer(os.Stderr),
		makeErrorFilter(),
	)

	logger = zap.New(zapcore.NewTee(coreInfo, coreError))
}
