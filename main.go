package main



import (
	http_ipfs_client "github.com/ipfs/go-ipfs-http-client"
	"fmt"
	 "context"
	p2ps "github.com/libp2p/go-libp2p-core/peer"
	paths "github.com/ipfs/interface-go-ipfs-core/path"
	core "github.com/ipfs/interface-go-ipfs-core"
	multiAddr "github.com/multiformats/go-multiaddr"
	gocid "github.com/ipfs/go-cid"
	mulhash "github.com/multiformats/go-multihash"
	"crypto/sha256"
	"encoding/hex"
	//"errors"
	//"reflect"
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
	
	resolved := paths.IpfsPath(cid)
	fmt.Println("resolved _ ipfs:", resolved)
	
	pId, err := p2ps.Decode(resolved.Cid().String()); if err != nil {
		fmt.Println("peer id error:", err.Error())
		return 
	}

	fmt.Println("@peer _ id:", pId)
	addr , err := multiAddr.NewMultiaddr(multaddr.String()); if err != nil {
		fmt.Println("multi-address error:", err.Error())
		return 
	}	
	fmt.Println("@multi-address:", addr)
	addr_info, err := p2ps.AddrInfoFromString("/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWKSG7DZhTNLBqkRzzkvSBbThZ1EpWBsDP54iq6mWLkiv6"); if err != nil {
		fmt.Println("peer multi-address error:", err.Error())
		return 
	}
	fmt.Println("@peer_multiaddr_info:", addr_info)	
	var object core.ObjectAPI = core.Object()
	fmt.Println("@object:", object)
}