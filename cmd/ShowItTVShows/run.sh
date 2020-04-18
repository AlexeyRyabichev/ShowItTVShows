#!/bin/bash
go build
nohup ./ShowItTVShows >>out.log 2>&1 &
