package staking

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 设置相关参数
const (
	// InfuraURL    = "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID" // 替换为你的Infura项目ID
	InfuraURL    = "https://sepolia-rollup.arbitrum.io/rpc"     // 替换为你的Infura项目ID
	ContractAddr = "0x669F929a3eb6009741C66d97F74D837C5e89622c" // 替换为合约地址
	// ContractAddr = "0x01571d0936d014521efab3FD2a9345D90B560624"                       // 替换为合约地址
	// PrivateKey = "9f381a543ab543a386941924c4734688d2367fb41acbe6fb629c74c95a4e6f61" // 替换为私钥
	PrivateKey = "795daf0d38010d9d1f919c8de96a44c80907837f08bbc0e07c45e202c5409478" // 替换为私钥

	// ABIFile = ".\\library\\contract\\staking.abi" // ABI文件路径
	ChainID = 421614
)

func linkClient() (client *ethclient.Client, auth *bind.TransactOpts, stakingContract *Staking) {
	// 1. 连接到以太坊客户端
	client, err := ethclient.Dial(InfuraURL)
	if err != nil {
		log.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer client.Close()
	// 2. 加载私钥
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Fatalf("私钥解析失败: %v", err)
	}

	// 3. 获取发送地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥转换失败")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 4. 获取交易参数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("获取 nonce 失败: %v", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("获取 gas 价格失败: %v", err)
	}
	chainID := big.NewInt(ChainID) // 替换为你的链 ID
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("创建交易认证失败: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // 发送的 ETH 数量，这里设置为 0
	auth.GasLimit = uint64(300000) // 自定义 Gas Limit
	auth.GasPrice = gasPrice

	// 5. 创建合约实例
	contractAddress := common.HexToAddress(ContractAddr)
	stakingContract, err = NewStaking(contractAddress, client)
	if err != nil {
		log.Fatalf("加载合约实例失败: %v", err)
	}
	return client, auth, stakingContract
}

func PausedStaking() (succ bool, err error) {
	client, auth, stakingContract := linkClient()
	// 6. 调用 `pause` 方法，传入 `私钥` 参数
	tx, err := stakingContract.Pause(auth) // 示例利率 1000
	if err != nil {
		log.Fatalf("调用合约方法失败: %v", err)
	}

	fmt.Printf("交易已发送, 交易哈希: %s\n", tx.Hash().Hex())

	// 7. 等待交易确认
	receipt, err := waitForTransaction(client, tx.Hash())
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}
	// 检查交易状态
	if receipt.Status == 1 {
		fmt.Println("交易成功，创建的空间ID在合约事件日志中。")
		return true, nil
	} else {
		fmt.Println("交易失败")
		return false, errors.New("transaction failed")
	}
}

// 等待交易确认函数
func waitForTransaction(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		time.Sleep(1 * time.Second)
	}
}

func StartSpace() (succ bool, err error) {
	client, auth, stakingContract := linkClient()
	// 6. 调用 `周期切换` 方法，传入 `私钥` 参数
	tx, err := stakingContract.StartSwitchPeriod(auth) // 示例利率 1000
	if err != nil {
		log.Fatalf("调用合约方法失败: %v", err)
	}

	fmt.Printf("交易已发送, 交易哈希: %s\n", tx.Hash().Hex())

	// 7. 等待交易确认
	receipt, err := waitForTransaction(client, tx.Hash())
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}
	// 检查交易状态
	if receipt.Status == 1 {
		fmt.Println("交易成功，创建的空间ID在合约事件日志中。")
		return true, nil
	} else {
		fmt.Println("交易失败")
		return false, errors.New("transaction failed")
	}

}

func ContinueSpace() (succ bool, err error) {
	client, auth, stakingContract := linkClient()
	// 6. 调用 `周期切换` 方法，传入 `私钥` 参数
	tx, err := stakingContract.ContinueSwitchPeriod(auth) // 示例利率 1000
	if err != nil {
		log.Fatalf("调用合约方法失败: %v", err)
	}

	fmt.Printf("交易已发送, 交易哈希: %s\n", tx.Hash().Hex())

	// 7. 等待交易确认
	receipt, err := waitForTransaction(client, tx.Hash())
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}
	// 检查交易状态
	if receipt.Status == 1 {
		fmt.Println("交易成功，创建的空间ID在合约事件日志中。")
		return true, nil
	} else {
		fmt.Println("交易失败")
		return false, errors.New("transaction failed")
	}

}

func AddNewWhiteAddress(address string) (succ bool, err error) {
	client, auth, stakingContract := linkClient()
	// 6. 调用 `周期切换` 方法，传入 `私钥` 参数
	// tx, err := stakingContract.AddToWhitelist(auth, common.HexToAddress("0x264eb13b9552Ecb0b20d3A6dc001fdaA16C13d2c"))
	tx, err := stakingContract.AddToWhitelist(auth, common.HexToAddress(address))
	if err != nil {
		log.Fatalf("调用合约方法失败: %v", err)
	}

	fmt.Printf("交易已发送, 交易哈希: %s\n", tx.Hash().Hex())

	// 7. 等待交易确认
	receipt, err := waitForTransaction(client, tx.Hash())
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}
	// 检查交易状态
	if receipt.Status == 1 {
		fmt.Println("交易成功，创建的空间ID在合约事件日志中。")
		return true, nil
	} else {
		fmt.Println("交易失败")
		return false, errors.New("transaction failed")
	}
}
