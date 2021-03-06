package utils

import (
	. "github.com/pingcap/check"
)

var _ = Suite(&TestSemverSuite{})

type TestSemverSuite struct{}

func (s *TestSemverSuite) TestSemverc(c *C) {
	cases := [][]interface{}{
		{"v0.0.1", "v0.0.1", true},
		{"0.0.1", "v0.0.1", true},
		{"invalid", "vinvalid", false},
		{"", "v", false},
		{"nightly", "nightly", true},
		{"Nightly", "Nightly", true},
	}

	for _, cas := range cases {
		v, e := FmtVer(cas[0].(string))
		c.Assert(v, Equals, cas[1].(string))
		c.Assert(e == nil, Equals, cas[2].(bool))
	}
}

func (s *TestSemverSuite) TestVersion(c *C) {
	c.Assert(Version("").IsValid(), IsFalse)
	c.Assert(Version("v3.0.").IsValid(), IsFalse)
	c.Assert(Version("").IsEmpty(), IsTrue)
	c.Assert(Version("").IsNightly(), IsFalse)
	c.Assert(Version("nightly").IsNightly(), IsTrue)
	c.Assert(Version("v1.2.3").String(), Equals, "v1.2.3")
}
