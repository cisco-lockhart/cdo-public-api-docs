# FtdChassisInfo

(Chassis devices managed by FMC only) Information on logical devices deployed on this chassis.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**logical_devices** | [**List[LogicalDevice]**](LogicalDevice.md) | The logical devices on this chassis that are managed on the same Security Cloud Control tenant as the chassis. A chassis can host logical devices managed on other tenants; those are not listed. | [optional] 
**num_logical_devices** | **int** | The number of logical devices on this chassis that are managed on the same Security Cloud Control tenant as the chassis. Always equal to the length of &#x60;logicalDevices&#x60;. Logical devices on the chassis that are managed on other tenants are not counted. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_chassis_info import FtdChassisInfo

# TODO update the JSON string below
json = "{}"
# create an instance of FtdChassisInfo from a JSON string
ftd_chassis_info_instance = FtdChassisInfo.from_json(json)
# print the JSON string representation of the object
print(FtdChassisInfo.to_json())

# convert the object into a dict
ftd_chassis_info_dict = ftd_chassis_info_instance.to_dict()
# create an instance of FtdChassisInfo from a dict
ftd_chassis_info_form_dict = ftd_chassis_info.from_dict(ftd_chassis_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


