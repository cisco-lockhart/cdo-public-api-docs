# DeploymentRunTranscriptDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployment_run_uid** | **str** | The unique identifier, represented as a UUID, of the deployment run the transcript is associated with. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device the transcript is associated with. | [optional] 
**transcript** | **str** | The transcript of the deployment run. Note: The transcript is returned as a string with &#x60; &#x60; as a line separator. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the deployment run transcript in SCC Firewall Manager. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.deployment_run_transcript_dto import DeploymentRunTranscriptDto

# TODO update the JSON string below
json = "{}"
# create an instance of DeploymentRunTranscriptDto from a JSON string
deployment_run_transcript_dto_instance = DeploymentRunTranscriptDto.from_json(json)
# print the JSON string representation of the object
print(DeploymentRunTranscriptDto.to_json())

# convert the object into a dict
deployment_run_transcript_dto_dict = deployment_run_transcript_dto_instance.to_dict()
# create an instance of DeploymentRunTranscriptDto from a dict
deployment_run_transcript_dto_form_dict = deployment_run_transcript_dto.from_dict(deployment_run_transcript_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


