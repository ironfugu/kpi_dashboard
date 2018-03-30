#!/usr/bin/env bash
set -ex
REPO=$GOPATH/src/bitbucket.org/maxim_yefremov/kpi_dashboard
cd $REPO/apps/kpi_dashboard
go install
cd $REPO
kpi_dashboard -alsologtostderr
