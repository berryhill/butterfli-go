package models

import (
	"time"
	// "fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/butterfli-go/store"
	"fmt"
)

type Account struct {
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Title		string           	`json:"title",bson:"title,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	//Posts []*Post 			`json:"posts",bson:"posts,omitempty"`
}

func NewAccount(username string, title string) *Account {
	a := new(Account)
	a.Id = bson.NewObjectId()
	a.Timestamp = time.Now()
	a.Username = username
	a.Title = title
	//a.Posts = []*Post{}

	return a
}

func (a *Account) Save() error {
	fmt.Print("saving this accnt ")
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToAccountsCollection(session, "accounts")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(&Account{
		Id: a.Id,
		Timestamp: a.Timestamp,
		Title: a.Title,
		Username: a.Username,
		//Posts: a.Posts
	})

	if err != nil {
		return err
	}
	return nil
}

func FindAccount(username string, title string) (*Account, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToAccountsCollection(session, "accounts")
	if err != nil {
		//panic(err)
		return &Account{}, err
	}

	account := Account{}
	err = collection.Find(bson.M{"username": username, "title": title}).One(&account)
	if err != nil {
		panic(err)
	}

	return &account, err
}

func FindAccountById(account_id string) (*Account, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToAccountsCollection(session, "accounts")
	if err != nil {
		//panic(err)
		return &Account{}, err
	}

	account := Account{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(account_id)}).One(&account)
	if err != nil {
		panic(err)
	}

	return &account, err
}

func GetAllAccounts(username string) ([]*Account, error){
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection := session.DB("test").C("accounts")

	accounts := []*Account{}

	err = collection.Find(bson.M{"username": username}).All(&accounts)

	return accounts, err
}

