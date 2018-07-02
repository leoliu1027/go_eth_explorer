package main

import (
	"fmt"
	"log"

	btcrpc "github.com/go-btcrpc"
)

type RPCClient struct {
	*btcrpc.RPCClient
}

var (
	client RPCClient
)

func connect_to_bitcoin_RPC(basicAuth *btcrpc.BasicAuth) *btcrpc.RPCClient {
	c := btcrpc.NewRPCClient("http://123.56.21.236:8332", basicAuth)
	return c
}

func (c *RPCClient) get_all_block_hashes() []string {
	var block_hash_list []string
	block_height, err := c.RPCClient.GetBlockCount()
	if err != nil {
		return nil
	}
	for j := 0; j <= int(block_height); j++ {
		block_hashes, _ := c.RPCClient.GetBlockHash(int32(j))
		block_hash_list = append(block_hash_list, block_hashes)
	}
	return block_hash_list
}

func (c *RPCClient) get_all_block_hashes_from_present_to_past() []string {
	var block_hash_list []string
	block_height, err := c.RPCClient.GetBlockCount()
	if err != nil {
		return nil
	}
	for j := 0; j <= int(block_height); j++ {
		block_hashes, _ := c.RPCClient.GetBlockHash(int32(j))
		block_hash_list = append(block_hash_list, block_hashes)
	}
	return block_hash_list
}

func (c *RPCClient) get_block_info(block_hash string) (*btcrpc.Block, error) {
	block, err := c.RPCClient.GetBlock(block_hash)
	if err != nil {
		log.Println(fmt.Sprintf("get block info error. block_hash %d", block_hash))
		return nil, nil
	}
	return block, err
}

func (c *RPCClient) update_block_info() {
	height_in_db, err := getCurrentBlockHeightInDB()
	// deleteing last two blocks and updating balance and degree table before synchronizing again
	if height_in_db >= 2 {
		for j := int(height_in_db) - 2 + 1; j <= int(height_in_db)+1; j++ {
			log.Println(fmt.Sprintf("delete pre block "))
		}
	}
	block_height, err := c.RPCClient.GetBlockCount()
	if err != nil {
		block_height = 0
	}
	if height_in_db+1 == int64(block_height) {
		log.Println(fmt.Sprintf("All blocks are up to date."))
	}

	for height := int64(529001); height <= int64(block_height)+1; height++ {
		block_hash, err := c.RPCClient.GetBlockHash(int32(height))
		if err != nil {
			log.Println(fmt.Sprintf("get block hash error. height %d", height))
			continue
		}
		block, err := c.get_block_info(block_hash)
		if err != nil {
			continue
		}
		insertBlockinfo(block)
		c.update_block_transaction_info(block)
	}
}

func (c *RPCClient) get_transaction_info(tx_hash []string) ([]*btcrpc.Transaction, error) {
	rawtx, err := c.RPCClient.GetRawTransactions(tx_hash)
	txs, err := c.RPCClient.DecodeRawTransactions(rawtx)
	return txs, err
}

func (c *RPCClient) update_block_transaction_info(block *btcrpc.Block) {
	tx_hashes := block.Txs
	tx_info, err := c.get_transaction_info(tx_hashes)
	if err != nil {
		return
	}
	for _, v := range tx_info {
		for _, vin := range v.Vin {
			insertTxIn(&vin, v.Txid, v.Hash)

		}
		//         process vout
		for _, vout := range v.Vout {
			insertTxOut(&vout, v.Txid, v.Hash)
		}
		insertTransaction(v)
	}
}

func main() {
	//     parse command line options
	basicAuth := &btcrpc.BasicAuth{
		Username: "testuser",
		Password: "123456",
	}
	client.RPCClient = connect_to_bitcoin_RPC(basicAuth)
	client.update_block_info()
}
