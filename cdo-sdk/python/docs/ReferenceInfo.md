# ReferenceInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the referenced object | [optional] 
**name** | **str** | The name of the referenced object | [optional] 
**object_type** | **str** | The object type | [optional] 

## Example

```python
from openapi_client.models.reference_info import ReferenceInfo

# TODO update the JSON string below
json = "{}"
# create an instance of ReferenceInfo from a JSON string
reference_info_instance = ReferenceInfo.from_json(json)
# print the JSON string representation of the object
print(ReferenceInfo.to_json())

# convert the object into a dict
reference_info_dict = reference_info_instance.to_dict()
# create an instance of ReferenceInfo from a dict
reference_info_form_dict = reference_info.from_dict(reference_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


