# dirtywords
Inspired by [gau](https://github.com/lc/gau), dirtywords builds targeted wordlists for a given domain using "dirty" knowledge from AlienVault's Open Threat Exchange, the Wayback Machine, and Common Crawl.

## Usage:
Example usage:
```
$ dirtywords -d example.com -nosubs -o example-list.txt
```

## Installation
To install, use:
```
go get -u github.com/un4gi/dirtywords
```