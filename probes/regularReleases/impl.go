// © 2023 Nokia
// Licensed under the Apache License 2.0
// SPDX-License-Identifier: Apache-2.0

package regularReleases

import (
	"embed"
	"fmt"
	"github.com/ossf/scorecard/v4/checker"
	"github.com/ossf/scorecard/v4/finding"
	//"github.com/ossf/scorecard/v4/probes/utils"
)

//go:embed *.yml
var fs embed.FS

const probe = "regularReleases"

func Run(raw *checker.RawResults) ([]finding.Finding, string, error) {
	rawReviewData := &raw.CodeReviewResults
	return CodeReviewRun(rawReviewData, fs, probe, finding.OutcomePositive, finding.OutcomeNegative)
}


/*
** Looks through the data and validates author and reviewers of a changeset
** Scorecard currently only supports GitHub revisions and generates a positive
** score in the case of other platforms. This probe is created to ensure that
** there are a number of unique reviewers for each changeset.
*/

func CodeReviewRun(reviewData *checker.CodeReviewData, fs embed.FS, probeID string,
	positiveOutcome, negativeOutcome finding.Outcome,
	) ([]finding.Finding, string, error) {
	var findings []finding.Finding
	
	return findings, probeID, nil
}