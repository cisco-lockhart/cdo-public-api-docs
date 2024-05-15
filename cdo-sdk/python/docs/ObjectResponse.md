# ObjectResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the object | [optional] 
**name** | **str** | The name of the object | [optional] 
**value** | [**SharedObjectValue**](SharedObjectValue.md) |  | [optional] 
**description** | **str** | The human-readable description of the object | [optional] 
**targets** | [**List[Target]**](Target.md) | Set of targets that contain the object. A target can be, for example, a device, service, or a shared policy (Ruleset). | [optional] 
**elements** | **List[str]** | A flattened list of the content value of the object | [optional] 
**references_info_from_default_and_overrides** | [**List[ReferenceInfo]**](ReferenceInfo.md) | List of objects referenced in the default content or in any of the overrides. | [optional] 
**tags** | **Dict[str, List[str]]** | The tags for the object | [optional] 
**labels** | **List[str]** | The labels for the object | [optional] 
**issues** | [**IssuesDto**](IssuesDto.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.object_response import ObjectResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ObjectResponse from a JSON string
object_response_instance = ObjectResponse.from_json(json)
# print the JSON string representation of the object
print(ObjectResponse.to_json())

# convert the object into a dict
object_response_dict = object_response_instance.to_dict()
# create an instance of ObjectResponse from a dict
object_response_form_dict = object_response.from_dict(object_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


