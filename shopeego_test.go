package shopeego

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetShopInfo(t *testing.T) {
	a := assert.New(t)
	c := NewClient(&ClientOptions{
		Secret: "9fac05968a6a2218f6a8f3c9a551d917e1646399286e3c32b823975a44bed87f",
	})
	resp, err := c.GetShopInfo(&GetShopInfoRequest{
		PartnerID: 844022,
		ShopID:    123456,
		Timestamp: int(time.Now().Unix()),
	})
	a.NoError(err)
	a.Equal("wow", resp)
}
