# RestorableDeviceDto

Every device in the deployment run, with indicators as to whether the device can be restored to the configuration deployed in this deployment run.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_name** | **str** | The name of the device. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 
**not_restorable_reason** | **str** | Why the device cannot be restored. Absent when the device is restorable. DEVICE_NOT_FOUND: the device no longer exists. NOT_ROLLBACK_ELIGIBLE: cdFMC cannot roll the device back to this deployment run. DEVICE_UNREACHABLE: the device is not online. DEPLOYMENT_IN_PROGRESS: the device is being deployed to. | [optional] 
**restorable** | **bool** | Whether the device can be restored to the configuration deployed in this deployment run. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.restorable_device_dto import RestorableDeviceDto

# TODO update the JSON string below
json = "{}"
# create an instance of RestorableDeviceDto from a JSON string
restorable_device_dto_instance = RestorableDeviceDto.from_json(json)
# print the JSON string representation of the object
print(RestorableDeviceDto.to_json())

# convert the object into a dict
restorable_device_dto_dict = restorable_device_dto_instance.to_dict()
# create an instance of RestorableDeviceDto from a dict
restorable_device_dto_form_dict = restorable_device_dto.from_dict(restorable_device_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


