# spotcsv

[![build](https://github.com/tamada/spotcsv/actions/workflows/build.yaml/badge.svg)](https://github.com/tamada/spotcsv/actions/workflows/build.yaml)
[![Coverage Status](https://coveralls.io/repos/github/tamada/spotcsv/badge.svg?branch=main)](https://coveralls.io/github/tamada/spotcsv?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/tamada/spotcsv)](https://goreportcard.com/report/github.com/tamada/spotcsv)

![Version](https://img.shields.io/badge/Version-$VERSION-blue)
[![License](https://img.shields.io/badge/License-MIT-blue)](LICENSE)

A tool for manipulating the CSV files.

## :speaking_head: Overview

This is a simple CLI tool that can manipulate CSV files.
This is an example in my course in the Graduate School.

## :runner: Usage

```sh
spotcsv <COMMANDs...> [CSV_FILE]
COMMAND
    headers
    head    <NUMBERS>
    tail    <NUMBERS>
    lines   <RANGES>
    index   <INDEXES>
    except  <INDEXES>
    format  <FORMATTER>
CSV_FILE
    The manipulation target CSV file. If not given, stdin is used.
```

### :walking: Demo​

```sh
$ cat testdata/sample.csv
ID,owner,repo
1,tamada,spotcsv
2,tamada,wildcherry
3,tamada,wildcat
4,tamada,sibling
5,tamada,btmeister
6,tamada,gixor
7,tamada,totebag
8,tamada,pochi
$ spotcsv headers testdata/sample.csv
ID
owner
repo
$ spotcsv lines 2-4 testdata/sample.csv
2,tamada,wildcherry
3,tamada,wildcat
$ spotcsv except 1 testdata/sample.csv
owner,repo
tamada,spotcsv
tamada,wildcherry
tamada,wildcat
tamada,sibling
tamada,btmeister
tamada,gixor
tamada,totebag
tamada,pochi
```

## :smile: About

### :man_office_worker: ​Authors :woman_office_worker:

- [Haruaki Tamada](https://tamada.github.io/) [:octocat:](https://github.com/tamada)

## :jack_o_lantern: Icon

![icon](docs/assets/icon.svg)

### :hammer_and_wrench: Related Tools

- [BurntSushi/xsv](https://github.com/BurntSushi/xsv)