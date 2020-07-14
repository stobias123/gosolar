package gosolar

import (
	"github.com/stretchr/testify/assert"
)

func (suite *TestSuite) TestFetchEmptyIP() {
	ip := suite.client.GetIP("10.199.152.200")
	nullIP := []IPAddress{}
	assert.Empty(suite.T(), nullIP, ip)
}

func (suite *TestSuite) TestFetchIP() {
	ip := suite.client.GetIP("10.199.152.0")
	nullIP := []IPAddress{}
	assert.NotEqual(suite.T(), nullIP, ip)
}

func (suite *TestSuite) TestReserveIP() {
	ip := suite.client.ReserveIP("10.199.152.0")
	nullIP := []IPAddress{}
	assert.NotEqual(suite.T(), nullIP, ip)
}

func (suite *TestSuite) TestReserveIPHostname() {
	ip := suite.client.ReserveIPForHostname("10.199.152.0", "test123")
	nullIP := []IPAddress{}
	assert.NotEqual(suite.T(), nullIP, ip)
}
