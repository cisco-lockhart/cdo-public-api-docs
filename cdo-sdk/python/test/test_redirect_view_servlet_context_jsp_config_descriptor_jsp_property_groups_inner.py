# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.4.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.redirect_view_servlet_context_jsp_config_descriptor_jsp_property_groups_inner import RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner

class TestRedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner(unittest.TestCase):
    """RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner:
        """Test RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner`
        """
        model = RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner()
        if include_optional:
            return RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner(
                el_ignored = '',
                scripting_invalid = '',
                page_encoding = '',
                is_xml = '',
                include_preludes = [
                    ''
                    ],
                include_codas = [
                    ''
                    ],
                deferred_syntax_allowed_as_literal = '',
                trim_directive_whitespaces = '',
                error_on_undeclared_namespace = '',
                buffer = '',
                default_content_type = '',
                url_patterns = [
                    ''
                    ]
            )
        else:
            return RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner(
        )
        """

    def testRedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner(self):
        """Test RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()