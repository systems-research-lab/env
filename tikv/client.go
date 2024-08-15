package tikv

import "github.com/tikv/client-go/v2/txnkv"

func main() {

	client, err := txnkv.NewClient([]string{""})
	txn, err := client.Begin(opts...)
	if err != nil {
		// ... handle error ...
	}

	if err := txn.Set([]byte("foo"), []byte("bar")); err != nil {
		// ... handle error ...
	}
	if err := txn.Delete([]byte("foo")); err != nil {
		// ... handle error ...
	}

}
