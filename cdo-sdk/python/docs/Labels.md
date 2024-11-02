# Labels

Labels used to identify/tag SCC entities.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**grouped_labels** | **Dict[str, List[str]]** | Groups of labels used to identify/tag SCC entities. | [optional] 
**ungrouped_labels** | **List[str]** | Set of free-labels used to identify/tag SCC entities. | [optional] 

## Example

```python
from cdo_sdk_python.models.labels import Labels

# TODO update the JSON string below
json = "{}"
# create an instance of Labels from a JSON string
labels_instance = Labels.from_json(json)
# print the JSON string representation of the object
print(Labels.to_json())

# convert the object into a dict
labels_dict = labels_instance.to_dict()
# create an instance of Labels from a dict
labels_form_dict = labels.from_dict(labels_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


