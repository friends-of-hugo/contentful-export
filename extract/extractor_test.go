package extract

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/bhsi-cinch/contentful-hugo/read"
)

type MockStore struct{}

func (ms MockStore) MkdirAll(path string, perm os.FileMode) error {
	return nil
}

func (ms MockStore) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return nil
}

func (ms MockStore) ReadFromFile(path string) (result []byte, err error) {
	return nil, nil
}

type MockGetter struct {
	JSON []string
}

type MockReaderCloser struct {
	Reader io.Reader
}

func (mrc MockReaderCloser) Read(p []byte) (n int, err error) {
	return mrc.Reader.Read(p)
}

func (mrc MockReaderCloser) Close() error {
	return nil
}

var count int

func (mg MockGetter) Get(url string) (result io.ReadCloser, err error) {
	mrc := MockReaderCloser{strings.NewReader(mg.JSON[count])}
	count++
	return mrc, nil
}

func TestExtractor(t *testing.T) {
	testContent := `
	{
		"sys": {
		  "type": "Array"
		},
		"total": 4,
		"skip": 0,
		"limit": 200,
		"items": [
		  {
			"sys": {
			  "space": {
				"sys": {
				  "type": "Link",
				  "linkType": "Space",
				  "id": "fp8h0eoshqd0"
				}
			  },
			  "id": "TU4YuzKtKC6QIqceaAOQM",
			  "type": "Entry",
			  "createdAt": "2017-08-04T13:01:18.942Z",
			  "updatedAt": "2017-08-12T05:07:50.535Z",
			  "revision": 6,
			  "contentType": {
				"sys": {
				  "type": "Link",
				  "linkType": "ContentType",
				  "id": "smallgroup"
				}
			  },
			  "locale": "en-US"
			},
			"fields": {
			  "title": "Joe & Joanne",
			  "slug": "joe-and-joanne",
			  "description": "Weekly bible study at Joe & Joanne place",
			  "locationText": "Deira",
			  "locationCoordinates": {
				"lon": 55.318323969841,
				"lat": 25.26726106496277
			  },
			  "weekday": "Sunday",
			  "time": "19:30",
			  "mainContent": "# Smallgroup\n\nJohn & Caroline host a group on Sunday evenings at 19:30.\n\nLocation: __Al Ghurair__\n\n\n---\n\n"
			}
		  },
		  {
			"sys": {
			  "space": {
				"sys": {
				  "type": "Link",
				  "linkType": "Space",
				  "id": "fp8h0eoshqd0"
				}
			  },
			  "id": "2ZZNV02YX6sI80Coc6MSwa",
			  "type": "Entry",
			  "createdAt": "2017-08-04T13:01:56.833Z",
			  "updatedAt": "2017-08-12T04:55:40.707Z",
			  "revision": 6,
			  "contentType": {
				"sys": {
				  "type": "Link",
				  "linkType": "ContentType",
				  "id": "smallgroup"
				}
			  },
			  "locale": "en-US"
			},
			"fields": {
			  "title": "Jim & Jane",
			  "slug": "jim-and-jane",
			  "description": "To do",
			  "locationText": "Silicon Oasis",
			  "locationCoordinates": {
				"lon": 55.38626380000005,
				"lat": 25.1279484
			  },
			  "weekday": "Tuesday",
			  "time": "19:30",
			  "mainContent": "To do: Add more text"
			}
		  },
		  {
			"sys": {
			  "space": {
				"sys": {
				  "type": "Link",
				  "linkType": "Space",
				  "id": "fp8h0eoshqd0"
				}
			  },
			  "id": "6CW1y3x6CWCKQoC8UwwAkC",
			  "type": "Entry",
			  "createdAt": "2017-08-04T13:05:35.498Z",
			  "updatedAt": "2017-08-12T04:52:41.342Z",
			  "revision": 14,
			  "contentType": {
				"sys": {
				  "type": "Link",
				  "linkType": "ContentType",
				  "id": "homepage"
				}
			  },
			  "locale": "en-US"
			},
			"fields": {
			  "mainHeader": "Welcome to the \"Test\" Site of Tubia!",
			  "test": 7,
			  "listOfSmallgroups": [
				{
				  "sys": {
					"type": "Link",
					"linkType": "Entry",
					"id": "2ZZNV02YX6sI80Coc6MSwa"
				  }
				},
				{
				  "sys": {
					"type": "Link",
					"linkType": "Entry",
					"id": "2P1NfVrIJ2yQQqOIiUewOW"
				  }
				},
				{
				  "sys": {
					"type": "Link",
					"linkType": "Entry",
					"id": "TU4YuzKtKC6QIqceaAOQM"
				  }
				}
			  ],
			  "mainContent": "This is the main content of the homepage!\n\nAnd here are a few more lines...\n\n# And this is a header"
			}
		  },
		  {
			"sys": {
			  "space": {
				"sys": {
				  "type": "Link",
				  "linkType": "Space",
				  "id": "fp8h0eoshqd0"
				}
			  },
			  "id": "2P1NfVrIJ2yQQqOIiUewOW",
			  "type": "Entry",
			  "createdAt": "2017-08-04T13:02:53.886Z",
			  "updatedAt": "2017-08-17T14:57:48.379Z",
			  "revision": 10,
			  "contentType": {
				"sys": {
				  "type": "Link",
				  "linkType": "ContentType",
				  "id": "smallgroup"
				}
			  },
			  "locale": "en-US"
			},
			"fields": {
			  "title": "Jaideep & Jyoti",
			  "slug": "jaideep-and-jyoti",
			  "description": "Short descr",
			  "locationText": "Garhoud",
			  "locationCoordinates": {
				"lon": 55.35490795969963,
				"lat": 25.239319871287446
			  },
			  "weekday": "Wednesday",
			  "time": "20:00",
			  "mainContent": "To do"
			}
		  }
		]
	  }
	`
	testTypes := `
	{
		"sys": {
		  "type": "Array"
		},
		"total": 3,
		"skip": 0,
		"limit": 200,
		"items": [
		  {
			"sys": {
			  "space": {
				"sys": {
				  "type": "Link",
				  "linkType": "Space",
				  "id": "fp8h0eoshqd0"
				}
			  },
			  "id": "2PqfXUJwE8qSYKuM0U6w8M",
			  "type": "ContentType",
			  "createdAt": "2017-03-16T17:51:06.624Z",
			  "updatedAt": "2017-03-16T17:51:06.624Z",
			  "revision": 1
			},
			"displayField": "productName",
			"name": "Product",
			"description": null,
			"fields": [
			  {
				"id": "productName",
				"name": "Product name",
				"type": "Text",
				"localized": false,
				"required": true,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "slug",
				"name": "Slug",
				"type": "Symbol",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "productDescription",
				"name": "Description",
				"type": "Text",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "sizetypecolor",
				"name": "Size/Type/Color",
				"type": "Symbol",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "image",
				"name": "Image",
				"type": "Array",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false,
				"items": {
				  "type": "Link",
				  "validations": [],
				  "linkType": "Asset"
				}
			  },
			  {
				"id": "tags",
				"name": "Tags",
				"type": "Array",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false,
				"items": {
				  "type": "Symbol",
				  "validations": []
				}
			  },
			  {
				"id": "categories",
				"name": "Categories",
				"type": "Array",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false,
				"items": {
				  "type": "Link",
				  "validations": [
					{
					  "linkContentType": [
						"6XwpTaSiiI2Ak2Ww0oi6qa"
					  ]
					}
				  ],
				  "linkType": "Entry"
				}
			  },
			  {
				"id": "price",
				"name": "Price",
				"type": "Number",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "brand",
				"name": "Brand",
				"type": "Link",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false,
				"linkType": "Entry"
			  },
			  {
				"id": "quantity",
				"name": "Quantity",
				"type": "Integer",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "sku",
				"name": "SKU",
				"type": "Symbol",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "website",
				"name": "Available at",
				"type": "Symbol",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  }
			]
		  },
		  {
			"sys": {
			  "space": {
				"sys": {
				  "type": "Link",
				  "linkType": "Space",
				  "id": "fp8h0eoshqd0"
				}
			  },
			  "id": "homepage",
			  "type": "ContentType",
			  "createdAt": "2017-08-04T12:59:45.163Z",
			  "updatedAt": "2017-08-11T16:40:02.258Z",
			  "revision": 2
			},
			"displayField": "mainHeader",
			"name": "Homepage",
			"description": "One and only one",
			"fields": [
			  {
				"id": "mainHeader",
				"name": "Main header",
				"type": "Symbol",
				"localized": false,
				"required": true,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "test",
				"name": "test",
				"type": "Integer",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "listOfSmallgroups",
				"name": "List of Smallgroups",
				"type": "Array",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false,
				"items": {
				  "type": "Link",
				  "validations": [
					{
					  "linkContentType": [
						"smallgroup"
					  ]
					}
				  ],
				  "linkType": "Entry"
				}
			  },
			  {
				"id": "mainContent",
				"name": "Main Content",
				"type": "Text",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  }
			]
		  },
		  {
			"sys": {
			  "space": {
				"sys": {
				  "type": "Link",
				  "linkType": "Space",
				  "id": "fp8h0eoshqd0"
				}
			  },
			  "id": "smallgroup",
			  "type": "ContentType",
			  "createdAt": "2017-08-04T11:42:20.401Z",
			  "updatedAt": "2017-08-17T14:57:11.135Z",
			  "revision": 9
			},
			"displayField": "title",
			"name": "Smallgroup",
			"description": "",
			"fields": [
			  {
				"id": "title",
				"name": "Title",
				"type": "Symbol",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "slug",
				"name": "Slug",
				"type": "Symbol",
				"localized": false,
				"required": true,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "description",
				"name": "Description",
				"type": "Text",
				"localized": true,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "locationText",
				"name": "LocationText",
				"type": "Symbol",
				"localized": false,
				"required": true,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "locationCoordinates",
				"name": "Location Coordinates",
				"type": "Location",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "weekday",
				"name": "Weekday",
				"type": "Symbol",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "time",
				"name": "Time",
				"type": "Symbol",
				"localized": false,
				"required": false,
				"disabled": false,
				"omitted": false
			  },
			  {
				"id": "mainContent",
				"name": "Main Content",
				"type": "Text",
				"localized": true,
				"required": false,
				"disabled": false,
				"omitted": false
			  }
			]
		  }
		]
	  }
	`

	extractor := Extractor{
		ReadConfig: read.ReadConfig{
			UsePreview:  false,
			SpaceID:     "my-fake-space-id",
			AccessToken: "my-fake-content-key",
			Locale:      "en-US",
		},
		Getter: MockGetter{[]string{testTypes, testContent}},
		RStore: MockStore{},
		WStore: MockStore{},
	}

	extractor.ProcessAll()

}
