package main

import (
	"github.com/Verybadwizard/datachart/chart"
)

func main() {

	c := chart.Create()
	c.Add("POST")
	c.Add("Test")
	c.Display()
}
