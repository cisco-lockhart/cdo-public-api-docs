# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.1.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.directory_group_page import DirectoryGroupPage

class TestDirectoryGroupPage(unittest.TestCase):
    """DirectoryGroupPage unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> DirectoryGroupPage:
        """Test DirectoryGroupPage
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `DirectoryGroupPage`
        """
        model = DirectoryGroupPage()
        if include_optional:
            return DirectoryGroupPage(
                count = 100,
                limit = 50,
                offset = 0,
                items = [
                    cdo_sdk_python.models.directory_group.DirectoryGroup(
                        name = 'myusername', 
                        role = 'ROLE_ADMIN', 
                        group_identifier = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                        issue_url = 'https://issue.com', 
                        notes = 'This is an example note on the directory group.', )
                    ]
            )
        else:
            return DirectoryGroupPage(
        )
        """

    def testDirectoryGroupPage(self):
        """Test DirectoryGroupPage"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()