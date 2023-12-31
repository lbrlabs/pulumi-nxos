// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can manage a Match Route in  Route-Map Rule Entry configuration.
//
// - API Documentation: [rtmapMatchRtDst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:MatchRtDst/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := nxos.NewRouteMapRuleEntryMatchRoute(ctx, "example", &nxos.RouteMapRuleEntryMatchRouteArgs{
//				Order:    pulumi.Int(10),
//				RuleName: pulumi.String("RULE1"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute example "sys/rpm/rtmap-[RULE1]/ent-[10]/mrtdst"
//
// ```
type RouteMapRuleEntryMatchRoute struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Route-Map Rule Entry order. - Range: `0`-`65535`
	Order pulumi.IntOutput `pulumi:"order"`
	// Route Map rule name.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
}

// NewRouteMapRuleEntryMatchRoute registers a new resource with the given unique name, arguments, and options.
func NewRouteMapRuleEntryMatchRoute(ctx *pulumi.Context,
	name string, args *RouteMapRuleEntryMatchRouteArgs, opts ...pulumi.ResourceOption) (*RouteMapRuleEntryMatchRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteMapRuleEntryMatchRoute
	err := ctx.RegisterResource("nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteMapRuleEntryMatchRoute gets an existing RouteMapRuleEntryMatchRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteMapRuleEntryMatchRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteMapRuleEntryMatchRouteState, opts ...pulumi.ResourceOption) (*RouteMapRuleEntryMatchRoute, error) {
	var resource RouteMapRuleEntryMatchRoute
	err := ctx.ReadResource("nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteMapRuleEntryMatchRoute resources.
type routeMapRuleEntryMatchRouteState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route-Map Rule Entry order. - Range: `0`-`65535`
	Order *int `pulumi:"order"`
	// Route Map rule name.
	RuleName *string `pulumi:"ruleName"`
}

type RouteMapRuleEntryMatchRouteState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Route-Map Rule Entry order. - Range: `0`-`65535`
	Order pulumi.IntPtrInput
	// Route Map rule name.
	RuleName pulumi.StringPtrInput
}

func (RouteMapRuleEntryMatchRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeMapRuleEntryMatchRouteState)(nil)).Elem()
}

type routeMapRuleEntryMatchRouteArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route-Map Rule Entry order. - Range: `0`-`65535`
	Order int `pulumi:"order"`
	// Route Map rule name.
	RuleName string `pulumi:"ruleName"`
}

// The set of arguments for constructing a RouteMapRuleEntryMatchRoute resource.
type RouteMapRuleEntryMatchRouteArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Route-Map Rule Entry order. - Range: `0`-`65535`
	Order pulumi.IntInput
	// Route Map rule name.
	RuleName pulumi.StringInput
}

func (RouteMapRuleEntryMatchRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeMapRuleEntryMatchRouteArgs)(nil)).Elem()
}

type RouteMapRuleEntryMatchRouteInput interface {
	pulumi.Input

	ToRouteMapRuleEntryMatchRouteOutput() RouteMapRuleEntryMatchRouteOutput
	ToRouteMapRuleEntryMatchRouteOutputWithContext(ctx context.Context) RouteMapRuleEntryMatchRouteOutput
}

func (*RouteMapRuleEntryMatchRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteMapRuleEntryMatchRoute)(nil)).Elem()
}

func (i *RouteMapRuleEntryMatchRoute) ToRouteMapRuleEntryMatchRouteOutput() RouteMapRuleEntryMatchRouteOutput {
	return i.ToRouteMapRuleEntryMatchRouteOutputWithContext(context.Background())
}

func (i *RouteMapRuleEntryMatchRoute) ToRouteMapRuleEntryMatchRouteOutputWithContext(ctx context.Context) RouteMapRuleEntryMatchRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapRuleEntryMatchRouteOutput)
}

func (i *RouteMapRuleEntryMatchRoute) ToOutput(ctx context.Context) pulumix.Output[*RouteMapRuleEntryMatchRoute] {
	return pulumix.Output[*RouteMapRuleEntryMatchRoute]{
		OutputState: i.ToRouteMapRuleEntryMatchRouteOutputWithContext(ctx).OutputState,
	}
}

// RouteMapRuleEntryMatchRouteArrayInput is an input type that accepts RouteMapRuleEntryMatchRouteArray and RouteMapRuleEntryMatchRouteArrayOutput values.
// You can construct a concrete instance of `RouteMapRuleEntryMatchRouteArrayInput` via:
//
//	RouteMapRuleEntryMatchRouteArray{ RouteMapRuleEntryMatchRouteArgs{...} }
type RouteMapRuleEntryMatchRouteArrayInput interface {
	pulumi.Input

	ToRouteMapRuleEntryMatchRouteArrayOutput() RouteMapRuleEntryMatchRouteArrayOutput
	ToRouteMapRuleEntryMatchRouteArrayOutputWithContext(context.Context) RouteMapRuleEntryMatchRouteArrayOutput
}

type RouteMapRuleEntryMatchRouteArray []RouteMapRuleEntryMatchRouteInput

func (RouteMapRuleEntryMatchRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteMapRuleEntryMatchRoute)(nil)).Elem()
}

func (i RouteMapRuleEntryMatchRouteArray) ToRouteMapRuleEntryMatchRouteArrayOutput() RouteMapRuleEntryMatchRouteArrayOutput {
	return i.ToRouteMapRuleEntryMatchRouteArrayOutputWithContext(context.Background())
}

func (i RouteMapRuleEntryMatchRouteArray) ToRouteMapRuleEntryMatchRouteArrayOutputWithContext(ctx context.Context) RouteMapRuleEntryMatchRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapRuleEntryMatchRouteArrayOutput)
}

func (i RouteMapRuleEntryMatchRouteArray) ToOutput(ctx context.Context) pulumix.Output[[]*RouteMapRuleEntryMatchRoute] {
	return pulumix.Output[[]*RouteMapRuleEntryMatchRoute]{
		OutputState: i.ToRouteMapRuleEntryMatchRouteArrayOutputWithContext(ctx).OutputState,
	}
}

// RouteMapRuleEntryMatchRouteMapInput is an input type that accepts RouteMapRuleEntryMatchRouteMap and RouteMapRuleEntryMatchRouteMapOutput values.
// You can construct a concrete instance of `RouteMapRuleEntryMatchRouteMapInput` via:
//
//	RouteMapRuleEntryMatchRouteMap{ "key": RouteMapRuleEntryMatchRouteArgs{...} }
type RouteMapRuleEntryMatchRouteMapInput interface {
	pulumi.Input

	ToRouteMapRuleEntryMatchRouteMapOutput() RouteMapRuleEntryMatchRouteMapOutput
	ToRouteMapRuleEntryMatchRouteMapOutputWithContext(context.Context) RouteMapRuleEntryMatchRouteMapOutput
}

type RouteMapRuleEntryMatchRouteMap map[string]RouteMapRuleEntryMatchRouteInput

func (RouteMapRuleEntryMatchRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteMapRuleEntryMatchRoute)(nil)).Elem()
}

func (i RouteMapRuleEntryMatchRouteMap) ToRouteMapRuleEntryMatchRouteMapOutput() RouteMapRuleEntryMatchRouteMapOutput {
	return i.ToRouteMapRuleEntryMatchRouteMapOutputWithContext(context.Background())
}

func (i RouteMapRuleEntryMatchRouteMap) ToRouteMapRuleEntryMatchRouteMapOutputWithContext(ctx context.Context) RouteMapRuleEntryMatchRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapRuleEntryMatchRouteMapOutput)
}

func (i RouteMapRuleEntryMatchRouteMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouteMapRuleEntryMatchRoute] {
	return pulumix.Output[map[string]*RouteMapRuleEntryMatchRoute]{
		OutputState: i.ToRouteMapRuleEntryMatchRouteMapOutputWithContext(ctx).OutputState,
	}
}

type RouteMapRuleEntryMatchRouteOutput struct{ *pulumi.OutputState }

func (RouteMapRuleEntryMatchRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteMapRuleEntryMatchRoute)(nil)).Elem()
}

func (o RouteMapRuleEntryMatchRouteOutput) ToRouteMapRuleEntryMatchRouteOutput() RouteMapRuleEntryMatchRouteOutput {
	return o
}

func (o RouteMapRuleEntryMatchRouteOutput) ToRouteMapRuleEntryMatchRouteOutputWithContext(ctx context.Context) RouteMapRuleEntryMatchRouteOutput {
	return o
}

func (o RouteMapRuleEntryMatchRouteOutput) ToOutput(ctx context.Context) pulumix.Output[*RouteMapRuleEntryMatchRoute] {
	return pulumix.Output[*RouteMapRuleEntryMatchRoute]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o RouteMapRuleEntryMatchRouteOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteMapRuleEntryMatchRoute) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Route-Map Rule Entry order. - Range: `0`-`65535`
func (o RouteMapRuleEntryMatchRouteOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *RouteMapRuleEntryMatchRoute) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// Route Map rule name.
func (o RouteMapRuleEntryMatchRouteOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMapRuleEntryMatchRoute) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

type RouteMapRuleEntryMatchRouteArrayOutput struct{ *pulumi.OutputState }

func (RouteMapRuleEntryMatchRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteMapRuleEntryMatchRoute)(nil)).Elem()
}

func (o RouteMapRuleEntryMatchRouteArrayOutput) ToRouteMapRuleEntryMatchRouteArrayOutput() RouteMapRuleEntryMatchRouteArrayOutput {
	return o
}

func (o RouteMapRuleEntryMatchRouteArrayOutput) ToRouteMapRuleEntryMatchRouteArrayOutputWithContext(ctx context.Context) RouteMapRuleEntryMatchRouteArrayOutput {
	return o
}

func (o RouteMapRuleEntryMatchRouteArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RouteMapRuleEntryMatchRoute] {
	return pulumix.Output[[]*RouteMapRuleEntryMatchRoute]{
		OutputState: o.OutputState,
	}
}

func (o RouteMapRuleEntryMatchRouteArrayOutput) Index(i pulumi.IntInput) RouteMapRuleEntryMatchRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteMapRuleEntryMatchRoute {
		return vs[0].([]*RouteMapRuleEntryMatchRoute)[vs[1].(int)]
	}).(RouteMapRuleEntryMatchRouteOutput)
}

type RouteMapRuleEntryMatchRouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapRuleEntryMatchRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteMapRuleEntryMatchRoute)(nil)).Elem()
}

func (o RouteMapRuleEntryMatchRouteMapOutput) ToRouteMapRuleEntryMatchRouteMapOutput() RouteMapRuleEntryMatchRouteMapOutput {
	return o
}

func (o RouteMapRuleEntryMatchRouteMapOutput) ToRouteMapRuleEntryMatchRouteMapOutputWithContext(ctx context.Context) RouteMapRuleEntryMatchRouteMapOutput {
	return o
}

func (o RouteMapRuleEntryMatchRouteMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouteMapRuleEntryMatchRoute] {
	return pulumix.Output[map[string]*RouteMapRuleEntryMatchRoute]{
		OutputState: o.OutputState,
	}
}

func (o RouteMapRuleEntryMatchRouteMapOutput) MapIndex(k pulumi.StringInput) RouteMapRuleEntryMatchRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteMapRuleEntryMatchRoute {
		return vs[0].(map[string]*RouteMapRuleEntryMatchRoute)[vs[1].(string)]
	}).(RouteMapRuleEntryMatchRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapRuleEntryMatchRouteInput)(nil)).Elem(), &RouteMapRuleEntryMatchRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapRuleEntryMatchRouteArrayInput)(nil)).Elem(), RouteMapRuleEntryMatchRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapRuleEntryMatchRouteMapInput)(nil)).Elem(), RouteMapRuleEntryMatchRouteMap{})
	pulumi.RegisterOutputType(RouteMapRuleEntryMatchRouteOutput{})
	pulumi.RegisterOutputType(RouteMapRuleEntryMatchRouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapRuleEntryMatchRouteMapOutput{})
}
