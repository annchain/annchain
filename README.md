# annchain

annchain 采用 dBFT 共识算法，兼容以太坊合约交易。

## 编译    

	go get github.com/annchain/annchain
	make ann


## 初始化节点

执行以下命令

	./ann init

初始化将会在 ~/.angine 目录下产生3个配置文件，分别是：

	config.toml           // 链运行所需要的参数
	genesis.json          // 创世块信息
	priv_validator.json   // 节点身份信息

[详细文档](../document/README.md)
