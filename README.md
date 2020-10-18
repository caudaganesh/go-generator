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
    -p, --pkgName string   package name for the generated interface
    -r, --target string    target struct for interface generator

### Proto Generator
    -f, --file string      file path to target struct
    -g, --goPkg string     go package for generated proto (default "./proto")
    -h, --help             help for proto
    -n, --name string      message name for generated proto
    -o, --output string    destination output of the result
    -p, --pkgName string   package name for generated proto (default "proto")
    -r, --target string    target struct