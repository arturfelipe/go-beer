package beer

import (
	"testing"

	check "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type BeerSuite struct {
	beer *Beer
}

var _ = check.Suite(&BeerSuite{})

func (s *BeerSuite) SetUpSuite(c *check.C) {
	s.beer = &Beer{}
}

func (s *BeerSuite) TestBeerData(c *check.C) {
	c.Check(s.beer.ID, check.Equals, "")
	c.Check(s.beer.Name, check.Equals, "")
	c.Check(s.beer.Description, check.Equals, "")
	c.Check(s.beer.Style, check.Equals, "")
	c.Check(s.beer.Abv, check.Equals, 0)
}
