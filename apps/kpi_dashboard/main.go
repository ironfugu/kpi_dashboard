package main

import (
	"fmt"

	"github.com/ypapax/kpi_dashboard"
)

func main() {
	context := kpi_dashboard.GetContext()

	kpi_dashboard.InitFlag(context)
	kpi_dashboard.Start(context)
	fmt.Println("Exiting program... ")
}
