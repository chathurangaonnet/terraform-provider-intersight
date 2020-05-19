package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceBiosPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBiosPolicyRead,
		Schema: map[string]*schema.Schema{
			"acs_control_gpu1state": {
				Description: "BIOS Token for setting ACS Control GPU-1 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu2state": {
				Description: "BIOS Token for setting ACS Control GPU-2 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu3state": {
				Description: "BIOS Token for setting ACS Control GPU-3 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu4state": {
				Description: "BIOS Token for setting ACS Control GPU-4 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu5state": {
				Description: "BIOS Token for setting ACS Control GPU-5 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu6state": {
				Description: "BIOS Token for setting ACS Control GPU-6 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu7state": {
				Description: "BIOS Token for setting ACS Control GPU-7 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_gpu8state": {
				Description: "BIOS Token for setting ACS Control GPU-8 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot11state": {
				Description: "BIOS Token for setting ACS Control Slot 11 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot12state": {
				Description: "BIOS Token for setting ACS Control Slot 12 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot13state": {
				Description: "BIOS Token for setting ACS Control Slot 13 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"acs_control_slot14state": {
				Description: "BIOS Token for setting ACS Control Slot 14 configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"adjacent_cache_line_prefetch": {
				Description: "BIOS Token for setting Adjacent Cache Line Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"all_usb_devices": {
				Description: "BIOS Token for setting All USB Devices configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"altitude": {
				Description: "BIOS Token for setting Altitude configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"aspm_support": {
				Description: "BIOS Token for setting ASPM Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"assert_nmi_on_perr": {
				Description: "BIOS Token for setting Assert NMI on PERR configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"assert_nmi_on_serr": {
				Description: "BIOS Token for setting Assert NMI on SERR configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"auto_cc_state": {
				Description: "BIOS Token for setting Autonomous Core C-state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"autonumous_cstate_enable": {
				Description: "BIOS Token for setting CPU Autonomous Cstate configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"baud_rate": {
				Description: "BIOS Token for setting Baud rate configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"bme_dma_mitigation": {
				Description: "BIOS Token for setting BME DMA Mitigation configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_option_num_retry": {
				Description: "BIOS Token for setting Number of Retries configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_option_re_cool_down": {
				Description: "BIOS Token for setting Cool Down Time  (sec) configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_option_retry": {
				Description: "BIOS Token for setting Boot option retry configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"boot_performance_mode": {
				Description: "BIOS Token for setting Boot Performance Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_cpb": {
				Description: "BIOS Token for setting Core Performance Boost configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_gen_downcore_ctrl": {
				Description: "BIOS Token for setting Downcore control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_global_cstate_ctrl": {
				Description: "BIOS Token for setting Global C-state Control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_l1stream_hw_prefetcher": {
				Description: "BIOS Token for setting L1 Stream HW Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_cpu_l2stream_hw_prefetcher": {
				Description: "BIOS Token for setting L2 Stream HW Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_determinism_slider": {
				Description: "BIOS Token for setting Determinism Slider configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_gnb_nb_iommu": {
				Description: "BIOS Token for setting IOMMU configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_mem_ctrl_bank_group_swap_ddr4": {
				Description: "BIOS Token for setting Bank Group Swap configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmn_mem_map_bank_interleave_ddr4": {
				Description: "BIOS Token for setting Chipselect Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_cmnc_tdp_ctl": {
				Description: "BIOS Token for setting cTDP Control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_df_cmn_mem_intlv": {
				Description: "BIOS Token for setting Memory interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cbs_df_cmn_mem_intlv_size": {
				Description: "BIOS Token for setting Memory interleaving size configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cdn_enable": {
				Description: "BIOS Token for setting Consistent Device Naming configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cdn_support": {
				Description: "BIOS Token for setting CDN Support for LOM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"channel_inter_leave": {
				Description: "BIOS Token for setting Channel Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cisco_adaptive_mem_training": {
				Description: "BIOS Token for setting Adaptive Memory Training configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cisco_debug_level": {
				Description: "BIOS Token for setting BIOS Techlog Level configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cisco_oprom_launch_optimization": {
				Description: "BIOS Token for setting OptionROM Launch Optimization configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cke_low_policy": {
				Description: "BIOS Token for setting CKE Low Policy configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"class_id": {
				Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"closed_loop_therm_throtl": {
				Description: "BIOS Token for setting Closed Loop Therm Throt configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cmci_enable": {
				Description: "BIOS Token for setting Processor CMCI configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"config_tdp": {
				Description: "BIOS Token for setting Config TDP configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"console_redirection": {
				Description: "BIOS Token for setting Console redirection configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"core_multi_processing": {
				Description: "BIOS Token for setting Core MultiProcessing configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_energy_performance": {
				Description: "BIOS Token for setting Energy Performance configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_frequency_floor": {
				Description: "BIOS Token for setting Frequency Floor Override configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_performance": {
				Description: "BIOS Token for setting CPU Performance configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cpu_power_management": {
				Description: "BIOS Token for setting Power Technology configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dcpmm_firmware_downgrade": {
				Description: "BIOS Token for setting DCPMM Firmware Downgrade configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"demand_scrub": {
				Description: "BIOS Token for setting Demand Scrub configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"direct_cache_access": {
				Description: "BIOS Token for setting Direct Cache Access Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dram_clock_throttling": {
				Description: "BIOS Token for setting DRAM Clock Throttling configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dram_refresh_rate": {
				Description: "BIOS Token for setting DRAM Refresh Rate configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"energy_efficient_turbo": {
				Description: "BIOS Token for setting Energy Efficient Turbo configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"eng_perf_tuning": {
				Description: "BIOS Token for setting Energy Performance Tuning configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enhanced_intel_speed_step_tech": {
				Description: "BIOS Token for setting Enhanced Intel Speedstep (R) Technology configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"epp_profile": {
				Description: "BIOS Token for setting EPP Profile configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"execute_disable_bit": {
				Description: "BIOS Token for setting Execute Disable Bit configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"extended_apic": {
				Description: "BIOS Token for setting Local X2 Apic configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"flow_control": {
				Description: "BIOS Token for setting Flow Control configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"frb2enable": {
				Description: "BIOS Token for setting FRB-2 Timer configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hardware_prefetch": {
				Description: "BIOS Token for setting Hardware Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hwpm_enable": {
				Description: "BIOS Token for setting CPU Hardware Power Management configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"imc_interleave": {
				Description: "BIOS Token for setting IMC Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_hyper_threading_tech": {
				Description: "BIOS Token for setting Intel HyperThreading Tech configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_speed_select": {
				Description: "BIOS Token for setting Intel Speed Select configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_turbo_boost_tech": {
				Description: "BIOS Token for setting Intel Turbo Boost Tech configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_virtualization_technology": {
				Description: "BIOS Token for setting Intel (R) VT configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vt_for_directed_io": {
				Description: "BIOS Token for setting Intel VT for directed IO configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtd_coherency_support": {
				Description: "BIOS Token for setting Intel (R) VT-d Coherency Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtd_interrupt_remapping": {
				Description: "BIOS Token for setting Intel (R) VT-d Interrupt Remapping configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtd_pass_through_dma_support": {
				Description: "BIOS Token for setting Intel (R) VT-d PassThrough DMA support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"intel_vtdats_support": {
				Description: "BIOS Token for setting Intel VTD ATS support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ioh_error_enable": {
				Description: "BIOS Token for setting IIO Error Enable configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ioh_resource": {
				Description: "BIOS Token for setting IOH Resource Allocation configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ip_prefetch": {
				Description: "BIOS Token for setting DCU IP Prefetcher configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv4pxe": {
				Description: "BIOS Token for setting IPv4 PXE Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ipv6pxe": {
				Description: "BIOS Token for setting IPV6 PXE Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"kti_prefetch": {
				Description: "BIOS Token for setting KTI Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"legacy_os_redirection": {
				Description: "BIOS Token for setting Legacy OS redirection configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"legacy_usb_support": {
				Description: "BIOS Token for setting Legacy USB Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"llc_prefetch": {
				Description: "BIOS Token for setting LLC Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port0state": {
				Description: "BIOS Token for setting LOM Port 0 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port1state": {
				Description: "BIOS Token for setting LOM Port 1 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port2state": {
				Description: "BIOS Token for setting LOM Port 2 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_port3state": {
				Description: "BIOS Token for setting LOM Port 3 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lom_ports_all_state": {
				Description: "BIOS Token for setting All Onboard LOM Ports configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"lv_ddr_mode": {
				Description: "BIOS Token for setting Low Voltage DDR Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"make_device_non_bootable": {
				Description: "BIOS Token for setting Make Device Non Bootable configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"memory_inter_leave": {
				Description: "BIOS Token for setting Memory Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"memory_mapped_io_above4gb": {
				Description: "BIOS Token for setting Memory mapped IO above 4GiB configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"memory_size_limit": {
				Description: "BIOS Token for setting Memory Size Limit in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mirroring_mode": {
				Description: "BIOS Token for setting Mirroring Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"mmcfg_base": {
				Description: "BIOS Token for setting MMCFG BASE configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Description: "Name of the concrete policy.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_stack": {
				Description: "BIOS Token for setting Network Stack configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"numa_optimized": {
				Description: "BIOS Token for setting NUMA optimized configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"onboard10gbit_lom": {
				Description: "BIOS Token for setting Onboard 10Gbit LOM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboard_gbit_lom": {
				Description: "BIOS Token for setting Onboard Gbit LOM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboard_scu_storage_support": {
				Description: "BIOS Token for setting Onboard SCU Storage Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"onboard_scu_storage_sw_stack": {
				Description: "BIOS Token for setting Onboard SCU Storage SW Stack configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"organization": {
				Description: "A reference to a organizationOrganization resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"os_boot_watchdog_timer": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"os_boot_watchdog_timer_policy": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer Policy configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"os_boot_watchdog_timer_timeout": {
				Description: "BIOS Token for setting OS Boot Watchdog Timer Timeout configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"out_of_band_mgmt_port": {
				Description: "BIOS Token for setting Out-of-Band Mgmt Port configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"package_cstate_limit": {
				Description: "BIOS Token for setting Package C State Limit configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_mode_config": {
				Description: "BIOS Token for setting Partial Memory Mirror Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_percent": {
				Description: "BIOS Token for setting Partial Mirror percentage configuration (0.00-50.00).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value1": {
				Description: "BIOS Token for setting Partial Mirror1 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value2": {
				Description: "BIOS Token for setting Partial Mirror2 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value3": {
				Description: "BIOS Token for setting Partial Mirror3 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"partial_mirror_value4": {
				Description: "BIOS Token for setting Partial Mirror4 Size in GiB configuration (0-65535).",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"patrol_scrub": {
				Description: "BIOS Token for setting Patrol Scrub configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"patrol_scrub_duration": {
				Description: "BIOS Token for setting Patrol Scrub Interval configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pc_ie_ras_support": {
				Description: "BIOS Token for setting PCIe RAS Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pc_ie_ssd_hot_plug_support": {
				Description: "BIOS Token for setting NVMe SSD Hot-Plug Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pch_usb30mode": {
				Description: "BIOS Token for setting xHCI Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pci_option_ro_ms": {
				Description: "BIOS Token for setting All PCIe Slots OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pci_rom_clp": {
				Description: "BIOS Token for setting PCI ROM CLP configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"permission_resources": {
				Description: "An array of relationships to moBaseMo resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"pop_support": {
				Description: "BIOS Token for setting Power ON Password configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"post_error_pause": {
				Description: "BIOS Token for setting POST Error Pause configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_c1e": {
				Description: "BIOS Token for setting Processor C1E configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_c3report": {
				Description: "BIOS Token for setting Processor C3 Report configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_c6report": {
				Description: "BIOS Token for setting Processor C6 Report configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"processor_cstate": {
				Description: "BIOS Token for setting CPU C State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"profiles": {
				Description: "An array of relationships to policyAbstractConfigProfile resources.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_id": {
							Description: "The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"link": {
							Description: "A URL to an instance of the 'mo.MoRef' class.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
				Computed: true,
			},
			"psata": {
				Description: "BIOS Token for setting P-SATA mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pstate_coord_type": {
				Description: "BIOS Token for setting P-STATE Coordination configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"putty_key_pad": {
				Description: "BIOS Token for setting Putty KeyPad configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"pwr_perf_tuning": {
				Description: "BIOS Token for setting Power Performance Tuning configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"qpi_link_frequency": {
				Description: "BIOS Token for setting QPI Link Frequency Select configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"qpi_snoop_mode": {
				Description: "BIOS Token for setting QPI Snoop Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"rank_inter_leave": {
				Description: "BIOS Token for setting Rank Interleaving configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"redirection_after_post": {
				Description: "BIOS Token for setting Redirection After BIOS POST configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sata_mode_select": {
				Description: "BIOS Token for setting SATA mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"select_memory_ras_configuration": {
				Description: "BIOS Token for setting SelectMemory RAS configuration configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"select_ppr_type": {
				Description: "BIOS Token for setting Select PPR Type configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial_port_aenable": {
				Description: "BIOS Token for setting Serial A Enable configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"single_pctl_enable": {
				Description: "BIOS Token for setting Single PCTL configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot10link_speed": {
				Description: "BIOS Token for setting PCIe Slot:10 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot10state": {
				Description: "BIOS Token for setting Slot 10 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot11link_speed": {
				Description: "BIOS Token for setting PCIe Slot:11 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot11state": {
				Description: "BIOS Token for setting Slot 11 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot12link_speed": {
				Description: "BIOS Token for setting PCIe Slot:12 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot12state": {
				Description: "BIOS Token for setting Slot 12 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot13state": {
				Description: "BIOS Token for setting PCIe Slot 13 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot14state": {
				Description: "BIOS Token for setting PCIe Slot 14 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot1state": {
				Description: "BIOS Token for setting Slot 1 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot2state": {
				Description: "BIOS Token for setting Slot 2 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot3link_speed": {
				Description: "BIOS Token for setting PCIe Slot:3 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot3state": {
				Description: "BIOS Token for setting Slot 3 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot4link_speed": {
				Description: "BIOS Token for setting PCIe Slot:4 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot4state": {
				Description: "BIOS Token for setting Slot 4 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:5 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot5state": {
				Description: "BIOS Token for setting Slot 5 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:6 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot6state": {
				Description: "BIOS Token for setting Slot 6 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot7link_speed": {
				Description: "BIOS Token for setting PCIe Slot:7 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot7state": {
				Description: "BIOS Token for setting Slot 7 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot8link_speed": {
				Description: "BIOS Token for setting PCIe Slot:8 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot8state": {
				Description: "BIOS Token for setting Slot 8 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot9link_speed": {
				Description: "BIOS Token for setting PCIe Slot:9 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot9state": {
				Description: "BIOS Token for setting Slot 9 state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_flom_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FLOM Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_nvme1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front Nvme1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_nvme2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front Nvme2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_front_slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Front2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu1state": {
				Description: "BIOS Token for setting GPU1 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu2state": {
				Description: "BIOS Token for setting GPU2 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu3state": {
				Description: "BIOS Token for setting GPU3 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu4state": {
				Description: "BIOS Token for setting GPU4 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu5state": {
				Description: "BIOS Token for setting GPU5 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu6state": {
				Description: "BIOS Token for setting GPU6 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu7state": {
				Description: "BIOS Token for setting GPU7 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_gpu8state": {
				Description: "BIOS Token for setting GPU8 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_hba_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:HBA Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_hba_state": {
				Description: "BIOS Token for setting PCIe Slot:HBA OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_lom1link": {
				Description: "BIOS Token for setting PCIe LOM:1 Link configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_lom2link": {
				Description: "BIOS Token for setting PCIe LOM:2 Link configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mezz_state": {
				Description: "BIOS Token for setting Slot Mezz state configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mlom_link_speed": {
				Description: "BIOS Token for setting PCIe Slot:MLOM Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mlom_state": {
				Description: "BIOS Token for setting PCIe Slot MLOM OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mraid_link_speed": {
				Description: "BIOS Token for setting MRAID Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_mraid_state": {
				Description: "BIOS Token for setting PCIe Slot MRAID OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n10state": {
				Description: "BIOS Token for setting PCIe Slot N10 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n11state": {
				Description: "BIOS Token for setting PCIe Slot N11 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n12state": {
				Description: "BIOS Token for setting PCIe Slot N12 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n13state": {
				Description: "BIOS Token for setting PCIe Slot N13 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n14state": {
				Description: "BIOS Token for setting PCIe Slot N14 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n15state": {
				Description: "BIOS Token for setting PCIe Slot N15 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n16state": {
				Description: "BIOS Token for setting PCIe Slot N16 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n17state": {
				Description: "BIOS Token for setting PCIe Slot N17 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n18state": {
				Description: "BIOS Token for setting PCIe Slot N18 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n19state": {
				Description: "BIOS Token for setting PCIe Slot N19 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n1state": {
				Description: "BIOS Token for setting PCIe Slot N1 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n20state": {
				Description: "BIOS Token for setting PCIe Slot N20 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n21state": {
				Description: "BIOS Token for setting PCIe Slot N21 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n22state": {
				Description: "BIOS Token for setting PCIe Slot N22 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n23state": {
				Description: "BIOS Token for setting PCIe Slot N23 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n24state": {
				Description: "BIOS Token for setting PCIe Slot N24 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n2state": {
				Description: "BIOS Token for setting PCIe Slot N2 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n3state": {
				Description: "BIOS Token for setting PCIe Slot N3 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n4state": {
				Description: "BIOS Token for setting PCIe Slot N4 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n5state": {
				Description: "BIOS Token for setting PCIe Slot N5 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n6state": {
				Description: "BIOS Token for setting PCIe Slot N6 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n7state": {
				Description: "BIOS Token for setting PCIe Slot N7 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n8state": {
				Description: "BIOS Token for setting PCIe Slot N8 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_n9state": {
				Description: "BIOS Token for setting PCIe Slot N9 OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_raid_link_speed": {
				Description: "BIOS Token for setting RAID Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_raid_state": {
				Description: "BIOS Token for setting PCIe Slot RAID OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Rear Nvme1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme1state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 1 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Rear Nvme2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme2state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 2 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme3state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 3 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme4state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 4 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme5state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 5 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme6state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 6 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme7state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 7 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_rear_nvme8state": {
				Description: "BIOS Token for setting PCIe Slot:Rear NVME 8 OptionRom configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser1slot3link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser1 Slot3 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2slot4link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot4 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2slot5link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot5 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_riser2slot6link_speed": {
				Description: "BIOS Token for setting PCIe Slot:Riser2 Slot6 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_sas_state": {
				Description: "BIOS Token for setting PCIe Slot:SAS OptionROM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_ssd_slot1link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FrontSSD1 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slot_ssd_slot2link_speed": {
				Description: "BIOS Token for setting PCIe Slot:FrontSSD2 Link Speed configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"smee": {
				Description: "BIOS Token for setting SMEE configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"smt_mode": {
				Description: "BIOS Token for setting SMT Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snc": {
				Description: "BIOS Token for setting Sub Numa Clustering configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sparing_mode": {
				Description: "BIOS Token for setting Sparing Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sr_iov": {
				Description: "BIOS Token for setting SR-IOV Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"streamer_prefetch": {
				Description: "BIOS Token for setting DCU Streamer Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"svm_mode": {
				Description: "BIOS Token for setting SVM Mode configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"terminal_type": {
				Description: "BIOS Token for setting Terminal Type configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tpm_control": {
				Description: "BIOS Token for setting Trusted Platform Module State configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tpm_support": {
				Description: "BIOS Token for setting TPM Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"txt_support": {
				Description: "BIOS Token for setting Intel Trusted Execution Technology Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"ucsm_boot_order_rule": {
				Description: "BIOS Token for setting Boot Order Rules configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_emul6064": {
				Description: "BIOS Token for setting Port 60/64 Emulation configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_front": {
				Description: "BIOS Token for setting USB Port Front configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_internal": {
				Description: "BIOS Token for setting USB Port Internal configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_kvm": {
				Description: "BIOS Token for setting USB Port KVM configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_rear": {
				Description: "BIOS Token for setting USB Port Rear configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_sd_card": {
				Description: "BIOS Token for setting USB Port SD Card configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_port_vmedia": {
				Description: "BIOS Token for setting USB Port VMedia configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usb_xhci_support": {
				Description: "BIOS Token for setting XHCI Legacy Support configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vga_priority": {
				Description: "BIOS Token for setting VGA Priority configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vmd_enable": {
				Description: "BIOS Token for setting VMD Enablement configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"work_load_config": {
				Description: "BIOS Token for setting Workload Configuration configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"xpt_prefetch": {
				Description: "BIOS Token for setting XPT Prefetch configuration.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceBiosPolicyRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var o = models.NewBiosPolicy()
	if v, ok := d.GetOk("acs_control_gpu1state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu1state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu2state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu2state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu3state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu3state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu4state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu4state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu5state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu5state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu6state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu6state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu7state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu7state(x)
	}
	if v, ok := d.GetOk("acs_control_gpu8state"); ok {
		x := (v.(string))
		o.SetAcsControlGpu8state(x)
	}
	if v, ok := d.GetOk("acs_control_slot11state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot11state(x)
	}
	if v, ok := d.GetOk("acs_control_slot12state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot12state(x)
	}
	if v, ok := d.GetOk("acs_control_slot13state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot13state(x)
	}
	if v, ok := d.GetOk("acs_control_slot14state"); ok {
		x := (v.(string))
		o.SetAcsControlSlot14state(x)
	}
	if v, ok := d.GetOk("adjacent_cache_line_prefetch"); ok {
		x := (v.(string))
		o.SetAdjacentCacheLinePrefetch(x)
	}
	if v, ok := d.GetOk("all_usb_devices"); ok {
		x := (v.(string))
		o.SetAllUsbDevices(x)
	}
	if v, ok := d.GetOk("altitude"); ok {
		x := (v.(string))
		o.SetAltitude(x)
	}
	if v, ok := d.GetOk("aspm_support"); ok {
		x := (v.(string))
		o.SetAspmSupport(x)
	}
	if v, ok := d.GetOk("assert_nmi_on_perr"); ok {
		x := (v.(string))
		o.SetAssertNmiOnPerr(x)
	}
	if v, ok := d.GetOk("assert_nmi_on_serr"); ok {
		x := (v.(string))
		o.SetAssertNmiOnSerr(x)
	}
	if v, ok := d.GetOk("auto_cc_state"); ok {
		x := (v.(string))
		o.SetAutoCcState(x)
	}
	if v, ok := d.GetOk("autonumous_cstate_enable"); ok {
		x := (v.(string))
		o.SetAutonumousCstateEnable(x)
	}
	if v, ok := d.GetOk("baud_rate"); ok {
		x := (v.(string))
		o.SetBaudRate(x)
	}
	if v, ok := d.GetOk("bme_dma_mitigation"); ok {
		x := (v.(string))
		o.SetBmeDmaMitigation(x)
	}
	if v, ok := d.GetOk("boot_option_num_retry"); ok {
		x := (v.(string))
		o.SetBootOptionNumRetry(x)
	}
	if v, ok := d.GetOk("boot_option_re_cool_down"); ok {
		x := (v.(string))
		o.SetBootOptionReCoolDown(x)
	}
	if v, ok := d.GetOk("boot_option_retry"); ok {
		x := (v.(string))
		o.SetBootOptionRetry(x)
	}
	if v, ok := d.GetOk("boot_performance_mode"); ok {
		x := (v.(string))
		o.SetBootPerformanceMode(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_cpb"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuCpb(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_gen_downcore_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuGenDowncoreCtrl(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_global_cstate_ctrl"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuGlobalCstateCtrl(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_l1stream_hw_prefetcher"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuL1streamHwPrefetcher(x)
	}
	if v, ok := d.GetOk("cbs_cmn_cpu_l2stream_hw_prefetcher"); ok {
		x := (v.(string))
		o.SetCbsCmnCpuL2streamHwPrefetcher(x)
	}
	if v, ok := d.GetOk("cbs_cmn_determinism_slider"); ok {
		x := (v.(string))
		o.SetCbsCmnDeterminismSlider(x)
	}
	if v, ok := d.GetOk("cbs_cmn_gnb_nb_iommu"); ok {
		x := (v.(string))
		o.SetCbsCmnGnbNbIommu(x)
	}
	if v, ok := d.GetOk("cbs_cmn_mem_ctrl_bank_group_swap_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemCtrlBankGroupSwapDdr4(x)
	}
	if v, ok := d.GetOk("cbs_cmn_mem_map_bank_interleave_ddr4"); ok {
		x := (v.(string))
		o.SetCbsCmnMemMapBankInterleaveDdr4(x)
	}
	if v, ok := d.GetOk("cbs_cmnc_tdp_ctl"); ok {
		x := (v.(string))
		o.SetCbsCmncTdpCtl(x)
	}
	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlv(x)
	}
	if v, ok := d.GetOk("cbs_df_cmn_mem_intlv_size"); ok {
		x := (v.(string))
		o.SetCbsDfCmnMemIntlvSize(x)
	}
	if v, ok := d.GetOk("cdn_enable"); ok {
		x := (v.(string))
		o.SetCdnEnable(x)
	}
	if v, ok := d.GetOk("cdn_support"); ok {
		x := (v.(string))
		o.SetCdnSupport(x)
	}
	if v, ok := d.GetOk("channel_inter_leave"); ok {
		x := (v.(string))
		o.SetChannelInterLeave(x)
	}
	if v, ok := d.GetOk("cisco_adaptive_mem_training"); ok {
		x := (v.(string))
		o.SetCiscoAdaptiveMemTraining(x)
	}
	if v, ok := d.GetOk("cisco_debug_level"); ok {
		x := (v.(string))
		o.SetCiscoDebugLevel(x)
	}
	if v, ok := d.GetOk("cisco_oprom_launch_optimization"); ok {
		x := (v.(string))
		o.SetCiscoOpromLaunchOptimization(x)
	}
	if v, ok := d.GetOk("cke_low_policy"); ok {
		x := (v.(string))
		o.SetCkeLowPolicy(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("closed_loop_therm_throtl"); ok {
		x := (v.(string))
		o.SetClosedLoopThermThrotl(x)
	}
	if v, ok := d.GetOk("cmci_enable"); ok {
		x := (v.(string))
		o.SetCmciEnable(x)
	}
	if v, ok := d.GetOk("config_tdp"); ok {
		x := (v.(string))
		o.SetConfigTdp(x)
	}
	if v, ok := d.GetOk("console_redirection"); ok {
		x := (v.(string))
		o.SetConsoleRedirection(x)
	}
	if v, ok := d.GetOk("core_multi_processing"); ok {
		x := (v.(string))
		o.SetCoreMultiProcessing(x)
	}
	if v, ok := d.GetOk("cpu_energy_performance"); ok {
		x := (v.(string))
		o.SetCpuEnergyPerformance(x)
	}
	if v, ok := d.GetOk("cpu_frequency_floor"); ok {
		x := (v.(string))
		o.SetCpuFrequencyFloor(x)
	}
	if v, ok := d.GetOk("cpu_performance"); ok {
		x := (v.(string))
		o.SetCpuPerformance(x)
	}
	if v, ok := d.GetOk("cpu_power_management"); ok {
		x := (v.(string))
		o.SetCpuPowerManagement(x)
	}
	if v, ok := d.GetOk("dcpmm_firmware_downgrade"); ok {
		x := (v.(string))
		o.SetDcpmmFirmwareDowngrade(x)
	}
	if v, ok := d.GetOk("demand_scrub"); ok {
		x := (v.(string))
		o.SetDemandScrub(x)
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.SetDescription(x)
	}
	if v, ok := d.GetOk("direct_cache_access"); ok {
		x := (v.(string))
		o.SetDirectCacheAccess(x)
	}
	if v, ok := d.GetOk("dram_clock_throttling"); ok {
		x := (v.(string))
		o.SetDramClockThrottling(x)
	}
	if v, ok := d.GetOk("dram_refresh_rate"); ok {
		x := (v.(string))
		o.SetDramRefreshRate(x)
	}
	if v, ok := d.GetOk("energy_efficient_turbo"); ok {
		x := (v.(string))
		o.SetEnergyEfficientTurbo(x)
	}
	if v, ok := d.GetOk("eng_perf_tuning"); ok {
		x := (v.(string))
		o.SetEngPerfTuning(x)
	}
	if v, ok := d.GetOk("enhanced_intel_speed_step_tech"); ok {
		x := (v.(string))
		o.SetEnhancedIntelSpeedStepTech(x)
	}
	if v, ok := d.GetOk("epp_profile"); ok {
		x := (v.(string))
		o.SetEppProfile(x)
	}
	if v, ok := d.GetOk("execute_disable_bit"); ok {
		x := (v.(string))
		o.SetExecuteDisableBit(x)
	}
	if v, ok := d.GetOk("extended_apic"); ok {
		x := (v.(string))
		o.SetExtendedApic(x)
	}
	if v, ok := d.GetOk("flow_control"); ok {
		x := (v.(string))
		o.SetFlowControl(x)
	}
	if v, ok := d.GetOk("frb2enable"); ok {
		x := (v.(string))
		o.SetFrb2enable(x)
	}
	if v, ok := d.GetOk("hardware_prefetch"); ok {
		x := (v.(string))
		o.SetHardwarePrefetch(x)
	}
	if v, ok := d.GetOk("hwpm_enable"); ok {
		x := (v.(string))
		o.SetHwpmEnable(x)
	}
	if v, ok := d.GetOk("imc_interleave"); ok {
		x := (v.(string))
		o.SetImcInterleave(x)
	}
	if v, ok := d.GetOk("intel_hyper_threading_tech"); ok {
		x := (v.(string))
		o.SetIntelHyperThreadingTech(x)
	}
	if v, ok := d.GetOk("intel_speed_select"); ok {
		x := (v.(string))
		o.SetIntelSpeedSelect(x)
	}
	if v, ok := d.GetOk("intel_turbo_boost_tech"); ok {
		x := (v.(string))
		o.SetIntelTurboBoostTech(x)
	}
	if v, ok := d.GetOk("intel_virtualization_technology"); ok {
		x := (v.(string))
		o.SetIntelVirtualizationTechnology(x)
	}
	if v, ok := d.GetOk("intel_vt_for_directed_io"); ok {
		x := (v.(string))
		o.SetIntelVtForDirectedIo(x)
	}
	if v, ok := d.GetOk("intel_vtd_coherency_support"); ok {
		x := (v.(string))
		o.SetIntelVtdCoherencySupport(x)
	}
	if v, ok := d.GetOk("intel_vtd_interrupt_remapping"); ok {
		x := (v.(string))
		o.SetIntelVtdInterruptRemapping(x)
	}
	if v, ok := d.GetOk("intel_vtd_pass_through_dma_support"); ok {
		x := (v.(string))
		o.SetIntelVtdPassThroughDmaSupport(x)
	}
	if v, ok := d.GetOk("intel_vtdats_support"); ok {
		x := (v.(string))
		o.SetIntelVtdatsSupport(x)
	}
	if v, ok := d.GetOk("ioh_error_enable"); ok {
		x := (v.(string))
		o.SetIohErrorEnable(x)
	}
	if v, ok := d.GetOk("ioh_resource"); ok {
		x := (v.(string))
		o.SetIohResource(x)
	}
	if v, ok := d.GetOk("ip_prefetch"); ok {
		x := (v.(string))
		o.SetIpPrefetch(x)
	}
	if v, ok := d.GetOk("ipv4pxe"); ok {
		x := (v.(string))
		o.SetIpv4pxe(x)
	}
	if v, ok := d.GetOk("ipv6pxe"); ok {
		x := (v.(string))
		o.SetIpv6pxe(x)
	}
	if v, ok := d.GetOk("kti_prefetch"); ok {
		x := (v.(string))
		o.SetKtiPrefetch(x)
	}
	if v, ok := d.GetOk("legacy_os_redirection"); ok {
		x := (v.(string))
		o.SetLegacyOsRedirection(x)
	}
	if v, ok := d.GetOk("legacy_usb_support"); ok {
		x := (v.(string))
		o.SetLegacyUsbSupport(x)
	}
	if v, ok := d.GetOk("llc_prefetch"); ok {
		x := (v.(string))
		o.SetLlcPrefetch(x)
	}
	if v, ok := d.GetOk("lom_port0state"); ok {
		x := (v.(string))
		o.SetLomPort0state(x)
	}
	if v, ok := d.GetOk("lom_port1state"); ok {
		x := (v.(string))
		o.SetLomPort1state(x)
	}
	if v, ok := d.GetOk("lom_port2state"); ok {
		x := (v.(string))
		o.SetLomPort2state(x)
	}
	if v, ok := d.GetOk("lom_port3state"); ok {
		x := (v.(string))
		o.SetLomPort3state(x)
	}
	if v, ok := d.GetOk("lom_ports_all_state"); ok {
		x := (v.(string))
		o.SetLomPortsAllState(x)
	}
	if v, ok := d.GetOk("lv_ddr_mode"); ok {
		x := (v.(string))
		o.SetLvDdrMode(x)
	}
	if v, ok := d.GetOk("make_device_non_bootable"); ok {
		x := (v.(string))
		o.SetMakeDeviceNonBootable(x)
	}
	if v, ok := d.GetOk("memory_inter_leave"); ok {
		x := (v.(string))
		o.SetMemoryInterLeave(x)
	}
	if v, ok := d.GetOk("memory_mapped_io_above4gb"); ok {
		x := (v.(string))
		o.SetMemoryMappedIoAbove4gb(x)
	}
	if v, ok := d.GetOk("memory_size_limit"); ok {
		x := (v.(string))
		o.SetMemorySizeLimit(x)
	}
	if v, ok := d.GetOk("mirroring_mode"); ok {
		x := (v.(string))
		o.SetMirroringMode(x)
	}
	if v, ok := d.GetOk("mmcfg_base"); ok {
		x := (v.(string))
		o.SetMmcfgBase(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("name"); ok {
		x := (v.(string))
		o.SetName(x)
	}
	if v, ok := d.GetOk("network_stack"); ok {
		x := (v.(string))
		o.SetNetworkStack(x)
	}
	if v, ok := d.GetOk("numa_optimized"); ok {
		x := (v.(string))
		o.SetNumaOptimized(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("onboard10gbit_lom"); ok {
		x := (v.(string))
		o.SetOnboard10gbitLom(x)
	}
	if v, ok := d.GetOk("onboard_gbit_lom"); ok {
		x := (v.(string))
		o.SetOnboardGbitLom(x)
	}
	if v, ok := d.GetOk("onboard_scu_storage_support"); ok {
		x := (v.(string))
		o.SetOnboardScuStorageSupport(x)
	}
	if v, ok := d.GetOk("onboard_scu_storage_sw_stack"); ok {
		x := (v.(string))
		o.SetOnboardScuStorageSwStack(x)
	}
	if v, ok := d.GetOk("os_boot_watchdog_timer"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimer(x)
	}
	if v, ok := d.GetOk("os_boot_watchdog_timer_policy"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimerPolicy(x)
	}
	if v, ok := d.GetOk("os_boot_watchdog_timer_timeout"); ok {
		x := (v.(string))
		o.SetOsBootWatchdogTimerTimeout(x)
	}
	if v, ok := d.GetOk("out_of_band_mgmt_port"); ok {
		x := (v.(string))
		o.SetOutOfBandMgmtPort(x)
	}
	if v, ok := d.GetOk("package_cstate_limit"); ok {
		x := (v.(string))
		o.SetPackageCstateLimit(x)
	}
	if v, ok := d.GetOk("partial_mirror_mode_config"); ok {
		x := (v.(string))
		o.SetPartialMirrorModeConfig(x)
	}
	if v, ok := d.GetOk("partial_mirror_percent"); ok {
		x := (v.(string))
		o.SetPartialMirrorPercent(x)
	}
	if v, ok := d.GetOk("partial_mirror_value1"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue1(x)
	}
	if v, ok := d.GetOk("partial_mirror_value2"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue2(x)
	}
	if v, ok := d.GetOk("partial_mirror_value3"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue3(x)
	}
	if v, ok := d.GetOk("partial_mirror_value4"); ok {
		x := (v.(string))
		o.SetPartialMirrorValue4(x)
	}
	if v, ok := d.GetOk("patrol_scrub"); ok {
		x := (v.(string))
		o.SetPatrolScrub(x)
	}
	if v, ok := d.GetOk("patrol_scrub_duration"); ok {
		x := (v.(string))
		o.SetPatrolScrubDuration(x)
	}
	if v, ok := d.GetOk("pc_ie_ras_support"); ok {
		x := (v.(string))
		o.SetPcIeRasSupport(x)
	}
	if v, ok := d.GetOk("pc_ie_ssd_hot_plug_support"); ok {
		x := (v.(string))
		o.SetPcIeSsdHotPlugSupport(x)
	}
	if v, ok := d.GetOk("pch_usb30mode"); ok {
		x := (v.(string))
		o.SetPchUsb30mode(x)
	}
	if v, ok := d.GetOk("pci_option_ro_ms"); ok {
		x := (v.(string))
		o.SetPciOptionRoMs(x)
	}
	if v, ok := d.GetOk("pci_rom_clp"); ok {
		x := (v.(string))
		o.SetPciRomClp(x)
	}
	if v, ok := d.GetOk("pop_support"); ok {
		x := (v.(string))
		o.SetPopSupport(x)
	}
	if v, ok := d.GetOk("post_error_pause"); ok {
		x := (v.(string))
		o.SetPostErrorPause(x)
	}
	if v, ok := d.GetOk("processor_c1e"); ok {
		x := (v.(string))
		o.SetProcessorC1e(x)
	}
	if v, ok := d.GetOk("processor_c3report"); ok {
		x := (v.(string))
		o.SetProcessorC3report(x)
	}
	if v, ok := d.GetOk("processor_c6report"); ok {
		x := (v.(string))
		o.SetProcessorC6report(x)
	}
	if v, ok := d.GetOk("processor_cstate"); ok {
		x := (v.(string))
		o.SetProcessorCstate(x)
	}
	if v, ok := d.GetOk("psata"); ok {
		x := (v.(string))
		o.SetPsata(x)
	}
	if v, ok := d.GetOk("pstate_coord_type"); ok {
		x := (v.(string))
		o.SetPstateCoordType(x)
	}
	if v, ok := d.GetOk("putty_key_pad"); ok {
		x := (v.(string))
		o.SetPuttyKeyPad(x)
	}
	if v, ok := d.GetOk("pwr_perf_tuning"); ok {
		x := (v.(string))
		o.SetPwrPerfTuning(x)
	}
	if v, ok := d.GetOk("qpi_link_frequency"); ok {
		x := (v.(string))
		o.SetQpiLinkFrequency(x)
	}
	if v, ok := d.GetOk("qpi_snoop_mode"); ok {
		x := (v.(string))
		o.SetQpiSnoopMode(x)
	}
	if v, ok := d.GetOk("rank_inter_leave"); ok {
		x := (v.(string))
		o.SetRankInterLeave(x)
	}
	if v, ok := d.GetOk("redirection_after_post"); ok {
		x := (v.(string))
		o.SetRedirectionAfterPost(x)
	}
	if v, ok := d.GetOk("sata_mode_select"); ok {
		x := (v.(string))
		o.SetSataModeSelect(x)
	}
	if v, ok := d.GetOk("select_memory_ras_configuration"); ok {
		x := (v.(string))
		o.SetSelectMemoryRasConfiguration(x)
	}
	if v, ok := d.GetOk("select_ppr_type"); ok {
		x := (v.(string))
		o.SetSelectPprType(x)
	}
	if v, ok := d.GetOk("serial_port_aenable"); ok {
		x := (v.(string))
		o.SetSerialPortAenable(x)
	}
	if v, ok := d.GetOk("single_pctl_enable"); ok {
		x := (v.(string))
		o.SetSinglePctlEnable(x)
	}
	if v, ok := d.GetOk("slot10link_speed"); ok {
		x := (v.(string))
		o.SetSlot10linkSpeed(x)
	}
	if v, ok := d.GetOk("slot10state"); ok {
		x := (v.(string))
		o.SetSlot10state(x)
	}
	if v, ok := d.GetOk("slot11link_speed"); ok {
		x := (v.(string))
		o.SetSlot11linkSpeed(x)
	}
	if v, ok := d.GetOk("slot11state"); ok {
		x := (v.(string))
		o.SetSlot11state(x)
	}
	if v, ok := d.GetOk("slot12link_speed"); ok {
		x := (v.(string))
		o.SetSlot12linkSpeed(x)
	}
	if v, ok := d.GetOk("slot12state"); ok {
		x := (v.(string))
		o.SetSlot12state(x)
	}
	if v, ok := d.GetOk("slot13state"); ok {
		x := (v.(string))
		o.SetSlot13state(x)
	}
	if v, ok := d.GetOk("slot14state"); ok {
		x := (v.(string))
		o.SetSlot14state(x)
	}
	if v, ok := d.GetOk("slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlot1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot1state"); ok {
		x := (v.(string))
		o.SetSlot1state(x)
	}
	if v, ok := d.GetOk("slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlot2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot2state"); ok {
		x := (v.(string))
		o.SetSlot2state(x)
	}
	if v, ok := d.GetOk("slot3link_speed"); ok {
		x := (v.(string))
		o.SetSlot3linkSpeed(x)
	}
	if v, ok := d.GetOk("slot3state"); ok {
		x := (v.(string))
		o.SetSlot3state(x)
	}
	if v, ok := d.GetOk("slot4link_speed"); ok {
		x := (v.(string))
		o.SetSlot4linkSpeed(x)
	}
	if v, ok := d.GetOk("slot4state"); ok {
		x := (v.(string))
		o.SetSlot4state(x)
	}
	if v, ok := d.GetOk("slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlot5linkSpeed(x)
	}
	if v, ok := d.GetOk("slot5state"); ok {
		x := (v.(string))
		o.SetSlot5state(x)
	}
	if v, ok := d.GetOk("slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlot6linkSpeed(x)
	}
	if v, ok := d.GetOk("slot6state"); ok {
		x := (v.(string))
		o.SetSlot6state(x)
	}
	if v, ok := d.GetOk("slot7link_speed"); ok {
		x := (v.(string))
		o.SetSlot7linkSpeed(x)
	}
	if v, ok := d.GetOk("slot7state"); ok {
		x := (v.(string))
		o.SetSlot7state(x)
	}
	if v, ok := d.GetOk("slot8link_speed"); ok {
		x := (v.(string))
		o.SetSlot8linkSpeed(x)
	}
	if v, ok := d.GetOk("slot8state"); ok {
		x := (v.(string))
		o.SetSlot8state(x)
	}
	if v, ok := d.GetOk("slot9link_speed"); ok {
		x := (v.(string))
		o.SetSlot9linkSpeed(x)
	}
	if v, ok := d.GetOk("slot9state"); ok {
		x := (v.(string))
		o.SetSlot9state(x)
	}
	if v, ok := d.GetOk("slot_flom_link_speed"); ok {
		x := (v.(string))
		o.SetSlotFlomLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontNvme2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontSlot5linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_front_slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlotFrontSlot6linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_gpu1state"); ok {
		x := (v.(string))
		o.SetSlotGpu1state(x)
	}
	if v, ok := d.GetOk("slot_gpu2state"); ok {
		x := (v.(string))
		o.SetSlotGpu2state(x)
	}
	if v, ok := d.GetOk("slot_gpu3state"); ok {
		x := (v.(string))
		o.SetSlotGpu3state(x)
	}
	if v, ok := d.GetOk("slot_gpu4state"); ok {
		x := (v.(string))
		o.SetSlotGpu4state(x)
	}
	if v, ok := d.GetOk("slot_gpu5state"); ok {
		x := (v.(string))
		o.SetSlotGpu5state(x)
	}
	if v, ok := d.GetOk("slot_gpu6state"); ok {
		x := (v.(string))
		o.SetSlotGpu6state(x)
	}
	if v, ok := d.GetOk("slot_gpu7state"); ok {
		x := (v.(string))
		o.SetSlotGpu7state(x)
	}
	if v, ok := d.GetOk("slot_gpu8state"); ok {
		x := (v.(string))
		o.SetSlotGpu8state(x)
	}
	if v, ok := d.GetOk("slot_hba_link_speed"); ok {
		x := (v.(string))
		o.SetSlotHbaLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_hba_state"); ok {
		x := (v.(string))
		o.SetSlotHbaState(x)
	}
	if v, ok := d.GetOk("slot_lom1link"); ok {
		x := (v.(string))
		o.SetSlotLom1link(x)
	}
	if v, ok := d.GetOk("slot_lom2link"); ok {
		x := (v.(string))
		o.SetSlotLom2link(x)
	}
	if v, ok := d.GetOk("slot_mezz_state"); ok {
		x := (v.(string))
		o.SetSlotMezzState(x)
	}
	if v, ok := d.GetOk("slot_mlom_link_speed"); ok {
		x := (v.(string))
		o.SetSlotMlomLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_mlom_state"); ok {
		x := (v.(string))
		o.SetSlotMlomState(x)
	}
	if v, ok := d.GetOk("slot_mraid_link_speed"); ok {
		x := (v.(string))
		o.SetSlotMraidLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_mraid_state"); ok {
		x := (v.(string))
		o.SetSlotMraidState(x)
	}
	if v, ok := d.GetOk("slot_n10state"); ok {
		x := (v.(string))
		o.SetSlotN10state(x)
	}
	if v, ok := d.GetOk("slot_n11state"); ok {
		x := (v.(string))
		o.SetSlotN11state(x)
	}
	if v, ok := d.GetOk("slot_n12state"); ok {
		x := (v.(string))
		o.SetSlotN12state(x)
	}
	if v, ok := d.GetOk("slot_n13state"); ok {
		x := (v.(string))
		o.SetSlotN13state(x)
	}
	if v, ok := d.GetOk("slot_n14state"); ok {
		x := (v.(string))
		o.SetSlotN14state(x)
	}
	if v, ok := d.GetOk("slot_n15state"); ok {
		x := (v.(string))
		o.SetSlotN15state(x)
	}
	if v, ok := d.GetOk("slot_n16state"); ok {
		x := (v.(string))
		o.SetSlotN16state(x)
	}
	if v, ok := d.GetOk("slot_n17state"); ok {
		x := (v.(string))
		o.SetSlotN17state(x)
	}
	if v, ok := d.GetOk("slot_n18state"); ok {
		x := (v.(string))
		o.SetSlotN18state(x)
	}
	if v, ok := d.GetOk("slot_n19state"); ok {
		x := (v.(string))
		o.SetSlotN19state(x)
	}
	if v, ok := d.GetOk("slot_n1state"); ok {
		x := (v.(string))
		o.SetSlotN1state(x)
	}
	if v, ok := d.GetOk("slot_n20state"); ok {
		x := (v.(string))
		o.SetSlotN20state(x)
	}
	if v, ok := d.GetOk("slot_n21state"); ok {
		x := (v.(string))
		o.SetSlotN21state(x)
	}
	if v, ok := d.GetOk("slot_n22state"); ok {
		x := (v.(string))
		o.SetSlotN22state(x)
	}
	if v, ok := d.GetOk("slot_n23state"); ok {
		x := (v.(string))
		o.SetSlotN23state(x)
	}
	if v, ok := d.GetOk("slot_n24state"); ok {
		x := (v.(string))
		o.SetSlotN24state(x)
	}
	if v, ok := d.GetOk("slot_n2state"); ok {
		x := (v.(string))
		o.SetSlotN2state(x)
	}
	if v, ok := d.GetOk("slot_n3state"); ok {
		x := (v.(string))
		o.SetSlotN3state(x)
	}
	if v, ok := d.GetOk("slot_n4state"); ok {
		x := (v.(string))
		o.SetSlotN4state(x)
	}
	if v, ok := d.GetOk("slot_n5state"); ok {
		x := (v.(string))
		o.SetSlotN5state(x)
	}
	if v, ok := d.GetOk("slot_n6state"); ok {
		x := (v.(string))
		o.SetSlotN6state(x)
	}
	if v, ok := d.GetOk("slot_n7state"); ok {
		x := (v.(string))
		o.SetSlotN7state(x)
	}
	if v, ok := d.GetOk("slot_n8state"); ok {
		x := (v.(string))
		o.SetSlotN8state(x)
	}
	if v, ok := d.GetOk("slot_n9state"); ok {
		x := (v.(string))
		o.SetSlotN9state(x)
	}
	if v, ok := d.GetOk("slot_raid_link_speed"); ok {
		x := (v.(string))
		o.SetSlotRaidLinkSpeed(x)
	}
	if v, ok := d.GetOk("slot_raid_state"); ok {
		x := (v.(string))
		o.SetSlotRaidState(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme1state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme1state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRearNvme2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme2state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme2state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme3state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme3state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme4state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme4state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme5state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme5state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme6state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme6state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme7state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme7state(x)
	}
	if v, ok := d.GetOk("slot_rear_nvme8state"); ok {
		x := (v.(string))
		o.SetSlotRearNvme8state(x)
	}
	if v, ok := d.GetOk("slot_riser1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser1slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser1slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser1slot3link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser1slot3linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2slot4link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot4linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2slot5link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot5linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_riser2slot6link_speed"); ok {
		x := (v.(string))
		o.SetSlotRiser2slot6linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_sas_state"); ok {
		x := (v.(string))
		o.SetSlotSasState(x)
	}
	if v, ok := d.GetOk("slot_ssd_slot1link_speed"); ok {
		x := (v.(string))
		o.SetSlotSsdSlot1linkSpeed(x)
	}
	if v, ok := d.GetOk("slot_ssd_slot2link_speed"); ok {
		x := (v.(string))
		o.SetSlotSsdSlot2linkSpeed(x)
	}
	if v, ok := d.GetOk("smee"); ok {
		x := (v.(string))
		o.SetSmee(x)
	}
	if v, ok := d.GetOk("smt_mode"); ok {
		x := (v.(string))
		o.SetSmtMode(x)
	}
	if v, ok := d.GetOk("snc"); ok {
		x := (v.(string))
		o.SetSnc(x)
	}
	if v, ok := d.GetOk("sparing_mode"); ok {
		x := (v.(string))
		o.SetSparingMode(x)
	}
	if v, ok := d.GetOk("sr_iov"); ok {
		x := (v.(string))
		o.SetSrIov(x)
	}
	if v, ok := d.GetOk("streamer_prefetch"); ok {
		x := (v.(string))
		o.SetStreamerPrefetch(x)
	}
	if v, ok := d.GetOk("svm_mode"); ok {
		x := (v.(string))
		o.SetSvmMode(x)
	}
	if v, ok := d.GetOk("terminal_type"); ok {
		x := (v.(string))
		o.SetTerminalType(x)
	}
	if v, ok := d.GetOk("tpm_control"); ok {
		x := (v.(string))
		o.SetTpmControl(x)
	}
	if v, ok := d.GetOk("tpm_support"); ok {
		x := (v.(string))
		o.SetTpmSupport(x)
	}
	if v, ok := d.GetOk("txt_support"); ok {
		x := (v.(string))
		o.SetTxtSupport(x)
	}
	if v, ok := d.GetOk("ucsm_boot_order_rule"); ok {
		x := (v.(string))
		o.SetUcsmBootOrderRule(x)
	}
	if v, ok := d.GetOk("usb_emul6064"); ok {
		x := (v.(string))
		o.SetUsbEmul6064(x)
	}
	if v, ok := d.GetOk("usb_port_front"); ok {
		x := (v.(string))
		o.SetUsbPortFront(x)
	}
	if v, ok := d.GetOk("usb_port_internal"); ok {
		x := (v.(string))
		o.SetUsbPortInternal(x)
	}
	if v, ok := d.GetOk("usb_port_kvm"); ok {
		x := (v.(string))
		o.SetUsbPortKvm(x)
	}
	if v, ok := d.GetOk("usb_port_rear"); ok {
		x := (v.(string))
		o.SetUsbPortRear(x)
	}
	if v, ok := d.GetOk("usb_port_sd_card"); ok {
		x := (v.(string))
		o.SetUsbPortSdCard(x)
	}
	if v, ok := d.GetOk("usb_port_vmedia"); ok {
		x := (v.(string))
		o.SetUsbPortVmedia(x)
	}
	if v, ok := d.GetOk("usb_xhci_support"); ok {
		x := (v.(string))
		o.SetUsbXhciSupport(x)
	}
	if v, ok := d.GetOk("vga_priority"); ok {
		x := (v.(string))
		o.SetVgaPriority(x)
	}
	if v, ok := d.GetOk("vmd_enable"); ok {
		x := (v.(string))
		o.SetVmdEnable(x)
	}
	if v, ok := d.GetOk("work_load_config"); ok {
		x := (v.(string))
		o.SetWorkLoadConfig(x)
	}
	if v, ok := d.GetOk("xpt_prefetch"); ok {
		x := (v.(string))
		o.SetXptPrefetch(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Errorf("Json Marshalling of data source failed with error : %+v", err)
	}
	result, _, err := conn.ApiClient.BiosApi.GetBiosPolicyList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if err != nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = models.NewBiosPolicy()
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu1state", (s.AcsControlGpu1state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu2state", (s.AcsControlGpu2state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu3state", (s.AcsControlGpu3state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu4state", (s.AcsControlGpu4state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu5state", (s.AcsControlGpu5state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu6state", (s.AcsControlGpu6state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu7state", (s.AcsControlGpu7state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_gpu8state", (s.AcsControlGpu8state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_slot11state", (s.AcsControlSlot11state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_slot12state", (s.AcsControlSlot12state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_slot13state", (s.AcsControlSlot13state)); err != nil {
				return err
			}
			if err := d.Set("acs_control_slot14state", (s.AcsControlSlot14state)); err != nil {
				return err
			}
			if err := d.Set("adjacent_cache_line_prefetch", (s.AdjacentCacheLinePrefetch)); err != nil {
				return err
			}
			if err := d.Set("all_usb_devices", (s.AllUsbDevices)); err != nil {
				return err
			}
			if err := d.Set("altitude", (s.Altitude)); err != nil {
				return err
			}
			if err := d.Set("aspm_support", (s.AspmSupport)); err != nil {
				return err
			}
			if err := d.Set("assert_nmi_on_perr", (s.AssertNmiOnPerr)); err != nil {
				return err
			}
			if err := d.Set("assert_nmi_on_serr", (s.AssertNmiOnSerr)); err != nil {
				return err
			}
			if err := d.Set("auto_cc_state", (s.AutoCcState)); err != nil {
				return err
			}
			if err := d.Set("autonumous_cstate_enable", (s.AutonumousCstateEnable)); err != nil {
				return err
			}
			if err := d.Set("baud_rate", (s.BaudRate)); err != nil {
				return err
			}
			if err := d.Set("bme_dma_mitigation", (s.BmeDmaMitigation)); err != nil {
				return err
			}
			if err := d.Set("boot_option_num_retry", (s.BootOptionNumRetry)); err != nil {
				return err
			}
			if err := d.Set("boot_option_re_cool_down", (s.BootOptionReCoolDown)); err != nil {
				return err
			}
			if err := d.Set("boot_option_retry", (s.BootOptionRetry)); err != nil {
				return err
			}
			if err := d.Set("boot_performance_mode", (s.BootPerformanceMode)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_cpu_cpb", (s.CbsCmnCpuCpb)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_cpu_gen_downcore_ctrl", (s.CbsCmnCpuGenDowncoreCtrl)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_cpu_global_cstate_ctrl", (s.CbsCmnCpuGlobalCstateCtrl)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_cpu_l1stream_hw_prefetcher", (s.CbsCmnCpuL1streamHwPrefetcher)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_cpu_l2stream_hw_prefetcher", (s.CbsCmnCpuL2streamHwPrefetcher)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_determinism_slider", (s.CbsCmnDeterminismSlider)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_gnb_nb_iommu", (s.CbsCmnGnbNbIommu)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_mem_ctrl_bank_group_swap_ddr4", (s.CbsCmnMemCtrlBankGroupSwapDdr4)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmn_mem_map_bank_interleave_ddr4", (s.CbsCmnMemMapBankInterleaveDdr4)); err != nil {
				return err
			}
			if err := d.Set("cbs_cmnc_tdp_ctl", (s.CbsCmncTdpCtl)); err != nil {
				return err
			}
			if err := d.Set("cbs_df_cmn_mem_intlv", (s.CbsDfCmnMemIntlv)); err != nil {
				return err
			}
			if err := d.Set("cbs_df_cmn_mem_intlv_size", (s.CbsDfCmnMemIntlvSize)); err != nil {
				return err
			}
			if err := d.Set("cdn_enable", (s.CdnEnable)); err != nil {
				return err
			}
			if err := d.Set("cdn_support", (s.CdnSupport)); err != nil {
				return err
			}
			if err := d.Set("channel_inter_leave", (s.ChannelInterLeave)); err != nil {
				return err
			}
			if err := d.Set("cisco_adaptive_mem_training", (s.CiscoAdaptiveMemTraining)); err != nil {
				return err
			}
			if err := d.Set("cisco_debug_level", (s.CiscoDebugLevel)); err != nil {
				return err
			}
			if err := d.Set("cisco_oprom_launch_optimization", (s.CiscoOpromLaunchOptimization)); err != nil {
				return err
			}
			if err := d.Set("cke_low_policy", (s.CkeLowPolicy)); err != nil {
				return err
			}
			if err := d.Set("class_id", (s.ClassId)); err != nil {
				return err
			}
			if err := d.Set("closed_loop_therm_throtl", (s.ClosedLoopThermThrotl)); err != nil {
				return err
			}
			if err := d.Set("cmci_enable", (s.CmciEnable)); err != nil {
				return err
			}
			if err := d.Set("config_tdp", (s.ConfigTdp)); err != nil {
				return err
			}
			if err := d.Set("console_redirection", (s.ConsoleRedirection)); err != nil {
				return err
			}
			if err := d.Set("core_multi_processing", (s.CoreMultiProcessing)); err != nil {
				return err
			}
			if err := d.Set("cpu_energy_performance", (s.CpuEnergyPerformance)); err != nil {
				return err
			}
			if err := d.Set("cpu_frequency_floor", (s.CpuFrequencyFloor)); err != nil {
				return err
			}
			if err := d.Set("cpu_performance", (s.CpuPerformance)); err != nil {
				return err
			}
			if err := d.Set("cpu_power_management", (s.CpuPowerManagement)); err != nil {
				return err
			}
			if err := d.Set("dcpmm_firmware_downgrade", (s.DcpmmFirmwareDowngrade)); err != nil {
				return err
			}
			if err := d.Set("demand_scrub", (s.DemandScrub)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("direct_cache_access", (s.DirectCacheAccess)); err != nil {
				return err
			}
			if err := d.Set("dram_clock_throttling", (s.DramClockThrottling)); err != nil {
				return err
			}
			if err := d.Set("dram_refresh_rate", (s.DramRefreshRate)); err != nil {
				return err
			}
			if err := d.Set("energy_efficient_turbo", (s.EnergyEfficientTurbo)); err != nil {
				return err
			}
			if err := d.Set("eng_perf_tuning", (s.EngPerfTuning)); err != nil {
				return err
			}
			if err := d.Set("enhanced_intel_speed_step_tech", (s.EnhancedIntelSpeedStepTech)); err != nil {
				return err
			}
			if err := d.Set("epp_profile", (s.EppProfile)); err != nil {
				return err
			}
			if err := d.Set("execute_disable_bit", (s.ExecuteDisableBit)); err != nil {
				return err
			}
			if err := d.Set("extended_apic", (s.ExtendedApic)); err != nil {
				return err
			}
			if err := d.Set("flow_control", (s.FlowControl)); err != nil {
				return err
			}
			if err := d.Set("frb2enable", (s.Frb2enable)); err != nil {
				return err
			}
			if err := d.Set("hardware_prefetch", (s.HardwarePrefetch)); err != nil {
				return err
			}
			if err := d.Set("hwpm_enable", (s.HwpmEnable)); err != nil {
				return err
			}
			if err := d.Set("imc_interleave", (s.ImcInterleave)); err != nil {
				return err
			}
			if err := d.Set("intel_hyper_threading_tech", (s.IntelHyperThreadingTech)); err != nil {
				return err
			}
			if err := d.Set("intel_speed_select", (s.IntelSpeedSelect)); err != nil {
				return err
			}
			if err := d.Set("intel_turbo_boost_tech", (s.IntelTurboBoostTech)); err != nil {
				return err
			}
			if err := d.Set("intel_virtualization_technology", (s.IntelVirtualizationTechnology)); err != nil {
				return err
			}
			if err := d.Set("intel_vt_for_directed_io", (s.IntelVtForDirectedIo)); err != nil {
				return err
			}
			if err := d.Set("intel_vtd_coherency_support", (s.IntelVtdCoherencySupport)); err != nil {
				return err
			}
			if err := d.Set("intel_vtd_interrupt_remapping", (s.IntelVtdInterruptRemapping)); err != nil {
				return err
			}
			if err := d.Set("intel_vtd_pass_through_dma_support", (s.IntelVtdPassThroughDmaSupport)); err != nil {
				return err
			}
			if err := d.Set("intel_vtdats_support", (s.IntelVtdatsSupport)); err != nil {
				return err
			}
			if err := d.Set("ioh_error_enable", (s.IohErrorEnable)); err != nil {
				return err
			}
			if err := d.Set("ioh_resource", (s.IohResource)); err != nil {
				return err
			}
			if err := d.Set("ip_prefetch", (s.IpPrefetch)); err != nil {
				return err
			}
			if err := d.Set("ipv4pxe", (s.Ipv4pxe)); err != nil {
				return err
			}
			if err := d.Set("ipv6pxe", (s.Ipv6pxe)); err != nil {
				return err
			}
			if err := d.Set("kti_prefetch", (s.KtiPrefetch)); err != nil {
				return err
			}
			if err := d.Set("legacy_os_redirection", (s.LegacyOsRedirection)); err != nil {
				return err
			}
			if err := d.Set("legacy_usb_support", (s.LegacyUsbSupport)); err != nil {
				return err
			}
			if err := d.Set("llc_prefetch", (s.LlcPrefetch)); err != nil {
				return err
			}
			if err := d.Set("lom_port0state", (s.LomPort0state)); err != nil {
				return err
			}
			if err := d.Set("lom_port1state", (s.LomPort1state)); err != nil {
				return err
			}
			if err := d.Set("lom_port2state", (s.LomPort2state)); err != nil {
				return err
			}
			if err := d.Set("lom_port3state", (s.LomPort3state)); err != nil {
				return err
			}
			if err := d.Set("lom_ports_all_state", (s.LomPortsAllState)); err != nil {
				return err
			}
			if err := d.Set("lv_ddr_mode", (s.LvDdrMode)); err != nil {
				return err
			}
			if err := d.Set("make_device_non_bootable", (s.MakeDeviceNonBootable)); err != nil {
				return err
			}
			if err := d.Set("memory_inter_leave", (s.MemoryInterLeave)); err != nil {
				return err
			}
			if err := d.Set("memory_mapped_io_above4gb", (s.MemoryMappedIoAbove4gb)); err != nil {
				return err
			}
			if err := d.Set("memory_size_limit", (s.MemorySizeLimit)); err != nil {
				return err
			}
			if err := d.Set("mirroring_mode", (s.MirroringMode)); err != nil {
				return err
			}
			if err := d.Set("mmcfg_base", (s.MmcfgBase)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("name", (s.Name)); err != nil {
				return err
			}
			if err := d.Set("network_stack", (s.NetworkStack)); err != nil {
				return err
			}
			if err := d.Set("numa_optimized", (s.NumaOptimized)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("onboard10gbit_lom", (s.Onboard10gbitLom)); err != nil {
				return err
			}
			if err := d.Set("onboard_gbit_lom", (s.OnboardGbitLom)); err != nil {
				return err
			}
			if err := d.Set("onboard_scu_storage_support", (s.OnboardScuStorageSupport)); err != nil {
				return err
			}
			if err := d.Set("onboard_scu_storage_sw_stack", (s.OnboardScuStorageSwStack)); err != nil {
				return err
			}

			if err := d.Set("organization", flattenMapOrganizationOrganizationRelationship(s.Organization, d)); err != nil {
				return err
			}
			if err := d.Set("os_boot_watchdog_timer", (s.OsBootWatchdogTimer)); err != nil {
				return err
			}
			if err := d.Set("os_boot_watchdog_timer_policy", (s.OsBootWatchdogTimerPolicy)); err != nil {
				return err
			}
			if err := d.Set("os_boot_watchdog_timer_timeout", (s.OsBootWatchdogTimerTimeout)); err != nil {
				return err
			}
			if err := d.Set("out_of_band_mgmt_port", (s.OutOfBandMgmtPort)); err != nil {
				return err
			}
			if err := d.Set("package_cstate_limit", (s.PackageCstateLimit)); err != nil {
				return err
			}
			if err := d.Set("partial_mirror_mode_config", (s.PartialMirrorModeConfig)); err != nil {
				return err
			}
			if err := d.Set("partial_mirror_percent", (s.PartialMirrorPercent)); err != nil {
				return err
			}
			if err := d.Set("partial_mirror_value1", (s.PartialMirrorValue1)); err != nil {
				return err
			}
			if err := d.Set("partial_mirror_value2", (s.PartialMirrorValue2)); err != nil {
				return err
			}
			if err := d.Set("partial_mirror_value3", (s.PartialMirrorValue3)); err != nil {
				return err
			}
			if err := d.Set("partial_mirror_value4", (s.PartialMirrorValue4)); err != nil {
				return err
			}
			if err := d.Set("patrol_scrub", (s.PatrolScrub)); err != nil {
				return err
			}
			if err := d.Set("patrol_scrub_duration", (s.PatrolScrubDuration)); err != nil {
				return err
			}
			if err := d.Set("pc_ie_ras_support", (s.PcIeRasSupport)); err != nil {
				return err
			}
			if err := d.Set("pc_ie_ssd_hot_plug_support", (s.PcIeSsdHotPlugSupport)); err != nil {
				return err
			}
			if err := d.Set("pch_usb30mode", (s.PchUsb30mode)); err != nil {
				return err
			}
			if err := d.Set("pci_option_ro_ms", (s.PciOptionRoMs)); err != nil {
				return err
			}
			if err := d.Set("pci_rom_clp", (s.PciRomClp)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRelationship(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("pop_support", (s.PopSupport)); err != nil {
				return err
			}
			if err := d.Set("post_error_pause", (s.PostErrorPause)); err != nil {
				return err
			}
			if err := d.Set("processor_c1e", (s.ProcessorC1e)); err != nil {
				return err
			}
			if err := d.Set("processor_c3report", (s.ProcessorC3report)); err != nil {
				return err
			}
			if err := d.Set("processor_c6report", (s.ProcessorC6report)); err != nil {
				return err
			}
			if err := d.Set("processor_cstate", (s.ProcessorCstate)); err != nil {
				return err
			}

			if err := d.Set("profiles", flattenListPolicyAbstractConfigProfileRelationship(s.Profiles, d)); err != nil {
				return err
			}
			if err := d.Set("psata", (s.Psata)); err != nil {
				return err
			}
			if err := d.Set("pstate_coord_type", (s.PstateCoordType)); err != nil {
				return err
			}
			if err := d.Set("putty_key_pad", (s.PuttyKeyPad)); err != nil {
				return err
			}
			if err := d.Set("pwr_perf_tuning", (s.PwrPerfTuning)); err != nil {
				return err
			}
			if err := d.Set("qpi_link_frequency", (s.QpiLinkFrequency)); err != nil {
				return err
			}
			if err := d.Set("qpi_snoop_mode", (s.QpiSnoopMode)); err != nil {
				return err
			}
			if err := d.Set("rank_inter_leave", (s.RankInterLeave)); err != nil {
				return err
			}
			if err := d.Set("redirection_after_post", (s.RedirectionAfterPost)); err != nil {
				return err
			}
			if err := d.Set("sata_mode_select", (s.SataModeSelect)); err != nil {
				return err
			}
			if err := d.Set("select_memory_ras_configuration", (s.SelectMemoryRasConfiguration)); err != nil {
				return err
			}
			if err := d.Set("select_ppr_type", (s.SelectPprType)); err != nil {
				return err
			}
			if err := d.Set("serial_port_aenable", (s.SerialPortAenable)); err != nil {
				return err
			}
			if err := d.Set("single_pctl_enable", (s.SinglePctlEnable)); err != nil {
				return err
			}
			if err := d.Set("slot10link_speed", (s.Slot10linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot10state", (s.Slot10state)); err != nil {
				return err
			}
			if err := d.Set("slot11link_speed", (s.Slot11linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot11state", (s.Slot11state)); err != nil {
				return err
			}
			if err := d.Set("slot12link_speed", (s.Slot12linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot12state", (s.Slot12state)); err != nil {
				return err
			}
			if err := d.Set("slot13state", (s.Slot13state)); err != nil {
				return err
			}
			if err := d.Set("slot14state", (s.Slot14state)); err != nil {
				return err
			}
			if err := d.Set("slot1link_speed", (s.Slot1linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot1state", (s.Slot1state)); err != nil {
				return err
			}
			if err := d.Set("slot2link_speed", (s.Slot2linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot2state", (s.Slot2state)); err != nil {
				return err
			}
			if err := d.Set("slot3link_speed", (s.Slot3linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot3state", (s.Slot3state)); err != nil {
				return err
			}
			if err := d.Set("slot4link_speed", (s.Slot4linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot4state", (s.Slot4state)); err != nil {
				return err
			}
			if err := d.Set("slot5link_speed", (s.Slot5linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot5state", (s.Slot5state)); err != nil {
				return err
			}
			if err := d.Set("slot6link_speed", (s.Slot6linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot6state", (s.Slot6state)); err != nil {
				return err
			}
			if err := d.Set("slot7link_speed", (s.Slot7linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot7state", (s.Slot7state)); err != nil {
				return err
			}
			if err := d.Set("slot8link_speed", (s.Slot8linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot8state", (s.Slot8state)); err != nil {
				return err
			}
			if err := d.Set("slot9link_speed", (s.Slot9linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot9state", (s.Slot9state)); err != nil {
				return err
			}
			if err := d.Set("slot_flom_link_speed", (s.SlotFlomLinkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_front_nvme1link_speed", (s.SlotFrontNvme1linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_front_nvme2link_speed", (s.SlotFrontNvme2linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_front_slot5link_speed", (s.SlotFrontSlot5linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_front_slot6link_speed", (s.SlotFrontSlot6linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu1state", (s.SlotGpu1state)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu2state", (s.SlotGpu2state)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu3state", (s.SlotGpu3state)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu4state", (s.SlotGpu4state)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu5state", (s.SlotGpu5state)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu6state", (s.SlotGpu6state)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu7state", (s.SlotGpu7state)); err != nil {
				return err
			}
			if err := d.Set("slot_gpu8state", (s.SlotGpu8state)); err != nil {
				return err
			}
			if err := d.Set("slot_hba_link_speed", (s.SlotHbaLinkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_hba_state", (s.SlotHbaState)); err != nil {
				return err
			}
			if err := d.Set("slot_lom1link", (s.SlotLom1link)); err != nil {
				return err
			}
			if err := d.Set("slot_lom2link", (s.SlotLom2link)); err != nil {
				return err
			}
			if err := d.Set("slot_mezz_state", (s.SlotMezzState)); err != nil {
				return err
			}
			if err := d.Set("slot_mlom_link_speed", (s.SlotMlomLinkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_mlom_state", (s.SlotMlomState)); err != nil {
				return err
			}
			if err := d.Set("slot_mraid_link_speed", (s.SlotMraidLinkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_mraid_state", (s.SlotMraidState)); err != nil {
				return err
			}
			if err := d.Set("slot_n10state", (s.SlotN10state)); err != nil {
				return err
			}
			if err := d.Set("slot_n11state", (s.SlotN11state)); err != nil {
				return err
			}
			if err := d.Set("slot_n12state", (s.SlotN12state)); err != nil {
				return err
			}
			if err := d.Set("slot_n13state", (s.SlotN13state)); err != nil {
				return err
			}
			if err := d.Set("slot_n14state", (s.SlotN14state)); err != nil {
				return err
			}
			if err := d.Set("slot_n15state", (s.SlotN15state)); err != nil {
				return err
			}
			if err := d.Set("slot_n16state", (s.SlotN16state)); err != nil {
				return err
			}
			if err := d.Set("slot_n17state", (s.SlotN17state)); err != nil {
				return err
			}
			if err := d.Set("slot_n18state", (s.SlotN18state)); err != nil {
				return err
			}
			if err := d.Set("slot_n19state", (s.SlotN19state)); err != nil {
				return err
			}
			if err := d.Set("slot_n1state", (s.SlotN1state)); err != nil {
				return err
			}
			if err := d.Set("slot_n20state", (s.SlotN20state)); err != nil {
				return err
			}
			if err := d.Set("slot_n21state", (s.SlotN21state)); err != nil {
				return err
			}
			if err := d.Set("slot_n22state", (s.SlotN22state)); err != nil {
				return err
			}
			if err := d.Set("slot_n23state", (s.SlotN23state)); err != nil {
				return err
			}
			if err := d.Set("slot_n24state", (s.SlotN24state)); err != nil {
				return err
			}
			if err := d.Set("slot_n2state", (s.SlotN2state)); err != nil {
				return err
			}
			if err := d.Set("slot_n3state", (s.SlotN3state)); err != nil {
				return err
			}
			if err := d.Set("slot_n4state", (s.SlotN4state)); err != nil {
				return err
			}
			if err := d.Set("slot_n5state", (s.SlotN5state)); err != nil {
				return err
			}
			if err := d.Set("slot_n6state", (s.SlotN6state)); err != nil {
				return err
			}
			if err := d.Set("slot_n7state", (s.SlotN7state)); err != nil {
				return err
			}
			if err := d.Set("slot_n8state", (s.SlotN8state)); err != nil {
				return err
			}
			if err := d.Set("slot_n9state", (s.SlotN9state)); err != nil {
				return err
			}
			if err := d.Set("slot_raid_link_speed", (s.SlotRaidLinkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_raid_state", (s.SlotRaidState)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme1link_speed", (s.SlotRearNvme1linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme1state", (s.SlotRearNvme1state)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme2link_speed", (s.SlotRearNvme2linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme2state", (s.SlotRearNvme2state)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme3state", (s.SlotRearNvme3state)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme4state", (s.SlotRearNvme4state)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme5state", (s.SlotRearNvme5state)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme6state", (s.SlotRearNvme6state)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme7state", (s.SlotRearNvme7state)); err != nil {
				return err
			}
			if err := d.Set("slot_rear_nvme8state", (s.SlotRearNvme8state)); err != nil {
				return err
			}
			if err := d.Set("slot_riser1link_speed", (s.SlotRiser1linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_riser1slot1link_speed", (s.SlotRiser1slot1linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_riser1slot2link_speed", (s.SlotRiser1slot2linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_riser1slot3link_speed", (s.SlotRiser1slot3linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_riser2link_speed", (s.SlotRiser2linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_riser2slot4link_speed", (s.SlotRiser2slot4linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_riser2slot5link_speed", (s.SlotRiser2slot5linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_riser2slot6link_speed", (s.SlotRiser2slot6linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_sas_state", (s.SlotSasState)); err != nil {
				return err
			}
			if err := d.Set("slot_ssd_slot1link_speed", (s.SlotSsdSlot1linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("slot_ssd_slot2link_speed", (s.SlotSsdSlot2linkSpeed)); err != nil {
				return err
			}
			if err := d.Set("smee", (s.Smee)); err != nil {
				return err
			}
			if err := d.Set("smt_mode", (s.SmtMode)); err != nil {
				return err
			}
			if err := d.Set("snc", (s.Snc)); err != nil {
				return err
			}
			if err := d.Set("sparing_mode", (s.SparingMode)); err != nil {
				return err
			}
			if err := d.Set("sr_iov", (s.SrIov)); err != nil {
				return err
			}
			if err := d.Set("streamer_prefetch", (s.StreamerPrefetch)); err != nil {
				return err
			}
			if err := d.Set("svm_mode", (s.SvmMode)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			if err := d.Set("terminal_type", (s.TerminalType)); err != nil {
				return err
			}
			if err := d.Set("tpm_control", (s.TpmControl)); err != nil {
				return err
			}
			if err := d.Set("tpm_support", (s.TpmSupport)); err != nil {
				return err
			}
			if err := d.Set("txt_support", (s.TxtSupport)); err != nil {
				return err
			}
			if err := d.Set("ucsm_boot_order_rule", (s.UcsmBootOrderRule)); err != nil {
				return err
			}
			if err := d.Set("usb_emul6064", (s.UsbEmul6064)); err != nil {
				return err
			}
			if err := d.Set("usb_port_front", (s.UsbPortFront)); err != nil {
				return err
			}
			if err := d.Set("usb_port_internal", (s.UsbPortInternal)); err != nil {
				return err
			}
			if err := d.Set("usb_port_kvm", (s.UsbPortKvm)); err != nil {
				return err
			}
			if err := d.Set("usb_port_rear", (s.UsbPortRear)); err != nil {
				return err
			}
			if err := d.Set("usb_port_sd_card", (s.UsbPortSdCard)); err != nil {
				return err
			}
			if err := d.Set("usb_port_vmedia", (s.UsbPortVmedia)); err != nil {
				return err
			}
			if err := d.Set("usb_xhci_support", (s.UsbXhciSupport)); err != nil {
				return err
			}
			if err := d.Set("vga_priority", (s.VgaPriority)); err != nil {
				return err
			}
			if err := d.Set("vmd_enable", (s.VmdEnable)); err != nil {
				return err
			}
			if err := d.Set("work_load_config", (s.WorkLoadConfig)); err != nil {
				return err
			}
			if err := d.Set("xpt_prefetch", (s.XptPrefetch)); err != nil {
				return err
			}
			d.SetId(s.GetMoid())
		}
	}
	return nil
}
