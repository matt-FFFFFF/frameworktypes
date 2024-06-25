package frameworktype

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/matt-FFFFFF/frameworktypes/to"
)

func TestToGoMap(t *testing.T) {
	ctx := context.Background()

	// Test map[string]string input
	inputMap, _ := basetypes.NewMapValue(types.StringType, map[string]attr.Value{
		"key1": types.StringValue("value1"),
		"key2": types.StringValue("value2"),
	})
	stringMapResult := ToGoMap[string](ctx, inputMap)
	expectedStringMapResult := map[string]*string{
		"key1": to.Ptr("value1"),
		"key2": to.Ptr("value2"),
	}
	if !reflect.DeepEqual(stringMapResult, expectedStringMapResult) {
		t.Errorf("Expected string map result to be %v, but got %v", expectedStringMapResult, stringMapResult)
	}

	// // Test map[string]int input
	// inputMap = basetypes.NewMapValue(map[string]interface{}{
	// 	"key1": 1,
	// 	"key2": 2,
	// })
	// intMapResult := ToGoMap[int](ctx, inputMap)
	// expectedIntMapResult := map[string]*int{
	// 	"key1": basetypes.Int(1),
	// 	"key2": basetypes.Int(2),
	// }
	// if !reflect.DeepEqual(intMapResult, expectedIntMapResult) {
	// 	t.Errorf("Expected int map result to be %v, but got %v", expectedIntMapResult, intMapResult)
	// }

	// // Test invalid input
	// invalidInput := basetypes.String("not a map")
	// invalidMapResult := ToGoMap[string](ctx, invalidInput)
	// if invalidMapResult != nil {
	// 	t.Errorf("Expected invalid map result to be nil, but got %v", invalidMapResult)
	// }
}
