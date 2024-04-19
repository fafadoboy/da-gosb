package models

import "fmt"

type City string

func (c City) Hash() string {
	return fmt.Sprint(c)
}

const (
	SEATTLE      City = "Seattle"
	SANFRANCISCO City = "SanFrancisco"
	LOSANGELES   City = "LosAngeles"
	RIVERSIDE    City = "Riverside"
	PHOENIX      City = "Phoenix"
	CHICAGO      City = "Chicago"
	BOSTON       City = "Boston"
	NEWYORK      City = "NewYork"
	ATLANTA      City = "Atlanta"
	MIAMI        City = "Miami"
	DALLAS       City = "Dallas"
	HOUSTON      City = "Houston"
	DETROIT      City = "Detroit"
	PHILADELPHIA City = "Philadelphia"
	WASHINGTON   City = "Washington"
)
