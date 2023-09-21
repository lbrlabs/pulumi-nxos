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

type QueuingQosPolicyMapMatchClassMap struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Class map name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy map name.
	PolicyMapName pulumi.StringOutput `pulumi:"policyMapName"`
}

// NewQueuingQosPolicyMapMatchClassMap registers a new resource with the given unique name, arguments, and options.
func NewQueuingQosPolicyMapMatchClassMap(ctx *pulumi.Context,
	name string, args *QueuingQosPolicyMapMatchClassMapArgs, opts ...pulumi.ResourceOption) (*QueuingQosPolicyMapMatchClassMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyMapName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyMapName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QueuingQosPolicyMapMatchClassMap
	err := ctx.RegisterResource("nxos:nxos/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueuingQosPolicyMapMatchClassMap gets an existing QueuingQosPolicyMapMatchClassMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueuingQosPolicyMapMatchClassMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueuingQosPolicyMapMatchClassMapState, opts ...pulumi.ResourceOption) (*QueuingQosPolicyMapMatchClassMap, error) {
	var resource QueuingQosPolicyMapMatchClassMap
	err := ctx.ReadResource("nxos:nxos/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueuingQosPolicyMapMatchClassMap resources.
type queuingQosPolicyMapMatchClassMapState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Class map name.
	Name *string `pulumi:"name"`
	// Policy map name.
	PolicyMapName *string `pulumi:"policyMapName"`
}

type QueuingQosPolicyMapMatchClassMapState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Class map name.
	Name pulumi.StringPtrInput
	// Policy map name.
	PolicyMapName pulumi.StringPtrInput
}

func (QueuingQosPolicyMapMatchClassMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicyMapMatchClassMapState)(nil)).Elem()
}

type queuingQosPolicyMapMatchClassMapArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Class map name.
	Name *string `pulumi:"name"`
	// Policy map name.
	PolicyMapName string `pulumi:"policyMapName"`
}

// The set of arguments for constructing a QueuingQosPolicyMapMatchClassMap resource.
type QueuingQosPolicyMapMatchClassMapArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Class map name.
	Name pulumi.StringPtrInput
	// Policy map name.
	PolicyMapName pulumi.StringInput
}

func (QueuingQosPolicyMapMatchClassMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queuingQosPolicyMapMatchClassMapArgs)(nil)).Elem()
}

type QueuingQosPolicyMapMatchClassMapInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapMatchClassMapOutput() QueuingQosPolicyMapMatchClassMapOutput
	ToQueuingQosPolicyMapMatchClassMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapOutput
}

func (*QueuingQosPolicyMapMatchClassMap) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicyMapMatchClassMap)(nil)).Elem()
}

func (i *QueuingQosPolicyMapMatchClassMap) ToQueuingQosPolicyMapMatchClassMapOutput() QueuingQosPolicyMapMatchClassMapOutput {
	return i.ToQueuingQosPolicyMapMatchClassMapOutputWithContext(context.Background())
}

func (i *QueuingQosPolicyMapMatchClassMap) ToQueuingQosPolicyMapMatchClassMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapMatchClassMapOutput)
}

func (i *QueuingQosPolicyMapMatchClassMap) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicyMapMatchClassMap] {
	return pulumix.Output[*QueuingQosPolicyMapMatchClassMap]{
		OutputState: i.ToQueuingQosPolicyMapMatchClassMapOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicyMapMatchClassMapArrayInput is an input type that accepts QueuingQosPolicyMapMatchClassMapArray and QueuingQosPolicyMapMatchClassMapArrayOutput values.
// You can construct a concrete instance of `QueuingQosPolicyMapMatchClassMapArrayInput` via:
//
//	QueuingQosPolicyMapMatchClassMapArray{ QueuingQosPolicyMapMatchClassMapArgs{...} }
type QueuingQosPolicyMapMatchClassMapArrayInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapMatchClassMapArrayOutput() QueuingQosPolicyMapMatchClassMapArrayOutput
	ToQueuingQosPolicyMapMatchClassMapArrayOutputWithContext(context.Context) QueuingQosPolicyMapMatchClassMapArrayOutput
}

type QueuingQosPolicyMapMatchClassMapArray []QueuingQosPolicyMapMatchClassMapInput

func (QueuingQosPolicyMapMatchClassMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicyMapMatchClassMap)(nil)).Elem()
}

func (i QueuingQosPolicyMapMatchClassMapArray) ToQueuingQosPolicyMapMatchClassMapArrayOutput() QueuingQosPolicyMapMatchClassMapArrayOutput {
	return i.ToQueuingQosPolicyMapMatchClassMapArrayOutputWithContext(context.Background())
}

func (i QueuingQosPolicyMapMatchClassMapArray) ToQueuingQosPolicyMapMatchClassMapArrayOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapMatchClassMapArrayOutput)
}

func (i QueuingQosPolicyMapMatchClassMapArray) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicyMapMatchClassMap] {
	return pulumix.Output[[]*QueuingQosPolicyMapMatchClassMap]{
		OutputState: i.ToQueuingQosPolicyMapMatchClassMapArrayOutputWithContext(ctx).OutputState,
	}
}

// QueuingQosPolicyMapMatchClassMapMapInput is an input type that accepts QueuingQosPolicyMapMatchClassMapMap and QueuingQosPolicyMapMatchClassMapMapOutput values.
// You can construct a concrete instance of `QueuingQosPolicyMapMatchClassMapMapInput` via:
//
//	QueuingQosPolicyMapMatchClassMapMap{ "key": QueuingQosPolicyMapMatchClassMapArgs{...} }
type QueuingQosPolicyMapMatchClassMapMapInput interface {
	pulumi.Input

	ToQueuingQosPolicyMapMatchClassMapMapOutput() QueuingQosPolicyMapMatchClassMapMapOutput
	ToQueuingQosPolicyMapMatchClassMapMapOutputWithContext(context.Context) QueuingQosPolicyMapMatchClassMapMapOutput
}

type QueuingQosPolicyMapMatchClassMapMap map[string]QueuingQosPolicyMapMatchClassMapInput

func (QueuingQosPolicyMapMatchClassMapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicyMapMatchClassMap)(nil)).Elem()
}

func (i QueuingQosPolicyMapMatchClassMapMap) ToQueuingQosPolicyMapMatchClassMapMapOutput() QueuingQosPolicyMapMatchClassMapMapOutput {
	return i.ToQueuingQosPolicyMapMatchClassMapMapOutputWithContext(context.Background())
}

func (i QueuingQosPolicyMapMatchClassMapMap) ToQueuingQosPolicyMapMatchClassMapMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuingQosPolicyMapMatchClassMapMapOutput)
}

func (i QueuingQosPolicyMapMatchClassMapMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMap] {
	return pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMap]{
		OutputState: i.ToQueuingQosPolicyMapMatchClassMapMapOutputWithContext(ctx).OutputState,
	}
}

type QueuingQosPolicyMapMatchClassMapOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapMatchClassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuingQosPolicyMapMatchClassMap)(nil)).Elem()
}

func (o QueuingQosPolicyMapMatchClassMapOutput) ToQueuingQosPolicyMapMatchClassMapOutput() QueuingQosPolicyMapMatchClassMapOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapOutput) ToQueuingQosPolicyMapMatchClassMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapOutput) ToOutput(ctx context.Context) pulumix.Output[*QueuingQosPolicyMapMatchClassMap] {
	return pulumix.Output[*QueuingQosPolicyMapMatchClassMap]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o QueuingQosPolicyMapMatchClassMapOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMapMatchClassMap) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Class map name.
func (o QueuingQosPolicyMapMatchClassMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMapMatchClassMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy map name.
func (o QueuingQosPolicyMapMatchClassMapOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v *QueuingQosPolicyMapMatchClassMap) pulumi.StringOutput { return v.PolicyMapName }).(pulumi.StringOutput)
}

type QueuingQosPolicyMapMatchClassMapArrayOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapMatchClassMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QueuingQosPolicyMapMatchClassMap)(nil)).Elem()
}

func (o QueuingQosPolicyMapMatchClassMapArrayOutput) ToQueuingQosPolicyMapMatchClassMapArrayOutput() QueuingQosPolicyMapMatchClassMapArrayOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapArrayOutput) ToQueuingQosPolicyMapMatchClassMapArrayOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapArrayOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*QueuingQosPolicyMapMatchClassMap] {
	return pulumix.Output[[]*QueuingQosPolicyMapMatchClassMap]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicyMapMatchClassMapArrayOutput) Index(i pulumi.IntInput) QueuingQosPolicyMapMatchClassMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QueuingQosPolicyMapMatchClassMap {
		return vs[0].([]*QueuingQosPolicyMapMatchClassMap)[vs[1].(int)]
	}).(QueuingQosPolicyMapMatchClassMapOutput)
}

type QueuingQosPolicyMapMatchClassMapMapOutput struct{ *pulumi.OutputState }

func (QueuingQosPolicyMapMatchClassMapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QueuingQosPolicyMapMatchClassMap)(nil)).Elem()
}

func (o QueuingQosPolicyMapMatchClassMapMapOutput) ToQueuingQosPolicyMapMatchClassMapMapOutput() QueuingQosPolicyMapMatchClassMapMapOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapMapOutput) ToQueuingQosPolicyMapMatchClassMapMapOutputWithContext(ctx context.Context) QueuingQosPolicyMapMatchClassMapMapOutput {
	return o
}

func (o QueuingQosPolicyMapMatchClassMapMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMap] {
	return pulumix.Output[map[string]*QueuingQosPolicyMapMatchClassMap]{
		OutputState: o.OutputState,
	}
}

func (o QueuingQosPolicyMapMatchClassMapMapOutput) MapIndex(k pulumi.StringInput) QueuingQosPolicyMapMatchClassMapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QueuingQosPolicyMapMatchClassMap {
		return vs[0].(map[string]*QueuingQosPolicyMapMatchClassMap)[vs[1].(string)]
	}).(QueuingQosPolicyMapMatchClassMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapMatchClassMapInput)(nil)).Elem(), &QueuingQosPolicyMapMatchClassMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapMatchClassMapArrayInput)(nil)).Elem(), QueuingQosPolicyMapMatchClassMapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueuingQosPolicyMapMatchClassMapMapInput)(nil)).Elem(), QueuingQosPolicyMapMatchClassMapMap{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapMatchClassMapOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapMatchClassMapArrayOutput{})
	pulumi.RegisterOutputType(QueuingQosPolicyMapMatchClassMapMapOutput{})
}
