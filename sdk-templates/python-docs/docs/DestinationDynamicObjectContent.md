# DestinationDynamicObjectContent

The destination dynamic object.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**elements** | **List[str]** | The list of elements. | [optional] 
**name** | **str** | The name | 
**type** | **str** | The type of the policy object. | [optional] 
**uid** | **str** | The unique identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.destination_dynamic_object_content import DestinationDynamicObjectContent

# TODO update the JSON string below
json = "{}"
# create an instance of DestinationDynamicObjectContent from a JSON string
destination_dynamic_object_content_instance = DestinationDynamicObjectContent.from_json(json)
# print the JSON string representation of the object
print(DestinationDynamicObjectContent.to_json())

# convert the object into a dict
destination_dynamic_object_content_dict = destination_dynamic_object_content_instance.to_dict()
# create an instance of DestinationDynamicObjectContent from a dict
destination_dynamic_object_content_form_dict = destination_dynamic_object_content.from_dict(destination_dynamic_object_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


