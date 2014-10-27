package models

import (
	"log"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type CheckList struct {
	Collection []Check `json:"results"`
}

type Check struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	Md5        string `json:"md5"`
	Created_at int    `json:"created_at"`
	Updated_at int    `json:"updated_at"`
}

func (self *Check) Save() {
	session, err := mgo.Dial("3.3.3.3")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("thingschange").C("check")
	err = c.Insert(self)
	if err != nil {
		log.Fatal(err)
	}

	// result := Person{}
	// err = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Phone:", result.Phone)

}

// has the page been updated?
func (self *Check) Changed(md5 string) bool {
	if self.Md5 != md5 {
		return true
	}
	return false
}
