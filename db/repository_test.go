package db

import (
	"testing"

	"github.com/arturfelipe/go-beer/beer"

	check "gopkg.in/check.v1"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type BeerRepositorySuite struct {
	repository *BeerRepository
	session    *mgo.Session
	collection *mgo.Collection
}

var _ = check.Suite(&BeerRepositorySuite{})

func (s *BeerRepositorySuite) SetUpSuite(c *check.C) {
	session, _ := mgo.Dial("localhost:27017/beer-test")
	s.session = session
	s.collection = s.session.DB("").C(BeerCollection)
	s.repository = &BeerRepository{session: s.session}
}

func (s *BeerRepositorySuite) TearDownTest(c *check.C) {
	s.collection.RemoveAll(bson.M{})
}

func (s *BeerRepositorySuite) TearDownSuite(c *check.C) {
	defer s.session.Close()
}

func (s *BeerRepositorySuite) TestCreate(c *check.C) {
	s.repository.Create(&beer.Beer{ID: "1", Name: "La Chouffe"})
	count, _ := s.collection.Count()
	c.Check(count, check.Equals, 1)
}
