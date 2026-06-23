# ManagedTenantUpgradeDto

Per-managed-tenant upgrade progress across this run.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error_msg** | **str** | Error message if the managed tenant&#39;s upgrade failed. | [optional] 
**managed_tenant_display_name** | **str** | Display name of the managed tenant in Security Cloud Control. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control. | [optional] 
**transaction_status** | **str** | The status of the managed tenant&#39;s upgrade transaction. One of: PENDING, IN_PROGRESS, CANCELLED, DONE, ERROR. | [optional] 
**transaction_uid** | **str** | The unique identifier, represented as a UUID, of the upgrade transaction on the managed tenant. Null if the upgrade could not be triggered on that tenant. | [optional] 
**upgrade_run_status** | **str** | The current status of the upgrade run on the managed tenant, as reported by the tenant. A granular pipeline step (e.g. UPLOADING_IMAGES, REGISTERING_NEW_RA_VPN_IMAGES, DONE); passed through verbatim from the managed tenant. | [optional] 
**upgrade_run_step_percentage_complete** | **float** | Percentage completion (0.0 to 1.0) of the managed tenant upgrade run&#39;s current step. Null when the tenant reports no active progress for the current step. | [optional] 
**upgrade_run_uid** | **str** | The unique identifier, represented as a UUID, of the upgrade run on the managed tenant. Use it to deep-link into the tenant&#39;s own upgrade-run detail. Null if the upgrade could not be triggered on that tenant. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.managed_tenant_upgrade_dto import ManagedTenantUpgradeDto

# TODO update the JSON string below
json = "{}"
# create an instance of ManagedTenantUpgradeDto from a JSON string
managed_tenant_upgrade_dto_instance = ManagedTenantUpgradeDto.from_json(json)
# print the JSON string representation of the object
print(ManagedTenantUpgradeDto.to_json())

# convert the object into a dict
managed_tenant_upgrade_dto_dict = managed_tenant_upgrade_dto_instance.to_dict()
# create an instance of ManagedTenantUpgradeDto from a dict
managed_tenant_upgrade_dto_form_dict = managed_tenant_upgrade_dto.from_dict(managed_tenant_upgrade_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


