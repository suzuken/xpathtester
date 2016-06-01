# XPath Tester

## Installation

	$ go get github.com/suzuken/xpathtester

## Usage

	Usage of xpathtester:
	  -url string
			url to fetch and extract (default "http://example.com")
	  -xpaths value
			xpaths is multiple xpath for extracting content. space separated. (default [])

## Examples

	$ xpathtester -url "https://github.com/suzuken" -xpaths '//h1 //h2'

## LICENSE

MIT
