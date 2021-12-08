#!/usr/bin/env bash

go build -tags 'game_server' -race -o bin/game_server github.com/chuang39/beego-boilerplate/src
