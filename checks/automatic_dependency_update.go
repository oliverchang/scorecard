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
	"strings"

	"github.com/ossf/scorecard/checker"
)

//nolint
func init() {
	registerCheck("Automatic-Dependency-Update", AutomaticDependencyUpdate)
}

//nolint
// AutomaticDependencyUpdate will check the repository if it contains Automatic dependecy update.
func AutomaticDependencyUpdate(c checker.Checker) checker.CheckResult {
	return CheckIfFileExists(c, fileExists)
}

// fileExists will validate the if frozen dependecies file name exists.
func fileExists(name string, logf func(s string, f ...interface{})) bool {
	return strings.Contains(name, strings.ToLower("dependabot.yml"))
}
