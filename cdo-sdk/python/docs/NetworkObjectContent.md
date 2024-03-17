# NetworkObjectContent

The content of a network object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**literal** | **str** | The literal content of the network object | 

## Example

```python
from cdo_sdk_python.models.network_object_content import NetworkObjectContent

# TODO update the JSON string below
json = "{}"
# create an instance of NetworkObjectContent from a JSON string
network_object_content_instance = NetworkObjectContent.from_json(json)
# print the JSON string representation of the object
print(NetworkObjectContent.to_json())

# convert the object into a dict
network_object_content_dict = network_object_content_instance.to_dict()
# create an instance of NetworkObjectContent from a dict
network_object_content_form_dict = network_object_content.from_dict(network_object_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


