package read

import (
	"io"
	"strconv"
)

const previewURL string = "https://preview.contentful.com"
const URL string = "https://cdn.contentful.com"

type Contentful struct {
	Getter     Getter
	ReadConfig ReadConfig
}

// Types will use Contentful's content_types endpoint to retrieve all content types from contentful
func (c *Contentful) Types() (rc io.ReadCloser, err error) {

	return c.get("/spaces/" +
		c.ReadConfig.SpaceID + "/content_types?access_token=" +
		c.ReadConfig.AccessToken + "&limit=200&locale=" +
		c.ReadConfig.Locale)
}

// Items will use Contentful's entires endpoint to retrieve all 'items' from contetnful
func (c *Contentful) Items(skip int) (rc io.ReadCloser, err error) {

	return c.get("/spaces/" +
		c.ReadConfig.SpaceID + "/entries?access_token=" +
		c.ReadConfig.AccessToken + "&limit=200&locale=" +
		c.ReadConfig.Locale + "&skip=" + strconv.Itoa(skip))
}

func (c *Contentful) get(endpoint string) (rc io.ReadCloser, err error) {
	urlBase := URL
	if c.ReadConfig.UsePreview {
		urlBase = previewURL
	}

	return c.Getter.Get(urlBase + endpoint)
}
