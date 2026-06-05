# AccessGroupCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applied_to** | **List[str]** | The set of device unique identifiers to which this Access Group was applied. Only valid for shared access group. | [optional] 
**entity_uid** | **str** | The unique identifier, represented as a UUID, of the device/manager associated with the Access Group. When creating shared Access Group, entityUid represents device that contains source Access Group  | 
**is_shared** | **bool** | The flag that identifies if access group is shared.  If set to true, appliedTo field should be provided as well and entityUid should point to source device. | [optional] 
**name** | **str** | A human-readable name for the Access Group. | [optional] 
**resources** | **List[Dict[str, object]]** | The set of of interface and direction pairs or global resource.  Resource is an attribute applicable only to devices and will not be propagated to appliedTo devices if Access Group is shared. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.access_group_create_input import AccessGroupCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AccessGroupCreateInput from a JSON string
access_group_create_input_instance = AccessGroupCreateInput.from_json(json)
# print the JSON string representation of the object
print(AccessGroupCreateInput.to_json())

# convert the object into a dict
access_group_create_input_dict = access_group_create_input_instance.to_dict()
# create an instance of AccessGroupCreateInput from a dict
access_group_create_input_form_dict = access_group_create_input.from_dict(access_group_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


