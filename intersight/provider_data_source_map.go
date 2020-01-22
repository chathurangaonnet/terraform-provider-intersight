package intersight

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func GetDataSourceMapping() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"intersight_equipment_locator_led":                               dataSourceEquipmentLocatorLed(),
		"intersight_iam_qualifier":                                       dataSourceIamQualifier(),
		"intersight_storage_pure_protection_group":                       dataSourceStoragePureProtectionGroup(),
		"intersight_hyperflex_health":                                    dataSourceHyperflexHealth(),
		"intersight_memory_persistent_memory_region":                     dataSourceMemoryPersistentMemoryRegion(),
		"intersight_search_tag_item":                                     dataSourceSearchTagItem(),
		"intersight_asset_device_contract_information":                   dataSourceAssetDeviceContractInformation(),
		"intersight_bios_policy":                                         dataSourceBiosPolicy(),
		"intersight_workflow_task_definition":                            dataSourceWorkflowTaskDefinition(),
		"intersight_boot_precision_policy":                               dataSourceBootPrecisionPolicy(),
		"intersight_storage_controller":                                  dataSourceStorageController(),
		"intersight_hyperflex_feature_limit_internal":                    dataSourceHyperflexFeatureLimitInternal(),
		"intersight_terminal_audit_log":                                  dataSourceTerminalAuditLog(),
		"intersight_syslog_policy":                                       dataSourceSyslogPolicy(),
		"intersight_iaas_connector_pack":                                 dataSourceIaasConnectorPack(),
		"intersight_search_search_item":                                  dataSourceSearchSearchItem(),
		"intersight_iam_idp":                                             dataSourceIamIdp(),
		"intersight_deviceconnector_policy":                              dataSourceDeviceconnectorPolicy(),
		"intersight_hyperflex_server_model":                              dataSourceHyperflexServerModel(),
		"intersight_iam_service_provider":                                dataSourceIamServiceProvider(),
		"intersight_storage_physical_disk_extension":                     dataSourceStoragePhysicalDiskExtension(),
		"intersight_iam_resource_limits":                                 dataSourceIamResourceLimits(),
		"intersight_os_install":                                          dataSourceOsInstall(),
		"intersight_appliance_backup":                                    dataSourceApplianceBackup(),
		"intersight_appliance_upgrade":                                   dataSourceApplianceUpgrade(),
		"intersight_workflow_task_meta":                                  dataSourceWorkflowTaskMeta(),
		"intersight_inventory_generic_inventory_holder":                  dataSourceInventoryGenericInventoryHolder(),
		"intersight_storage_enclosure_disk_slot_ep":                      dataSourceStorageEnclosureDiskSlotEp(),
		"intersight_iam_trust_point":                                     dataSourceIamTrustPoint(),
		"intersight_iaas_license_info":                                   dataSourceIaasLicenseInfo(),
		"intersight_storage_enclosure_disk":                              dataSourceStorageEnclosureDisk(),
		"intersight_niaapi_apic_cco_post":                                dataSourceNiaapiApicCcoPost(),
		"intersight_asset_device_configuration":                          dataSourceAssetDeviceConfiguration(),
		"intersight_hyperflex_local_credential_policy":                   dataSourceHyperflexLocalCredentialPolicy(),
		"intersight_workflow_workflow_definition":                        dataSourceWorkflowWorkflowDefinition(),
		"intersight_workflow_catalog":                                    dataSourceWorkflowCatalog(),
		"intersight_tam_security_advisory":                               dataSourceTamSecurityAdvisory(),
		"intersight_sdcard_policy":                                       dataSourceSdcardPolicy(),
		"intersight_cond_alarm":                                          dataSourceCondAlarm(),
		"intersight_resource_license_resource_count":                     dataSourceResourceLicenseResourceCount(),
		"intersight_hyperflex_ext_iscsi_storage_policy":                  dataSourceHyperflexExtIscsiStoragePolicy(),
		"intersight_storage_pure_replication_schedule":                   dataSourceStoragePureReplicationSchedule(),
		"intersight_workflow_build_task_meta_owner":                      dataSourceWorkflowBuildTaskMetaOwner(),
		"intersight_compute_rack_unit":                                   dataSourceComputeRackUnit(),
		"intersight_storage_pure_port":                                   dataSourceStoragePurePort(),
		"intersight_hyperflex_node_profile":                              dataSourceHyperflexNodeProfile(),
		"intersight_port_sub_group":                                      dataSourcePortSubGroup(),
		"intersight_memory_array":                                        dataSourceMemoryArray(),
		"intersight_appliance_system_info":                               dataSourceApplianceSystemInfo(),
		"intersight_hyperflex_proxy_setting_policy":                      dataSourceHyperflexProxySettingPolicy(),
		"intersight_license_smartlicense_token":                          dataSourceLicenseSmartlicenseToken(),
		"intersight_storage_flex_flash_controller_props":                 dataSourceStorageFlexFlashControllerProps(),
		"intersight_os_configuration_file":                               dataSourceOsConfigurationFile(),
		"intersight_hyperflex_ext_fc_storage_policy":                     dataSourceHyperflexExtFcStoragePolicy(),
		"intersight_boot_device_boot_mode":                               dataSourceBootDeviceBootMode(),
		"intersight_workflow_build_task_meta":                            dataSourceWorkflowBuildTaskMeta(),
		"intersight_softwarerepository_catalog":                          dataSourceSoftwarerepositoryCatalog(),
		"intersight_storage_vd_member_ep":                                dataSourceStorageVdMemberEp(),
		"intersight_storage_pure_array":                                  dataSourceStoragePureArray(),
		"intersight_hcl_hyperflex_software_compatibility_info":           dataSourceHclHyperflexSoftwareCompatibilityInfo(),
		"intersight_pci_switch":                                          dataSourcePciSwitch(),
		"intersight_memory_persistent_memory_config_result":              dataSourceMemoryPersistentMemoryConfigResult(),
		"intersight_niaapi_dcnm_cco_post":                                dataSourceNiaapiDcnmCcoPost(),
		"intersight_workflow_workflow_meta":                              dataSourceWorkflowWorkflowMeta(),
		"intersight_network_element":                                     dataSourceNetworkElement(),
		"intersight_fault_instance":                                      dataSourceFaultInstance(),
		"intersight_storage_flex_util_controller":                        dataSourceStorageFlexUtilController(),
		"intersight_hyperflex_node_config_policy":                        dataSourceHyperflexNodeConfigPolicy(),
		"intersight_iam_ldap_policy":                                     dataSourceIamLdapPolicy(),
		"intersight_firmware_upgrade_status":                             dataSourceFirmwareUpgradeStatus(),
		"intersight_networkconfig_policy":                                dataSourceNetworkconfigPolicy(),
		"intersight_sol_policy":                                          dataSourceSolPolicy(),
		"intersight_externalsite_authorization":                          dataSourceExternalsiteAuthorization(),
		"intersight_compute_blade":                                       dataSourceComputeBlade(),
		"intersight_iam_ldap_group":                                      dataSourceIamLdapGroup(),
		"intersight_forecast_catalog":                                    dataSourceForecastCatalog(),
		"intersight_hyperflex_feature_limit_external":                    dataSourceHyperflexFeatureLimitExternal(),
		"intersight_appliance_restore":                                   dataSourceApplianceRestore(),
		"intersight_workflow_custom_data_type_definition":                dataSourceWorkflowCustomDataTypeDefinition(),
		"intersight_storage_enclosure":                                   dataSourceStorageEnclosure(),
		"intersight_iam_user_preference":                                 dataSourceIamUserPreference(),
		"intersight_equipment_io_expander":                               dataSourceEquipmentIoExpander(),
		"intersight_graphics_controller":                                 dataSourceGraphicsController(),
		"intersight_appliance_image_bundle":                              dataSourceApplianceImageBundle(),
		"intersight_equipment_fex":                                       dataSourceEquipmentFex(),
		"intersight_hcl_exempted_catalog":                                dataSourceHclExemptedCatalog(),
		"intersight_storage_storage_policy":                              dataSourceStorageStoragePolicy(),
		"intersight_software_hcl_meta":                                   dataSourceSoftwareHclMeta(),
		"intersight_niaapi_apic_sweol":                                   dataSourceNiaapiApicSweol(),
		"intersight_niatelemetry_nia_license_state":                      dataSourceNiatelemetryNiaLicenseState(),
		"intersight_compute_server_setting":                              dataSourceComputeServerSetting(),
		"intersight_iam_certificate":                                     dataSourceIamCertificate(),
		"intersight_ipmioverlan_policy":                                  dataSourceIpmioverlanPolicy(),
		"intersight_storage_virtual_drive":                               dataSourceStorageVirtualDrive(),
		"intersight_vnic_eth_qos_policy":                                 dataSourceVnicEthQosPolicy(),
		"intersight_vnic_fc_qos_policy":                                  dataSourceVnicFcQosPolicy(),
		"intersight_firmware_server_configuration_utility_distributable": dataSourceFirmwareServerConfigurationUtilityDistributable(),
		"intersight_hcl_driver_image":                                    dataSourceHclDriverImage(),
		"intersight_pci_coprocessor_card":                                dataSourcePciCoprocessorCard(),
		"intersight_compute_board":                                       dataSourceComputeBoard(),
		"intersight_vnic_eth_network_policy":                             dataSourceVnicEthNetworkPolicy(),
		"intersight_iam_user_group":                                      dataSourceIamUserGroup(),
		"intersight_iam_api_key":                                         dataSourceIamApiKey(),
		"intersight_aaa_audit_record":                                    dataSourceAaaAuditRecord(),
		"intersight_firmware_running_firmware":                           dataSourceFirmwareRunningFirmware(),
		"intersight_iam_end_point_user":                                  dataSourceIamEndPointUser(),
		"intersight_server_config_import":                                dataSourceServerConfigImport(),
		"intersight_iam_security_holder":                                 dataSourceIamSecurityHolder(),
		"intersight_ucsd_backup_info":                                    dataSourceUcsdBackupInfo(),
		"intersight_hyperflex_capability_info":                           dataSourceHyperflexCapabilityInfo(),
		"intersight_niaapi_dcnm_hweol":                                   dataSourceNiaapiDcnmHweol(),
		"intersight_storage_pure_volume_snapshot":                        dataSourceStoragePureVolumeSnapshot(),
		"intersight_iam_resource_roles":                                  dataSourceIamResourceRoles(),
		"intersight_adapter_config_policy":                               dataSourceAdapterConfigPolicy(),
		"intersight_organization_organization":                           dataSourceOrganizationOrganization(),
		"intersight_resource_group":                                      dataSourceResourceGroup(),
		"intersight_equipment_fan":                                       dataSourceEquipmentFan(),
		"intersight_storage_flex_flash_virtual_drive":                    dataSourceStorageFlexFlashVirtualDrive(),
		"intersight_firmware_driver_distributable":                       dataSourceFirmwareDriverDistributable(),
		"intersight_bios_boot_mode":                                      dataSourceBiosBootMode(),
		"intersight_iaas_ucsd_info":                                      dataSourceIaasUcsdInfo(),
		"intersight_storage_flex_flash_controller":                       dataSourceStorageFlexFlashController(),
		"intersight_iam_idp_reference":                                   dataSourceIamIdpReference(),
		"intersight_appliance_setup_info":                                dataSourceApplianceSetupInfo(),
		"intersight_adapter_host_fc_interface":                           dataSourceAdapterHostFcInterface(),
		"intersight_vnic_fc_if":                                          dataSourceVnicFcIf(),
		"intersight_niaapi_dcnm_release_recommend":                       dataSourceNiaapiDcnmReleaseRecommend(),
		"intersight_iam_end_point_privilege":                             dataSourceIamEndPointPrivilege(),
		"intersight_compute_physical_summary":                            dataSourceComputePhysicalSummary(),
		"intersight_recovery_backup_profile":                             dataSourceRecoveryBackupProfile(),
		"intersight_niatelemetry_nia_inventory":                          dataSourceNiatelemetryNiaInventory(),
		"intersight_firmware_distributable":                              dataSourceFirmwareDistributable(),
		"intersight_resource_membership":                                 dataSourceResourceMembership(),
		"intersight_management_interface":                                dataSourceManagementInterface(),
		"intersight_niaapi_dcnm_latest_maintained_release":               dataSourceNiaapiDcnmLatestMaintainedRelease(),
		"intersight_server_config_result":                                dataSourceServerConfigResult(),
		"intersight_tam_advisory_info":                                   dataSourceTamAdvisoryInfo(),
		"intersight_forecast_definition":                                 dataSourceForecastDefinition(),
		"intersight_recovery_config_result_entry":                        dataSourceRecoveryConfigResultEntry(),
		"intersight_iaas_ucsd_managed_infra":                             dataSourceIaasUcsdManagedInfra(),
		"intersight_equipment_shared_io_module":                          dataSourceEquipmentSharedIoModule(),
		"intersight_hyperflex_ucsm_config_policy":                        dataSourceHyperflexUcsmConfigPolicy(),
		"intersight_smtp_policy":                                         dataSourceSmtpPolicy(),
		"intersight_asset_device_registration":                           dataSourceAssetDeviceRegistration(),
		"intersight_ntp_policy":                                          dataSourceNtpPolicy(),
		"intersight_telemetry_time_series":                               dataSourceTelemetryTimeSeries(),
		"intersight_ls_service_profile":                                  dataSourceLsServiceProfile(),
		"intersight_hyperflex_hxdp_version":                              dataSourceHyperflexHxdpVersion(),
		"intersight_iam_privilege_set":                                   dataSourceIamPrivilegeSet(),
		"intersight_vnic_fc_network_policy":                              dataSourceVnicFcNetworkPolicy(),
		"intersight_storage_physical_disk":                               dataSourceStoragePhysicalDisk(),
		"intersight_fc_physical_port":                                    dataSourceFcPhysicalPort(),
		"intersight_cond_hcl_status":                                     dataSourceCondHclStatus(),
		"intersight_appliance_diag_setting":                              dataSourceApplianceDiagSetting(),
		"intersight_tam_advisory_instance":                               dataSourceTamAdvisoryInstance(),
		"intersight_snmp_policy":                                         dataSourceSnmpPolicy(),
		"intersight_hyperflex_app_catalog":                               dataSourceHyperflexAppCatalog(),
		"intersight_storage_sas_port":                                    dataSourceStorageSasPort(),
		"intersight_vnic_san_connectivity_policy":                        dataSourceVnicSanConnectivityPolicy(),
		"intersight_resource_membership_holder":                          dataSourceResourceMembershipHolder(),
		"intersight_iam_session":                                         dataSourceIamSession(),
		"intersight_kvm_policy":                                          dataSourceKvmPolicy(),
		"intersight_niaapi_apic_release_recommend":                       dataSourceNiaapiApicReleaseRecommend(),
		"intersight_memory_persistent_memory_namespace":                  dataSourceMemoryPersistentMemoryNamespace(),
		"intersight_iam_private_key_spec":                                dataSourceIamPrivateKeySpec(),
		"intersight_resource_group_member":                               dataSourceResourceGroupMember(),
		"intersight_iam_domain_group":                                    dataSourceIamDomainGroup(),
		"intersight_vnic_eth_if":                                         dataSourceVnicEthIf(),
		"intersight_pci_link":                                            dataSourcePciLink(),
		"intersight_inventory_device_info":                               dataSourceInventoryDeviceInfo(),
		"intersight_storage_pure_protection_group_snapshot":              dataSourceStoragePureProtectionGroupSnapshot(),
		"intersight_storage_pure_disk":                                   dataSourceStoragePureDisk(),
		"intersight_iam_end_point_user_role":                             dataSourceIamEndPointUserRole(),
		"intersight_management_entity":                                   dataSourceManagementEntity(),
		"intersight_asset_cluster_member":                                dataSourceAssetClusterMember(),
		"intersight_workflow_batch_api_executor":                         dataSourceWorkflowBatchApiExecutor(),
		"intersight_appliance_node_info":                                 dataSourceApplianceNodeInfo(),
		"intersight_recovery_on_demand_backup":                           dataSourceRecoveryOnDemandBackup(),
		"intersight_ssh_policy":                                          dataSourceSshPolicy(),
		"intersight_processor_unit":                                      dataSourceProcessorUnit(),
		"intersight_niaapi_apic_hweol":                                   dataSourceNiaapiApicHweol(),
		"intersight_vnic_fc_adapter_policy":                              dataSourceVnicFcAdapterPolicy(),
		"intersight_software_hyperflex_distributable":                    dataSourceSoftwareHyperflexDistributable(),
		"intersight_firmware_upgrade":                                    dataSourceFirmwareUpgrade(),
		"intersight_hcl_operating_system":                                dataSourceHclOperatingSystem(),
		"intersight_appliance_data_export_policy":                        dataSourceApplianceDataExportPolicy(),
		"intersight_equipment_io_card":                                   dataSourceEquipmentIoCard(),
		"intersight_top_system":                                          dataSourceTopSystem(),
		"intersight_adapter_unit":                                        dataSourceAdapterUnit(),
		"intersight_hyperflex_software_version_policy":                   dataSourceHyperflexSoftwareVersionPolicy(),
		"intersight_hyperflex_config_result_entry":                       dataSourceHyperflexConfigResultEntry(),
		"intersight_port_group":                                          dataSourcePortGroup(),
		"intersight_vnic_eth_adapter_policy":                             dataSourceVnicEthAdapterPolicy(),
		"intersight_cond_hcl_status_job":                                 dataSourceCondHclStatusJob(),
		"intersight_softwarerepository_authorization":                    dataSourceSoftwarerepositoryAuthorization(),
		"intersight_iam_permission":                                      dataSourceIamPermission(),
		"intersight_pci_device":                                          dataSourcePciDevice(),
		"intersight_appliance_certificate_setting":                       dataSourceApplianceCertificateSetting(),
		"intersight_recovery_backup_config_policy":                       dataSourceRecoveryBackupConfigPolicy(),
		"intersight_iam_role":                                            dataSourceIamRole(),
		"intersight_hyperflex_cluster_storage_policy":                    dataSourceHyperflexClusterStoragePolicy(),
		"intersight_equipment_fan_module":                                dataSourceEquipmentFanModule(),
		"intersight_niaapi_apic_latest_maintained_release":               dataSourceNiaapiApicLatestMaintainedRelease(),
		"intersight_hyperflex_server_firmware_version":                   dataSourceHyperflexServerFirmwareVersion(),
		"intersight_storage_disk_group_policy":                           dataSourceStorageDiskGroupPolicy(),
		"intersight_ether_physical_port":                                 dataSourceEtherPhysicalPort(),
		"intersight_storage_sas_expander":                                dataSourceStorageSasExpander(),
		"intersight_iaas_device_status":                                  dataSourceIaasDeviceStatus(),
		"intersight_storage_pure_host":                                   dataSourceStoragePureHost(),
		"intersight_tam_advisory_count":                                  dataSourceTamAdvisoryCount(),
		"intersight_iam_app_registration":                                dataSourceIamAppRegistration(),
		"intersight_graphics_card":                                       dataSourceGraphicsCard(),
		"intersight_hyperflex_alarm":                                     dataSourceHyperflexAlarm(),
		"intersight_hyperflex_sys_config_policy":                         dataSourceHyperflexSysConfigPolicy(),
		"intersight_recovery_restore":                                    dataSourceRecoveryRestore(),
		"intersight_iam_system":                                          dataSourceIamSystem(),
		"intersight_adapter_ext_eth_interface":                           dataSourceAdapterExtEthInterface(),
		"intersight_iam_user":                                            dataSourceIamUser(),
		"intersight_iam_end_point_role":                                  dataSourceIamEndPointRole(),
		"intersight_niaapi_dcnm_sweol":                                   dataSourceNiaapiDcnmSweol(),
		"intersight_workflow_task_info":                                  dataSourceWorkflowTaskInfo(),
		"intersight_iam_account":                                         dataSourceIamAccount(),
		"intersight_appliance_device_claim":                              dataSourceApplianceDeviceClaim(),
		"intersight_equipment_rack_enclosure_slot":                       dataSourceEquipmentRackEnclosureSlot(),
		"intersight_niaapi_apic_field_notice":                            dataSourceNiaapiApicFieldNotice(),
		"intersight_memory_unit":                                         dataSourceMemoryUnit(),
		"intersight_recovery_schedule_config_policy":                     dataSourceRecoveryScheduleConfigPolicy(),
		"intersight_storage_flex_util_physical_drive":                    dataSourceStorageFlexUtilPhysicalDrive(),
		"intersight_management_controller":                               dataSourceManagementController(),
		"intersight_niaapi_dcnm_field_notice":                            dataSourceNiaapiDcnmFieldNotice(),
		"intersight_appliance_upgrade_policy":                            dataSourceApplianceUpgradePolicy(),
		"intersight_equipment_chassis":                                   dataSourceEquipmentChassis(),
		"intersight_iam_privilege":                                       dataSourceIamPrivilege(),
		"intersight_niaapi_file_downloader":                              dataSourceNiaapiFileDownloader(),
		"intersight_appliance_backup_policy":                             dataSourceApplianceBackupPolicy(),
		"intersight_cond_hcl_status_detail":                              dataSourceCondHclStatusDetail(),
		"intersight_iam_end_point_user_policy":                           dataSourceIamEndPointUserPolicy(),
		"intersight_workflow_workflow_info":                              dataSourceWorkflowWorkflowInfo(),
		"intersight_adapter_host_eth_interface":                          dataSourceAdapterHostEthInterface(),
		"intersight_hyperflex_cluster_network_policy":                    dataSourceHyperflexClusterNetworkPolicy(),
		"intersight_os_catalog":                                          dataSourceOsCatalog(),
		"intersight_server_profile":                                      dataSourceServerProfile(),
		"intersight_niaapi_version_regex":                                dataSourceNiaapiVersionRegex(),
		"intersight_memory_persistent_memory_configuration":              dataSourceMemoryPersistentMemoryConfiguration(),
		"intersight_firmware_eula":                                       dataSourceFirmwareEula(),
		"intersight_storage_pure_host_lun":                               dataSourceStoragePureHostLun(),
		"intersight_storage_pure_volume":                                 dataSourceStoragePureVolume(),
		"intersight_server_config_result_entry":                          dataSourceServerConfigResultEntry(),
		"intersight_storage_pure_snapshot_schedule":                      dataSourceStoragePureSnapshotSchedule(),
		"intersight_hyperflex_auto_support_policy":                       dataSourceHyperflexAutoSupportPolicy(),
		"intersight_inventory_dn_mo_binding":                             dataSourceInventoryDnMoBinding(),
		"intersight_security_unit":                                       dataSourceSecurityUnit(),
		"intersight_iam_o_auth_token":                                    dataSourceIamOAuthToken(),
		"intersight_adapter_host_iscsi_interface":                        dataSourceAdapterHostIscsiInterface(),
		"intersight_storage_virtual_drive_extension":                     dataSourceStorageVirtualDriveExtension(),
		"intersight_niaapi_nia_metadata":                                 dataSourceNiaapiNiaMetadata(),
		"intersight_vnic_lan_connectivity_policy":                        dataSourceVnicLanConnectivityPolicy(),
		"intersight_storage_physical_disk_usage":                         dataSourceStoragePhysicalDiskUsage(),
		"intersight_server_config_change_detail":                         dataSourceServerConfigChangeDetail(),
		"intersight_memory_persistent_memory_unit":                       dataSourceMemoryPersistentMemoryUnit(),
		"intersight_iam_certificate_request":                             dataSourceIamCertificateRequest(),
		"intersight_iaas_most_run_tasks":                                 dataSourceIaasMostRunTasks(),
		"intersight_vmedia_policy":                                       dataSourceVmediaPolicy(),
		"intersight_hcl_service_status":                                  dataSourceHclServiceStatus(),
		"intersight_equipment_rack_enclosure":                            dataSourceEquipmentRackEnclosure(),
		"intersight_iam_ldap_provider":                                   dataSourceIamLdapProvider(),
		"intersight_workflow_pending_dynamic_workflow_info":              dataSourceWorkflowPendingDynamicWorkflowInfo(),
		"intersight_bios_unit":                                           dataSourceBiosUnit(),
		"intersight_license_account_license_data":                        dataSourceLicenseAccountLicenseData(),
		"intersight_license_license_info":                                dataSourceLicenseLicenseInfo(),
		"intersight_asset_managed_device":                                dataSourceAssetManagedDevice(),
		"intersight_storage_flex_util_virtual_drive":                     dataSourceStorageFlexUtilVirtualDrive(),
		"intersight_equipment_switch_card":                               dataSourceEquipmentSwitchCard(),
		"intersight_storage_pure_controller":                             dataSourceStoragePureController(),
		"intersight_hyperflex_vcenter_config_policy":                     dataSourceHyperflexVcenterConfigPolicy(),
		"intersight_forecast_instance":                                   dataSourceForecastInstance(),
		"intersight_hyperflex_cluster":                                   dataSourceHyperflexCluster(),
		"intersight_meta_definition":                                     dataSourceMetaDefinition(),
		"intersight_equipment_psu":                                       dataSourceEquipmentPsu(),
		"intersight_storage_flex_flash_physical_drive":                   dataSourceStorageFlexFlashPhysicalDrive(),
		"intersight_network_element_summary":                             dataSourceNetworkElementSummary(),
		"intersight_softwarerepository_operating_system_file":            dataSourceSoftwarerepositoryOperatingSystemFile(),
		"intersight_hyperflex_cluster_profile":                           dataSourceHyperflexClusterProfile(),
		"intersight_hyperflex_config_result":                             dataSourceHyperflexConfigResult(),
		"intersight_inventory_generic_inventory":                         dataSourceInventoryGenericInventory(),
		"intersight_iam_session_limits":                                  dataSourceIamSessionLimits(),
		"intersight_asset_device_connector_manager":                      dataSourceAssetDeviceConnectorManager(),
		"intersight_equipment_tpm":                                       dataSourceEquipmentTpm(),
		"intersight_hcl_operating_system_vendor":                         dataSourceHclOperatingSystemVendor(),
		"intersight_equipment_device_summary":                            dataSourceEquipmentDeviceSummary(),
		"intersight_memory_persistent_memory_namespace_config_result":    dataSourceMemoryPersistentMemoryNamespaceConfigResult(),
		"intersight_license_customer_op":                                 dataSourceLicenseCustomerOp(),
		"intersight_iam_resource_permission":                             dataSourceIamResourcePermission(),
		"intersight_recovery_config_result":                              dataSourceRecoveryConfigResult(),
		"intersight_equipment_system_io_controller":                      dataSourceEquipmentSystemIoController(),
		"intersight_hyperflex_node":                                      dataSourceHyperflexNode(),
	}
}
