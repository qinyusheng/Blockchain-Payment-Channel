package levdb

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var dbA2B *leveldb.DB
var dbB2A *leveldb.DB

//打开数据库
func init()  {
	fmt.Println("init")
	var err error
	//数据存储路径和一些初始文件
	dbA2B,err = leveldb.OpenFile("./levelDb/A2B",nil)
	if err != nil {
		log.Fatalln(err)
	}
	dbB2A,err = leveldb.OpenFile("./levelDb/B2A",nil)
	if err != nil {
		log.Fatalln(err)
	}
}

//存入数据
func SaveA2B(key string,value string)  {
	dbA2B.Put([]byte(key),[]byte(value),nil)
}

//取出数据
func GetA2B(key string) ([]byte, error)  {
	fmt.Println("A2B key = ", key)
	value,err := dbA2B.Get([]byte(key),nil)
	return value, err
}


//遍历数据库并返回键值
func SeekA2B() []string {
	var r []string
	iter := dbA2B.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := string(iter.Key())
		r = append(r, key)
	}
	iter.Release()
	err := iter.Error()
	if err != nil{
		fmt.Println("iter.Error() = ", err)
	}
	return r
}

//遍历数据库并返回数据数量
func CountA2B() int {
	var num int
	num = 0
	iter := dbA2B.NewIterator(nil, nil)
	for iter.Next() {
		num++
	}
	iter.Release()
	err := iter.Error()
	if err != nil{
		fmt.Println("iter.Error() = ", err)
	}
	return num
}

//删除数据
func DelA2B(key string)  {
	err := dbA2B.Delete([]byte(key), nil)
	if err != nil{
		fmt.Println("db.Delete() = ", err)
	}
	return
}

//存入数据
func SaveB2A(key string,value string)  {
	dbB2A.Put([]byte(key),[]byte(value),nil)
}

//取出数据
func GetB2A(key string) ([]byte, error)  {
	value,err := dbB2A.Get([]byte(key),nil)
	return value, err
}

//遍历数据库并返回键值
func SeekB2A() []string {
	var r []string
	iter := dbB2A.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := string(iter.Key())
		r = append(r, key)
	}
	iter.Release()
	err := iter.Error()
	if err != nil{
		fmt.Println("iter.Error() = ", err)
	}
	return r
}

//遍历数据库并返回数据数量
func CountB2A() int {
	var num int
	num = 0
	iter := dbB2A.NewIterator(nil, nil)
	for iter.Next() {
		num++
	}
	iter.Release()
	err := iter.Error()
	if err != nil{
		fmt.Println("iter.Error() = ", err)
	}
	return num
}

//删除数据
func DelB2A(key string)  {
	err := dbB2A.Delete([]byte(key), nil)
	if err != nil{
		fmt.Println("db.Delete() = ", err)
	}
	return
}