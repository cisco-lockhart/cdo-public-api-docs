# DirectoryGroupPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[DirectoryGroup]**](DirectoryGroup.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.directory_group_page import DirectoryGroupPage

# TODO update the JSON string below
json = "{}"
# create an instance of DirectoryGroupPage from a JSON string
directory_group_page_instance = DirectoryGroupPage.from_json(json)
# print the JSON string representation of the object
print(DirectoryGroupPage.to_json())

# convert the object into a dict
directory_group_page_dict = directory_group_page_instance.to_dict()
# create an instance of DirectoryGroupPage from a dict
directory_group_page_form_dict = directory_group_page.from_dict(directory_group_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


