package main

//"github.com/go-btcrpc"
import (
	btcrpc "github.com/go-btcrpc"
	"github.com/go_eth_explorer/app/db/mysql"
)

func getCurrentBlockHeightInDB() (int64, error) {

	var id int64

	if err := mysql.DBQueryRow(
		"select max(Height) from btcblockinfo ",
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil

}

func insertTxIn(vin *btcrpc.Vin, txId string, txHash string) (int64, error) {
	var id int64
	if ok := vin.IsCoinBase(); !ok {
		result, err := mysql.DBExec(
			"replace INTO txvin (Txid,txHash, pretxid,Vout, created_at, updated_at) values (?,?,?,?,now(),now())",
			txId,
			txHash,
			vin.Txid,
			vin.Vout,
		)
		if err != nil {
			panic(err)
		}

		id, err = result.LastInsertId()
		if err != nil {
			panic(err)
		}

		return id, nil
	}
	if ok := vin.IsCoinBase(); ok {
		result, err := mysql.DBExec(
			"replace INTO txvin (Txid,txHash, pretxid,Vout, created_at, updated_at) values (?,?,?,?,now(),now())",
			txId,
			txHash,
			"",
			-1,
		)
		if err != nil {
			panic(err)
		}

		id, err = result.LastInsertId()
		if err != nil {
			panic(err)
		}

		return id, nil
	}
	return 0, nil
}

func insertTxOut(vout *btcrpc.Vout, txId string, txHash string) (int, error) {

	for _, v := range vout.ScriptPubKey.Addresses {
		_, err := mysql.DBExec(
			"replace INTO txvout (Txid,txHash, Value, N, address, created_at, updated_at) values (?,?,?,?,?,now(),now())",
			txId,
			txHash,
			vout.Value,
			vout.N,
			v,
		)
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
	}
	return len(vout.ScriptPubKey.Addresses), nil
}

func insertTransaction(tx *btcrpc.Transaction) (int64, error) {
	var id int64
	result, err := mysql.DBExec(
		"replace INTO btcTransaction (Txid,txHash, Version, Size, Vsize,Locktime, created_at, updated_at) values (?,?,?,?,?,?,now(),now())",
		tx.Txid,
		tx.Hash,
		tx.Version,
		tx.Size,
		tx.Vsize,
		tx.Locktime,
	)
	if err != nil {
		panic(err)
	}

	id, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id, nil
}

func insertBlockinfo(blk *btcrpc.Block) (int64, error) {
	var id int64

	result, err := mysql.DBExec(
		"replace INTO btcblockinfo (Hash,Version, Time, Size,Nonce, Weight, Difficulty,Merkleroot,NextBlockhash,Confirmations,Height, created_at, updated_at) values (?,?,?,?,?,?,?,?,?,?,?,now(),now())",
		blk.Hash,
		blk.Version,
		blk.Time,
		blk.Size,
		blk.Nonce,
		blk.Weight,
		blk.Difficulty,
		blk.Merkleroot,
		blk.NextBlockhash,
		blk.Confirmations,
		blk.Height,
	)
	if err != nil {
		panic(err)
	}

	id, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id, nil
}
