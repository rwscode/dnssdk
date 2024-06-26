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

package internal

type (
	LineListResp struct {
		List []LineListRespLine `json:"list"`
	}
	LineListRespLine struct {
		Id   string `json:"id"`   // 线路id
		Name string `json:"name"` // 线路名称
	}
	DomainListResp struct {
		Total uint                   `json:"total"`
		List  []DomainListRespDomain `json:"list"`
	}
	DomainListRespDomain struct {
		Id          string   `json:"id"`           // 域名id => xxxxxxxxxxxx
		Name        string   `json:"name"`         // 域名名称 => example.com
		DnsServer   []string `json:"dns_server"`   // DNS服务器 => [ns1.com, ns2.com]
		RecordCount uint     `json:"record_count"` // 记录数 => 100
		Remark      string   `json:"remark"`       // 备注
		CreateTime  string   `json:"create_time"`  // 创建时间 => 2022-09-27 08:09:25
	}
	DomainAddResp struct {
		Id        string   `json:"id"`         // 域名id => xxxxxxxxxxxx
		DnsServer []string `json:"dns_server"` // DNS服务器 => [ns1.com, ns2.com]
	}
	RecordListResp struct {
		Total uint                   `json:"total"`
		List  []RecordListRespRecord `json:"list"`
	}
	RecordListRespRecord struct {
		Id         string `json:"id"`          // id => xxxxxxxxxxxx
		Record     string `json:"record"`      // 主机记录 => www
		Name       string `json:"name"`        // 名称 => www.example.com
		Type       string `json:"type"`        // 类型 => A
		Value      string `json:"value"`       // 记录值 => 1.1.1.1
		Line       string `json:"line"`        // 线路 => default
		TTL        uint   `json:"ttl"`         // TTL => 60
		MX         uint16 `json:"mx"`          // MX => 1
		Weight     uint   `json:"weight"`      // 权重 => 5
		Remark     string `json:"remark"`      // 备注
		Status     string `json:"status"`      // 状态
		CreateTime string `json:"create_time"` // 创建时间 => 2022-09-27 08:09:25
		UpdateTime string `json:"update_time"` // 修改时间 => 2022-09-27 08:09:25
	}
	RecordAddResp    struct{ RecordListRespRecord }
	RecordUpdateResp RecordAddResp
)
