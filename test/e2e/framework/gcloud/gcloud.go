/*
Copyright 2021 TriggerMesh Inc.

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

// Package gcloud contains helpers to interact with Google Cloud services.
package gcloud

import (
	"os"

	"github.com/triggermesh/triggermesh/test/e2e/framework"
)

const pubsubLabelOwnerResource = "io-triggermesh_owner-resource"

// PubSubResourceID returns a deterministic Pub/Sub resource ID matching the given framework.Framework.
func PubSubResourceID(f *framework.Framework) string {
	return f.UniqueName
}

// TagsFor returns a set of resource tags matching the given framework.Framework.
func TagsFor(f *framework.Framework) map[string]string {
	return map[string]string{
		pubsubLabelOwnerResource: f.UniqueName,
	}
}

// GetCreds returns the Google Cloud creds read from the environment.
func GetCreds(credsEnvVar string) string {
	return os.Getenv(credsEnvVar)
}

// GetProject returns the Google Cloud project read from the environment.
func GetProject(projectEnvVar string) string {
	return os.Getenv(projectEnvVar)
}