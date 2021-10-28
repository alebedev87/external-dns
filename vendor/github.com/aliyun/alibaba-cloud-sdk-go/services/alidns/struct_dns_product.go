package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DnsProduct is a nested struct in alidns response
type DnsProduct struct {
	InstanceId            string `json:"InstanceId" xml:"InstanceId"`
	VersionCode           string `json:"VersionCode" xml:"VersionCode"`
	VersionName           string `json:"VersionName" xml:"VersionName"`
	StartTime             string `json:"StartTime" xml:"StartTime"`
	EndTime               string `json:"EndTime" xml:"EndTime"`
	StartTimestamp        int64  `json:"StartTimestamp" xml:"StartTimestamp"`
	EndTimestamp          int64  `json:"EndTimestamp" xml:"EndTimestamp"`
	Domain                string `json:"Domain" xml:"Domain"`
	BindCount             int64  `json:"BindCount" xml:"BindCount"`
	BindUsedCount         int64  `json:"BindUsedCount" xml:"BindUsedCount"`
	TTLMinValue           int64  `json:"TTLMinValue" xml:"TTLMinValue"`
	SubDomainLevel        int64  `json:"SubDomainLevel" xml:"SubDomainLevel"`
	DnsSLBCount           int64  `json:"DnsSLBCount" xml:"DnsSLBCount"`
	URLForwardCount       int64  `json:"URLForwardCount" xml:"URLForwardCount"`
	DDosDefendFlow        int64  `json:"DDosDefendFlow" xml:"DDosDefendFlow"`
	DDosDefendQuery       int64  `json:"DDosDefendQuery" xml:"DDosDefendQuery"`
	OverseaDDosDefendFlow int64  `json:"OverseaDDosDefendFlow" xml:"OverseaDDosDefendFlow"`
	SearchEngineLines     string `json:"SearchEngineLines" xml:"SearchEngineLines"`
	ISPLines              string `json:"ISPLines" xml:"ISPLines"`
	ISPRegionLines        string `json:"ISPRegionLines" xml:"ISPRegionLines"`
	OverseaLine           string `json:"OverseaLine" xml:"OverseaLine"`
	MonitorNodeCount      int64  `json:"MonitorNodeCount" xml:"MonitorNodeCount"`
	MonitorFrequency      int64  `json:"MonitorFrequency" xml:"MonitorFrequency"`
	MonitorTaskCount      int64  `json:"MonitorTaskCount" xml:"MonitorTaskCount"`
	RegionLines           bool   `json:"RegionLines" xml:"RegionLines"`
	Gslb                  bool   `json:"Gslb" xml:"Gslb"`
	InClean               bool   `json:"InClean" xml:"InClean"`
	InBlackHole           bool   `json:"InBlackHole" xml:"InBlackHole"`
	BindDomainCount       int64  `json:"BindDomainCount" xml:"BindDomainCount"`
	BindDomainUsedCount   int64  `json:"BindDomainUsedCount" xml:"BindDomainUsedCount"`
	DnsSecurity           string `json:"DnsSecurity" xml:"DnsSecurity"`
	PaymentType           string `json:"PaymentType" xml:"PaymentType"`
}
