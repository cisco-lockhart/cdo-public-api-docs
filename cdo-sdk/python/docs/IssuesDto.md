# IssuesDto

The issues for the object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**unused_target_ids** | **List[str]** |  | [optional] 

## Example

```python
from cdo-python-sdk.models.issues_dto import IssuesDto

# TODO update the JSON string below
json = "{}"
# create an instance of IssuesDto from a JSON string
issues_dto_instance = IssuesDto.from_json(json)
# print the JSON string representation of the object
print(IssuesDto.to_json())

# convert the object into a dict
issues_dto_dict = issues_dto_instance.to_dict()
# create an instance of IssuesDto from a dict
issues_dto_form_dict = issues_dto.from_dict(issues_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


