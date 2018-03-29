package main

import (
	"fmt"

	"bitbucket.org/maxim_yefremov/kpi_dashboard"
)

func main() {
	context := kpi_dashboard.GetContext()

	kpi_dashboard.InitFlag(context)
	kpi_dashboard.Start(context)
	fmt.Println("Exiting program... ")
}
