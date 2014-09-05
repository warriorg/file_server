package server

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type FileInfo struct {
	Key  string `bson:"_id"`
	Name string
	Type string
	Size int
	Data []byte
	Date time.Time
}

var sess *mgo.Session

func init() {
	fmt.Println("repository init")

	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Printf("连接数据库失败")
		os.Exit(1)
	}
	// defer sess.Close()

	sess.SetMode(mgo.Eventual, true)
	sess.SetSyncTimeout(0)

	// sess.SetSafe(&mgo.Safe{})
	// err = sess.DB("file").DropDatabase()
	// if err != nil {
	// 	fmt.Printf("删除数据库失败:%v\n", err)
	// 	os.Exit(1)
	// }

	// col := sess.DB("file").C("file")
	// doc := FileInfo{Key: bson.NewObjectId().String(), Name: "哈哈哈", Date: time.Now()}
	// col.Insert(doc)
}

func Save(file *FileInfo) {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Printf("连接数据库失败")
		os.Exit(1)
	}
	defer sess.Close()
	file.Key = bson.NewObjectId().Hex()

	sess.SetMode(mgo.Eventual, true)
	sess.SetSyncTimeout(0)

	col := sess.DB("file").C("file")
	col.Insert(file)
}

func Find(key string) *FileInfo {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Printf("连接数据库失败")
		os.Exit(1)
	}
	defer sess.Close()

	sess.SetMode(mgo.Eventual, true)
	sess.SetSyncTimeout(0)

	col := sess.DB("file").C("file")
	file := &FileInfo{}
	err = col.Find(bson.M{"_id": key}).One(&file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return file
}

func Delete(key string) {
	sess, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Printf("连接数据库失败")
		os.Exit(1)
	}
	defer sess.Close()

	sess.SetMode(mgo.Eventual, true)
	sess.SetSyncTimeout(0)

	col := sess.DB("file").C("file")
	col.Remove(bson.M{"_id": key})
}
