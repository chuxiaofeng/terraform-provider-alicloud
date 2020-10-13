package nas

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

// FileSystem is a nested struct in nas response
type FileSystem struct {
	FileSystemId          string                            `json:"FileSystemId" xml:"FileSystemId"`
	Description           string                            `json:"Description" xml:"Description"`
	CreateTime            string                            `json:"CreateTime" xml:"CreateTime"`
	ExpiredTime           string                            `json:"ExpiredTime" xml:"ExpiredTime"`
	RegionId              string                            `json:"RegionId" xml:"RegionId"`
	ZoneId                string                            `json:"ZoneId" xml:"ZoneId"`
	ProtocolType          string                            `json:"ProtocolType" xml:"ProtocolType"`
	StorageType           string                            `json:"StorageType" xml:"StorageType"`
	FileSystemType        string                            `json:"FileSystemType" xml:"FileSystemType"`
	EncryptType           int                               `json:"EncryptType" xml:"EncryptType"`
	MeteredSize           int64                             `json:"MeteredSize" xml:"MeteredSize"`
	MeteredIASize         int64                             `json:"MeteredIASize" xml:"MeteredIASize"`
	Bandwidth             int64                             `json:"Bandwidth" xml:"Bandwidth"`
	Capacity              int64                             `json:"Capacity" xml:"Capacity"`
	AutoSnapshotPolicyId  string                            `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
	Status                string                            `json:"Status" xml:"Status"`
	ChargeType            string                            `json:"ChargeType" xml:"ChargeType"`
	MountTargetCountLimit int64                             `json:"MountTargetCountLimit" xml:"MountTargetCountLimit"`
	NasNamespaceId        string                            `json:"NasNamespaceId" xml:"NasNamespaceId"`
	KMSKeyId              string                            `json:"KMSKeyId" xml:"KMSKeyId"`
	Version               string                            `json:"Version" xml:"Version"`
	SupportedFeatures     SupportedFeatures                 `json:"SupportedFeatures" xml:"SupportedFeatures"`
	Ldap                  Ldap                              `json:"Ldap" xml:"Ldap"`
	MountTargets          MountTargetsInDescribeFileSystems `json:"MountTargets" xml:"MountTargets"`
	Packages              PackagesInDescribeFileSystems     `json:"Packages" xml:"Packages"`
	Tags                  TagsInDescribeFileSystems         `json:"Tags" xml:"Tags"`
}