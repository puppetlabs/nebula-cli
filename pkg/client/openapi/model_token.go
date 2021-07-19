/*
 * Relay API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v20200615
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Token - An access token
type Token struct {
	UserToken *UserToken
}

// UserTokenAsToken is a convenience function that returns UserToken wrapped in Token
func UserTokenAsToken(v *UserToken) Token {
	return Token{ UserToken: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Token) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UserToken
	err = json.Unmarshal(data, &dst.UserToken)
	if err == nil {
		jsonUserToken, _ := json.Marshal(dst.UserToken)
		if string(jsonUserToken) == "{}" { // empty struct
			dst.UserToken = nil
		} else {
			match++
		}
	} else {
		dst.UserToken = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UserToken = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(Token)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(Token)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Token) MarshalJSON() ([]byte, error) {
	if src.UserToken != nil {
		return json.Marshal(&src.UserToken)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Token) GetActualInstance() (interface{}) {
	if obj.UserToken != nil {
		return obj.UserToken
	}

	// all schemas are nil
	return nil
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

