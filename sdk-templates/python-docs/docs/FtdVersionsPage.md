# FtdVersionsPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**items** | [**List[FtdVersion]**](FtdVersion.md) | The list of items retrieved. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 

## Example

```python
from cdo_sdk_python.models.ftd_versions_page import FtdVersionsPage

# TODO update the JSON string below
json = "{}"
# create an instance of FtdVersionsPage from a JSON string
ftd_versions_page_instance = FtdVersionsPage.from_json(json)
# print the JSON string representation of the object
print(FtdVersionsPage.to_json())

# convert the object into a dict
ftd_versions_page_dict = ftd_versions_page_instance.to_dict()
# create an instance of FtdVersionsPage from a dict
ftd_versions_page_from_dict = FtdVersionsPage.from_dict(ftd_versions_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


