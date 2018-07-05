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



## 配置新的链网络

#### 1 配置链网络创世信息

配置新的链网络，需要配置链的创世信息 ( genesis.json )，主要有初始验证节点、初始token分配、初始share分配三项。

##### 1.1 配置初始验证节点

初始搭建一个新的链网络时，需要配置链初始节点信息。genesis.json 文件中有一个 validators 数组，表示链的初始状态下，有哪些验证节点。假设我们初始有4个验证节点，则需要将这4个节点的信息填写进validators数组。像下边这样：

	"validators": [
		{
			"pub_key": [
				1,
				"BEED23BC7F3427809FACB27A2A15AA4929FE30AE44C4523A70A261299F52AAE9"
			],
			"amount": 100,
			"is_ca": true
		},
		{
			"pub_key": [
				1,
				"4485349A302B4F08B08D0602B6E205B587127DECF1C326C0D1038612E8966E1A"
			],
			"amount": 100,
			"is_ca": true
		},
		...
	],

如果单节点运行，数组中配置自己就可以了。

##### 1.2 配置初始token分配

token是链网络中用于价值交互以及用作交易手续费的数字信息，在这里可以配置新的链网络中，token的初始分配情况，像下边这样：
		
	"init_token": [
		{"address":"B9A383673C09366D05A97323F682624DC11670CC","amount":"100000000000","extra":"hello11111"},
		{"address":"9F403EC1C16317683961BEA249FEFCE88CE619D1","amount":"100000000000","extra":"hello22222"}
    ],

##### 1.3 配置初始share分配

share是链网络中用于参与共识的份额信息，在这里可以配置新的链网络中，share的初始分配情况，像下边这样：
	
	"init_share": [
		{"pubkey":"9946622A1233E31BBE09CC439D8C217D98D78BC94A436C6E34CA997C8253C009","amount":"10000000","extra":"hello22222"},
		{"pubkey":"59923B75050F05BBBF789781965760038A4E4C048A050B8A9CE7ABF3D70D3F11","amount":"10000000","extra":"hello22222"},
	],

#### 2 配置区块链 ID

genesis.json 配置文件中 chain\_id 参数表示当前链的 ID，需要加入同一个链网络的节点，需保持 chain\_id 一致。

#### 3 配置连接 seeds

在 config.toml 配置文件中 seeds 参数表示节点接入链网络的入口，在这里我们统一填写其中一个节点的地址即可：

	seeds = "172.28.228.75:46656"

这里的端口默认为 46656，如有冲突，可以修改参数 p2p_laddr 

## 接入已有的链网络

接入已有的链网络，跟配置新的链网络相比，区别在于不需要步骤1，配置区块链ID为目标网络的区块链ID，配置链接seeds之后即可。

## 启动链

配置好之后，各个节点执行启动命令

	./ann run

也可以使其在后台启动

	nohup ./ann run &

启动之后，可以在 ~/.angine/angine-XXX（XXX是上边提到的 chain\_id）目录中看到链执行的日志，查看该日志可以看到链运行输出的信息

## 附录

#### 1、config.toml 配置项

	environment = "development"               // 日志级别，支持development和production
	p2p_laddr = "tcp://0.0.0.0:46656"         // P2P 监听端口
	rpc_laddr = "tcp://0.0.0.0:46657"         // 本地RPC命令监听端口
	log_path = ""                             // 日志路径
	seeds = "172.28.228.75:46656"             // 加入链网络的入口地址
	auth_by_ca = false                        // 加入链网络时是否使用CA认证
	enable_incentive = true                   // 是否启用激励
	db_backend = "leveldb"                    // 底层数据库，暂不可改
	moniker = "anonymous"                     // 暂不支持修改
	signbyca = ""                             // auth_by_ca=true 时有效，CA节点给当前节点的签名

#### 2、genesis.json 配置项

	{
		"genesis_time": "0001-01-01T00:00:00Z",        // 无需配置
		"chain_id": "ann",                             // 区块链ID
		"validators": null,                            // 初始验证节点数组
		"app_hash": "",                                // 无需配置
		"plugins": "specialop,suspect,querycache",     // 插件
		"init_token": null,                            // 初始token分配数组
		"init_share": null                             // 初始share分配数组
	}

#### 3、priv_validator.json 配置项

	{
		"pub_key": [                                  // 节点公钥
			1,                                        // 公钥加密算法，代表ED25519，暂不支持修改
			"5BE3622901BA3A93F81119C262EAB3B8321180FBD9CE49B00D05A1335F8BD14F"
		],
		"last_height": 0,                             // 共识状态，无需修改
		"last_round": 0,                              // 共识状态，无需修改
		"last_step": 0,                               // 共识状态，无需修改
		"last_signature": null,                       // 共识状态，无需修改
		"last_signbytes": "",                         // 共识状态，无需修改
		"priv_key": [                                 // 节点私钥
			1,
			"8A0F79C3434808CADF184000649E368C0C0A53206BC84066A6437BDE39E493595BE3622901BA3A93F81119C262EAB3B8321180FBD9CE49B00D05A1335F8BD14F"
		]
	}