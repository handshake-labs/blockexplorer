package actions

import (
	"errors"
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/node"
	// "github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
	// "log"
	"os"
)

type GetMempoolTxsParams struct {
	Page int16 `json:"page"`
}

type GetMempoolTxsResult struct {
	Txs []MempoolTx `json:"txs""`
}

func GetMempoolTxs(ctx *Context, params *GetMempoolTxsParams) (*GetMempoolTxsResult, error) {
	nc := node.NewClient(os.Getenv("NODE_API_ORIGIN"), os.Getenv("NODE_API_KEY"))
	mempoolTxs, err := nc.GetMempool(ctx)
	var txs []MempoolTx
	var result GetMempoolTxsResult
	for _, mempoolTx := range *mempoolTxs {
		var tx MempoolTx
		copier.Copy(&tx, &mempoolTx)
		var nullArray []TxOutput //i need to free this field, as it's incorrectly populateds by copier and will be parsed below
		tx.TxOutputs = nullArray
		for _, txOutput := range mempoolTx.TxOutputs {
			txOut := TxOutput{}
			txOut.CovenantAction = db.CovenantAction(txOutput.Covenant.CovenantAction)
			copier.Copy(&txOut, &txOutput)
			covenantItems := txOutput.Covenant.CovenantItems
			if txOut.CovenantAction != "NONE" {
				name, _ := ctx.db.GetNameByNameHash(ctx, covenantItems[0])
				txOut.Name = name
			}
			switch txOut.CovenantAction { //this code is duplicate of what is done in db insertion, need to refactor it
			case "NONE":
			case "CLAIM":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantName = covenantItems[2]
			case "OPEN":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantName = covenantItems[2]
			case "BID":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantName = covenantItems[2]
				txOut.CovenantBidHash = covenantItems[3]
			case "REVEAL":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantNonce = covenantItems[2]
			case "REDEEM":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
			case "REGISTER":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantRecordData = covenantItems[2]
				txOut.CovenantBlockHash = covenantItems[3]
			case "UPDATE":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantRecordData = covenantItems[2]
			case "RENEW":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantBlockHash = covenantItems[2]
			case "TRANSFER":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantVersion = covenantItems[2]
				txOut.CovenantAddress = covenantItems[3]
			case "FINALIZE":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
				txOut.CovenantName = covenantItems[2]
				txOut.CovenantClaimHeight = covenantItems[4]
				txOut.CovenantRenewalCount = covenantItems[5]
				txOut.CovenantBlockHash = covenantItems[6]
			case "REVOKE":
				txOut.CovenantNameHash = covenantItems[0]
				txOut.CovenantHeight = covenantItems[1]
			default:
				return nil, errors.New("Unknown covenant action")
			}
			tx.TxOutputs = append(tx.TxOutputs, txOut)
		}
		txs = append(txs, tx)
		result.Txs = txs
	}
	return &result, err
}
