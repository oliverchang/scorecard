// Copyright 2020 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checks

import (
	"fmt"

	"github.com/google/go-github/v32/github"
	"github.com/ossf/scorecard/checker"
)

const fuzzingStr = "Fuzzing"

func init() {
	registerCheck(fuzzingStr, Fuzzing)
}

func Fuzzing(c checker.CheckRequest) checker.CheckResult {
	url := fmt.Sprintf("github.com/%s/%s", c.Owner, c.Repo)
	searchString := url + " repo:google/oss-fuzz in:file filename:project.yaml"
	results, _, err := c.Client.Search.Code(c.Ctx, searchString, &github.SearchOptions{})
	if err != nil {
		return checker.MakeRetryResult(fuzzingStr, err)
	}

	if *results.Total > 0 {
		c.Logf("found project in OSS-Fuzz")
		return checker.CheckResult{
			Name:       fuzzingStr,
			Pass:       true,
			Confidence: checker.MaxResultConfidence,
		}
	}

	return checker.CheckResult{
		Name:       fuzzingStr,
		Pass:       false,
		Confidence: checker.MaxResultConfidence,
	}
}
