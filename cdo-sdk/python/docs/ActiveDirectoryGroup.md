# ActiveDirectoryGroup

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the Active Directory Group. CDO does not support special characters for this field. | [optional] 
**role** | [**UserRole**](UserRole.md) |  | [optional] 
**group_identifier** | **str** | The unique identifier of the Active Directory Group in your Identity Provider (IdP). | [optional] 
**issuer_url** | **str** | The Identity Provider (IdP) URL, which Cisco Defense Orchestrator will use to validate SAML assertions during the sign-in process. | [optional] 
**notes** | **str** | Any notes that are applicable to this Active Directory Group. | [optional] 

## Example

```python
from cdo_sdk_python.models.active_directory_group import ActiveDirectoryGroup

# TODO update the JSON string below
json = "{}"
# create an instance of ActiveDirectoryGroup from a JSON string
active_directory_group_instance = ActiveDirectoryGroup.from_json(json)
# print the JSON string representation of the object
print(ActiveDirectoryGroup.to_json())

# convert the object into a dict
active_directory_group_dict = active_directory_group_instance.to_dict()
# create an instance of ActiveDirectoryGroup from a dict
active_directory_group_form_dict = active_directory_group.from_dict(active_directory_group_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


