package read

import (
	"io"
	"strconv"
)

type Contentful struct {
	Getter     Getter
	ReadConfig ReadConfig
}

func (c *Contentful) Types() (rc io.ReadCloser, err error) {
	return c.Getter.Get(c.ReadConfig.UrlBase + "/spaces/" +
		c.ReadConfig.SpaceID + "/content_types?access_token=" +
		c.ReadConfig.AccessToken + "&limit=200&locale" +
		c.ReadConfig.Locale)
}

func (c *Contentful) Items(skip int) (rc io.ReadCloser, err error) {

	return c.Getter.Get(c.ReadConfig.UrlBase + "/spaces/" +
		c.ReadConfig.SpaceID + "/entries?access_token=" +
		c.ReadConfig.AccessToken + "&limit=200&locale" +
		"&skip=" + strconv.Itoa(skip) +
		c.ReadConfig.Locale)
}
