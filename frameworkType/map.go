package frameworktype

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ToGoMap[T any](ctx context.Context, input attr.Value) map[string]*T {
	inputMap, ok := input.(basetypes.MapValue)
	if !ok {
		return nil
	}
	switch inputMap.ElementType(ctx) {
	case types.StringType:
		res := make(map[string]*string)
		diags := inputMap.ElementsAs(ctx, &res, false)
		if diags.ErrorsCount() > 0 {
			return nil
		}
		return any(res).(map[string]*T)
	}
	return nil
}
