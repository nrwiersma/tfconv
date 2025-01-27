package tfconv_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/api/resource"
)

type likeAString string

type TestObject struct {
	Str         string                 `json:"str"`
	Alias       StrAlias               `json:"alias"`
	Int         int                    `json:"int"`
	Float       float64                `json:"float,omitempty"`
	Bool        bool                   `json:"bool"`
	Slice       []T                    `json:"slice"`
	StringSlice []string               `json:"stringSlice"`
	Map         map[string]int         `json:"map"`
	MapConvert  map[likeAString]string `json:"mapConvert"`
	Struct      *T                     `json:"struct"`
	Q           resource.Quantity      `json:"q"`
	QPtr        *resource.Quantity     `json:"qPtr"`
}

type T struct {
	A string `json:"a"`
	B *int   `json:"b"`
	C *int   `json:"c"`
}

type StrAlias string

func testObjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"alias": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"bool": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"float": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"int": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"map": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"map_convert": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"q": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"q_ptr": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			}},
		},
		"slice": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"a": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"string_slice": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"str": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"struct": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"a": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"b": {
						Type:     schema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"value": {
								Type:     schema.TypeInt,
								Required: true,
							},
						}},
					},
					"c": {
						Type:     schema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"value": {
								Type:     schema.TypeInt,
								Required: true,
							},
						}},
					},
				},
			},
		},
	}
}

func ptrOf[T any](v T) *T {
	return &v
}
