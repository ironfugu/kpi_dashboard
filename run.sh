#!/usr/bin/env bash
set -ex
cd $GOPATH/src/bitbucket.org/maxim_yefremov/kpi_dashboard/apps/kpi_dashboard
go install
kpi_dashboard -alsologtostderr
