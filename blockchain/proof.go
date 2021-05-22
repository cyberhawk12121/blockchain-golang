package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Take data from the block
// create a counter (nonce) which starts at 0 and goes till infinity
// create a hash of the data plus the counter
// check the hash to see if it meets a set of requirements, i.e., the difficulty level

// Requirements:
// First few bytes must contain 0s

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int // It's the requirement of the difficulty level
}

func CreateProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) // Lsh= Left shift
	pow := &ProofOfWork{b, target}
	return pow
}

// It combines different data which is then passed to initHash to be hashed
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		// we'll join these 4 byte[] instead of just previous hash and the data to make the current hash
		[][]byte{
			pow.Block.Prev_Hash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{}, // This is just a seperator
	)
	return data
}

// The initHash gives the hash which is evaluated here, i.e., if the hash is a worthy problem
func (pow *ProofOfWork) Run() (int, []byte) {
	var initHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		initHash.SetBytes(hash[:])
		if initHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

// To validate the POW
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// util
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
