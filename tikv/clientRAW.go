package main

import (
	"context"
	"fmt"
	"github.com/tikv/client-go/v2/txnkv"
	"strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		client, err := txnkv.NewClient([]string{"192.168.0.216:2379"})
		if err != nil {
			panic(err)
		}
		txn, err := client.Begin()
		if err != nil {
			panic(err)
		}
		concatenated := "foobarnew" + strconv.Itoa(i)
		fmt.Sprintln(concatenated)
		if err := txn.Set([]byte(concatenated), []byte("bar")); err != nil {
			{
				panic(err)
			}
		}
		if err := txn.Commit(context.TODO()); err != nil {
			panic(err)
		}

		if err := client.Close(); err != nil {
			panic(err)
		}

	}

}
