# Contentful Hugo Extractor

<img src="https://www.contentful.com/assets/images/logos/contentful-dark-1a51a42b.svg" width="180" /> &nbsp; <img src="https://gohugo.io/img/hugo-logo.png" width="150" />

This tool extracts all content from your Contentful space and makes it easily consubable by Hugo. You can run it locally or as part of a CI server like Travis.

## Usage

The tool requires two environment variables to work:

```
export CONTENTFUL_KEY=YOUR-ACCESS-KEY-HERE
export SPACE_ID=YOUR-ID-HERE
contentful-hugo
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

While the default output is in TOML format, it is also possible to output content in YAML format. To achieve this, ensure there is a config.toml file in the work directory of contentful-hugo with the following content:

```
encoding = "yaml"
```
