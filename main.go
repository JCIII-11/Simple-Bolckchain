package main

import "fmt"
import "crypto/sha256"
import "encoding/hex"
import "encoding/json"
import "time"

type Goblock struct{
	Prevhash string
	Timestamp string
	Hash string
	index int
	data string
}
func getHashSruct(a interface{}) (string) {
	data, err:= json.Marshal(a)
	if(err != nil){
		fmt.Println("Error:");
	}
		hash := sha256.Sum256(data)

		return hex.EncodeToString(hash[:])
}
func mkBlock(prevBolck Goblock, data string) Goblock{
	var newBlock Goblock
	newBlock.index = prevBolck.index+1
	newBlock.Timestamp = time.Now().String()
	newBlock.Prevhash = prevBolck.Hash
	newBlock.data = data
	newBlock.Hash = getHashSruct(newBlock)
	
	return newBlock;
}
var Blockchain[] Goblock

func genesis()Goblock{
	return Goblock{
		index: 0,
		Timestamp: time.Now().String(),
		Prevhash: "0",
		Hash: "",
		data: "First_Block",
	}
}
func addBlock(data string){
	B := mkBlock(Blockchain[len(Blockchain)-1], data)
	Blockchain = append(Blockchain, B)
}
var z bool;
func main(){
		frBlock := genesis()
		frBlock.Hash = getHashSruct(frBlock)
		Blockchain = append(Blockchain, frBlock)
		var com int;
		z = true
		var dat string
		for z{
			fmt.Println("What do you want:")
			fmt.Println("1: Add Block:")
			fmt.Println("2: Get Blockchain:")
			fmt.Println("3: Exit")
			fmt.Scanln(&com)
			if com == 1{
				fmt.Println("Type a data: ")
				fmt.Scanln(&dat)
				addBlock(dat)
			}
			if com == 2{
				for i := 0; i < len(Blockchain); i++{
					fmt.Printf("Index: %d\n",Blockchain[i].index)
					fmt.Printf("Data: %s\n",Blockchain[i].data)
					fmt.Printf("Timestamp: %s\n",Blockchain[i].Timestamp)
					fmt.Printf("SHA-256: %s\n",Blockchain[i].Hash)
					fmt.Println("--------------")
				}
			}
			if com == 3{
				break
			}
		}
}