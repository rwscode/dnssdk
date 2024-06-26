// Copyright 2024 dnsdk Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dnsdk

func CloudflareSupporter[T any](supportFunc SupportFunc[T, *CloudflareSupportOpts]) supporter[T, *CloudflareSupportOpts] {
	return &defaultSupporter[T, *CloudflareSupportOpts]{ApiType: ApiTypeCloudflare, SupportFunc: supportFunc}
}

type CloudflareSupportOpts struct{ email, apiKey string }

func NewCloudflareSupportOpts(email string, apiKey string) *CloudflareSupportOpts {
	return &CloudflareSupportOpts{email, apiKey}
}
