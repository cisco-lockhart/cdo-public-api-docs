# AccessGroup


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of Access Group in Security Cloud Control. | 
**name** | **str** | The name of Access Group in CDO. Access Group names are unique in Security Cloud Control. | 
**entity_uid** | **str** | The unique identifier, represented as a UUID, of the device/manager associated with the Access Group. | 
**shared_access_group_uid** | **str** | The unique identifier, represented as a UUID, of the shared access group manager associated with the Access Group. | 
**is_shared** | **bool** | The flag that identifies if access group is shared. | [optional] 
**applied_to** | **List[str]** | The set of device unique identifiers to which this Access Group was applied. Only valid for shared access group. | [optional] 
**resources** | **List[Dict[str, object]]** | The set of of interface and direction pairs or global resource. | [optional] 
**created_date** | **datetime** | The time (in UTC) at which Access Group was created, represented using the RFC-3339 standard. | [optional] 
**updated_date** | **datetime** | The time (in UTC) at which Access Group was updated, represented using the RFC-3339 standard. | [optional] 

## Example

```python
from cdo_sdk_python.models.access_group import AccessGroup

# TODO update the JSON string below
json = "{}"
# create an instance of AccessGroup from a JSON string
access_group_instance = AccessGroup.from_json(json)
# print the JSON string representation of the object
print(AccessGroup.to_json())

# convert the object into a dict
access_group_dict = access_group_instance.to_dict()
# create an instance of AccessGroup from a dict
access_group_form_dict = access_group.from_dict(access_group_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


