# IssuesCount


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**total_objects** | **int** | Total number of objects for the tenant | [optional] 
**un_associated** | **int** | Number of UnAssociated Objects | [optional] 
**un_useds** | **int** | Number of unAssigned Objects | [optional] 
**duplicates** | **int** | Number of duplicate Objects | [optional] 

## Example

```python
from cdo_sdk_python.models.issues_count import IssuesCount

# TODO update the JSON string below
json = "{}"
# create an instance of IssuesCount from a JSON string
issues_count_instance = IssuesCount.from_json(json)
# print the JSON string representation of the object
print(IssuesCount.to_json())

# convert the object into a dict
issues_count_dict = issues_count_instance.to_dict()
# create an instance of IssuesCount from a dict
issues_count_form_dict = issues_count.from_dict(issues_count_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


