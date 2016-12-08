package models

import (
	//"encoding/json"

	//"labix.org/v2/mgo"
	"gopkg.in/mgo.v2/bson"
	//"log"
	//"sync"
	//"fmt"
	//"time"
	"github.com/butterfli-go/store"
	"time"
	"fmt"
)

type AccountCreds struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Account	string           `json:"account",bson:"account,omitempty"`
	ConsumerKey		string           `json:"consumerKey",bson:"consumerKey,omitempty"`
	ConsumerSecret		string           `json:"consumerSecret",bson:"consumerSecret,omitempty"`
	AccessToken		string           `json:"accessToken",bson:"accessToken,omitempty"`
	AccessTokenSecret		string           `json:"accessTokenSecret",bson:"accessTokenSecret,omitempty"`
}

func NewAccountCreds(username string, accountId string, consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *AccountCreds {
	a := new(AccountCreds)
	a.Id = bson.NewObjectId()
	a.Timestamp = time.Now()
	a.Username = username
	a.Account = accountId
	a.ConsumerKey = consumerKey
	a.ConsumerSecret = consumerSecret
	a.AccessToken = accessToken
	a.AccessTokenSecret = accessTokenSecret

	return a
}

func (a *AccountCreds) Save() error {
	fmt.Print("saving! from the top ")
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToAccountCredCollection(session, "accountCreds")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(&AccountCreds{
		Id: a.Id,
		Timestamp: a.Timestamp,
		Username: a.Username,
		Account: a.Account,
		ConsumerKey: a.ConsumerKey,
		ConsumerSecret: a.ConsumerSecret,
		AccessToken: a.AccessToken,
		AccessTokenSecret: a.AccessTokenSecret})


	if err != nil {
		return err
	}

	return nil
}


func FindAccountCredsById(accountCredsId string) (*AccountCreds, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToPostsCollection(session, "accountCreds")
	if err != nil {
		//panic(err)
		return &AccountCreds{}, err
	}

	accountCreds := AccountCreds{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(accountCredsId)}).One(&accountCreds)
	if err != nil {
		panic(err)
		//return &post, err
	}

	return &accountCreds, err
}


func FindAccountCredsByAccountId(accountId string) (*AccountCreds, error) {
	fmt.Print("begin FACBAI")
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToAccountCredCollection(session, "accountCreds")
	if err != nil {
		//panic(err)
		return &AccountCreds{}, err
	}

	accountCreds := AccountCreds{}
	err = collection.Find(bson.M{"account": accountId}).One(&accountCreds)
	if err != nil {
		panic(err)
		//return &post, err
	}

	return &accountCreds, err
}