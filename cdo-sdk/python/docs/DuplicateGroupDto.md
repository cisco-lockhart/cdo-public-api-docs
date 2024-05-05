# DuplicateGroupDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_uid** | **str** | The unique identifier of the target to search in. | [optional] 
**unified_object_uids** | **List[str]** | The set of unique identifiers of the duplicate objects. | [optional] 

## Example

```python
from cdo_sdk_python.models.duplicate_group_dto import DuplicateGroupDto

# TODO update the JSON string below
json = "{}"
# create an instance of DuplicateGroupDto from a JSON string
duplicate_group_dto_instance = DuplicateGroupDto.from_json(json)
# print the JSON string representation of the object
print(DuplicateGroupDto.to_json())

# convert the object into a dict
duplicate_group_dto_dict = duplicate_group_dto_instance.to_dict()
# create an instance of DuplicateGroupDto from a dict
duplicate_group_dto_form_dict = duplicate_group_dto.from_dict(duplicate_group_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


