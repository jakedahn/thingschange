package models

import (
	"fmt"
	"testing"

	"./"
)

func getSampleCheck() models.Check {
	return models.Check{
		Owner: "23",
		Url:   "http://google.com/23",
		Md5:   "23",
	}
}

func getArrayOfSampleChecks(count int) models.CheckList {
	var checks = []models.Check{}

	for i := 0; i < count; i++ {
		c := models.Check{
			Owner: fmt.Sprintf("%s", i),
			Url:   fmt.Sprintf("http://google.com/%s", i),
			Md5:   fmt.Sprintf("%s", i),
		}
		checks = append(checks, c)
	}
	return models.CheckList{Collection: checks}
}

func TestCheckChanged(t *testing.T) {
	check := getSampleCheck()

	if check.Changed("23") == true {
		t.Errorf("Changed should be returning false, but it is returning true")
	}

	if !check.Changed("46") {
		t.Errorf("Changed should be returning true, but it is returning false")
	}
}
