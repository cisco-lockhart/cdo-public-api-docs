# ChangeItemDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployment_type** | **str** |  | 

## Example

```python
from scc_firewall_manager_sdk.models.change_item_dto import ChangeItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of ChangeItemDto from a JSON string
change_item_dto_instance = ChangeItemDto.from_json(json)
# print the JSON string representation of the object
print(ChangeItemDto.to_json())

# convert the object into a dict
change_item_dto_dict = change_item_dto_instance.to_dict()
# create an instance of ChangeItemDto from a dict
change_item_dto_form_dict = change_item_dto.from_dict(change_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


