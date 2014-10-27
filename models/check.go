package models

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mgoSession *mgo.Session
)

func setupMongoConn() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"3.3.3.3"},
		Timeout:  60 * time.Second,
		Database: "thingschange",
	}

	mgoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mgoSession.SetMode(mgo.Monotonic, true)
	// defer mgoSession.Close()

	return mgoSession
}

type CheckList struct {
	Collection []Check `json:"results"`
}

type Check struct {
	Owner      string `json:"-"`
	Url        string `json:"url"`
	Md5        string `json:"md5"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
}

func (self *Check) Save() {
	mgo := setupMongoConn()
	c := mgo.DB("thingschange").C("check")
	err := c.Insert(self)
	if err != nil {
		log.Fatal(err)
	}
}

func GetChecksByOwner(api_key string) {
	mgo := setupMongoConn()
	c := mgo.DB("thingschange").C("check")

	log.Println("#####")
	log.Println(api_key)
	log.Println(c.Find(bson.M{}))
	log.Println("#####")
	// result := CheckList{Collection: {}}

	// err = c.Find(bson.M{"owner": api_key})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("res:", result)

}

// has the page been updated?
func (self *Check) Changed(md5 string) bool {
	if self.Md5 != md5 {
		return true
	}
	return false
}
