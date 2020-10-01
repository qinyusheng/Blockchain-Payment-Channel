package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bitgoin/address"
	"log"
	"time"
)
import "./tx"

func main() {
	fmt.Print("123");
	wif := "928Qr9J5oAC6AYieWJ3fG3dZDjuC7BFVUqgu4GsvRVpoXiTaJJf"
	//n3Bp1hbgtmwDtjQTpa6BnPPCA8fTymsiZy
	txKey, err := address.FromWIF(wif, address.BitcoinTest)
	if err != nil {
		panic(err)
	}
	adr := txKey.PublicKey.Address()
	log.Println("address for tx=", adr)

	wif2 := "92DUfNPumHzpCkKjmeqiSEDB1PU67eWbyUgYHhK9ziM7NEbqjnK"
	//ms5repuZHtBrKRE93FdWqz8JEo6d8ikM3k
	txKey2, err := address.FromWIF(wif2, address.BitcoinTest)
	if err != nil {
		panic(err)
	}
	txhashes := []string{
		"12c2f61d839b2b38146715e4dfc0fd914906253920480298816f108513e53e5c",
		"12c2f61d839b2b38146715e4dfc0fd988806253920480298816f108513e53e5c",
	}
	values := []uint64{100 * tx.Unit, 150 * tx.Unit}
	script, err := hex.DecodeString("76a914d94987ba89c258372030bc9d610f89547757896488ac")
	if err != nil {
		panic(err)
	}

	utxos := make(tx.UTXOs, len(txhashes))
	for i, h := range txhashes {
		var ha []byte
		ha, err = hex.DecodeString(h)
		if err != nil {
			panic(err)
		}
		ha = tx.Reverse(ha)
		utxos[i] = &tx.UTXO{
			Key:     txKey,
			TxHash:  ha,
			Value:   values[i],
			Script:  script,
			TxIndex: uint32(i + 1),
		}
	}

	payer := tx.NewMicroPayer(txKey, txKey2.PublicKey, 200*tx.Unit, 0.001*tx.Unit)
	payee := tx.NewMicroPayee(txKey.PublicKey, txKey2, 200*tx.Unit, 0.001*tx.Unit)
	locktime := uint32(time.Now().Add(time.Hour).Unix())

	bond, refund, err := payer.CreateBond(locktime, utxos, txKey.PublicKey.Address())
	if err != nil {
		panic(err)
	}
	sign, err := payee.SignRefund(refund, locktime)
	if err != nil {
		panic(err)
	}

	if err = payer.SignRefund(refund, sign); err != nil {
		panic(err)
	}
	if err = payee.CheckBond(refund, bond); err != nil {
		panic(err)
	}

	signIP, err := payer.SignIncremented(0.001 * tx.Unit)
	if err != nil {
		panic(err)
	}
	log.Println(hex.EncodeToString(signIP))
	tx, err := payee.IncrementedTx(0.001*tx.Unit, signIP)
	if err != nil {
		panic(err)
	}
	bbond, err := bond.Pack()
	if err != nil {
		panic(err)
	}
	bref, err := refund.Pack()
	if err != nil {
		panic(err)
	}
	btx, err := tx.Pack()
	if err != nil {
		panic(err)
	}
	log.Print("bond ", hex.EncodeToString(bbond))
	log.Print("refund ", hex.EncodeToString(bref))
	log.Print("incremented tx ", hex.EncodeToString(btx))
}

