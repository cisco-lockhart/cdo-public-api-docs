# FtdChangeItemDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**added_references** | [**List[Reference]**](Reference.md) | The list of references added to this entity as part of the change. | [optional] 
**added_values** | [**List[Value]**](Value.md) | The list of values added to this entity as part of the change. | [optional] 
**change_item_type** | **str** | Indicates if this item represents an actual entity change or a grouping header. | [optional] 
**change_type** | **str** | The type of the change. | [optional] 
**children** | [**List[FtdChangeItemDto]**](FtdChangeItemDto.md) | The children of this change item. This represents changes that rest within the group represented by this change item. | [optional] 
**deleted_references** | [**List[Reference]**](Reference.md) | The list of references deleted from this entity as part of the change. | [optional] 
**deleted_values** | [**List[Value]**](Value.md) | The list of values removed from this entity as part of the change. | [optional] 
**deployment_type** | **str** | Discriminator identifying this change item as FTD. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device the changed entity is associated with. | [optional] 
**entity_type** | **str** | The type of the changed entity. | [optional] 
**entity_uid** | **str** | The unique identifier, represented as a UUID, of the changed entity. | [optional] 
**error_msg** | **str** | An error message indicating why a diff could not be generated for the changed entity. | [optional] 
**last_updated_by_users** | **List[str]** | The list of users who have made the changes to this entity. | [optional] 
**name** | **str** | The name of the changed entity. | [optional] 
**parent_uid** | **str** | The unique identifier, represented as a UUID, of the &#39;parent&#39; of this change item. The parentUid, if set, can be used to identify the group that this change belongs to. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of this change item. Note: this is not the unique identifier of the changed entity. | [optional] 
**updated_values** | [**List[UpdatedValue]**](UpdatedValue.md) | The list of values updated on this entity as part of the change. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_change_item_dto import FtdChangeItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of FtdChangeItemDto from a JSON string
ftd_change_item_dto_instance = FtdChangeItemDto.from_json(json)
# print the JSON string representation of the object
print(FtdChangeItemDto.to_json())

# convert the object into a dict
ftd_change_item_dto_dict = ftd_change_item_dto_instance.to_dict()
# create an instance of FtdChangeItemDto from a dict
ftd_change_item_dto_form_dict = ftd_change_item_dto.from_dict(ftd_change_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


