package main

import (
	"fmt"
	"go_chain/blockchain"
	"strconv"
)

func main() {
	// fmt.Println("Hello")
	chain := blockchain.InitBlockChain()
	chain.AddBlock("2nd Block")
	chain.AddBlock("3rd Block")
	chain.AddBlock("4th Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash:--> %x", block.Prev_Hash)
		fmt.Printf("\nData:--> %s", block.Data)
		fmt.Printf("\nCurrent Hash:--> %x", block.Hash)

		pow := blockchain.CreateProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
