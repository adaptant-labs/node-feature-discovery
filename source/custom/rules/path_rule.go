/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rules

import (
	"fmt"
	"os"
)

// Rule that matches on matching file system paths in the system
type PathRule []string

// Match file system paths on provided list
func (paths *PathRule) Match() (bool, error) {
	for _, path := range *paths {
		exists, err := fileExists(path)
		if err != nil {
			return false, fmt.Errorf("failed to validate path. #{err.Error()}")
		}
		if exists != true {
			return false, fmt.Errorf("path does not exist: %s", path)
		}
	}

	return true, nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}
