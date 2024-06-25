# DirectoryGroup

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the directory group in CDO. | [optional] 
**role** | [**UserRole**](UserRole.md) |  | [optional] 
**group_identifier** | **str** | The unique identifier of the directory group in your IdP. | [optional] 
**issue_url** | **str** | The URL for reporting issues related to the directory group. | [optional] 
**notes** | **str** | Free-form notes on the directory group. | [optional] 

## Example

```python
from cdo_sdk_python.models.directory_group import DirectoryGroup

# TODO update the JSON string below
json = "{}"
# create an instance of DirectoryGroup from a JSON string
directory_group_instance = DirectoryGroup.from_json(json)
# print the JSON string representation of the object
print(DirectoryGroup.to_json())

# convert the object into a dict
directory_group_dict = directory_group_instance.to_dict()
# create an instance of DirectoryGroup from a dict
directory_group_form_dict = directory_group.from_dict(directory_group_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


