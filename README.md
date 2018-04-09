# Contentful Hugo Extractor

<img src="https://d33wubrfki0l68.cloudfront.net/21d38ec2ccdfaacf6adc0b9921add9d18406493a/e1bcd/assets/images/logos/contentful-dark.svg" width="180" /> &nbsp; <img src="https://gohugo.io/img/hugo-logo.png" width="150" />

This tool extracts all content from your Contentful space and makes it easily consumable by Hugo. You can run it locally or as part of a CI server like Travis.

## Install

### Go Install Method

Assuming Go (1.10 +) is installed as well as [dep](https://golang.github.io/dep/)

``` sh
go get -u "github.com/bhsi-cinch/contentful-hugo"
cd "$GOPATH/src/github.com/bhsi-cinch/contentful-hugo"
dep ensure
go install
```

## Usage

``` sh
contentful-hugo [Flags]

Flags:
 -space-id=value      "Id of the contentful space from which to extract content. If not present will default to an environment variable named `$CONTENTFUL_API_SPACE`"
 -api-token=value     "API Key used to authenticate with contentful for content delivery. If not present will default to an environment variable named `$CONTENTFUL_API_KEY`"
 -config-file=value   "Path to the config TOML file to load. Defauls to `./extract-config.tml`"
 ```

The tool requires two parameters to work, a contentful space id and API key. These can be provided as command line flags or as environment variables

_As environment vars..._

``` sh
export CONTENTFUL_API_KEY=YOUR-CONTENT-DELIVERY-API-ACCESS-TOKEN-HERE
export CONTENTFUL_API_SPACE=YOUR-SPACE-ID-HERE

contentful-hugo
```

_As flags..._

``` sh
contentful-hugo -space-id=[YOUR-ID-HERE] -api-token=[YOUR-ACCESS-KEY-HERE] -config-file="./export-conf.toml"

```

## Expected output

Contentful Hugo Extractor stores all content under the /content directory. For each content type, it makes a subdirectory. For each item, it creates a markdown file with the all properties in TOML format.

Special cases:

- Items of type Homepage are stored as /content/_index
  - Note that there can only be one such item
- Fields named mainContent are used for the main content of the markdown file
- File names are based on the ID of the item to make it easily referencable from related items (for the machine, not humans)

## Configuration

Use the `--config-file` command line flag to provide the location of a TOML configuration to laod or ensure that there is a `extract-config.toml` file in the work directory of contentful-hugo

### Configure YAML output

While the default output is in TOML format, it is also possible to output content in YAML format. Use the following key in your config file:

``` yaml
encoding = "yaml"
```

### Configure Hugo Page Bundles

`contentful-hugo` will export each content type in contentful into its own content directory `./content/` and, since hugo treats each rootlevel content directory as a [Section][1], you will end up having a hugo section for each contentful content type. Hugo allows you to provide [Section][1] level configuration for its [Page Bundles](https://gohugo.io/content-management/page-bundles) by dropping a file named `_index.md` in the section's content directory. It is likely that you'll want to provide such configuration for some sections. 

For example, let's say you need to make a section [headless](https://gohugo.io/content-management/page-bundles/#headless-bundle). Pretend that you have a contentful content type with the id `question` and you have some questions in your contentful content model which you intend to reference in a seperate `FAQ` page. After a `contentful-hugo` export, you might the following directory structure:

``` none
./
|   content
|   |   _index.md
|   |   question
|   |   |   12h3jk213n.md   //question 1
|   |   |   sdfer343sn.md   //question 2
|   |   page
|   |   |   sdf234dd32.md   //FAQ page - refs questions in its frontmatter
|   layouts
|   |   _default
|   |   |   single.html
|   |   page
|   |   |   single.html //question refs are loaded via .Site.GetPage
```

Without any further confuguration, hugo would generate a HTML file for the page using the `./layouts/page/single.html` layout template but it would aslo generate HTML files for the questions using the `./layouts/_default/single.html` layout template. To prevent this from happening you would create the following file under the path `./content/question/index.md`: 

``` toml
+++
headless = true
+++
```

If you need this kind of configuration, the `contentful-hugo` export process can generate this `index.md` file for you.  Simply provide the TOML to use in your config file:

``` toml
encoding = "toml"
[section]
[section.question]
 headless = true
```

You can nest as many tables as you need under the `[sections]` and if the nested table name matches a contentful content type id than the configuration provided will be propagated to the section's `index.md` frontmatter. 

[1]: https://gohugo.io/content-management/sections/