#!/bin/bash
goyacc ./parser.y 
go install github.com/AYJiaYou/restest/cmd/restest
restest
