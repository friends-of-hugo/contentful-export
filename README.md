# Contentful Hugo Extractor

<img src="https://d33wubrfki0l68.cloudfront.net/21d38ec2ccdfaacf6adc0b9921add9d18406493a/e1bcd/assets/images/logos/contentful-dark.svg" width="180" /> &nbsp; <img src="https://gohugo.io/img/hugo-logo.png" width="150" />

This tool extracts all content from your Contentful space and makes it easily consumable by Hugo. You can run it locally or as part of a CI server like Travis.

## Install

### Go Install Method

Assuming Go (1.10 +) is installed as well as [dep](https://golang.github.io/dep/)
```
go get -u github.com/icyitscold/contentful-hugo
cd $GOPATH/src/github.com/icyitscold/contentful-hugo
dep ensure
go install
```

## Usage

The tool requires two parameters to work, a contentful space id and API key. These can be provided as command line flags or as environment variables

```
export CONTENTFUL_API_KEY=YOUR-ACCESS-KEY-HERE
export CONTENTFUL_API_SPACE=YOUR-ID-HERE
contentful-hugo
```

```
contentful-hugo --space-id [my_space_id] --api-key [my_content_delivery_key]

```

## Expected output

Contentful Hugo Extractor stores all content under the /content directory. For each content type, it makes a subdirectory. For each item, it creates a markdown file with the all properties in TOML format.

Special cases:
 - Items of type Homepage are stored as /content/_index
   - Note that there can only be one such item
 - Fields named mainContent are used for the main content of the markdown file
 - File names are based on the ID of the item to make it easily referencable from related items (for the machine, not humans)

## Configuration

### Configure YAML output

While the default output is in TOML format, it is also possible to output content in YAML format. To achieve this, ensure there is a `extract-config.toml` file in the work directory of contentful-hugo with the following content:

```
encoding = "yaml"
```
