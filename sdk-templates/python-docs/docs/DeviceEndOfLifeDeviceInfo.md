# DeviceEndOfLifeDeviceInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**host** | **str** | The address of the device without the port. Security Cloud Control connects to the device at this address. | [optional] 
**model_number** | **str** | The hardware, or virtualized hardware platform, that the device is running on. | [optional] 
**name** | **str** | The name of the device. Device names are unique in Security Cloud Control. | [optional] 
**software_version** | **str** | The version of the software running on the device. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_end_of_life_device_info import DeviceEndOfLifeDeviceInfo

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceEndOfLifeDeviceInfo from a JSON string
device_end_of_life_device_info_instance = DeviceEndOfLifeDeviceInfo.from_json(json)
# print the JSON string representation of the object
print(DeviceEndOfLifeDeviceInfo.to_json())

# convert the object into a dict
device_end_of_life_device_info_dict = device_end_of_life_device_info_instance.to_dict()
# create an instance of DeviceEndOfLifeDeviceInfo from a dict
device_end_of_life_device_info_form_dict = device_end_of_life_device_info.from_dict(device_end_of_life_device_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


