# SourceDynamicObjectContent

The source dynamic object.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**elements** | **List[str]** | The list of elements. | [optional] 
**name** | **str** | The name | 
**type** | **str** | The type of the policy object. | [optional] 
**uid** | **str** | The unique identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.source_dynamic_object_content import SourceDynamicObjectContent

# TODO update the JSON string below
json = "{}"
# create an instance of SourceDynamicObjectContent from a JSON string
source_dynamic_object_content_instance = SourceDynamicObjectContent.from_json(json)
# print the JSON string representation of the object
print(SourceDynamicObjectContent.to_json())

# convert the object into a dict
source_dynamic_object_content_dict = source_dynamic_object_content_instance.to_dict()
# create an instance of SourceDynamicObjectContent from a dict
source_dynamic_object_content_form_dict = source_dynamic_object_content.from_dict(source_dynamic_object_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


