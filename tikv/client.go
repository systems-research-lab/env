package main

import (
	"context"
	"fmt"
	"github.com/tikv/client-go/v2/rawkv"
	"strconv"
)

func main() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		client, err := rawkv.NewClientWithOpts(context.TODO(), []string{"192.168.0.216:2379"})
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		concatenated := "foobarnew" + strconv.Itoa(i)
		val := "bar" + strconv.Itoa(i)
		fmt.Println(concatenated)
		if err := client.Put(context.TODO(), []byte(concatenated), []byte(val)); err != nil {
			{
				panic(err)
			}
		}
		v, err := client.Get(context.TODO(), []byte(concatenated))
		if err != nil {
			panic(err)
		}
		fmt.Println("Get")
		fmt.Println(v)

		//if err := txn.Commit(context.TODO()); err != nil {
		//	panic(err)
		//}

		if err := client.Close(); err != nil {
			panic(err)
		}

	}

}
