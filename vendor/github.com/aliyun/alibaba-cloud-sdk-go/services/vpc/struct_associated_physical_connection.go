package vpc

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

// AssociatedPhysicalConnection is a nested struct in vpc response
type AssociatedPhysicalConnection struct {
	CircuitCode                      string `json:"CircuitCode" xml:"CircuitCode"`
	VlanInterfaceId                  string `json:"VlanInterfaceId" xml:"VlanInterfaceId"`
	LocalGatewayIp                   string `json:"LocalGatewayIp" xml:"LocalGatewayIp"`
	PeerGatewayIp                    string `json:"PeerGatewayIp" xml:"PeerGatewayIp"`
	PeeringSubnetMask                string `json:"PeeringSubnetMask" xml:"PeeringSubnetMask"`
	PhysicalConnectionId             string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
	PhysicalConnectionStatus         string `json:"PhysicalConnectionStatus" xml:"PhysicalConnectionStatus"`
	PhysicalConnectionBusinessStatus string `json:"PhysicalConnectionBusinessStatus" xml:"PhysicalConnectionBusinessStatus"`
	PhysicalConnectionOwnerUid       string `json:"PhysicalConnectionOwnerUid" xml:"PhysicalConnectionOwnerUid"`
	VlanId                           string `json:"VlanId" xml:"VlanId"`
	LocalIpv6GatewayIp               string `json:"LocalIpv6GatewayIp" xml:"LocalIpv6GatewayIp"`
	PeerIpv6GatewayIp                string `json:"PeerIpv6GatewayIp" xml:"PeerIpv6GatewayIp"`
	PeeringIpv6SubnetMask            string `json:"PeeringIpv6SubnetMask" xml:"PeeringIpv6SubnetMask"`
	Status                           string `json:"Status" xml:"Status"`
	EnableIpv6                       bool   `json:"EnableIpv6" xml:"EnableIpv6"`
}
