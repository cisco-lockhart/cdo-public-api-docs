# GroupContent

The contents of an object group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**literals** | [**List[SingleContent]**](SingleContent.md) | List of content literals | [optional] 
**referenced_object_uids** | **List[str]** | Set of UIDs of the group&#39;s referenced objects | [optional] 

## Example

```python
from cdo_python_sdk.models.group_content import GroupContent

# TODO update the JSON string below
json = "{}"
# create an instance of GroupContent from a JSON string
group_content_instance = GroupContent.from_json(json)
# print the JSON string representation of the object
print(GroupContent.to_json())

# convert the object into a dict
group_content_dict = group_content_instance.to_dict()
# create an instance of GroupContent from a dict
group_content_form_dict = group_content.from_dict(group_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


