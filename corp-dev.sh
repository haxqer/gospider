#!/bin/sh


env GOOS=linux go build
env GOOS=linux go build ./services/itemsave
env GOOS=linux go build ./services/spiderhttp



