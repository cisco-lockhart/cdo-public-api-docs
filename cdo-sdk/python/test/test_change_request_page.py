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

from cdo_sdk_python.models.change_request_page import ChangeRequestPage

class TestChangeRequestPage(unittest.TestCase):
    """ChangeRequestPage unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ChangeRequestPage:
        """Test ChangeRequestPage
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ChangeRequestPage`
        """
        model = ChangeRequestPage()
        if include_optional:
            return ChangeRequestPage(
                count = 100,
                limit = 50,
                offset = 0,
                items = [
                    cdo_sdk_python.models.change_request.ChangeRequest(
                        uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                        name = 'my change request', 
                        description = 'my change request description', )
                    ]
            )
        else:
            return ChangeRequestPage(
        )
        """

    def testChangeRequestPage(self):
        """Test ChangeRequestPage"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
