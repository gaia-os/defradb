// Copyright 2023 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package field

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestSchemaUpdatesAddSimpleErrorsAddingSchema(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add schema fails",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/-", "value": {"Name": "Books"} }
					]
				`,
				ExpectedError: "unknown collection, adding collections via patch is not supported. Name: Books",
			},
			testUtils.Request{
				Request: `query {
					Users {
						Name
					}
				}`,
				Results: []map[string]any{},
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesAddSimpleErrorsAddingCollectionProp(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add collection property fails",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/Users/-", "value": {"Name": "Books"} }
					]
				`,
				ExpectedError: `json: unknown field "-"`,
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesAddSimpleErrorsAddingSchemaProp(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add schema property fails",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/Users/Schema/-", "value": {"Foo": "Bar"} }
					]
				`,
				ExpectedError: `json: unknown field "-"`,
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesAddSimpleErrorsAddingUnsupportedCollectionProp(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add to unsupported collection prop",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/Users/Foo/Fields/-", "value": {"Name": "Email", "Kind": 11} }
					]
				`,
				ExpectedError: "add operation does not apply: doc is missing path",
			},
			testUtils.Request{
				Request: `query {
					Users {
						Name
					}
				}`,
				Results: []map[string]any{},
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesAddSimpleErrorsAddingUnsupportedSchemaProp(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add to unsupported schema prop",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/Users/Schema/Foo/-", "value": {"Name": "Email", "Kind": 11} }
					]
				`,
				ExpectedError: "add operation does not apply: doc is missing path",
			},
			testUtils.Request{
				Request: `query {
					Users {
						Name
					}
				}`,
				Results: []map[string]any{},
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}
