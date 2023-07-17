// © 2023 Nokia
// Licensed under the Apache License 2.0
// SPDX-License-Identifier: Apache-2.0

package regularReleases

import (
	"embed"
	"fmt"
	"time"
	"github.com/ossf/scorecard/v4/checker"
	"github.com/ossf/scorecard/v4/finding"
	//"github.com/ossf/scorecard/v4/probes/utils"
)

//go:embed *.yml
var fs embed.FS

var timeOneYearAgo = time.Now().AddDate(-1, 0, 0)

const probe = "regularReleases"

// Currently working but the wrong raw data is retrieved.
func Run(raw *checker.RawResults) ([]finding.Finding, string, error) {
	rawMaintainedData := &raw.MaintainedResults
	createdAt := rawMaintainedData.CreatedAt
	return runReleases(createdAt)
	//fmt.Printf("[ALL GUD IN THE HOOD]\n%v\n%v", createdAt, timeOneYearAgo)

}

func runReleases(createdAt time.Time) ([]finding.Finding, string, error) {
	var findings []finding.Finding
	if releasedWithinOneYear(createdAt, timeOneYearAgo) == true {
		f, err := finding.NewPositive(fs, probe, printTime(createdAt), nil)
		if err != nil {
			return nil, probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
	} else {
		f, err := finding.NewNegative(fs, probe, printTime(createdAt), nil)
		if err != nil {
			return nil, probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
	}
	return findings, probe, nil
}

func printTime(createdAt time.Time) string {
	return fmt.Sprintf("latest release: %s", createdAt.Format("Jan _2 2006 15:04:05"))
}

func releasedWithinOneYear(createdAt time.Time, timeOneYearAgo time.Time) bool {
	// handle errors?
	if createdAt.After(timeOneYearAgo) == true {
		return true
	} else {
		return false
	}
}
