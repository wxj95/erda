// Code generated by protoc-gen-go-form. DO NOT EDIT.
// Source: action.proto

package pb

import (
	json "encoding/json"
	url "net/url"
	strconv "strconv"
	strings "strings"

	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/urlenc" package it is being compiled against.
var _ urlenc.URLValuesUnmarshaler = (*Action)(nil)
var _ urlenc.URLValuesUnmarshaler = (*ActionNameWithVersionQuery)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionListRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionListResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionSaveRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionSaveResponse)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionDeleteRequest)(nil)
var _ urlenc.URLValuesUnmarshaler = (*PipelineActionDeleteResponse)(nil)

// Action implement urlenc.URLValuesUnmarshaler.
func (m *Action) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "ID":
				m.ID = vals[0]
			case "timeCreated":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
			case "timeCreated.seconds":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeCreated.Seconds = val
			case "timeCreated.nanos":
				if m.TimeCreated == nil {
					m.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeCreated.Nanos = int32(val)
			case "timeUpdated":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "timeUpdated.seconds":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.TimeUpdated.Seconds = val
			case "timeUpdated.nanos":
				if m.TimeUpdated == nil {
					m.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.TimeUpdated.Nanos = int32(val)
			case "name":
				m.Name = vals[0]
			case "category":
				m.Category = vals[0]
			case "display_name":
				m.DisplayName = vals[0]
			case "logo_url":
				m.LogoUrl = vals[0]
			case "desc":
				m.Desc = vals[0]
			case "readme":
				m.Readme = vals[0]
			case "dice":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Dice = val
					} else {
						m.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "spec":
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Spec = val
					} else {
						m.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "version":
				m.Version = vals[0]
			case "location":
				m.Location = vals[0]
			case "isPublic":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsPublic = val
			case "is_default":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDefault = val
			case "is_delete":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDelete = val
			case "softDeletedAt":
				if m.SoftDeletedAt == nil {
					m.SoftDeletedAt = &timestamppb.Timestamp{}
				}
			case "softDeletedAt.seconds":
				if m.SoftDeletedAt == nil {
					m.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.SoftDeletedAt.Seconds = val
			case "softDeletedAt.nanos":
				if m.SoftDeletedAt == nil {
					m.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.SoftDeletedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}

// ActionNameWithVersionQuery implement urlenc.URLValuesUnmarshaler.
func (m *ActionNameWithVersionQuery) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "version":
				m.Version = vals[0]
			case "is_default":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsDefault = val
			case "locationFilter":
				m.LocationFilter = vals[0]
			}
		}
	}
	return nil
}

// PipelineActionListRequest implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionListRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "categories":
				m.Categories = vals
			case "locations":
				m.Locations = vals
			case "is_public":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.IsPublic = val
			case "yamlFormat":
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.YamlFormat = val
			}
		}
	}
	return nil
}

// PipelineActionListResponse implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionListResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}

// PipelineActionSaveRequest implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionSaveRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "readme":
				m.Readme = vals[0]
			case "dice":
				m.Dice = vals[0]
			case "spec":
				m.Spec = vals[0]
			case "location":
				m.Location = vals[0]
			}
		}
	}
	return nil
}

// PipelineActionSaveResponse implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionSaveResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "action":
				if m.Action == nil {
					m.Action = &Action{}
				}
			case "action.ID":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.ID = vals[0]
			case "action.timeCreated":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.TimeCreated == nil {
					m.Action.TimeCreated = &timestamppb.Timestamp{}
				}
			case "action.timeCreated.seconds":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.TimeCreated == nil {
					m.Action.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Action.TimeCreated.Seconds = val
			case "action.timeCreated.nanos":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.TimeCreated == nil {
					m.Action.TimeCreated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Action.TimeCreated.Nanos = int32(val)
			case "action.timeUpdated":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.TimeUpdated == nil {
					m.Action.TimeUpdated = &timestamppb.Timestamp{}
				}
			case "action.timeUpdated.seconds":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.TimeUpdated == nil {
					m.Action.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Action.TimeUpdated.Seconds = val
			case "action.timeUpdated.nanos":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.TimeUpdated == nil {
					m.Action.TimeUpdated = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Action.TimeUpdated.Nanos = int32(val)
			case "action.name":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.Name = vals[0]
			case "action.category":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.Category = vals[0]
			case "action.display_name":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.DisplayName = vals[0]
			case "action.logo_url":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.LogoUrl = vals[0]
			case "action.desc":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.Desc = vals[0]
			case "action.readme":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.Readme = vals[0]
			case "action.dice":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Dice = val
					} else {
						m.Action.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "action.dice.null_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Dice = val
					} else {
						m.Action.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "action.dice.number_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Dice = val
					} else {
						m.Action.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "action.dice.string_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Dice = val
					} else {
						m.Action.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "action.dice.bool_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Dice = val
					} else {
						m.Action.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "action.dice.struct_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Dice = val
					} else {
						m.Action.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "action.dice.list_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Dice = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Dice = val
					} else {
						m.Action.Dice = structpb.NewStringValue(vals[0])
					}
				}
			case "action.spec":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Spec = val
					} else {
						m.Action.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "action.spec.null_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Spec = val
					} else {
						m.Action.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "action.spec.number_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Spec = val
					} else {
						m.Action.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "action.spec.string_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Spec = val
					} else {
						m.Action.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "action.spec.bool_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Spec = val
					} else {
						m.Action.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "action.spec.struct_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Spec = val
					} else {
						m.Action.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "action.spec.list_value":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if len(vals) > 1 {
					var list []interface{}
					for _, text := range vals {
						var v interface{}
						err := json.NewDecoder(strings.NewReader(text)).Decode(&v)
						if err != nil {
							list = append(list, v)
						} else {
							list = append(list, text)
						}
					}
					val, _ := structpb.NewList(list)
					m.Action.Spec = structpb.NewListValue(val)
				} else {
					var v interface{}
					err := json.NewDecoder(strings.NewReader(vals[0])).Decode(&v)
					if err != nil {
						val, _ := structpb.NewValue(v)
						m.Action.Spec = val
					} else {
						m.Action.Spec = structpb.NewStringValue(vals[0])
					}
				}
			case "action.version":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.Version = vals[0]
			case "action.location":
				if m.Action == nil {
					m.Action = &Action{}
				}
				m.Action.Location = vals[0]
			case "action.isPublic":
				if m.Action == nil {
					m.Action = &Action{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Action.IsPublic = val
			case "action.is_default":
				if m.Action == nil {
					m.Action = &Action{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Action.IsDefault = val
			case "action.is_delete":
				if m.Action == nil {
					m.Action = &Action{}
				}
				val, err := strconv.ParseBool(vals[0])
				if err != nil {
					return err
				}
				m.Action.IsDelete = val
			case "action.softDeletedAt":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.SoftDeletedAt == nil {
					m.Action.SoftDeletedAt = &timestamppb.Timestamp{}
				}
			case "action.softDeletedAt.seconds":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.SoftDeletedAt == nil {
					m.Action.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 64)
				if err != nil {
					return err
				}
				m.Action.SoftDeletedAt.Seconds = val
			case "action.softDeletedAt.nanos":
				if m.Action == nil {
					m.Action = &Action{}
				}
				if m.Action.SoftDeletedAt == nil {
					m.Action.SoftDeletedAt = &timestamppb.Timestamp{}
				}
				val, err := strconv.ParseInt(vals[0], 10, 32)
				if err != nil {
					return err
				}
				m.Action.SoftDeletedAt.Nanos = int32(val)
			}
		}
	}
	return nil
}

// PipelineActionDeleteRequest implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionDeleteRequest) UnmarshalURLValues(prefix string, values url.Values) error {
	for key, vals := range values {
		if len(vals) > 0 {
			switch prefix + key {
			case "name":
				m.Name = vals[0]
			case "version":
				m.Version = vals[0]
			case "location":
				m.Location = vals[0]
			}
		}
	}
	return nil
}

// PipelineActionDeleteResponse implement urlenc.URLValuesUnmarshaler.
func (m *PipelineActionDeleteResponse) UnmarshalURLValues(prefix string, values url.Values) error {
	return nil
}
