/*
Copyright 2018 Idealnaya rabota LLC
Licensed under Multy.io license.
See LICENSE for details
*/
package store

import (
	"time"
)

const (
	TxStatusAppearedInMempoolIncoming = 1
	TxStatusAppearedInBlockIncoming   = 2

	TxStatusAppearedInMempoolOutcoming = 3
	TxStatusAppearedInBlockOutcoming   = 4

	TxStatusInBlockConfirmedIncoming  = 5
	TxStatusInBlockConfirmedOutcoming = 6

	// TxStatusInBlockConfirmed = 5

	// TxStatusRejectedFromBlock = -1
)

// User represents a single app user
type User struct {
	UserID     string      `bson:"userID"`     // User uqnique identifier
	Devices    []Device    `bson:"devices"`    // All user devices
	Wallets    []Wallet    `bson:"wallets"`    // All user addresses in all chains
	WalletsETH []WalletETH `bson:"walletsEth"` // All user addresses in all chains
}

type BTCTransaction struct {
	Hash    string                `json:"hash"`
	Txid    string                `json:"txid"`
	Time    time.Time             `json:"time"`
	Outputs map[string]*BtcOutput `json:"outputs"` // addresses to outputs, key = address
}

type BtcOutput struct {
	Address     string  `json:"address"`
	Amount      float64 `json:"amount"`
	TxIndex     uint32  `json:"txIndex"`
	TxOutScript string  `json:"txOutScript"`
}

type TxInfo struct {
	Type    string  `json:"type"`
	TxHash  string  `json:"txhash"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

// Device represents a single users device.
type Device struct {
	DeviceID       string `bson:"deviceID"`       // Device uqnique identifier (MAC address of device)
	PushToken      string `bson:"pushToken"`      // Firebase
	JWT            string `bson:"JWT"`            // Device JSON Web Token
	LastActionTime int64  `bson:"lastActionTime"` // Last action time from current device
	LastActionIP   string `bson:"lastActionIP"`   // IP from last session
	DeviceType     int    `bson:"deviceType"`     // 1 - IOS, 2 - Android
}

const (
	WalletStatusOK      = "ok"
	WalletStatusDeleted = "deleted"
)

// Wallet Specifies a concrete wallet of user.
type Wallet struct {
	// Currency of wallet.
	CurrencyID int `bson:"currencyID"`
	// Currency of wallet.
	NetworkID int `bson:"networkID"`

	//wallet identifier
	WalletIndex int `bson:"walletIndex"`

	//wallet identifier
	WalletName string `bson:"walletName"`

	LastActionTime int64 `bson:"lastActionTime"`

	DateOfCreation int64 `bson:"dateOfCreation"`

	// All addresses assigned to this wallet.
	Adresses []Address `bson:"addresses"`

	Status string `bson:"status"`
}

type RatesRecord struct {
	Category int    `json:"category" bson:"category"`
	TxHash   string `json:"txHash" bson:"txHash"`
}

type Address struct {
	AddressIndex   int    `json:"addressIndex" bson:"addressIndex"`
	Address        string `json:"address" bson:"address"`
	LastActionTime int64  `json:"lastActionTime" bson:"lastActionTime"`
}
type WalletsSelect struct {
	Wallets []struct {
		Addresses []struct {
			AddressIndex int    `bson:"addressIndex"`
			Address      string `bson:"address"`
		} `bson:"addresses"`
		WalletIndex int `bson:"walletIndex"`
	} `bson:"wallets"`
}

type WalletForTx struct {
	UserId      string           `json:"userid"`
	WalletIndex int              `json:"walletindex"`
	Address     AddressForWallet `json:"address"`
}

type AddressForWallet struct {
	AddressIndex    int    `json:"addressindex"`
	AddressOutIndex int    `json:"addresoutindex"`
	Address         string `json:"address"`
	Amount          int64  `json:"amount"`
}

// the way how user transations store in db
type MultyTX struct {
	UserId            string                `json:"userid"`
	TxID              string                `json:"txid"`
	TxHash            string                `json:"txhash"`
	TxOutScript       string                `json:"txoutscript"`
	TxAddress         []string              `json:"addresses"` //this is major addresses of the transaction (if send - inputs addresses of our user, if get - outputs addresses of our user)
	TxStatus          int                   `json:"txstatus"`
	TxOutAmount       int64                 `json:"txoutamount"`
	BlockTime         int64                 `json:"blocktime"`
	BlockHeight       int64                 `json:"blockheight"`
	Confirmations     int                   `json:"confirmations"`
	TxFee             int64                 `json:"txfee"`
	MempoolTime       int64                 `json:"mempooltime"`
	StockExchangeRate []ExchangeRatesRecord `json:"stockexchangerate"`
	TxInputs          []AddresAmount        `json:"txinputs"`
	TxOutputs         []AddresAmount        `json:"txoutputs"`
	WalletsInput      []WalletForTx         `json:"walletsinput"`  //here we storing all wallets and addresses that took part in Inputs of the transaction
	WalletsOutput     []WalletForTx         `json:"walletsoutput"` //here we storing all wallets and addresses that took part in Outputs of the transaction
}

type AddresAmount struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

type TxRecord struct {
	UserID       string    `json:"userid"`
	Transactions []MultyTX `json:"transactions"`
}

// ExchangeRatesRecord presents record with exchanges from rate stock
// with additional information, such as date and exchange stock
type ExchangeRatesRecord struct {
	Exchanges     ExchangeRates `json:"exchanges"`
	Timestamp     int64         `json:"timestamp"`
	StockExchange string        `json:"stock_exchange"`
}

// ExchangeRates stores exchange rates
type ExchangeRates struct {
	EURtoBTC float64 `json:"eur_btc"`
	USDtoBTC float64 `json:"usd_btc"`
	ETHtoBTC float64 `json:"eth_btc"`

	ETHtoUSD float64 `json:"eth_usd"`
	ETHtoEUR float64 `json:"eth_eur"`

	BTCtoUSD float64 `json:"btc_usd"`
}

type RatesAPIBitstamp struct {
	Date  string `json:"date"`
	Price string `json:"price"`
}
type SpendableOutputs struct {
	TxID              string                `json:"txid"`
	TxOutID           int                   `json:"txoutid"`
	TxOutAmount       int64                 `json:"txoutamount"`
	TxOutScript       string                `json:"txoutscript"`
	Address           string                `json:"address"`
	UserID            string                `json:"userid"`
	WalletIndex       int                   `json:"walletindex"`
	AddressIndex      int                   `json:"addressindex"`
	TxStatus          int                   `json:"txstatus"`
	StockExchangeRate []ExchangeRatesRecord `json:"stockexchangerate"`
}

type WalletETH struct {
	// Currency of wallet.
	CurrencyID int `bson:"currencyID"`
	// Currency of wallet.
	NetworkID int `bson:"networkID"`

	//wallet identifier
	WalletIndex int `bson:"walletIndex"`

	//wallet identifier
	WalletName string `bson:"walletName"`

	LastActionTime int64 `bson:"lastActionTime"`

	DateOfCreation int64 `bson:"dateOfCreation"`

	// All addresses assigned to this wallet.
	Adresses []Address `bson:"addresses"`

	// Wallet status
	Status string `bson:"status"`

	// Balance of the eth wallet in wei
	Balance int64 `bson:"balance"`

	// Nonce of the wallet - index of the last transaction
	Nonce int64 `bson:"nonce"`
}

type TransactionETH struct {
	UserID            string                `json:"userid"`
	Hash              string                `json:"hash"`
	From              string                `json:"from"`
	To                string                `json:"to"`
	Amount            int64                 `json:"amount"`
	GasPrice          int64                 `json:"gasprice"`
	GasLimit          int64                 `json:"gaslimit"`
	Nonce             int                   `json:"nonce"`
	Status            int                   `json:"status"`
	BlockTime         int64                 `json:"blocktime"`
	PoolTime          int64                 `json:"tpooltime"`
	BlockHeight       int64                 `json:"blockheight"`
	StockExchangeRate []ExchangeRatesRecord `json:"stockexchangerate"`
}

type CoinType struct {
	СurrencyID int
	NetworkID  int
	GRPCUrl    string
}

type MempoolRecord struct {
	Category int    `json:"category"`
	HashTX   string `json:"hashTX"`
}

type DeleteSpendableOutput struct {
	UserID  string
	TxID    string
	Address string
}

type DonationInfo struct {
	FeatureCode     int
	DonationAddress string
}

type AddressExtended struct {
	UserID       string
	WalletIndex  int
	AddressIndex int
}

type ServerConfig struct {
	BranchName string `json:"branch"`
	CommitHash string `json:"commit"`
	Build      string `json:"build_time"`
	Tag        string `json:"tag"`
}
