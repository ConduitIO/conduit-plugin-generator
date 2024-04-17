// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-connector-sdk/tree/main/cmd/paramgen

package generator

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (Config) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"burst.generateTime": {
			Default:     "1s",
			Description: "The amount of time the generator is generating records in a burst. Has an effect only if `burst.sleepTime` is set.",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"burst.sleepTime": {
			Default:     "",
			Description: "The time the generator \"sleeps\" between bursts.",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"collections.*.format.options.*": {
			Default:     "",
			Description: "The options for the format type selected, which are:   1. For raw and structured: pairs of field names and field types, where the type can be one of: int, string, time, bool.   2. For the file format: a path to the file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"collections.*.format.options.path": {
			Default:     "",
			Description: "Path to the input file (only applicable if the format type is file).",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"collections.*.format.type": {
			Default:     "",
			Description: "The format of the generated payload data (raw, structured, file).",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"raw", "structured", "file"}},
			},
		},
		"collections.*.operation": {
			Default:     "create",
			Description: "Comma separated list of record operations to generate. Allowed values are \"create\", \"update\", \"delete\", \"snapshot\".",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"format.options.*": {
			Default:     "",
			Description: "The options for the format type selected, which are:   1. For raw and structured: pairs of field names and field types, where the type can be one of: int, string, time, bool.   2. For the file format: a path to the file.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"format.options.path": {
			Default:     "",
			Description: "Path to the input file (only applicable if the format type is file).",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"format.type": {
			Default:     "",
			Description: "The format of the generated payload data (raw, structured, file).",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"raw", "structured", "file"}},
			},
		},
		"operation": {
			Default:     "create",
			Description: "Comma separated list of record operations to generate. Allowed values are \"create\", \"update\", \"delete\", \"snapshot\".",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"rate": {
			Default:     "",
			Description: "The maximum rate in records per second, at which records are generated (0 means no rate limit).",
			Type:        sdk.ParameterTypeFloat,
			Validations: []sdk.Validation{},
		},
		"readTime": {
			Default:     "",
			Description: "The time it takes to 'read' a record. Deprecated: use `rate` instead.",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"recordCount": {
			Default:     "",
			Description: "Number of records to be generated (0 means infinite).",
			Type:        sdk.ParameterTypeInt,
			Validations: []sdk.Validation{
				sdk.ValidationGreaterThan{Value: -1},
			},
		},
	}
}
