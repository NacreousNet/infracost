package ibm

import "github.com/infracost/infracost/internal/schema"

var ResourceRegistry []*schema.RegistryItem = []*schema.RegistryItem{
	getIsInstanceRegistryItem(),
	getIbmIsVpcRegistryItem(),
	getIbmCosBucketRegistryItem(),
	getIsFloatingIpRegistryItem(),
	getIsFlowLogRegistryItem(),
	getContainerVpcWorkerPoolRegistryItem(),
	getContainerVpcClusterRegistryItem(),
	getResourceInstanceRegistryItem(),
	getIsVolumeRegistryItem(),
	getIsVpnGatewayRegistryItem(),
	getTgGatewayRegistryItem(),
	getCloudantRegistryItem(),
	getPiInstanceRegistryItem(),
	getIsLbRegistryItem(),
	getIsPublicGatewayRegistryItem(),
	getIbmPiVolumeRegistryItem(),
	getDatabaseRegistryItem(),
}

// FreeResources grouped alphabetically
var FreeResources = []string{
	"ibm_appid_action_url",
	"ibm_appid_apm",
	"ibm_appid_application",
	"ibm_appid_application_roles",
	"ibm_appid_application_scopes",
	"ibm_appid_audit_status",
	"ibm_appid_cloud_directory_template",
	"ibm_appid_cloud_directory_user",
	"ibm_appid_idp_cloud_directory",
	"ibm_appid_idp_custom",
	"ibm_appid_idp_facebook",
	"ibm_appid_idp_google",
	"ibm_appid_idp_saml",
	"ibm_appid_languages",
	"ibm_appid_mfa",
	"ibm_appid_mfa_channel",
	"ibm_appid_password_regex",
	"ibm_appid_redirect_urls",
	"ibm_appid_role",
	"ibm_appid_theme_color",
	"ibm_appid_theme_text",
	"ibm_appid_token_config",
	"ibm_appid_user_roles",
	"ibm_atracker_route",
	"ibm_atracker_target",
	"ibm_cbr_rule",
	"ibm_cbr_zone",
	"ibm_cd_tekton_pipeline",
	"ibm_cd_tekton_pipeline_definition",
	"ibm_cd_tekton_pipeline_property",
	"ibm_cd_tekton_pipeline_trigger",
	"ibm_cd_tekton_pipeline_trigger_property",
	"ibm_cd_toolchain",
	"ibm_cd_toolchain_tool_appconfig",
	"ibm_cd_toolchain_tool_artifactory",
	"ibm_cd_toolchain_tool_bitbucketgit",
	"ibm_cd_toolchain_tool_custom",
	"ibm_cd_toolchain_tool_devopsinsights",
	"ibm_cd_toolchain_tool_githubconsolidated",
	"ibm_cd_toolchain_tool_gitlab",
	"ibm_cd_toolchain_tool_hashicorpvault",
	"ibm_cd_toolchain_tool_hostedgit",
	"ibm_cd_toolchain_tool_jenkins",
	"ibm_cd_toolchain_tool_jira",
	"ibm_cd_toolchain_tool_keyprotect",
	"ibm_cd_toolchain_tool_nexus",
	"ibm_cd_toolchain_tool_pagerduty",
	"ibm_cd_toolchain_tool_pipeline",
	"ibm_cd_toolchain_tool_privateworker",
	"ibm_cd_toolchain_tool_saucelabs",
	"ibm_cd_toolchain_tool_secretsmanager",
	"ibm_cd_toolchain_tool_securitycompliance",
	"ibm_cd_toolchain_tool_slack",
	"ibm_cd_toolchain_tool_sonarqube",
	"ibm_cloud_shell_account_settings",
	"ibm_container_addons",
	"ibm_cos_bucket_object_lock_configuration",
	"ibm_dns_custom_resolver",
	"ibm_iam_access_group",
	"ibm_iam_access_group_account_settings",
	"ibm_iam_access_group_dynamic_rule",
	"ibm_iam_access_group_members",
	"ibm_iam_access_group_policy",
	"ibm_iam_account_settings",
	"ibm_iam_authorization_policy",
	"ibm_iam_service_api_key",
	"ibm_iam_service_id",
	"ibm_iam_service_policy",
	"ibm_iam_trusted_profile",
	"ibm_iam_trusted_profile_claim_rule",
	"ibm_iam_trusted_profile_link",
	"ibm_iam_trusted_profile_policy",
	"ibm_is_lb_listener",
	"ibm_is_lb_pool",
	"ibm_is_lb_pool_member",
	"ibm_is_network_acl",
	"ibm_is_network_acl_rule",
	"ibm_is_placement_group",
	"ibm_is_security_group",
	"ibm_is_security_group_rule",
	"ibm_is_security_group_target",
	"ibm_is_ssh_key",
	"ibm_is_subnet",
	"ibm_is_subnet_public_gateway_attachment",
	"ibm_is_subnet_reserved_ip",
	"ibm_is_virtual_endpoint_gateway",
	"ibm_is_virtual_endpoint_gateway_ip",
	"ibm_is_vpc_address_prefix",
	"ibm_is_vpc_dns_resolution_binding",
	"ibm_is_vpc_routing_table",
	"ibm_is_vpc_routing_table_route",
	"ibm_is_vpn_gateway_connection",
	"ibm_kms_instance_policies",
	"ibm_kms_key",
	"ibm_kms_key_policies",
	"ibm_kms_key_rings",
	"ibm_pi_capture",
	"ibm_pi_cloud_connection",
	"ibm_pi_cloud_connection_network_attach",
	"ibm_pi_console_language",
	"ibm_pi_dhcp",
	"ibm_pi_ike_policy",
	"ibm_pi_image",
	"ibm_pi_instance_action",
	"ibm_pi_ipsec_policy",
	"ibm_pi_key",
	"ibm_pi_network",
	"ibm_pi_network_port",
	"ibm_pi_network_port_attach",
	"ibm_pi_placement_group",
	"ibm_pi_shared_processor_pool",
	"ibm_pi_spp_placement_group",
	"ibm_pi_vpn_connection",
	"ibm_resource_group",
	"ibm_resource_key",
	"ibm_resource_tag",
	"ibm_scc_account_settings",
	"ibm_scc_control_library",
	"ibm_scc_instance_settings",
	"ibm_scc_posture_collector",
	"ibm_scc_posture_credential",
	"ibm_scc_posture_profile_import",
	"ibm_scc_posture_scan_initiate_validation",
	"ibm_scc_posture_scope",
	"ibm_scc_profile",
	"ibm_scc_profile_attachment",
	"ibm_scc_provider_type_instance",
	"ibm_scc_rule",
	"ibm_scc_rule_attachment",
	"ibm_scc_template",
	"ibm_scc_template_attachment",
	"ibm_sm_arbitrary_secret",
	"ibm_sm_en_registration",
	"ibm_sm_iam_credentials_configuration",
	"ibm_sm_imported_certificate",
	"ibm_sm_private_certificate",
	"ibm_sm_private_certificate_configuration_intermediate_ca",
	"ibm_sm_private_certificate_configuration_root_ca",
	"ibm_sm_private_certificate_configuration_template",
	"ibm_sm_public_certificate_configuration_ca_lets_encrypt",
	"ibm_sm_public_certificate_configuration_dns_cis",
	"ibm_sm_secret_group",
	"ibm_sm_service_credentials_secret",
	"ibm_sm_username_password_secret",
	"ibm_tg_connection",
}

var UsageOnlyResources = []string{
	"",
}
