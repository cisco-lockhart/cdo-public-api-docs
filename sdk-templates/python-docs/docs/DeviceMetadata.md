# DeviceMetadata


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**additional** | **Dict[str, str]** | Additional metadata as a map of key-value pairs. | [optional] 
**secure_access** | [**SecureAccessMetadata**](SecureAccessMetadata.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_metadata import DeviceMetadata

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceMetadata from a JSON string
device_metadata_instance = DeviceMetadata.from_json(json)
# print the JSON string representation of the object
print(DeviceMetadata.to_json())

# convert the object into a dict
device_metadata_dict = device_metadata_instance.to_dict()
# create an instance of DeviceMetadata from a dict
device_metadata_form_dict = device_metadata.from_dict(device_metadata_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


