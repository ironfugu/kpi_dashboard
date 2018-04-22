#!/usr/bin/env bash
set -ex
REPO=$GOPATH/src/github.com/ypapax/kpi_dashboard
cd $REPO/apps/kpi_dashboard
go get ./...
go install
cd $REPO
kpi_dashboard -alsologtostderr
