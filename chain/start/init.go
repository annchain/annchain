package start

import (
	"github.com/spf13/viper"

	"github.com/annchain/annchain/angine"
)

func Initfiles(conf *viper.Viper) {
	angine.Initialize(&angine.Tunes{Conf: conf},"")
}