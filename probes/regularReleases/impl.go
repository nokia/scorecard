// © 2023 Nokia
// Licensed under the Apache License 2.0
// SPDX-License-Identifier: Apache-2.0

package regularReleases

import (
	"embed"
	"fmt"
	"time"

	"github.com/ossf/scorecard/v4/checker"
	"github.com/ossf/scorecard/v4/clients"
	"github.com/ossf/scorecard/v4/finding"
	//"github.com/ossf/scorecard/v4/probes/utils"
)

//go:embed *.yml
var fs embed.FS

var timeOneYearAgo = time.Now().AddDate(-1, 0, 0)

const probe = "regularReleases"

func Run(raw *checker.RawResults) ([]finding.Finding, string, error) {
	rawReleaseData := raw.SignedReleasesResults.Releases
	return runReleases(rawReleaseData)
}

// Looks through the releases, and returns positive findings if at least one was released in the past year.
// If no new releases are found, returns findings with negative outcome.
func runReleases(releases []clients.Release) ([]finding.Finding, string, error) {
	var findings []finding.Finding
	numReleasesWithinYear := 0
	//fmt.Printf("\n[ASSETS]\n%v\n", releases[0].Assets)
	if len(releases) == 0 /*|| len(releases[0].Assets) == 0*/ {
		f, err := finding.NewNegative(fs, probe, fmt.Sprintf("Found no releases for the project."), nil)
		if err != nil {
			return nil, probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
		return findings, probe, nil
	}
	for x := range releases {
		if len(releases[x].Assets) == 0 {
			continue
		}
		timeCreated := releases[x].Assets[0].CreatedAt
		if releasedWithinOneYear(timeCreated, timeOneYearAgo) == true {
			numReleasesWithinYear += 1
			} else {
				break
			}
	}
	if numReleasesWithinYear > 0 {
		f, err := finding.NewPositive(fs, probe, fmt.Sprintf("Found %v release(s) within the last year", numReleasesWithinYear), nil)
		if err != nil {
			return nil, probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
		timeCreated := releases[0].Assets[0].CreatedAt
		if releasedWithinOneYear(timeCreated, timeOneYearAgo) == true {
		f, err := finding.NewPositive(fs, probe, printTime(timeCreated), nil)
			if err != nil {
			return nil, probe, fmt.Errorf("create finding: %w", err)
			}
				findings = append(findings, *f)
		}
	} else {
		f, err := finding.NewNegative(fs, probe, fmt.Sprintf("Found 0 releases within the last year"), nil)
		if err != nil {
			return nil, probe, fmt.Errorf("create finding: %w", err)
		}
		findings = append(findings, *f)
	}
	return findings, probe, nil
}

func printTime(timeCreated time.Time) string {
	return fmt.Sprintf("latest release: %s", timeCreated.Format("Jan _2 2006 15:04:05"))
}

func releasedWithinOneYear(timeCreated time.Time, timeOneYearAgo time.Time) bool {
	// handle time errors?
	if timeCreated.After(timeOneYearAgo) == true {
		return true
	} else {
		return false
	}
}
