// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This resource can manage the queuing QoS policy map configuration.
//
// - API Documentation: [ipqosPMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:PMapInst/)
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
//			_, err := nxos.NewQueuingQosPolicyMap(ctx, "example", &nxos.QueuingQosPolicyMapArgs{
//				MatchType: pulumi.String("match-any"),
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
//	$ pulumi import nxos:index/queuingQosPolicyMap:QueuingQosPolicyMap example "sys/ipqos/queuing/p/name-[PM1]"
//
// ```
type QueuingQosPolicyMap struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType pulumi.StringOutput `pulumi:"matchType"`
	// Policy map name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewQueuingQosPolicyMap registers a new resource with the given unique name, arguments, and options.
func NewQueuingQosPolicyMap(ctx *pulumi.Context,
	name string, args *QueuingQosPolicyMapArgs, opts ...pulumi.ResourceOption) (*QueuingQosPolicyMap, error) {
	if args == nil {
		args = &QueuingQosPolicyMapArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueuingQosPolicyMap
	err := ctx.RegisterResource("nxos:index/queuingQosPolicyMap:QueuingQosPolicyMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueuingQosPolicyMap gets an existing QueuingQosPolicyMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueuingQosPolicyMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueuingQosPolicyMapState, opts ...pulumi.ResourceOption) (*QueuingQosPolicyMap, error) {
	var resource QueuingQosPolicyMap
	err := ctx.ReadResource("nxos:index/queuingQosPolicyMap:QueuingQosPolicyMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueuingQosPolicyMap resources.
type queuingQosPolicyMapState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType *string `pulumi:"matchType"`
	// Policy map name.
	Name *string `pulumi:"name"`
}

type QueuingQosPolicyMapState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType pulumi.StringPtrInput
	// Policy map name.
	Name pulumi.StringPtrInput
}

func (QueuingQosPolicyMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicyMapState)(nil)).Elem()
}

type queuingQosPolicyMapArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType *string `pulumi:"matchType"`
	// Policy map name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a QueuingQosPolicyMap resource.
type QueuingQosPolicyMapArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
	MatchType pulumi.StringPtrInput
	// Policy map name.
	Name pulumi.StringPtrInput
}

func (QueuingQosPolicyMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicyMapArgs)(nil)).Elem()
}

type QueuingQosPolicyMapInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapOutput() QueuingQosPolicyMapOutput
	ToQueuingQosPolicyMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapOutput
}

func (*QueuingQosPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicyMap)(nil)).Elem()
}

func (i *QueuingQosPolicyMap) ToQueuingQosPolicyMapOutput() QueuingQosPolicyMapOutput {
	return i.ToQueuingQosPolicyMapOutputWithContext(context.Background())
}

func (i *QueuingQosPolicyMap) ToQueuingQosPolicyMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapOutput)
}

func (i *QueuingQosPolicyMap) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicyMap] {
	return pulumix.Output[*QueuingQosPolicyMap]{
		OutputState: i.ToQueuingQosPolicyMapOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicyMapArrayInput is an input type that accepts QueuingQosPolicyMapArray and QueuingQosPolicyMapArrayOutput values.
// You can construct a concrete instance of `QueuingQosPolicyMapArrayInput` via:
//
//	QueuingQosPolicyMapArray{ QueuingQosPolicyMapArgs{...} }
type QueuingQosPolicyMapArrayInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapArrayOutput() QueuingQosPolicyMapArrayOutput
	ToQueuingQosPolicyMapArrayOutputWithContext(context.Context) QueuingQosPolicyMapArrayOutput
}

type QueuingQosPolicyMapArray []QueuingQosPolicyMapInput

func (QueuingQosPolicyMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicyMap)(nil)).Elem()
}

func (i QueuingQosPolicyMapArray) ToQueuingQosPolicyMapArrayOutput() QueuingQosPolicyMapArrayOutput {
	return i.ToQueuingQosPolicyMapArrayOutputWithContext(context.Background())
}

func (i QueuingQosPolicyMapArray) ToQueuingQosPolicyMapArrayOutputWithContext(ctx context.Context) QueuingQosPolicyMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapArrayOutput)
}

func (i QueuingQosPolicyMapArray) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicyMap] {
	return pulumix.Output[[]*QueuingQosPolicyMap]{
		OutputState: i.ToQueuingQosPolicyMapArrayOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicyMapMapInput is an input type that accepts QueuingQosPolicyMapMap and QueuingQosPolicyMapMapOutput values.
// You can construct a concrete instance of `QueuingQosPolicyMapMapInput` via:
//
//	QueuingQosPolicyMapMap{ "key": QueuingQosPolicyMapArgs{...} }
type QueuingQosPolicyMapMapInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapMapOutput() QueuingQosPolicyMapMapOutput
	ToQueuingQosPolicyMapMapOutputWithContext(context.Context) QueuingQosPolicyMapMapOutput
}

type QueuingQosPolicyMapMap map[string]QueuingQosPolicyMapInput

func (QueuingQosPolicyMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicyMap)(nil)).Elem()
}

func (i QueuingQosPolicyMapMap) ToQueuingQosPolicyMapMapOutput() QueuingQosPolicyMapMapOutput {
	return i.ToQueuingQosPolicyMapMapOutputWithContext(context.Background())
}

func (i QueuingQosPolicyMapMap) ToQueuingQosPolicyMapMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapMapOutput)
}

func (i QueuingQosPolicyMapMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicyMap] {
	return pulumix.Output[map[string]*QueuingQosPolicyMap]{
		OutputState: i.ToQueuingQosPolicyMapMapOutputWithContext(ctx).OutputState,
	}
}

type QueuingQosPolicyMapOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicyMap)(nil)).Elem()
}

func (o QueuingQosPolicyMapOutput) ToQueuingQosPolicyMapOutput() QueuingQosPolicyMapOutput {
	return o
}

func (o QueuingQosPolicyMapOutput) ToQueuingQosPolicyMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapOutput {
	return o
}

func (o QueuingQosPolicyMapOutput) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicyMap] {
	return pulumix.Output[*QueuingQosPolicyMap]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o QueuingQosPolicyMapOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMap) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
func (o QueuingQosPolicyMapOutput) MatchType() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMap) pulumi.StringOutput { return v.MatchType }).(pulumi.StringOutput)
}

// Policy map name.
func (o QueuingQosPolicyMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type QueuingQosPolicyMapArrayOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicyMap)(nil)).Elem()
}

func (o QueuingQosPolicyMapArrayOutput) ToQueuingQosPolicyMapArrayOutput() QueuingQosPolicyMapArrayOutput {
	return o
}

func (o QueuingQosPolicyMapArrayOutput) ToQueuingQosPolicyMapArrayOutputWithContext(ctx context.Context) QueuingQosPolicyMapArrayOutput {
	return o
}

func (o QueuingQosPolicyMapArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicyMap] {
	return pulumix.Output[[]*QueuingQosPolicyMap]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicyMapArrayOutput) Index(i pulumi.IntInput) QueuingQosPolicyMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QueuingQosPolicyMap {
		return vs[0].([]*QueuingQosPolicyMap)[vs[1].(int)]
	}).(QueuingQosPolicyMapOutput)
}

type QueuingQosPolicyMapMapOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicyMap)(nil)).Elem()
}

func (o QueuingQosPolicyMapMapOutput) ToQueuingQosPolicyMapMapOutput() QueuingQosPolicyMapMapOutput {
	return o
}

func (o QueuingQosPolicyMapMapOutput) ToQueuingQosPolicyMapMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMapOutput {
	return o
}

func (o QueuingQosPolicyMapMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicyMap] {
	return pulumix.Output[map[string]*QueuingQosPolicyMap]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicyMapMapOutput) MapIndex(k pulumi.StringInput) QueuingQosPolicyMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QueuingQosPolicyMap {
		return vs[0].(map[string]*QueuingQosPolicyMap)[vs[1].(string)]
	}).(QueuingQosPolicyMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapInput)(nil)).Elem(), &QueuingQosPolicyMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapArrayInput)(nil)).Elem(), QueuingQosPolicyMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapMapInput)(nil)).Elem(), QueuingQosPolicyMapMap{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapArrayOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapMapOutput{})
}