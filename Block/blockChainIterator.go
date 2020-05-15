package Block

import "github.com/boltdb/bolt"

//区块链读取数据 迭代器
type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{
		bc.tip,
		bc.DB,
	}
	return bci
}

func (i *BlockChainIterator) Next() *BlockData {
	var block *BlockData

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodeBlock := b.Get(i.currentHash)
		block = Deserialize(encodeBlock)
		return nil
	})
	if err != nil {
		return nil
	}

	i.currentHash = block.PrevBlockHash

	return block
}
