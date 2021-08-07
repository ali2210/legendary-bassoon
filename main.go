package main



import (
	http_ipfs_client "github.com/ipfs/go-ipfs-http-client"
	"fmt"
	"context"
	core "github.com/ipfs/interface-go-ipfs-core/path"
	gocid "github.com/ipfs/go-cid"
	mulhash "github.com/multiformats/go-multihash"
	 "crypto/sha256"
	"encoding/hex"
	//"strings"
)

func main() {
	multaddr, err := http_ipfs_client.ApiAddr("~/.ipfs"); if err != nil {
		fmt.Println("multi-address error:", err.Error())
		return 
	}
	fmt.Println("@multi-addr:", multaddr)

	client_http_api , err := http_ipfs_client.NewApi(multaddr); if err != nil {
		fmt.Println("http client error:", err.Error())
		return 
	}

	fmt.Println("@http-client:", client_http_api)
	block := client_http_api.Block()
	fmt.Println("@block:", block)
	
	hash := sha256.Sum256([]byte("hello-world"))
	hex_Data, err  := hex.DecodeString(hex.EncodeToString(hash[:])); if err != nil {
		fmt.Println("decode string error:", err.Error())
		return 
	}	
	
	mulhashdata, err := mulhash.EncodeName(hex_Data, "sha2-256"); if err != nil {
		fmt.Println("multi-hash encode name error:", err.Error())
		return 
	}
	
	mulHash := hex.EncodeToString(mulhashdata)
	multihash , err:= mulhash.FromHexString(mulHash); if err != nil {
		fmt.Println("multi-hash hex error:", err.Error())
		return 
	}
	
	cid := gocid.NewCidV0(multihash)
	fmt.Println("@go - cid: ", cid)
	
	resolved := core.IpfsPath(cid)
	fmt.Println("resolved _ ipfs:", resolved)
	block_value, err := block.Stat(context.Background(), core.New("/ipfs/"+ resolved.Cid().String())); if err != nil {
	 	fmt.Println("block error:", err.Error())
	 	return 
	 }
	fmt.Println("@block_stats:", block_value.Path(), block_value.Size())
}