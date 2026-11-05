# RestorableDeviceListDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | [**List[RestorableDeviceDto]**](RestorableDeviceDto.md) | Every device in the deployment run, with indicators as to whether the device can be restored to the configuration deployed in this deployment run. | 

## Example

```python
from scc_firewall_manager_sdk.models.restorable_device_list_dto import RestorableDeviceListDto

# TODO update the JSON string below
json = "{}"
# create an instance of RestorableDeviceListDto from a JSON string
restorable_device_list_dto_instance = RestorableDeviceListDto.from_json(json)
# print the JSON string representation of the object
print(RestorableDeviceListDto.to_json())

# convert the object into a dict
restorable_device_list_dto_dict = restorable_device_list_dto_instance.to_dict()
# create an instance of RestorableDeviceListDto from a dict
restorable_device_list_dto_form_dict = restorable_device_list_dto.from_dict(restorable_device_list_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


