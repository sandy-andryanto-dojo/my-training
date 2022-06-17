package main

// import (
// 	"fmt"
// 	"hash/crc32"
// )

// func main(){
// 	h := crc32.NewIEEE()
// 	h.Write([]byte("test"))
// 	v := h.Sum32()
// 	fmt.Println(v)
// }

// import (
// 	"fmt"
// 	"hash/crc32"
// 	"io/ioutil"
// )

// func getHash(filename string) (uint32, error){
// 	bs, err := ioutil.ReadFile(filename)
// 	if err != nil{
// 		return 0, err
// 	}
// 	h := crc32.NewIEEE()
// 	h.Write(bs)
// 	return h.Sum32(), nil
// }

// func main(){
// 	h1, err :=  getHash("test1.txt")
// 	if err != nil{
// 		return
// 	}
// 	h2, err := getHash("test2.txt")
// 	if err != nil{
// 		return
// 	}
// 	fmt.Println(h1, h2, h1 == h2)
// }


import (
	"fmt"
	"crypto/sha1"
)

func main(){
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}


























