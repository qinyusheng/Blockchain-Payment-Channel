package net1

import (
	"bytes"
	"encoding/json"
	"github.com/bitgoin/tx1"
)

// 1号信息容器
type Mes1 struct {
	 Transaction []byte
	 Sign []byte
}

type Mes2 struct {
	T1 []byte
	T2 []byte
}


type Mes3 struct {
	T1 []byte
	T2 []byte
	T3 []byte
}

type RealCTx struct {
	Version  uint32 `json:"version"`
	TxIn     []tx.TxIn  `json:"txin"`
	TxOut    []tx.TxOut `json:"txout"`
	Locktime uint32 `json:"locktime"`
	Addr 	 string `json:"addr"`
}

func CTx2Real(ctx tx.CTx) RealCTx{
	var reall RealCTx
	reall.Addr = ctx.Addr
	reall.Locktime = ctx.Locktime
	reall.Version = ctx.Version

	var tar int = 0
	for _,in := range ctx.TxIn {
		tar = 1
		reall.TxIn = append(reall.TxIn, *in)
	}
	if tar == 0 {
		ctx.TxIn = nil;
	}

	tar = 0
	for _,out := range ctx.TxOut {
		tar = 1
		reall.TxOut = append(reall.TxOut, *out)
	}
	if tar == 0 {
		ctx.TxIn = nil;
	}

	return reall
}

func Real2Ctx (reall RealCTx) tx.CTx {
	var ctx tx.CTx
	ctx.Addr = reall.Addr
	ctx.Version = reall.Version
	ctx.Locktime = reall.Locktime

	var tar int = 0
	for _,in := range reall.TxIn {
		tar = 1
		a := in
		ctx.TxIn = append(ctx.TxIn, &a)
	}
	if tar == 0 {
		ctx.TxIn = nil;
	}

	tar = 0
	for _,out := range reall.TxOut {
		tar = 1
		b := out
		ctx.TxOut = append(ctx.TxOut, &b)
	}
	if tar == 0 {
		ctx.TxOut = nil;
	}


	return ctx
}

func Tx2Ctx (t tx.Tx) tx.CTx {
	var ctx tx.CTx
	ctx.Locktime = t.Locktime
	ctx.Version = t.Version
	ctx.TxOut = t.TxOut[:]
	ctx.TxIn = t.TxIn[:]
	return ctx
}

func Ctx2Tx (ctx tx.CTx) tx.Tx {
	var t tx.Tx
	t.TxIn = ctx.TxIn[:]
	t.TxOut = ctx.TxOut[:]
	t.Version = ctx.Version
	t.Locktime = ctx.Locktime
	return t
}

func Ctx2Byte(ctx tx.CTx) []byte {
	b,_ := json.Marshal(CTx2Real(ctx))
	
	var c tx.CTx
	json.Unmarshal(b, &c)
	return b
}

func Byte2Ctx(bs []byte) tx.CTx {
	var reall RealCTx
	err := json.Unmarshal(bs, &reall)
	if err!=nil {
		panic(err)
	}
	
	return Real2Ctx(reall)
}

func Tx2Byte(t tx.Tx) []byte {
	return Ctx2Byte(Tx2Ctx(t))
}

func Byte2Tx(bs []byte) tx.Tx {
	return Ctx2Tx(Byte2Ctx(bs))
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
