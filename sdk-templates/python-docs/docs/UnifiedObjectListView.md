# UnifiedObjectListView


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** |  | [optional] 
**issues** | [**IssuesDto**](IssuesDto.md) |  | [optional] 
**labels** | **List[str]** |  | [optional] 
**name** | **str** |  | [optional] 
**override_ids** | **List[str]** |  | [optional] 
**read_only** | **bool** |  | [optional] 
**tags** | **Dict[str, List[str]]** |  | [optional] 
**target_ids** | **List[str]** |  | [optional] 
**uid** | **str** |  | [optional] 
**value** | [**SharedObjectValue**](SharedObjectValue.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.unified_object_list_view import UnifiedObjectListView

# TODO update the JSON string below
json = "{}"
# create an instance of UnifiedObjectListView from a JSON string
unified_object_list_view_instance = UnifiedObjectListView.from_json(json)
# print the JSON string representation of the object
print(UnifiedObjectListView.to_json())

# convert the object into a dict
unified_object_list_view_dict = unified_object_list_view_instance.to_dict()
# create an instance of UnifiedObjectListView from a dict
unified_object_list_view_form_dict = unified_object_list_view.from_dict(unified_object_list_view_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


