package main

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	// "log"

	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/node"
	"github.com/handshake-labs/blockexplorer/pkg/types"

	"github.com/jinzhu/copier"
)

func deleteMempool(pg *sql.DB, nc *node.Client) error {
	q := db.New(pg)
	err := q.DeleteMempool(context.Background())
	return err
}

func syncMempool(pg *sql.DB, nc *node.Client) error {
	deleteMempool(pg, nc)
	q := db.New(pg)
	txs, err := nc.GetMempool(context.Background())
	for _, transaction := range *txs {
		tx := node.Transaction{}
		copier.Copy(&tx, &transaction)
		if err := insertTransaction(q, &tx, nil, nil); err != nil {
			return err
		}
	}
	return err
}

func syncBlocks(pg *sql.DB, nc *node.Client) error {
	height, hash, err := getSyncedHead(pg, nc)
	q := db.New(pg)
	if err != nil {
		return err
	}
	maxHeight, err := nc.GetBlocksHeight(context.Background())
	if err != nil {
		return err
	}
	for height < maxHeight {
		height += 1
		block, err := nc.GetBlockByHeight(context.Background(), height)
		if err != nil {
			return err
		}
		if !bytes.Equal(hash, block.PrevBlockHash) {
			break
		}
		err = syncBlock(pg, block)
		if err != nil {
			return err
		}
		hash = block.Hash
	}
	q.RefreshViews(context.Background())
	return nil
}

func getSyncedHead(pg *sql.DB, nc *node.Client) (int32, types.Bytes, error) {
	q := db.New(pg)
	height, err := q.GetBlocksMaxHeight(context.Background())
	if err != nil {
		return -1, nil, err
	}
	for height >= 0 {
		dbHash, err := q.GetBlockHashByHeight(context.Background(), height)
		if err != nil {
			return -1, nil, err
		}
		nodeHash, err := nc.GetBlockHashByHeight(context.Background(), height)
		if err != nil {
			return -1, nil, err
		}
		if bytes.Equal(dbHash, nodeHash) {
			if err := q.DeleteBlocksAfterHeight(context.Background(), height); err != nil {
				return -1, nil, err
			}
			return height, dbHash, nil
		}
		height -= 1
	}
	return -1, nil, nil
}

func insertTransaction(q *db.Queries, transaction *node.Transaction, blockHash *types.Bytes, index *int32) error {
	transactionParams := db.InsertTransactionParams{}
	copier.Copy(&transactionParams, &transaction)
	var err error
	transactionParams.BlockHash = blockHash
	var nullableIndex sql.NullInt32
	if index == nil {
		nullableIndex.Valid = false
	} else {
		nullableIndex.Valid = true
		nullableIndex.Int32 = *index
	}
	transactionParams.Index = nullableIndex
	if err = q.InsertTransaction(context.Background(), transactionParams); err != nil {
		return err
	}
	for index, txInput := range transaction.TxInputs {
		txInputParams := db.InsertTxInputParams{}
		txInputParams.Txid = transactionParams.Txid
		txInputParams.Index = int64(index)
		copier.Copy(&txInputParams, &txInput)
		if err := q.InsertTxInput(context.Background(), txInputParams); err != nil {
			return err
		}
	}
	for index, txOutput := range transaction.TxOutputs {
		txOutputParams := db.InsertTxOutputParams{}
		txOutputParams.Txid = transactionParams.Txid
		txOutputParams.CovenantAction = db.CovenantAction(txOutput.Covenant.CovenantAction)
		copier.Copy(&txOutputParams, &txOutput)
		txOutputParams.Index = int32(index)
		covenantItems := txOutput.Covenant.CovenantItems
		switch txOutputParams.CovenantAction {
		case "NONE":
		case "CLAIM":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantName = &covenantItems[2]
		case "OPEN":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantName = &covenantItems[2]
		case "BID":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantName = &covenantItems[2]
			txOutputParams.CovenantBidHash = &covenantItems[3]
		case "REVEAL":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantNonce = &covenantItems[2]
		case "REDEEM":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
		case "REGISTER":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantRecordData = &covenantItems[2]
			txOutputParams.CovenantBlockHash = &covenantItems[3]
		case "UPDATE":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantRecordData = &covenantItems[2]
		case "RENEW":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantBlockHash = &covenantItems[2]
		case "TRANSFER":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantVersion = &covenantItems[2]
			txOutputParams.CovenantAddress = &covenantItems[3]
		case "FINALIZE":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
			txOutputParams.CovenantName = &covenantItems[2]
			txOutputParams.CovenantClaimHeight = &covenantItems[4]
			txOutputParams.CovenantRenewalCount = &covenantItems[5]
			txOutputParams.CovenantBlockHash = &covenantItems[6]
		case "REVOKE":
			txOutputParams.CovenantNameHash = &covenantItems[0]
			txOutputParams.CovenantHeight = &covenantItems[1]
		default:
			return errors.New("Unknown covenant action")
		}
		if err := q.InsertTxOutput(context.Background(), txOutputParams); err != nil {
			return err
		}
	}
	return nil
}

func syncBlock(pg *sql.DB, block *node.Block) error {
	tx, err := pg.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q := db.New(tx)
	blockParams := db.InsertBlockParams{}
	copier.Copy(&blockParams, &block)
	if err := q.InsertBlock(context.Background(), blockParams); err != nil {
		return err
	}
	for tx_index, transaction := range block.Transactions {
		ind := int32(tx_index)
		if err := insertTransaction(q, &transaction, &blockParams.Hash, &ind); err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
