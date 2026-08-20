# CdoTokenInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enterprise_id** | **str** | The unique identifier, represented as a UUID, of the Security Cloud Control enterprise this token is associated with. | [optional] 
**expires_at** | **datetime** | The time (UTC; represented using the RFC-3339 standard) the token expires. If this field is missing, the token will never expire. | [optional] 
**name** | **str** | The name of the user this token belongs to. The user can be API-only or a human. | [optional] 
**product_instance_id** | **str** | The unique identifier, represented as a UUID, of the firewall product instance this token is associated with. A Security Cloud Control enterprise has subscriptions to multiple products; this is the unique identifier of the firewall product instance. | [optional] 
**roles** | [**List[UserRole]**](UserRole.md) | The role of the user this token belongs to. The user can be API-only or a human. | [optional] 
**tenant_name** | **str** | The name of the enterprise (tenant) in Security Cloud Control. | [optional] 
**tenant_pay_type** | **str** | The pay type of the tenant this token is associated with. | [optional] 
**tenant_uid** | **str** | The unique identifier, represented as a UUID, of the tenant this token is associated with. While a user can be associated with multiple tenants, a token is associated with a single tenant. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the user this token belongs to. The user can be API-only or a human. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.cdo_token_info import CdoTokenInfo

# TODO update the JSON string below
json = "{}"
# create an instance of CdoTokenInfo from a JSON string
cdo_token_info_instance = CdoTokenInfo.from_json(json)
# print the JSON string representation of the object
print(CdoTokenInfo.to_json())

# convert the object into a dict
cdo_token_info_dict = cdo_token_info_instance.to_dict()
# create an instance of CdoTokenInfo from a dict
cdo_token_info_form_dict = cdo_token_info.from_dict(cdo_token_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


