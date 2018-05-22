package main

import (
	"zhqGo/MyhashMap/HashMp"
)

func main() {
	////测试HashMp包下的散列因子是否平均分布
	//for i := 'a'; i < 'a'+16; i++ {
	//	//fmt.Printf("%c", i)
	//	s := fmt.Sprintf("%c",i)
	//	//fmt.Println(s)
	//	fmt.Println(s,"=",HashMp.HashCode(s))
	//
	//}

	//fmt.Println(HashMp.HashCode("abc"))
	HashMp.CreateBulet()
	//HashMp.AddKeyValue("e", "hello")
	//HashMp.AddKeyValue("g", "world")
	//HashMp.GetValueByKey("e")

	HashMp.AddKeyValue("name", "sweet")
	HashMp.AddKeyValue("hello", "world")
	//HashMp.GetValueByKey("头结点0")
}
