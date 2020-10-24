[![Go Report Card](https://goreportcard.com/badge/github.com/caudaganesh/go-generator)](https://goreportcard.com/report/github.com/caudaganesh/go-generator)
[![codecov](https://codecov.io/gh/caudaganesh/go-generator/branch/master/graph/badge.svg?token=TJXKV5O5EL)](https://codecov.io/gh/caudaganesh/go-generator)

# go-generator

## Introduction

This repo contains multiple golang generators

## Prerequisite
- go 1.12 later

## Installation
    make install

## Available Commands
    gogen -h : show the help
    gogen proto : generating proto
    gogen interface : generating interface

## Flags
### Interface Generator
    -c, --comment string   comment for the generated interface
    -f, --file string      file path of the target struct
    -h, --help             help for interface
    -n, --name string      name for the generated interface
    -o, --output string    destination output of the result
    -p, --pkg string       package of the target struct
    -e, --pkgName string   package name for the generated interface
    -r, --target string    target struct for interface generator

### Proto Generator
    -f, --file string      file path to target struct
    -g, --goPkg string     go package for generated proto (default "./proto")
    -h, --help             help for proto
    -n, --name string      message name for generated proto
    -o, --output string    destination output of the result
    -p, --pkgName string   package name for generated proto (default "proto")
    -r, --target string    target struct

### App Layer Generator
    -p, --pkg string      package containing struct
    -r, --target string   target struct