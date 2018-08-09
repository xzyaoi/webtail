# WebTail

[![Build Status](https://travis-ci.org/xzyaoi/webtail.svg?branch=master)](https://travis-ci.org/xzyaoi/webtail)
[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/xzyaoi)
[![Go Report Card](https://goreportcard.com/badge/github.com/xzyaoi/webtail)](https://goreportcard.com/report/github.com/xzyaoi/webtail)

> Tail your Logs to Browser

## Getting Started

## Configuration

When initiating WebTail Serve, you need a ```.webtail.toml``` in the folder where the webtail executable program locates.

```
[[program]]
name = "test"
stderr = "./test.log"
stdout = "./nohup.out"

[[program]]
name = "test2"
stderr = "./nohup.out"
stdout = "./test.log"

```

## Development

## To Dos

* [ ] Config Port in toml file