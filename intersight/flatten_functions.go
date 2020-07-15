package intersight

import (
	"encoding/json"
	"log"
	"reflect"

	models "github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenAdditionalProperties(m map[string]interface{}) string {
	var s string
	if len(m) > 0 {
		j, err := json.Marshal(m)
		if err != nil {
			log.Printf("Error occurred while flattening and json parsing: %s", err)
		} else {
			s = string(j)
		}
	}
	return s
}
func flattenMoMoRef(ref *models.MoMoRef) map[string]interface{} {
	x := make(map[string]interface{})
	x["additional_properties"] = flattenAdditionalProperties(ref.AdditionalProperties)
	x["class_id"] = ref.ClassId
	x["moid"] = ref.Moid
	x["object_type"] = ref.ObjectType
	return x
}

func flattenListAdapterAdapterConfig(p []models.AdapterAdapterConfig, d *schema.ResourceData) []map[string]interface{} {
	var adapteradapterconfigs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		adapteradapterconfig := make(map[string]interface{})
		adapteradapterconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		adapteradapterconfig["class_id"] = item.ClassId
		adapteradapterconfig["dce_interface_settings"] = (func(p []models.AdapterDceInterfaceSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterdceinterfacesettingss []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				adapterdceinterfacesettings := make(map[string]interface{})
				adapterdceinterfacesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				adapterdceinterfacesettings["class_id"] = item.ClassId
				adapterdceinterfacesettings["fec_mode"] = item.FecMode
				adapterdceinterfacesettings["interface_id"] = item.InterfaceId
				adapterdceinterfacesettings["object_type"] = item.ObjectType
				adapterdceinterfacesettingss = append(adapterdceinterfacesettingss, adapterdceinterfacesettings)
			}
			return adapterdceinterfacesettingss
		})(item.GetDceInterfaceSettings(), d)
		adapteradapterconfig["eth_settings"] = (func(p models.AdapterEthSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterethsettingss []map[string]interface{}
			var ret models.AdapterEthSettings
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			adapterethsettings := make(map[string]interface{})
			adapterethsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			adapterethsettings["class_id"] = item.ClassId
			adapterethsettings["lldp_enabled"] = item.LldpEnabled
			adapterethsettings["object_type"] = item.ObjectType

			adapterethsettingss = append(adapterethsettingss, adapterethsettings)
			return adapterethsettingss
		})(item.GetEthSettings(), d)
		adapteradapterconfig["fc_settings"] = (func(p models.AdapterFcSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterfcsettingss []map[string]interface{}
			var ret models.AdapterFcSettings
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			adapterfcsettings := make(map[string]interface{})
			adapterfcsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			adapterfcsettings["class_id"] = item.ClassId
			adapterfcsettings["fip_enabled"] = item.FipEnabled
			adapterfcsettings["object_type"] = item.ObjectType

			adapterfcsettingss = append(adapterfcsettingss, adapterfcsettings)
			return adapterfcsettingss
		})(item.GetFcSettings(), d)
		adapteradapterconfig["object_type"] = item.ObjectType
		adapteradapterconfig["port_channel_settings"] = (func(p models.AdapterPortChannelSettings, d *schema.ResourceData) []map[string]interface{} {
			var adapterportchannelsettingss []map[string]interface{}
			var ret models.AdapterPortChannelSettings
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			adapterportchannelsettings := make(map[string]interface{})
			adapterportchannelsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			adapterportchannelsettings["class_id"] = item.ClassId
			adapterportchannelsettings["enabled"] = item.Enabled
			adapterportchannelsettings["object_type"] = item.ObjectType

			adapterportchannelsettingss = append(adapterportchannelsettingss, adapterportchannelsettings)
			return adapterportchannelsettingss
		})(item.GetPortChannelSettings(), d)
		adapteradapterconfig["slot_id"] = item.SlotId
		adapteradapterconfigs = append(adapteradapterconfigs, adapteradapterconfig)
	}
	return adapteradapterconfigs
}
func flattenListAdapterExtEthInterfaceRelationship(p []models.AdapterExtEthInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterextethinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterextethinterfacerelationship := flattenMoMoRef(item)
		adapterextethinterfacerelationships = append(adapterextethinterfacerelationships, adapterextethinterfacerelationship)
	}
	return adapterextethinterfacerelationships
}
func flattenListAdapterHostEthInterfaceRelationship(p []models.AdapterHostEthInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostethinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterhostethinterfacerelationship := flattenMoMoRef(item)
		adapterhostethinterfacerelationships = append(adapterhostethinterfacerelationships, adapterhostethinterfacerelationship)
	}
	return adapterhostethinterfacerelationships
}
func flattenListAdapterHostFcInterfaceRelationship(p []models.AdapterHostFcInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostfcinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterhostfcinterfacerelationship := flattenMoMoRef(item)
		adapterhostfcinterfacerelationships = append(adapterhostfcinterfacerelationships, adapterhostfcinterfacerelationship)
	}
	return adapterhostfcinterfacerelationships
}
func flattenListAdapterHostIscsiInterfaceRelationship(p []models.AdapterHostIscsiInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterhostiscsiinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterhostiscsiinterfacerelationship := flattenMoMoRef(item)
		adapterhostiscsiinterfacerelationships = append(adapterhostiscsiinterfacerelationships, adapterhostiscsiinterfacerelationship)
	}
	return adapterhostiscsiinterfacerelationships
}
func flattenListAdapterUnitRelationship(p []models.AdapterUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		adapterunitrelationship := flattenMoMoRef(item)
		adapterunitrelationships = append(adapterunitrelationships, adapterunitrelationship)
	}
	return adapterunitrelationships
}
func flattenListApplianceDataExportPolicyRelationship(p []models.ApplianceDataExportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		appliancedataexportpolicyrelationship := flattenMoMoRef(item)
		appliancedataexportpolicyrelationships = append(appliancedataexportpolicyrelationships, appliancedataexportpolicyrelationship)
	}
	return appliancedataexportpolicyrelationships
}
func flattenListApplianceKeyValuePair(p []models.ApplianceKeyValuePair, d *schema.ResourceData) []map[string]interface{} {
	var appliancekeyvaluepairs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		appliancekeyvaluepair := make(map[string]interface{})
		appliancekeyvaluepair["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		appliancekeyvaluepair["class_id"] = item.ClassId
		appliancekeyvaluepair["key"] = item.Key
		appliancekeyvaluepair["object_type"] = item.ObjectType
		appliancekeyvaluepair["value"] = item.Value
		appliancekeyvaluepairs = append(appliancekeyvaluepairs, appliancekeyvaluepair)
	}
	return appliancekeyvaluepairs
}
func flattenListAssetClusterMemberRelationship(p []models.AssetClusterMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		assetclustermemberrelationship := flattenMoMoRef(item)
		assetclustermemberrelationships = append(assetclustermemberrelationships, assetclustermemberrelationship)
	}
	return assetclustermemberrelationships
}
func flattenListAssetConnection(p []models.AssetConnection, d *schema.ResourceData) []map[string]interface{} {
	var assetconnections []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		assetconnection := make(map[string]interface{})
		assetconnection["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetconnection["class_id"] = item.ClassId
		assetconnection["object_type"] = item.ObjectType
		assetconnections = append(assetconnections, assetconnection)
	}
	return assetconnections
}
func flattenListAssetService(p []models.AssetService, d *schema.ResourceData) []map[string]interface{} {
	var assetservices []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		assetservice := make(map[string]interface{})
		assetservice["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetservice["class_id"] = item.ClassId
		assetservice["object_type"] = item.ObjectType
		assetservice["options"] = (func(p models.AssetServiceOptions, d *schema.ResourceData) []map[string]interface{} {
			var assetserviceoptionss []map[string]interface{}
			var ret models.AssetServiceOptions
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			assetserviceoptions := make(map[string]interface{})
			assetserviceoptions["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			assetserviceoptions["class_id"] = item.ClassId
			assetserviceoptions["object_type"] = item.ObjectType

			assetserviceoptionss = append(assetserviceoptionss, assetserviceoptions)
			return assetserviceoptionss
		})(item.GetOptions(), d)
		assetservice["status"] = item.Status
		assetservice["status_error_reason"] = item.StatusErrorReason
		assetservices = append(assetservices, assetservice)
	}
	return assetservices
}
func flattenListBiosBootDeviceRelationship(p []models.BiosBootDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosbootdevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		biosbootdevicerelationship := flattenMoMoRef(item)
		biosbootdevicerelationships = append(biosbootdevicerelationships, biosbootdevicerelationship)
	}
	return biosbootdevicerelationships
}
func flattenListBiosUnitRelationship(p []models.BiosUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		biosunitrelationship := flattenMoMoRef(item)
		biosunitrelationships = append(biosunitrelationships, biosunitrelationship)
	}
	return biosunitrelationships
}
func flattenListBootDeviceBase(p []models.BootDeviceBase, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		bootdevicebase := make(map[string]interface{})
		bootdevicebase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		bootdevicebase["class_id"] = item.ClassId
		bootdevicebase["enabled"] = item.Enabled
		bootdevicebase["name"] = item.Name
		bootdevicebase["object_type"] = item.ObjectType
		bootdevicebases = append(bootdevicebases, bootdevicebase)
	}
	return bootdevicebases
}
func flattenListCapabilityCapabilityRelationship(p []models.CapabilityCapabilityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var capabilitycapabilityrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		capabilitycapabilityrelationship := flattenMoMoRef(item)
		capabilitycapabilityrelationships = append(capabilitycapabilityrelationships, capabilitycapabilityrelationship)
	}
	return capabilitycapabilityrelationships
}
func flattenListCapabilityPortRange(p []models.CapabilityPortRange, d *schema.ResourceData) []map[string]interface{} {
	var capabilityportranges []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		capabilityportrange := make(map[string]interface{})
		capabilityportrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		capabilityportrange["class_id"] = item.ClassId
		capabilityportrange["end_port_id"] = item.EndPortId
		capabilityportrange["end_slot_id"] = item.EndSlotId
		capabilityportrange["object_type"] = item.ObjectType
		capabilityportrange["start_port_id"] = item.StartPortId
		capabilityportrange["start_slot_id"] = item.StartSlotId
		capabilityportranges = append(capabilityportranges, capabilityportrange)
	}
	return capabilityportranges
}
func flattenListComputeBladeRelationship(p []models.ComputeBladeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computebladerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		computebladerelationship := flattenMoMoRef(item)
		computebladerelationships = append(computebladerelationships, computebladerelationship)
	}
	return computebladerelationships
}
func flattenListComputeIpAddress(p []models.ComputeIpAddress, d *schema.ResourceData) []map[string]interface{} {
	var computeipaddresss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		computeipaddress := make(map[string]interface{})
		computeipaddress["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		computeipaddress["address"] = item.Address
		computeipaddress["category"] = item.Category
		computeipaddress["class_id"] = item.ClassId
		computeipaddress["default_gateway"] = item.DefaultGateway
		computeipaddress["dn"] = item.Dn
		computeipaddress["http_port"] = item.HttpPort
		computeipaddress["https_port"] = item.HttpsPort
		computeipaddress["kvm_port"] = item.KvmPort
		computeipaddress["kvm_vlan"] = item.KvmVlan
		computeipaddress["name"] = item.Name
		computeipaddress["object_type"] = item.ObjectType
		computeipaddress["subnet"] = item.Subnet
		computeipaddress["type"] = item.Type
		computeipaddresss = append(computeipaddresss, computeipaddress)
	}
	return computeipaddresss
}
func flattenListComputeRackUnitRelationship(p []models.ComputeRackUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		computerackunitrelationship := flattenMoMoRef(item)
		computerackunitrelationships = append(computerackunitrelationships, computerackunitrelationship)
	}
	return computerackunitrelationships
}
func flattenListCondHclStatusDetailRelationship(p []models.CondHclStatusDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusdetailrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		condhclstatusdetailrelationship := flattenMoMoRef(item)
		condhclstatusdetailrelationships = append(condhclstatusdetailrelationships, condhclstatusdetailrelationship)
	}
	return condhclstatusdetailrelationships
}
func flattenListConfigExportedItemRelationship(p []models.ConfigExportedItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configexporteditemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		configexporteditemrelationship := flattenMoMoRef(item)
		configexporteditemrelationships = append(configexporteditemrelationships, configexporteditemrelationship)
	}
	return configexporteditemrelationships
}
func flattenListConfigImportedItemRelationship(p []models.ConfigImportedItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configimporteditemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		configimporteditemrelationship := flattenMoMoRef(item)
		configimporteditemrelationships = append(configimporteditemrelationships, configimporteditemrelationship)
	}
	return configimporteditemrelationships
}
func flattenListConfigMoRef(p []models.ConfigMoRef, d *schema.ResourceData) []map[string]interface{} {
	var configmorefs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		configmoref := make(map[string]interface{})
		configmoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		configmoref["class_id"] = item.ClassId
		configmoref["moid"] = item.Moid
		configmoref["object_type"] = item.ObjectType
		configmorefs = append(configmorefs, configmoref)
	}
	return configmorefs
}
func flattenListConnectorpackConnectorPackUpdate(p []models.ConnectorpackConnectorPackUpdate, d *schema.ResourceData) []map[string]interface{} {
	var connectorpackconnectorpackupdates []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		connectorpackconnectorpackupdate := make(map[string]interface{})
		connectorpackconnectorpackupdate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		connectorpackconnectorpackupdate["class_id"] = item.ClassId
		connectorpackconnectorpackupdate["current_version"] = item.CurrentVersion
		connectorpackconnectorpackupdate["name"] = item.Name
		connectorpackconnectorpackupdate["new_version"] = item.NewVersion
		connectorpackconnectorpackupdate["object_type"] = item.ObjectType
		connectorpackconnectorpackupdates = append(connectorpackconnectorpackupdates, connectorpackconnectorpackupdate)
	}
	return connectorpackconnectorpackupdates
}
func flattenListContentComplexType(p []models.ContentComplexType, d *schema.ResourceData) []map[string]interface{} {
	var contentcomplextypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		contentcomplextype := make(map[string]interface{})
		contentcomplextype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		contentcomplextype["class_id"] = item.ClassId
		contentcomplextype["name"] = item.Name
		contentcomplextype["object_type"] = item.ObjectType
		contentcomplextype["parameters"] = (func(p []models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
			var contentbaseparameters []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				contentbaseparameter := make(map[string]interface{})
				contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
				contentbaseparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				contentbaseparameter["class_id"] = item.ClassId
				contentbaseparameter["complex_type"] = item.ComplexType
				contentbaseparameter["item_type"] = item.ItemType
				contentbaseparameter["name"] = item.Name
				contentbaseparameter["object_type"] = item.ObjectType
				contentbaseparameter["path"] = item.Path
				contentbaseparameter["type"] = item.Type
				contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
			}
			return contentbaseparameters
		})(item.GetParameters(), d)
		contentcomplextypes = append(contentcomplextypes, contentcomplextype)
	}
	return contentcomplextypes
}
func flattenListContentParameter(p []models.ContentParameter, d *schema.ResourceData) []map[string]interface{} {
	var contentparameters []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		contentparameter := make(map[string]interface{})
		contentparameter["accept_single_value"] = item.AcceptSingleValue
		contentparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		contentparameter["class_id"] = item.ClassId
		contentparameter["complex_type"] = item.ComplexType
		contentparameter["item_type"] = item.ItemType
		contentparameter["name"] = item.Name
		contentparameter["object_type"] = item.ObjectType
		contentparameter["path"] = item.Path
		contentparameter["type"] = item.Type
		contentparameters = append(contentparameters, contentparameter)
	}
	return contentparameters
}
func flattenListEquipmentFanRelationship(p []models.EquipmentFanRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentfanrelationship := flattenMoMoRef(item)
		equipmentfanrelationships = append(equipmentfanrelationships, equipmentfanrelationship)
	}
	return equipmentfanrelationships
}
func flattenListEquipmentFanModuleRelationship(p []models.EquipmentFanModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentfanmodulerelationship := flattenMoMoRef(item)
		equipmentfanmodulerelationships = append(equipmentfanmodulerelationships, equipmentfanmodulerelationship)
	}
	return equipmentfanmodulerelationships
}
func flattenListEquipmentIoCardRelationship(p []models.EquipmentIoCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentiocardrelationship := flattenMoMoRef(item)
		equipmentiocardrelationships = append(equipmentiocardrelationships, equipmentiocardrelationship)
	}
	return equipmentiocardrelationships
}
func flattenListEquipmentIoCardIdentity(p []models.EquipmentIoCardIdentity, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardidentitys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		equipmentiocardidentity := make(map[string]interface{})
		equipmentiocardidentity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		equipmentiocardidentity["class_id"] = item.ClassId
		equipmentiocardidentity["io_card_moid"] = item.IoCardMoid
		equipmentiocardidentity["module_id"] = item.ModuleId
		equipmentiocardidentity["network_element_moid"] = item.NetworkElementMoid
		equipmentiocardidentity["object_type"] = item.ObjectType
		equipmentiocardidentity["switch_id"] = item.SwitchId
		equipmentiocardidentitys = append(equipmentiocardidentitys, equipmentiocardidentity)
	}
	return equipmentiocardidentitys
}
func flattenListEquipmentIoExpanderRelationship(p []models.EquipmentIoExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentioexpanderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentioexpanderrelationship := flattenMoMoRef(item)
		equipmentioexpanderrelationships = append(equipmentioexpanderrelationships, equipmentioexpanderrelationship)
	}
	return equipmentioexpanderrelationships
}
func flattenListEquipmentPsuRelationship(p []models.EquipmentPsuRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentpsurelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentpsurelationship := flattenMoMoRef(item)
		equipmentpsurelationships = append(equipmentpsurelationships, equipmentpsurelationship)
	}
	return equipmentpsurelationships
}
func flattenListEquipmentRackEnclosureSlotRelationship(p []models.EquipmentRackEnclosureSlotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentrackenclosureslotrelationship := flattenMoMoRef(item)
		equipmentrackenclosureslotrelationships = append(equipmentrackenclosureslotrelationships, equipmentrackenclosureslotrelationship)
	}
	return equipmentrackenclosureslotrelationships
}
func flattenListEquipmentSwitchCardRelationship(p []models.EquipmentSwitchCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentswitchcardrelationship := flattenMoMoRef(item)
		equipmentswitchcardrelationships = append(equipmentswitchcardrelationships, equipmentswitchcardrelationship)
	}
	return equipmentswitchcardrelationships
}
func flattenListEquipmentSystemIoControllerRelationship(p []models.EquipmentSystemIoControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmentsystemiocontrollerrelationship := flattenMoMoRef(item)
		equipmentsystemiocontrollerrelationships = append(equipmentsystemiocontrollerrelationships, equipmentsystemiocontrollerrelationship)
	}
	return equipmentsystemiocontrollerrelationships
}
func flattenListEquipmentTpmRelationship(p []models.EquipmentTpmRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmenttpmrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		equipmenttpmrelationship := flattenMoMoRef(item)
		equipmenttpmrelationships = append(equipmenttpmrelationships, equipmenttpmrelationship)
	}
	return equipmenttpmrelationships
}
func flattenListEtherHostPortRelationship(p []models.EtherHostPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherhostportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		etherhostportrelationship := flattenMoMoRef(item)
		etherhostportrelationships = append(etherhostportrelationships, etherhostportrelationship)
	}
	return etherhostportrelationships
}
func flattenListEtherNetworkPortRelationship(p []models.EtherNetworkPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ethernetworkportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ethernetworkportrelationship := flattenMoMoRef(item)
		ethernetworkportrelationships = append(ethernetworkportrelationships, ethernetworkportrelationship)
	}
	return ethernetworkportrelationships
}
func flattenListEtherPhysicalPortRelationship(p []models.EtherPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		etherphysicalportrelationship := flattenMoMoRef(item)
		etherphysicalportrelationships = append(etherphysicalportrelationships, etherphysicalportrelationship)
	}
	return etherphysicalportrelationships
}
func flattenListEtherPortChannelRelationship(p []models.EtherPortChannelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherportchannelrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		etherportchannelrelationship := flattenMoMoRef(item)
		etherportchannelrelationships = append(etherportchannelrelationships, etherportchannelrelationship)
	}
	return etherportchannelrelationships
}
func flattenListFabricConfigChangeDetailRelationship(p []models.FabricConfigChangeDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricconfigchangedetailrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fabricconfigchangedetailrelationship := flattenMoMoRef(item)
		fabricconfigchangedetailrelationships = append(fabricconfigchangedetailrelationships, fabricconfigchangedetailrelationship)
	}
	return fabricconfigchangedetailrelationships
}
func flattenListFabricConfigResultEntryRelationship(p []models.FabricConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fabricconfigresultentryrelationship := flattenMoMoRef(item)
		fabricconfigresultentryrelationships = append(fabricconfigresultentryrelationships, fabricconfigresultentryrelationship)
	}
	return fabricconfigresultentryrelationships
}
func flattenListFabricPortIdentifier(p []models.FabricPortIdentifier, d *schema.ResourceData) []map[string]interface{} {
	var fabricportidentifiers []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		fabricportidentifier := make(map[string]interface{})
		fabricportidentifier["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		fabricportidentifier["aggregate_port_id"] = item.AggregatePortId
		fabricportidentifier["class_id"] = item.ClassId
		fabricportidentifier["object_type"] = item.ObjectType
		fabricportidentifier["port_id"] = item.PortId
		fabricportidentifier["slot_id"] = item.SlotId
		fabricportidentifiers = append(fabricportidentifiers, fabricportidentifier)
	}
	return fabricportidentifiers
}
func flattenListFabricQosClass(p []models.FabricQosClass, d *schema.ResourceData) []map[string]interface{} {
	var fabricqosclasss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		fabricqosclass := make(map[string]interface{})
		fabricqosclass["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		fabricqosclass["admin_state"] = item.AdminState
		fabricqosclass["bandwidth_percent"] = item.BandwidthPercent
		fabricqosclass["class_id"] = item.ClassId
		fabricqosclass["cos"] = item.Cos
		fabricqosclass["mtu"] = item.Mtu
		fabricqosclass["multicast_optimize"] = item.MulticastOptimize
		fabricqosclass["name"] = item.Name
		fabricqosclass["object_type"] = item.ObjectType
		fabricqosclass["packet_drop"] = item.PacketDrop
		fabricqosclass["weight"] = item.Weight
		fabricqosclasss = append(fabricqosclasss, fabricqosclass)
	}
	return fabricqosclasss
}
func flattenListFabricSwitchProfileRelationship(p []models.FabricSwitchProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricswitchprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fabricswitchprofilerelationship := flattenMoMoRef(item)
		fabricswitchprofilerelationships = append(fabricswitchprofilerelationships, fabricswitchprofilerelationship)
	}
	return fabricswitchprofilerelationships
}
func flattenListFcPhysicalPortRelationship(p []models.FcPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcphysicalportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fcphysicalportrelationship := flattenMoMoRef(item)
		fcphysicalportrelationships = append(fcphysicalportrelationships, fcphysicalportrelationship)
	}
	return fcphysicalportrelationships
}
func flattenListFcPortChannelRelationship(p []models.FcPortChannelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcportchannelrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fcportchannelrelationship := flattenMoMoRef(item)
		fcportchannelrelationships = append(fcportchannelrelationships, fcportchannelrelationship)
	}
	return fcportchannelrelationships
}
func flattenListFcpoolBlock(p []models.FcpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		fcpoolblock := make(map[string]interface{})
		fcpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		fcpoolblock["class_id"] = item.ClassId
		fcpoolblock["from"] = item.From
		fcpoolblock["object_type"] = item.ObjectType
		fcpoolblock["size"] = item.Size
		fcpoolblock["to"] = item.To
		fcpoolblocks = append(fcpoolblocks, fcpoolblock)
	}
	return fcpoolblocks
}
func flattenListFcpoolFcBlockRelationship(p []models.FcpoolFcBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolfcblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		fcpoolfcblockrelationship := flattenMoMoRef(item)
		fcpoolfcblockrelationships = append(fcpoolfcblockrelationships, fcpoolfcblockrelationship)
	}
	return fcpoolfcblockrelationships
}
func flattenListFirmwareBaseImpact(p []models.FirmwareBaseImpact, d *schema.ResourceData) []map[string]interface{} {
	var firmwarebaseimpacts []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		firmwarebaseimpact := make(map[string]interface{})
		firmwarebaseimpact["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarebaseimpact["class_id"] = item.ClassId
		firmwarebaseimpact["computation_error"] = item.ComputationError
		firmwarebaseimpact["computation_status_detail"] = item.ComputationStatusDetail
		firmwarebaseimpact["domain_name"] = item.DomainName
		firmwarebaseimpact["end_point"] = item.EndPoint
		firmwarebaseimpact["firmware_version"] = item.FirmwareVersion
		firmwarebaseimpact["impact_type"] = item.ImpactType
		firmwarebaseimpact["model"] = item.Model
		firmwarebaseimpact["object_type"] = item.ObjectType
		firmwarebaseimpact["target_firmware_version"] = item.TargetFirmwareVersion
		firmwarebaseimpact["version_impact"] = item.VersionImpact
		firmwarebaseimpacts = append(firmwarebaseimpacts, firmwarebaseimpact)
	}
	return firmwarebaseimpacts
}
func flattenListFirmwareComponentMeta(p []models.FirmwareComponentMeta, d *schema.ResourceData) []map[string]interface{} {
	var firmwarecomponentmetas []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		firmwarecomponentmeta := make(map[string]interface{})
		firmwarecomponentmeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarecomponentmeta["class_id"] = item.ClassId
		firmwarecomponentmeta["component_label"] = item.ComponentLabel
		firmwarecomponentmeta["component_type"] = item.ComponentType
		firmwarecomponentmeta["description"] = item.Description
		firmwarecomponentmeta["disruption"] = item.Disruption
		firmwarecomponentmeta["is_oob_supported"] = item.IsOobSupported
		firmwarecomponentmeta["model"] = item.Model
		firmwarecomponentmeta["object_type"] = item.ObjectType
		firmwarecomponentmeta["oob_manageability"] = item.OobManageability
		firmwarecomponentmeta["packed_version"] = item.PackedVersion
		firmwarecomponentmeta["redfish_url"] = item.RedfishUrl
		firmwarecomponentmeta["vendor"] = item.Vendor
		firmwarecomponentmetas = append(firmwarecomponentmetas, firmwarecomponentmeta)
	}
	return firmwarecomponentmetas
}
func flattenListFirmwareDistributableMetaRelationship(p []models.FirmwareDistributableMetaRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredistributablemetarelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		firmwaredistributablemetarelationship := flattenMoMoRef(item)
		firmwaredistributablemetarelationships = append(firmwaredistributablemetarelationships, firmwaredistributablemetarelationship)
	}
	return firmwaredistributablemetarelationships
}
func flattenListFirmwareFirmwareInventory(p []models.FirmwareFirmwareInventory, d *schema.ResourceData) []map[string]interface{} {
	var firmwarefirmwareinventorys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		firmwarefirmwareinventory := make(map[string]interface{})
		firmwarefirmwareinventory["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarefirmwareinventory["category"] = item.Category
		firmwarefirmwareinventory["class_id"] = item.ClassId
		firmwarefirmwareinventory["label"] = item.Label
		firmwarefirmwareinventory["model"] = item.Model
		firmwarefirmwareinventory["object_type"] = item.ObjectType
		firmwarefirmwareinventory["update_uri"] = item.UpdateUri
		firmwarefirmwareinventory["vendor"] = item.Vendor
		firmwarefirmwareinventory["nr_version"] = item.Version
		firmwarefirmwareinventorys = append(firmwarefirmwareinventorys, firmwarefirmwareinventory)
	}
	return firmwarefirmwareinventorys
}
func flattenListFirmwareRunningFirmwareRelationship(p []models.FirmwareRunningFirmwareRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		firmwarerunningfirmwarerelationship := flattenMoMoRef(item)
		firmwarerunningfirmwarerelationships = append(firmwarerunningfirmwarerelationships, firmwarerunningfirmwarerelationship)
	}
	return firmwarerunningfirmwarerelationships
}
func flattenListForecastDefinitionRelationship(p []models.ForecastDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		forecastdefinitionrelationship := flattenMoMoRef(item)
		forecastdefinitionrelationships = append(forecastdefinitionrelationships, forecastdefinitionrelationship)
	}
	return forecastdefinitionrelationships
}
func flattenListGraphicsCardRelationship(p []models.GraphicsCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		graphicscardrelationship := flattenMoMoRef(item)
		graphicscardrelationships = append(graphicscardrelationships, graphicscardrelationship)
	}
	return graphicscardrelationships
}
func flattenListGraphicsControllerRelationship(p []models.GraphicsControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		graphicscontrollerrelationship := flattenMoMoRef(item)
		graphicscontrollerrelationships = append(graphicscontrollerrelationships, graphicscontrollerrelationship)
	}
	return graphicscontrollerrelationships
}
func flattenListHclConstraint(p []models.HclConstraint, d *schema.ResourceData) []map[string]interface{} {
	var hclconstraints []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hclconstraint := make(map[string]interface{})
		hclconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hclconstraint["class_id"] = item.ClassId
		hclconstraint["constraint_name"] = item.ConstraintName
		hclconstraint["constraint_value"] = item.ConstraintValue
		hclconstraint["object_type"] = item.ObjectType
		hclconstraints = append(hclconstraints, hclconstraint)
	}
	return hclconstraints
}
func flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(p []models.HclHyperflexSoftwareCompatibilityInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hclhyperflexsoftwarecompatibilityinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hclhyperflexsoftwarecompatibilityinforelationship := flattenMoMoRef(item)
		hclhyperflexsoftwarecompatibilityinforelationships = append(hclhyperflexsoftwarecompatibilityinforelationships, hclhyperflexsoftwarecompatibilityinforelationship)
	}
	return hclhyperflexsoftwarecompatibilityinforelationships
}
func flattenListHclOperatingSystemRelationship(p []models.HclOperatingSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hcloperatingsystemrelationship := flattenMoMoRef(item)
		hcloperatingsystemrelationships = append(hcloperatingsystemrelationships, hcloperatingsystemrelationship)
	}
	return hcloperatingsystemrelationships
}
func flattenListHyperflexAlarmRelationship(p []models.HyperflexAlarmRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexalarmrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexalarmrelationship := flattenMoMoRef(item)
		hyperflexalarmrelationships = append(hyperflexalarmrelationships, hyperflexalarmrelationship)
	}
	return hyperflexalarmrelationships
}
func flattenListHyperflexCapabilityInfoRelationship(p []models.HyperflexCapabilityInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexcapabilityinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexcapabilityinforelationship := flattenMoMoRef(item)
		hyperflexcapabilityinforelationships = append(hyperflexcapabilityinforelationships, hyperflexcapabilityinforelationship)
	}
	return hyperflexcapabilityinforelationships
}
func flattenListHyperflexClusterProfileRelationship(p []models.HyperflexClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexclusterprofilerelationship := flattenMoMoRef(item)
		hyperflexclusterprofilerelationships = append(hyperflexclusterprofilerelationships, hyperflexclusterprofilerelationship)
	}
	return hyperflexclusterprofilerelationships
}
func flattenListHyperflexConfigResultEntryRelationship(p []models.HyperflexConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexconfigresultentryrelationship := flattenMoMoRef(item)
		hyperflexconfigresultentryrelationships = append(hyperflexconfigresultentryrelationships, hyperflexconfigresultentryrelationship)
	}
	return hyperflexconfigresultentryrelationships
}
func flattenListHyperflexFeatureLimitEntry(p []models.HyperflexFeatureLimitEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitentrys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexfeaturelimitentry := make(map[string]interface{})
		hyperflexfeaturelimitentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexfeaturelimitentry["class_id"] = item.ClassId
		hyperflexfeaturelimitentry["constraint"] = (func(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			var ret models.HyperflexAppSettingConstraint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexappsettingconstraint["class_id"] = item.ClassId
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.GetConstraint(), d)
		hyperflexfeaturelimitentry["name"] = item.Name
		hyperflexfeaturelimitentry["object_type"] = item.ObjectType
		hyperflexfeaturelimitentry["value"] = item.Value
		hyperflexfeaturelimitentrys = append(hyperflexfeaturelimitentrys, hyperflexfeaturelimitentry)
	}
	return hyperflexfeaturelimitentrys
}
func flattenListHyperflexHxZoneResiliencyInfoDt(p []models.HyperflexHxZoneResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxzoneresiliencyinfodts []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexhxzoneresiliencyinfodt := make(map[string]interface{})
		hyperflexhxzoneresiliencyinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexhxzoneresiliencyinfodt["class_id"] = item.ClassId
		hyperflexhxzoneresiliencyinfodt["name"] = item.Name
		hyperflexhxzoneresiliencyinfodt["object_type"] = item.ObjectType
		hyperflexhxzoneresiliencyinfodt["resiliency_info"] = (func(p models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexhxresiliencyinfodts []map[string]interface{}
			var ret models.HyperflexHxResiliencyInfoDt
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexhxresiliencyinfodt := make(map[string]interface{})
			hyperflexhxresiliencyinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexhxresiliencyinfodt["class_id"] = item.ClassId
			hyperflexhxresiliencyinfodt["data_replication_factor"] = item.DataReplicationFactor
			hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.HddFailuresTolerable
			hyperflexhxresiliencyinfodt["messages"] = item.Messages
			hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.NodeFailuresTolerable
			hyperflexhxresiliencyinfodt["object_type"] = item.ObjectType
			hyperflexhxresiliencyinfodt["policy_compliance"] = item.PolicyCompliance
			hyperflexhxresiliencyinfodt["resiliency_state"] = item.ResiliencyState
			hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.SsdFailuresTolerable

			hyperflexhxresiliencyinfodts = append(hyperflexhxresiliencyinfodts, hyperflexhxresiliencyinfodt)
			return hyperflexhxresiliencyinfodts
		})(item.GetResiliencyInfo(), d)
		hyperflexhxzoneresiliencyinfodts = append(hyperflexhxzoneresiliencyinfodts, hyperflexhxzoneresiliencyinfodt)
	}
	return hyperflexhxzoneresiliencyinfodts
}
func flattenListHyperflexHxdpVersionRelationship(p []models.HyperflexHxdpVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxdpversionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexhxdpversionrelationship := flattenMoMoRef(item)
		hyperflexhxdpversionrelationships = append(hyperflexhxdpversionrelationships, hyperflexhxdpversionrelationship)
	}
	return hyperflexhxdpversionrelationships
}
func flattenListHyperflexNamedVlan(p []models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvlans []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexnamedvlan := make(map[string]interface{})
		hyperflexnamedvlan["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexnamedvlan["class_id"] = item.ClassId
		hyperflexnamedvlan["name"] = item.Name
		hyperflexnamedvlan["object_type"] = item.ObjectType
		hyperflexnamedvlan["vlan_id"] = item.VlanId
		hyperflexnamedvlans = append(hyperflexnamedvlans, hyperflexnamedvlan)
	}
	return hyperflexnamedvlans
}
func flattenListHyperflexNodeRelationship(p []models.HyperflexNodeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnoderelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexnoderelationship := flattenMoMoRef(item)
		hyperflexnoderelationships = append(hyperflexnoderelationships, hyperflexnoderelationship)
	}
	return hyperflexnoderelationships
}
func flattenListHyperflexNodeProfileRelationship(p []models.HyperflexNodeProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		hyperflexnodeprofilerelationship := flattenMoMoRef(item)
		hyperflexnodeprofilerelationships = append(hyperflexnodeprofilerelationships, hyperflexnodeprofilerelationship)
	}
	return hyperflexnodeprofilerelationships
}
func flattenListHyperflexServerFirmwareVersionEntry(p []models.HyperflexServerFirmwareVersionEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionentrys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexserverfirmwareversionentry := make(map[string]interface{})
		hyperflexserverfirmwareversionentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexserverfirmwareversionentry["class_id"] = item.ClassId
		hyperflexserverfirmwareversionentry["constraint"] = (func(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			var ret models.HyperflexAppSettingConstraint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexappsettingconstraint["class_id"] = item.ClassId
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.GetConstraint(), d)
		hyperflexserverfirmwareversionentry["label"] = item.Label
		hyperflexserverfirmwareversionentry["name"] = item.Name
		hyperflexserverfirmwareversionentry["object_type"] = item.ObjectType
		hyperflexserverfirmwareversionentry["value"] = item.Value
		hyperflexserverfirmwareversionentrys = append(hyperflexserverfirmwareversionentrys, hyperflexserverfirmwareversionentry)
	}
	return hyperflexserverfirmwareversionentrys
}
func flattenListHyperflexServerModelEntry(p []models.HyperflexServerModelEntry, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelentrys []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		hyperflexservermodelentry := make(map[string]interface{})
		hyperflexservermodelentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexservermodelentry["class_id"] = item.ClassId
		hyperflexservermodelentry["constraint"] = (func(p models.HyperflexAppSettingConstraint, d *schema.ResourceData) []map[string]interface{} {
			var hyperflexappsettingconstraints []map[string]interface{}
			var ret models.HyperflexAppSettingConstraint
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			hyperflexappsettingconstraint := make(map[string]interface{})
			hyperflexappsettingconstraint["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexappsettingconstraint["class_id"] = item.ClassId
			hyperflexappsettingconstraint["hxdp_version"] = item.HxdpVersion
			hyperflexappsettingconstraint["hypervisor_type"] = item.HypervisorType
			hyperflexappsettingconstraint["mgmt_platform"] = item.MgmtPlatform
			hyperflexappsettingconstraint["object_type"] = item.ObjectType
			hyperflexappsettingconstraint["server_model"] = item.ServerModel

			hyperflexappsettingconstraints = append(hyperflexappsettingconstraints, hyperflexappsettingconstraint)
			return hyperflexappsettingconstraints
		})(item.GetConstraint(), d)
		hyperflexservermodelentry["name"] = item.Name
		hyperflexservermodelentry["object_type"] = item.ObjectType
		hyperflexservermodelentry["value"] = item.Value
		hyperflexservermodelentrys = append(hyperflexservermodelentrys, hyperflexservermodelentry)
	}
	return hyperflexservermodelentrys
}
func flattenListIaasConnectorPackRelationship(p []models.IaasConnectorPackRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasconnectorpackrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iaasconnectorpackrelationship := flattenMoMoRef(item)
		iaasconnectorpackrelationships = append(iaasconnectorpackrelationships, iaasconnectorpackrelationship)
	}
	return iaasconnectorpackrelationships
}
func flattenListIaasDeviceStatusRelationship(p []models.IaasDeviceStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasdevicestatusrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iaasdevicestatusrelationship := flattenMoMoRef(item)
		iaasdevicestatusrelationships = append(iaasdevicestatusrelationships, iaasdevicestatusrelationship)
	}
	return iaasdevicestatusrelationships
}
func flattenListIaasLicenseKeysInfo(p []models.IaasLicenseKeysInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicensekeysinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iaaslicensekeysinfo := make(map[string]interface{})
		iaaslicensekeysinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iaaslicensekeysinfo["class_id"] = item.ClassId
		iaaslicensekeysinfo["nr_count"] = item.Count
		iaaslicensekeysinfo["expiration_date"] = item.ExpirationDate
		iaaslicensekeysinfo["license_id"] = item.LicenseId
		iaaslicensekeysinfo["object_type"] = item.ObjectType
		iaaslicensekeysinfo["pid"] = item.Pid
		iaaslicensekeysinfos = append(iaaslicensekeysinfos, iaaslicensekeysinfo)
	}
	return iaaslicensekeysinfos
}
func flattenListIaasLicenseUtilizationInfo(p []models.IaasLicenseUtilizationInfo, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseutilizationinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iaaslicenseutilizationinfo := make(map[string]interface{})
		iaaslicenseutilizationinfo["actual_used"] = item.ActualUsed
		iaaslicenseutilizationinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iaaslicenseutilizationinfo["class_id"] = item.ClassId
		iaaslicenseutilizationinfo["label"] = item.Label
		iaaslicenseutilizationinfo["licensed_limit"] = item.LicensedLimit
		iaaslicenseutilizationinfo["object_type"] = item.ObjectType
		iaaslicenseutilizationinfo["sku"] = item.Sku
		iaaslicenseutilizationinfos = append(iaaslicenseutilizationinfos, iaaslicenseutilizationinfo)
	}
	return iaaslicenseutilizationinfos
}
func flattenListIaasMostRunTasksRelationship(p []models.IaasMostRunTasksRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasmostruntasksrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iaasmostruntasksrelationship := flattenMoMoRef(item)
		iaasmostruntasksrelationships = append(iaasmostruntasksrelationships, iaasmostruntasksrelationship)
	}
	return iaasmostruntasksrelationships
}
func flattenListIamAccountPermissions(p []models.IamAccountPermissions, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountpermissionss []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iamaccountpermissions := make(map[string]interface{})
		iamaccountpermissions["account_identifier"] = item.AccountIdentifier
		iamaccountpermissions["account_name"] = item.AccountName
		iamaccountpermissions["account_status"] = item.AccountStatus
		iamaccountpermissions["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iamaccountpermissions["class_id"] = item.ClassId
		iamaccountpermissions["object_type"] = item.ObjectType
		iamaccountpermissions["permissions"] = (func(p []models.IamPermissionReference, d *schema.ResourceData) []map[string]interface{} {
			var iampermissionreferences []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				iampermissionreference := make(map[string]interface{})
				iampermissionreference["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				iampermissionreference["class_id"] = item.ClassId
				iampermissionreference["object_type"] = item.ObjectType
				iampermissionreference["permission_identifier"] = item.PermissionIdentifier
				iampermissionreference["permission_name"] = item.PermissionName
				iampermissionreferences = append(iampermissionreferences, iampermissionreference)
			}
			return iampermissionreferences
		})(item.GetPermissions(), d)
		iamaccountpermissionss = append(iamaccountpermissionss, iamaccountpermissions)
	}
	return iamaccountpermissionss
}
func flattenListIamApiKeyRelationship(p []models.IamApiKeyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamapikeyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamapikeyrelationship := flattenMoMoRef(item)
		iamapikeyrelationships = append(iamapikeyrelationships, iamapikeyrelationship)
	}
	return iamapikeyrelationships
}
func flattenListIamAppRegistrationRelationship(p []models.IamAppRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamappregistrationrelationship := flattenMoMoRef(item)
		iamappregistrationrelationships = append(iamappregistrationrelationships, iamappregistrationrelationship)
	}
	return iamappregistrationrelationships
}
func flattenListIamDomainGroupRelationship(p []models.IamDomainGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamdomaingrouprelationship := flattenMoMoRef(item)
		iamdomaingrouprelationships = append(iamdomaingrouprelationships, iamdomaingrouprelationship)
	}
	return iamdomaingrouprelationships
}
func flattenListIamEndPointPrivilegeRelationship(p []models.IamEndPointPrivilegeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointprivilegerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamendpointprivilegerelationship := flattenMoMoRef(item)
		iamendpointprivilegerelationships = append(iamendpointprivilegerelationships, iamendpointprivilegerelationship)
	}
	return iamendpointprivilegerelationships
}
func flattenListIamEndPointRoleRelationship(p []models.IamEndPointRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointrolerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamendpointrolerelationship := flattenMoMoRef(item)
		iamendpointrolerelationships = append(iamendpointrolerelationships, iamendpointrolerelationship)
	}
	return iamendpointrolerelationships
}
func flattenListIamEndPointUserRoleRelationship(p []models.IamEndPointUserRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrolerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamendpointuserrolerelationship := flattenMoMoRef(item)
		iamendpointuserrolerelationships = append(iamendpointuserrolerelationships, iamendpointuserrolerelationship)
	}
	return iamendpointuserrolerelationships
}
func flattenListIamFeatureDefinition(p []models.IamFeatureDefinition, d *schema.ResourceData) []map[string]interface{} {
	var iamfeaturedefinitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iamfeaturedefinition := make(map[string]interface{})
		iamfeaturedefinition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iamfeaturedefinition["class_id"] = item.ClassId
		iamfeaturedefinition["feature"] = item.Feature
		iamfeaturedefinition["object_type"] = item.ObjectType
		iamfeaturedefinitions = append(iamfeaturedefinitions, iamfeaturedefinition)
	}
	return iamfeaturedefinitions
}
func flattenListIamGroupPermissionToRoles(p []models.IamGroupPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iamgrouppermissiontoroless []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iamgrouppermissiontoroles := make(map[string]interface{})
		iamgrouppermissiontoroles["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iamgrouppermissiontoroles["class_id"] = item.ClassId
		iamgrouppermissiontoroles["group"] = (func(p models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			var ret models.CmrfCmRf
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			cmrfcmrf := make(map[string]interface{})
			cmrfcmrf["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			cmrfcmrf["class_id"] = item.ClassId
			cmrfcmrf["moid"] = item.Moid
			cmrfcmrf["object_type"] = item.ObjectType

			cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			return cmrfcmrfs
		})(item.GetGroup(), d)
		iamgrouppermissiontoroles["object_type"] = item.ObjectType
		iamgrouppermissiontoroles["orgs"] = (func(p []models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				cmrfcmrf := make(map[string]interface{})
				cmrfcmrf["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				cmrfcmrf["class_id"] = item.ClassId
				cmrfcmrf["moid"] = item.Moid
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			}
			return cmrfcmrfs
		})(item.GetOrgs(), d)
		iamgrouppermissiontoroless = append(iamgrouppermissiontoroless, iamgrouppermissiontoroles)
	}
	return iamgrouppermissiontoroless
}
func flattenListIamIdpRelationship(p []models.IamIdpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamidprelationship := flattenMoMoRef(item)
		iamidprelationships = append(iamidprelationships, iamidprelationship)
	}
	return iamidprelationships
}
func flattenListIamIdpReferenceRelationship(p []models.IamIdpReferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamidpreferencerelationship := flattenMoMoRef(item)
		iamidpreferencerelationships = append(iamidpreferencerelationships, iamidpreferencerelationship)
	}
	return iamidpreferencerelationships
}
func flattenListIamLdapGroupRelationship(p []models.IamLdapGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldapgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamldapgrouprelationship := flattenMoMoRef(item)
		iamldapgrouprelationships = append(iamldapgrouprelationships, iamldapgrouprelationship)
	}
	return iamldapgrouprelationships
}
func flattenListIamLdapProviderRelationship(p []models.IamLdapProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldapproviderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamldapproviderrelationship := flattenMoMoRef(item)
		iamldapproviderrelationships = append(iamldapproviderrelationships, iamldapproviderrelationship)
	}
	return iamldapproviderrelationships
}
func flattenListIamOAuthTokenRelationship(p []models.IamOAuthTokenRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamoauthtokenrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamoauthtokenrelationship := flattenMoMoRef(item)
		iamoauthtokenrelationships = append(iamoauthtokenrelationships, iamoauthtokenrelationship)
	}
	return iamoauthtokenrelationships
}
func flattenListIamPermissionRelationship(p []models.IamPermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iampermissionrelationship := flattenMoMoRef(item)
		iampermissionrelationships = append(iampermissionrelationships, iampermissionrelationship)
	}
	return iampermissionrelationships
}
func flattenListIamPermissionToRoles(p []models.IamPermissionToRoles, d *schema.ResourceData) []map[string]interface{} {
	var iampermissiontoroless []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		iampermissiontoroles := make(map[string]interface{})
		iampermissiontoroles["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		iampermissiontoroles["class_id"] = item.ClassId
		iampermissiontoroles["object_type"] = item.ObjectType
		iampermissiontoroles["permission"] = (func(p models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			var ret models.CmrfCmRf
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			cmrfcmrf := make(map[string]interface{})
			cmrfcmrf["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			cmrfcmrf["class_id"] = item.ClassId
			cmrfcmrf["moid"] = item.Moid
			cmrfcmrf["object_type"] = item.ObjectType

			cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			return cmrfcmrfs
		})(item.GetPermission(), d)
		iampermissiontoroles["roles"] = (func(p []models.CmrfCmRf, d *schema.ResourceData) []map[string]interface{} {
			var cmrfcmrfs []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				cmrfcmrf := make(map[string]interface{})
				cmrfcmrf["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				cmrfcmrf["class_id"] = item.ClassId
				cmrfcmrf["moid"] = item.Moid
				cmrfcmrf["object_type"] = item.ObjectType
				cmrfcmrfs = append(cmrfcmrfs, cmrfcmrf)
			}
			return cmrfcmrfs
		})(item.GetRoles(), d)
		iampermissiontoroless = append(iampermissiontoroless, iampermissiontoroles)
	}
	return iampermissiontoroless
}
func flattenListIamPrivilegeRelationship(p []models.IamPrivilegeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamprivilegerelationship := flattenMoMoRef(item)
		iamprivilegerelationships = append(iamprivilegerelationships, iamprivilegerelationship)
	}
	return iamprivilegerelationships
}
func flattenListIamPrivilegeSetRelationship(p []models.IamPrivilegeSetRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivilegesetrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamprivilegesetrelationship := flattenMoMoRef(item)
		iamprivilegesetrelationships = append(iamprivilegesetrelationships, iamprivilegesetrelationship)
	}
	return iamprivilegesetrelationships
}
func flattenListIamResourcePermissionRelationship(p []models.IamResourcePermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcepermissionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamresourcepermissionrelationship := flattenMoMoRef(item)
		iamresourcepermissionrelationships = append(iamresourcepermissionrelationships, iamresourcepermissionrelationship)
	}
	return iamresourcepermissionrelationships
}
func flattenListIamResourceRolesRelationship(p []models.IamResourceRolesRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcerolesrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamresourcerolesrelationship := flattenMoMoRef(item)
		iamresourcerolesrelationships = append(iamresourcerolesrelationships, iamresourcerolesrelationship)
	}
	return iamresourcerolesrelationships
}
func flattenListIamRoleRelationship(p []models.IamRoleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamrolerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamrolerelationship := flattenMoMoRef(item)
		iamrolerelationships = append(iamrolerelationships, iamrolerelationship)
	}
	return iamrolerelationships
}
func flattenListIamSessionRelationship(p []models.IamSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamsessionrelationship := flattenMoMoRef(item)
		iamsessionrelationships = append(iamsessionrelationships, iamsessionrelationship)
	}
	return iamsessionrelationships
}
func flattenListIamUserRelationship(p []models.IamUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamuserrelationship := flattenMoMoRef(item)
		iamuserrelationships = append(iamuserrelationships, iamuserrelationship)
	}
	return iamuserrelationships
}
func flattenListIamUserGroupRelationship(p []models.IamUserGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamusergrouprelationship := flattenMoMoRef(item)
		iamusergrouprelationships = append(iamusergrouprelationships, iamusergrouprelationship)
	}
	return iamusergrouprelationships
}
func flattenListIamUserPreferenceRelationship(p []models.IamUserPreferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserpreferencerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		iamuserpreferencerelationship := flattenMoMoRef(item)
		iamuserpreferencerelationships = append(iamuserpreferencerelationships, iamuserpreferencerelationship)
	}
	return iamuserpreferencerelationships
}
func flattenListInventoryGenericInventoryRelationship(p []models.InventoryGenericInventoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		inventorygenericinventoryrelationship := flattenMoMoRef(item)
		inventorygenericinventoryrelationships = append(inventorygenericinventoryrelationships, inventorygenericinventoryrelationship)
	}
	return inventorygenericinventoryrelationships
}
func flattenListInventoryGenericInventoryHolderRelationship(p []models.InventoryGenericInventoryHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		inventorygenericinventoryholderrelationship := flattenMoMoRef(item)
		inventorygenericinventoryholderrelationships = append(inventorygenericinventoryholderrelationships, inventorygenericinventoryholderrelationship)
	}
	return inventorygenericinventoryholderrelationships
}
func flattenListIppoolIpBlock(p []models.IppoolIpBlock, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		ippoolipblock := make(map[string]interface{})
		ippoolipblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		ippoolipblock["class_id"] = item.ClassId
		ippoolipblock["from"] = item.From
		ippoolipblock["object_type"] = item.ObjectType
		ippoolipblock["size"] = item.Size
		ippoolipblock["to"] = item.To
		ippoolipblocks = append(ippoolipblocks, ippoolipblock)
	}
	return ippoolipblocks
}
func flattenListIppoolShadowBlockRelationship(p []models.IppoolShadowBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ippoolshadowblockrelationship := flattenMoMoRef(item)
		ippoolshadowblockrelationships = append(ippoolshadowblockrelationships, ippoolshadowblockrelationship)
	}
	return ippoolshadowblockrelationships
}
func flattenListIppoolShadowPoolRelationship(p []models.IppoolShadowPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowpoolrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		ippoolshadowpoolrelationship := flattenMoMoRef(item)
		ippoolshadowpoolrelationships = append(ippoolshadowpoolrelationships, ippoolshadowpoolrelationship)
	}
	return ippoolshadowpoolrelationships
}
func flattenListLicenseLicenseInfoRelationship(p []models.LicenseLicenseInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenselicenseinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		licenselicenseinforelationship := flattenMoMoRef(item)
		licenselicenseinforelationships = append(licenselicenseinforelationships, licenselicenseinforelationship)
	}
	return licenselicenseinforelationships
}
func flattenListMacpoolBlock(p []models.MacpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var macpoolblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		macpoolblock := make(map[string]interface{})
		macpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		macpoolblock["class_id"] = item.ClassId
		macpoolblock["from"] = item.From
		macpoolblock["object_type"] = item.ObjectType
		macpoolblock["size"] = item.Size
		macpoolblock["to"] = item.To
		macpoolblocks = append(macpoolblocks, macpoolblock)
	}
	return macpoolblocks
}
func flattenListMacpoolIdBlockRelationship(p []models.MacpoolIdBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolidblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		macpoolidblockrelationship := flattenMoMoRef(item)
		macpoolidblockrelationships = append(macpoolidblockrelationships, macpoolidblockrelationship)
	}
	return macpoolidblockrelationships
}
func flattenListManagementInterfaceRelationship(p []models.ManagementInterfaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managementinterfacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		managementinterfacerelationship := flattenMoMoRef(item)
		managementinterfacerelationships = append(managementinterfacerelationships, managementinterfacerelationship)
	}
	return managementinterfacerelationships
}
func flattenListMemoryArrayRelationship(p []models.MemoryArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memoryarrayrelationship := flattenMoMoRef(item)
		memoryarrayrelationships = append(memoryarrayrelationships, memoryarrayrelationship)
	}
	return memoryarrayrelationships
}
func flattenListMemoryPersistentMemoryGoal(p []models.MemoryPersistentMemoryGoal, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorygoals []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		memorypersistentmemorygoal := make(map[string]interface{})
		memorypersistentmemorygoal["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		memorypersistentmemorygoal["class_id"] = item.ClassId
		memorypersistentmemorygoal["memory_mode_percentage"] = item.MemoryModePercentage
		memorypersistentmemorygoal["object_type"] = item.ObjectType
		memorypersistentmemorygoal["persistent_memory_type"] = item.PersistentMemoryType
		memorypersistentmemorygoal["socket_id"] = item.SocketId
		memorypersistentmemorygoals = append(memorypersistentmemorygoals, memorypersistentmemorygoal)
	}
	return memorypersistentmemorygoals
}
func flattenListMemoryPersistentMemoryLogicalNamespace(p []models.MemoryPersistentMemoryLogicalNamespace, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylogicalnamespaces []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		memorypersistentmemorylogicalnamespace := make(map[string]interface{})
		memorypersistentmemorylogicalnamespace["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		memorypersistentmemorylogicalnamespace["capacity"] = item.Capacity
		memorypersistentmemorylogicalnamespace["class_id"] = item.ClassId
		memorypersistentmemorylogicalnamespace["mode"] = item.Mode
		memorypersistentmemorylogicalnamespace["name"] = item.Name
		memorypersistentmemorylogicalnamespace["object_type"] = item.ObjectType
		memorypersistentmemorylogicalnamespace["socket_id"] = item.SocketId
		memorypersistentmemorylogicalnamespace["socket_memory_id"] = item.SocketMemoryId
		memorypersistentmemorylogicalnamespaces = append(memorypersistentmemorylogicalnamespaces, memorypersistentmemorylogicalnamespace)
	}
	return memorypersistentmemorylogicalnamespaces
}
func flattenListMemoryPersistentMemoryNamespaceRelationship(p []models.MemoryPersistentMemoryNamespaceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespacerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemorynamespacerelationship := flattenMoMoRef(item)
		memorypersistentmemorynamespacerelationships = append(memorypersistentmemorynamespacerelationships, memorypersistentmemorynamespacerelationship)
	}
	return memorypersistentmemorynamespacerelationships
}
func flattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(p []models.MemoryPersistentMemoryNamespaceConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorynamespaceconfigresultrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemorynamespaceconfigresultrelationship := flattenMoMoRef(item)
		memorypersistentmemorynamespaceconfigresultrelationships = append(memorypersistentmemorynamespaceconfigresultrelationships, memorypersistentmemorynamespaceconfigresultrelationship)
	}
	return memorypersistentmemorynamespaceconfigresultrelationships
}
func flattenListMemoryPersistentMemoryRegionRelationship(p []models.MemoryPersistentMemoryRegionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemoryregionrelationship := flattenMoMoRef(item)
		memorypersistentmemoryregionrelationships = append(memorypersistentmemoryregionrelationships, memorypersistentmemoryregionrelationship)
	}
	return memorypersistentmemoryregionrelationships
}
func flattenListMemoryPersistentMemoryUnitRelationship(p []models.MemoryPersistentMemoryUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memorypersistentmemoryunitrelationship := flattenMoMoRef(item)
		memorypersistentmemoryunitrelationships = append(memorypersistentmemoryunitrelationships, memorypersistentmemoryunitrelationship)
	}
	return memorypersistentmemoryunitrelationships
}
func flattenListMemoryUnitRelationship(p []models.MemoryUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		memoryunitrelationship := flattenMoMoRef(item)
		memoryunitrelationships = append(memoryunitrelationships, memoryunitrelationship)
	}
	return memoryunitrelationships
}
func flattenListMetaAccessPrivilege(p []models.MetaAccessPrivilege, d *schema.ResourceData) []map[string]interface{} {
	var metaaccessprivileges []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		metaaccessprivilege := make(map[string]interface{})
		metaaccessprivilege["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		metaaccessprivilege["class_id"] = item.ClassId
		metaaccessprivilege["method"] = item.Method
		metaaccessprivilege["object_type"] = item.ObjectType
		metaaccessprivilege["privilege"] = item.Privilege
		metaaccessprivileges = append(metaaccessprivileges, metaaccessprivilege)
	}
	return metaaccessprivileges
}
func flattenListMetaPropDefinition(p []models.MetaPropDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metapropdefinitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		metapropdefinition := make(map[string]interface{})
		metapropdefinition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		metapropdefinition["api_access"] = item.ApiAccess
		metapropdefinition["class_id"] = item.ClassId
		metapropdefinition["name"] = item.Name
		metapropdefinition["object_type"] = item.ObjectType
		metapropdefinition["op_security"] = item.OpSecurity
		metapropdefinition["search_weight"] = item.SearchWeight
		metapropdefinitions = append(metapropdefinitions, metapropdefinition)
	}
	return metapropdefinitions
}
func flattenListMetaRelationshipDefinition(p []models.MetaRelationshipDefinition, d *schema.ResourceData) []map[string]interface{} {
	var metarelationshipdefinitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		metarelationshipdefinition := make(map[string]interface{})
		metarelationshipdefinition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		metarelationshipdefinition["api_access"] = item.ApiAccess
		metarelationshipdefinition["class_id"] = item.ClassId
		metarelationshipdefinition["collection"] = item.Collection
		metarelationshipdefinition["export"] = item.Export
		metarelationshipdefinition["export_with_peer"] = item.ExportWithPeer
		metarelationshipdefinition["name"] = item.Name
		metarelationshipdefinition["object_type"] = item.ObjectType
		metarelationshipdefinition["type"] = item.Type
		metarelationshipdefinitions = append(metarelationshipdefinitions, metarelationshipdefinition)
	}
	return metarelationshipdefinitions
}
func flattenListMoTag(p []models.MoTag, d *schema.ResourceData) []map[string]interface{} {
	var motags []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		motag := make(map[string]interface{})
		motag["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		motag["key"] = item.Key
		motag["value"] = item.Value
		motags = append(motags, motag)
	}
	return motags
}
func flattenListNetworkElementRelationship(p []models.NetworkElementRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		networkelementrelationship := flattenMoMoRef(item)
		networkelementrelationships = append(networkelementrelationships, networkelementrelationship)
	}
	return networkelementrelationships
}
func flattenListNiaapiDetail(p []models.NiaapiDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapidetails []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		niaapidetail := make(map[string]interface{})
		niaapidetail["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapidetail["chksum"] = item.Chksum
		niaapidetail["class_id"] = item.ClassId
		niaapidetail["filename"] = item.Filename
		niaapidetail["name"] = item.Name
		niaapidetail["object_type"] = item.ObjectType
		niaapidetails = append(niaapidetails, niaapidetail)
	}
	return niaapidetails
}
func flattenListNiaapiRevisionInfo(p []models.NiaapiRevisionInfo, d *schema.ResourceData) []map[string]interface{} {
	var niaapirevisioninfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		niaapirevisioninfo := make(map[string]interface{})
		niaapirevisioninfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapirevisioninfo["class_id"] = item.ClassId
		niaapirevisioninfo["date_published"] = item.DatePublished
		niaapirevisioninfo["object_type"] = item.ObjectType
		niaapirevisioninfo["revision_comment"] = item.RevisionComment
		niaapirevisioninfo["revision_no"] = item.RevisionNo
		niaapirevisioninfos = append(niaapirevisioninfos, niaapirevisioninfo)
	}
	return niaapirevisioninfos
}
func flattenListOnpremImagePackage(p []models.OnpremImagePackage, d *schema.ResourceData) []map[string]interface{} {
	var onpremimagepackages []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		onpremimagepackage := make(map[string]interface{})
		onpremimagepackage["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		onpremimagepackage["class_id"] = item.ClassId
		onpremimagepackage["file_path"] = item.FilePath
		onpremimagepackage["file_sha"] = item.FileSha
		onpremimagepackage["file_size"] = item.FileSize
		onpremimagepackage["file_time"] = item.FileTime
		onpremimagepackage["filename"] = item.Filename
		onpremimagepackage["name"] = item.Name
		onpremimagepackage["object_type"] = item.ObjectType
		onpremimagepackage["package_type"] = item.PackageType
		onpremimagepackage["nr_version"] = item.Version
		onpremimagepackages = append(onpremimagepackages, onpremimagepackage)
	}
	return onpremimagepackages
}
func flattenListOnpremUpgradeNote(p []models.OnpremUpgradeNote, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradenotes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		onpremupgradenote := make(map[string]interface{})
		onpremupgradenote["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		onpremupgradenote["class_id"] = item.ClassId
		onpremupgradenote["message"] = item.Message
		onpremupgradenote["object_type"] = item.ObjectType
		onpremupgradenotes = append(onpremupgradenotes, onpremupgradenote)
	}
	return onpremupgradenotes
}
func flattenListOnpremUpgradePhase(p []models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradephases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		onpremupgradephase := make(map[string]interface{})
		onpremupgradephase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		onpremupgradephase["class_id"] = item.ClassId
		onpremupgradephase["elapsed_time"] = item.ElapsedTime
		onpremupgradephase["end_time"] = item.EndTime
		onpremupgradephase["failed"] = item.Failed
		onpremupgradephase["message"] = item.Message
		onpremupgradephase["name"] = item.Name
		onpremupgradephase["object_type"] = item.ObjectType
		onpremupgradephase["start_time"] = item.StartTime
		onpremupgradephases = append(onpremupgradephases, onpremupgradephase)
	}
	return onpremupgradephases
}
func flattenListOrganizationOrganizationRelationship(p []models.OrganizationOrganizationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		organizationorganizationrelationship := flattenMoMoRef(item)
		organizationorganizationrelationships = append(organizationorganizationrelationships, organizationorganizationrelationship)
	}
	return organizationorganizationrelationships
}
func flattenListOsConfigurationFileRelationship(p []models.OsConfigurationFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		osconfigurationfilerelationship := flattenMoMoRef(item)
		osconfigurationfilerelationships = append(osconfigurationfilerelationships, osconfigurationfilerelationship)
	}
	return osconfigurationfilerelationships
}
func flattenListOsDistributionRelationship(p []models.OsDistributionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osdistributionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		osdistributionrelationship := flattenMoMoRef(item)
		osdistributionrelationships = append(osdistributionrelationships, osdistributionrelationship)
	}
	return osdistributionrelationships
}
func flattenListOsPlaceHolder(p []models.OsPlaceHolder, d *schema.ResourceData) []map[string]interface{} {
	var osplaceholders []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		osplaceholder := make(map[string]interface{})
		osplaceholder["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osplaceholder["class_id"] = item.ClassId
		osplaceholder["is_value_set"] = item.IsValueSet
		osplaceholder["object_type"] = item.ObjectType
		osplaceholder["type"] = (func(p models.WorkflowPrimitiveDataType, d *schema.ResourceData) []map[string]interface{} {
			var workflowprimitivedatatypes []map[string]interface{}
			var ret models.WorkflowPrimitiveDataType
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowprimitivedatatype := make(map[string]interface{})
			workflowprimitivedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowprimitivedatatype["class_id"] = item.ClassId
			workflowprimitivedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.ClassId
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowprimitivedatatype["description"] = item.Description
			workflowprimitivedatatype["label"] = item.Label
			workflowprimitivedatatype["name"] = item.Name
			workflowprimitivedatatype["object_type"] = item.ObjectType
			workflowprimitivedatatype["properties"] = (func(p models.WorkflowPrimitiveDataProperty, d *schema.ResourceData) []map[string]interface{} {
				var workflowprimitivedatapropertys []map[string]interface{}
				var ret models.WorkflowPrimitiveDataProperty
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowprimitivedataproperty := make(map[string]interface{})
				workflowprimitivedataproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowprimitivedataproperty["class_id"] = item.ClassId
				workflowprimitivedataproperty["constraints"] = (func(p models.WorkflowConstraints, d *schema.ResourceData) []map[string]interface{} {
					var workflowconstraintss []map[string]interface{}
					var ret models.WorkflowConstraints
					if reflect.DeepEqual(ret, p) {
						return nil
					}
					item := p
					workflowconstraints := make(map[string]interface{})
					workflowconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					workflowconstraints["class_id"] = item.ClassId
					workflowconstraints["enum_list"] = (func(p []models.WorkflowEnumEntry, d *schema.ResourceData) []map[string]interface{} {
						var workflowenumentrys []map[string]interface{}
						if len(p) == 0 {
							return nil
						}
						for _, item := range p {
							workflowenumentry := make(map[string]interface{})
							workflowenumentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							workflowenumentry["class_id"] = item.ClassId
							workflowenumentry["label"] = item.Label
							workflowenumentry["object_type"] = item.ObjectType
							workflowenumentry["value"] = item.Value
							workflowenumentrys = append(workflowenumentrys, workflowenumentry)
						}
						return workflowenumentrys
					})(item.GetEnumList(), d)
					workflowconstraints["max"] = item.Max
					workflowconstraints["min"] = item.Min
					workflowconstraints["object_type"] = item.ObjectType
					workflowconstraints["regex"] = item.Regex

					workflowconstraintss = append(workflowconstraintss, workflowconstraints)
					return workflowconstraintss
				})(item.GetConstraints(), d)
				workflowprimitivedataproperty["inventory_selector"] = (func(p []models.WorkflowMoReferenceProperty, d *schema.ResourceData) []map[string]interface{} {
					var workflowmoreferencepropertys []map[string]interface{}
					if len(p) == 0 {
						return nil
					}
					for _, item := range p {
						workflowmoreferenceproperty := make(map[string]interface{})
						workflowmoreferenceproperty["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
						workflowmoreferenceproperty["class_id"] = item.ClassId
						workflowmoreferenceproperty["display_attributes"] = item.DisplayAttributes
						workflowmoreferenceproperty["object_type"] = item.ObjectType
						workflowmoreferenceproperty["selector"] = item.Selector
						workflowmoreferenceproperty["value_attribute"] = item.ValueAttribute
						workflowmoreferencepropertys = append(workflowmoreferencepropertys, workflowmoreferenceproperty)
					}
					return workflowmoreferencepropertys
				})(item.GetInventorySelector(), d)
				workflowprimitivedataproperty["object_type"] = item.ObjectType
				workflowprimitivedataproperty["secure"] = item.Secure
				workflowprimitivedataproperty["type"] = item.Type

				workflowprimitivedatapropertys = append(workflowprimitivedatapropertys, workflowprimitivedataproperty)
				return workflowprimitivedatapropertys
			})(item.GetProperties(), d)
			workflowprimitivedatatype["required"] = item.Required

			workflowprimitivedatatypes = append(workflowprimitivedatatypes, workflowprimitivedatatype)
			return workflowprimitivedatatypes
		})(item.GetType(), d)
		osplaceholder["value"] = item.Value
		osplaceholders = append(osplaceholders, osplaceholder)
	}
	return osplaceholders
}
func flattenListPciCoprocessorCardRelationship(p []models.PciCoprocessorCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcicoprocessorcardrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pcicoprocessorcardrelationship := flattenMoMoRef(item)
		pcicoprocessorcardrelationships = append(pcicoprocessorcardrelationships, pcicoprocessorcardrelationship)
	}
	return pcicoprocessorcardrelationships
}
func flattenListPciDeviceRelationship(p []models.PciDeviceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcidevicerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pcidevicerelationship := flattenMoMoRef(item)
		pcidevicerelationships = append(pcidevicerelationships, pcidevicerelationship)
	}
	return pcidevicerelationships
}
func flattenListPciLinkRelationship(p []models.PciLinkRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pcilinkrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pcilinkrelationship := flattenMoMoRef(item)
		pcilinkrelationships = append(pcilinkrelationships, pcilinkrelationship)
	}
	return pcilinkrelationships
}
func flattenListPciSwitchRelationship(p []models.PciSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		pciswitchrelationship := flattenMoMoRef(item)
		pciswitchrelationships = append(pciswitchrelationships, pciswitchrelationship)
	}
	return pciswitchrelationships
}
func flattenListPolicyAbstractConfigProfileRelationship(p []models.PolicyAbstractConfigProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractconfigprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		policyabstractconfigprofilerelationship := flattenMoMoRef(item)
		policyabstractconfigprofilerelationships = append(policyabstractconfigprofilerelationships, policyabstractconfigprofilerelationship)
	}
	return policyabstractconfigprofilerelationships
}
func flattenListPolicyinventoryJobInfo(p []models.PolicyinventoryJobInfo, d *schema.ResourceData) []map[string]interface{} {
	var policyinventoryjobinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		policyinventoryjobinfo := make(map[string]interface{})
		policyinventoryjobinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		policyinventoryjobinfo["class_id"] = item.ClassId
		policyinventoryjobinfo["execution_status"] = item.ExecutionStatus
		policyinventoryjobinfo["last_scheduled_time"] = item.LastScheduledTime
		policyinventoryjobinfo["object_type"] = item.ObjectType
		policyinventoryjobinfo["policy_id"] = item.PolicyId
		policyinventoryjobinfo["policy_name"] = item.PolicyName
		policyinventoryjobinfos = append(policyinventoryjobinfos, policyinventoryjobinfo)
	}
	return policyinventoryjobinfos
}
func flattenListPortGroupRelationship(p []models.PortGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		portgrouprelationship := flattenMoMoRef(item)
		portgrouprelationships = append(portgrouprelationships, portgrouprelationship)
	}
	return portgrouprelationships
}
func flattenListPortMacBindingRelationship(p []models.PortMacBindingRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portmacbindingrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		portmacbindingrelationship := flattenMoMoRef(item)
		portmacbindingrelationships = append(portmacbindingrelationships, portmacbindingrelationship)
	}
	return portmacbindingrelationships
}
func flattenListPortSubGroupRelationship(p []models.PortSubGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		portsubgrouprelationship := flattenMoMoRef(item)
		portsubgrouprelationships = append(portsubgrouprelationships, portsubgrouprelationship)
	}
	return portsubgrouprelationships
}
func flattenListProcessorUnitRelationship(p []models.ProcessorUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var processorunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		processorunitrelationship := flattenMoMoRef(item)
		processorunitrelationships = append(processorunitrelationships, processorunitrelationship)
	}
	return processorunitrelationships
}
func flattenListRecoveryBackupProfileRelationship(p []models.RecoveryBackupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		recoverybackupprofilerelationship := flattenMoMoRef(item)
		recoverybackupprofilerelationships = append(recoverybackupprofilerelationships, recoverybackupprofilerelationship)
	}
	return recoverybackupprofilerelationships
}
func flattenListRecoveryConfigResultEntryRelationship(p []models.RecoveryConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		recoveryconfigresultentryrelationship := flattenMoMoRef(item)
		recoveryconfigresultentryrelationships = append(recoveryconfigresultentryrelationships, recoveryconfigresultentryrelationship)
	}
	return recoveryconfigresultentryrelationships
}
func flattenListResourceGroupRelationship(p []models.ResourceGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		resourcegrouprelationship := flattenMoMoRef(item)
		resourcegrouprelationships = append(resourcegrouprelationships, resourcegrouprelationship)
	}
	return resourcegrouprelationships
}
func flattenListResourcePerTypeCombinedSelector(p []models.ResourcePerTypeCombinedSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourcepertypecombinedselectors []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		resourcepertypecombinedselector := make(map[string]interface{})
		resourcepertypecombinedselector["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		resourcepertypecombinedselector["class_id"] = item.ClassId
		resourcepertypecombinedselector["combined_selector"] = item.CombinedSelector
		resourcepertypecombinedselector["empty_filter"] = item.EmptyFilter
		resourcepertypecombinedselector["object_type"] = item.ObjectType
		resourcepertypecombinedselector["selector_object_type"] = item.SelectorObjectType
		resourcepertypecombinedselectors = append(resourcepertypecombinedselectors, resourcepertypecombinedselector)
	}
	return resourcepertypecombinedselectors
}
func flattenListResourceSelector(p []models.ResourceSelector, d *schema.ResourceData) []map[string]interface{} {
	var resourceselectors []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		resourceselector := make(map[string]interface{})
		resourceselector["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		resourceselector["class_id"] = item.ClassId
		resourceselector["object_type"] = item.ObjectType
		resourceselector["selector"] = item.Selector
		resourceselectors = append(resourceselectors, resourceselector)
	}
	return resourceselectors
}
func flattenListSdcardPartition(p []models.SdcardPartition, d *schema.ResourceData) []map[string]interface{} {
	var sdcardpartitions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sdcardpartition := make(map[string]interface{})
		sdcardpartition["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sdcardpartition["class_id"] = item.ClassId
		sdcardpartition["object_type"] = item.ObjectType
		sdcardpartition["type"] = item.Type
		sdcardpartition["virtual_drives"] = (func(p []models.SdcardVirtualDrive, d *schema.ResourceData) []map[string]interface{} {
			var sdcardvirtualdrives []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				sdcardvirtualdrive := make(map[string]interface{})
				sdcardvirtualdrive["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				sdcardvirtualdrive["class_id"] = item.ClassId
				sdcardvirtualdrive["enable"] = item.Enable
				sdcardvirtualdrive["object_type"] = item.ObjectType
				sdcardvirtualdrives = append(sdcardvirtualdrives, sdcardvirtualdrive)
			}
			return sdcardvirtualdrives
		})(item.GetVirtualDrives(), d)
		sdcardpartitions = append(sdcardpartitions, sdcardpartition)
	}
	return sdcardpartitions
}
func flattenListSdwanNetworkConfigurationType(p []models.SdwanNetworkConfigurationType, d *schema.ResourceData) []map[string]interface{} {
	var sdwannetworkconfigurationtypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sdwannetworkconfigurationtype := make(map[string]interface{})
		sdwannetworkconfigurationtype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sdwannetworkconfigurationtype["class_id"] = item.ClassId
		sdwannetworkconfigurationtype["network_type"] = item.NetworkType
		sdwannetworkconfigurationtype["object_type"] = item.ObjectType
		sdwannetworkconfigurationtype["port_group"] = item.PortGroup
		sdwannetworkconfigurationtype["vlan"] = item.Vlan
		sdwannetworkconfigurationtypes = append(sdwannetworkconfigurationtypes, sdwannetworkconfigurationtype)
	}
	return sdwannetworkconfigurationtypes
}
func flattenListSdwanProfileRelationship(p []models.SdwanProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		sdwanprofilerelationship := flattenMoMoRef(item)
		sdwanprofilerelationships = append(sdwanprofilerelationships, sdwanprofilerelationship)
	}
	return sdwanprofilerelationships
}
func flattenListSdwanRouterNodeRelationship(p []models.SdwanRouterNodeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouternoderelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		sdwanrouternoderelationship := flattenMoMoRef(item)
		sdwanrouternoderelationships = append(sdwanrouternoderelationships, sdwanrouternoderelationship)
	}
	return sdwanrouternoderelationships
}
func flattenListSdwanTemplateInputsType(p []models.SdwanTemplateInputsType, d *schema.ResourceData) []map[string]interface{} {
	var sdwantemplateinputstypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sdwantemplateinputstype := make(map[string]interface{})
		sdwantemplateinputstype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sdwantemplateinputstype["class_id"] = item.ClassId
		sdwantemplateinputstype["editable"] = item.Editable
		sdwantemplateinputstype["key"] = item.Key
		sdwantemplateinputstype["object_type"] = item.ObjectType
		sdwantemplateinputstype["required"] = item.Required
		sdwantemplateinputstype["template"] = item.Template
		sdwantemplateinputstype["title"] = item.Title
		sdwantemplateinputstype["type"] = item.Type
		sdwantemplateinputstype["value"] = item.Value
		sdwantemplateinputstypes = append(sdwantemplateinputstypes, sdwantemplateinputstype)
	}
	return sdwantemplateinputstypes
}
func flattenListSecurityUnitRelationship(p []models.SecurityUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var securityunitrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		securityunitrelationship := flattenMoMoRef(item)
		securityunitrelationships = append(securityunitrelationships, securityunitrelationship)
	}
	return securityunitrelationships
}
func flattenListServerConfigChangeDetailRelationship(p []models.ServerConfigChangeDetailRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigchangedetailrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		serverconfigchangedetailrelationship := flattenMoMoRef(item)
		serverconfigchangedetailrelationships = append(serverconfigchangedetailrelationships, serverconfigchangedetailrelationship)
	}
	return serverconfigchangedetailrelationships
}
func flattenListServerConfigResultEntryRelationship(p []models.ServerConfigResultEntryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultentryrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		serverconfigresultentryrelationship := flattenMoMoRef(item)
		serverconfigresultentryrelationships = append(serverconfigresultentryrelationships, serverconfigresultentryrelationship)
	}
	return serverconfigresultentryrelationships
}
func flattenListSnmpTrap(p []models.SnmpTrap, d *schema.ResourceData) []map[string]interface{} {
	var snmptraps []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		snmptrap := make(map[string]interface{})
		snmptrap["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		snmptrap["class_id"] = item.ClassId
		snmptrap["destination"] = item.Destination
		snmptrap["enabled"] = item.Enabled
		snmptrap["object_type"] = item.ObjectType
		snmptrap["port"] = item.Port
		snmptrap["type"] = item.Type
		snmptrap["user"] = item.User
		snmptrap["nr_version"] = item.Version
		snmptraps = append(snmptraps, snmptrap)
	}
	return snmptraps
}
func flattenListSnmpUser(p []models.SnmpUser, d *schema.ResourceData) []map[string]interface{} {
	var snmpusers []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		snmpuser := make(map[string]interface{})
		snmpuser["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		auth_password_x := d.Get("snmp_users").([]interface{})
		auth_password_y := auth_password_x[0].(map[string]interface{})
		snmpuser["auth_password"] = auth_password_y["auth_password"]
		snmpuser["auth_type"] = item.AuthType
		snmpuser["class_id"] = item.ClassId
		snmpuser["is_auth_password_set"] = item.IsAuthPasswordSet
		snmpuser["is_privacy_password_set"] = item.IsPrivacyPasswordSet
		snmpuser["name"] = item.Name
		snmpuser["object_type"] = item.ObjectType
		privacy_password_x := d.Get("snmp_users").([]interface{})
		privacy_password_y := privacy_password_x[0].(map[string]interface{})
		snmpuser["privacy_password"] = privacy_password_y["privacy_password"]
		snmpuser["privacy_type"] = item.PrivacyType
		snmpuser["security_level"] = item.SecurityLevel
		snmpusers = append(snmpusers, snmpuser)
	}
	return snmpusers
}
func flattenListStorageBaseInitiator(p []models.StorageBaseInitiator, d *schema.ResourceData) []map[string]interface{} {
	var storagebaseinitiators []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagebaseinitiator := make(map[string]interface{})
		storagebaseinitiator["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagebaseinitiator["class_id"] = item.ClassId
		storagebaseinitiator["iqn"] = item.Iqn
		storagebaseinitiator["name"] = item.Name
		storagebaseinitiator["object_type"] = item.ObjectType
		storagebaseinitiator["type"] = item.Type
		storagebaseinitiator["wwn"] = item.Wwn
		storagebaseinitiators = append(storagebaseinitiators, storagebaseinitiator)
	}
	return storagebaseinitiators
}
func flattenListStorageControllerRelationship(p []models.StorageControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagecontrollerrelationship := flattenMoMoRef(item)
		storagecontrollerrelationships = append(storagecontrollerrelationships, storagecontrollerrelationship)
	}
	return storagecontrollerrelationships
}
func flattenListStorageDiskGroupRelationship(p []models.StorageDiskGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagediskgrouprelationship := flattenMoMoRef(item)
		storagediskgrouprelationships = append(storagediskgrouprelationships, storagediskgrouprelationship)
	}
	return storagediskgrouprelationships
}
func flattenListStorageDiskGroupPolicyRelationship(p []models.StorageDiskGroupPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouppolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagediskgrouppolicyrelationship := flattenMoMoRef(item)
		storagediskgrouppolicyrelationships = append(storagediskgrouppolicyrelationships, storagediskgrouppolicyrelationship)
	}
	return storagediskgrouppolicyrelationships
}
func flattenListStorageEnclosureRelationship(p []models.StorageEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageenclosurerelationship := flattenMoMoRef(item)
		storageenclosurerelationships = append(storageenclosurerelationships, storageenclosurerelationship)
	}
	return storageenclosurerelationships
}
func flattenListStorageEnclosureDiskRelationship(p []models.StorageEnclosureDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurediskrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageenclosurediskrelationship := flattenMoMoRef(item)
		storageenclosurediskrelationships = append(storageenclosurediskrelationships, storageenclosurediskrelationship)
	}
	return storageenclosurediskrelationships
}
func flattenListStorageEnclosureDiskSlotEpRelationship(p []models.StorageEnclosureDiskSlotEpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosuredisksloteprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageenclosuredisksloteprelationship := flattenMoMoRef(item)
		storageenclosuredisksloteprelationships = append(storageenclosuredisksloteprelationships, storageenclosuredisksloteprelationship)
	}
	return storageenclosuredisksloteprelationships
}
func flattenListStorageFlexFlashControllerRelationship(p []models.StorageFlexFlashControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashcontrollerrelationship := flattenMoMoRef(item)
		storageflexflashcontrollerrelationships = append(storageflexflashcontrollerrelationships, storageflexflashcontrollerrelationship)
	}
	return storageflexflashcontrollerrelationships
}
func flattenListStorageFlexFlashControllerPropsRelationship(p []models.StorageFlexFlashControllerPropsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerpropsrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashcontrollerpropsrelationship := flattenMoMoRef(item)
		storageflexflashcontrollerpropsrelationships = append(storageflexflashcontrollerpropsrelationships, storageflexflashcontrollerpropsrelationship)
	}
	return storageflexflashcontrollerpropsrelationships
}
func flattenListStorageFlexFlashPhysicalDriveRelationship(p []models.StorageFlexFlashPhysicalDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashphysicaldriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashphysicaldriverelationship := flattenMoMoRef(item)
		storageflexflashphysicaldriverelationships = append(storageflexflashphysicaldriverelationships, storageflexflashphysicaldriverelationship)
	}
	return storageflexflashphysicaldriverelationships
}
func flattenListStorageFlexFlashVirtualDriveRelationship(p []models.StorageFlexFlashVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashvirtualdriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexflashvirtualdriverelationship := flattenMoMoRef(item)
		storageflexflashvirtualdriverelationships = append(storageflexflashvirtualdriverelationships, storageflexflashvirtualdriverelationship)
	}
	return storageflexflashvirtualdriverelationships
}
func flattenListStorageFlexUtilControllerRelationship(p []models.StorageFlexUtilControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexutilcontrollerrelationship := flattenMoMoRef(item)
		storageflexutilcontrollerrelationships = append(storageflexutilcontrollerrelationships, storageflexutilcontrollerrelationship)
	}
	return storageflexutilcontrollerrelationships
}
func flattenListStorageFlexUtilPhysicalDriveRelationship(p []models.StorageFlexUtilPhysicalDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilphysicaldriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexutilphysicaldriverelationship := flattenMoMoRef(item)
		storageflexutilphysicaldriverelationships = append(storageflexutilphysicaldriverelationships, storageflexutilphysicaldriverelationship)
	}
	return storageflexutilphysicaldriverelationships
}
func flattenListStorageFlexUtilVirtualDriveRelationship(p []models.StorageFlexUtilVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilvirtualdriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageflexutilvirtualdriverelationship := flattenMoMoRef(item)
		storageflexutilvirtualdriverelationships = append(storageflexutilvirtualdriverelationships, storageflexutilvirtualdriverelationship)
	}
	return storageflexutilvirtualdriverelationships
}
func flattenListStorageItemRelationship(p []models.StorageItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageitemrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storageitemrelationship := flattenMoMoRef(item)
		storageitemrelationships = append(storageitemrelationships, storageitemrelationship)
	}
	return storageitemrelationships
}
func flattenListStorageLocalDisk(p []models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
	var storagelocaldisks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagelocaldisk := make(map[string]interface{})
		storagelocaldisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagelocaldisk["class_id"] = item.ClassId
		storagelocaldisk["object_type"] = item.ObjectType
		storagelocaldisk["slot_number"] = item.SlotNumber
		storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
	}
	return storagelocaldisks
}
func flattenListStoragePhysicalDiskRelationship(p []models.StoragePhysicalDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagephysicaldiskrelationship := flattenMoMoRef(item)
		storagephysicaldiskrelationships = append(storagephysicaldiskrelationships, storagephysicaldiskrelationship)
	}
	return storagephysicaldiskrelationships
}
func flattenListStoragePhysicalDiskExtensionRelationship(p []models.StoragePhysicalDiskExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskextensionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagephysicaldiskextensionrelationship := flattenMoMoRef(item)
		storagephysicaldiskextensionrelationships = append(storagephysicaldiskextensionrelationships, storagephysicaldiskextensionrelationship)
	}
	return storagephysicaldiskextensionrelationships
}
func flattenListStoragePhysicalDiskUsageRelationship(p []models.StoragePhysicalDiskUsageRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskusagerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagephysicaldiskusagerelationship := flattenMoMoRef(item)
		storagephysicaldiskusagerelationships = append(storagephysicaldiskusagerelationships, storagephysicaldiskusagerelationship)
	}
	return storagephysicaldiskusagerelationships
}
func flattenListStoragePureHostRelationship(p []models.StoragePureHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagepurehostrelationship := flattenMoMoRef(item)
		storagepurehostrelationships = append(storagepurehostrelationships, storagepurehostrelationship)
	}
	return storagepurehostrelationships
}
func flattenListStoragePureHostGroupRelationship(p []models.StoragePureHostGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagepurehostgrouprelationship := flattenMoMoRef(item)
		storagepurehostgrouprelationships = append(storagepurehostgrouprelationships, storagepurehostgrouprelationship)
	}
	return storagepurehostgrouprelationships
}
func flattenListStoragePureReplicationBlackout(p []models.StoragePureReplicationBlackout, d *schema.ResourceData) []map[string]interface{} {
	var storagepurereplicationblackouts []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagepurereplicationblackout := make(map[string]interface{})
		storagepurereplicationblackout["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagepurereplicationblackout["class_id"] = item.ClassId
		storagepurereplicationblackout["end"] = item.End
		storagepurereplicationblackout["object_type"] = item.ObjectType
		storagepurereplicationblackout["start"] = item.Start
		storagepurereplicationblackouts = append(storagepurereplicationblackouts, storagepurereplicationblackout)
	}
	return storagepurereplicationblackouts
}
func flattenListStoragePureVolumeRelationship(p []models.StoragePureVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurevolumerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagepurevolumerelationship := flattenMoMoRef(item)
		storagepurevolumerelationships = append(storagepurevolumerelationships, storagepurevolumerelationship)
	}
	return storagepurevolumerelationships
}
func flattenListStorageSasExpanderRelationship(p []models.StorageSasExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagesasexpanderrelationship := flattenMoMoRef(item)
		storagesasexpanderrelationships = append(storagesasexpanderrelationships, storagesasexpanderrelationship)
	}
	return storagesasexpanderrelationships
}
func flattenListStorageSasPortRelationship(p []models.StorageSasPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasportrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagesasportrelationship := flattenMoMoRef(item)
		storagesasportrelationships = append(storagesasportrelationships, storagesasportrelationship)
	}
	return storagesasportrelationships
}
func flattenListStorageSpanRelationship(p []models.StorageSpanRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagespanrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagespanrelationship := flattenMoMoRef(item)
		storagespanrelationships = append(storagespanrelationships, storagespanrelationship)
	}
	return storagespanrelationships
}
func flattenListStorageSpanGroup(p []models.StorageSpanGroup, d *schema.ResourceData) []map[string]interface{} {
	var storagespangroups []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagespangroup := make(map[string]interface{})
		storagespangroup["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagespangroup["class_id"] = item.ClassId
		storagespangroup["disks"] = (func(p []models.StorageLocalDisk, d *schema.ResourceData) []map[string]interface{} {
			var storagelocaldisks []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				storagelocaldisk := make(map[string]interface{})
				storagelocaldisk["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				storagelocaldisk["class_id"] = item.ClassId
				storagelocaldisk["object_type"] = item.ObjectType
				storagelocaldisk["slot_number"] = item.SlotNumber
				storagelocaldisks = append(storagelocaldisks, storagelocaldisk)
			}
			return storagelocaldisks
		})(item.GetDisks(), d)
		storagespangroup["object_type"] = item.ObjectType
		storagespangroups = append(storagespangroups, storagespangroup)
	}
	return storagespangroups
}
func flattenListStorageStoragePolicyRelationship(p []models.StorageStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagestoragepolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagestoragepolicyrelationship := flattenMoMoRef(item)
		storagestoragepolicyrelationships = append(storagestoragepolicyrelationships, storagestoragepolicyrelationship)
	}
	return storagestoragepolicyrelationships
}
func flattenListStorageVdMemberEpRelationship(p []models.StorageVdMemberEpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevdmembereprelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagevdmembereprelationship := flattenMoMoRef(item)
		storagevdmembereprelationships = append(storagevdmembereprelationships, storagevdmembereprelationship)
	}
	return storagevdmembereprelationships
}
func flattenListStorageVirtualDriveRelationship(p []models.StorageVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagevirtualdriverelationship := flattenMoMoRef(item)
		storagevirtualdriverelationships = append(storagevirtualdriverelationships, storagevirtualdriverelationship)
	}
	return storagevirtualdriverelationships
}
func flattenListStorageVirtualDriveConfig(p []models.StorageVirtualDriveConfig, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveconfigs []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		storagevirtualdriveconfig := make(map[string]interface{})
		storagevirtualdriveconfig["access_policy"] = item.AccessPolicy
		storagevirtualdriveconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		storagevirtualdriveconfig["boot_drive"] = item.BootDrive
		storagevirtualdriveconfig["class_id"] = item.ClassId
		storagevirtualdriveconfig["disk_group_name"] = item.DiskGroupName
		storagevirtualdriveconfig["disk_group_policy"] = item.DiskGroupPolicy
		storagevirtualdriveconfig["drive_cache"] = item.DriveCache
		storagevirtualdriveconfig["expand_to_available"] = item.ExpandToAvailable
		storagevirtualdriveconfig["io_policy"] = item.IoPolicy
		storagevirtualdriveconfig["name"] = item.Name
		storagevirtualdriveconfig["object_type"] = item.ObjectType
		storagevirtualdriveconfig["read_policy"] = item.ReadPolicy
		storagevirtualdriveconfig["size"] = item.Size
		storagevirtualdriveconfig["write_policy"] = item.WritePolicy
		storagevirtualdriveconfigs = append(storagevirtualdriveconfigs, storagevirtualdriveconfig)
	}
	return storagevirtualdriveconfigs
}
func flattenListStorageVirtualDriveExtensionRelationship(p []models.StorageVirtualDriveExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		storagevirtualdriveextensionrelationship := flattenMoMoRef(item)
		storagevirtualdriveextensionrelationships = append(storagevirtualdriveextensionrelationships, storagevirtualdriveextensionrelationship)
	}
	return storagevirtualdriveextensionrelationships
}
func flattenListSyslogLocalClientBase(p []models.SyslogLocalClientBase, d *schema.ResourceData) []map[string]interface{} {
	var sysloglocalclientbases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		sysloglocalclientbase := make(map[string]interface{})
		sysloglocalclientbase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		sysloglocalclientbase["class_id"] = item.ClassId
		sysloglocalclientbase["min_severity"] = item.MinSeverity
		sysloglocalclientbase["object_type"] = item.ObjectType
		sysloglocalclientbases = append(sysloglocalclientbases, sysloglocalclientbase)
	}
	return sysloglocalclientbases
}
func flattenListSyslogRemoteClientBase(p []models.SyslogRemoteClientBase, d *schema.ResourceData) []map[string]interface{} {
	var syslogremoteclientbases []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		syslogremoteclientbase := make(map[string]interface{})
		syslogremoteclientbase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		syslogremoteclientbase["class_id"] = item.ClassId
		syslogremoteclientbase["enabled"] = item.Enabled
		syslogremoteclientbase["hostname"] = item.Hostname
		syslogremoteclientbase["min_severity"] = item.MinSeverity
		syslogremoteclientbase["object_type"] = item.ObjectType
		syslogremoteclientbase["port"] = item.Port
		syslogremoteclientbase["protocol"] = item.Protocol
		syslogremoteclientbases = append(syslogremoteclientbases, syslogremoteclientbase)
	}
	return syslogremoteclientbases
}
func flattenListTamAction(p []models.TamAction, d *schema.ResourceData) []map[string]interface{} {
	var tamactions []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		tamaction := make(map[string]interface{})
		tamaction["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		tamaction["affected_object_type"] = item.AffectedObjectType
		tamaction["alert_type"] = item.AlertType
		tamaction["class_id"] = item.ClassId
		tamaction["identifiers"] = (func(p []models.TamIdentifiers, d *schema.ResourceData) []map[string]interface{} {
			var tamidentifierss []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamidentifiers := make(map[string]interface{})
				tamidentifiers["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamidentifiers["class_id"] = item.ClassId
				tamidentifiers["name"] = item.Name
				tamidentifiers["object_type"] = item.ObjectType
				tamidentifiers["value"] = item.Value
				tamidentifierss = append(tamidentifierss, tamidentifiers)
			}
			return tamidentifierss
		})(item.GetIdentifiers(), d)
		tamaction["name"] = item.Name
		tamaction["object_type"] = item.ObjectType
		tamaction["operation_type"] = item.OperationType
		tamaction["queries"] = (func(p []models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamqueryentry["class_id"] = item.ClassId
				tamqueryentry["name"] = item.Name
				tamqueryentry["object_type"] = item.ObjectType
				tamqueryentry["priority"] = item.Priority
				tamqueryentry["query"] = item.Query
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.GetQueries(), d)
		tamaction["type"] = item.Type
		tamactions = append(tamactions, tamaction)
	}
	return tamactions
}
func flattenListTamApiDataSource(p []models.TamApiDataSource, d *schema.ResourceData) []map[string]interface{} {
	var tamapidatasources []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		tamapidatasource := make(map[string]interface{})
		tamapidatasource["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		tamapidatasource["class_id"] = item.ClassId
		tamapidatasource["mo_type"] = item.MoType
		tamapidatasource["name"] = item.Name
		tamapidatasource["object_type"] = item.ObjectType
		tamapidatasource["queries"] = (func(p []models.TamQueryEntry, d *schema.ResourceData) []map[string]interface{} {
			var tamqueryentrys []map[string]interface{}
			if len(p) == 0 {
				return nil
			}
			for _, item := range p {
				tamqueryentry := make(map[string]interface{})
				tamqueryentry["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				tamqueryentry["class_id"] = item.ClassId
				tamqueryentry["name"] = item.Name
				tamqueryentry["object_type"] = item.ObjectType
				tamqueryentry["priority"] = item.Priority
				tamqueryentry["query"] = item.Query
				tamqueryentrys = append(tamqueryentrys, tamqueryentry)
			}
			return tamqueryentrys
		})(item.GetQueries(), d)
		tamapidatasource["type"] = item.Type
		tamapidatasources = append(tamapidatasources, tamapidatasource)
	}
	return tamapidatasources
}
func flattenListUcsdConnectorPack(p []models.UcsdConnectorPack, d *schema.ResourceData) []map[string]interface{} {
	var ucsdconnectorpacks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		ucsdconnectorpack := make(map[string]interface{})
		ucsdconnectorpack["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		ucsdconnectorpack["class_id"] = item.ClassId
		ucsdconnectorpack["connector_feature"] = item.ConnectorFeature
		ucsdconnectorpack["dependency_names"] = item.DependencyNames
		ucsdconnectorpack["downloaded_version"] = item.DownloadedVersion
		ucsdconnectorpack["name"] = item.Name
		ucsdconnectorpack["object_type"] = item.ObjectType
		ucsdconnectorpack["services"] = item.Services
		ucsdconnectorpack["state"] = item.State
		ucsdconnectorpack["nr_version"] = item.Version
		ucsdconnectorpacks = append(ucsdconnectorpacks, ucsdconnectorpack)
	}
	return ucsdconnectorpacks
}
func flattenListUuidpoolBlockRelationship(p []models.UuidpoolBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolblockrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		uuidpoolblockrelationship := flattenMoMoRef(item)
		uuidpoolblockrelationships = append(uuidpoolblockrelationships, uuidpoolblockrelationship)
	}
	return uuidpoolblockrelationships
}
func flattenListUuidpoolUuidBlock(p []models.UuidpoolUuidBlock, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluuidblocks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		uuidpooluuidblock := make(map[string]interface{})
		uuidpooluuidblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		uuidpooluuidblock["class_id"] = item.ClassId
		uuidpooluuidblock["from"] = item.From
		uuidpooluuidblock["object_type"] = item.ObjectType
		uuidpooluuidblock["size"] = item.Size
		uuidpooluuidblock["to"] = item.To
		uuidpooluuidblocks = append(uuidpooluuidblocks, uuidpooluuidblock)
	}
	return uuidpooluuidblocks
}
func flattenListVirtualizationVmwareDatastoreRelationship(p []models.VirtualizationVmwareDatastoreRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatastorerelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		virtualizationvmwaredatastorerelationship := flattenMoMoRef(item)
		virtualizationvmwaredatastorerelationships = append(virtualizationvmwaredatastorerelationships, virtualizationvmwaredatastorerelationship)
	}
	return virtualizationvmwaredatastorerelationships
}
func flattenListVirtualizationVmwareHostRelationship(p []models.VirtualizationVmwareHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		virtualizationvmwarehostrelationship := flattenMoMoRef(item)
		virtualizationvmwarehostrelationships = append(virtualizationvmwarehostrelationships, virtualizationvmwarehostrelationship)
	}
	return virtualizationvmwarehostrelationships
}
func flattenListVmediaMapping(p []models.VmediaMapping, d *schema.ResourceData) []map[string]interface{} {
	var vmediamappings []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		vmediamapping := make(map[string]interface{})
		vmediamapping["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		vmediamapping["authentication_protocol"] = item.AuthenticationProtocol
		vmediamapping["class_id"] = item.ClassId
		vmediamapping["device_type"] = item.DeviceType
		vmediamapping["host_name"] = item.HostName
		vmediamapping["is_password_set"] = item.IsPasswordSet
		vmediamapping["mount_options"] = item.MountOptions
		vmediamapping["mount_protocol"] = item.MountProtocol
		vmediamapping["object_type"] = item.ObjectType
		password_x := d.Get("mappings").([]interface{})
		password_y := password_x[0].(map[string]interface{})
		vmediamapping["password"] = password_y["password"]
		vmediamapping["remote_file"] = item.RemoteFile
		vmediamapping["remote_path"] = item.RemotePath
		vmediamapping["username"] = item.Username
		vmediamapping["volume_name"] = item.VolumeName
		vmediamappings = append(vmediamappings, vmediamapping)
	}
	return vmediamappings
}
func flattenListVnicEthIfRelationship(p []models.VnicEthIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethifrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		vnicethifrelationship := flattenMoMoRef(item)
		vnicethifrelationships = append(vnicethifrelationships, vnicethifrelationship)
	}
	return vnicethifrelationships
}
func flattenListVnicEthNetworkPolicyRelationship(p []models.VnicEthNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethnetworkpolicyrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		vnicethnetworkpolicyrelationship := flattenMoMoRef(item)
		vnicethnetworkpolicyrelationships = append(vnicethnetworkpolicyrelationships, vnicethnetworkpolicyrelationship)
	}
	return vnicethnetworkpolicyrelationships
}
func flattenListVnicFcIfRelationship(p []models.VnicFcIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcifrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		vnicfcifrelationship := flattenMoMoRef(item)
		vnicfcifrelationships = append(vnicfcifrelationships, vnicfcifrelationship)
	}
	return vnicfcifrelationships
}
func flattenListWorkflowApi(p []models.WorkflowApi, d *schema.ResourceData) []map[string]interface{} {
	var workflowapis []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowapi := make(map[string]interface{})
		workflowapi["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowapi["body"] = item.Body
		workflowapi["class_id"] = item.ClassId
		workflowapi["content_type"] = item.ContentType
		workflowapi["name"] = item.Name
		workflowapi["object_type"] = item.ObjectType
		workflowapi["outcomes"] = item.Outcomes
		workflowapi["response_spec"] = (func(p models.ContentGrammar, d *schema.ResourceData) []map[string]interface{} {
			var contentgrammars []map[string]interface{}
			var ret models.ContentGrammar
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			contentgrammar := make(map[string]interface{})
			contentgrammar["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			contentgrammar["class_id"] = item.ClassId
			contentgrammar["error_parameters"] = (func(p []models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var contentbaseparameters []map[string]interface{}
				if len(p) == 0 {
					return nil
				}
				for _, item := range p {
					contentbaseparameter := make(map[string]interface{})
					contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
					contentbaseparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					contentbaseparameter["class_id"] = item.ClassId
					contentbaseparameter["complex_type"] = item.ComplexType
					contentbaseparameter["item_type"] = item.ItemType
					contentbaseparameter["name"] = item.Name
					contentbaseparameter["object_type"] = item.ObjectType
					contentbaseparameter["path"] = item.Path
					contentbaseparameter["type"] = item.Type
					contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
				}
				return contentbaseparameters
			})(item.GetErrorParameters(), d)
			contentgrammar["object_type"] = item.ObjectType
			contentgrammar["parameters"] = (func(p []models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
				var contentbaseparameters []map[string]interface{}
				if len(p) == 0 {
					return nil
				}
				for _, item := range p {
					contentbaseparameter := make(map[string]interface{})
					contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
					contentbaseparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					contentbaseparameter["class_id"] = item.ClassId
					contentbaseparameter["complex_type"] = item.ComplexType
					contentbaseparameter["item_type"] = item.ItemType
					contentbaseparameter["name"] = item.Name
					contentbaseparameter["object_type"] = item.ObjectType
					contentbaseparameter["path"] = item.Path
					contentbaseparameter["type"] = item.Type
					contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
				}
				return contentbaseparameters
			})(item.GetParameters(), d)
			contentgrammar["types"] = (func(p []models.ContentComplexType, d *schema.ResourceData) []map[string]interface{} {
				var contentcomplextypes []map[string]interface{}
				if len(p) == 0 {
					return nil
				}
				for _, item := range p {
					contentcomplextype := make(map[string]interface{})
					contentcomplextype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
					contentcomplextype["class_id"] = item.ClassId
					contentcomplextype["name"] = item.Name
					contentcomplextype["object_type"] = item.ObjectType
					contentcomplextype["parameters"] = (func(p []models.ContentBaseParameter, d *schema.ResourceData) []map[string]interface{} {
						var contentbaseparameters []map[string]interface{}
						if len(p) == 0 {
							return nil
						}
						for _, item := range p {
							contentbaseparameter := make(map[string]interface{})
							contentbaseparameter["accept_single_value"] = item.AcceptSingleValue
							contentbaseparameter["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
							contentbaseparameter["class_id"] = item.ClassId
							contentbaseparameter["complex_type"] = item.ComplexType
							contentbaseparameter["item_type"] = item.ItemType
							contentbaseparameter["name"] = item.Name
							contentbaseparameter["object_type"] = item.ObjectType
							contentbaseparameter["path"] = item.Path
							contentbaseparameter["type"] = item.Type
							contentbaseparameters = append(contentbaseparameters, contentbaseparameter)
						}
						return contentbaseparameters
					})(item.GetParameters(), d)
					contentcomplextypes = append(contentcomplextypes, contentcomplextype)
				}
				return contentcomplextypes
			})(item.GetTypes(), d)

			contentgrammars = append(contentgrammars, contentgrammar)
			return contentgrammars
		})(item.GetResponseSpec(), d)
		workflowapi["skip_on_condition"] = item.SkipOnCondition
		workflowapi["start_delay"] = item.StartDelay
		workflowapi["timeout"] = item.Timeout
		workflowapis = append(workflowapis, workflowapi)
	}
	return workflowapis
}
func flattenListWorkflowBaseDataType(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
	var workflowbasedatatypes []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowbasedatatype := make(map[string]interface{})
		workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowbasedatatype["class_id"] = item.ClassId
		workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
			var workflowdefaultvalues []map[string]interface{}
			var ret models.WorkflowDefaultValue
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			workflowdefaultvalue := make(map[string]interface{})
			workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowdefaultvalue["class_id"] = item.ClassId
			workflowdefaultvalue["object_type"] = item.ObjectType
			workflowdefaultvalue["override"] = item.Override
			workflowdefaultvalue["value"] = item.Value

			workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
			return workflowdefaultvalues
		})(item.GetDefault(), d)
		workflowbasedatatype["description"] = item.Description
		workflowbasedatatype["label"] = item.Label
		workflowbasedatatype["name"] = item.Name
		workflowbasedatatype["object_type"] = item.ObjectType
		workflowbasedatatype["required"] = item.Required
		workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
	}
	return workflowbasedatatypes
}
func flattenListWorkflowDynamicWorkflowActionTaskList(p []models.WorkflowDynamicWorkflowActionTaskList, d *schema.ResourceData) []map[string]interface{} {
	var workflowdynamicworkflowactiontasklists []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowdynamicworkflowactiontasklist := make(map[string]interface{})
		workflowdynamicworkflowactiontasklist["action"] = item.Action
		workflowdynamicworkflowactiontasklist["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowdynamicworkflowactiontasklist["class_id"] = item.ClassId
		workflowdynamicworkflowactiontasklist["object_type"] = item.ObjectType
		workflowdynamicworkflowactiontasklist["tasks"] = item.Tasks
		workflowdynamicworkflowactiontasklists = append(workflowdynamicworkflowactiontasklists, workflowdynamicworkflowactiontasklist)
	}
	return workflowdynamicworkflowactiontasklists
}
func flattenListWorkflowMessage(p []models.WorkflowMessage, d *schema.ResourceData) []map[string]interface{} {
	var workflowmessages []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowmessage := make(map[string]interface{})
		workflowmessage["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowmessage["class_id"] = item.ClassId
		workflowmessage["message"] = item.Message
		workflowmessage["object_type"] = item.ObjectType
		workflowmessage["severity"] = item.Severity
		workflowmessages = append(workflowmessages, workflowmessage)
	}
	return workflowmessages
}
func flattenListWorkflowTaskDefinitionRelationship(p []models.WorkflowTaskDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		workflowtaskdefinitionrelationship := flattenMoMoRef(item)
		workflowtaskdefinitionrelationships = append(workflowtaskdefinitionrelationships, workflowtaskdefinitionrelationship)
	}
	return workflowtaskdefinitionrelationships
}
func flattenListWorkflowTaskInfoRelationship(p []models.WorkflowTaskInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		workflowtaskinforelationship := flattenMoMoRef(item)
		workflowtaskinforelationships = append(workflowtaskinforelationships, workflowtaskinforelationship)
	}
	return workflowtaskinforelationships
}
func flattenListWorkflowTaskRetryInfo(p []models.WorkflowTaskRetryInfo, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskretryinfos []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowtaskretryinfo := make(map[string]interface{})
		workflowtaskretryinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowtaskretryinfo["class_id"] = item.ClassId
		workflowtaskretryinfo["object_type"] = item.ObjectType
		workflowtaskretryinfo["status"] = item.Status
		workflowtaskretryinfo["task_inst_id"] = item.TaskInstId
		workflowtaskretryinfos = append(workflowtaskretryinfos, workflowtaskretryinfo)
	}
	return workflowtaskretryinfos
}
func flattenListWorkflowWorkflowInfoRelationship(p []models.WorkflowWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforelationships []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		item := item.MoMoRef
		workflowworkflowinforelationship := flattenMoMoRef(item)
		workflowworkflowinforelationships = append(workflowworkflowinforelationships, workflowworkflowinforelationship)
	}
	return workflowworkflowinforelationships
}
func flattenListWorkflowWorkflowTask(p []models.WorkflowWorkflowTask, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowtasks []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		workflowworkflowtask := make(map[string]interface{})
		workflowworkflowtask["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowworkflowtask["class_id"] = item.ClassId
		workflowworkflowtask["description"] = item.Description
		workflowworkflowtask["label"] = item.Label
		workflowworkflowtask["name"] = item.Name
		workflowworkflowtask["object_type"] = item.ObjectType
		workflowworkflowtasks = append(workflowworkflowtasks, workflowworkflowtask)
	}
	return workflowworkflowtasks
}
func flattenListX509Certificate(p []models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var x509certificates []map[string]interface{}
	if len(p) == 0 {
		return nil
	}
	for _, item := range p {
		x509certificate := make(map[string]interface{})
		x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		x509certificate["class_id"] = item.ClassId
		x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetIssuer(), d)
		x509certificate["object_type"] = item.ObjectType
		x509certificate["pem_certificate"] = item.PemCertificate
		x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
		x509certificate["signature_algorithm"] = item.SignatureAlgorithm
		x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetSubject(), d)
		x509certificates = append(x509certificates, x509certificate)
	}
	return x509certificates
}
func flattenMapAdapterUnitRelationship(p models.AdapterUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var adapterunitrelationships []map[string]interface{}
	var ret models.AdapterUnitRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	adapterunitrelationship := make(map[string]interface{})
	adapterunitrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	adapterunitrelationship["class_id"] = item.ClassId
	adapterunitrelationship["moid"] = item.Moid
	adapterunitrelationship["object_type"] = item.ObjectType
	adapterunitrelationship["selector"] = item.Selector

	adapterunitrelationships = append(adapterunitrelationships, adapterunitrelationship)
	return adapterunitrelationships
}
func flattenMapApplianceDataExportPolicyRelationship(p models.ApplianceDataExportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var appliancedataexportpolicyrelationships []map[string]interface{}
	var ret models.ApplianceDataExportPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	appliancedataexportpolicyrelationship := make(map[string]interface{})
	appliancedataexportpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	appliancedataexportpolicyrelationship["class_id"] = item.ClassId
	appliancedataexportpolicyrelationship["moid"] = item.Moid
	appliancedataexportpolicyrelationship["object_type"] = item.ObjectType
	appliancedataexportpolicyrelationship["selector"] = item.Selector

	appliancedataexportpolicyrelationships = append(appliancedataexportpolicyrelationships, appliancedataexportpolicyrelationship)
	return appliancedataexportpolicyrelationships
}
func flattenMapApplianceImageBundleRelationship(p models.ApplianceImageBundleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var applianceimagebundlerelationships []map[string]interface{}
	var ret models.ApplianceImageBundleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	applianceimagebundlerelationship := make(map[string]interface{})
	applianceimagebundlerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	applianceimagebundlerelationship["class_id"] = item.ClassId
	applianceimagebundlerelationship["moid"] = item.Moid
	applianceimagebundlerelationship["object_type"] = item.ObjectType
	applianceimagebundlerelationship["selector"] = item.Selector

	applianceimagebundlerelationships = append(applianceimagebundlerelationships, applianceimagebundlerelationship)
	return applianceimagebundlerelationships
}
func flattenMapAssetClusterMemberRelationship(p models.AssetClusterMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetclustermemberrelationships []map[string]interface{}
	var ret models.AssetClusterMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetclustermemberrelationship := make(map[string]interface{})
	assetclustermemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetclustermemberrelationship["class_id"] = item.ClassId
	assetclustermemberrelationship["moid"] = item.Moid
	assetclustermemberrelationship["object_type"] = item.ObjectType
	assetclustermemberrelationship["selector"] = item.Selector

	assetclustermemberrelationships = append(assetclustermemberrelationships, assetclustermemberrelationship)
	return assetclustermemberrelationships
}
func flattenMapAssetContractInformation(p models.AssetContractInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetcontractinformations []map[string]interface{}
	var ret models.AssetContractInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetcontractinformation := make(map[string]interface{})
	assetcontractinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetcontractinformation["bill_to"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetBillTo(), d)
	assetcontractinformation["bill_to_global_ultimate"] = (func(p models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {
		var assetglobalultimates []map[string]interface{}
		var ret models.AssetGlobalUltimate
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetglobalultimate := make(map[string]interface{})
		assetglobalultimate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetglobalultimate["class_id"] = item.ClassId
		assetglobalultimate["id"] = item.Id
		assetglobalultimate["name"] = item.Name
		assetglobalultimate["object_type"] = item.ObjectType

		assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
		return assetglobalultimates
	})(item.GetBillToGlobalUltimate(), d)
	assetcontractinformation["class_id"] = item.ClassId
	assetcontractinformation["contract_number"] = item.ContractNumber
	assetcontractinformation["line_status"] = item.LineStatus
	assetcontractinformation["object_type"] = item.ObjectType

	assetcontractinformations = append(assetcontractinformations, assetcontractinformation)
	return assetcontractinformations
}
func flattenMapAssetCustomerInformation(p models.AssetCustomerInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetcustomerinformations []map[string]interface{}
	var ret models.AssetCustomerInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetcustomerinformation := make(map[string]interface{})
	assetcustomerinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetcustomerinformation["address"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetAddress(), d)
	assetcustomerinformation["class_id"] = item.ClassId
	assetcustomerinformation["id"] = item.Id
	assetcustomerinformation["name"] = item.Name
	assetcustomerinformation["object_type"] = item.ObjectType

	assetcustomerinformations = append(assetcustomerinformations, assetcustomerinformation)
	return assetcustomerinformations
}
func flattenMapAssetDeviceClaimRelationship(p models.AssetDeviceClaimRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceclaimrelationships []map[string]interface{}
	var ret models.AssetDeviceClaimRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceclaimrelationship := make(map[string]interface{})
	assetdeviceclaimrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceclaimrelationship["class_id"] = item.ClassId
	assetdeviceclaimrelationship["moid"] = item.Moid
	assetdeviceclaimrelationship["object_type"] = item.ObjectType
	assetdeviceclaimrelationship["selector"] = item.Selector

	assetdeviceclaimrelationships = append(assetdeviceclaimrelationships, assetdeviceclaimrelationship)
	return assetdeviceclaimrelationships
}
func flattenMapAssetDeviceConfigurationRelationship(p models.AssetDeviceConfigurationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconfigurationrelationships []map[string]interface{}
	var ret models.AssetDeviceConfigurationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceconfigurationrelationship := make(map[string]interface{})
	assetdeviceconfigurationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceconfigurationrelationship["class_id"] = item.ClassId
	assetdeviceconfigurationrelationship["moid"] = item.Moid
	assetdeviceconfigurationrelationship["object_type"] = item.ObjectType
	assetdeviceconfigurationrelationship["selector"] = item.Selector

	assetdeviceconfigurationrelationships = append(assetdeviceconfigurationrelationships, assetdeviceconfigurationrelationship)
	return assetdeviceconfigurationrelationships
}
func flattenMapAssetDeviceConnectionRelationship(p models.AssetDeviceConnectionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceconnectionrelationships []map[string]interface{}
	var ret models.AssetDeviceConnectionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceconnectionrelationship := make(map[string]interface{})
	assetdeviceconnectionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceconnectionrelationship["class_id"] = item.ClassId
	assetdeviceconnectionrelationship["moid"] = item.Moid
	assetdeviceconnectionrelationship["object_type"] = item.ObjectType
	assetdeviceconnectionrelationship["selector"] = item.Selector

	assetdeviceconnectionrelationships = append(assetdeviceconnectionrelationships, assetdeviceconnectionrelationship)
	return assetdeviceconnectionrelationships
}
func flattenMapAssetDeviceRegistrationRelationship(p models.AssetDeviceRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var assetdeviceregistrationrelationships []map[string]interface{}
	var ret models.AssetDeviceRegistrationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	assetdeviceregistrationrelationship := make(map[string]interface{})
	assetdeviceregistrationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetdeviceregistrationrelationship["class_id"] = item.ClassId
	assetdeviceregistrationrelationship["moid"] = item.Moid
	assetdeviceregistrationrelationship["object_type"] = item.ObjectType
	assetdeviceregistrationrelationship["selector"] = item.Selector

	assetdeviceregistrationrelationships = append(assetdeviceregistrationrelationships, assetdeviceregistrationrelationship)
	return assetdeviceregistrationrelationships
}
func flattenMapAssetGlobalUltimate(p models.AssetGlobalUltimate, d *schema.ResourceData) []map[string]interface{} {
	var assetglobalultimates []map[string]interface{}
	var ret models.AssetGlobalUltimate
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetglobalultimate := make(map[string]interface{})
	assetglobalultimate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetglobalultimate["class_id"] = item.ClassId
	assetglobalultimate["id"] = item.Id
	assetglobalultimate["name"] = item.Name
	assetglobalultimate["object_type"] = item.ObjectType

	assetglobalultimates = append(assetglobalultimates, assetglobalultimate)
	return assetglobalultimates
}
func flattenMapAssetManagedDeviceStatus(p models.AssetManagedDeviceStatus, d *schema.ResourceData) []map[string]interface{} {
	var assetmanageddevicestatuss []map[string]interface{}
	var ret models.AssetManagedDeviceStatus
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetmanageddevicestatus := make(map[string]interface{})
	assetmanageddevicestatus["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetmanageddevicestatus["class_id"] = item.ClassId
	assetmanageddevicestatus["cloud_port"] = item.CloudPort
	assetmanageddevicestatus["connection_failure_reason"] = item.ConnectionFailureReason
	assetmanageddevicestatus["connection_status"] = item.ConnectionStatus
	assetmanageddevicestatus["error_code"] = item.ErrorCode
	assetmanageddevicestatus["error_reason"] = item.ErrorReason
	assetmanageddevicestatus["object_type"] = item.ObjectType
	assetmanageddevicestatus["process_id"] = item.ProcessId
	assetmanageddevicestatus["server_port"] = item.ServerPort
	assetmanageddevicestatus["state"] = item.State

	assetmanageddevicestatuss = append(assetmanageddevicestatuss, assetmanageddevicestatus)
	return assetmanageddevicestatuss
}
func flattenMapAssetParentConnectionSignature(p models.AssetParentConnectionSignature, d *schema.ResourceData) []map[string]interface{} {
	var assetparentconnectionsignatures []map[string]interface{}
	var ret models.AssetParentConnectionSignature
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetparentconnectionsignature := make(map[string]interface{})
	assetparentconnectionsignature["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetparentconnectionsignature["class_id"] = item.ClassId
	assetparentconnectionsignature["device_id"] = item.DeviceId
	assetparentconnectionsignature["node_id"] = item.NodeId
	assetparentconnectionsignature["object_type"] = item.ObjectType

	assetparentconnectionsignatures = append(assetparentconnectionsignatures, assetparentconnectionsignature)
	return assetparentconnectionsignatures
}
func flattenMapAssetProductInformation(p models.AssetProductInformation, d *schema.ResourceData) []map[string]interface{} {
	var assetproductinformations []map[string]interface{}
	var ret models.AssetProductInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetproductinformation := make(map[string]interface{})
	assetproductinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetproductinformation["bill_to"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetBillTo(), d)
	assetproductinformation["class_id"] = item.ClassId
	assetproductinformation["description"] = item.Description
	assetproductinformation["family"] = item.Family
	assetproductinformation["group"] = item.Group
	assetproductinformation["number"] = item.Number
	assetproductinformation["object_type"] = item.ObjectType
	assetproductinformation["ship_to"] = (func(p models.AssetAddressInformation, d *schema.ResourceData) []map[string]interface{} {
		var assetaddressinformations []map[string]interface{}
		var ret models.AssetAddressInformation
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		assetaddressinformation := make(map[string]interface{})
		assetaddressinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		assetaddressinformation["address1"] = item.Address1
		assetaddressinformation["address2"] = item.Address2
		assetaddressinformation["city"] = item.City
		assetaddressinformation["class_id"] = item.ClassId
		assetaddressinformation["country"] = item.Country
		assetaddressinformation["location"] = item.Location
		assetaddressinformation["name"] = item.Name
		assetaddressinformation["object_type"] = item.ObjectType
		assetaddressinformation["postal_code"] = item.PostalCode
		assetaddressinformation["state"] = item.State

		assetaddressinformations = append(assetaddressinformations, assetaddressinformation)
		return assetaddressinformations
	})(item.GetShipTo(), d)
	assetproductinformation["sub_type"] = item.SubType

	assetproductinformations = append(assetproductinformations, assetproductinformation)
	return assetproductinformations
}
func flattenMapAssetSudiInfo(p models.AssetSudiInfo, d *schema.ResourceData) []map[string]interface{} {
	var assetsudiinfos []map[string]interface{}
	var ret models.AssetSudiInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	assetsudiinfo := make(map[string]interface{})
	assetsudiinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	assetsudiinfo["class_id"] = item.ClassId
	assetsudiinfo["object_type"] = item.ObjectType
	assetsudiinfo["pid"] = item.Pid
	assetsudiinfo["serial_number"] = item.SerialNumber
	assetsudiinfo["signature"] = item.Signature
	assetsudiinfo["status"] = item.Status
	assetsudiinfo["sudi_certificate"] = (func(p models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
		var x509certificates []map[string]interface{}
		var ret models.X509Certificate
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		x509certificate := make(map[string]interface{})
		x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		x509certificate["class_id"] = item.ClassId
		x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetIssuer(), d)
		x509certificate["object_type"] = item.ObjectType
		x509certificate["pem_certificate"] = item.PemCertificate
		x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
		x509certificate["signature_algorithm"] = item.SignatureAlgorithm
		x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
			var pkixdistinguishednames []map[string]interface{}
			var ret models.PkixDistinguishedName
			if reflect.DeepEqual(ret, p) {
				return nil
			}
			item := p
			pkixdistinguishedname := make(map[string]interface{})
			pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			pkixdistinguishedname["class_id"] = item.ClassId
			pkixdistinguishedname["common_name"] = item.CommonName
			pkixdistinguishedname["country"] = item.Country
			pkixdistinguishedname["locality"] = item.Locality
			pkixdistinguishedname["object_type"] = item.ObjectType
			pkixdistinguishedname["organization"] = item.Organization
			pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
			pkixdistinguishedname["state"] = item.State

			pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
			return pkixdistinguishednames
		})(item.GetSubject(), d)

		x509certificates = append(x509certificates, x509certificate)
		return x509certificates
	})(item.GetSudiCertificate(), d)

	assetsudiinfos = append(assetsudiinfos, assetsudiinfo)
	return assetsudiinfos
}
func flattenMapBiosBootModeRelationship(p models.BiosBootModeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosbootmoderelationships []map[string]interface{}
	var ret models.BiosBootModeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biosbootmoderelationship := make(map[string]interface{})
	biosbootmoderelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biosbootmoderelationship["class_id"] = item.ClassId
	biosbootmoderelationship["moid"] = item.Moid
	biosbootmoderelationship["object_type"] = item.ObjectType
	biosbootmoderelationship["selector"] = item.Selector

	biosbootmoderelationships = append(biosbootmoderelationships, biosbootmoderelationship)
	return biosbootmoderelationships
}
func flattenMapBiosSystemBootOrderRelationship(p models.BiosSystemBootOrderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biossystembootorderrelationships []map[string]interface{}
	var ret models.BiosSystemBootOrderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biossystembootorderrelationship := make(map[string]interface{})
	biossystembootorderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biossystembootorderrelationship["class_id"] = item.ClassId
	biossystembootorderrelationship["moid"] = item.Moid
	biossystembootorderrelationship["object_type"] = item.ObjectType
	biossystembootorderrelationship["selector"] = item.Selector

	biossystembootorderrelationships = append(biossystembootorderrelationships, biossystembootorderrelationship)
	return biossystembootorderrelationships
}
func flattenMapBiosUnitRelationship(p models.BiosUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var biosunitrelationships []map[string]interface{}
	var ret models.BiosUnitRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	biosunitrelationship := make(map[string]interface{})
	biosunitrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	biosunitrelationship["class_id"] = item.ClassId
	biosunitrelationship["moid"] = item.Moid
	biosunitrelationship["object_type"] = item.ObjectType
	biosunitrelationship["selector"] = item.Selector

	biosunitrelationships = append(biosunitrelationships, biosunitrelationship)
	return biosunitrelationships
}
func flattenMapBootDeviceBootModeRelationship(p models.BootDeviceBootModeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var bootdevicebootmoderelationships []map[string]interface{}
	var ret models.BootDeviceBootModeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	bootdevicebootmoderelationship := make(map[string]interface{})
	bootdevicebootmoderelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	bootdevicebootmoderelationship["class_id"] = item.ClassId
	bootdevicebootmoderelationship["moid"] = item.Moid
	bootdevicebootmoderelationship["object_type"] = item.ObjectType
	bootdevicebootmoderelationship["selector"] = item.Selector

	bootdevicebootmoderelationships = append(bootdevicebootmoderelationships, bootdevicebootmoderelationship)
	return bootdevicebootmoderelationships
}
func flattenMapCapabilitySectionRelationship(p models.CapabilitySectionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var capabilitysectionrelationships []map[string]interface{}
	var ret models.CapabilitySectionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	capabilitysectionrelationship := make(map[string]interface{})
	capabilitysectionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	capabilitysectionrelationship["class_id"] = item.ClassId
	capabilitysectionrelationship["moid"] = item.Moid
	capabilitysectionrelationship["object_type"] = item.ObjectType
	capabilitysectionrelationship["selector"] = item.Selector

	capabilitysectionrelationships = append(capabilitysectionrelationships, capabilitysectionrelationship)
	return capabilitysectionrelationships
}
func flattenMapCommCredential(p models.CommCredential, d *schema.ResourceData) []map[string]interface{} {
	var commcredentials []map[string]interface{}
	var ret models.CommCredential
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	commcredential := make(map[string]interface{})
	commcredential["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	commcredential["class_id"] = item.ClassId
	commcredential["is_password_set"] = item.IsPasswordSet
	commcredential["object_type"] = item.ObjectType
	password_x := d.Get("credential").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	commcredential["password"] = password_y["password"]
	commcredential["username"] = item.Username

	commcredentials = append(commcredentials, commcredential)
	return commcredentials
}
func flattenMapCommIpV4Interface(p models.CommIpV4Interface, d *schema.ResourceData) []map[string]interface{} {
	var commipv4interfaces []map[string]interface{}
	var ret models.CommIpV4Interface
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	commipv4interface := make(map[string]interface{})
	commipv4interface["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	commipv4interface["class_id"] = item.ClassId
	commipv4interface["gateway"] = item.Gateway
	commipv4interface["ip_address"] = item.IpAddress
	commipv4interface["netmask"] = item.Netmask
	commipv4interface["object_type"] = item.ObjectType

	commipv4interfaces = append(commipv4interfaces, commipv4interface)
	return commipv4interfaces
}
func flattenMapComputeBladeRelationship(p models.ComputeBladeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computebladerelationships []map[string]interface{}
	var ret models.ComputeBladeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computebladerelationship := make(map[string]interface{})
	computebladerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computebladerelationship["class_id"] = item.ClassId
	computebladerelationship["moid"] = item.Moid
	computebladerelationship["object_type"] = item.ObjectType
	computebladerelationship["selector"] = item.Selector

	computebladerelationships = append(computebladerelationships, computebladerelationship)
	return computebladerelationships
}
func flattenMapComputeBoardRelationship(p models.ComputeBoardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computeboardrelationships []map[string]interface{}
	var ret models.ComputeBoardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computeboardrelationship := make(map[string]interface{})
	computeboardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computeboardrelationship["class_id"] = item.ClassId
	computeboardrelationship["moid"] = item.Moid
	computeboardrelationship["object_type"] = item.ObjectType
	computeboardrelationship["selector"] = item.Selector

	computeboardrelationships = append(computeboardrelationships, computeboardrelationship)
	return computeboardrelationships
}
func flattenMapComputePersistentMemoryOperation(p models.ComputePersistentMemoryOperation, d *schema.ResourceData) []map[string]interface{} {
	var computepersistentmemoryoperations []map[string]interface{}
	var ret models.ComputePersistentMemoryOperation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computepersistentmemoryoperation := make(map[string]interface{})
	computepersistentmemoryoperation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computepersistentmemoryoperation["admin_action"] = item.AdminAction
	computepersistentmemoryoperation["class_id"] = item.ClassId
	computepersistentmemoryoperation["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	computepersistentmemoryoperation["modules"] = (func(p []models.ComputePersistentMemoryModule, d *schema.ResourceData) []map[string]interface{} {
		var computepersistentmemorymodules []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			computepersistentmemorymodule := make(map[string]interface{})
			computepersistentmemorymodule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			computepersistentmemorymodule["class_id"] = item.ClassId
			computepersistentmemorymodule["object_type"] = item.ObjectType
			computepersistentmemorymodule["socket_id"] = item.SocketId
			computepersistentmemorymodule["socket_memory_id"] = item.SocketMemoryId
			computepersistentmemorymodules = append(computepersistentmemorymodules, computepersistentmemorymodule)
		}
		return computepersistentmemorymodules
	})(item.GetModules(), d)
	computepersistentmemoryoperation["object_type"] = item.ObjectType
	secure_passphrase_x := d.Get("persistent_memory_operation").([]interface{})
	secure_passphrase_y := secure_passphrase_x[0].(map[string]interface{})
	computepersistentmemoryoperation["secure_passphrase"] = secure_passphrase_y["secure_passphrase"]

	computepersistentmemoryoperations = append(computepersistentmemoryoperations, computepersistentmemoryoperation)
	return computepersistentmemoryoperations
}
func flattenMapComputePhysicalRelationship(p models.ComputePhysicalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computephysicalrelationships []map[string]interface{}
	var ret models.ComputePhysicalRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computephysicalrelationship := make(map[string]interface{})
	computephysicalrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computephysicalrelationship["class_id"] = item.ClassId
	computephysicalrelationship["moid"] = item.Moid
	computephysicalrelationship["object_type"] = item.ObjectType
	computephysicalrelationship["selector"] = item.Selector

	computephysicalrelationships = append(computephysicalrelationships, computephysicalrelationship)
	return computephysicalrelationships
}
func flattenMapComputeRackUnitRelationship(p models.ComputeRackUnitRelationship, d *schema.ResourceData) []map[string]interface{} {
	var computerackunitrelationships []map[string]interface{}
	var ret models.ComputeRackUnitRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	computerackunitrelationship := make(map[string]interface{})
	computerackunitrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computerackunitrelationship["class_id"] = item.ClassId
	computerackunitrelationship["moid"] = item.Moid
	computerackunitrelationship["object_type"] = item.ObjectType
	computerackunitrelationship["selector"] = item.Selector

	computerackunitrelationships = append(computerackunitrelationships, computerackunitrelationship)
	return computerackunitrelationships
}
func flattenMapComputeServerConfig(p models.ComputeServerConfig, d *schema.ResourceData) []map[string]interface{} {
	var computeserverconfigs []map[string]interface{}
	var ret models.ComputeServerConfig
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	computeserverconfig := make(map[string]interface{})
	computeserverconfig["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	computeserverconfig["asset_tag"] = item.AssetTag
	computeserverconfig["class_id"] = item.ClassId
	computeserverconfig["object_type"] = item.ObjectType
	computeserverconfig["user_label"] = item.UserLabel

	computeserverconfigs = append(computeserverconfigs, computeserverconfig)
	return computeserverconfigs
}
func flattenMapCondHclStatusRelationship(p models.CondHclStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var condhclstatusrelationships []map[string]interface{}
	var ret models.CondHclStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	condhclstatusrelationship := make(map[string]interface{})
	condhclstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	condhclstatusrelationship["class_id"] = item.ClassId
	condhclstatusrelationship["moid"] = item.Moid
	condhclstatusrelationship["object_type"] = item.ObjectType
	condhclstatusrelationship["selector"] = item.Selector

	condhclstatusrelationships = append(condhclstatusrelationships, condhclstatusrelationship)
	return condhclstatusrelationships
}
func flattenMapConfigExportedItemRelationship(p models.ConfigExportedItemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configexporteditemrelationships []map[string]interface{}
	var ret models.ConfigExportedItemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	configexporteditemrelationship := make(map[string]interface{})
	configexporteditemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configexporteditemrelationship["class_id"] = item.ClassId
	configexporteditemrelationship["moid"] = item.Moid
	configexporteditemrelationship["object_type"] = item.ObjectType
	configexporteditemrelationship["selector"] = item.Selector

	configexporteditemrelationships = append(configexporteditemrelationships, configexporteditemrelationship)
	return configexporteditemrelationships
}
func flattenMapConfigExporterRelationship(p models.ConfigExporterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configexporterrelationships []map[string]interface{}
	var ret models.ConfigExporterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	configexporterrelationship := make(map[string]interface{})
	configexporterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configexporterrelationship["class_id"] = item.ClassId
	configexporterrelationship["moid"] = item.Moid
	configexporterrelationship["object_type"] = item.ObjectType
	configexporterrelationship["selector"] = item.Selector

	configexporterrelationships = append(configexporterrelationships, configexporterrelationship)
	return configexporterrelationships
}
func flattenMapConfigImporterRelationship(p models.ConfigImporterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var configimporterrelationships []map[string]interface{}
	var ret models.ConfigImporterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	configimporterrelationship := make(map[string]interface{})
	configimporterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configimporterrelationship["class_id"] = item.ClassId
	configimporterrelationship["moid"] = item.Moid
	configimporterrelationship["object_type"] = item.ObjectType
	configimporterrelationship["selector"] = item.Selector

	configimporterrelationships = append(configimporterrelationships, configimporterrelationship)
	return configimporterrelationships
}
func flattenMapConfigMoRef(p models.ConfigMoRef, d *schema.ResourceData) []map[string]interface{} {
	var configmorefs []map[string]interface{}
	var ret models.ConfigMoRef
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	configmoref := make(map[string]interface{})
	configmoref["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	configmoref["class_id"] = item.ClassId
	configmoref["moid"] = item.Moid
	configmoref["object_type"] = item.ObjectType

	configmorefs = append(configmorefs, configmoref)
	return configmorefs
}
func flattenMapConnectorFileChecksum(p models.ConnectorFileChecksum, d *schema.ResourceData) []map[string]interface{} {
	var connectorfilechecksums []map[string]interface{}
	var ret models.ConnectorFileChecksum
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	connectorfilechecksum := make(map[string]interface{})
	connectorfilechecksum["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	connectorfilechecksum["class_id"] = item.ClassId
	connectorfilechecksum["hash_algorithm"] = item.HashAlgorithm
	connectorfilechecksum["object_type"] = item.ObjectType

	connectorfilechecksums = append(connectorfilechecksums, connectorfilechecksum)
	return connectorfilechecksums
}
func flattenMapConnectorPlatformParamBase(p models.ConnectorPlatformParamBase, d *schema.ResourceData) []map[string]interface{} {
	var connectorplatformparambases []map[string]interface{}
	var ret models.ConnectorPlatformParamBase
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	connectorplatformparambase := make(map[string]interface{})
	connectorplatformparambase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	connectorplatformparambase["class_id"] = item.ClassId
	connectorplatformparambase["object_type"] = item.ObjectType

	connectorplatformparambases = append(connectorplatformparambases, connectorplatformparambase)
	return connectorplatformparambases
}
func flattenMapEquipmentChassisRelationship(p models.EquipmentChassisRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentchassisrelationships []map[string]interface{}
	var ret models.EquipmentChassisRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentchassisrelationship := make(map[string]interface{})
	equipmentchassisrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentchassisrelationship["class_id"] = item.ClassId
	equipmentchassisrelationship["moid"] = item.Moid
	equipmentchassisrelationship["object_type"] = item.ObjectType
	equipmentchassisrelationship["selector"] = item.Selector

	equipmentchassisrelationships = append(equipmentchassisrelationships, equipmentchassisrelationship)
	return equipmentchassisrelationships
}
func flattenMapEquipmentFanModuleRelationship(p models.EquipmentFanModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfanmodulerelationships []map[string]interface{}
	var ret models.EquipmentFanModuleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentfanmodulerelationship := make(map[string]interface{})
	equipmentfanmodulerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentfanmodulerelationship["class_id"] = item.ClassId
	equipmentfanmodulerelationship["moid"] = item.Moid
	equipmentfanmodulerelationship["object_type"] = item.ObjectType
	equipmentfanmodulerelationship["selector"] = item.Selector

	equipmentfanmodulerelationships = append(equipmentfanmodulerelationships, equipmentfanmodulerelationship)
	return equipmentfanmodulerelationships
}
func flattenMapEquipmentFexRelationship(p models.EquipmentFexRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentfexrelationships []map[string]interface{}
	var ret models.EquipmentFexRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentfexrelationship := make(map[string]interface{})
	equipmentfexrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentfexrelationship["class_id"] = item.ClassId
	equipmentfexrelationship["moid"] = item.Moid
	equipmentfexrelationship["object_type"] = item.ObjectType
	equipmentfexrelationship["selector"] = item.Selector

	equipmentfexrelationships = append(equipmentfexrelationships, equipmentfexrelationship)
	return equipmentfexrelationships
}
func flattenMapEquipmentIoCardRelationship(p models.EquipmentIoCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardrelationships []map[string]interface{}
	var ret models.EquipmentIoCardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentiocardrelationship := make(map[string]interface{})
	equipmentiocardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentiocardrelationship["class_id"] = item.ClassId
	equipmentiocardrelationship["moid"] = item.Moid
	equipmentiocardrelationship["object_type"] = item.ObjectType
	equipmentiocardrelationship["selector"] = item.Selector

	equipmentiocardrelationships = append(equipmentiocardrelationships, equipmentiocardrelationship)
	return equipmentiocardrelationships
}
func flattenMapEquipmentIoCardBaseRelationship(p models.EquipmentIoCardBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentiocardbaserelationships []map[string]interface{}
	var ret models.EquipmentIoCardBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentiocardbaserelationship := make(map[string]interface{})
	equipmentiocardbaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentiocardbaserelationship["class_id"] = item.ClassId
	equipmentiocardbaserelationship["moid"] = item.Moid
	equipmentiocardbaserelationship["object_type"] = item.ObjectType
	equipmentiocardbaserelationship["selector"] = item.Selector

	equipmentiocardbaserelationships = append(equipmentiocardbaserelationships, equipmentiocardbaserelationship)
	return equipmentiocardbaserelationships
}
func flattenMapEquipmentLocatorLedRelationship(p models.EquipmentLocatorLedRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentlocatorledrelationships []map[string]interface{}
	var ret models.EquipmentLocatorLedRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentlocatorledrelationship := make(map[string]interface{})
	equipmentlocatorledrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentlocatorledrelationship["class_id"] = item.ClassId
	equipmentlocatorledrelationship["moid"] = item.Moid
	equipmentlocatorledrelationship["object_type"] = item.ObjectType
	equipmentlocatorledrelationship["selector"] = item.Selector

	equipmentlocatorledrelationships = append(equipmentlocatorledrelationships, equipmentlocatorledrelationship)
	return equipmentlocatorledrelationships
}
func flattenMapEquipmentPhysicalIdentityRelationship(p models.EquipmentPhysicalIdentityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentphysicalidentityrelationships []map[string]interface{}
	var ret models.EquipmentPhysicalIdentityRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentphysicalidentityrelationship := make(map[string]interface{})
	equipmentphysicalidentityrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentphysicalidentityrelationship["class_id"] = item.ClassId
	equipmentphysicalidentityrelationship["moid"] = item.Moid
	equipmentphysicalidentityrelationship["object_type"] = item.ObjectType
	equipmentphysicalidentityrelationship["selector"] = item.Selector

	equipmentphysicalidentityrelationships = append(equipmentphysicalidentityrelationships, equipmentphysicalidentityrelationship)
	return equipmentphysicalidentityrelationships
}
func flattenMapEquipmentPsuControlRelationship(p models.EquipmentPsuControlRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentpsucontrolrelationships []map[string]interface{}
	var ret models.EquipmentPsuControlRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentpsucontrolrelationship := make(map[string]interface{})
	equipmentpsucontrolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentpsucontrolrelationship["class_id"] = item.ClassId
	equipmentpsucontrolrelationship["moid"] = item.Moid
	equipmentpsucontrolrelationship["object_type"] = item.ObjectType
	equipmentpsucontrolrelationship["selector"] = item.Selector

	equipmentpsucontrolrelationships = append(equipmentpsucontrolrelationships, equipmentpsucontrolrelationship)
	return equipmentpsucontrolrelationships
}
func flattenMapEquipmentRackEnclosureRelationship(p models.EquipmentRackEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosurerelationships []map[string]interface{}
	var ret models.EquipmentRackEnclosureRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentrackenclosurerelationship := make(map[string]interface{})
	equipmentrackenclosurerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentrackenclosurerelationship["class_id"] = item.ClassId
	equipmentrackenclosurerelationship["moid"] = item.Moid
	equipmentrackenclosurerelationship["object_type"] = item.ObjectType
	equipmentrackenclosurerelationship["selector"] = item.Selector

	equipmentrackenclosurerelationships = append(equipmentrackenclosurerelationships, equipmentrackenclosurerelationship)
	return equipmentrackenclosurerelationships
}
func flattenMapEquipmentRackEnclosureSlotRelationship(p models.EquipmentRackEnclosureSlotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentrackenclosureslotrelationships []map[string]interface{}
	var ret models.EquipmentRackEnclosureSlotRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentrackenclosureslotrelationship := make(map[string]interface{})
	equipmentrackenclosureslotrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentrackenclosureslotrelationship["class_id"] = item.ClassId
	equipmentrackenclosureslotrelationship["moid"] = item.Moid
	equipmentrackenclosureslotrelationship["object_type"] = item.ObjectType
	equipmentrackenclosureslotrelationship["selector"] = item.Selector

	equipmentrackenclosureslotrelationships = append(equipmentrackenclosureslotrelationships, equipmentrackenclosureslotrelationship)
	return equipmentrackenclosureslotrelationships
}
func flattenMapEquipmentSharedIoModuleRelationship(p models.EquipmentSharedIoModuleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsharediomodulerelationships []map[string]interface{}
	var ret models.EquipmentSharedIoModuleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentsharediomodulerelationship := make(map[string]interface{})
	equipmentsharediomodulerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentsharediomodulerelationship["class_id"] = item.ClassId
	equipmentsharediomodulerelationship["moid"] = item.Moid
	equipmentsharediomodulerelationship["object_type"] = item.ObjectType
	equipmentsharediomodulerelationship["selector"] = item.Selector

	equipmentsharediomodulerelationships = append(equipmentsharediomodulerelationships, equipmentsharediomodulerelationship)
	return equipmentsharediomodulerelationships
}
func flattenMapEquipmentSwitchCardRelationship(p models.EquipmentSwitchCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentswitchcardrelationships []map[string]interface{}
	var ret models.EquipmentSwitchCardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentswitchcardrelationship := make(map[string]interface{})
	equipmentswitchcardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentswitchcardrelationship["class_id"] = item.ClassId
	equipmentswitchcardrelationship["moid"] = item.Moid
	equipmentswitchcardrelationship["object_type"] = item.ObjectType
	equipmentswitchcardrelationship["selector"] = item.Selector

	equipmentswitchcardrelationships = append(equipmentswitchcardrelationships, equipmentswitchcardrelationship)
	return equipmentswitchcardrelationships
}
func flattenMapEquipmentSystemIoControllerRelationship(p models.EquipmentSystemIoControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var equipmentsystemiocontrollerrelationships []map[string]interface{}
	var ret models.EquipmentSystemIoControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	equipmentsystemiocontrollerrelationship := make(map[string]interface{})
	equipmentsystemiocontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	equipmentsystemiocontrollerrelationship["class_id"] = item.ClassId
	equipmentsystemiocontrollerrelationship["moid"] = item.Moid
	equipmentsystemiocontrollerrelationship["object_type"] = item.ObjectType
	equipmentsystemiocontrollerrelationship["selector"] = item.Selector

	equipmentsystemiocontrollerrelationships = append(equipmentsystemiocontrollerrelationships, equipmentsystemiocontrollerrelationship)
	return equipmentsystemiocontrollerrelationships
}
func flattenMapEtherPhysicalPortRelationship(p models.EtherPhysicalPortRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportrelationships []map[string]interface{}
	var ret models.EtherPhysicalPortRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	etherphysicalportrelationship := make(map[string]interface{})
	etherphysicalportrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	etherphysicalportrelationship["class_id"] = item.ClassId
	etherphysicalportrelationship["moid"] = item.Moid
	etherphysicalportrelationship["object_type"] = item.ObjectType
	etherphysicalportrelationship["selector"] = item.Selector

	etherphysicalportrelationships = append(etherphysicalportrelationships, etherphysicalportrelationship)
	return etherphysicalportrelationships
}
func flattenMapEtherPhysicalPortBaseRelationship(p models.EtherPhysicalPortBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var etherphysicalportbaserelationships []map[string]interface{}
	var ret models.EtherPhysicalPortBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	etherphysicalportbaserelationship := make(map[string]interface{})
	etherphysicalportbaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	etherphysicalportbaserelationship["class_id"] = item.ClassId
	etherphysicalportbaserelationship["moid"] = item.Moid
	etherphysicalportbaserelationship["object_type"] = item.ObjectType
	etherphysicalportbaserelationship["selector"] = item.Selector

	etherphysicalportbaserelationships = append(etherphysicalportbaserelationships, etherphysicalportbaserelationship)
	return etherphysicalportbaserelationships
}
func flattenMapFabricConfigResultRelationship(p models.FabricConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricconfigresultrelationships []map[string]interface{}
	var ret models.FabricConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricconfigresultrelationship := make(map[string]interface{})
	fabricconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricconfigresultrelationship["class_id"] = item.ClassId
	fabricconfigresultrelationship["moid"] = item.Moid
	fabricconfigresultrelationship["object_type"] = item.ObjectType
	fabricconfigresultrelationship["selector"] = item.Selector

	fabricconfigresultrelationships = append(fabricconfigresultrelationships, fabricconfigresultrelationship)
	return fabricconfigresultrelationships
}
func flattenMapFabricEthNetworkPolicyRelationship(p models.FabricEthNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricethnetworkpolicyrelationships []map[string]interface{}
	var ret models.FabricEthNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricethnetworkpolicyrelationship := make(map[string]interface{})
	fabricethnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricethnetworkpolicyrelationship["class_id"] = item.ClassId
	fabricethnetworkpolicyrelationship["moid"] = item.Moid
	fabricethnetworkpolicyrelationship["object_type"] = item.ObjectType
	fabricethnetworkpolicyrelationship["selector"] = item.Selector

	fabricethnetworkpolicyrelationships = append(fabricethnetworkpolicyrelationships, fabricethnetworkpolicyrelationship)
	return fabricethnetworkpolicyrelationships
}
func flattenMapFabricFcNetworkPolicyRelationship(p models.FabricFcNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricfcnetworkpolicyrelationships []map[string]interface{}
	var ret models.FabricFcNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricfcnetworkpolicyrelationship := make(map[string]interface{})
	fabricfcnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricfcnetworkpolicyrelationship["class_id"] = item.ClassId
	fabricfcnetworkpolicyrelationship["moid"] = item.Moid
	fabricfcnetworkpolicyrelationship["object_type"] = item.ObjectType
	fabricfcnetworkpolicyrelationship["selector"] = item.Selector

	fabricfcnetworkpolicyrelationships = append(fabricfcnetworkpolicyrelationships, fabricfcnetworkpolicyrelationship)
	return fabricfcnetworkpolicyrelationships
}
func flattenMapFabricLldpSettings(p models.FabricLldpSettings, d *schema.ResourceData) []map[string]interface{} {
	var fabriclldpsettingss []map[string]interface{}
	var ret models.FabricLldpSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fabriclldpsettings := make(map[string]interface{})
	fabriclldpsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabriclldpsettings["class_id"] = item.ClassId
	fabriclldpsettings["object_type"] = item.ObjectType
	fabriclldpsettings["receive_enabled"] = item.ReceiveEnabled
	fabriclldpsettings["transmit_enabled"] = item.TransmitEnabled

	fabriclldpsettingss = append(fabriclldpsettingss, fabriclldpsettings)
	return fabriclldpsettingss
}
func flattenMapFabricPortPolicyRelationship(p models.FabricPortPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricportpolicyrelationships []map[string]interface{}
	var ret models.FabricPortPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricportpolicyrelationship := make(map[string]interface{})
	fabricportpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricportpolicyrelationship["class_id"] = item.ClassId
	fabricportpolicyrelationship["moid"] = item.Moid
	fabricportpolicyrelationship["object_type"] = item.ObjectType
	fabricportpolicyrelationship["selector"] = item.Selector

	fabricportpolicyrelationships = append(fabricportpolicyrelationships, fabricportpolicyrelationship)
	return fabricportpolicyrelationships
}
func flattenMapFabricSwitchClusterProfileRelationship(p models.FabricSwitchClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricswitchclusterprofilerelationships []map[string]interface{}
	var ret models.FabricSwitchClusterProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricswitchclusterprofilerelationship := make(map[string]interface{})
	fabricswitchclusterprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricswitchclusterprofilerelationship["class_id"] = item.ClassId
	fabricswitchclusterprofilerelationship["moid"] = item.Moid
	fabricswitchclusterprofilerelationship["object_type"] = item.ObjectType
	fabricswitchclusterprofilerelationship["selector"] = item.Selector

	fabricswitchclusterprofilerelationships = append(fabricswitchclusterprofilerelationships, fabricswitchclusterprofilerelationship)
	return fabricswitchclusterprofilerelationships
}
func flattenMapFabricSwitchProfileRelationship(p models.FabricSwitchProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fabricswitchprofilerelationships []map[string]interface{}
	var ret models.FabricSwitchProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fabricswitchprofilerelationship := make(map[string]interface{})
	fabricswitchprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fabricswitchprofilerelationship["class_id"] = item.ClassId
	fabricswitchprofilerelationship["moid"] = item.Moid
	fabricswitchprofilerelationship["object_type"] = item.ObjectType
	fabricswitchprofilerelationship["selector"] = item.Selector

	fabricswitchprofilerelationships = append(fabricswitchprofilerelationships, fabricswitchprofilerelationship)
	return fabricswitchprofilerelationships
}
func flattenMapFcpoolBlock(p models.FcpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolblocks []map[string]interface{}
	var ret models.FcpoolBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	fcpoolblock := make(map[string]interface{})
	fcpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolblock["class_id"] = item.ClassId
	fcpoolblock["from"] = item.From
	fcpoolblock["object_type"] = item.ObjectType
	fcpoolblock["size"] = item.Size
	fcpoolblock["to"] = item.To

	fcpoolblocks = append(fcpoolblocks, fcpoolblock)
	return fcpoolblocks
}
func flattenMapFcpoolFcBlockRelationship(p models.FcpoolFcBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolfcblockrelationships []map[string]interface{}
	var ret models.FcpoolFcBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolfcblockrelationship := make(map[string]interface{})
	fcpoolfcblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolfcblockrelationship["class_id"] = item.ClassId
	fcpoolfcblockrelationship["moid"] = item.Moid
	fcpoolfcblockrelationship["object_type"] = item.ObjectType
	fcpoolfcblockrelationship["selector"] = item.Selector

	fcpoolfcblockrelationships = append(fcpoolfcblockrelationships, fcpoolfcblockrelationship)
	return fcpoolfcblockrelationships
}
func flattenMapFcpoolLeaseRelationship(p models.FcpoolLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolleaserelationships []map[string]interface{}
	var ret models.FcpoolLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolleaserelationship := make(map[string]interface{})
	fcpoolleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolleaserelationship["class_id"] = item.ClassId
	fcpoolleaserelationship["moid"] = item.Moid
	fcpoolleaserelationship["object_type"] = item.ObjectType
	fcpoolleaserelationship["selector"] = item.Selector

	fcpoolleaserelationships = append(fcpoolleaserelationships, fcpoolleaserelationship)
	return fcpoolleaserelationships
}
func flattenMapFcpoolPoolRelationship(p models.FcpoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolpoolrelationships []map[string]interface{}
	var ret models.FcpoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolpoolrelationship := make(map[string]interface{})
	fcpoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolpoolrelationship["class_id"] = item.ClassId
	fcpoolpoolrelationship["moid"] = item.Moid
	fcpoolpoolrelationship["object_type"] = item.ObjectType
	fcpoolpoolrelationship["selector"] = item.Selector

	fcpoolpoolrelationships = append(fcpoolpoolrelationships, fcpoolpoolrelationship)
	return fcpoolpoolrelationships
}
func flattenMapFcpoolPoolMemberRelationship(p models.FcpoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpoolpoolmemberrelationships []map[string]interface{}
	var ret models.FcpoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpoolpoolmemberrelationship := make(map[string]interface{})
	fcpoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpoolpoolmemberrelationship["class_id"] = item.ClassId
	fcpoolpoolmemberrelationship["moid"] = item.Moid
	fcpoolpoolmemberrelationship["object_type"] = item.ObjectType
	fcpoolpoolmemberrelationship["selector"] = item.Selector

	fcpoolpoolmemberrelationships = append(fcpoolpoolmemberrelationships, fcpoolpoolmemberrelationship)
	return fcpoolpoolmemberrelationships
}
func flattenMapFcpoolUniverseRelationship(p models.FcpoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var fcpooluniverserelationships []map[string]interface{}
	var ret models.FcpoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	fcpooluniverserelationship := make(map[string]interface{})
	fcpooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	fcpooluniverserelationship["class_id"] = item.ClassId
	fcpooluniverserelationship["moid"] = item.Moid
	fcpooluniverserelationship["object_type"] = item.ObjectType
	fcpooluniverserelationship["selector"] = item.Selector

	fcpooluniverserelationships = append(fcpooluniverserelationships, fcpooluniverserelationship)
	return fcpooluniverserelationships
}
func flattenMapFirmwareDirectDownload(p models.FirmwareDirectDownload, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredirectdownloads []map[string]interface{}
	var ret models.FirmwareDirectDownload
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	firmwaredirectdownload := make(map[string]interface{})
	firmwaredirectdownload["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwaredirectdownload["class_id"] = item.ClassId
	firmwaredirectdownload["http_server"] = (func(p models.FirmwareHttpServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		var ret models.FirmwareHttpServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarehttpserver := make(map[string]interface{})
		firmwarehttpserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarehttpserver["class_id"] = item.ClassId
		firmwarehttpserver["location_link"] = item.LocationLink
		firmwarehttpserver["mount_options"] = item.MountOptions
		firmwarehttpserver["object_type"] = item.ObjectType

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.GetHttpServer(), d)
	firmwaredirectdownload["image_source"] = item.ImageSource
	firmwaredirectdownload["is_password_set"] = item.IsPasswordSet
	firmwaredirectdownload["object_type"] = item.ObjectType
	password_x := d.Get("direct_download").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	firmwaredirectdownload["password"] = password_y["password"]
	firmwaredirectdownload["upgradeoption"] = item.Upgradeoption
	firmwaredirectdownload["username"] = item.Username

	firmwaredirectdownloads = append(firmwaredirectdownloads, firmwaredirectdownload)
	return firmwaredirectdownloads
}
func flattenMapFirmwareDistributableRelationship(p models.FirmwareDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwaredistributablerelationships []map[string]interface{}
	var ret models.FirmwareDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwaredistributablerelationship := make(map[string]interface{})
	firmwaredistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwaredistributablerelationship["class_id"] = item.ClassId
	firmwaredistributablerelationship["moid"] = item.Moid
	firmwaredistributablerelationship["object_type"] = item.ObjectType
	firmwaredistributablerelationship["selector"] = item.Selector

	firmwaredistributablerelationships = append(firmwaredistributablerelationships, firmwaredistributablerelationship)
	return firmwaredistributablerelationships
}
func flattenMapFirmwareNetworkShare(p models.FirmwareNetworkShare, d *schema.ResourceData) []map[string]interface{} {
	var firmwarenetworkshares []map[string]interface{}
	var ret models.FirmwareNetworkShare
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	firmwarenetworkshare := make(map[string]interface{})
	firmwarenetworkshare["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwarenetworkshare["cifs_server"] = (func(p models.FirmwareCifsServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarecifsservers []map[string]interface{}
		var ret models.FirmwareCifsServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarecifsserver := make(map[string]interface{})
		firmwarecifsserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarecifsserver["class_id"] = item.ClassId
		firmwarecifsserver["file_location"] = item.FileLocation
		firmwarecifsserver["mount_options"] = item.MountOptions
		firmwarecifsserver["object_type"] = item.ObjectType
		firmwarecifsserver["remote_file"] = item.RemoteFile
		firmwarecifsserver["remote_ip"] = item.RemoteIp
		firmwarecifsserver["remote_share"] = item.RemoteShare

		firmwarecifsservers = append(firmwarecifsservers, firmwarecifsserver)
		return firmwarecifsservers
	})(item.GetCifsServer(), d)
	firmwarenetworkshare["class_id"] = item.ClassId
	firmwarenetworkshare["http_server"] = (func(p models.FirmwareHttpServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarehttpservers []map[string]interface{}
		var ret models.FirmwareHttpServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarehttpserver := make(map[string]interface{})
		firmwarehttpserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarehttpserver["class_id"] = item.ClassId
		firmwarehttpserver["location_link"] = item.LocationLink
		firmwarehttpserver["mount_options"] = item.MountOptions
		firmwarehttpserver["object_type"] = item.ObjectType

		firmwarehttpservers = append(firmwarehttpservers, firmwarehttpserver)
		return firmwarehttpservers
	})(item.GetHttpServer(), d)
	firmwarenetworkshare["is_password_set"] = item.IsPasswordSet
	firmwarenetworkshare["map_type"] = item.MapType
	firmwarenetworkshare["nfs_server"] = (func(p models.FirmwareNfsServer, d *schema.ResourceData) []map[string]interface{} {
		var firmwarenfsservers []map[string]interface{}
		var ret models.FirmwareNfsServer
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		firmwarenfsserver := make(map[string]interface{})
		firmwarenfsserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		firmwarenfsserver["class_id"] = item.ClassId
		firmwarenfsserver["file_location"] = item.FileLocation
		firmwarenfsserver["mount_options"] = item.MountOptions
		firmwarenfsserver["object_type"] = item.ObjectType
		firmwarenfsserver["remote_file"] = item.RemoteFile
		firmwarenfsserver["remote_ip"] = item.RemoteIp
		firmwarenfsserver["remote_share"] = item.RemoteShare

		firmwarenfsservers = append(firmwarenfsservers, firmwarenfsserver)
		return firmwarenfsservers
	})(item.GetNfsServer(), d)
	firmwarenetworkshare["object_type"] = item.ObjectType
	password_x := d.Get("network_share").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	firmwarenetworkshare["password"] = password_y["password"]
	firmwarenetworkshare["upgradeoption"] = item.Upgradeoption
	firmwarenetworkshare["username"] = item.Username

	firmwarenetworkshares = append(firmwarenetworkshares, firmwarenetworkshare)
	return firmwarenetworkshares
}
func flattenMapFirmwareRunningFirmwareRelationship(p models.FirmwareRunningFirmwareRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwarerunningfirmwarerelationships []map[string]interface{}
	var ret models.FirmwareRunningFirmwareRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwarerunningfirmwarerelationship := make(map[string]interface{})
	firmwarerunningfirmwarerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwarerunningfirmwarerelationship["class_id"] = item.ClassId
	firmwarerunningfirmwarerelationship["moid"] = item.Moid
	firmwarerunningfirmwarerelationship["object_type"] = item.ObjectType
	firmwarerunningfirmwarerelationship["selector"] = item.Selector

	firmwarerunningfirmwarerelationships = append(firmwarerunningfirmwarerelationships, firmwarerunningfirmwarerelationship)
	return firmwarerunningfirmwarerelationships
}
func flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(p models.FirmwareServerConfigurationUtilityDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareserverconfigurationutilitydistributablerelationships []map[string]interface{}
	var ret models.FirmwareServerConfigurationUtilityDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareserverconfigurationutilitydistributablerelationship := make(map[string]interface{})
	firmwareserverconfigurationutilitydistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareserverconfigurationutilitydistributablerelationship["class_id"] = item.ClassId
	firmwareserverconfigurationutilitydistributablerelationship["moid"] = item.Moid
	firmwareserverconfigurationutilitydistributablerelationship["object_type"] = item.ObjectType
	firmwareserverconfigurationutilitydistributablerelationship["selector"] = item.Selector

	firmwareserverconfigurationutilitydistributablerelationships = append(firmwareserverconfigurationutilitydistributablerelationships, firmwareserverconfigurationutilitydistributablerelationship)
	return firmwareserverconfigurationutilitydistributablerelationships
}
func flattenMapFirmwareUpgradeBaseRelationship(p models.FirmwareUpgradeBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradebaserelationships []map[string]interface{}
	var ret models.FirmwareUpgradeBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareupgradebaserelationship := make(map[string]interface{})
	firmwareupgradebaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareupgradebaserelationship["class_id"] = item.ClassId
	firmwareupgradebaserelationship["moid"] = item.Moid
	firmwareupgradebaserelationship["object_type"] = item.ObjectType
	firmwareupgradebaserelationship["selector"] = item.Selector

	firmwareupgradebaserelationships = append(firmwareupgradebaserelationships, firmwareupgradebaserelationship)
	return firmwareupgradebaserelationships
}
func flattenMapFirmwareUpgradeImpactStatusRelationship(p models.FirmwareUpgradeImpactStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradeimpactstatusrelationships []map[string]interface{}
	var ret models.FirmwareUpgradeImpactStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareupgradeimpactstatusrelationship := make(map[string]interface{})
	firmwareupgradeimpactstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareupgradeimpactstatusrelationship["class_id"] = item.ClassId
	firmwareupgradeimpactstatusrelationship["moid"] = item.Moid
	firmwareupgradeimpactstatusrelationship["object_type"] = item.ObjectType
	firmwareupgradeimpactstatusrelationship["selector"] = item.Selector

	firmwareupgradeimpactstatusrelationships = append(firmwareupgradeimpactstatusrelationships, firmwareupgradeimpactstatusrelationship)
	return firmwareupgradeimpactstatusrelationships
}
func flattenMapFirmwareUpgradeStatusRelationship(p models.FirmwareUpgradeStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var firmwareupgradestatusrelationships []map[string]interface{}
	var ret models.FirmwareUpgradeStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	firmwareupgradestatusrelationship := make(map[string]interface{})
	firmwareupgradestatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	firmwareupgradestatusrelationship["class_id"] = item.ClassId
	firmwareupgradestatusrelationship["moid"] = item.Moid
	firmwareupgradestatusrelationship["object_type"] = item.ObjectType
	firmwareupgradestatusrelationship["selector"] = item.Selector

	firmwareupgradestatusrelationships = append(firmwareupgradestatusrelationships, firmwareupgradestatusrelationship)
	return firmwareupgradestatusrelationships
}
func flattenMapForecastCatalogRelationship(p models.ForecastCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastcatalogrelationships []map[string]interface{}
	var ret models.ForecastCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	forecastcatalogrelationship := make(map[string]interface{})
	forecastcatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastcatalogrelationship["class_id"] = item.ClassId
	forecastcatalogrelationship["moid"] = item.Moid
	forecastcatalogrelationship["object_type"] = item.ObjectType
	forecastcatalogrelationship["selector"] = item.Selector

	forecastcatalogrelationships = append(forecastcatalogrelationships, forecastcatalogrelationship)
	return forecastcatalogrelationships
}
func flattenMapForecastDefinitionRelationship(p models.ForecastDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var forecastdefinitionrelationships []map[string]interface{}
	var ret models.ForecastDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	forecastdefinitionrelationship := make(map[string]interface{})
	forecastdefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastdefinitionrelationship["class_id"] = item.ClassId
	forecastdefinitionrelationship["moid"] = item.Moid
	forecastdefinitionrelationship["object_type"] = item.ObjectType
	forecastdefinitionrelationship["selector"] = item.Selector

	forecastdefinitionrelationships = append(forecastdefinitionrelationships, forecastdefinitionrelationship)
	return forecastdefinitionrelationships
}
func flattenMapForecastModel(p models.ForecastModel, d *schema.ResourceData) []map[string]interface{} {
	var forecastmodels []map[string]interface{}
	var ret models.ForecastModel
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	forecastmodel := make(map[string]interface{})
	forecastmodel["accuracy"] = item.Accuracy
	forecastmodel["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	forecastmodel["class_id"] = item.ClassId
	forecastmodel["model_data"] = item.ModelData
	forecastmodel["model_type"] = item.ModelType
	forecastmodel["object_type"] = item.ObjectType

	forecastmodels = append(forecastmodels, forecastmodel)
	return forecastmodels
}
func flattenMapGraphicsCardRelationship(p models.GraphicsCardRelationship, d *schema.ResourceData) []map[string]interface{} {
	var graphicscardrelationships []map[string]interface{}
	var ret models.GraphicsCardRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	graphicscardrelationship := make(map[string]interface{})
	graphicscardrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	graphicscardrelationship["class_id"] = item.ClassId
	graphicscardrelationship["moid"] = item.Moid
	graphicscardrelationship["object_type"] = item.ObjectType
	graphicscardrelationship["selector"] = item.Selector

	graphicscardrelationships = append(graphicscardrelationships, graphicscardrelationship)
	return graphicscardrelationships
}
func flattenMapHclOperatingSystemRelationship(p models.HclOperatingSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemrelationships []map[string]interface{}
	var ret models.HclOperatingSystemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hcloperatingsystemrelationship := make(map[string]interface{})
	hcloperatingsystemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hcloperatingsystemrelationship["class_id"] = item.ClassId
	hcloperatingsystemrelationship["moid"] = item.Moid
	hcloperatingsystemrelationship["object_type"] = item.ObjectType
	hcloperatingsystemrelationship["selector"] = item.Selector

	hcloperatingsystemrelationships = append(hcloperatingsystemrelationships, hcloperatingsystemrelationship)
	return hcloperatingsystemrelationships
}
func flattenMapHclOperatingSystemVendorRelationship(p models.HclOperatingSystemVendorRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hcloperatingsystemvendorrelationships []map[string]interface{}
	var ret models.HclOperatingSystemVendorRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hcloperatingsystemvendorrelationship := make(map[string]interface{})
	hcloperatingsystemvendorrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hcloperatingsystemvendorrelationship["class_id"] = item.ClassId
	hcloperatingsystemvendorrelationship["moid"] = item.Moid
	hcloperatingsystemvendorrelationship["object_type"] = item.ObjectType
	hcloperatingsystemvendorrelationship["selector"] = item.Selector

	hcloperatingsystemvendorrelationships = append(hcloperatingsystemvendorrelationships, hcloperatingsystemvendorrelationship)
	return hcloperatingsystemvendorrelationships
}
func flattenMapHyperflexAppCatalogRelationship(p models.HyperflexAppCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexappcatalogrelationships []map[string]interface{}
	var ret models.HyperflexAppCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexappcatalogrelationship := make(map[string]interface{})
	hyperflexappcatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexappcatalogrelationship["class_id"] = item.ClassId
	hyperflexappcatalogrelationship["moid"] = item.Moid
	hyperflexappcatalogrelationship["object_type"] = item.ObjectType
	hyperflexappcatalogrelationship["selector"] = item.Selector

	hyperflexappcatalogrelationships = append(hyperflexappcatalogrelationships, hyperflexappcatalogrelationship)
	return hyperflexappcatalogrelationships
}
func flattenMapHyperflexAutoSupportPolicyRelationship(p models.HyperflexAutoSupportPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexautosupportpolicyrelationships []map[string]interface{}
	var ret models.HyperflexAutoSupportPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexautosupportpolicyrelationship := make(map[string]interface{})
	hyperflexautosupportpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexautosupportpolicyrelationship["class_id"] = item.ClassId
	hyperflexautosupportpolicyrelationship["moid"] = item.Moid
	hyperflexautosupportpolicyrelationship["object_type"] = item.ObjectType
	hyperflexautosupportpolicyrelationship["selector"] = item.Selector

	hyperflexautosupportpolicyrelationships = append(hyperflexautosupportpolicyrelationships, hyperflexautosupportpolicyrelationship)
	return hyperflexautosupportpolicyrelationships
}
func flattenMapHyperflexClusterRelationship(p models.HyperflexClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterrelationships []map[string]interface{}
	var ret models.HyperflexClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusterrelationship := make(map[string]interface{})
	hyperflexclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusterrelationship["class_id"] = item.ClassId
	hyperflexclusterrelationship["moid"] = item.Moid
	hyperflexclusterrelationship["object_type"] = item.ObjectType
	hyperflexclusterrelationship["selector"] = item.Selector

	hyperflexclusterrelationships = append(hyperflexclusterrelationships, hyperflexclusterrelationship)
	return hyperflexclusterrelationships
}
func flattenMapHyperflexClusterNetworkPolicyRelationship(p models.HyperflexClusterNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusternetworkpolicyrelationships []map[string]interface{}
	var ret models.HyperflexClusterNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusternetworkpolicyrelationship := make(map[string]interface{})
	hyperflexclusternetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusternetworkpolicyrelationship["class_id"] = item.ClassId
	hyperflexclusternetworkpolicyrelationship["moid"] = item.Moid
	hyperflexclusternetworkpolicyrelationship["object_type"] = item.ObjectType
	hyperflexclusternetworkpolicyrelationship["selector"] = item.Selector

	hyperflexclusternetworkpolicyrelationships = append(hyperflexclusternetworkpolicyrelationships, hyperflexclusternetworkpolicyrelationship)
	return hyperflexclusternetworkpolicyrelationships
}
func flattenMapHyperflexClusterProfileRelationship(p models.HyperflexClusterProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterprofilerelationships []map[string]interface{}
	var ret models.HyperflexClusterProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusterprofilerelationship := make(map[string]interface{})
	hyperflexclusterprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusterprofilerelationship["class_id"] = item.ClassId
	hyperflexclusterprofilerelationship["moid"] = item.Moid
	hyperflexclusterprofilerelationship["object_type"] = item.ObjectType
	hyperflexclusterprofilerelationship["selector"] = item.Selector

	hyperflexclusterprofilerelationships = append(hyperflexclusterprofilerelationships, hyperflexclusterprofilerelationship)
	return hyperflexclusterprofilerelationships
}
func flattenMapHyperflexClusterStoragePolicyRelationship(p models.HyperflexClusterStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexclusterstoragepolicyrelationships []map[string]interface{}
	var ret models.HyperflexClusterStoragePolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexclusterstoragepolicyrelationship := make(map[string]interface{})
	hyperflexclusterstoragepolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexclusterstoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexclusterstoragepolicyrelationship["moid"] = item.Moid
	hyperflexclusterstoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexclusterstoragepolicyrelationship["selector"] = item.Selector

	hyperflexclusterstoragepolicyrelationships = append(hyperflexclusterstoragepolicyrelationships, hyperflexclusterstoragepolicyrelationship)
	return hyperflexclusterstoragepolicyrelationships
}
func flattenMapHyperflexConfigResultRelationship(p models.HyperflexConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexconfigresultrelationships []map[string]interface{}
	var ret models.HyperflexConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexconfigresultrelationship := make(map[string]interface{})
	hyperflexconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexconfigresultrelationship["class_id"] = item.ClassId
	hyperflexconfigresultrelationship["moid"] = item.Moid
	hyperflexconfigresultrelationship["object_type"] = item.ObjectType
	hyperflexconfigresultrelationship["selector"] = item.Selector

	hyperflexconfigresultrelationships = append(hyperflexconfigresultrelationships, hyperflexconfigresultrelationship)
	return hyperflexconfigresultrelationships
}
func flattenMapHyperflexExtFcStoragePolicyRelationship(p models.HyperflexExtFcStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextfcstoragepolicyrelationships []map[string]interface{}
	var ret models.HyperflexExtFcStoragePolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexextfcstoragepolicyrelationship := make(map[string]interface{})
	hyperflexextfcstoragepolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexextfcstoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexextfcstoragepolicyrelationship["moid"] = item.Moid
	hyperflexextfcstoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexextfcstoragepolicyrelationship["selector"] = item.Selector

	hyperflexextfcstoragepolicyrelationships = append(hyperflexextfcstoragepolicyrelationships, hyperflexextfcstoragepolicyrelationship)
	return hyperflexextfcstoragepolicyrelationships
}
func flattenMapHyperflexExtIscsiStoragePolicyRelationship(p models.HyperflexExtIscsiStoragePolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexextiscsistoragepolicyrelationships []map[string]interface{}
	var ret models.HyperflexExtIscsiStoragePolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexextiscsistoragepolicyrelationship := make(map[string]interface{})
	hyperflexextiscsistoragepolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexextiscsistoragepolicyrelationship["class_id"] = item.ClassId
	hyperflexextiscsistoragepolicyrelationship["moid"] = item.Moid
	hyperflexextiscsistoragepolicyrelationship["object_type"] = item.ObjectType
	hyperflexextiscsistoragepolicyrelationship["selector"] = item.Selector

	hyperflexextiscsistoragepolicyrelationships = append(hyperflexextiscsistoragepolicyrelationships, hyperflexextiscsistoragepolicyrelationship)
	return hyperflexextiscsistoragepolicyrelationships
}
func flattenMapHyperflexFeatureLimitExternalRelationship(p models.HyperflexFeatureLimitExternalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitexternalrelationships []map[string]interface{}
	var ret models.HyperflexFeatureLimitExternalRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexfeaturelimitexternalrelationship := make(map[string]interface{})
	hyperflexfeaturelimitexternalrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexfeaturelimitexternalrelationship["class_id"] = item.ClassId
	hyperflexfeaturelimitexternalrelationship["moid"] = item.Moid
	hyperflexfeaturelimitexternalrelationship["object_type"] = item.ObjectType
	hyperflexfeaturelimitexternalrelationship["selector"] = item.Selector

	hyperflexfeaturelimitexternalrelationships = append(hyperflexfeaturelimitexternalrelationships, hyperflexfeaturelimitexternalrelationship)
	return hyperflexfeaturelimitexternalrelationships
}
func flattenMapHyperflexFeatureLimitInternalRelationship(p models.HyperflexFeatureLimitInternalRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexfeaturelimitinternalrelationships []map[string]interface{}
	var ret models.HyperflexFeatureLimitInternalRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexfeaturelimitinternalrelationship := make(map[string]interface{})
	hyperflexfeaturelimitinternalrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexfeaturelimitinternalrelationship["class_id"] = item.ClassId
	hyperflexfeaturelimitinternalrelationship["moid"] = item.Moid
	hyperflexfeaturelimitinternalrelationship["object_type"] = item.ObjectType
	hyperflexfeaturelimitinternalrelationship["selector"] = item.Selector

	hyperflexfeaturelimitinternalrelationships = append(hyperflexfeaturelimitinternalrelationships, hyperflexfeaturelimitinternalrelationship)
	return hyperflexfeaturelimitinternalrelationships
}
func flattenMapHyperflexHealthRelationship(p models.HyperflexHealthRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhealthrelationships []map[string]interface{}
	var ret models.HyperflexHealthRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexhealthrelationship := make(map[string]interface{})
	hyperflexhealthrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhealthrelationship["class_id"] = item.ClassId
	hyperflexhealthrelationship["moid"] = item.Moid
	hyperflexhealthrelationship["object_type"] = item.ObjectType
	hyperflexhealthrelationship["selector"] = item.Selector

	hyperflexhealthrelationships = append(hyperflexhealthrelationships, hyperflexhealthrelationship)
	return hyperflexhealthrelationships
}
func flattenMapHyperflexHxNetworkAddressDt(p models.HyperflexHxNetworkAddressDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxnetworkaddressdts []map[string]interface{}
	var ret models.HyperflexHxNetworkAddressDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxnetworkaddressdt := make(map[string]interface{})
	hyperflexhxnetworkaddressdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxnetworkaddressdt["address"] = item.Address
	hyperflexhxnetworkaddressdt["class_id"] = item.ClassId
	hyperflexhxnetworkaddressdt["fqdn"] = item.Fqdn
	hyperflexhxnetworkaddressdt["ip"] = item.Ip
	hyperflexhxnetworkaddressdt["object_type"] = item.ObjectType

	hyperflexhxnetworkaddressdts = append(hyperflexhxnetworkaddressdts, hyperflexhxnetworkaddressdt)
	return hyperflexhxnetworkaddressdts
}
func flattenMapHyperflexHxResiliencyInfoDt(p models.HyperflexHxResiliencyInfoDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxresiliencyinfodts []map[string]interface{}
	var ret models.HyperflexHxResiliencyInfoDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxresiliencyinfodt := make(map[string]interface{})
	hyperflexhxresiliencyinfodt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxresiliencyinfodt["class_id"] = item.ClassId
	hyperflexhxresiliencyinfodt["data_replication_factor"] = item.DataReplicationFactor
	hyperflexhxresiliencyinfodt["hdd_failures_tolerable"] = item.HddFailuresTolerable
	hyperflexhxresiliencyinfodt["messages"] = item.Messages
	hyperflexhxresiliencyinfodt["node_failures_tolerable"] = item.NodeFailuresTolerable
	hyperflexhxresiliencyinfodt["object_type"] = item.ObjectType
	hyperflexhxresiliencyinfodt["policy_compliance"] = item.PolicyCompliance
	hyperflexhxresiliencyinfodt["resiliency_state"] = item.ResiliencyState
	hyperflexhxresiliencyinfodt["ssd_failures_tolerable"] = item.SsdFailuresTolerable

	hyperflexhxresiliencyinfodts = append(hyperflexhxresiliencyinfodts, hyperflexhxresiliencyinfodt)
	return hyperflexhxresiliencyinfodts
}
func flattenMapHyperflexHxUuIdDt(p models.HyperflexHxUuIdDt, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexhxuuiddts []map[string]interface{}
	var ret models.HyperflexHxUuIdDt
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexhxuuiddt := make(map[string]interface{})
	hyperflexhxuuiddt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexhxuuiddt["class_id"] = item.ClassId
	hyperflexhxuuiddt["links"] = (func(p []models.HyperflexHxLinkDt, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexhxlinkdts []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			hyperflexhxlinkdt := make(map[string]interface{})
			hyperflexhxlinkdt["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			hyperflexhxlinkdt["class_id"] = item.ClassId
			hyperflexhxlinkdt["comments"] = item.Comments
			hyperflexhxlinkdt["href"] = item.Href
			hyperflexhxlinkdt["method"] = item.Method
			hyperflexhxlinkdt["object_type"] = item.ObjectType
			hyperflexhxlinkdt["rel"] = item.Rel
			hyperflexhxlinkdts = append(hyperflexhxlinkdts, hyperflexhxlinkdt)
		}
		return hyperflexhxlinkdts
	})(item.GetLinks(), d)
	hyperflexhxuuiddt["object_type"] = item.ObjectType
	hyperflexhxuuiddt["uuid"] = item.Uuid

	hyperflexhxuuiddts = append(hyperflexhxuuiddts, hyperflexhxuuiddt)
	return hyperflexhxuuiddts
}
func flattenMapHyperflexIpAddrRange(p models.HyperflexIpAddrRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexipaddrranges []map[string]interface{}
	var ret models.HyperflexIpAddrRange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexipaddrrange := make(map[string]interface{})
	hyperflexipaddrrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexipaddrrange["class_id"] = item.ClassId
	hyperflexipaddrrange["end_addr"] = item.EndAddr
	hyperflexipaddrrange["gateway"] = item.Gateway
	hyperflexipaddrrange["netmask"] = item.Netmask
	hyperflexipaddrrange["object_type"] = item.ObjectType
	hyperflexipaddrrange["start_addr"] = item.StartAddr

	hyperflexipaddrranges = append(hyperflexipaddrranges, hyperflexipaddrrange)
	return hyperflexipaddrranges
}
func flattenMapHyperflexLocalCredentialPolicyRelationship(p models.HyperflexLocalCredentialPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlocalcredentialpolicyrelationships []map[string]interface{}
	var ret models.HyperflexLocalCredentialPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexlocalcredentialpolicyrelationship := make(map[string]interface{})
	hyperflexlocalcredentialpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexlocalcredentialpolicyrelationship["class_id"] = item.ClassId
	hyperflexlocalcredentialpolicyrelationship["moid"] = item.Moid
	hyperflexlocalcredentialpolicyrelationship["object_type"] = item.ObjectType
	hyperflexlocalcredentialpolicyrelationship["selector"] = item.Selector

	hyperflexlocalcredentialpolicyrelationships = append(hyperflexlocalcredentialpolicyrelationships, hyperflexlocalcredentialpolicyrelationship)
	return hyperflexlocalcredentialpolicyrelationships
}
func flattenMapHyperflexLogicalAvailabilityZone(p models.HyperflexLogicalAvailabilityZone, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexlogicalavailabilityzones []map[string]interface{}
	var ret models.HyperflexLogicalAvailabilityZone
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexlogicalavailabilityzone := make(map[string]interface{})
	hyperflexlogicalavailabilityzone["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexlogicalavailabilityzone["auto_config"] = item.AutoConfig
	hyperflexlogicalavailabilityzone["class_id"] = item.ClassId
	hyperflexlogicalavailabilityzone["object_type"] = item.ObjectType

	hyperflexlogicalavailabilityzones = append(hyperflexlogicalavailabilityzones, hyperflexlogicalavailabilityzone)
	return hyperflexlogicalavailabilityzones
}
func flattenMapHyperflexMacAddrPrefixRange(p models.HyperflexMacAddrPrefixRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexmacaddrprefixranges []map[string]interface{}
	var ret models.HyperflexMacAddrPrefixRange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexmacaddrprefixrange := make(map[string]interface{})
	hyperflexmacaddrprefixrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexmacaddrprefixrange["class_id"] = item.ClassId
	hyperflexmacaddrprefixrange["end_addr"] = item.EndAddr
	hyperflexmacaddrprefixrange["object_type"] = item.ObjectType
	hyperflexmacaddrprefixrange["start_addr"] = item.StartAddr

	hyperflexmacaddrprefixranges = append(hyperflexmacaddrprefixranges, hyperflexmacaddrprefixrange)
	return hyperflexmacaddrprefixranges
}
func flattenMapHyperflexNamedVlan(p models.HyperflexNamedVlan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvlans []map[string]interface{}
	var ret models.HyperflexNamedVlan
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexnamedvlan := make(map[string]interface{})
	hyperflexnamedvlan["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexnamedvlan["class_id"] = item.ClassId
	hyperflexnamedvlan["name"] = item.Name
	hyperflexnamedvlan["object_type"] = item.ObjectType
	hyperflexnamedvlan["vlan_id"] = item.VlanId

	hyperflexnamedvlans = append(hyperflexnamedvlans, hyperflexnamedvlan)
	return hyperflexnamedvlans
}
func flattenMapHyperflexNamedVsan(p models.HyperflexNamedVsan, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnamedvsans []map[string]interface{}
	var ret models.HyperflexNamedVsan
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexnamedvsan := make(map[string]interface{})
	hyperflexnamedvsan["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexnamedvsan["class_id"] = item.ClassId
	hyperflexnamedvsan["name"] = item.Name
	hyperflexnamedvsan["object_type"] = item.ObjectType
	hyperflexnamedvsan["vsan_id"] = item.VsanId

	hyperflexnamedvsans = append(hyperflexnamedvsans, hyperflexnamedvsan)
	return hyperflexnamedvsans
}
func flattenMapHyperflexNodeConfigPolicyRelationship(p models.HyperflexNodeConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexnodeconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexNodeConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexnodeconfigpolicyrelationship := make(map[string]interface{})
	hyperflexnodeconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexnodeconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexnodeconfigpolicyrelationship["moid"] = item.Moid
	hyperflexnodeconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexnodeconfigpolicyrelationship["selector"] = item.Selector

	hyperflexnodeconfigpolicyrelationships = append(hyperflexnodeconfigpolicyrelationships, hyperflexnodeconfigpolicyrelationship)
	return hyperflexnodeconfigpolicyrelationships
}
func flattenMapHyperflexProxySettingPolicyRelationship(p models.HyperflexProxySettingPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexproxysettingpolicyrelationships []map[string]interface{}
	var ret models.HyperflexProxySettingPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexproxysettingpolicyrelationship := make(map[string]interface{})
	hyperflexproxysettingpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexproxysettingpolicyrelationship["class_id"] = item.ClassId
	hyperflexproxysettingpolicyrelationship["moid"] = item.Moid
	hyperflexproxysettingpolicyrelationship["object_type"] = item.ObjectType
	hyperflexproxysettingpolicyrelationship["selector"] = item.Selector

	hyperflexproxysettingpolicyrelationships = append(hyperflexproxysettingpolicyrelationships, hyperflexproxysettingpolicyrelationship)
	return hyperflexproxysettingpolicyrelationships
}
func flattenMapHyperflexServerFirmwareVersionRelationship(p models.HyperflexServerFirmwareVersionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexserverfirmwareversionrelationships []map[string]interface{}
	var ret models.HyperflexServerFirmwareVersionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexserverfirmwareversionrelationship := make(map[string]interface{})
	hyperflexserverfirmwareversionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexserverfirmwareversionrelationship["class_id"] = item.ClassId
	hyperflexserverfirmwareversionrelationship["moid"] = item.Moid
	hyperflexserverfirmwareversionrelationship["object_type"] = item.ObjectType
	hyperflexserverfirmwareversionrelationship["selector"] = item.Selector

	hyperflexserverfirmwareversionrelationships = append(hyperflexserverfirmwareversionrelationships, hyperflexserverfirmwareversionrelationship)
	return hyperflexserverfirmwareversionrelationships
}
func flattenMapHyperflexServerModelRelationship(p models.HyperflexServerModelRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexservermodelrelationships []map[string]interface{}
	var ret models.HyperflexServerModelRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexservermodelrelationship := make(map[string]interface{})
	hyperflexservermodelrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexservermodelrelationship["class_id"] = item.ClassId
	hyperflexservermodelrelationship["moid"] = item.Moid
	hyperflexservermodelrelationship["object_type"] = item.ObjectType
	hyperflexservermodelrelationship["selector"] = item.Selector

	hyperflexservermodelrelationships = append(hyperflexservermodelrelationships, hyperflexservermodelrelationship)
	return hyperflexservermodelrelationships
}
func flattenMapHyperflexSoftwareVersionPolicyRelationship(p models.HyperflexSoftwareVersionPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsoftwareversionpolicyrelationships []map[string]interface{}
	var ret models.HyperflexSoftwareVersionPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexsoftwareversionpolicyrelationship := make(map[string]interface{})
	hyperflexsoftwareversionpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsoftwareversionpolicyrelationship["class_id"] = item.ClassId
	hyperflexsoftwareversionpolicyrelationship["moid"] = item.Moid
	hyperflexsoftwareversionpolicyrelationship["object_type"] = item.ObjectType
	hyperflexsoftwareversionpolicyrelationship["selector"] = item.Selector

	hyperflexsoftwareversionpolicyrelationships = append(hyperflexsoftwareversionpolicyrelationships, hyperflexsoftwareversionpolicyrelationship)
	return hyperflexsoftwareversionpolicyrelationships
}
func flattenMapHyperflexSummary(p models.HyperflexSummary, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsummarys []map[string]interface{}
	var ret models.HyperflexSummary
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexsummary := make(map[string]interface{})
	hyperflexsummary["active_nodes"] = item.ActiveNodes
	hyperflexsummary["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsummary["address"] = item.Address
	hyperflexsummary["boottime"] = item.Boottime
	hyperflexsummary["class_id"] = item.ClassId
	hyperflexsummary["cluster_access_policy"] = item.ClusterAccessPolicy
	hyperflexsummary["compression_savings"] = item.CompressionSavings
	hyperflexsummary["data_replication_compliance"] = item.DataReplicationCompliance
	hyperflexsummary["data_replication_factor"] = item.DataReplicationFactor
	hyperflexsummary["deduplication_savings"] = item.DeduplicationSavings
	hyperflexsummary["downtime"] = item.Downtime
	hyperflexsummary["free_capacity"] = item.FreeCapacity
	hyperflexsummary["healing_info"] = (func(p models.HyperflexStPlatformClusterHealingInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterhealinginfos []map[string]interface{}
		var ret models.HyperflexStPlatformClusterHealingInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexstplatformclusterhealinginfo := make(map[string]interface{})
		hyperflexstplatformclusterhealinginfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexstplatformclusterhealinginfo["class_id"] = item.ClassId
		hyperflexstplatformclusterhealinginfo["estimated_completion_time_in_seconds"] = item.EstimatedCompletionTimeInSeconds
		hyperflexstplatformclusterhealinginfo["in_progress"] = item.InProgress
		hyperflexstplatformclusterhealinginfo["messages"] = item.Messages
		hyperflexstplatformclusterhealinginfo["messages_iterator"] = item.MessagesIterator
		hyperflexstplatformclusterhealinginfo["messages_size"] = item.MessagesSize
		hyperflexstplatformclusterhealinginfo["object_type"] = item.ObjectType
		hyperflexstplatformclusterhealinginfo["percent_complete"] = item.PercentComplete

		hyperflexstplatformclusterhealinginfos = append(hyperflexstplatformclusterhealinginfos, hyperflexstplatformclusterhealinginfo)
		return hyperflexstplatformclusterhealinginfos
	})(item.GetHealingInfo(), d)
	hyperflexsummary["name"] = item.Name
	hyperflexsummary["object_type"] = item.ObjectType
	hyperflexsummary["resiliency_details"] = item.ResiliencyDetails
	hyperflexsummary["resiliency_details_size"] = item.ResiliencyDetailsSize
	hyperflexsummary["resiliency_info"] = (func(p models.HyperflexStPlatformClusterResiliencyInfo, d *schema.ResourceData) []map[string]interface{} {
		var hyperflexstplatformclusterresiliencyinfos []map[string]interface{}
		var ret models.HyperflexStPlatformClusterResiliencyInfo
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		hyperflexstplatformclusterresiliencyinfo := make(map[string]interface{})
		hyperflexstplatformclusterresiliencyinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		hyperflexstplatformclusterresiliencyinfo["class_id"] = item.ClassId
		hyperflexstplatformclusterresiliencyinfo["hdd_failures_tolerable"] = item.HddFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["messages"] = item.Messages
		hyperflexstplatformclusterresiliencyinfo["messages_iterator"] = item.MessagesIterator
		hyperflexstplatformclusterresiliencyinfo["messages_size"] = item.MessagesSize
		hyperflexstplatformclusterresiliencyinfo["node_failures_tolerable"] = item.NodeFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["object_type"] = item.ObjectType
		hyperflexstplatformclusterresiliencyinfo["ssd_failures_tolerable"] = item.SsdFailuresTolerable
		hyperflexstplatformclusterresiliencyinfo["state"] = item.State

		hyperflexstplatformclusterresiliencyinfos = append(hyperflexstplatformclusterresiliencyinfos, hyperflexstplatformclusterresiliencyinfo)
		return hyperflexstplatformclusterresiliencyinfos
	})(item.GetResiliencyInfo(), d)
	hyperflexsummary["space_status"] = item.SpaceStatus
	hyperflexsummary["state"] = item.State
	hyperflexsummary["total_capacity"] = item.TotalCapacity
	hyperflexsummary["total_savings"] = item.TotalSavings
	hyperflexsummary["uptime"] = item.Uptime
	hyperflexsummary["used_capacity"] = item.UsedCapacity

	hyperflexsummarys = append(hyperflexsummarys, hyperflexsummary)
	return hyperflexsummarys
}
func flattenMapHyperflexSysConfigPolicyRelationship(p models.HyperflexSysConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexsysconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexSysConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexsysconfigpolicyrelationship := make(map[string]interface{})
	hyperflexsysconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexsysconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexsysconfigpolicyrelationship["moid"] = item.Moid
	hyperflexsysconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexsysconfigpolicyrelationship["selector"] = item.Selector

	hyperflexsysconfigpolicyrelationships = append(hyperflexsysconfigpolicyrelationships, hyperflexsysconfigpolicyrelationship)
	return hyperflexsysconfigpolicyrelationships
}
func flattenMapHyperflexUcsmConfigPolicyRelationship(p models.HyperflexUcsmConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexucsmconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexUcsmConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexucsmconfigpolicyrelationship := make(map[string]interface{})
	hyperflexucsmconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexucsmconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexucsmconfigpolicyrelationship["moid"] = item.Moid
	hyperflexucsmconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexucsmconfigpolicyrelationship["selector"] = item.Selector

	hyperflexucsmconfigpolicyrelationships = append(hyperflexucsmconfigpolicyrelationships, hyperflexucsmconfigpolicyrelationship)
	return hyperflexucsmconfigpolicyrelationships
}
func flattenMapHyperflexVcenterConfigPolicyRelationship(p models.HyperflexVcenterConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexvcenterconfigpolicyrelationships []map[string]interface{}
	var ret models.HyperflexVcenterConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	hyperflexvcenterconfigpolicyrelationship := make(map[string]interface{})
	hyperflexvcenterconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexvcenterconfigpolicyrelationship["class_id"] = item.ClassId
	hyperflexvcenterconfigpolicyrelationship["moid"] = item.Moid
	hyperflexvcenterconfigpolicyrelationship["object_type"] = item.ObjectType
	hyperflexvcenterconfigpolicyrelationship["selector"] = item.Selector

	hyperflexvcenterconfigpolicyrelationships = append(hyperflexvcenterconfigpolicyrelationships, hyperflexvcenterconfigpolicyrelationship)
	return hyperflexvcenterconfigpolicyrelationships
}
func flattenMapHyperflexWwxnPrefixRange(p models.HyperflexWwxnPrefixRange, d *schema.ResourceData) []map[string]interface{} {
	var hyperflexwwxnprefixranges []map[string]interface{}
	var ret models.HyperflexWwxnPrefixRange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	hyperflexwwxnprefixrange := make(map[string]interface{})
	hyperflexwwxnprefixrange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	hyperflexwwxnprefixrange["class_id"] = item.ClassId
	hyperflexwwxnprefixrange["end_addr"] = item.EndAddr
	hyperflexwwxnprefixrange["object_type"] = item.ObjectType
	hyperflexwwxnprefixrange["start_addr"] = item.StartAddr

	hyperflexwwxnprefixranges = append(hyperflexwwxnprefixranges, hyperflexwwxnprefixrange)
	return hyperflexwwxnprefixranges
}
func flattenMapIaasLicenseInfoRelationship(p models.IaasLicenseInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaaslicenseinforelationships []map[string]interface{}
	var ret models.IaasLicenseInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iaaslicenseinforelationship := make(map[string]interface{})
	iaaslicenseinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iaaslicenseinforelationship["class_id"] = item.ClassId
	iaaslicenseinforelationship["moid"] = item.Moid
	iaaslicenseinforelationship["object_type"] = item.ObjectType
	iaaslicenseinforelationship["selector"] = item.Selector

	iaaslicenseinforelationships = append(iaaslicenseinforelationships, iaaslicenseinforelationship)
	return iaaslicenseinforelationships
}
func flattenMapIaasUcsdInfoRelationship(p models.IaasUcsdInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdinforelationships []map[string]interface{}
	var ret models.IaasUcsdInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iaasucsdinforelationship := make(map[string]interface{})
	iaasucsdinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iaasucsdinforelationship["class_id"] = item.ClassId
	iaasucsdinforelationship["moid"] = item.Moid
	iaasucsdinforelationship["object_type"] = item.ObjectType
	iaasucsdinforelationship["selector"] = item.Selector

	iaasucsdinforelationships = append(iaasucsdinforelationships, iaasucsdinforelationship)
	return iaasucsdinforelationships
}
func flattenMapIaasUcsdManagedInfraRelationship(p models.IaasUcsdManagedInfraRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iaasucsdmanagedinfrarelationships []map[string]interface{}
	var ret models.IaasUcsdManagedInfraRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iaasucsdmanagedinfrarelationship := make(map[string]interface{})
	iaasucsdmanagedinfrarelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iaasucsdmanagedinfrarelationship["class_id"] = item.ClassId
	iaasucsdmanagedinfrarelationship["moid"] = item.Moid
	iaasucsdmanagedinfrarelationship["object_type"] = item.ObjectType
	iaasucsdmanagedinfrarelationship["selector"] = item.Selector

	iaasucsdmanagedinfrarelationships = append(iaasucsdmanagedinfrarelationships, iaasucsdmanagedinfrarelationship)
	return iaasucsdmanagedinfrarelationships
}
func flattenMapIamAccountRelationship(p models.IamAccountRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamaccountrelationships []map[string]interface{}
	var ret models.IamAccountRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamaccountrelationship := make(map[string]interface{})
	iamaccountrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamaccountrelationship["class_id"] = item.ClassId
	iamaccountrelationship["moid"] = item.Moid
	iamaccountrelationship["object_type"] = item.ObjectType
	iamaccountrelationship["selector"] = item.Selector

	iamaccountrelationships = append(iamaccountrelationships, iamaccountrelationship)
	return iamaccountrelationships
}
func flattenMapIamAppRegistrationRelationship(p models.IamAppRegistrationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamappregistrationrelationships []map[string]interface{}
	var ret models.IamAppRegistrationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamappregistrationrelationship := make(map[string]interface{})
	iamappregistrationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamappregistrationrelationship["class_id"] = item.ClassId
	iamappregistrationrelationship["moid"] = item.Moid
	iamappregistrationrelationship["object_type"] = item.ObjectType
	iamappregistrationrelationship["selector"] = item.Selector

	iamappregistrationrelationships = append(iamappregistrationrelationships, iamappregistrationrelationship)
	return iamappregistrationrelationships
}
func flattenMapIamCertificateRelationship(p models.IamCertificateRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterelationships []map[string]interface{}
	var ret models.IamCertificateRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamcertificaterelationship := make(map[string]interface{})
	iamcertificaterelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamcertificaterelationship["class_id"] = item.ClassId
	iamcertificaterelationship["moid"] = item.Moid
	iamcertificaterelationship["object_type"] = item.ObjectType
	iamcertificaterelationship["selector"] = item.Selector

	iamcertificaterelationships = append(iamcertificaterelationships, iamcertificaterelationship)
	return iamcertificaterelationships
}
func flattenMapIamCertificateRequestRelationship(p models.IamCertificateRequestRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamcertificaterequestrelationships []map[string]interface{}
	var ret models.IamCertificateRequestRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamcertificaterequestrelationship := make(map[string]interface{})
	iamcertificaterequestrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamcertificaterequestrelationship["class_id"] = item.ClassId
	iamcertificaterequestrelationship["moid"] = item.Moid
	iamcertificaterequestrelationship["object_type"] = item.ObjectType
	iamcertificaterequestrelationship["selector"] = item.Selector

	iamcertificaterequestrelationships = append(iamcertificaterequestrelationships, iamcertificaterequestrelationship)
	return iamcertificaterequestrelationships
}
func flattenMapIamClientMeta(p models.IamClientMeta, d *schema.ResourceData) []map[string]interface{} {
	var iamclientmetas []map[string]interface{}
	var ret models.IamClientMeta
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamclientmeta := make(map[string]interface{})
	iamclientmeta["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamclientmeta["class_id"] = item.ClassId
	iamclientmeta["device_model"] = item.DeviceModel
	iamclientmeta["object_type"] = item.ObjectType
	iamclientmeta["user_agent"] = item.UserAgent

	iamclientmetas = append(iamclientmetas, iamclientmeta)
	return iamclientmetas
}
func flattenMapIamDomainGroupRelationship(p models.IamDomainGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamdomaingrouprelationships []map[string]interface{}
	var ret models.IamDomainGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamdomaingrouprelationship := make(map[string]interface{})
	iamdomaingrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamdomaingrouprelationship["class_id"] = item.ClassId
	iamdomaingrouprelationship["moid"] = item.Moid
	iamdomaingrouprelationship["object_type"] = item.ObjectType
	iamdomaingrouprelationship["selector"] = item.Selector

	iamdomaingrouprelationships = append(iamdomaingrouprelationships, iamdomaingrouprelationship)
	return iamdomaingrouprelationships
}
func flattenMapIamEndPointPasswordProperties(p models.IamEndPointPasswordProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointpasswordpropertiess []map[string]interface{}
	var ret models.IamEndPointPasswordProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamendpointpasswordproperties := make(map[string]interface{})
	iamendpointpasswordproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamendpointpasswordproperties["class_id"] = item.ClassId
	iamendpointpasswordproperties["enable_password_expiry"] = item.EnablePasswordExpiry
	iamendpointpasswordproperties["enforce_strong_password"] = item.EnforceStrongPassword
	iamendpointpasswordproperties["force_send_password"] = item.ForceSendPassword
	iamendpointpasswordproperties["grace_period"] = item.GracePeriod
	iamendpointpasswordproperties["notification_period"] = item.NotificationPeriod
	iamendpointpasswordproperties["object_type"] = item.ObjectType
	iamendpointpasswordproperties["password_expiry_duration"] = item.PasswordExpiryDuration
	iamendpointpasswordproperties["password_history"] = item.PasswordHistory

	iamendpointpasswordpropertiess = append(iamendpointpasswordpropertiess, iamendpointpasswordproperties)
	return iamendpointpasswordpropertiess
}
func flattenMapIamEndPointUserRelationship(p models.IamEndPointUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserrelationships []map[string]interface{}
	var ret models.IamEndPointUserRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamendpointuserrelationship := make(map[string]interface{})
	iamendpointuserrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamendpointuserrelationship["class_id"] = item.ClassId
	iamendpointuserrelationship["moid"] = item.Moid
	iamendpointuserrelationship["object_type"] = item.ObjectType
	iamendpointuserrelationship["selector"] = item.Selector

	iamendpointuserrelationships = append(iamendpointuserrelationships, iamendpointuserrelationship)
	return iamendpointuserrelationships
}
func flattenMapIamEndPointUserPolicyRelationship(p models.IamEndPointUserPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamendpointuserpolicyrelationships []map[string]interface{}
	var ret models.IamEndPointUserPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamendpointuserpolicyrelationship := make(map[string]interface{})
	iamendpointuserpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamendpointuserpolicyrelationship["class_id"] = item.ClassId
	iamendpointuserpolicyrelationship["moid"] = item.Moid
	iamendpointuserpolicyrelationship["object_type"] = item.ObjectType
	iamendpointuserpolicyrelationship["selector"] = item.Selector

	iamendpointuserpolicyrelationships = append(iamendpointuserpolicyrelationships, iamendpointuserpolicyrelationship)
	return iamendpointuserpolicyrelationships
}
func flattenMapIamIdpRelationship(p models.IamIdpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidprelationships []map[string]interface{}
	var ret models.IamIdpRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamidprelationship := make(map[string]interface{})
	iamidprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamidprelationship["class_id"] = item.ClassId
	iamidprelationship["moid"] = item.Moid
	iamidprelationship["object_type"] = item.ObjectType
	iamidprelationship["selector"] = item.Selector

	iamidprelationships = append(iamidprelationships, iamidprelationship)
	return iamidprelationships
}
func flattenMapIamIdpReferenceRelationship(p models.IamIdpReferenceRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamidpreferencerelationships []map[string]interface{}
	var ret models.IamIdpReferenceRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamidpreferencerelationship := make(map[string]interface{})
	iamidpreferencerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamidpreferencerelationship["class_id"] = item.ClassId
	iamidpreferencerelationship["moid"] = item.Moid
	iamidpreferencerelationship["object_type"] = item.ObjectType
	iamidpreferencerelationship["selector"] = item.Selector

	iamidpreferencerelationships = append(iamidpreferencerelationships, iamidpreferencerelationship)
	return iamidpreferencerelationships
}
func flattenMapIamLdapBaseProperties(p models.IamLdapBaseProperties, d *schema.ResourceData) []map[string]interface{} {
	var iamldapbasepropertiess []map[string]interface{}
	var ret models.IamLdapBaseProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamldapbaseproperties := make(map[string]interface{})
	iamldapbaseproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamldapbaseproperties["attribute"] = item.Attribute
	iamldapbaseproperties["base_dn"] = item.BaseDn
	iamldapbaseproperties["bind_dn"] = item.BindDn
	iamldapbaseproperties["bind_method"] = item.BindMethod
	iamldapbaseproperties["class_id"] = item.ClassId
	iamldapbaseproperties["domain"] = item.Domain
	iamldapbaseproperties["enable_encryption"] = item.EnableEncryption
	iamldapbaseproperties["enable_group_authorization"] = item.EnableGroupAuthorization
	iamldapbaseproperties["filter"] = item.Filter
	iamldapbaseproperties["group_attribute"] = item.GroupAttribute
	iamldapbaseproperties["is_password_set"] = item.IsPasswordSet
	iamldapbaseproperties["nested_group_search_depth"] = item.NestedGroupSearchDepth
	iamldapbaseproperties["object_type"] = item.ObjectType
	password_x := d.Get("base_properties").([]interface{})
	password_y := password_x[0].(map[string]interface{})
	iamldapbaseproperties["password"] = password_y["password"]
	iamldapbaseproperties["timeout"] = item.Timeout

	iamldapbasepropertiess = append(iamldapbasepropertiess, iamldapbaseproperties)
	return iamldapbasepropertiess
}
func flattenMapIamLdapDnsParameters(p models.IamLdapDnsParameters, d *schema.ResourceData) []map[string]interface{} {
	var iamldapdnsparameterss []map[string]interface{}
	var ret models.IamLdapDnsParameters
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	iamldapdnsparameters := make(map[string]interface{})
	iamldapdnsparameters["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamldapdnsparameters["class_id"] = item.ClassId
	iamldapdnsparameters["object_type"] = item.ObjectType
	iamldapdnsparameters["search_domain"] = item.SearchDomain
	iamldapdnsparameters["search_forest"] = item.SearchForest
	iamldapdnsparameters["nr_source"] = item.Source

	iamldapdnsparameterss = append(iamldapdnsparameterss, iamldapdnsparameters)
	return iamldapdnsparameterss
}
func flattenMapIamLdapPolicyRelationship(p models.IamLdapPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamldappolicyrelationships []map[string]interface{}
	var ret models.IamLdapPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamldappolicyrelationship := make(map[string]interface{})
	iamldappolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamldappolicyrelationship["class_id"] = item.ClassId
	iamldappolicyrelationship["moid"] = item.Moid
	iamldappolicyrelationship["object_type"] = item.ObjectType
	iamldappolicyrelationship["selector"] = item.Selector

	iamldappolicyrelationships = append(iamldappolicyrelationships, iamldappolicyrelationship)
	return iamldappolicyrelationships
}
func flattenMapIamLocalUserPasswordRelationship(p models.IamLocalUserPasswordRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamlocaluserpasswordrelationships []map[string]interface{}
	var ret models.IamLocalUserPasswordRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamlocaluserpasswordrelationship := make(map[string]interface{})
	iamlocaluserpasswordrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamlocaluserpasswordrelationship["class_id"] = item.ClassId
	iamlocaluserpasswordrelationship["moid"] = item.Moid
	iamlocaluserpasswordrelationship["object_type"] = item.ObjectType
	iamlocaluserpasswordrelationship["selector"] = item.Selector

	iamlocaluserpasswordrelationships = append(iamlocaluserpasswordrelationships, iamlocaluserpasswordrelationship)
	return iamlocaluserpasswordrelationships
}
func flattenMapIamPermissionRelationship(p models.IamPermissionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iampermissionrelationships []map[string]interface{}
	var ret models.IamPermissionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iampermissionrelationship := make(map[string]interface{})
	iampermissionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iampermissionrelationship["class_id"] = item.ClassId
	iampermissionrelationship["moid"] = item.Moid
	iampermissionrelationship["object_type"] = item.ObjectType
	iampermissionrelationship["selector"] = item.Selector

	iampermissionrelationships = append(iampermissionrelationships, iampermissionrelationship)
	return iampermissionrelationships
}
func flattenMapIamPrivateKeySpecRelationship(p models.IamPrivateKeySpecRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamprivatekeyspecrelationships []map[string]interface{}
	var ret models.IamPrivateKeySpecRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamprivatekeyspecrelationship := make(map[string]interface{})
	iamprivatekeyspecrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamprivatekeyspecrelationship["class_id"] = item.ClassId
	iamprivatekeyspecrelationship["moid"] = item.Moid
	iamprivatekeyspecrelationship["object_type"] = item.ObjectType
	iamprivatekeyspecrelationship["selector"] = item.Selector

	iamprivatekeyspecrelationships = append(iamprivatekeyspecrelationships, iamprivatekeyspecrelationship)
	return iamprivatekeyspecrelationships
}
func flattenMapIamQualifierRelationship(p models.IamQualifierRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamqualifierrelationships []map[string]interface{}
	var ret models.IamQualifierRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamqualifierrelationship := make(map[string]interface{})
	iamqualifierrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamqualifierrelationship["class_id"] = item.ClassId
	iamqualifierrelationship["moid"] = item.Moid
	iamqualifierrelationship["object_type"] = item.ObjectType
	iamqualifierrelationship["selector"] = item.Selector

	iamqualifierrelationships = append(iamqualifierrelationships, iamqualifierrelationship)
	return iamqualifierrelationships
}
func flattenMapIamResourceLimitsRelationship(p models.IamResourceLimitsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamresourcelimitsrelationships []map[string]interface{}
	var ret models.IamResourceLimitsRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamresourcelimitsrelationship := make(map[string]interface{})
	iamresourcelimitsrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamresourcelimitsrelationship["class_id"] = item.ClassId
	iamresourcelimitsrelationship["moid"] = item.Moid
	iamresourcelimitsrelationship["object_type"] = item.ObjectType
	iamresourcelimitsrelationship["selector"] = item.Selector

	iamresourcelimitsrelationships = append(iamresourcelimitsrelationships, iamresourcelimitsrelationship)
	return iamresourcelimitsrelationships
}
func flattenMapIamSecurityHolderRelationship(p models.IamSecurityHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsecurityholderrelationships []map[string]interface{}
	var ret models.IamSecurityHolderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsecurityholderrelationship := make(map[string]interface{})
	iamsecurityholderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsecurityholderrelationship["class_id"] = item.ClassId
	iamsecurityholderrelationship["moid"] = item.Moid
	iamsecurityholderrelationship["object_type"] = item.ObjectType
	iamsecurityholderrelationship["selector"] = item.Selector

	iamsecurityholderrelationships = append(iamsecurityholderrelationships, iamsecurityholderrelationship)
	return iamsecurityholderrelationships
}
func flattenMapIamServiceProviderRelationship(p models.IamServiceProviderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamserviceproviderrelationships []map[string]interface{}
	var ret models.IamServiceProviderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamserviceproviderrelationship := make(map[string]interface{})
	iamserviceproviderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamserviceproviderrelationship["class_id"] = item.ClassId
	iamserviceproviderrelationship["moid"] = item.Moid
	iamserviceproviderrelationship["object_type"] = item.ObjectType
	iamserviceproviderrelationship["selector"] = item.Selector

	iamserviceproviderrelationships = append(iamserviceproviderrelationships, iamserviceproviderrelationship)
	return iamserviceproviderrelationships
}
func flattenMapIamSessionRelationship(p models.IamSessionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionrelationships []map[string]interface{}
	var ret models.IamSessionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsessionrelationship := make(map[string]interface{})
	iamsessionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsessionrelationship["class_id"] = item.ClassId
	iamsessionrelationship["moid"] = item.Moid
	iamsessionrelationship["object_type"] = item.ObjectType
	iamsessionrelationship["selector"] = item.Selector

	iamsessionrelationships = append(iamsessionrelationships, iamsessionrelationship)
	return iamsessionrelationships
}
func flattenMapIamSessionLimitsRelationship(p models.IamSessionLimitsRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsessionlimitsrelationships []map[string]interface{}
	var ret models.IamSessionLimitsRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsessionlimitsrelationship := make(map[string]interface{})
	iamsessionlimitsrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsessionlimitsrelationship["class_id"] = item.ClassId
	iamsessionlimitsrelationship["moid"] = item.Moid
	iamsessionlimitsrelationship["object_type"] = item.ObjectType
	iamsessionlimitsrelationship["selector"] = item.Selector

	iamsessionlimitsrelationships = append(iamsessionlimitsrelationships, iamsessionlimitsrelationship)
	return iamsessionlimitsrelationships
}
func flattenMapIamSystemRelationship(p models.IamSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamsystemrelationships []map[string]interface{}
	var ret models.IamSystemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamsystemrelationship := make(map[string]interface{})
	iamsystemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamsystemrelationship["class_id"] = item.ClassId
	iamsystemrelationship["moid"] = item.Moid
	iamsystemrelationship["object_type"] = item.ObjectType
	iamsystemrelationship["selector"] = item.Selector

	iamsystemrelationships = append(iamsystemrelationships, iamsystemrelationship)
	return iamsystemrelationships
}
func flattenMapIamUserRelationship(p models.IamUserRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamuserrelationships []map[string]interface{}
	var ret models.IamUserRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamuserrelationship := make(map[string]interface{})
	iamuserrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamuserrelationship["class_id"] = item.ClassId
	iamuserrelationship["moid"] = item.Moid
	iamuserrelationship["object_type"] = item.ObjectType
	iamuserrelationship["selector"] = item.Selector

	iamuserrelationships = append(iamuserrelationships, iamuserrelationship)
	return iamuserrelationships
}
func flattenMapIamUserGroupRelationship(p models.IamUserGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var iamusergrouprelationships []map[string]interface{}
	var ret models.IamUserGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	iamusergrouprelationship := make(map[string]interface{})
	iamusergrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	iamusergrouprelationship["class_id"] = item.ClassId
	iamusergrouprelationship["moid"] = item.Moid
	iamusergrouprelationship["object_type"] = item.ObjectType
	iamusergrouprelationship["selector"] = item.Selector

	iamusergrouprelationships = append(iamusergrouprelationships, iamusergrouprelationship)
	return iamusergrouprelationships
}
func flattenMapInfraHardwareInfo(p models.InfraHardwareInfo, d *schema.ResourceData) []map[string]interface{} {
	var infrahardwareinfos []map[string]interface{}
	var ret models.InfraHardwareInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	infrahardwareinfo := make(map[string]interface{})
	infrahardwareinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	infrahardwareinfo["class_id"] = item.ClassId
	infrahardwareinfo["cpu_cores"] = item.CpuCores
	infrahardwareinfo["cpu_speed"] = item.CpuSpeed
	infrahardwareinfo["memory_size"] = item.MemorySize
	infrahardwareinfo["object_type"] = item.ObjectType

	infrahardwareinfos = append(infrahardwareinfos, infrahardwareinfo)
	return infrahardwareinfos
}
func flattenMapInventoryBaseRelationship(p models.InventoryBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorybaserelationships []map[string]interface{}
	var ret models.InventoryBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	inventorybaserelationship := make(map[string]interface{})
	inventorybaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventorybaserelationship["class_id"] = item.ClassId
	inventorybaserelationship["moid"] = item.Moid
	inventorybaserelationship["object_type"] = item.ObjectType
	inventorybaserelationship["selector"] = item.Selector

	inventorybaserelationships = append(inventorybaserelationships, inventorybaserelationship)
	return inventorybaserelationships
}
func flattenMapInventoryDeviceInfoRelationship(p models.InventoryDeviceInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorydeviceinforelationships []map[string]interface{}
	var ret models.InventoryDeviceInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	inventorydeviceinforelationship := make(map[string]interface{})
	inventorydeviceinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventorydeviceinforelationship["class_id"] = item.ClassId
	inventorydeviceinforelationship["moid"] = item.Moid
	inventorydeviceinforelationship["object_type"] = item.ObjectType
	inventorydeviceinforelationship["selector"] = item.Selector

	inventorydeviceinforelationships = append(inventorydeviceinforelationships, inventorydeviceinforelationship)
	return inventorydeviceinforelationships
}
func flattenMapInventoryEpInfo(p models.InventoryEpInfo, d *schema.ResourceData) []map[string]interface{} {
	var inventoryepinfos []map[string]interface{}
	var ret models.InventoryEpInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	inventoryepinfo := make(map[string]interface{})
	inventoryepinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventoryepinfo["class_id"] = item.ClassId
	inventoryepinfo["connection_status"] = item.ConnectionStatus
	inventoryepinfo["ep_type"] = item.EpType
	inventoryepinfo["ip"] = item.Ip
	inventoryepinfo["mac_address"] = item.MacAddress
	inventoryepinfo["member_identity"] = item.MemberIdentity
	inventoryepinfo["model"] = item.Model
	inventoryepinfo["object_type"] = item.ObjectType
	inventoryepinfo["serial"] = item.Serial
	inventoryepinfo["server_registration_id"] = item.ServerRegistrationId
	inventoryepinfo["vendor"] = item.Vendor

	inventoryepinfos = append(inventoryepinfos, inventoryepinfo)
	return inventoryepinfos
}
func flattenMapInventoryGenericInventoryHolderRelationship(p models.InventoryGenericInventoryHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventorygenericinventoryholderrelationships []map[string]interface{}
	var ret models.InventoryGenericInventoryHolderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	inventorygenericinventoryholderrelationship := make(map[string]interface{})
	inventorygenericinventoryholderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventorygenericinventoryholderrelationship["class_id"] = item.ClassId
	inventorygenericinventoryholderrelationship["moid"] = item.Moid
	inventorygenericinventoryholderrelationship["object_type"] = item.ObjectType
	inventorygenericinventoryholderrelationship["selector"] = item.Selector

	inventorygenericinventoryholderrelationships = append(inventorygenericinventoryholderrelationships, inventorygenericinventoryholderrelationship)
	return inventorygenericinventoryholderrelationships
}
func flattenMapInventoryUemConnectionRelationship(p models.InventoryUemConnectionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var inventoryuemconnectionrelationships []map[string]interface{}
	var ret models.InventoryUemConnectionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	inventoryuemconnectionrelationship := make(map[string]interface{})
	inventoryuemconnectionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	inventoryuemconnectionrelationship["class_id"] = item.ClassId
	inventoryuemconnectionrelationship["moid"] = item.Moid
	inventoryuemconnectionrelationship["object_type"] = item.ObjectType
	inventoryuemconnectionrelationship["selector"] = item.Selector

	inventoryuemconnectionrelationships = append(inventoryuemconnectionrelationships, inventoryuemconnectionrelationship)
	return inventoryuemconnectionrelationships
}
func flattenMapIppoolIpBlock(p models.IppoolIpBlock, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipblocks []map[string]interface{}
	var ret models.IppoolIpBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	ippoolipblock := make(map[string]interface{})
	ippoolipblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipblock["class_id"] = item.ClassId
	ippoolipblock["from"] = item.From
	ippoolipblock["object_type"] = item.ObjectType
	ippoolipblock["size"] = item.Size
	ippoolipblock["to"] = item.To

	ippoolipblocks = append(ippoolipblocks, ippoolipblock)
	return ippoolipblocks
}
func flattenMapIppoolIpLeaseRelationship(p models.IppoolIpLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipleaserelationships []map[string]interface{}
	var ret models.IppoolIpLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolipleaserelationship := make(map[string]interface{})
	ippoolipleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipleaserelationship["class_id"] = item.ClassId
	ippoolipleaserelationship["moid"] = item.Moid
	ippoolipleaserelationship["object_type"] = item.ObjectType
	ippoolipleaserelationship["selector"] = item.Selector

	ippoolipleaserelationships = append(ippoolipleaserelationships, ippoolipleaserelationship)
	return ippoolipleaserelationships
}
func flattenMapIppoolIpV4Config(p models.IppoolIpV4Config, d *schema.ResourceData) []map[string]interface{} {
	var ippoolipv4configs []map[string]interface{}
	var ret models.IppoolIpV4Config
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	ippoolipv4config := make(map[string]interface{})
	ippoolipv4config["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolipv4config["class_id"] = item.ClassId
	ippoolipv4config["gateway"] = item.Gateway
	ippoolipv4config["netmask"] = item.Netmask
	ippoolipv4config["object_type"] = item.ObjectType
	ippoolipv4config["primary_dns"] = item.PrimaryDns
	ippoolipv4config["secondary_dns"] = item.SecondaryDns

	ippoolipv4configs = append(ippoolipv4configs, ippoolipv4config)
	return ippoolipv4configs
}
func flattenMapIppoolPoolRelationship(p models.IppoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolpoolrelationships []map[string]interface{}
	var ret models.IppoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolpoolrelationship := make(map[string]interface{})
	ippoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolpoolrelationship["class_id"] = item.ClassId
	ippoolpoolrelationship["moid"] = item.Moid
	ippoolpoolrelationship["object_type"] = item.ObjectType
	ippoolpoolrelationship["selector"] = item.Selector

	ippoolpoolrelationships = append(ippoolpoolrelationships, ippoolpoolrelationship)
	return ippoolpoolrelationships
}
func flattenMapIppoolPoolMemberRelationship(p models.IppoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolpoolmemberrelationships []map[string]interface{}
	var ret models.IppoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolpoolmemberrelationship := make(map[string]interface{})
	ippoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolpoolmemberrelationship["class_id"] = item.ClassId
	ippoolpoolmemberrelationship["moid"] = item.Moid
	ippoolpoolmemberrelationship["object_type"] = item.ObjectType
	ippoolpoolmemberrelationship["selector"] = item.Selector

	ippoolpoolmemberrelationships = append(ippoolpoolmemberrelationships, ippoolpoolmemberrelationship)
	return ippoolpoolmemberrelationships
}
func flattenMapIppoolShadowBlockRelationship(p models.IppoolShadowBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowblockrelationships []map[string]interface{}
	var ret models.IppoolShadowBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolshadowblockrelationship := make(map[string]interface{})
	ippoolshadowblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolshadowblockrelationship["class_id"] = item.ClassId
	ippoolshadowblockrelationship["moid"] = item.Moid
	ippoolshadowblockrelationship["object_type"] = item.ObjectType
	ippoolshadowblockrelationship["selector"] = item.Selector

	ippoolshadowblockrelationships = append(ippoolshadowblockrelationships, ippoolshadowblockrelationship)
	return ippoolshadowblockrelationships
}
func flattenMapIppoolShadowPoolRelationship(p models.IppoolShadowPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippoolshadowpoolrelationships []map[string]interface{}
	var ret models.IppoolShadowPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippoolshadowpoolrelationship := make(map[string]interface{})
	ippoolshadowpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippoolshadowpoolrelationship["class_id"] = item.ClassId
	ippoolshadowpoolrelationship["moid"] = item.Moid
	ippoolshadowpoolrelationship["object_type"] = item.ObjectType
	ippoolshadowpoolrelationship["selector"] = item.Selector

	ippoolshadowpoolrelationships = append(ippoolshadowpoolrelationships, ippoolshadowpoolrelationship)
	return ippoolshadowpoolrelationships
}
func flattenMapIppoolUniverseRelationship(p models.IppoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var ippooluniverserelationships []map[string]interface{}
	var ret models.IppoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	ippooluniverserelationship := make(map[string]interface{})
	ippooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	ippooluniverserelationship["class_id"] = item.ClassId
	ippooluniverserelationship["moid"] = item.Moid
	ippooluniverserelationship["object_type"] = item.ObjectType
	ippooluniverserelationship["selector"] = item.Selector

	ippooluniverserelationships = append(ippooluniverserelationships, ippooluniverserelationship)
	return ippooluniverserelationships
}
func flattenMapLicenseAccountLicenseDataRelationship(p models.LicenseAccountLicenseDataRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licenseaccountlicensedatarelationships []map[string]interface{}
	var ret models.LicenseAccountLicenseDataRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licenseaccountlicensedatarelationship := make(map[string]interface{})
	licenseaccountlicensedatarelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licenseaccountlicensedatarelationship["class_id"] = item.ClassId
	licenseaccountlicensedatarelationship["moid"] = item.Moid
	licenseaccountlicensedatarelationship["object_type"] = item.ObjectType
	licenseaccountlicensedatarelationship["selector"] = item.Selector

	licenseaccountlicensedatarelationships = append(licenseaccountlicensedatarelationships, licenseaccountlicensedatarelationship)
	return licenseaccountlicensedatarelationships
}
func flattenMapLicenseCustomerOpRelationship(p models.LicenseCustomerOpRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licensecustomeroprelationships []map[string]interface{}
	var ret models.LicenseCustomerOpRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licensecustomeroprelationship := make(map[string]interface{})
	licensecustomeroprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licensecustomeroprelationship["class_id"] = item.ClassId
	licensecustomeroprelationship["moid"] = item.Moid
	licensecustomeroprelationship["object_type"] = item.ObjectType
	licensecustomeroprelationship["selector"] = item.Selector

	licensecustomeroprelationships = append(licensecustomeroprelationships, licensecustomeroprelationship)
	return licensecustomeroprelationships
}
func flattenMapLicenseSmartlicenseTokenRelationship(p models.LicenseSmartlicenseTokenRelationship, d *schema.ResourceData) []map[string]interface{} {
	var licensesmartlicensetokenrelationships []map[string]interface{}
	var ret models.LicenseSmartlicenseTokenRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	licensesmartlicensetokenrelationship := make(map[string]interface{})
	licensesmartlicensetokenrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	licensesmartlicensetokenrelationship["class_id"] = item.ClassId
	licensesmartlicensetokenrelationship["moid"] = item.Moid
	licensesmartlicensetokenrelationship["object_type"] = item.ObjectType
	licensesmartlicensetokenrelationship["selector"] = item.Selector

	licensesmartlicensetokenrelationships = append(licensesmartlicensetokenrelationships, licensesmartlicensetokenrelationship)
	return licensesmartlicensetokenrelationships
}
func flattenMapMacpoolBlock(p models.MacpoolBlock, d *schema.ResourceData) []map[string]interface{} {
	var macpoolblocks []map[string]interface{}
	var ret models.MacpoolBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	macpoolblock := make(map[string]interface{})
	macpoolblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolblock["class_id"] = item.ClassId
	macpoolblock["from"] = item.From
	macpoolblock["object_type"] = item.ObjectType
	macpoolblock["size"] = item.Size
	macpoolblock["to"] = item.To

	macpoolblocks = append(macpoolblocks, macpoolblock)
	return macpoolblocks
}
func flattenMapMacpoolIdBlockRelationship(p models.MacpoolIdBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolidblockrelationships []map[string]interface{}
	var ret models.MacpoolIdBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolidblockrelationship := make(map[string]interface{})
	macpoolidblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolidblockrelationship["class_id"] = item.ClassId
	macpoolidblockrelationship["moid"] = item.Moid
	macpoolidblockrelationship["object_type"] = item.ObjectType
	macpoolidblockrelationship["selector"] = item.Selector

	macpoolidblockrelationships = append(macpoolidblockrelationships, macpoolidblockrelationship)
	return macpoolidblockrelationships
}
func flattenMapMacpoolLeaseRelationship(p models.MacpoolLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolleaserelationships []map[string]interface{}
	var ret models.MacpoolLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolleaserelationship := make(map[string]interface{})
	macpoolleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolleaserelationship["class_id"] = item.ClassId
	macpoolleaserelationship["moid"] = item.Moid
	macpoolleaserelationship["object_type"] = item.ObjectType
	macpoolleaserelationship["selector"] = item.Selector

	macpoolleaserelationships = append(macpoolleaserelationships, macpoolleaserelationship)
	return macpoolleaserelationships
}
func flattenMapMacpoolPoolRelationship(p models.MacpoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolpoolrelationships []map[string]interface{}
	var ret models.MacpoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolpoolrelationship := make(map[string]interface{})
	macpoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolpoolrelationship["class_id"] = item.ClassId
	macpoolpoolrelationship["moid"] = item.Moid
	macpoolpoolrelationship["object_type"] = item.ObjectType
	macpoolpoolrelationship["selector"] = item.Selector

	macpoolpoolrelationships = append(macpoolpoolrelationships, macpoolpoolrelationship)
	return macpoolpoolrelationships
}
func flattenMapMacpoolPoolMemberRelationship(p models.MacpoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpoolpoolmemberrelationships []map[string]interface{}
	var ret models.MacpoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpoolpoolmemberrelationship := make(map[string]interface{})
	macpoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpoolpoolmemberrelationship["class_id"] = item.ClassId
	macpoolpoolmemberrelationship["moid"] = item.Moid
	macpoolpoolmemberrelationship["object_type"] = item.ObjectType
	macpoolpoolmemberrelationship["selector"] = item.Selector

	macpoolpoolmemberrelationships = append(macpoolpoolmemberrelationships, macpoolpoolmemberrelationship)
	return macpoolpoolmemberrelationships
}
func flattenMapMacpoolUniverseRelationship(p models.MacpoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var macpooluniverserelationships []map[string]interface{}
	var ret models.MacpoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	macpooluniverserelationship := make(map[string]interface{})
	macpooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	macpooluniverserelationship["class_id"] = item.ClassId
	macpooluniverserelationship["moid"] = item.Moid
	macpooluniverserelationship["object_type"] = item.ObjectType
	macpooluniverserelationship["selector"] = item.Selector

	macpooluniverserelationships = append(macpooluniverserelationships, macpooluniverserelationship)
	return macpooluniverserelationships
}
func flattenMapManagementControllerRelationship(p models.ManagementControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managementcontrollerrelationships []map[string]interface{}
	var ret models.ManagementControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	managementcontrollerrelationship := make(map[string]interface{})
	managementcontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	managementcontrollerrelationship["class_id"] = item.ClassId
	managementcontrollerrelationship["moid"] = item.Moid
	managementcontrollerrelationship["object_type"] = item.ObjectType
	managementcontrollerrelationship["selector"] = item.Selector

	managementcontrollerrelationships = append(managementcontrollerrelationships, managementcontrollerrelationship)
	return managementcontrollerrelationships
}
func flattenMapManagementEntityRelationship(p models.ManagementEntityRelationship, d *schema.ResourceData) []map[string]interface{} {
	var managemententityrelationships []map[string]interface{}
	var ret models.ManagementEntityRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	managemententityrelationship := make(map[string]interface{})
	managemententityrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	managemententityrelationship["class_id"] = item.ClassId
	managemententityrelationship["moid"] = item.Moid
	managemententityrelationship["object_type"] = item.ObjectType
	managemententityrelationship["selector"] = item.Selector

	managemententityrelationships = append(managemententityrelationships, managemententityrelationship)
	return managemententityrelationships
}
func flattenMapMemoryArrayRelationship(p models.MemoryArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memoryarrayrelationships []map[string]interface{}
	var ret models.MemoryArrayRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memoryarrayrelationship := make(map[string]interface{})
	memoryarrayrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memoryarrayrelationship["class_id"] = item.ClassId
	memoryarrayrelationship["moid"] = item.Moid
	memoryarrayrelationship["object_type"] = item.ObjectType
	memoryarrayrelationship["selector"] = item.Selector

	memoryarrayrelationships = append(memoryarrayrelationships, memoryarrayrelationship)
	return memoryarrayrelationships
}
func flattenMapMemoryPersistentMemoryConfigResultRelationship(p models.MemoryPersistentMemoryConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigresultrelationships []map[string]interface{}
	var ret models.MemoryPersistentMemoryConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memorypersistentmemoryconfigresultrelationship := make(map[string]interface{})
	memorypersistentmemoryconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemoryconfigresultrelationship["class_id"] = item.ClassId
	memorypersistentmemoryconfigresultrelationship["moid"] = item.Moid
	memorypersistentmemoryconfigresultrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryconfigresultrelationship["selector"] = item.Selector

	memorypersistentmemoryconfigresultrelationships = append(memorypersistentmemoryconfigresultrelationships, memorypersistentmemoryconfigresultrelationship)
	return memorypersistentmemoryconfigresultrelationships
}
func flattenMapMemoryPersistentMemoryConfigurationRelationship(p models.MemoryPersistentMemoryConfigurationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryconfigurationrelationships []map[string]interface{}
	var ret models.MemoryPersistentMemoryConfigurationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memorypersistentmemoryconfigurationrelationship := make(map[string]interface{})
	memorypersistentmemoryconfigurationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemoryconfigurationrelationship["class_id"] = item.ClassId
	memorypersistentmemoryconfigurationrelationship["moid"] = item.Moid
	memorypersistentmemoryconfigurationrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryconfigurationrelationship["selector"] = item.Selector

	memorypersistentmemoryconfigurationrelationships = append(memorypersistentmemoryconfigurationrelationships, memorypersistentmemoryconfigurationrelationship)
	return memorypersistentmemoryconfigurationrelationships
}
func flattenMapMemoryPersistentMemoryLocalSecurity(p models.MemoryPersistentMemoryLocalSecurity, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemorylocalsecuritys []map[string]interface{}
	var ret models.MemoryPersistentMemoryLocalSecurity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	memorypersistentmemorylocalsecurity := make(map[string]interface{})
	memorypersistentmemorylocalsecurity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemorylocalsecurity["class_id"] = item.ClassId
	memorypersistentmemorylocalsecurity["enabled"] = item.Enabled
	memorypersistentmemorylocalsecurity["is_secure_passphrase_set"] = item.IsSecurePassphraseSet
	memorypersistentmemorylocalsecurity["object_type"] = item.ObjectType
	secure_passphrase_x := d.Get("local_security").([]interface{})
	secure_passphrase_y := secure_passphrase_x[0].(map[string]interface{})
	memorypersistentmemorylocalsecurity["secure_passphrase"] = secure_passphrase_y["secure_passphrase"]

	memorypersistentmemorylocalsecuritys = append(memorypersistentmemorylocalsecuritys, memorypersistentmemorylocalsecurity)
	return memorypersistentmemorylocalsecuritys
}
func flattenMapMemoryPersistentMemoryRegionRelationship(p models.MemoryPersistentMemoryRegionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var memorypersistentmemoryregionrelationships []map[string]interface{}
	var ret models.MemoryPersistentMemoryRegionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	memorypersistentmemoryregionrelationship := make(map[string]interface{})
	memorypersistentmemoryregionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	memorypersistentmemoryregionrelationship["class_id"] = item.ClassId
	memorypersistentmemoryregionrelationship["moid"] = item.Moid
	memorypersistentmemoryregionrelationship["object_type"] = item.ObjectType
	memorypersistentmemoryregionrelationship["selector"] = item.Selector

	memorypersistentmemoryregionrelationships = append(memorypersistentmemoryregionrelationships, memorypersistentmemoryregionrelationship)
	return memorypersistentmemoryregionrelationships
}
func flattenMapMoBaseMoRelationship(p models.MoBaseMoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var mobasemorelationships []map[string]interface{}
	var ret models.MoBaseMoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	mobasemorelationship := make(map[string]interface{})
	mobasemorelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	mobasemorelationship["class_id"] = item.ClassId
	mobasemorelationship["moid"] = item.Moid
	mobasemorelationship["object_type"] = item.ObjectType
	mobasemorelationship["selector"] = item.Selector

	mobasemorelationships = append(mobasemorelationships, mobasemorelationship)
	return mobasemorelationships
}
func flattenMapNetworkElementRelationship(p models.NetworkElementRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkelementrelationships []map[string]interface{}
	var ret models.NetworkElementRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	networkelementrelationship := make(map[string]interface{})
	networkelementrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	networkelementrelationship["class_id"] = item.ClassId
	networkelementrelationship["moid"] = item.Moid
	networkelementrelationship["object_type"] = item.ObjectType
	networkelementrelationship["selector"] = item.Selector

	networkelementrelationships = append(networkelementrelationships, networkelementrelationship)
	return networkelementrelationships
}
func flattenMapNetworkFcZoneInfoRelationship(p models.NetworkFcZoneInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkfczoneinforelationships []map[string]interface{}
	var ret models.NetworkFcZoneInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	networkfczoneinforelationship := make(map[string]interface{})
	networkfczoneinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	networkfczoneinforelationship["class_id"] = item.ClassId
	networkfczoneinforelationship["moid"] = item.Moid
	networkfczoneinforelationship["object_type"] = item.ObjectType
	networkfczoneinforelationship["selector"] = item.Selector

	networkfczoneinforelationships = append(networkfczoneinforelationships, networkfczoneinforelationship)
	return networkfczoneinforelationships
}
func flattenMapNetworkVlanPortInfoRelationship(p models.NetworkVlanPortInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var networkvlanportinforelationships []map[string]interface{}
	var ret models.NetworkVlanPortInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	networkvlanportinforelationship := make(map[string]interface{})
	networkvlanportinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	networkvlanportinforelationship["class_id"] = item.ClassId
	networkvlanportinforelationship["moid"] = item.Moid
	networkvlanportinforelationship["object_type"] = item.ObjectType
	networkvlanportinforelationship["selector"] = item.Selector

	networkvlanportinforelationships = append(networkvlanportinforelationships, networkvlanportinforelationship)
	return networkvlanportinforelationships
}
func flattenMapNiaapiNewReleaseDetail(p models.NiaapiNewReleaseDetail, d *schema.ResourceData) []map[string]interface{} {
	var niaapinewreleasedetails []map[string]interface{}
	var ret models.NiaapiNewReleaseDetail
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niaapinewreleasedetail := make(map[string]interface{})
	niaapinewreleasedetail["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niaapinewreleasedetail["class_id"] = item.ClassId
	niaapinewreleasedetail["description"] = item.Description
	niaapinewreleasedetail["link"] = item.Link
	niaapinewreleasedetail["object_type"] = item.ObjectType
	niaapinewreleasedetail["release_note_link"] = item.ReleaseNoteLink
	niaapinewreleasedetail["release_note_link_title"] = item.ReleaseNoteLinkTitle
	niaapinewreleasedetail["software_download_link"] = item.SoftwareDownloadLink
	niaapinewreleasedetail["software_download_link_title"] = item.SoftwareDownloadLinkTitle
	niaapinewreleasedetail["title"] = item.Title
	niaapinewreleasedetail["nr_version"] = item.Version

	niaapinewreleasedetails = append(niaapinewreleasedetails, niaapinewreleasedetail)
	return niaapinewreleasedetails
}
func flattenMapNiaapiVersionRegexPlatform(p models.NiaapiVersionRegexPlatform, d *schema.ResourceData) []map[string]interface{} {
	var niaapiversionregexplatforms []map[string]interface{}
	var ret models.NiaapiVersionRegexPlatform
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niaapiversionregexplatform := make(map[string]interface{})
	niaapiversionregexplatform["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niaapiversionregexplatform["anyllregex"] = item.Anyllregex
	niaapiversionregexplatform["class_id"] = item.ClassId
	niaapiversionregexplatform["currentlltrain"] = (func(p models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		var ret models.NiaapiSoftwareRegex
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapisoftwareregex["class_id"] = item.ClassId
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetCurrentlltrain(), d)
	niaapiversionregexplatform["latestsltrain"] = (func(p models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		var ret models.NiaapiSoftwareRegex
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapisoftwareregex["class_id"] = item.ClassId
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetLatestsltrain(), d)
	niaapiversionregexplatform["object_type"] = item.ObjectType
	niaapiversionregexplatform["sltrain"] = (func(p []models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			niaapisoftwareregex := make(map[string]interface{})
			niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			niaapisoftwareregex["class_id"] = item.ClassId
			niaapisoftwareregex["object_type"] = item.ObjectType
			niaapisoftwareregex["regex"] = item.Regex
			niaapisoftwareregex["software_version"] = item.SoftwareVersion
			niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		}
		return niaapisoftwareregexs
	})(item.GetSltrain(), d)
	niaapiversionregexplatform["upcominglltrain"] = (func(p models.NiaapiSoftwareRegex, d *schema.ResourceData) []map[string]interface{} {
		var niaapisoftwareregexs []map[string]interface{}
		var ret models.NiaapiSoftwareRegex
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		niaapisoftwareregex := make(map[string]interface{})
		niaapisoftwareregex["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		niaapisoftwareregex["class_id"] = item.ClassId
		niaapisoftwareregex["object_type"] = item.ObjectType
		niaapisoftwareregex["regex"] = item.Regex
		niaapisoftwareregex["software_version"] = item.SoftwareVersion

		niaapisoftwareregexs = append(niaapisoftwareregexs, niaapisoftwareregex)
		return niaapisoftwareregexs
	})(item.GetUpcominglltrain(), d)

	niaapiversionregexplatforms = append(niaapiversionregexplatforms, niaapiversionregexplatform)
	return niaapiversionregexplatforms
}
func flattenMapNiatelemetryDiskinfo(p models.NiatelemetryDiskinfo, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrydiskinfos []map[string]interface{}
	var ret models.NiatelemetryDiskinfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	niatelemetrydiskinfo := make(map[string]interface{})
	niatelemetrydiskinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrydiskinfo["class_id"] = item.ClassId
	niatelemetrydiskinfo["free"] = item.Free
	niatelemetrydiskinfo["name"] = item.Name
	niatelemetrydiskinfo["object_type"] = item.ObjectType
	niatelemetrydiskinfo["total"] = item.Total
	niatelemetrydiskinfo["used"] = item.Used

	niatelemetrydiskinfos = append(niatelemetrydiskinfos, niatelemetrydiskinfo)
	return niatelemetrydiskinfos
}
func flattenMapNiatelemetryNiaInventoryRelationship(p models.NiatelemetryNiaInventoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetryniainventoryrelationships []map[string]interface{}
	var ret models.NiatelemetryNiaInventoryRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	niatelemetryniainventoryrelationship := make(map[string]interface{})
	niatelemetryniainventoryrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetryniainventoryrelationship["class_id"] = item.ClassId
	niatelemetryniainventoryrelationship["moid"] = item.Moid
	niatelemetryniainventoryrelationship["object_type"] = item.ObjectType
	niatelemetryniainventoryrelationship["selector"] = item.Selector

	niatelemetryniainventoryrelationships = append(niatelemetryniainventoryrelationships, niatelemetryniainventoryrelationship)
	return niatelemetryniainventoryrelationships
}
func flattenMapNiatelemetryNiaLicenseStateRelationship(p models.NiatelemetryNiaLicenseStateRelationship, d *schema.ResourceData) []map[string]interface{} {
	var niatelemetrynialicensestaterelationships []map[string]interface{}
	var ret models.NiatelemetryNiaLicenseStateRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	niatelemetrynialicensestaterelationship := make(map[string]interface{})
	niatelemetrynialicensestaterelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	niatelemetrynialicensestaterelationship["class_id"] = item.ClassId
	niatelemetrynialicensestaterelationship["moid"] = item.Moid
	niatelemetrynialicensestaterelationship["object_type"] = item.ObjectType
	niatelemetrynialicensestaterelationship["selector"] = item.Selector

	niatelemetrynialicensestaterelationships = append(niatelemetrynialicensestaterelationships, niatelemetrynialicensestaterelationship)
	return niatelemetrynialicensestaterelationships
}
func flattenMapOnpremSchedule(p models.OnpremSchedule, d *schema.ResourceData) []map[string]interface{} {
	var onpremschedules []map[string]interface{}
	var ret models.OnpremSchedule
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	onpremschedule := make(map[string]interface{})
	onpremschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	onpremschedule["class_id"] = item.ClassId
	onpremschedule["day_of_month"] = item.DayOfMonth
	onpremschedule["day_of_week"] = item.DayOfWeek
	onpremschedule["month_of_year"] = item.MonthOfYear
	onpremschedule["object_type"] = item.ObjectType
	onpremschedule["repeat_interval"] = item.RepeatInterval
	onpremschedule["time_of_day"] = item.TimeOfDay
	onpremschedule["time_zone"] = item.TimeZone
	onpremschedule["week_of_month"] = item.WeekOfMonth

	onpremschedules = append(onpremschedules, onpremschedule)
	return onpremschedules
}
func flattenMapOnpremUpgradePhase(p models.OnpremUpgradePhase, d *schema.ResourceData) []map[string]interface{} {
	var onpremupgradephases []map[string]interface{}
	var ret models.OnpremUpgradePhase
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	onpremupgradephase := make(map[string]interface{})
	onpremupgradephase["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	onpremupgradephase["class_id"] = item.ClassId
	onpremupgradephase["elapsed_time"] = item.ElapsedTime
	onpremupgradephase["end_time"] = item.EndTime
	onpremupgradephase["failed"] = item.Failed
	onpremupgradephase["message"] = item.Message
	onpremupgradephase["name"] = item.Name
	onpremupgradephase["object_type"] = item.ObjectType
	onpremupgradephase["start_time"] = item.StartTime

	onpremupgradephases = append(onpremupgradephases, onpremupgradephase)
	return onpremupgradephases
}
func flattenMapOrganizationOrganizationRelationship(p models.OrganizationOrganizationRelationship, d *schema.ResourceData) []map[string]interface{} {
	var organizationorganizationrelationships []map[string]interface{}
	var ret models.OrganizationOrganizationRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	organizationorganizationrelationship := make(map[string]interface{})
	organizationorganizationrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	organizationorganizationrelationship["class_id"] = item.ClassId
	organizationorganizationrelationship["moid"] = item.Moid
	organizationorganizationrelationship["object_type"] = item.ObjectType
	organizationorganizationrelationship["selector"] = item.Selector

	organizationorganizationrelationships = append(organizationorganizationrelationships, organizationorganizationrelationship)
	return organizationorganizationrelationships
}
func flattenMapOsAnswers(p models.OsAnswers, d *schema.ResourceData) []map[string]interface{} {
	var osanswerss []map[string]interface{}
	var ret models.OsAnswers
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	osanswers := make(map[string]interface{})
	osanswers["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	answer_file_x := d.Get("answers").([]interface{})
	answer_file_y := answer_file_x[0].(map[string]interface{})
	osanswers["answer_file"] = answer_file_y["answer_file"]
	osanswers["class_id"] = item.ClassId
	osanswers["hostname"] = item.Hostname
	osanswers["ip_config_type"] = item.IpConfigType
	osanswers["ip_configuration"] = (func(p models.OsIpConfiguration, d *schema.ResourceData) []map[string]interface{} {
		var osipconfigurations []map[string]interface{}
		var ret models.OsIpConfiguration
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		osipconfiguration := make(map[string]interface{})
		osipconfiguration["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		osipconfiguration["class_id"] = item.ClassId
		osipconfiguration["object_type"] = item.ObjectType

		osipconfigurations = append(osipconfigurations, osipconfiguration)
		return osipconfigurations
	})(item.GetIpConfiguration(), d)
	osanswers["is_answer_file_set"] = item.IsAnswerFileSet
	osanswers["is_root_password_crypted"] = item.IsRootPasswordCrypted
	osanswers["is_root_password_set"] = item.IsRootPasswordSet
	osanswers["nameserver"] = item.Nameserver
	osanswers["object_type"] = item.ObjectType
	osanswers["product_key"] = item.ProductKey
	root_password_x := d.Get("answers").([]interface{})
	root_password_y := root_password_x[0].(map[string]interface{})
	osanswers["root_password"] = root_password_y["root_password"]
	osanswers["nr_source"] = item.Source

	osanswerss = append(osanswerss, osanswers)
	return osanswerss
}
func flattenMapOsCatalogRelationship(p models.OsCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var oscatalogrelationships []map[string]interface{}
	var ret models.OsCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	oscatalogrelationship := make(map[string]interface{})
	oscatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	oscatalogrelationship["class_id"] = item.ClassId
	oscatalogrelationship["moid"] = item.Moid
	oscatalogrelationship["object_type"] = item.ObjectType
	oscatalogrelationship["selector"] = item.Selector

	oscatalogrelationships = append(oscatalogrelationships, oscatalogrelationship)
	return oscatalogrelationships
}
func flattenMapOsConfigurationFileRelationship(p models.OsConfigurationFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var osconfigurationfilerelationships []map[string]interface{}
	var ret models.OsConfigurationFileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	osconfigurationfilerelationship := make(map[string]interface{})
	osconfigurationfilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	osconfigurationfilerelationship["class_id"] = item.ClassId
	osconfigurationfilerelationship["moid"] = item.Moid
	osconfigurationfilerelationship["object_type"] = item.ObjectType
	osconfigurationfilerelationship["selector"] = item.Selector

	osconfigurationfilerelationships = append(osconfigurationfilerelationships, osconfigurationfilerelationship)
	return osconfigurationfilerelationships
}
func flattenMapOsOperatingSystemParameters(p models.OsOperatingSystemParameters, d *schema.ResourceData) []map[string]interface{} {
	var osoperatingsystemparameterss []map[string]interface{}
	var ret models.OsOperatingSystemParameters
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	osoperatingsystemparameters := make(map[string]interface{})
	osoperatingsystemparameters["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	osoperatingsystemparameters["class_id"] = item.ClassId
	osoperatingsystemparameters["object_type"] = item.ObjectType

	osoperatingsystemparameterss = append(osoperatingsystemparameterss, osoperatingsystemparameters)
	return osoperatingsystemparameterss
}
func flattenMapPciSwitchRelationship(p models.PciSwitchRelationship, d *schema.ResourceData) []map[string]interface{} {
	var pciswitchrelationships []map[string]interface{}
	var ret models.PciSwitchRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	pciswitchrelationship := make(map[string]interface{})
	pciswitchrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pciswitchrelationship["class_id"] = item.ClassId
	pciswitchrelationship["moid"] = item.Moid
	pciswitchrelationship["object_type"] = item.ObjectType
	pciswitchrelationship["selector"] = item.Selector

	pciswitchrelationships = append(pciswitchrelationships, pciswitchrelationship)
	return pciswitchrelationships
}
func flattenMapPkixDistinguishedName(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
	var pkixdistinguishednames []map[string]interface{}
	var ret models.PkixDistinguishedName
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	pkixdistinguishedname := make(map[string]interface{})
	pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pkixdistinguishedname["class_id"] = item.ClassId
	pkixdistinguishedname["common_name"] = item.CommonName
	pkixdistinguishedname["country"] = item.Country
	pkixdistinguishedname["locality"] = item.Locality
	pkixdistinguishedname["object_type"] = item.ObjectType
	pkixdistinguishedname["organization"] = item.Organization
	pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
	pkixdistinguishedname["state"] = item.State

	pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
	return pkixdistinguishednames
}
func flattenMapPkixKeyGenerationSpec(p models.PkixKeyGenerationSpec, d *schema.ResourceData) []map[string]interface{} {
	var pkixkeygenerationspecs []map[string]interface{}
	var ret models.PkixKeyGenerationSpec
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	pkixkeygenerationspec := make(map[string]interface{})
	pkixkeygenerationspec["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pkixkeygenerationspec["class_id"] = item.ClassId
	pkixkeygenerationspec["name"] = item.Name
	pkixkeygenerationspec["object_type"] = item.ObjectType

	pkixkeygenerationspecs = append(pkixkeygenerationspecs, pkixkeygenerationspec)
	return pkixkeygenerationspecs
}
func flattenMapPkixSubjectAlternateName(p models.PkixSubjectAlternateName, d *schema.ResourceData) []map[string]interface{} {
	var pkixsubjectalternatenames []map[string]interface{}
	var ret models.PkixSubjectAlternateName
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	pkixsubjectalternatename := make(map[string]interface{})
	pkixsubjectalternatename["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	pkixsubjectalternatename["class_id"] = item.ClassId
	pkixsubjectalternatename["dns_name"] = item.DnsName
	pkixsubjectalternatename["email_address"] = item.EmailAddress
	pkixsubjectalternatename["ip_address"] = item.IpAddress
	pkixsubjectalternatename["object_type"] = item.ObjectType
	pkixsubjectalternatename["uri"] = item.Uri

	pkixsubjectalternatenames = append(pkixsubjectalternatenames, pkixsubjectalternatename)
	return pkixsubjectalternatenames
}
func flattenMapPolicyAbstractConfigProfileRelationship(p models.PolicyAbstractConfigProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractconfigprofilerelationships []map[string]interface{}
	var ret models.PolicyAbstractConfigProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	policyabstractconfigprofilerelationship := make(map[string]interface{})
	policyabstractconfigprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyabstractconfigprofilerelationship["class_id"] = item.ClassId
	policyabstractconfigprofilerelationship["moid"] = item.Moid
	policyabstractconfigprofilerelationship["object_type"] = item.ObjectType
	policyabstractconfigprofilerelationship["selector"] = item.Selector

	policyabstractconfigprofilerelationships = append(policyabstractconfigprofilerelationships, policyabstractconfigprofilerelationship)
	return policyabstractconfigprofilerelationships
}
func flattenMapPolicyAbstractProfileRelationship(p models.PolicyAbstractProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var policyabstractprofilerelationships []map[string]interface{}
	var ret models.PolicyAbstractProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	policyabstractprofilerelationship := make(map[string]interface{})
	policyabstractprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyabstractprofilerelationship["class_id"] = item.ClassId
	policyabstractprofilerelationship["moid"] = item.Moid
	policyabstractprofilerelationship["object_type"] = item.ObjectType
	policyabstractprofilerelationship["selector"] = item.Selector

	policyabstractprofilerelationships = append(policyabstractprofilerelationships, policyabstractprofilerelationship)
	return policyabstractprofilerelationships
}
func flattenMapPolicyConfigChange(p models.PolicyConfigChange, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigchanges []map[string]interface{}
	var ret models.PolicyConfigChange
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	policyconfigchange := make(map[string]interface{})
	policyconfigchange["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyconfigchange["changes"] = item.Changes
	policyconfigchange["class_id"] = item.ClassId
	policyconfigchange["disruptions"] = item.Disruptions
	policyconfigchange["object_type"] = item.ObjectType

	policyconfigchanges = append(policyconfigchanges, policyconfigchange)
	return policyconfigchanges
}
func flattenMapPolicyConfigContext(p models.PolicyConfigContext, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigcontexts []map[string]interface{}
	var ret models.PolicyConfigContext
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	policyconfigcontext := make(map[string]interface{})
	policyconfigcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyconfigcontext["class_id"] = item.ClassId
	policyconfigcontext["config_state"] = item.ConfigState
	policyconfigcontext["control_action"] = item.ControlAction
	policyconfigcontext["error_state"] = item.ErrorState
	policyconfigcontext["object_type"] = item.ObjectType
	policyconfigcontext["oper_state"] = item.OperState

	policyconfigcontexts = append(policyconfigcontexts, policyconfigcontext)
	return policyconfigcontexts
}
func flattenMapPolicyConfigResultContext(p models.PolicyConfigResultContext, d *schema.ResourceData) []map[string]interface{} {
	var policyconfigresultcontexts []map[string]interface{}
	var ret models.PolicyConfigResultContext
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	policyconfigresultcontext := make(map[string]interface{})
	policyconfigresultcontext["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	policyconfigresultcontext["class_id"] = item.ClassId
	policyconfigresultcontext["entity_data"] = item.EntityData
	policyconfigresultcontext["entity_moid"] = item.EntityMoid
	policyconfigresultcontext["entity_name"] = item.EntityName
	policyconfigresultcontext["entity_type"] = item.EntityType
	policyconfigresultcontext["object_type"] = item.ObjectType

	policyconfigresultcontexts = append(policyconfigresultcontexts, policyconfigresultcontext)
	return policyconfigresultcontexts
}
func flattenMapPortGroupRelationship(p models.PortGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portgrouprelationships []map[string]interface{}
	var ret models.PortGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	portgrouprelationship := make(map[string]interface{})
	portgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	portgrouprelationship["class_id"] = item.ClassId
	portgrouprelationship["moid"] = item.Moid
	portgrouprelationship["object_type"] = item.ObjectType
	portgrouprelationship["selector"] = item.Selector

	portgrouprelationships = append(portgrouprelationships, portgrouprelationship)
	return portgrouprelationships
}
func flattenMapPortInterfaceBaseRelationship(p models.PortInterfaceBaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portinterfacebaserelationships []map[string]interface{}
	var ret models.PortInterfaceBaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	portinterfacebaserelationship := make(map[string]interface{})
	portinterfacebaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	portinterfacebaserelationship["class_id"] = item.ClassId
	portinterfacebaserelationship["moid"] = item.Moid
	portinterfacebaserelationship["object_type"] = item.ObjectType
	portinterfacebaserelationship["selector"] = item.Selector

	portinterfacebaserelationships = append(portinterfacebaserelationships, portinterfacebaserelationship)
	return portinterfacebaserelationships
}
func flattenMapPortSubGroupRelationship(p models.PortSubGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var portsubgrouprelationships []map[string]interface{}
	var ret models.PortSubGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	portsubgrouprelationship := make(map[string]interface{})
	portsubgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	portsubgrouprelationship["class_id"] = item.ClassId
	portsubgrouprelationship["moid"] = item.Moid
	portsubgrouprelationship["object_type"] = item.ObjectType
	portsubgrouprelationship["selector"] = item.Selector

	portsubgrouprelationships = append(portsubgrouprelationships, portsubgrouprelationship)
	return portsubgrouprelationships
}
func flattenMapRecoveryAbstractBackupInfoRelationship(p models.RecoveryAbstractBackupInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryabstractbackupinforelationships []map[string]interface{}
	var ret models.RecoveryAbstractBackupInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoveryabstractbackupinforelationship := make(map[string]interface{})
	recoveryabstractbackupinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryabstractbackupinforelationship["class_id"] = item.ClassId
	recoveryabstractbackupinforelationship["moid"] = item.Moid
	recoveryabstractbackupinforelationship["object_type"] = item.ObjectType
	recoveryabstractbackupinforelationship["selector"] = item.Selector

	recoveryabstractbackupinforelationships = append(recoveryabstractbackupinforelationships, recoveryabstractbackupinforelationship)
	return recoveryabstractbackupinforelationships
}
func flattenMapRecoveryBackupConfigPolicyRelationship(p models.RecoveryBackupConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupconfigpolicyrelationships []map[string]interface{}
	var ret models.RecoveryBackupConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoverybackupconfigpolicyrelationship := make(map[string]interface{})
	recoverybackupconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoverybackupconfigpolicyrelationship["class_id"] = item.ClassId
	recoverybackupconfigpolicyrelationship["moid"] = item.Moid
	recoverybackupconfigpolicyrelationship["object_type"] = item.ObjectType
	recoverybackupconfigpolicyrelationship["selector"] = item.Selector

	recoverybackupconfigpolicyrelationships = append(recoverybackupconfigpolicyrelationships, recoverybackupconfigpolicyrelationship)
	return recoverybackupconfigpolicyrelationships
}
func flattenMapRecoveryBackupProfileRelationship(p models.RecoveryBackupProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupprofilerelationships []map[string]interface{}
	var ret models.RecoveryBackupProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoverybackupprofilerelationship := make(map[string]interface{})
	recoverybackupprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoverybackupprofilerelationship["class_id"] = item.ClassId
	recoverybackupprofilerelationship["moid"] = item.Moid
	recoverybackupprofilerelationship["object_type"] = item.ObjectType
	recoverybackupprofilerelationship["selector"] = item.Selector

	recoverybackupprofilerelationships = append(recoverybackupprofilerelationships, recoverybackupprofilerelationship)
	return recoverybackupprofilerelationships
}
func flattenMapRecoveryBackupSchedule(p models.RecoveryBackupSchedule, d *schema.ResourceData) []map[string]interface{} {
	var recoverybackupschedules []map[string]interface{}
	var ret models.RecoveryBackupSchedule
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	recoverybackupschedule := make(map[string]interface{})
	recoverybackupschedule["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoverybackupschedule["class_id"] = item.ClassId
	recoverybackupschedule["execution_time"] = item.ExecutionTime
	recoverybackupschedule["frequency_unit"] = item.FrequencyUnit
	recoverybackupschedule["hours"] = item.Hours
	recoverybackupschedule["object_type"] = item.ObjectType

	recoverybackupschedules = append(recoverybackupschedules, recoverybackupschedule)
	return recoverybackupschedules
}
func flattenMapRecoveryConfigParams(p models.RecoveryConfigParams, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigparamss []map[string]interface{}
	var ret models.RecoveryConfigParams
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	recoveryconfigparams := make(map[string]interface{})
	recoveryconfigparams["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryconfigparams["class_id"] = item.ClassId
	recoveryconfigparams["object_type"] = item.ObjectType

	recoveryconfigparamss = append(recoveryconfigparamss, recoveryconfigparams)
	return recoveryconfigparamss
}
func flattenMapRecoveryConfigResultRelationship(p models.RecoveryConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryconfigresultrelationships []map[string]interface{}
	var ret models.RecoveryConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoveryconfigresultrelationship := make(map[string]interface{})
	recoveryconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryconfigresultrelationship["class_id"] = item.ClassId
	recoveryconfigresultrelationship["moid"] = item.Moid
	recoveryconfigresultrelationship["object_type"] = item.ObjectType
	recoveryconfigresultrelationship["selector"] = item.Selector

	recoveryconfigresultrelationships = append(recoveryconfigresultrelationships, recoveryconfigresultrelationship)
	return recoveryconfigresultrelationships
}
func flattenMapRecoveryScheduleConfigPolicyRelationship(p models.RecoveryScheduleConfigPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var recoveryscheduleconfigpolicyrelationships []map[string]interface{}
	var ret models.RecoveryScheduleConfigPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	recoveryscheduleconfigpolicyrelationship := make(map[string]interface{})
	recoveryscheduleconfigpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	recoveryscheduleconfigpolicyrelationship["class_id"] = item.ClassId
	recoveryscheduleconfigpolicyrelationship["moid"] = item.Moid
	recoveryscheduleconfigpolicyrelationship["object_type"] = item.ObjectType
	recoveryscheduleconfigpolicyrelationship["selector"] = item.Selector

	recoveryscheduleconfigpolicyrelationships = append(recoveryscheduleconfigpolicyrelationships, recoveryscheduleconfigpolicyrelationship)
	return recoveryscheduleconfigpolicyrelationships
}
func flattenMapResourceGroupRelationship(p models.ResourceGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcegrouprelationships []map[string]interface{}
	var ret models.ResourceGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	resourcegrouprelationship := make(map[string]interface{})
	resourcegrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	resourcegrouprelationship["class_id"] = item.ClassId
	resourcegrouprelationship["moid"] = item.Moid
	resourcegrouprelationship["object_type"] = item.ObjectType
	resourcegrouprelationship["selector"] = item.Selector

	resourcegrouprelationships = append(resourcegrouprelationships, resourcegrouprelationship)
	return resourcegrouprelationships
}
func flattenMapResourceMembershipHolderRelationship(p models.ResourceMembershipHolderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var resourcemembershipholderrelationships []map[string]interface{}
	var ret models.ResourceMembershipHolderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	resourcemembershipholderrelationship := make(map[string]interface{})
	resourcemembershipholderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	resourcemembershipholderrelationship["class_id"] = item.ClassId
	resourcemembershipholderrelationship["moid"] = item.Moid
	resourcemembershipholderrelationship["object_type"] = item.ObjectType
	resourcemembershipholderrelationship["selector"] = item.Selector

	resourcemembershipholderrelationships = append(resourcemembershipholderrelationships, resourcemembershipholderrelationship)
	return resourcemembershipholderrelationships
}
func flattenMapSdwanProfileRelationship(p models.SdwanProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanprofilerelationships []map[string]interface{}
	var ret models.SdwanProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	sdwanprofilerelationship := make(map[string]interface{})
	sdwanprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	sdwanprofilerelationship["class_id"] = item.ClassId
	sdwanprofilerelationship["moid"] = item.Moid
	sdwanprofilerelationship["object_type"] = item.ObjectType
	sdwanprofilerelationship["selector"] = item.Selector

	sdwanprofilerelationships = append(sdwanprofilerelationships, sdwanprofilerelationship)
	return sdwanprofilerelationships
}
func flattenMapSdwanRouterPolicyRelationship(p models.SdwanRouterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanrouterpolicyrelationships []map[string]interface{}
	var ret models.SdwanRouterPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	sdwanrouterpolicyrelationship := make(map[string]interface{})
	sdwanrouterpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	sdwanrouterpolicyrelationship["class_id"] = item.ClassId
	sdwanrouterpolicyrelationship["moid"] = item.Moid
	sdwanrouterpolicyrelationship["object_type"] = item.ObjectType
	sdwanrouterpolicyrelationship["selector"] = item.Selector

	sdwanrouterpolicyrelationships = append(sdwanrouterpolicyrelationships, sdwanrouterpolicyrelationship)
	return sdwanrouterpolicyrelationships
}
func flattenMapSdwanVmanageAccountPolicyRelationship(p models.SdwanVmanageAccountPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var sdwanvmanageaccountpolicyrelationships []map[string]interface{}
	var ret models.SdwanVmanageAccountPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	sdwanvmanageaccountpolicyrelationship := make(map[string]interface{})
	sdwanvmanageaccountpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	sdwanvmanageaccountpolicyrelationship["class_id"] = item.ClassId
	sdwanvmanageaccountpolicyrelationship["moid"] = item.Moid
	sdwanvmanageaccountpolicyrelationship["object_type"] = item.ObjectType
	sdwanvmanageaccountpolicyrelationship["selector"] = item.Selector

	sdwanvmanageaccountpolicyrelationships = append(sdwanvmanageaccountpolicyrelationships, sdwanvmanageaccountpolicyrelationship)
	return sdwanvmanageaccountpolicyrelationships
}
func flattenMapServerConfigResultRelationship(p models.ServerConfigResultRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverconfigresultrelationships []map[string]interface{}
	var ret models.ServerConfigResultRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	serverconfigresultrelationship := make(map[string]interface{})
	serverconfigresultrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	serverconfigresultrelationship["class_id"] = item.ClassId
	serverconfigresultrelationship["moid"] = item.Moid
	serverconfigresultrelationship["object_type"] = item.ObjectType
	serverconfigresultrelationship["selector"] = item.Selector

	serverconfigresultrelationships = append(serverconfigresultrelationships, serverconfigresultrelationship)
	return serverconfigresultrelationships
}
func flattenMapServerProfileRelationship(p models.ServerProfileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var serverprofilerelationships []map[string]interface{}
	var ret models.ServerProfileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	serverprofilerelationship := make(map[string]interface{})
	serverprofilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	serverprofilerelationship["class_id"] = item.ClassId
	serverprofilerelationship["moid"] = item.Moid
	serverprofilerelationship["object_type"] = item.ObjectType
	serverprofilerelationship["selector"] = item.Selector

	serverprofilerelationships = append(serverprofilerelationships, serverprofilerelationship)
	return serverprofilerelationships
}
func flattenMapSoftwareHyperflexDistributableRelationship(p models.SoftwareHyperflexDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarehyperflexdistributablerelationships []map[string]interface{}
	var ret models.SoftwareHyperflexDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarehyperflexdistributablerelationship := make(map[string]interface{})
	softwarehyperflexdistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarehyperflexdistributablerelationship["class_id"] = item.ClassId
	softwarehyperflexdistributablerelationship["moid"] = item.Moid
	softwarehyperflexdistributablerelationship["object_type"] = item.ObjectType
	softwarehyperflexdistributablerelationship["selector"] = item.Selector

	softwarehyperflexdistributablerelationships = append(softwarehyperflexdistributablerelationships, softwarehyperflexdistributablerelationship)
	return softwarehyperflexdistributablerelationships
}
func flattenMapSoftwareSolutionDistributableRelationship(p models.SoftwareSolutionDistributableRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwaresolutiondistributablerelationships []map[string]interface{}
	var ret models.SoftwareSolutionDistributableRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwaresolutiondistributablerelationship := make(map[string]interface{})
	softwaresolutiondistributablerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwaresolutiondistributablerelationship["class_id"] = item.ClassId
	softwaresolutiondistributablerelationship["moid"] = item.Moid
	softwaresolutiondistributablerelationship["object_type"] = item.ObjectType
	softwaresolutiondistributablerelationship["selector"] = item.Selector

	softwaresolutiondistributablerelationships = append(softwaresolutiondistributablerelationships, softwaresolutiondistributablerelationship)
	return softwaresolutiondistributablerelationships
}
func flattenMapSoftwarerepositoryCatalogRelationship(p models.SoftwarerepositoryCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositorycatalogrelationships []map[string]interface{}
	var ret models.SoftwarerepositoryCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositorycatalogrelationship := make(map[string]interface{})
	softwarerepositorycatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositorycatalogrelationship["class_id"] = item.ClassId
	softwarerepositorycatalogrelationship["moid"] = item.Moid
	softwarerepositorycatalogrelationship["object_type"] = item.ObjectType
	softwarerepositorycatalogrelationship["selector"] = item.Selector

	softwarerepositorycatalogrelationships = append(softwarerepositorycatalogrelationships, softwarerepositorycatalogrelationship)
	return softwarerepositorycatalogrelationships
}
func flattenMapSoftwarerepositoryFileRelationship(p models.SoftwarerepositoryFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryfilerelationships []map[string]interface{}
	var ret models.SoftwarerepositoryFileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositoryfilerelationship := make(map[string]interface{})
	softwarerepositoryfilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryfilerelationship["class_id"] = item.ClassId
	softwarerepositoryfilerelationship["moid"] = item.Moid
	softwarerepositoryfilerelationship["object_type"] = item.ObjectType
	softwarerepositoryfilerelationship["selector"] = item.Selector

	softwarerepositoryfilerelationships = append(softwarerepositoryfilerelationships, softwarerepositoryfilerelationship)
	return softwarerepositoryfilerelationships
}
func flattenMapSoftwarerepositoryFileServer(p models.SoftwarerepositoryFileServer, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryfileservers []map[string]interface{}
	var ret models.SoftwarerepositoryFileServer
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	softwarerepositoryfileserver := make(map[string]interface{})
	softwarerepositoryfileserver["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryfileserver["class_id"] = item.ClassId
	softwarerepositoryfileserver["object_type"] = item.ObjectType

	softwarerepositoryfileservers = append(softwarerepositoryfileservers, softwarerepositoryfileserver)
	return softwarerepositoryfileservers
}
func flattenMapSoftwarerepositoryOperatingSystemFileRelationship(p models.SoftwarerepositoryOperatingSystemFileRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryoperatingsystemfilerelationships []map[string]interface{}
	var ret models.SoftwarerepositoryOperatingSystemFileRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositoryoperatingsystemfilerelationship := make(map[string]interface{})
	softwarerepositoryoperatingsystemfilerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryoperatingsystemfilerelationship["class_id"] = item.ClassId
	softwarerepositoryoperatingsystemfilerelationship["moid"] = item.Moid
	softwarerepositoryoperatingsystemfilerelationship["object_type"] = item.ObjectType
	softwarerepositoryoperatingsystemfilerelationship["selector"] = item.Selector

	softwarerepositoryoperatingsystemfilerelationships = append(softwarerepositoryoperatingsystemfilerelationships, softwarerepositoryoperatingsystemfilerelationship)
	return softwarerepositoryoperatingsystemfilerelationships
}
func flattenMapSoftwarerepositoryReleaseRelationship(p models.SoftwarerepositoryReleaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var softwarerepositoryreleaserelationships []map[string]interface{}
	var ret models.SoftwarerepositoryReleaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	softwarerepositoryreleaserelationship := make(map[string]interface{})
	softwarerepositoryreleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	softwarerepositoryreleaserelationship["class_id"] = item.ClassId
	softwarerepositoryreleaserelationship["moid"] = item.Moid
	softwarerepositoryreleaserelationship["object_type"] = item.ObjectType
	softwarerepositoryreleaserelationship["selector"] = item.Selector

	softwarerepositoryreleaserelationships = append(softwarerepositoryreleaserelationships, softwarerepositoryreleaserelationship)
	return softwarerepositoryreleaserelationships
}
func flattenMapStorageBaseCapacity(p models.StorageBaseCapacity, d *schema.ResourceData) []map[string]interface{} {
	var storagebasecapacitys []map[string]interface{}
	var ret models.StorageBaseCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	storagebasecapacity := make(map[string]interface{})
	storagebasecapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagebasecapacity["available"] = item.Available
	storagebasecapacity["class_id"] = item.ClassId
	storagebasecapacity["free"] = item.Free
	storagebasecapacity["object_type"] = item.ObjectType
	storagebasecapacity["total"] = item.Total
	storagebasecapacity["used"] = item.Used

	storagebasecapacitys = append(storagebasecapacitys, storagebasecapacity)
	return storagebasecapacitys
}
func flattenMapStorageControllerRelationship(p models.StorageControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagecontrollerrelationships []map[string]interface{}
	var ret models.StorageControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagecontrollerrelationship := make(map[string]interface{})
	storagecontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagecontrollerrelationship["class_id"] = item.ClassId
	storagecontrollerrelationship["moid"] = item.Moid
	storagecontrollerrelationship["object_type"] = item.ObjectType
	storagecontrollerrelationship["selector"] = item.Selector

	storagecontrollerrelationships = append(storagecontrollerrelationships, storagecontrollerrelationship)
	return storagecontrollerrelationships
}
func flattenMapStorageDiskGroupRelationship(p models.StorageDiskGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagediskgrouprelationships []map[string]interface{}
	var ret models.StorageDiskGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagediskgrouprelationship := make(map[string]interface{})
	storagediskgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagediskgrouprelationship["class_id"] = item.ClassId
	storagediskgrouprelationship["moid"] = item.Moid
	storagediskgrouprelationship["object_type"] = item.ObjectType
	storagediskgrouprelationship["selector"] = item.Selector

	storagediskgrouprelationships = append(storagediskgrouprelationships, storagediskgrouprelationship)
	return storagediskgrouprelationships
}
func flattenMapStorageEnclosureRelationship(p models.StorageEnclosureRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageenclosurerelationships []map[string]interface{}
	var ret models.StorageEnclosureRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storageenclosurerelationship := make(map[string]interface{})
	storageenclosurerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storageenclosurerelationship["class_id"] = item.ClassId
	storageenclosurerelationship["moid"] = item.Moid
	storageenclosurerelationship["object_type"] = item.ObjectType
	storageenclosurerelationship["selector"] = item.Selector

	storageenclosurerelationships = append(storageenclosurerelationships, storageenclosurerelationship)
	return storageenclosurerelationships
}
func flattenMapStorageFlexFlashControllerRelationship(p models.StorageFlexFlashControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexflashcontrollerrelationships []map[string]interface{}
	var ret models.StorageFlexFlashControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storageflexflashcontrollerrelationship := make(map[string]interface{})
	storageflexflashcontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storageflexflashcontrollerrelationship["class_id"] = item.ClassId
	storageflexflashcontrollerrelationship["moid"] = item.Moid
	storageflexflashcontrollerrelationship["object_type"] = item.ObjectType
	storageflexflashcontrollerrelationship["selector"] = item.Selector

	storageflexflashcontrollerrelationships = append(storageflexflashcontrollerrelationships, storageflexflashcontrollerrelationship)
	return storageflexflashcontrollerrelationships
}
func flattenMapStorageFlexUtilControllerRelationship(p models.StorageFlexUtilControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storageflexutilcontrollerrelationships []map[string]interface{}
	var ret models.StorageFlexUtilControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storageflexutilcontrollerrelationship := make(map[string]interface{})
	storageflexutilcontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storageflexutilcontrollerrelationship["class_id"] = item.ClassId
	storageflexutilcontrollerrelationship["moid"] = item.Moid
	storageflexutilcontrollerrelationship["object_type"] = item.ObjectType
	storageflexutilcontrollerrelationship["selector"] = item.Selector

	storageflexutilcontrollerrelationships = append(storageflexutilcontrollerrelationships, storageflexutilcontrollerrelationship)
	return storageflexutilcontrollerrelationships
}
func flattenMapStoragePhysicalDiskRelationship(p models.StoragePhysicalDiskRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagephysicaldiskrelationships []map[string]interface{}
	var ret models.StoragePhysicalDiskRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagephysicaldiskrelationship := make(map[string]interface{})
	storagephysicaldiskrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagephysicaldiskrelationship["class_id"] = item.ClassId
	storagephysicaldiskrelationship["moid"] = item.Moid
	storagephysicaldiskrelationship["object_type"] = item.ObjectType
	storagephysicaldiskrelationship["selector"] = item.Selector

	storagephysicaldiskrelationships = append(storagephysicaldiskrelationships, storagephysicaldiskrelationship)
	return storagephysicaldiskrelationships
}
func flattenMapStoragePureArrayRelationship(p models.StoragePureArrayRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurearrayrelationships []map[string]interface{}
	var ret models.StoragePureArrayRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurearrayrelationship := make(map[string]interface{})
	storagepurearrayrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurearrayrelationship["class_id"] = item.ClassId
	storagepurearrayrelationship["moid"] = item.Moid
	storagepurearrayrelationship["object_type"] = item.ObjectType
	storagepurearrayrelationship["selector"] = item.Selector

	storagepurearrayrelationships = append(storagepurearrayrelationships, storagepurearrayrelationship)
	return storagepurearrayrelationships
}
func flattenMapStoragePureControllerRelationship(p models.StoragePureControllerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurecontrollerrelationships []map[string]interface{}
	var ret models.StoragePureControllerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurecontrollerrelationship := make(map[string]interface{})
	storagepurecontrollerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurecontrollerrelationship["class_id"] = item.ClassId
	storagepurecontrollerrelationship["moid"] = item.Moid
	storagepurecontrollerrelationship["object_type"] = item.ObjectType
	storagepurecontrollerrelationship["selector"] = item.Selector

	storagepurecontrollerrelationships = append(storagepurecontrollerrelationships, storagepurecontrollerrelationship)
	return storagepurecontrollerrelationships
}
func flattenMapStoragePureHostRelationship(p models.StoragePureHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostrelationships []map[string]interface{}
	var ret models.StoragePureHostRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurehostrelationship := make(map[string]interface{})
	storagepurehostrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurehostrelationship["class_id"] = item.ClassId
	storagepurehostrelationship["moid"] = item.Moid
	storagepurehostrelationship["object_type"] = item.ObjectType
	storagepurehostrelationship["selector"] = item.Selector

	storagepurehostrelationships = append(storagepurehostrelationships, storagepurehostrelationship)
	return storagepurehostrelationships
}
func flattenMapStoragePureHostGroupRelationship(p models.StoragePureHostGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostgrouprelationships []map[string]interface{}
	var ret models.StoragePureHostGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurehostgrouprelationship := make(map[string]interface{})
	storagepurehostgrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurehostgrouprelationship["class_id"] = item.ClassId
	storagepurehostgrouprelationship["moid"] = item.Moid
	storagepurehostgrouprelationship["object_type"] = item.ObjectType
	storagepurehostgrouprelationship["selector"] = item.Selector

	storagepurehostgrouprelationships = append(storagepurehostgrouprelationships, storagepurehostgrouprelationship)
	return storagepurehostgrouprelationships
}
func flattenMapStoragePureHostUtilization(p models.StoragePureHostUtilization, d *schema.ResourceData) []map[string]interface{} {
	var storagepurehostutilizations []map[string]interface{}
	var ret models.StoragePureHostUtilization
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	storagepurehostutilization := make(map[string]interface{})
	storagepurehostutilization["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurehostutilization["available"] = item.Available
	storagepurehostutilization["class_id"] = item.ClassId
	storagepurehostutilization["data_reduction"] = item.DataReduction
	storagepurehostutilization["free"] = item.Free
	storagepurehostutilization["object_type"] = item.ObjectType
	storagepurehostutilization["snapshot"] = item.Snapshot
	storagepurehostutilization["thin_provisioned"] = item.ThinProvisioned
	storagepurehostutilization["total"] = item.Total
	storagepurehostutilization["total_reduction"] = item.TotalReduction
	storagepurehostutilization["used"] = item.Used
	storagepurehostutilization["volume"] = item.Volume

	storagepurehostutilizations = append(storagepurehostutilizations, storagepurehostutilization)
	return storagepurehostutilizations
}
func flattenMapStoragePureProtectionGroupRelationship(p models.StoragePureProtectionGroupRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepureprotectiongrouprelationships []map[string]interface{}
	var ret models.StoragePureProtectionGroupRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepureprotectiongrouprelationship := make(map[string]interface{})
	storagepureprotectiongrouprelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepureprotectiongrouprelationship["class_id"] = item.ClassId
	storagepureprotectiongrouprelationship["moid"] = item.Moid
	storagepureprotectiongrouprelationship["object_type"] = item.ObjectType
	storagepureprotectiongrouprelationship["selector"] = item.Selector

	storagepureprotectiongrouprelationships = append(storagepureprotectiongrouprelationships, storagepureprotectiongrouprelationship)
	return storagepureprotectiongrouprelationships
}
func flattenMapStoragePureProtectionGroupSnapshotRelationship(p models.StoragePureProtectionGroupSnapshotRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepureprotectiongroupsnapshotrelationships []map[string]interface{}
	var ret models.StoragePureProtectionGroupSnapshotRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepureprotectiongroupsnapshotrelationship := make(map[string]interface{})
	storagepureprotectiongroupsnapshotrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepureprotectiongroupsnapshotrelationship["class_id"] = item.ClassId
	storagepureprotectiongroupsnapshotrelationship["moid"] = item.Moid
	storagepureprotectiongroupsnapshotrelationship["object_type"] = item.ObjectType
	storagepureprotectiongroupsnapshotrelationship["selector"] = item.Selector

	storagepureprotectiongroupsnapshotrelationships = append(storagepureprotectiongroupsnapshotrelationships, storagepureprotectiongroupsnapshotrelationship)
	return storagepureprotectiongroupsnapshotrelationships
}
func flattenMapStoragePureVolumeRelationship(p models.StoragePureVolumeRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagepurevolumerelationships []map[string]interface{}
	var ret models.StoragePureVolumeRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagepurevolumerelationship := make(map[string]interface{})
	storagepurevolumerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagepurevolumerelationship["class_id"] = item.ClassId
	storagepurevolumerelationship["moid"] = item.Moid
	storagepurevolumerelationship["object_type"] = item.ObjectType
	storagepurevolumerelationship["selector"] = item.Selector

	storagepurevolumerelationships = append(storagepurevolumerelationships, storagepurevolumerelationship)
	return storagepurevolumerelationships
}
func flattenMapStorageSasExpanderRelationship(p models.StorageSasExpanderRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagesasexpanderrelationships []map[string]interface{}
	var ret models.StorageSasExpanderRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagesasexpanderrelationship := make(map[string]interface{})
	storagesasexpanderrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagesasexpanderrelationship["class_id"] = item.ClassId
	storagesasexpanderrelationship["moid"] = item.Moid
	storagesasexpanderrelationship["object_type"] = item.ObjectType
	storagesasexpanderrelationship["selector"] = item.Selector

	storagesasexpanderrelationships = append(storagesasexpanderrelationships, storagesasexpanderrelationship)
	return storagesasexpanderrelationships
}
func flattenMapStorageVirtualDriveRelationship(p models.StorageVirtualDriveRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriverelationships []map[string]interface{}
	var ret models.StorageVirtualDriveRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagevirtualdriverelationship := make(map[string]interface{})
	storagevirtualdriverelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagevirtualdriverelationship["class_id"] = item.ClassId
	storagevirtualdriverelationship["moid"] = item.Moid
	storagevirtualdriverelationship["object_type"] = item.ObjectType
	storagevirtualdriverelationship["selector"] = item.Selector

	storagevirtualdriverelationships = append(storagevirtualdriverelationships, storagevirtualdriverelationship)
	return storagevirtualdriverelationships
}
func flattenMapStorageVirtualDriveExtensionRelationship(p models.StorageVirtualDriveExtensionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var storagevirtualdriveextensionrelationships []map[string]interface{}
	var ret models.StorageVirtualDriveExtensionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	storagevirtualdriveextensionrelationship := make(map[string]interface{})
	storagevirtualdriveextensionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	storagevirtualdriveextensionrelationship["class_id"] = item.ClassId
	storagevirtualdriveextensionrelationship["moid"] = item.Moid
	storagevirtualdriveextensionrelationship["object_type"] = item.ObjectType
	storagevirtualdriveextensionrelationship["selector"] = item.Selector

	storagevirtualdriveextensionrelationships = append(storagevirtualdriveextensionrelationships, storagevirtualdriveextensionrelationship)
	return storagevirtualdriveextensionrelationships
}
func flattenMapTamBaseAdvisoryRelationship(p models.TamBaseAdvisoryRelationship, d *schema.ResourceData) []map[string]interface{} {
	var tambaseadvisoryrelationships []map[string]interface{}
	var ret models.TamBaseAdvisoryRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	tambaseadvisoryrelationship := make(map[string]interface{})
	tambaseadvisoryrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	tambaseadvisoryrelationship["class_id"] = item.ClassId
	tambaseadvisoryrelationship["moid"] = item.Moid
	tambaseadvisoryrelationship["object_type"] = item.ObjectType
	tambaseadvisoryrelationship["selector"] = item.Selector

	tambaseadvisoryrelationships = append(tambaseadvisoryrelationships, tambaseadvisoryrelationship)
	return tambaseadvisoryrelationships
}
func flattenMapTamBaseAdvisoryDetails(p models.TamBaseAdvisoryDetails, d *schema.ResourceData) []map[string]interface{} {
	var tambaseadvisorydetailss []map[string]interface{}
	var ret models.TamBaseAdvisoryDetails
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	tambaseadvisorydetails := make(map[string]interface{})
	tambaseadvisorydetails["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	tambaseadvisorydetails["class_id"] = item.ClassId
	tambaseadvisorydetails["description"] = item.Description
	tambaseadvisorydetails["object_type"] = item.ObjectType

	tambaseadvisorydetailss = append(tambaseadvisorydetailss, tambaseadvisorydetails)
	return tambaseadvisorydetailss
}
func flattenMapTamSeverity(p models.TamSeverity, d *schema.ResourceData) []map[string]interface{} {
	var tamseveritys []map[string]interface{}
	var ret models.TamSeverity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	tamseverity := make(map[string]interface{})
	tamseverity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	tamseverity["class_id"] = item.ClassId
	tamseverity["object_type"] = item.ObjectType

	tamseveritys = append(tamseveritys, tamseverity)
	return tamseveritys
}
func flattenMapTechsupportmanagementTechSupportBundleRelationship(p models.TechsupportmanagementTechSupportBundleRelationship, d *schema.ResourceData) []map[string]interface{} {
	var techsupportmanagementtechsupportbundlerelationships []map[string]interface{}
	var ret models.TechsupportmanagementTechSupportBundleRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	techsupportmanagementtechsupportbundlerelationship := make(map[string]interface{})
	techsupportmanagementtechsupportbundlerelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	techsupportmanagementtechsupportbundlerelationship["class_id"] = item.ClassId
	techsupportmanagementtechsupportbundlerelationship["moid"] = item.Moid
	techsupportmanagementtechsupportbundlerelationship["object_type"] = item.ObjectType
	techsupportmanagementtechsupportbundlerelationship["selector"] = item.Selector

	techsupportmanagementtechsupportbundlerelationships = append(techsupportmanagementtechsupportbundlerelationships, techsupportmanagementtechsupportbundlerelationship)
	return techsupportmanagementtechsupportbundlerelationships
}
func flattenMapTechsupportmanagementTechSupportStatusRelationship(p models.TechsupportmanagementTechSupportStatusRelationship, d *schema.ResourceData) []map[string]interface{} {
	var techsupportmanagementtechsupportstatusrelationships []map[string]interface{}
	var ret models.TechsupportmanagementTechSupportStatusRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	techsupportmanagementtechsupportstatusrelationship := make(map[string]interface{})
	techsupportmanagementtechsupportstatusrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	techsupportmanagementtechsupportstatusrelationship["class_id"] = item.ClassId
	techsupportmanagementtechsupportstatusrelationship["moid"] = item.Moid
	techsupportmanagementtechsupportstatusrelationship["object_type"] = item.ObjectType
	techsupportmanagementtechsupportstatusrelationship["selector"] = item.Selector

	techsupportmanagementtechsupportstatusrelationships = append(techsupportmanagementtechsupportstatusrelationships, techsupportmanagementtechsupportstatusrelationship)
	return techsupportmanagementtechsupportstatusrelationships
}
func flattenMapTopSystemRelationship(p models.TopSystemRelationship, d *schema.ResourceData) []map[string]interface{} {
	var topsystemrelationships []map[string]interface{}
	var ret models.TopSystemRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	topsystemrelationship := make(map[string]interface{})
	topsystemrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	topsystemrelationship["class_id"] = item.ClassId
	topsystemrelationship["moid"] = item.Moid
	topsystemrelationship["object_type"] = item.ObjectType
	topsystemrelationship["selector"] = item.Selector

	topsystemrelationships = append(topsystemrelationships, topsystemrelationship)
	return topsystemrelationships
}
func flattenMapUuidpoolBlockRelationship(p models.UuidpoolBlockRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolblockrelationships []map[string]interface{}
	var ret models.UuidpoolBlockRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpoolblockrelationship := make(map[string]interface{})
	uuidpoolblockrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpoolblockrelationship["class_id"] = item.ClassId
	uuidpoolblockrelationship["moid"] = item.Moid
	uuidpoolblockrelationship["object_type"] = item.ObjectType
	uuidpoolblockrelationship["selector"] = item.Selector

	uuidpoolblockrelationships = append(uuidpoolblockrelationships, uuidpoolblockrelationship)
	return uuidpoolblockrelationships
}
func flattenMapUuidpoolPoolRelationship(p models.UuidpoolPoolRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolpoolrelationships []map[string]interface{}
	var ret models.UuidpoolPoolRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpoolpoolrelationship := make(map[string]interface{})
	uuidpoolpoolrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpoolpoolrelationship["class_id"] = item.ClassId
	uuidpoolpoolrelationship["moid"] = item.Moid
	uuidpoolpoolrelationship["object_type"] = item.ObjectType
	uuidpoolpoolrelationship["selector"] = item.Selector

	uuidpoolpoolrelationships = append(uuidpoolpoolrelationships, uuidpoolpoolrelationship)
	return uuidpoolpoolrelationships
}
func flattenMapUuidpoolPoolMemberRelationship(p models.UuidpoolPoolMemberRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpoolpoolmemberrelationships []map[string]interface{}
	var ret models.UuidpoolPoolMemberRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpoolpoolmemberrelationship := make(map[string]interface{})
	uuidpoolpoolmemberrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpoolpoolmemberrelationship["class_id"] = item.ClassId
	uuidpoolpoolmemberrelationship["moid"] = item.Moid
	uuidpoolpoolmemberrelationship["object_type"] = item.ObjectType
	uuidpoolpoolmemberrelationship["selector"] = item.Selector

	uuidpoolpoolmemberrelationships = append(uuidpoolpoolmemberrelationships, uuidpoolpoolmemberrelationship)
	return uuidpoolpoolmemberrelationships
}
func flattenMapUuidpoolUniverseRelationship(p models.UuidpoolUniverseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluniverserelationships []map[string]interface{}
	var ret models.UuidpoolUniverseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpooluniverserelationship := make(map[string]interface{})
	uuidpooluniverserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpooluniverserelationship["class_id"] = item.ClassId
	uuidpooluniverserelationship["moid"] = item.Moid
	uuidpooluniverserelationship["object_type"] = item.ObjectType
	uuidpooluniverserelationship["selector"] = item.Selector

	uuidpooluniverserelationships = append(uuidpooluniverserelationships, uuidpooluniverserelationship)
	return uuidpooluniverserelationships
}
func flattenMapUuidpoolUuidBlock(p models.UuidpoolUuidBlock, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluuidblocks []map[string]interface{}
	var ret models.UuidpoolUuidBlock
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	uuidpooluuidblock := make(map[string]interface{})
	uuidpooluuidblock["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpooluuidblock["class_id"] = item.ClassId
	uuidpooluuidblock["from"] = item.From
	uuidpooluuidblock["object_type"] = item.ObjectType
	uuidpooluuidblock["size"] = item.Size
	uuidpooluuidblock["to"] = item.To

	uuidpooluuidblocks = append(uuidpooluuidblocks, uuidpooluuidblock)
	return uuidpooluuidblocks
}
func flattenMapUuidpoolUuidLeaseRelationship(p models.UuidpoolUuidLeaseRelationship, d *schema.ResourceData) []map[string]interface{} {
	var uuidpooluuidleaserelationships []map[string]interface{}
	var ret models.UuidpoolUuidLeaseRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	uuidpooluuidleaserelationship := make(map[string]interface{})
	uuidpooluuidleaserelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	uuidpooluuidleaserelationship["class_id"] = item.ClassId
	uuidpooluuidleaserelationship["moid"] = item.Moid
	uuidpooluuidleaserelationship["object_type"] = item.ObjectType
	uuidpooluuidleaserelationship["selector"] = item.Selector

	uuidpooluuidleaserelationships = append(uuidpooluuidleaserelationships, uuidpooluuidleaserelationship)
	return uuidpooluuidleaserelationships
}
func flattenMapVirtualizationComputeCapacity(p models.VirtualizationComputeCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcomputecapacitys []map[string]interface{}
	var ret models.VirtualizationComputeCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationcomputecapacity := make(map[string]interface{})
	virtualizationcomputecapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationcomputecapacity["capacity"] = item.Capacity
	virtualizationcomputecapacity["class_id"] = item.ClassId
	virtualizationcomputecapacity["free"] = item.Free
	virtualizationcomputecapacity["object_type"] = item.ObjectType
	virtualizationcomputecapacity["used"] = item.Used

	virtualizationcomputecapacitys = append(virtualizationcomputecapacitys, virtualizationcomputecapacity)
	return virtualizationcomputecapacitys
}
func flattenMapVirtualizationCpuInfo(p models.VirtualizationCpuInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationcpuinfos []map[string]interface{}
	var ret models.VirtualizationCpuInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationcpuinfo := make(map[string]interface{})
	virtualizationcpuinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationcpuinfo["class_id"] = item.ClassId
	virtualizationcpuinfo["cores"] = item.Cores
	virtualizationcpuinfo["description"] = item.Description
	virtualizationcpuinfo["object_type"] = item.ObjectType
	virtualizationcpuinfo["sockets"] = item.Sockets
	virtualizationcpuinfo["speed"] = item.Speed
	virtualizationcpuinfo["vendor"] = item.Vendor

	virtualizationcpuinfos = append(virtualizationcpuinfos, virtualizationcpuinfo)
	return virtualizationcpuinfos
}
func flattenMapVirtualizationGuestInfo(p models.VirtualizationGuestInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationguestinfos []map[string]interface{}
	var ret models.VirtualizationGuestInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationguestinfo := make(map[string]interface{})
	virtualizationguestinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationguestinfo["class_id"] = item.ClassId
	virtualizationguestinfo["hostname"] = item.Hostname
	virtualizationguestinfo["ip_address"] = item.IpAddress
	virtualizationguestinfo["name"] = item.Name
	virtualizationguestinfo["object_type"] = item.ObjectType
	virtualizationguestinfo["operating_system"] = item.OperatingSystem

	virtualizationguestinfos = append(virtualizationguestinfos, virtualizationguestinfo)
	return virtualizationguestinfos
}
func flattenMapVirtualizationMemoryCapacity(p models.VirtualizationMemoryCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationmemorycapacitys []map[string]interface{}
	var ret models.VirtualizationMemoryCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationmemorycapacity := make(map[string]interface{})
	virtualizationmemorycapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationmemorycapacity["capacity"] = item.Capacity
	virtualizationmemorycapacity["class_id"] = item.ClassId
	virtualizationmemorycapacity["free"] = item.Free
	virtualizationmemorycapacity["object_type"] = item.ObjectType
	virtualizationmemorycapacity["used"] = item.Used

	virtualizationmemorycapacitys = append(virtualizationmemorycapacitys, virtualizationmemorycapacity)
	return virtualizationmemorycapacitys
}
func flattenMapVirtualizationProductInfo(p models.VirtualizationProductInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationproductinfos []map[string]interface{}
	var ret models.VirtualizationProductInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationproductinfo := make(map[string]interface{})
	virtualizationproductinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationproductinfo["class_id"] = item.ClassId
	virtualizationproductinfo["object_type"] = item.ObjectType
	virtualizationproductinfo["product_name"] = item.ProductName
	virtualizationproductinfo["product_type"] = item.ProductType
	virtualizationproductinfo["product_vendor"] = item.ProductVendor
	virtualizationproductinfo["nr_version"] = item.Version

	virtualizationproductinfos = append(virtualizationproductinfos, virtualizationproductinfo)
	return virtualizationproductinfos
}
func flattenMapVirtualizationStorageCapacity(p models.VirtualizationStorageCapacity, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationstoragecapacitys []map[string]interface{}
	var ret models.VirtualizationStorageCapacity
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationstoragecapacity := make(map[string]interface{})
	virtualizationstoragecapacity["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationstoragecapacity["capacity"] = item.Capacity
	virtualizationstoragecapacity["class_id"] = item.ClassId
	virtualizationstoragecapacity["free"] = item.Free
	virtualizationstoragecapacity["object_type"] = item.ObjectType
	virtualizationstoragecapacity["used"] = item.Used

	virtualizationstoragecapacitys = append(virtualizationstoragecapacitys, virtualizationstoragecapacity)
	return virtualizationstoragecapacitys
}
func flattenMapVirtualizationVmwareClusterRelationship(p models.VirtualizationVmwareClusterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareclusterrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareClusterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwareclusterrelationship := make(map[string]interface{})
	virtualizationvmwareclusterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwareclusterrelationship["class_id"] = item.ClassId
	virtualizationvmwareclusterrelationship["moid"] = item.Moid
	virtualizationvmwareclusterrelationship["object_type"] = item.ObjectType
	virtualizationvmwareclusterrelationship["selector"] = item.Selector

	virtualizationvmwareclusterrelationships = append(virtualizationvmwareclusterrelationships, virtualizationvmwareclusterrelationship)
	return virtualizationvmwareclusterrelationships
}
func flattenMapVirtualizationVmwareDatacenterRelationship(p models.VirtualizationVmwareDatacenterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwaredatacenterrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareDatacenterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwaredatacenterrelationship := make(map[string]interface{})
	virtualizationvmwaredatacenterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwaredatacenterrelationship["class_id"] = item.ClassId
	virtualizationvmwaredatacenterrelationship["moid"] = item.Moid
	virtualizationvmwaredatacenterrelationship["object_type"] = item.ObjectType
	virtualizationvmwaredatacenterrelationship["selector"] = item.Selector

	virtualizationvmwaredatacenterrelationships = append(virtualizationvmwaredatacenterrelationships, virtualizationvmwaredatacenterrelationship)
	return virtualizationvmwaredatacenterrelationships
}
func flattenMapVirtualizationVmwareHostRelationship(p models.VirtualizationVmwareHostRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarehostrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareHostRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarehostrelationship := make(map[string]interface{})
	virtualizationvmwarehostrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarehostrelationship["class_id"] = item.ClassId
	virtualizationvmwarehostrelationship["moid"] = item.Moid
	virtualizationvmwarehostrelationship["object_type"] = item.ObjectType
	virtualizationvmwarehostrelationship["selector"] = item.Selector

	virtualizationvmwarehostrelationships = append(virtualizationvmwarehostrelationships, virtualizationvmwarehostrelationship)
	return virtualizationvmwarehostrelationships
}
func flattenMapVirtualizationVmwareRemoteDisplayInfo(p models.VirtualizationVmwareRemoteDisplayInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareremotedisplayinfos []map[string]interface{}
	var ret models.VirtualizationVmwareRemoteDisplayInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwareremotedisplayinfo := make(map[string]interface{})
	virtualizationvmwareremotedisplayinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwareremotedisplayinfo["class_id"] = item.ClassId
	virtualizationvmwareremotedisplayinfo["object_type"] = item.ObjectType
	virtualizationvmwareremotedisplayinfo["remote_display_password"] = item.RemoteDisplayPassword
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_key"] = item.RemoteDisplayVncKey
	virtualizationvmwareremotedisplayinfo["remote_display_vnc_port"] = item.RemoteDisplayVncPort

	virtualizationvmwareremotedisplayinfos = append(virtualizationvmwareremotedisplayinfos, virtualizationvmwareremotedisplayinfo)
	return virtualizationvmwareremotedisplayinfos
}
func flattenMapVirtualizationVmwareResourceConsumption(p models.VirtualizationVmwareResourceConsumption, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwareresourceconsumptions []map[string]interface{}
	var ret models.VirtualizationVmwareResourceConsumption
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwareresourceconsumption := make(map[string]interface{})
	virtualizationvmwareresourceconsumption["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwareresourceconsumption["class_id"] = item.ClassId
	virtualizationvmwareresourceconsumption["cpu_consumed"] = item.CpuConsumed
	virtualizationvmwareresourceconsumption["memory_consumed"] = item.MemoryConsumed
	virtualizationvmwareresourceconsumption["object_type"] = item.ObjectType

	virtualizationvmwareresourceconsumptions = append(virtualizationvmwareresourceconsumptions, virtualizationvmwareresourceconsumption)
	return virtualizationvmwareresourceconsumptions
}
func flattenMapVirtualizationVmwareVcenterRelationship(p models.VirtualizationVmwareVcenterRelationship, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevcenterrelationships []map[string]interface{}
	var ret models.VirtualizationVmwareVcenterRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	virtualizationvmwarevcenterrelationship := make(map[string]interface{})
	virtualizationvmwarevcenterrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevcenterrelationship["class_id"] = item.ClassId
	virtualizationvmwarevcenterrelationship["moid"] = item.Moid
	virtualizationvmwarevcenterrelationship["object_type"] = item.ObjectType
	virtualizationvmwarevcenterrelationship["selector"] = item.Selector

	virtualizationvmwarevcenterrelationships = append(virtualizationvmwarevcenterrelationships, virtualizationvmwarevcenterrelationship)
	return virtualizationvmwarevcenterrelationships
}
func flattenMapVirtualizationVmwareVmCpuShareInfo(p models.VirtualizationVmwareVmCpuShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmcpushareinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmCpuShareInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmcpushareinfo := make(map[string]interface{})
	virtualizationvmwarevmcpushareinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmcpushareinfo["class_id"] = item.ClassId
	virtualizationvmwarevmcpushareinfo["cpu_limit"] = item.CpuLimit
	virtualizationvmwarevmcpushareinfo["cpu_overhead_limit"] = item.CpuOverheadLimit
	virtualizationvmwarevmcpushareinfo["cpu_reservation"] = item.CpuReservation
	virtualizationvmwarevmcpushareinfo["cpu_shares"] = item.CpuShares
	virtualizationvmwarevmcpushareinfo["object_type"] = item.ObjectType

	virtualizationvmwarevmcpushareinfos = append(virtualizationvmwarevmcpushareinfos, virtualizationvmwarevmcpushareinfo)
	return virtualizationvmwarevmcpushareinfos
}
func flattenMapVirtualizationVmwareVmCpuSocketInfo(p models.VirtualizationVmwareVmCpuSocketInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmcpusocketinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmCpuSocketInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmcpusocketinfo := make(map[string]interface{})
	virtualizationvmwarevmcpusocketinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmcpusocketinfo["class_id"] = item.ClassId
	virtualizationvmwarevmcpusocketinfo["cores_per_socket"] = item.CoresPerSocket
	virtualizationvmwarevmcpusocketinfo["num_cpus"] = item.NumCpus
	virtualizationvmwarevmcpusocketinfo["num_sockets"] = item.NumSockets
	virtualizationvmwarevmcpusocketinfo["object_type"] = item.ObjectType

	virtualizationvmwarevmcpusocketinfos = append(virtualizationvmwarevmcpusocketinfos, virtualizationvmwarevmcpusocketinfo)
	return virtualizationvmwarevmcpusocketinfos
}
func flattenMapVirtualizationVmwareVmDiskCommitInfo(p models.VirtualizationVmwareVmDiskCommitInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmdiskcommitinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmDiskCommitInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmdiskcommitinfo := make(map[string]interface{})
	virtualizationvmwarevmdiskcommitinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmdiskcommitinfo["class_id"] = item.ClassId
	virtualizationvmwarevmdiskcommitinfo["committed_disk"] = item.CommittedDisk
	virtualizationvmwarevmdiskcommitinfo["object_type"] = item.ObjectType
	virtualizationvmwarevmdiskcommitinfo["un_committed_disk"] = item.UnCommittedDisk
	virtualizationvmwarevmdiskcommitinfo["unshared_disk"] = item.UnsharedDisk

	virtualizationvmwarevmdiskcommitinfos = append(virtualizationvmwarevmdiskcommitinfos, virtualizationvmwarevmdiskcommitinfo)
	return virtualizationvmwarevmdiskcommitinfos
}
func flattenMapVirtualizationVmwareVmMemoryShareInfo(p models.VirtualizationVmwareVmMemoryShareInfo, d *schema.ResourceData) []map[string]interface{} {
	var virtualizationvmwarevmmemoryshareinfos []map[string]interface{}
	var ret models.VirtualizationVmwareVmMemoryShareInfo
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	virtualizationvmwarevmmemoryshareinfo := make(map[string]interface{})
	virtualizationvmwarevmmemoryshareinfo["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	virtualizationvmwarevmmemoryshareinfo["class_id"] = item.ClassId
	virtualizationvmwarevmmemoryshareinfo["mem_limit"] = item.MemLimit
	virtualizationvmwarevmmemoryshareinfo["mem_overhead_limit"] = item.MemOverheadLimit
	virtualizationvmwarevmmemoryshareinfo["mem_reservation"] = item.MemReservation
	virtualizationvmwarevmmemoryshareinfo["mem_shares"] = item.MemShares
	virtualizationvmwarevmmemoryshareinfo["object_type"] = item.ObjectType

	virtualizationvmwarevmmemoryshareinfos = append(virtualizationvmwarevmmemoryshareinfos, virtualizationvmwarevmmemoryshareinfo)
	return virtualizationvmwarevmmemoryshareinfos
}
func flattenMapVnicArfsSettings(p models.VnicArfsSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicarfssettingss []map[string]interface{}
	var ret models.VnicArfsSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicarfssettings := make(map[string]interface{})
	vnicarfssettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicarfssettings["class_id"] = item.ClassId
	vnicarfssettings["enabled"] = item.Enabled
	vnicarfssettings["object_type"] = item.ObjectType

	vnicarfssettingss = append(vnicarfssettingss, vnicarfssettings)
	return vnicarfssettingss
}
func flattenMapVnicCdn(p models.VnicCdn, d *schema.ResourceData) []map[string]interface{} {
	var vniccdns []map[string]interface{}
	var ret models.VnicCdn
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vniccdn := make(map[string]interface{})
	vniccdn["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniccdn["class_id"] = item.ClassId
	vniccdn["object_type"] = item.ObjectType
	vniccdn["nr_source"] = item.Source
	vniccdn["value"] = item.Value

	vniccdns = append(vniccdns, vniccdn)
	return vniccdns
}
func flattenMapVnicCompletionQueueSettings(p models.VnicCompletionQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vniccompletionqueuesettingss []map[string]interface{}
	var ret models.VnicCompletionQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vniccompletionqueuesettings := make(map[string]interface{})
	vniccompletionqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniccompletionqueuesettings["class_id"] = item.ClassId
	vniccompletionqueuesettings["nr_count"] = item.Count
	vniccompletionqueuesettings["object_type"] = item.ObjectType
	vniccompletionqueuesettings["ring_size"] = item.RingSize

	vniccompletionqueuesettingss = append(vniccompletionqueuesettingss, vniccompletionqueuesettings)
	return vniccompletionqueuesettingss
}
func flattenMapVnicEthAdapterPolicyRelationship(p models.VnicEthAdapterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethadapterpolicyrelationships []map[string]interface{}
	var ret models.VnicEthAdapterPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethadapterpolicyrelationship := make(map[string]interface{})
	vnicethadapterpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethadapterpolicyrelationship["class_id"] = item.ClassId
	vnicethadapterpolicyrelationship["moid"] = item.Moid
	vnicethadapterpolicyrelationship["object_type"] = item.ObjectType
	vnicethadapterpolicyrelationship["selector"] = item.Selector

	vnicethadapterpolicyrelationships = append(vnicethadapterpolicyrelationships, vnicethadapterpolicyrelationship)
	return vnicethadapterpolicyrelationships
}
func flattenMapVnicEthIfRelationship(p models.VnicEthIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethifrelationships []map[string]interface{}
	var ret models.VnicEthIfRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethifrelationship := make(map[string]interface{})
	vnicethifrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethifrelationship["class_id"] = item.ClassId
	vnicethifrelationship["moid"] = item.Moid
	vnicethifrelationship["object_type"] = item.ObjectType
	vnicethifrelationship["selector"] = item.Selector

	vnicethifrelationships = append(vnicethifrelationships, vnicethifrelationship)
	return vnicethifrelationships
}
func flattenMapVnicEthInterruptSettings(p models.VnicEthInterruptSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethinterruptsettingss []map[string]interface{}
	var ret models.VnicEthInterruptSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicethinterruptsettings := make(map[string]interface{})
	vnicethinterruptsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethinterruptsettings["class_id"] = item.ClassId
	vnicethinterruptsettings["coalescing_time"] = item.CoalescingTime
	vnicethinterruptsettings["coalescing_type"] = item.CoalescingType
	vnicethinterruptsettings["nr_count"] = item.Count
	vnicethinterruptsettings["mode"] = item.Mode
	vnicethinterruptsettings["object_type"] = item.ObjectType

	vnicethinterruptsettingss = append(vnicethinterruptsettingss, vnicethinterruptsettings)
	return vnicethinterruptsettingss
}
func flattenMapVnicEthNetworkPolicyRelationship(p models.VnicEthNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethnetworkpolicyrelationships []map[string]interface{}
	var ret models.VnicEthNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethnetworkpolicyrelationship := make(map[string]interface{})
	vnicethnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethnetworkpolicyrelationship["class_id"] = item.ClassId
	vnicethnetworkpolicyrelationship["moid"] = item.Moid
	vnicethnetworkpolicyrelationship["object_type"] = item.ObjectType
	vnicethnetworkpolicyrelationship["selector"] = item.Selector

	vnicethnetworkpolicyrelationships = append(vnicethnetworkpolicyrelationships, vnicethnetworkpolicyrelationship)
	return vnicethnetworkpolicyrelationships
}
func flattenMapVnicEthQosPolicyRelationship(p models.VnicEthQosPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicethqospolicyrelationships []map[string]interface{}
	var ret models.VnicEthQosPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicethqospolicyrelationship := make(map[string]interface{})
	vnicethqospolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethqospolicyrelationship["class_id"] = item.ClassId
	vnicethqospolicyrelationship["moid"] = item.Moid
	vnicethqospolicyrelationship["object_type"] = item.ObjectType
	vnicethqospolicyrelationship["selector"] = item.Selector

	vnicethqospolicyrelationships = append(vnicethqospolicyrelationships, vnicethqospolicyrelationship)
	return vnicethqospolicyrelationships
}
func flattenMapVnicEthRxQueueSettings(p models.VnicEthRxQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethrxqueuesettingss []map[string]interface{}
	var ret models.VnicEthRxQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicethrxqueuesettings := make(map[string]interface{})
	vnicethrxqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethrxqueuesettings["class_id"] = item.ClassId
	vnicethrxqueuesettings["nr_count"] = item.Count
	vnicethrxqueuesettings["object_type"] = item.ObjectType
	vnicethrxqueuesettings["ring_size"] = item.RingSize

	vnicethrxqueuesettingss = append(vnicethrxqueuesettingss, vnicethrxqueuesettings)
	return vnicethrxqueuesettingss
}
func flattenMapVnicEthTxQueueSettings(p models.VnicEthTxQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicethtxqueuesettingss []map[string]interface{}
	var ret models.VnicEthTxQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicethtxqueuesettings := make(map[string]interface{})
	vnicethtxqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicethtxqueuesettings["class_id"] = item.ClassId
	vnicethtxqueuesettings["nr_count"] = item.Count
	vnicethtxqueuesettings["object_type"] = item.ObjectType
	vnicethtxqueuesettings["ring_size"] = item.RingSize

	vnicethtxqueuesettingss = append(vnicethtxqueuesettingss, vnicethtxqueuesettings)
	return vnicethtxqueuesettingss
}
func flattenMapVnicFcAdapterPolicyRelationship(p models.VnicFcAdapterPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcadapterpolicyrelationships []map[string]interface{}
	var ret models.VnicFcAdapterPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcadapterpolicyrelationship := make(map[string]interface{})
	vnicfcadapterpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcadapterpolicyrelationship["class_id"] = item.ClassId
	vnicfcadapterpolicyrelationship["moid"] = item.Moid
	vnicfcadapterpolicyrelationship["object_type"] = item.ObjectType
	vnicfcadapterpolicyrelationship["selector"] = item.Selector

	vnicfcadapterpolicyrelationships = append(vnicfcadapterpolicyrelationships, vnicfcadapterpolicyrelationship)
	return vnicfcadapterpolicyrelationships
}
func flattenMapVnicFcErrorRecoverySettings(p models.VnicFcErrorRecoverySettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcerrorrecoverysettingss []map[string]interface{}
	var ret models.VnicFcErrorRecoverySettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicfcerrorrecoverysettings := make(map[string]interface{})
	vnicfcerrorrecoverysettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcerrorrecoverysettings["class_id"] = item.ClassId
	vnicfcerrorrecoverysettings["enabled"] = item.Enabled
	vnicfcerrorrecoverysettings["io_retry_count"] = item.IoRetryCount
	vnicfcerrorrecoverysettings["io_retry_timeout"] = item.IoRetryTimeout
	vnicfcerrorrecoverysettings["link_down_timeout"] = item.LinkDownTimeout
	vnicfcerrorrecoverysettings["object_type"] = item.ObjectType
	vnicfcerrorrecoverysettings["port_down_timeout"] = item.PortDownTimeout

	vnicfcerrorrecoverysettingss = append(vnicfcerrorrecoverysettingss, vnicfcerrorrecoverysettings)
	return vnicfcerrorrecoverysettingss
}
func flattenMapVnicFcIfRelationship(p models.VnicFcIfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcifrelationships []map[string]interface{}
	var ret models.VnicFcIfRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcifrelationship := make(map[string]interface{})
	vnicfcifrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcifrelationship["class_id"] = item.ClassId
	vnicfcifrelationship["moid"] = item.Moid
	vnicfcifrelationship["object_type"] = item.ObjectType
	vnicfcifrelationship["selector"] = item.Selector

	vnicfcifrelationships = append(vnicfcifrelationships, vnicfcifrelationship)
	return vnicfcifrelationships
}
func flattenMapVnicFcInterruptSettings(p models.VnicFcInterruptSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcinterruptsettingss []map[string]interface{}
	var ret models.VnicFcInterruptSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicfcinterruptsettings := make(map[string]interface{})
	vnicfcinterruptsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcinterruptsettings["class_id"] = item.ClassId
	vnicfcinterruptsettings["mode"] = item.Mode
	vnicfcinterruptsettings["object_type"] = item.ObjectType

	vnicfcinterruptsettingss = append(vnicfcinterruptsettingss, vnicfcinterruptsettings)
	return vnicfcinterruptsettingss
}
func flattenMapVnicFcNetworkPolicyRelationship(p models.VnicFcNetworkPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcnetworkpolicyrelationships []map[string]interface{}
	var ret models.VnicFcNetworkPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcnetworkpolicyrelationship := make(map[string]interface{})
	vnicfcnetworkpolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcnetworkpolicyrelationship["class_id"] = item.ClassId
	vnicfcnetworkpolicyrelationship["moid"] = item.Moid
	vnicfcnetworkpolicyrelationship["object_type"] = item.ObjectType
	vnicfcnetworkpolicyrelationship["selector"] = item.Selector

	vnicfcnetworkpolicyrelationships = append(vnicfcnetworkpolicyrelationships, vnicfcnetworkpolicyrelationship)
	return vnicfcnetworkpolicyrelationships
}
func flattenMapVnicFcQosPolicyRelationship(p models.VnicFcQosPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqospolicyrelationships []map[string]interface{}
	var ret models.VnicFcQosPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicfcqospolicyrelationship := make(map[string]interface{})
	vnicfcqospolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcqospolicyrelationship["class_id"] = item.ClassId
	vnicfcqospolicyrelationship["moid"] = item.Moid
	vnicfcqospolicyrelationship["object_type"] = item.ObjectType
	vnicfcqospolicyrelationship["selector"] = item.Selector

	vnicfcqospolicyrelationships = append(vnicfcqospolicyrelationships, vnicfcqospolicyrelationship)
	return vnicfcqospolicyrelationships
}
func flattenMapVnicFcQueueSettings(p models.VnicFcQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicfcqueuesettingss []map[string]interface{}
	var ret models.VnicFcQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicfcqueuesettings := make(map[string]interface{})
	vnicfcqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicfcqueuesettings["class_id"] = item.ClassId
	vnicfcqueuesettings["nr_count"] = item.Count
	vnicfcqueuesettings["object_type"] = item.ObjectType
	vnicfcqueuesettings["ring_size"] = item.RingSize

	vnicfcqueuesettingss = append(vnicfcqueuesettingss, vnicfcqueuesettings)
	return vnicfcqueuesettingss
}
func flattenMapVnicFlogiSettings(p models.VnicFlogiSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicflogisettingss []map[string]interface{}
	var ret models.VnicFlogiSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicflogisettings := make(map[string]interface{})
	vnicflogisettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicflogisettings["class_id"] = item.ClassId
	vnicflogisettings["object_type"] = item.ObjectType
	vnicflogisettings["retries"] = item.Retries
	vnicflogisettings["timeout"] = item.Timeout

	vnicflogisettingss = append(vnicflogisettingss, vnicflogisettings)
	return vnicflogisettingss
}
func flattenMapVnicLanConnectivityPolicyRelationship(p models.VnicLanConnectivityPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vniclanconnectivitypolicyrelationships []map[string]interface{}
	var ret models.VnicLanConnectivityPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vniclanconnectivitypolicyrelationship := make(map[string]interface{})
	vniclanconnectivitypolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vniclanconnectivitypolicyrelationship["class_id"] = item.ClassId
	vniclanconnectivitypolicyrelationship["moid"] = item.Moid
	vniclanconnectivitypolicyrelationship["object_type"] = item.ObjectType
	vniclanconnectivitypolicyrelationship["selector"] = item.Selector

	vniclanconnectivitypolicyrelationships = append(vniclanconnectivitypolicyrelationships, vniclanconnectivitypolicyrelationship)
	return vniclanconnectivitypolicyrelationships
}
func flattenMapVnicNvgreSettings(p models.VnicNvgreSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicnvgresettingss []map[string]interface{}
	var ret models.VnicNvgreSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicnvgresettings := make(map[string]interface{})
	vnicnvgresettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicnvgresettings["class_id"] = item.ClassId
	vnicnvgresettings["enabled"] = item.Enabled
	vnicnvgresettings["object_type"] = item.ObjectType

	vnicnvgresettingss = append(vnicnvgresettingss, vnicnvgresettings)
	return vnicnvgresettingss
}
func flattenMapVnicPlacementSettings(p models.VnicPlacementSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicplacementsettingss []map[string]interface{}
	var ret models.VnicPlacementSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicplacementsettings := make(map[string]interface{})
	vnicplacementsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicplacementsettings["class_id"] = item.ClassId
	vnicplacementsettings["id"] = item.Id
	vnicplacementsettings["object_type"] = item.ObjectType
	vnicplacementsettings["pci_link"] = item.PciLink
	vnicplacementsettings["switch_id"] = item.SwitchId
	vnicplacementsettings["uplink"] = item.Uplink

	vnicplacementsettingss = append(vnicplacementsettingss, vnicplacementsettings)
	return vnicplacementsettingss
}
func flattenMapVnicPlogiSettings(p models.VnicPlogiSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicplogisettingss []map[string]interface{}
	var ret models.VnicPlogiSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicplogisettings := make(map[string]interface{})
	vnicplogisettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicplogisettings["class_id"] = item.ClassId
	vnicplogisettings["object_type"] = item.ObjectType
	vnicplogisettings["retries"] = item.Retries
	vnicplogisettings["timeout"] = item.Timeout

	vnicplogisettingss = append(vnicplogisettingss, vnicplogisettings)
	return vnicplogisettingss
}
func flattenMapVnicRoceSettings(p models.VnicRoceSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicrocesettingss []map[string]interface{}
	var ret models.VnicRoceSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicrocesettings := make(map[string]interface{})
	vnicrocesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicrocesettings["class_id"] = item.ClassId
	vnicrocesettings["class_of_service"] = item.ClassOfService
	vnicrocesettings["enabled"] = item.Enabled
	vnicrocesettings["memory_regions"] = item.MemoryRegions
	vnicrocesettings["object_type"] = item.ObjectType
	vnicrocesettings["queue_pairs"] = item.QueuePairs
	vnicrocesettings["resource_groups"] = item.ResourceGroups
	vnicrocesettings["nr_version"] = item.Version

	vnicrocesettingss = append(vnicrocesettingss, vnicrocesettings)
	return vnicrocesettingss
}
func flattenMapVnicRssHashSettings(p models.VnicRssHashSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicrsshashsettingss []map[string]interface{}
	var ret models.VnicRssHashSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicrsshashsettings := make(map[string]interface{})
	vnicrsshashsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicrsshashsettings["class_id"] = item.ClassId
	vnicrsshashsettings["ipv4_hash"] = item.Ipv4Hash
	vnicrsshashsettings["ipv6_ext_hash"] = item.Ipv6ExtHash
	vnicrsshashsettings["ipv6_hash"] = item.Ipv6Hash
	vnicrsshashsettings["object_type"] = item.ObjectType
	vnicrsshashsettings["tcp_ipv4_hash"] = item.TcpIpv4Hash
	vnicrsshashsettings["tcp_ipv6_ext_hash"] = item.TcpIpv6ExtHash
	vnicrsshashsettings["tcp_ipv6_hash"] = item.TcpIpv6Hash
	vnicrsshashsettings["udp_ipv4_hash"] = item.UdpIpv4Hash
	vnicrsshashsettings["udp_ipv6_hash"] = item.UdpIpv6Hash

	vnicrsshashsettingss = append(vnicrsshashsettingss, vnicrsshashsettings)
	return vnicrsshashsettingss
}
func flattenMapVnicSanConnectivityPolicyRelationship(p models.VnicSanConnectivityPolicyRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vnicsanconnectivitypolicyrelationships []map[string]interface{}
	var ret models.VnicSanConnectivityPolicyRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vnicsanconnectivitypolicyrelationship := make(map[string]interface{})
	vnicsanconnectivitypolicyrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicsanconnectivitypolicyrelationship["class_id"] = item.ClassId
	vnicsanconnectivitypolicyrelationship["moid"] = item.Moid
	vnicsanconnectivitypolicyrelationship["object_type"] = item.ObjectType
	vnicsanconnectivitypolicyrelationship["selector"] = item.Selector

	vnicsanconnectivitypolicyrelationships = append(vnicsanconnectivitypolicyrelationships, vnicsanconnectivitypolicyrelationship)
	return vnicsanconnectivitypolicyrelationships
}
func flattenMapVnicScsiQueueSettings(p models.VnicScsiQueueSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicscsiqueuesettingss []map[string]interface{}
	var ret models.VnicScsiQueueSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicscsiqueuesettings := make(map[string]interface{})
	vnicscsiqueuesettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicscsiqueuesettings["class_id"] = item.ClassId
	vnicscsiqueuesettings["nr_count"] = item.Count
	vnicscsiqueuesettings["object_type"] = item.ObjectType
	vnicscsiqueuesettings["ring_size"] = item.RingSize

	vnicscsiqueuesettingss = append(vnicscsiqueuesettingss, vnicscsiqueuesettings)
	return vnicscsiqueuesettingss
}
func flattenMapVnicTcpOffloadSettings(p models.VnicTcpOffloadSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnictcpoffloadsettingss []map[string]interface{}
	var ret models.VnicTcpOffloadSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnictcpoffloadsettings := make(map[string]interface{})
	vnictcpoffloadsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnictcpoffloadsettings["class_id"] = item.ClassId
	vnictcpoffloadsettings["large_receive"] = item.LargeReceive
	vnictcpoffloadsettings["large_send"] = item.LargeSend
	vnictcpoffloadsettings["object_type"] = item.ObjectType
	vnictcpoffloadsettings["rx_checksum"] = item.RxChecksum
	vnictcpoffloadsettings["tx_checksum"] = item.TxChecksum

	vnictcpoffloadsettingss = append(vnictcpoffloadsettingss, vnictcpoffloadsettings)
	return vnictcpoffloadsettingss
}
func flattenMapVnicUsnicSettings(p models.VnicUsnicSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicusnicsettingss []map[string]interface{}
	var ret models.VnicUsnicSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicusnicsettings := make(map[string]interface{})
	vnicusnicsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicusnicsettings["class_id"] = item.ClassId
	vnicusnicsettings["cos"] = item.Cos
	vnicusnicsettings["nr_count"] = item.Count
	vnicusnicsettings["object_type"] = item.ObjectType
	vnicusnicsettings["usnic_adapter_policy"] = item.UsnicAdapterPolicy

	vnicusnicsettingss = append(vnicusnicsettingss, vnicusnicsettings)
	return vnicusnicsettingss
}
func flattenMapVnicVlanSettings(p models.VnicVlanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvlansettingss []map[string]interface{}
	var ret models.VnicVlanSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvlansettings := make(map[string]interface{})
	vnicvlansettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvlansettings["allowed_vlans"] = item.AllowedVlans
	vnicvlansettings["class_id"] = item.ClassId
	vnicvlansettings["default_vlan"] = item.DefaultVlan
	vnicvlansettings["mode"] = item.Mode
	vnicvlansettings["object_type"] = item.ObjectType

	vnicvlansettingss = append(vnicvlansettingss, vnicvlansettings)
	return vnicvlansettingss
}
func flattenMapVnicVmqSettings(p models.VnicVmqSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvmqsettingss []map[string]interface{}
	var ret models.VnicVmqSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvmqsettings := make(map[string]interface{})
	vnicvmqsettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvmqsettings["class_id"] = item.ClassId
	vnicvmqsettings["enabled"] = item.Enabled
	vnicvmqsettings["multi_queue_support"] = item.MultiQueueSupport
	vnicvmqsettings["num_interrupts"] = item.NumInterrupts
	vnicvmqsettings["num_sub_vnics"] = item.NumSubVnics
	vnicvmqsettings["num_vmqs"] = item.NumVmqs
	vnicvmqsettings["object_type"] = item.ObjectType
	vnicvmqsettings["vmmq_adapter_policy"] = item.VmmqAdapterPolicy

	vnicvmqsettingss = append(vnicvmqsettingss, vnicvmqsettings)
	return vnicvmqsettingss
}
func flattenMapVnicVsanSettings(p models.VnicVsanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvsansettingss []map[string]interface{}
	var ret models.VnicVsanSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvsansettings := make(map[string]interface{})
	vnicvsansettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvsansettings["class_id"] = item.ClassId
	vnicvsansettings["id"] = item.Id
	vnicvsansettings["object_type"] = item.ObjectType

	vnicvsansettingss = append(vnicvsansettingss, vnicvsansettings)
	return vnicvsansettingss
}
func flattenMapVnicVxlanSettings(p models.VnicVxlanSettings, d *schema.ResourceData) []map[string]interface{} {
	var vnicvxlansettingss []map[string]interface{}
	var ret models.VnicVxlanSettings
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	vnicvxlansettings := make(map[string]interface{})
	vnicvxlansettings["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vnicvxlansettings["class_id"] = item.ClassId
	vnicvxlansettings["enabled"] = item.Enabled
	vnicvxlansettings["object_type"] = item.ObjectType

	vnicvxlansettingss = append(vnicvxlansettingss, vnicvxlansettings)
	return vnicvxlansettingss
}
func flattenMapVrfVrfRelationship(p models.VrfVrfRelationship, d *schema.ResourceData) []map[string]interface{} {
	var vrfvrfrelationships []map[string]interface{}
	var ret models.VrfVrfRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	vrfvrfrelationship := make(map[string]interface{})
	vrfvrfrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	vrfvrfrelationship["class_id"] = item.ClassId
	vrfvrfrelationship["moid"] = item.Moid
	vrfvrfrelationship["object_type"] = item.ObjectType
	vrfvrfrelationship["selector"] = item.Selector

	vrfvrfrelationships = append(vrfvrfrelationships, vrfvrfrelationship)
	return vrfvrfrelationships
}
func flattenMapWorkflowCatalogRelationship(p models.WorkflowCatalogRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowcatalogrelationships []map[string]interface{}
	var ret models.WorkflowCatalogRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowcatalogrelationship := make(map[string]interface{})
	workflowcatalogrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowcatalogrelationship["class_id"] = item.ClassId
	workflowcatalogrelationship["moid"] = item.Moid
	workflowcatalogrelationship["object_type"] = item.ObjectType
	workflowcatalogrelationship["selector"] = item.Selector

	workflowcatalogrelationships = append(workflowcatalogrelationships, workflowcatalogrelationship)
	return workflowcatalogrelationships
}
func flattenMapWorkflowErrorResponseHandlerRelationship(p models.WorkflowErrorResponseHandlerRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowerrorresponsehandlerrelationships []map[string]interface{}
	var ret models.WorkflowErrorResponseHandlerRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowerrorresponsehandlerrelationship := make(map[string]interface{})
	workflowerrorresponsehandlerrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowerrorresponsehandlerrelationship["class_id"] = item.ClassId
	workflowerrorresponsehandlerrelationship["moid"] = item.Moid
	workflowerrorresponsehandlerrelationship["object_type"] = item.ObjectType
	workflowerrorresponsehandlerrelationship["selector"] = item.Selector

	workflowerrorresponsehandlerrelationships = append(workflowerrorresponsehandlerrelationships, workflowerrorresponsehandlerrelationship)
	return workflowerrorresponsehandlerrelationships
}
func flattenMapWorkflowInternalProperties(p models.WorkflowInternalProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowinternalpropertiess []map[string]interface{}
	var ret models.WorkflowInternalProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowinternalproperties := make(map[string]interface{})
	workflowinternalproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowinternalproperties["base_task_type"] = item.BaseTaskType
	workflowinternalproperties["class_id"] = item.ClassId
	workflowinternalproperties["constraints"] = (func(p models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
		var workflowtaskconstraintss []map[string]interface{}
		var ret models.WorkflowTaskConstraints
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		workflowtaskconstraints := make(map[string]interface{})
		workflowtaskconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		workflowtaskconstraints["class_id"] = item.ClassId
		workflowtaskconstraints["object_type"] = item.ObjectType
		workflowtaskconstraints["target_data_type"] = item.TargetDataType

		workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
		return workflowtaskconstraintss
	})(item.GetConstraints(), d)
	workflowinternalproperties["internal"] = item.Internal
	workflowinternalproperties["object_type"] = item.ObjectType
	workflowinternalproperties["owner"] = item.Owner

	workflowinternalpropertiess = append(workflowinternalpropertiess, workflowinternalproperties)
	return workflowinternalpropertiess
}
func flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(p models.WorkflowPendingDynamicWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowpendingdynamicworkflowinforelationships []map[string]interface{}
	var ret models.WorkflowPendingDynamicWorkflowInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowpendingdynamicworkflowinforelationship := make(map[string]interface{})
	workflowpendingdynamicworkflowinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowpendingdynamicworkflowinforelationship["class_id"] = item.ClassId
	workflowpendingdynamicworkflowinforelationship["moid"] = item.Moid
	workflowpendingdynamicworkflowinforelationship["object_type"] = item.ObjectType
	workflowpendingdynamicworkflowinforelationship["selector"] = item.Selector

	workflowpendingdynamicworkflowinforelationships = append(workflowpendingdynamicworkflowinforelationships, workflowpendingdynamicworkflowinforelationship)
	return workflowpendingdynamicworkflowinforelationships
}
func flattenMapWorkflowProperties(p models.WorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowpropertiess []map[string]interface{}
	var ret models.WorkflowProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowproperties := make(map[string]interface{})
	workflowproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowproperties["class_id"] = item.ClassId
	workflowproperties["external_meta"] = item.ExternalMeta
	workflowproperties["input_definition"] = (func(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowbasedatatype["class_id"] = item.ClassId
			workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.ClassId
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowbasedatatype["description"] = item.Description
			workflowbasedatatype["label"] = item.Label
			workflowbasedatatype["name"] = item.Name
			workflowbasedatatype["object_type"] = item.ObjectType
			workflowbasedatatype["required"] = item.Required
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.GetInputDefinition(), d)
	workflowproperties["object_type"] = item.ObjectType
	workflowproperties["output_definition"] = (func(p []models.WorkflowBaseDataType, d *schema.ResourceData) []map[string]interface{} {
		var workflowbasedatatypes []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowbasedatatype := make(map[string]interface{})
			workflowbasedatatype["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowbasedatatype["class_id"] = item.ClassId
			workflowbasedatatype["default"] = (func(p models.WorkflowDefaultValue, d *schema.ResourceData) []map[string]interface{} {
				var workflowdefaultvalues []map[string]interface{}
				var ret models.WorkflowDefaultValue
				if reflect.DeepEqual(ret, p) {
					return nil
				}
				item := p
				workflowdefaultvalue := make(map[string]interface{})
				workflowdefaultvalue["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
				workflowdefaultvalue["class_id"] = item.ClassId
				workflowdefaultvalue["object_type"] = item.ObjectType
				workflowdefaultvalue["override"] = item.Override
				workflowdefaultvalue["value"] = item.Value

				workflowdefaultvalues = append(workflowdefaultvalues, workflowdefaultvalue)
				return workflowdefaultvalues
			})(item.GetDefault(), d)
			workflowbasedatatype["description"] = item.Description
			workflowbasedatatype["label"] = item.Label
			workflowbasedatatype["name"] = item.Name
			workflowbasedatatype["object_type"] = item.ObjectType
			workflowbasedatatype["required"] = item.Required
			workflowbasedatatypes = append(workflowbasedatatypes, workflowbasedatatype)
		}
		return workflowbasedatatypes
	})(item.GetOutputDefinition(), d)
	workflowproperties["retry_count"] = item.RetryCount
	workflowproperties["retry_delay"] = item.RetryDelay
	workflowproperties["retry_policy"] = item.RetryPolicy
	workflowproperties["support_status"] = item.SupportStatus
	workflowproperties["timeout"] = item.Timeout
	workflowproperties["timeout_policy"] = item.TimeoutPolicy

	workflowpropertiess = append(workflowpropertiess, workflowproperties)
	return workflowpropertiess
}
func flattenMapWorkflowTaskConstraints(p models.WorkflowTaskConstraints, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskconstraintss []map[string]interface{}
	var ret models.WorkflowTaskConstraints
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowtaskconstraints := make(map[string]interface{})
	workflowtaskconstraints["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowtaskconstraints["class_id"] = item.ClassId
	workflowtaskconstraints["object_type"] = item.ObjectType
	workflowtaskconstraints["target_data_type"] = item.TargetDataType

	workflowtaskconstraintss = append(workflowtaskconstraintss, workflowtaskconstraints)
	return workflowtaskconstraintss
}
func flattenMapWorkflowTaskDefinitionRelationship(p models.WorkflowTaskDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskdefinitionrelationships []map[string]interface{}
	var ret models.WorkflowTaskDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowtaskdefinitionrelationship := make(map[string]interface{})
	workflowtaskdefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowtaskdefinitionrelationship["class_id"] = item.ClassId
	workflowtaskdefinitionrelationship["moid"] = item.Moid
	workflowtaskdefinitionrelationship["object_type"] = item.ObjectType
	workflowtaskdefinitionrelationship["selector"] = item.Selector

	workflowtaskdefinitionrelationships = append(workflowtaskdefinitionrelationships, workflowtaskdefinitionrelationship)
	return workflowtaskdefinitionrelationships
}
func flattenMapWorkflowTaskInfoRelationship(p models.WorkflowTaskInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowtaskinforelationships []map[string]interface{}
	var ret models.WorkflowTaskInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowtaskinforelationship := make(map[string]interface{})
	workflowtaskinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowtaskinforelationship["class_id"] = item.ClassId
	workflowtaskinforelationship["moid"] = item.Moid
	workflowtaskinforelationship["object_type"] = item.ObjectType
	workflowtaskinforelationship["selector"] = item.Selector

	workflowtaskinforelationships = append(workflowtaskinforelationships, workflowtaskinforelationship)
	return workflowtaskinforelationships
}
func flattenMapWorkflowValidationInformation(p models.WorkflowValidationInformation, d *schema.ResourceData) []map[string]interface{} {
	var workflowvalidationinformations []map[string]interface{}
	var ret models.WorkflowValidationInformation
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowvalidationinformation := make(map[string]interface{})
	workflowvalidationinformation["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowvalidationinformation["class_id"] = item.ClassId
	workflowvalidationinformation["object_type"] = item.ObjectType
	workflowvalidationinformation["state"] = item.State
	workflowvalidationinformation["validation_error"] = (func(p []models.WorkflowValidationError, d *schema.ResourceData) []map[string]interface{} {
		var workflowvalidationerrors []map[string]interface{}
		if len(p) == 0 {
			return nil
		}
		for _, item := range p {
			workflowvalidationerror := make(map[string]interface{})
			workflowvalidationerror["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
			workflowvalidationerror["class_id"] = item.ClassId
			workflowvalidationerror["error_log"] = item.ErrorLog
			workflowvalidationerror["field"] = item.Field
			workflowvalidationerror["object_type"] = item.ObjectType
			workflowvalidationerror["task_name"] = item.TaskName
			workflowvalidationerror["transition_name"] = item.TransitionName
			workflowvalidationerrors = append(workflowvalidationerrors, workflowvalidationerror)
		}
		return workflowvalidationerrors
	})(item.GetValidationError(), d)

	workflowvalidationinformations = append(workflowvalidationinformations, workflowvalidationinformation)
	return workflowvalidationinformations
}
func flattenMapWorkflowWorkflowDefinitionRelationship(p models.WorkflowWorkflowDefinitionRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowdefinitionrelationships []map[string]interface{}
	var ret models.WorkflowWorkflowDefinitionRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowworkflowdefinitionrelationship := make(map[string]interface{})
	workflowworkflowdefinitionrelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowdefinitionrelationship["class_id"] = item.ClassId
	workflowworkflowdefinitionrelationship["moid"] = item.Moid
	workflowworkflowdefinitionrelationship["object_type"] = item.ObjectType
	workflowworkflowdefinitionrelationship["selector"] = item.Selector

	workflowworkflowdefinitionrelationships = append(workflowworkflowdefinitionrelationships, workflowworkflowdefinitionrelationship)
	return workflowworkflowdefinitionrelationships
}
func flattenMapWorkflowWorkflowInfoRelationship(p models.WorkflowWorkflowInfoRelationship, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinforelationships []map[string]interface{}
	var ret models.WorkflowWorkflowInfoRelationship
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	x := p
	item := x.MoMoRef
	workflowworkflowinforelationship := make(map[string]interface{})
	workflowworkflowinforelationship["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowinforelationship["class_id"] = item.ClassId
	workflowworkflowinforelationship["moid"] = item.Moid
	workflowworkflowinforelationship["object_type"] = item.ObjectType
	workflowworkflowinforelationship["selector"] = item.Selector

	workflowworkflowinforelationships = append(workflowworkflowinforelationships, workflowworkflowinforelationship)
	return workflowworkflowinforelationships
}
func flattenMapWorkflowWorkflowInfoProperties(p models.WorkflowWorkflowInfoProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowinfopropertiess []map[string]interface{}
	var ret models.WorkflowWorkflowInfoProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowworkflowinfoproperties := make(map[string]interface{})
	workflowworkflowinfoproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowinfoproperties["class_id"] = item.ClassId
	workflowworkflowinfoproperties["object_type"] = item.ObjectType
	workflowworkflowinfoproperties["retryable"] = item.Retryable

	workflowworkflowinfopropertiess = append(workflowworkflowinfopropertiess, workflowworkflowinfoproperties)
	return workflowworkflowinfopropertiess
}
func flattenMapWorkflowWorkflowProperties(p models.WorkflowWorkflowProperties, d *schema.ResourceData) []map[string]interface{} {
	var workflowworkflowpropertiess []map[string]interface{}
	var ret models.WorkflowWorkflowProperties
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	workflowworkflowproperties := make(map[string]interface{})
	workflowworkflowproperties["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	workflowworkflowproperties["class_id"] = item.ClassId
	workflowworkflowproperties["external_meta"] = item.ExternalMeta
	workflowworkflowproperties["object_type"] = item.ObjectType
	workflowworkflowproperties["retryable"] = item.Retryable
	workflowworkflowproperties["support_status"] = item.SupportStatus

	workflowworkflowpropertiess = append(workflowworkflowpropertiess, workflowworkflowproperties)
	return workflowworkflowpropertiess
}
func flattenMapX509Certificate(p models.X509Certificate, d *schema.ResourceData) []map[string]interface{} {
	var x509certificates []map[string]interface{}
	var ret models.X509Certificate
	if reflect.DeepEqual(ret, p) {
		return nil
	}
	item := p
	x509certificate := make(map[string]interface{})
	x509certificate["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
	x509certificate["class_id"] = item.ClassId
	x509certificate["issuer"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		var ret models.PkixDistinguishedName
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		pkixdistinguishedname := make(map[string]interface{})
		pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		pkixdistinguishedname["class_id"] = item.ClassId
		pkixdistinguishedname["common_name"] = item.CommonName
		pkixdistinguishedname["country"] = item.Country
		pkixdistinguishedname["locality"] = item.Locality
		pkixdistinguishedname["object_type"] = item.ObjectType
		pkixdistinguishedname["organization"] = item.Organization
		pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
		pkixdistinguishedname["state"] = item.State

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.GetIssuer(), d)
	x509certificate["object_type"] = item.ObjectType
	x509certificate["pem_certificate"] = item.PemCertificate
	x509certificate["sha256_fingerprint"] = item.Sha256Fingerprint
	x509certificate["signature_algorithm"] = item.SignatureAlgorithm
	x509certificate["subject"] = (func(p models.PkixDistinguishedName, d *schema.ResourceData) []map[string]interface{} {
		var pkixdistinguishednames []map[string]interface{}
		var ret models.PkixDistinguishedName
		if reflect.DeepEqual(ret, p) {
			return nil
		}
		item := p
		pkixdistinguishedname := make(map[string]interface{})
		pkixdistinguishedname["additional_properties"] = flattenAdditionalProperties(item.AdditionalProperties)
		pkixdistinguishedname["class_id"] = item.ClassId
		pkixdistinguishedname["common_name"] = item.CommonName
		pkixdistinguishedname["country"] = item.Country
		pkixdistinguishedname["locality"] = item.Locality
		pkixdistinguishedname["object_type"] = item.ObjectType
		pkixdistinguishedname["organization"] = item.Organization
		pkixdistinguishedname["organizational_unit"] = item.OrganizationalUnit
		pkixdistinguishedname["state"] = item.State

		pkixdistinguishednames = append(pkixdistinguishednames, pkixdistinguishedname)
		return pkixdistinguishednames
	})(item.GetSubject(), d)

	x509certificates = append(x509certificates, x509certificate)
	return x509certificates
}
