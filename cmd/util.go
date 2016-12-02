/*
Copyright 2015 All rights reserved.
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

package main

import (
	"fmt"

	"github.com/fcantournet/vault-sync/pkg/api"
	"github.com/fcantournet/vault-sync/pkg/utils"
)

// parseConfigFiles parses the configuration files and extracts the resources
func parseConfigFiles(files []string) (*resources, error) {
	r := new(resources)

	// step: iterate the configuration files and decode
	for _, c := range files {
		cfg := new(api.Config)
		if err := utils.DecodeFile(c, cfg); err != nil {
			return nil, fmt.Errorf("unable to decode the file: %s, error: %s", c, err)
		}
		// step: appends the elements
		r.users = append(r.users, cfg.Users...)
		r.backends = append(r.backends, cfg.Backends...)
		r.secrets = append(r.secrets, cfg.Secrets...)
		r.auths = append(r.auths, cfg.Auths...)
	}

	return r, nil
}
