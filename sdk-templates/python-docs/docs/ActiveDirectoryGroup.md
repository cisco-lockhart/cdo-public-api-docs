# ActiveDirectoryGroup


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**group_identifier** | **str** | The unique identifier of the user group in your Identity Provider (IdP). | [optional] 
**issuer_url** | **str** | The Identity Provider (IdP) URL, which Cisco Defense Orchestrator will use to validate SAML assertions during the sign-in process. | [optional] 
**name** | **str** | The name of the user group. Security Cloud Control does not support special characters for this field. | [optional] 
**notes** | **str** | Any notes that are applicable to this user group. | [optional] 
**role** | [**UserRole**](UserRole.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the Active Directory Group in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.active_directory_group import ActiveDirectoryGroup

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


