// Code generated by zenrpc; DO NOT EDIT.

package testdata

import (
	"context"
	"encoding/json"

	"github.com/vmkteam/zenrpc/v2"
	"github.com/vmkteam/zenrpc/v2/smd"

	"github.com/vmkteam/zenrpc/v2/testdata/model"
)

var RPC = struct {
	ArithService     struct{ Sum, Positive, DoSomething, GetPoints, DoSomethingWithPoint, Multiply, CheckError, CheckZenRPCError, Divide, Pow, Pi, SumArray string }
	CatalogueService struct{ First, Second, Third string }
	PhoneBook        struct{ Get, ValidateSearch, ById, Delete, Remove, Save string }
	PrintService     struct{ PrintRequiredDefault, PrintOptionalWithDefault, PrintRequired, PrintOptional string }
}{
	ArithService: struct{ Sum, Positive, DoSomething, GetPoints, DoSomethingWithPoint, Multiply, CheckError, CheckZenRPCError, Divide, Pow, Pi, SumArray string }{
		Sum:                  "sum",
		Positive:             "positive",
		DoSomething:          "dosomething",
		GetPoints:            "getpoints",
		DoSomethingWithPoint: "dosomethingwithpoint",
		Multiply:             "multiply",
		CheckError:           "checkerror",
		CheckZenRPCError:     "checkzenrpcerror",
		Divide:               "divide",
		Pow:                  "pow",
		Pi:                   "pi",
		SumArray:             "sumarray",
	},
	CatalogueService: struct{ First, Second, Third string }{
		First:  "first",
		Second: "second",
		Third:  "third",
	},
	PhoneBook: struct{ Get, ValidateSearch, ById, Delete, Remove, Save string }{
		Get:            "get",
		ValidateSearch: "validatesearch",
		ById:           "byid",
		Delete:         "delete",
		Remove:         "remove",
		Save:           "save",
	},
	PrintService: struct{ PrintRequiredDefault, PrintOptionalWithDefault, PrintRequired, PrintOptional string }{
		PrintRequiredDefault:     "printrequireddefault",
		PrintOptionalWithDefault: "printoptionalwithdefault",
		PrintRequired:            "printrequired",
		PrintOptional:            "printoptional",
	},
}

func (ArithService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"Sum": {
				Description: `Sum sums two digits and returns error with error code as result and IP from context.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "a",
						Type: smd.Integer,
					},
					{
						Name: "b",
						Type: smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"Positive": {
				Parameters: []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"DoSomething": {
				Parameters: []smd.JSONSchema{},
			},
			"GetPoints": {
				Parameters: []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type: smd.Array,
					Items: map[string]string{
						"$ref": "#/definitions/model.Point",
					},
					Definitions: map[string]smd.Definition{
						"model.Point": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name:        "X",
									Description: `coordinate`,
									Type:        smd.Integer,
								},
								{
									Name:        "Y",
									Description: `coordinate`,
									Type:        smd.Integer,
								},
								{
									Name: "ConnectedObject",
									Ref:  "#/definitions/objects.AbstractObject",
									Type: smd.Object,
								},
							},
						},
						"objects.AbstractObject": {
							Type:       "object",
							Properties: smd.PropertyList{},
						},
					},
				},
			},
			"DoSomethingWithPoint": {
				Parameters: []smd.JSONSchema{
					{
						Name: "p",
						Type: smd.Object,
						Properties: smd.PropertyList{
							{
								Name:        "X",
								Description: `coordinate`,
								Type:        smd.Integer,
							},
							{
								Name:        "Y",
								Description: `coordinate`,
								Type:        smd.Integer,
							},
							{
								Name: "ConnectedObject",
								Ref:  "#/definitions/objects.AbstractObject",
								Type: smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"objects.AbstractObject": {
								Type:       "object",
								Properties: smd.PropertyList{},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Object,
					Properties: smd.PropertyList{
						{
							Name:        "X",
							Description: `coordinate`,
							Type:        smd.Integer,
						},
						{
							Name:        "Y",
							Description: `coordinate`,
							Type:        smd.Integer,
						},
						{
							Name: "ConnectedObject",
							Ref:  "#/definitions/objects.AbstractObject",
							Type: smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"objects.AbstractObject": {
							Type:       "object",
							Properties: smd.PropertyList{},
						},
					},
				},
			},
			"Multiply": {
				Description: `Multiply multiples two digits and returns result.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "a",
						Type: smd.Integer,
					},
					{
						Name: "b",
						Type: smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Integer,
				},
			},
			"CheckError": {
				Description: `CheckError throws error is isErr true.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "isErr",
						Type: smd.Boolean,
					},
				},
				Errors: map[int]string{
					500: "test error",
				},
			},
			"CheckZenRPCError": {
				Description: `CheckError throws zenrpc error is isErr true.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "isErr",
						Type: smd.Boolean,
					},
				},
				Errors: map[int]string{
					500: "test error",
				},
			},
			"Divide": {
				Description: `Divide divides two numbers.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "a",
						Description: `the a`,
						Type:        smd.Integer,
					},
					{
						Name:        "b",
						Description: `the b`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
					Properties: smd.PropertyList{
						{
							Name:        "Quo",
							Description: `Quo docs`,
							Type:        smd.Integer,
						},
						{
							Name:        "rem",
							Description: `Rem docs`,
							Type:        smd.Integer,
						},
					},
				},
				Errors: map[int]string{
					401:    "we do not serve 1",
					-32603: "divide by zero",
				},
			},
			"Pow": {
				Description: `Pow returns x**y, the base-x exponential of y. If Exp is not set then default value is 2.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "base",
						Type: smd.Float,
					},
					{
						Name:        "exp",
						Optional:    true,
						Description: `exponent could be empty`,
						Type:        smd.Float,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Float,
				},
			},
			"Pi": {
				Description: `PI returns math.Pi.`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type: smd.Float,
				},
			},
			"SumArray": {
				Description: `SumArray returns sum all items from array`,
				Parameters: []smd.JSONSchema{
					{
						Name:     "array",
						Optional: true,
						Type:     smd.Array,
						Items: map[string]string{
							"type": smd.Float,
						},
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Float,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (as ArithService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.ArithService.Sum:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(as.Sum(ctx, args.A, args.B))

	case RPC.ArithService.Positive:
		resp.Set(as.Positive())

	case RPC.ArithService.DoSomething:
		as.DoSomething()

	case RPC.ArithService.GetPoints:
		resp.Set(as.GetPoints())

	case RPC.ArithService.DoSomethingWithPoint:
		var args = struct {
			P model.Point `json:"p"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"p"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(as.DoSomethingWithPoint(args.P))

	case RPC.ArithService.Multiply:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(as.Multiply(args.A, args.B))

	case RPC.ArithService.CheckError:
		var args = struct {
			IsErr bool `json:"isErr"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"isErr"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(as.CheckError(args.IsErr))

	case RPC.ArithService.CheckZenRPCError:
		var args = struct {
			IsErr bool `json:"isErr"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"isErr"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(as.CheckZenRPCError(args.IsErr))

	case RPC.ArithService.Divide:
		var args = struct {
			A int `json:"a"`
			B int `json:"b"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"a", "b"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(as.Divide(args.A, args.B))

	case RPC.ArithService.Pow:
		var args = struct {
			Base float64  `json:"base"`
			Exp  *float64 `json:"exp"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"base", "exp"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:exp=2 	exponent could be empty
		if args.Exp == nil {
			var v float64 = 2
			args.Exp = &v
		}

		resp.Set(as.Pow(args.Base, args.Exp))

	case RPC.ArithService.Pi:
		resp.Set(as.Pi())

	case RPC.ArithService.SumArray:
		var args = struct {
			Array *[]float64 `json:"array"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"array"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:array=[]float64{1,2,4}
		if args.Array == nil {
			var v []float64 = []float64{1, 2, 4}
			args.Array = &v
		}

		resp.Set(as.SumArray(args.Array))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (CatalogueService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"First": {
				Parameters: []smd.JSONSchema{
					{
						Name: "groups",
						Type: smd.Array,
						Items: map[string]string{
							"$ref": "#/definitions/Group",
						},
						Definitions: map[string]smd.Definition{
							"Group": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "title",
										Type: smd.String,
									},
									{
										Name: "nodes",
										Type: smd.Array,
										Items: map[string]string{
											"$ref": "#/definitions/Group",
										},
									},
									{
										Name: "groups",
										Type: smd.Array,
										Items: map[string]string{
											"$ref": "#/definitions/Group",
										},
									},
									{
										Name:     "child",
										Optional: true,
										Ref:      "#/definitions/Group",
										Type:     smd.Object,
									},
									{
										Name: "sub",
										Ref:  "#/definitions/SubGroup",
										Type: smd.Object,
									},
								},
							},
							"SubGroup": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "title",
										Type: smd.String,
									},
									{
										Name: "nodes",
										Type: smd.Array,
										Items: map[string]string{
											"$ref": "#/definitions/Group",
										},
									},
								},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"Second": {
				Parameters: []smd.JSONSchema{
					{
						Name: "campaigns",
						Type: smd.Array,
						Items: map[string]string{
							"$ref": "#/definitions/Campaign",
						},
						Definitions: map[string]smd.Definition{
							"Campaign": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "groups",
										Type: smd.Array,
										Items: map[string]string{
											"$ref": "#/definitions/Group",
										},
									},
								},
							},
							"Group": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "title",
										Type: smd.String,
									},
									{
										Name: "nodes",
										Type: smd.Array,
										Items: map[string]string{
											"$ref": "#/definitions/Group",
										},
									},
									{
										Name: "groups",
										Type: smd.Array,
										Items: map[string]string{
											"$ref": "#/definitions/Group",
										},
									},
									{
										Name:     "child",
										Optional: true,
										Ref:      "#/definitions/Group",
										Type:     smd.Object,
									},
									{
										Name: "sub",
										Ref:  "#/definitions/SubGroup",
										Type: smd.Object,
									},
								},
							},
							"SubGroup": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "id",
										Type: smd.Integer,
									},
									{
										Name: "title",
										Type: smd.String,
									},
									{
										Name: "nodes",
										Type: smd.Array,
										Items: map[string]string{
											"$ref": "#/definitions/Group",
										},
									},
								},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"Third": {
				Parameters: []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Type: smd.Object,
					Properties: smd.PropertyList{
						{
							Name: "id",
							Type: smd.Integer,
						},
						{
							Name: "groups",
							Type: smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/Group",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"Group": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "title",
									Type: smd.String,
								},
								{
									Name: "nodes",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Group",
									},
								},
								{
									Name: "groups",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Group",
									},
								},
								{
									Name:     "child",
									Optional: true,
									Ref:      "#/definitions/Group",
									Type:     smd.Object,
								},
								{
									Name: "sub",
									Ref:  "#/definitions/SubGroup",
									Type: smd.Object,
								},
							},
						},
						"SubGroup": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "id",
									Type: smd.Integer,
								},
								{
									Name: "title",
									Type: smd.String,
								},
								{
									Name: "nodes",
									Type: smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Group",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s CatalogueService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.CatalogueService.First:
		var args = struct {
			Groups Groups `json:"groups"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"groups"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.First(args.Groups))

	case RPC.CatalogueService.Second:
		var args = struct {
			Campaigns []Campaign `json:"campaigns"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"campaigns"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Second(args.Campaigns))

	case RPC.CatalogueService.Third:
		resp.Set(s.Third())

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (PhoneBook) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"Get": {
				Description: `Get returns all people from DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "search",
						Type: smd.Object,
						Properties: smd.PropertyList{
							{
								Name:        "ByName",
								Optional:    true,
								Description: `ByName is filter for searching person by first name or last name.`,
								Type:        smd.String,
							},
							{
								Name:     "ByType",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "ByPhone",
								Type: smd.String,
							},
							{
								Name:     "ByAddress",
								Optional: true,
								Ref:      "#/definitions/Address",
								Type:     smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"Address": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "Street",
										Type: smd.String,
									},
									{
										Name: "City",
										Type: smd.String,
									},
								},
							},
						},
					},
					{
						Name:        "page",
						Optional:    true,
						Description: `current page`,
						Type:        smd.Integer,
					},
					{
						Name:        "count",
						Optional:    true,
						Description: `page size`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Array,
					Items: map[string]string{
						"$ref": "#/definitions/Person",
					},
					Definitions: map[string]smd.Definition{
						"Person": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name:        "ID",
									Description: `ID is Unique Identifier for person`,
									Type:        smd.Integer,
								},
								{
									Name: "FirstName",
									Type: smd.String,
								},
								{
									Name: "LastName",
									Type: smd.String,
								},
								{
									Name:        "Phone",
									Description: `Phone is main phone`,
									Type:        smd.String,
								},
								{
									Name:     "WorkPhone",
									Optional: true,
									Type:     smd.String,
								},
								{
									Name: "Mobile",
									Type: smd.Array,
									Items: map[string]string{
										"type": smd.String,
									},
								},
								{
									Name:        "Deleted",
									Description: `Deleted is flag for`,
									Type:        smd.Boolean,
								},
								{
									Name:        "Addresses",
									Description: `Addresses Could be nil or len() == 0.`,
									Type:        smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/Address",
									},
								},
								{
									Name:     "address",
									Optional: true,
									Ref:      "#/definitions/Address",
									Type:     smd.Object,
								},
							},
						},
						"Address": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "Street",
									Type: smd.String,
								},
								{
									Name: "City",
									Type: smd.String,
								},
							},
						},
					},
				},
			},
			"ValidateSearch": {
				Description: `ValidateSearch returns given search as result.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "search",
						Optional:    true,
						Description: `search object`,
						Type:        smd.Object,
						Properties: smd.PropertyList{
							{
								Name:        "ByName",
								Optional:    true,
								Description: `ByName is filter for searching person by first name or last name.`,
								Type:        smd.String,
							},
							{
								Name:     "ByType",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "ByPhone",
								Type: smd.String,
							},
							{
								Name:     "ByAddress",
								Optional: true,
								Ref:      "#/definitions/Address",
								Type:     smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"Address": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "Street",
										Type: smd.String,
									},
									{
										Name: "City",
										Type: smd.String,
									},
								},
							},
						},
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
					Properties: smd.PropertyList{
						{
							Name:        "ByName",
							Optional:    true,
							Description: `ByName is filter for searching person by first name or last name.`,
							Type:        smd.String,
						},
						{
							Name:     "ByType",
							Optional: true,
							Type:     smd.String,
						},
						{
							Name: "ByPhone",
							Type: smd.String,
						},
						{
							Name:     "ByAddress",
							Optional: true,
							Ref:      "#/definitions/Address",
							Type:     smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"Address": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "Street",
									Type: smd.String,
								},
								{
									Name: "City",
									Type: smd.String,
								},
							},
						},
					},
				},
			},
			"ById": {
				Description: `ById returns Person from DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Description: `person id`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Optional: true,
					Type:     smd.Object,
					Properties: smd.PropertyList{
						{
							Name:        "ID",
							Description: `ID is Unique Identifier for person`,
							Type:        smd.Integer,
						},
						{
							Name: "FirstName",
							Type: smd.String,
						},
						{
							Name: "LastName",
							Type: smd.String,
						},
						{
							Name:        "Phone",
							Description: `Phone is main phone`,
							Type:        smd.String,
						},
						{
							Name:     "WorkPhone",
							Optional: true,
							Type:     smd.String,
						},
						{
							Name: "Mobile",
							Type: smd.Array,
							Items: map[string]string{
								"type": smd.String,
							},
						},
						{
							Name:        "Deleted",
							Description: `Deleted is flag for`,
							Type:        smd.Boolean,
						},
						{
							Name:        "Addresses",
							Description: `Addresses Could be nil or len() == 0.`,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/Address",
							},
						},
						{
							Name:     "address",
							Optional: true,
							Ref:      "#/definitions/Address",
							Type:     smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"Address": {
							Type: "object",
							Properties: smd.PropertyList{
								{
									Name: "Street",
									Type: smd.String,
								},
								{
									Name: "City",
									Type: smd.String,
								},
							},
						},
					},
				},
				Errors: map[int]string{
					404: "person was not found",
				},
			},
			"Delete": {
				Description: `Delete marks person as deleted.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Description: `person id`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Boolean,
				},
			},
			"Remove": {
				Description: `Removes deletes person from DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "id",
						Description: `person id`,
						Type:        smd.Integer,
					},
				},
				Returns: smd.JSONSchema{
					Description: `operation result`,
					Type:        smd.Boolean,
				},
			},
			"Save": {
				Description: `Save saves person to DB.`,
				Parameters: []smd.JSONSchema{
					{
						Name: "p",
						Type: smd.Object,
						Properties: smd.PropertyList{
							{
								Name:        "ID",
								Description: `ID is Unique Identifier for person`,
								Type:        smd.Integer,
							},
							{
								Name: "FirstName",
								Type: smd.String,
							},
							{
								Name: "LastName",
								Type: smd.String,
							},
							{
								Name:        "Phone",
								Description: `Phone is main phone`,
								Type:        smd.String,
							},
							{
								Name:     "WorkPhone",
								Optional: true,
								Type:     smd.String,
							},
							{
								Name: "Mobile",
								Type: smd.Array,
								Items: map[string]string{
									"type": smd.String,
								},
							},
							{
								Name:        "Deleted",
								Description: `Deleted is flag for`,
								Type:        smd.Boolean,
							},
							{
								Name:        "Addresses",
								Description: `Addresses Could be nil or len() == 0.`,
								Type:        smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/Address",
								},
							},
							{
								Name:     "address",
								Optional: true,
								Ref:      "#/definitions/Address",
								Type:     smd.Object,
							},
						},
						Definitions: map[string]smd.Definition{
							"Address": {
								Type: "object",
								Properties: smd.PropertyList{
									{
										Name: "Street",
										Type: smd.String,
									},
									{
										Name: "City",
										Type: smd.String,
									},
								},
							},
						},
					},
					{
						Name:        "replace",
						Optional:    true,
						Description: `update person if exist`,
						Type:        smd.Boolean,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.Integer,
				},
				Errors: map[int]string{
					400: "invalid request",
					401: "use replace=true",
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (pb PhoneBook) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.PhoneBook.Get:
		var args = struct {
			Search PersonSearch `json:"search"`
			Page   *int         `json:"page"`
			Count  *int         `json:"count"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"search", "page", "count"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:count=50 page size
		if args.Count == nil {
			var v int = 50
			args.Count = &v
		}

		//zenrpc:page=0 current page
		if args.Page == nil {
			var v int = 0
			args.Page = &v
		}

		resp.Set(pb.Get(args.Search, args.Page, args.Count))

	case RPC.PhoneBook.ValidateSearch:
		var args = struct {
			Search *PersonSearch `json:"search"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"search"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(pb.ValidateSearch(args.Search))

	case RPC.PhoneBook.ById:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(pb.ById(args.Id))

	case RPC.PhoneBook.Delete:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(pb.Delete(args.Id))

	case RPC.PhoneBook.Remove:
		var args = struct {
			Id uint64 `json:"id"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"id"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(pb.Remove(args.Id))

	case RPC.PhoneBook.Save:
		var args = struct {
			P       Person `json:"p"`
			Replace *bool  `json:"replace"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"p", "replace"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:replace=false update person if exist
		if args.Replace == nil {
			var v bool = false
			args.Replace = &v
		}

		resp.Set(pb.Save(args.P, args.Replace))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (PrintService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Methods: map[string]smd.Service{
			"PrintRequiredDefault": {
				Parameters: []smd.JSONSchema{
					{
						Name:     "s",
						Optional: true,
						Type:     smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.String,
				},
			},
			"PrintOptionalWithDefault": {
				Parameters: []smd.JSONSchema{
					{
						Name:     "s",
						Optional: true,
						Type:     smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.String,
				},
			},
			"PrintRequired": {
				Parameters: []smd.JSONSchema{
					{
						Name: "s",
						Type: smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.String,
				},
			},
			"PrintOptional": {
				Parameters: []smd.JSONSchema{
					{
						Name:     "s",
						Optional: true,
						Type:     smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Type: smd.String,
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s PrintService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.PrintService.PrintRequiredDefault:
		var args = struct {
			S *string `json:"s"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"s"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:s="test"
		if args.S == nil {
			var v string = "test"
			args.S = &v
		}

		resp.Set(s.PrintRequiredDefault(*args.S))

	case RPC.PrintService.PrintOptionalWithDefault:
		var args = struct {
			S *string `json:"s"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"s"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		//zenrpc:s="test"
		if args.S == nil {
			var v string = "test"
			args.S = &v
		}

		resp.Set(s.PrintOptionalWithDefault(args.S))

	case RPC.PrintService.PrintRequired:
		var args = struct {
			S string `json:"s"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"s"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.PrintRequired(args.S))

	case RPC.PrintService.PrintOptional:
		var args = struct {
			S *string `json:"s"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"s"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.PrintOptional(args.S))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
