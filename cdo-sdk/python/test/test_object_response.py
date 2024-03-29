# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.object_response import ObjectResponse

class TestObjectResponse(unittest.TestCase):
    """ObjectResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ObjectResponse:
        """Test ObjectResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ObjectResponse`
        """
        model = ObjectResponse()
        if include_optional:
            return ObjectResponse(
                uid = '',
                name = 'my-object',
                value = cdo_sdk_python.models.shared_object_value.SharedObjectValue(
                    object_type = 'NETWORK_OBJECT', 
                    default_content = {"literal":"a:b:c::1"}, 
                    overrides = [
                        {"targetId":"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx","content":"a:b:c::2"}
                        ], ),
                description = 'My object description',
                targets = [
                    cdo_sdk_python.models.target.Target(
                        id = '[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx]', 
                        display_name = 'My ASA', 
                        type = 'ASA', )
                    ],
                elements = [a:b:c::1],
                references_info_from_default_and_overrides = [
                    cdo_sdk_python.models.reference_info.ReferenceInfo(
                        uid = '', 
                        name = 'another-object', 
                        object_type = 'NETWORK_OBJECT', )
                    ],
                tags = {"location":["London","Head-office"]},
                labels = ["migration"],
                issues = {"unusedTargetIds":["xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"]}
            )
        else:
            return ObjectResponse(
        )
        """

    def testObjectResponse(self):
        """Test ObjectResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
