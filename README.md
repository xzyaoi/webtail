# WebTail

[![Build Status](https://travis-ci.org/xzyaoi/webtail.svg?branch=master)](https://travis-ci.org/xzyaoi/webtail)
[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/xzyaoi)
[![Go Report Card](https://goreportcard.com/badge/github.com/xzyaoi/webtail)](https://goreportcard.com/report/github.com/xzyaoi/webtail)

> Tail your Logs to Browser

## Getting Started

1. Download binary executable program for you platform from [Releases](https://github.com/xzyaoi/webtail/releases)

2. ``` tar -zxvf webtail_VERSION_PLATFORM.tar.gz ```

3. ``` chmod +x webtail ```

4. Put Configuration File (```.webtail.toml```) with the binary program.

5. Start Server by ``` ./webtail ```

6. Open [http://webtail.cvtron.xyz](http://webtail.cvtron.xyz), and typein your server ip address and port number(8080 as default). Such as ```8.8.8.8:8080```

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
* [ ] Use token to guard the socket
* [ ] Allow Optional stdout/stderr
* [ ] Multiple Program Selection in webpage