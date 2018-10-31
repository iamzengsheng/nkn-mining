package transactionPayload

import (
	"NKNMining/network/chainDataTypes/rpcApiResponse"
)

type Register struct {
	Asset  rpcApiResponse.Asset `json:"asset"`
	Amount string               `json:"amount"`
}
