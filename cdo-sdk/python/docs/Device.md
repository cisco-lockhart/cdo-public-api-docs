# Device


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the device in CDO. | [optional] 
**name** | **str** | The name of the device in CDO. Device names are unique in CDO. | 
**device_type** | [**EntityType**](EntityType.md) |  | 
**connector_type** | [**ConnectorType**](ConnectorType.md) |  | [optional] 
**connector_uid** | **str** | The unique identifier of the Secure Device Connector (SDC) that will be used to communicate with the device. This value is not required if the connector type selected is Cloud Connector (CDG)The name of the Secure Device Connector (SDC) that will be used to communicate with the device. This value is not required if the connector type selected is Cloud Connector (CDG). | [optional] 
**address** | **str** | The address of the device, in &#x60;host:port&#x60; format. CDO connects to the device at this address. | [optional] 
**device_role** | [**DeviceRole**](DeviceRole.md) |  | [optional] 
**serial** | **str** | The serial number of the device. This is typically used for licensing, and is not the same as the chassis&#39; serial number. | [optional] 
**chassis_serial** | **str** | The serial number on the chassis of the device (ASA-only). This is typically used to type up to Cisco SmartNet, and is not the same as the serial number. | [optional] 
**software_version** | **str** | The version of the software running on the device. | [optional] 
**connectivity_state** | [**ConnectivityState**](ConnectivityState.md) |  | [optional] 
**config_state** | [**ConfigState**](ConfigState.md) |  | [optional] 
**conflict_detection_state** | [**ConflictDetectionState**](ConflictDetectionState.md) |  | [optional] 
**notes** | **str** | Free-form notes on the device. | [optional] 
**asdm_version** | **str** | (ASAs only) Version of the ASDM device manager running on the device. | [optional] 
**asa_failover_mode** | [**AsaFailoverMode**](AsaFailoverMode.md) |  | [optional] 
**asa_failover_state** | **str** | (High Availability ASAs only) Failover state of this device. | [optional] 
**asa_failover_mate** | [**AsaFailoverMate**](AsaFailoverMate.md) |  | [optional] 
**asa_license_entitlements** | **Dict[str, str]** | (ASAs only) Map of ASA License entitlements. | [optional] 
**ftd_licenses** | **List[str]** | (FTDs only) List of FTD License entitlements. | [optional] 
**snort_version** | **str** | (FTDs only) List of FTD License entitlements. | [optional] 
**ftd_performance_tier** | **str** | (FTDvs only) The FTDv supports performance-tiered licensing that provides different throughput levels and VPN connection limits based on deployment requirements. This field specifies the performance tier of the FTD. | [optional] 
**cd_fmc_info** | [**CdFmcInfo**](CdFmcInfo.md) |  | [optional] 
**on_prem_fmc_info** | [**OnPremFmcInfo**](OnPremFmcInfo.md) |  | [optional] 
**meraki_deployment_mode** | [**MerakiDeploymentMode**](MerakiDeploymentMode.md) |  | [optional] 
**meraki_network** | [**Network**](Network.md) |  | [optional] 
**state** | **str** | The device state. | [optional] 
**state_machine_details** | [**StateMachineDetails**](StateMachineDetails.md) |  | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 

## Example

```python
from openapi_client.models.device import Device

# TODO update the JSON string below
json = "{}"
# create an instance of Device from a JSON string
device_instance = Device.from_json(json)
# print the JSON string representation of the object
print(Device.to_json())

# convert the object into a dict
device_dict = device_instance.to_dict()
# create an instance of Device from a dict
device_form_dict = device.from_dict(device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


