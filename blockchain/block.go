package blockchain

type BlockChain struct {
	Blocks []*Block // Pointers of blocks, made public when it starts with capital letter, so Block instead of block
}

// []byte is like char[] datatype in C/C++
type Block struct {
	//slice of uint[]
	Hash      []byte
	Data      []byte
	Prev_Hash []byte
	Nonce     int
}

// This runs first which creates the Hash for the blocks using the previous hash value and combines them
// // And then puts the value of created hash in the Block instance "b"	See --> b.Hash = hash[:]
// func (b *Block) CreateHash() {

// 	// Step 1: Use previous blocks hash and join it with data
// 	info := bytes.Join([][]byte{b.Data, b.Prev_Hash}, []byte{}) // The other parameter []byte slice is just a seperator which is used for joining the hashes
// 	// Step 2: Create a new SHA256 Hash using using that combined data (Last step)
// 	hash := sha256.Sum256(info)
// 	// Step 3: Now give that new hash value to the "hash" variable of the current hash
// 	b.Hash = hash[:]
// 	fmt.Println(reflect.TypeOf(info))
// }

// This will create a Block using the new data and Hash value that we just created
func CreateBlock(data string, prevHash []byte) *Block {
	// Step 1: Create a block reference with just an empty slice of byte as it's current []Hash
	// And put the data and prevHash values simply
	block := &Block{[]byte{}, []byte(data), prevHash, 0} // converted the string to []byte("String")
	// Step 2: This will then just populate the current Hash of the current block
	// block.CreateHash()
	// Block created now!

	pow := CreateProof(block)
	nonce, hash := pow.Run()
	// put the nonce and hash in the current block
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) { // It takes the data
	// Using the chain pointer of the Blockchain we'll access the contents of the blocks in the current BlockChain
	// Step 1: Get the previous block (I didn't know we can get the struct using <struct_name>[0,1,2..n] this way)
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// Step 2: Create a new &Block instance using previous block's Hash.
	new := CreateBlock(data, prevBlock.Hash)
	// Step 3: Append the new block in the BlockChain (Why tf are struct instances treated like a slice/array)
	chain.Blocks = append(chain.Blocks, new) // new block has been added
}

// Creating first Block
func Genesis() *Block {
	s := "Sameer"
	return CreateBlock(s, []byte{})
}

// Creating first Blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
