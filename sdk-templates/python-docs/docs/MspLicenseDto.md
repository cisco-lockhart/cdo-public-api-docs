# MspLicenseDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compliance_status** | **str** | Indicates whether the license is currently compliant with Smart Licensing restrictions. | [optional] 
**expiry_date** | **datetime** | (Only for TERM licenses) (Only for TERM licenses) The expiry date date, in RFC 3339 format, for this license. | [optional] 
**last_updated_time** | **datetime** | The time, in RFC 3339 format, that the license usage for this Smart Account was updated. | [optional] 
**managed_tenants** | [**List[ManagedTenantDto]**](ManagedTenantDto.md) | The list of managed tenants associated with this license. | [optional] 
**name** | **str** | The name of the license. | [optional] 
**num_available** | **int** | The number of licenses available for use. | [optional] 
**num_in_use** | **int** | The number of licenses currently in use. | [optional] 
**num_purchased** | **int** | The total number of licenses purchased. | [optional] 
**start_date** | **datetime** | (Only for TERM licenses) The activation date, in RFC 3339 format, for this license. | [optional] 
**type** | **str** | The type of the license. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the license in Security Cloud Control. | [optional] 
**virtual_account_uid** | **str** | The unique identifier, represented as a UUID, of the virtual account associated with this license. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_license_dto import MspLicenseDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspLicenseDto from a JSON string
msp_license_dto_instance = MspLicenseDto.from_json(json)
# print the JSON string representation of the object
print(MspLicenseDto.to_json())

# convert the object into a dict
msp_license_dto_dict = msp_license_dto_instance.to_dict()
# create an instance of MspLicenseDto from a dict
msp_license_dto_form_dict = msp_license_dto.from_dict(msp_license_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


