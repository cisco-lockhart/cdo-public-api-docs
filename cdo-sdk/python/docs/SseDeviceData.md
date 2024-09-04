# SseDeviceData

SSX related info for the registered device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sse_device_id** | **str** |  | [optional] 
**sse_protocol_version** | **str** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.sse_device_data import SseDeviceData

# TODO update the JSON string below
json = "{}"
# create an instance of SseDeviceData from a JSON string
sse_device_data_instance = SseDeviceData.from_json(json)
# print the JSON string representation of the object
print(SseDeviceData.to_json())

# convert the object into a dict
sse_device_data_dict = sse_device_data_instance.to_dict()
# create an instance of SseDeviceData from a dict
sse_device_data_form_dict = sse_device_data.from_dict(sse_device_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


