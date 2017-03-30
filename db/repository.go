package db

import (
	"errors"

	"github.com/arturfelipe/go-beer/beer"
	mgo "gopkg.in/mgo.v2"
)

// BeerCollection : Beer mongo collection
const BeerCollection = "beer"

// ErrDuplicatedBeer : Beer duplicated error
var ErrDuplicatedBeer = errors.New("Duplicated beer")

// Beer repository
type BeerRepository struct {
	session *mgo.Session
}

// Create a beer
func (r *BeerRepository) Create(p *beer.Beer) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(BeerCollection)
	err := collection.Insert(p)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedBeer
		}
	}
	return err
}
