# DeviceEndOfLifeReport


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_end_of_life** | [**DeviceEndOfLife**](DeviceEndOfLife.md) |  | [optional] 
**devices** | [**List[DeviceEndOfLifeDeviceInfo]**](DeviceEndOfLifeDeviceInfo.md) | An inventory of End-Of-Life affected devices. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_end_of_life_report import DeviceEndOfLifeReport

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceEndOfLifeReport from a JSON string
device_end_of_life_report_instance = DeviceEndOfLifeReport.from_json(json)
# print the JSON string representation of the object
print(DeviceEndOfLifeReport.to_json())

# convert the object into a dict
device_end_of_life_report_dict = device_end_of_life_report_instance.to_dict()
# create an instance of DeviceEndOfLifeReport from a dict
device_end_of_life_report_form_dict = device_end_of_life_report.from_dict(device_end_of_life_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


