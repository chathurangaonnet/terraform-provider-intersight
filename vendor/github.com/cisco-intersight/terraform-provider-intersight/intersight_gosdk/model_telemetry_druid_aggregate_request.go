/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-08T07:46:03Z.
 *
 * API version: 0.0.1-15983
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"fmt"
)

// TelemetryDruidAggregateRequest - Exposes a REST endpoint for performing queries against Druid time series data. View Telemetry allows [POST of a Druid query](http://druid.io/docs/latest/querying/querying). Manage Telemetry allows [READ of broker status](http://druid.io/docs/latest/operations/api-reference.html#broker).
type TelemetryDruidAggregateRequest struct {
	TelemetryDruidDataSourceMetadataRequest *TelemetryDruidDataSourceMetadataRequest
	TelemetryDruidGroupByRequest            *TelemetryDruidGroupByRequest
	TelemetryDruidScanRequest               *TelemetryDruidScanRequest
	TelemetryDruidSegmentMetadataRequest    *TelemetryDruidSegmentMetadataRequest
	TelemetryDruidTimeBoundaryRequest       *TelemetryDruidTimeBoundaryRequest
	TelemetryDruidTimeSeriesRequest         *TelemetryDruidTimeSeriesRequest
	TelemetryDruidTopNRequest               *TelemetryDruidTopNRequest
}

// TelemetryDruidDataSourceMetadataRequestAsTelemetryDruidAggregateRequest is a convenience function that returns TelemetryDruidDataSourceMetadataRequest wrapped in TelemetryDruidAggregateRequest
func TelemetryDruidDataSourceMetadataRequestAsTelemetryDruidAggregateRequest(v *TelemetryDruidDataSourceMetadataRequest) TelemetryDruidAggregateRequest {
	return TelemetryDruidAggregateRequest{TelemetryDruidDataSourceMetadataRequest: v}
}

// TelemetryDruidGroupByRequestAsTelemetryDruidAggregateRequest is a convenience function that returns TelemetryDruidGroupByRequest wrapped in TelemetryDruidAggregateRequest
func TelemetryDruidGroupByRequestAsTelemetryDruidAggregateRequest(v *TelemetryDruidGroupByRequest) TelemetryDruidAggregateRequest {
	return TelemetryDruidAggregateRequest{TelemetryDruidGroupByRequest: v}
}

// TelemetryDruidScanRequestAsTelemetryDruidAggregateRequest is a convenience function that returns TelemetryDruidScanRequest wrapped in TelemetryDruidAggregateRequest
func TelemetryDruidScanRequestAsTelemetryDruidAggregateRequest(v *TelemetryDruidScanRequest) TelemetryDruidAggregateRequest {
	return TelemetryDruidAggregateRequest{TelemetryDruidScanRequest: v}
}

// TelemetryDruidSegmentMetadataRequestAsTelemetryDruidAggregateRequest is a convenience function that returns TelemetryDruidSegmentMetadataRequest wrapped in TelemetryDruidAggregateRequest
func TelemetryDruidSegmentMetadataRequestAsTelemetryDruidAggregateRequest(v *TelemetryDruidSegmentMetadataRequest) TelemetryDruidAggregateRequest {
	return TelemetryDruidAggregateRequest{TelemetryDruidSegmentMetadataRequest: v}
}

// TelemetryDruidTimeBoundaryRequestAsTelemetryDruidAggregateRequest is a convenience function that returns TelemetryDruidTimeBoundaryRequest wrapped in TelemetryDruidAggregateRequest
func TelemetryDruidTimeBoundaryRequestAsTelemetryDruidAggregateRequest(v *TelemetryDruidTimeBoundaryRequest) TelemetryDruidAggregateRequest {
	return TelemetryDruidAggregateRequest{TelemetryDruidTimeBoundaryRequest: v}
}

// TelemetryDruidTimeSeriesRequestAsTelemetryDruidAggregateRequest is a convenience function that returns TelemetryDruidTimeSeriesRequest wrapped in TelemetryDruidAggregateRequest
func TelemetryDruidTimeSeriesRequestAsTelemetryDruidAggregateRequest(v *TelemetryDruidTimeSeriesRequest) TelemetryDruidAggregateRequest {
	return TelemetryDruidAggregateRequest{TelemetryDruidTimeSeriesRequest: v}
}

// TelemetryDruidTopNRequestAsTelemetryDruidAggregateRequest is a convenience function that returns TelemetryDruidTopNRequest wrapped in TelemetryDruidAggregateRequest
func TelemetryDruidTopNRequestAsTelemetryDruidAggregateRequest(v *TelemetryDruidTopNRequest) TelemetryDruidAggregateRequest {
	return TelemetryDruidAggregateRequest{TelemetryDruidTopNRequest: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidAggregateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'dataSourceMetadata'
	if jsonDict["queryType"] == "dataSourceMetadata" {
		// try to unmarshal JSON data into TelemetryDruidDataSourceMetadataRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidDataSourceMetadataRequest)
		if err == nil {
			jsonTelemetryDruidDataSourceMetadataRequest, _ := json.Marshal(dst.TelemetryDruidDataSourceMetadataRequest)
			if string(jsonTelemetryDruidDataSourceMetadataRequest) == "{}" { // empty struct
				dst.TelemetryDruidDataSourceMetadataRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidDataSourceMetadataRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidDataSourceMetadataRequest = nil
		}
	}

	// check if the discriminator value is 'groupBy'
	if jsonDict["queryType"] == "groupBy" {
		// try to unmarshal JSON data into TelemetryDruidGroupByRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidGroupByRequest)
		if err == nil {
			jsonTelemetryDruidGroupByRequest, _ := json.Marshal(dst.TelemetryDruidGroupByRequest)
			if string(jsonTelemetryDruidGroupByRequest) == "{}" { // empty struct
				dst.TelemetryDruidGroupByRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidGroupByRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidGroupByRequest = nil
		}
	}

	// check if the discriminator value is 'scan'
	if jsonDict["queryType"] == "scan" {
		// try to unmarshal JSON data into TelemetryDruidScanRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidScanRequest)
		if err == nil {
			jsonTelemetryDruidScanRequest, _ := json.Marshal(dst.TelemetryDruidScanRequest)
			if string(jsonTelemetryDruidScanRequest) == "{}" { // empty struct
				dst.TelemetryDruidScanRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidScanRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidScanRequest = nil
		}
	}

	// check if the discriminator value is 'segmentMetadata'
	if jsonDict["queryType"] == "segmentMetadata" {
		// try to unmarshal JSON data into TelemetryDruidSegmentMetadataRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidSegmentMetadataRequest)
		if err == nil {
			jsonTelemetryDruidSegmentMetadataRequest, _ := json.Marshal(dst.TelemetryDruidSegmentMetadataRequest)
			if string(jsonTelemetryDruidSegmentMetadataRequest) == "{}" { // empty struct
				dst.TelemetryDruidSegmentMetadataRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidSegmentMetadataRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidSegmentMetadataRequest = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidDataSourceMetadataRequest'
	if jsonDict["queryType"] == "telemetry.DruidDataSourceMetadataRequest" {
		// try to unmarshal JSON data into TelemetryDruidDataSourceMetadataRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidDataSourceMetadataRequest)
		if err == nil {
			jsonTelemetryDruidDataSourceMetadataRequest, _ := json.Marshal(dst.TelemetryDruidDataSourceMetadataRequest)
			if string(jsonTelemetryDruidDataSourceMetadataRequest) == "{}" { // empty struct
				dst.TelemetryDruidDataSourceMetadataRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidDataSourceMetadataRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidDataSourceMetadataRequest = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidGroupByRequest'
	if jsonDict["queryType"] == "telemetry.DruidGroupByRequest" {
		// try to unmarshal JSON data into TelemetryDruidGroupByRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidGroupByRequest)
		if err == nil {
			jsonTelemetryDruidGroupByRequest, _ := json.Marshal(dst.TelemetryDruidGroupByRequest)
			if string(jsonTelemetryDruidGroupByRequest) == "{}" { // empty struct
				dst.TelemetryDruidGroupByRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidGroupByRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidGroupByRequest = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidScanRequest'
	if jsonDict["queryType"] == "telemetry.DruidScanRequest" {
		// try to unmarshal JSON data into TelemetryDruidScanRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidScanRequest)
		if err == nil {
			jsonTelemetryDruidScanRequest, _ := json.Marshal(dst.TelemetryDruidScanRequest)
			if string(jsonTelemetryDruidScanRequest) == "{}" { // empty struct
				dst.TelemetryDruidScanRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidScanRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidScanRequest = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidSegmentMetadataRequest'
	if jsonDict["queryType"] == "telemetry.DruidSegmentMetadataRequest" {
		// try to unmarshal JSON data into TelemetryDruidSegmentMetadataRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidSegmentMetadataRequest)
		if err == nil {
			jsonTelemetryDruidSegmentMetadataRequest, _ := json.Marshal(dst.TelemetryDruidSegmentMetadataRequest)
			if string(jsonTelemetryDruidSegmentMetadataRequest) == "{}" { // empty struct
				dst.TelemetryDruidSegmentMetadataRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidSegmentMetadataRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidSegmentMetadataRequest = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidTimeBoundaryRequest'
	if jsonDict["queryType"] == "telemetry.DruidTimeBoundaryRequest" {
		// try to unmarshal JSON data into TelemetryDruidTimeBoundaryRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidTimeBoundaryRequest)
		if err == nil {
			jsonTelemetryDruidTimeBoundaryRequest, _ := json.Marshal(dst.TelemetryDruidTimeBoundaryRequest)
			if string(jsonTelemetryDruidTimeBoundaryRequest) == "{}" { // empty struct
				dst.TelemetryDruidTimeBoundaryRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidTimeBoundaryRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidTimeBoundaryRequest = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidTimeSeriesRequest'
	if jsonDict["queryType"] == "telemetry.DruidTimeSeriesRequest" {
		// try to unmarshal JSON data into TelemetryDruidTimeSeriesRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidTimeSeriesRequest)
		if err == nil {
			jsonTelemetryDruidTimeSeriesRequest, _ := json.Marshal(dst.TelemetryDruidTimeSeriesRequest)
			if string(jsonTelemetryDruidTimeSeriesRequest) == "{}" { // empty struct
				dst.TelemetryDruidTimeSeriesRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidTimeSeriesRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidTimeSeriesRequest = nil
		}
	}

	// check if the discriminator value is 'telemetry.DruidTopNRequest'
	if jsonDict["queryType"] == "telemetry.DruidTopNRequest" {
		// try to unmarshal JSON data into TelemetryDruidTopNRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidTopNRequest)
		if err == nil {
			jsonTelemetryDruidTopNRequest, _ := json.Marshal(dst.TelemetryDruidTopNRequest)
			if string(jsonTelemetryDruidTopNRequest) == "{}" { // empty struct
				dst.TelemetryDruidTopNRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidTopNRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidTopNRequest = nil
		}
	}

	// check if the discriminator value is 'timeBoundary'
	if jsonDict["queryType"] == "timeBoundary" {
		// try to unmarshal JSON data into TelemetryDruidTimeBoundaryRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidTimeBoundaryRequest)
		if err == nil {
			jsonTelemetryDruidTimeBoundaryRequest, _ := json.Marshal(dst.TelemetryDruidTimeBoundaryRequest)
			if string(jsonTelemetryDruidTimeBoundaryRequest) == "{}" { // empty struct
				dst.TelemetryDruidTimeBoundaryRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidTimeBoundaryRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidTimeBoundaryRequest = nil
		}
	}

	// check if the discriminator value is 'timeseries'
	if jsonDict["queryType"] == "timeseries" {
		// try to unmarshal JSON data into TelemetryDruidTimeSeriesRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidTimeSeriesRequest)
		if err == nil {
			jsonTelemetryDruidTimeSeriesRequest, _ := json.Marshal(dst.TelemetryDruidTimeSeriesRequest)
			if string(jsonTelemetryDruidTimeSeriesRequest) == "{}" { // empty struct
				dst.TelemetryDruidTimeSeriesRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidTimeSeriesRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidTimeSeriesRequest = nil
		}
	}

	// check if the discriminator value is 'topN'
	if jsonDict["queryType"] == "topN" {
		// try to unmarshal JSON data into TelemetryDruidTopNRequest
		err = json.Unmarshal(data, &dst.TelemetryDruidTopNRequest)
		if err == nil {
			jsonTelemetryDruidTopNRequest, _ := json.Marshal(dst.TelemetryDruidTopNRequest)
			if string(jsonTelemetryDruidTopNRequest) == "{}" { // empty struct
				dst.TelemetryDruidTopNRequest = nil
			} else {
				return nil // data stored in dst.TelemetryDruidTopNRequest, return on the first match
			}
		} else {
			dst.TelemetryDruidTopNRequest = nil
		}
	}

	// try to unmarshal data into TelemetryDruidDataSourceMetadataRequest
	err = json.Unmarshal(data, &dst.TelemetryDruidDataSourceMetadataRequest)
	if err == nil {
		jsonTelemetryDruidDataSourceMetadataRequest, _ := json.Marshal(dst.TelemetryDruidDataSourceMetadataRequest)
		if string(jsonTelemetryDruidDataSourceMetadataRequest) == "{}" { // empty struct
			dst.TelemetryDruidDataSourceMetadataRequest = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidDataSourceMetadataRequest = nil
	}

	// try to unmarshal data into TelemetryDruidGroupByRequest
	err = json.Unmarshal(data, &dst.TelemetryDruidGroupByRequest)
	if err == nil {
		jsonTelemetryDruidGroupByRequest, _ := json.Marshal(dst.TelemetryDruidGroupByRequest)
		if string(jsonTelemetryDruidGroupByRequest) == "{}" { // empty struct
			dst.TelemetryDruidGroupByRequest = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidGroupByRequest = nil
	}

	// try to unmarshal data into TelemetryDruidScanRequest
	err = json.Unmarshal(data, &dst.TelemetryDruidScanRequest)
	if err == nil {
		jsonTelemetryDruidScanRequest, _ := json.Marshal(dst.TelemetryDruidScanRequest)
		if string(jsonTelemetryDruidScanRequest) == "{}" { // empty struct
			dst.TelemetryDruidScanRequest = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidScanRequest = nil
	}

	// try to unmarshal data into TelemetryDruidSegmentMetadataRequest
	err = json.Unmarshal(data, &dst.TelemetryDruidSegmentMetadataRequest)
	if err == nil {
		jsonTelemetryDruidSegmentMetadataRequest, _ := json.Marshal(dst.TelemetryDruidSegmentMetadataRequest)
		if string(jsonTelemetryDruidSegmentMetadataRequest) == "{}" { // empty struct
			dst.TelemetryDruidSegmentMetadataRequest = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidSegmentMetadataRequest = nil
	}

	// try to unmarshal data into TelemetryDruidTimeBoundaryRequest
	err = json.Unmarshal(data, &dst.TelemetryDruidTimeBoundaryRequest)
	if err == nil {
		jsonTelemetryDruidTimeBoundaryRequest, _ := json.Marshal(dst.TelemetryDruidTimeBoundaryRequest)
		if string(jsonTelemetryDruidTimeBoundaryRequest) == "{}" { // empty struct
			dst.TelemetryDruidTimeBoundaryRequest = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidTimeBoundaryRequest = nil
	}

	// try to unmarshal data into TelemetryDruidTimeSeriesRequest
	err = json.Unmarshal(data, &dst.TelemetryDruidTimeSeriesRequest)
	if err == nil {
		jsonTelemetryDruidTimeSeriesRequest, _ := json.Marshal(dst.TelemetryDruidTimeSeriesRequest)
		if string(jsonTelemetryDruidTimeSeriesRequest) == "{}" { // empty struct
			dst.TelemetryDruidTimeSeriesRequest = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidTimeSeriesRequest = nil
	}

	// try to unmarshal data into TelemetryDruidTopNRequest
	err = json.Unmarshal(data, &dst.TelemetryDruidTopNRequest)
	if err == nil {
		jsonTelemetryDruidTopNRequest, _ := json.Marshal(dst.TelemetryDruidTopNRequest)
		if string(jsonTelemetryDruidTopNRequest) == "{}" { // empty struct
			dst.TelemetryDruidTopNRequest = nil
		} else {
			match++
		}
	} else {
		dst.TelemetryDruidTopNRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TelemetryDruidDataSourceMetadataRequest = nil
		dst.TelemetryDruidGroupByRequest = nil
		dst.TelemetryDruidScanRequest = nil
		dst.TelemetryDruidSegmentMetadataRequest = nil
		dst.TelemetryDruidTimeBoundaryRequest = nil
		dst.TelemetryDruidTimeSeriesRequest = nil
		dst.TelemetryDruidTopNRequest = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TelemetryDruidAggregateRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TelemetryDruidAggregateRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidAggregateRequest) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidDataSourceMetadataRequest != nil {
		return json.Marshal(&src.TelemetryDruidDataSourceMetadataRequest)
	}

	if src.TelemetryDruidGroupByRequest != nil {
		return json.Marshal(&src.TelemetryDruidGroupByRequest)
	}

	if src.TelemetryDruidScanRequest != nil {
		return json.Marshal(&src.TelemetryDruidScanRequest)
	}

	if src.TelemetryDruidSegmentMetadataRequest != nil {
		return json.Marshal(&src.TelemetryDruidSegmentMetadataRequest)
	}

	if src.TelemetryDruidTimeBoundaryRequest != nil {
		return json.Marshal(&src.TelemetryDruidTimeBoundaryRequest)
	}

	if src.TelemetryDruidTimeSeriesRequest != nil {
		return json.Marshal(&src.TelemetryDruidTimeSeriesRequest)
	}

	if src.TelemetryDruidTopNRequest != nil {
		return json.Marshal(&src.TelemetryDruidTopNRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidAggregateRequest) GetActualInstance() interface{} {
	if obj.TelemetryDruidDataSourceMetadataRequest != nil {
		return obj.TelemetryDruidDataSourceMetadataRequest
	}

	if obj.TelemetryDruidGroupByRequest != nil {
		return obj.TelemetryDruidGroupByRequest
	}

	if obj.TelemetryDruidScanRequest != nil {
		return obj.TelemetryDruidScanRequest
	}

	if obj.TelemetryDruidSegmentMetadataRequest != nil {
		return obj.TelemetryDruidSegmentMetadataRequest
	}

	if obj.TelemetryDruidTimeBoundaryRequest != nil {
		return obj.TelemetryDruidTimeBoundaryRequest
	}

	if obj.TelemetryDruidTimeSeriesRequest != nil {
		return obj.TelemetryDruidTimeSeriesRequest
	}

	if obj.TelemetryDruidTopNRequest != nil {
		return obj.TelemetryDruidTopNRequest
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidAggregateRequest struct {
	value *TelemetryDruidAggregateRequest
	isSet bool
}

func (v NullableTelemetryDruidAggregateRequest) Get() *TelemetryDruidAggregateRequest {
	return v.value
}

func (v *NullableTelemetryDruidAggregateRequest) Set(val *TelemetryDruidAggregateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidAggregateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidAggregateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidAggregateRequest(val *TelemetryDruidAggregateRequest) *NullableTelemetryDruidAggregateRequest {
	return &NullableTelemetryDruidAggregateRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidAggregateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidAggregateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
