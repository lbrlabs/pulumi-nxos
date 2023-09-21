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

type DefaultQosPolicyInterfaceIn struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
}

// NewDefaultQosPolicyInterfaceIn registers a new resource with the given unique name, arguments, and options.
func NewDefaultQosPolicyInterfaceIn(ctx *pulumi.Context,
	name string, args *DefaultQosPolicyInterfaceInArgs, opts ...pulumi.ResourceOption) (*DefaultQosPolicyInterfaceIn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DefaultQosPolicyInterfaceIn
	err := ctx.RegisterResource("nxos:nxos/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultQosPolicyInterfaceIn gets an existing DefaultQosPolicyInterfaceIn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultQosPolicyInterfaceIn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultQosPolicyInterfaceInState, opts ...pulumi.ResourceOption) (*DefaultQosPolicyInterfaceIn, error) {
	var resource DefaultQosPolicyInterfaceIn
	err := ctx.ReadResource("nxos:nxos/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultQosPolicyInterfaceIn resources.
type defaultQosPolicyInterfaceInState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId *string `pulumi:"interfaceId"`
}

type DefaultQosPolicyInterfaceInState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringPtrInput
}

func (DefaultQosPolicyInterfaceInState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultQosPolicyInterfaceInState)(nil)).Elem()
}

type defaultQosPolicyInterfaceInArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
}

// The set of arguments for constructing a DefaultQosPolicyInterfaceIn resource.
type DefaultQosPolicyInterfaceInArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringInput
}

func (DefaultQosPolicyInterfaceInArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultQosPolicyInterfaceInArgs)(nil)).Elem()
}

type DefaultQosPolicyInterfaceInInput interface {
	pulumi.Input

	ToDefaultQosPolicyInterfaceInOutput() DefaultQosPolicyInterfaceInOutput
	ToDefaultQosPolicyInterfaceInOutputWithContext(ctx context.Context) DefaultQosPolicyInterfaceInOutput
}

func (*DefaultQosPolicyInterfaceIn) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultQosPolicyInterfaceIn)(nil)).Elem()
}

func (i *DefaultQosPolicyInterfaceIn) ToDefaultQosPolicyInterfaceInOutput() DefaultQosPolicyInterfaceInOutput {
	return i.ToDefaultQosPolicyInterfaceInOutputWithContext(context.Background())
}

func (i *DefaultQosPolicyInterfaceIn) ToDefaultQosPolicyInterfaceInOutputWithContext(ctx context.Context) DefaultQosPolicyInterfaceInOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultQosPolicyInterfaceInOutput)
}

func (i *DefaultQosPolicyInterfaceIn) ToOutput(ctx context.Context) pulumix.Output[*DefaultQosPolicyInterfaceIn] {
	return pulumix.Output[*DefaultQosPolicyInterfaceIn]{
		OutputState: i.ToDefaultQosPolicyInterfaceInOutputWithContext(ctx).OutputState,
	}
}

// DefaultQosPolicyInterfaceInArrayInput is an input type that accepts DefaultQosPolicyInterfaceInArray and DefaultQosPolicyInterfaceInArrayOutput values.
// You can construct a concrete instance of `DefaultQosPolicyInterfaceInArrayInput` via:
//
//	DefaultQosPolicyInterfaceInArray{ DefaultQosPolicyInterfaceInArgs{...} }
type DefaultQosPolicyInterfaceInArrayInput interface {
	pulumi.Input

	ToDefaultQosPolicyInterfaceInArrayOutput() DefaultQosPolicyInterfaceInArrayOutput
	ToDefaultQosPolicyInterfaceInArrayOutputWithContext(context.Context) DefaultQosPolicyInterfaceInArrayOutput
}

type DefaultQosPolicyInterfaceInArray []DefaultQosPolicyInterfaceInInput

func (DefaultQosPolicyInterfaceInArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultQosPolicyInterfaceIn)(nil)).Elem()
}

func (i DefaultQosPolicyInterfaceInArray) ToDefaultQosPolicyInterfaceInArrayOutput() DefaultQosPolicyInterfaceInArrayOutput {
	return i.ToDefaultQosPolicyInterfaceInArrayOutputWithContext(context.Background())
}

func (i DefaultQosPolicyInterfaceInArray) ToDefaultQosPolicyInterfaceInArrayOutputWithContext(ctx context.Context) DefaultQosPolicyInterfaceInArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultQosPolicyInterfaceInArrayOutput)
}

func (i DefaultQosPolicyInterfaceInArray) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultQosPolicyInterfaceIn] {
	return pulumix.Output[[]*DefaultQosPolicyInterfaceIn]{
		OutputState: i.ToDefaultQosPolicyInterfaceInArrayOutputWithContext(ctx).OutputState,
	}
}

// DefaultQosPolicyInterfaceInMapInput is an input type that accepts DefaultQosPolicyInterfaceInMap and DefaultQosPolicyInterfaceInMapOutput values.
// You can construct a concrete instance of `DefaultQosPolicyInterfaceInMapInput` via:
//
//	DefaultQosPolicyInterfaceInMap{ "key": DefaultQosPolicyInterfaceInArgs{...} }
type DefaultQosPolicyInterfaceInMapInput interface {
	pulumi.Input

	ToDefaultQosPolicyInterfaceInMapOutput() DefaultQosPolicyInterfaceInMapOutput
	ToDefaultQosPolicyInterfaceInMapOutputWithContext(context.Context) DefaultQosPolicyInterfaceInMapOutput
}

type DefaultQosPolicyInterfaceInMap map[string]DefaultQosPolicyInterfaceInInput

func (DefaultQosPolicyInterfaceInMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultQosPolicyInterfaceIn)(nil)).Elem()
}

func (i DefaultQosPolicyInterfaceInMap) ToDefaultQosPolicyInterfaceInMapOutput() DefaultQosPolicyInterfaceInMapOutput {
	return i.ToDefaultQosPolicyInterfaceInMapOutputWithContext(context.Background())
}

func (i DefaultQosPolicyInterfaceInMap) ToDefaultQosPolicyInterfaceInMapOutputWithContext(ctx context.Context) DefaultQosPolicyInterfaceInMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultQosPolicyInterfaceInMapOutput)
}

func (i DefaultQosPolicyInterfaceInMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultQosPolicyInterfaceIn] {
	return pulumix.Output[map[string]*DefaultQosPolicyInterfaceIn]{
		OutputState: i.ToDefaultQosPolicyInterfaceInMapOutputWithContext(ctx).OutputState,
	}
}

type DefaultQosPolicyInterfaceInOutput struct{ *pulumi.OutputState }

func (DefaultQosPolicyInterfaceInOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultQosPolicyInterfaceIn)(nil)).Elem()
}

func (o DefaultQosPolicyInterfaceInOutput) ToDefaultQosPolicyInterfaceInOutput() DefaultQosPolicyInterfaceInOutput {
	return o
}

func (o DefaultQosPolicyInterfaceInOutput) ToDefaultQosPolicyInterfaceInOutputWithContext(ctx context.Context) DefaultQosPolicyInterfaceInOutput {
	return o
}

func (o DefaultQosPolicyInterfaceInOutput) ToOutput(ctx context.Context) pulumix.Output[*DefaultQosPolicyInterfaceIn] {
	return pulumix.Output[*DefaultQosPolicyInterfaceIn]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o DefaultQosPolicyInterfaceInOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultQosPolicyInterfaceIn) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
func (o DefaultQosPolicyInterfaceInOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultQosPolicyInterfaceIn) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

type DefaultQosPolicyInterfaceInArrayOutput struct{ *pulumi.OutputState }

func (DefaultQosPolicyInterfaceInArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultQosPolicyInterfaceIn)(nil)).Elem()
}

func (o DefaultQosPolicyInterfaceInArrayOutput) ToDefaultQosPolicyInterfaceInArrayOutput() DefaultQosPolicyInterfaceInArrayOutput {
	return o
}

func (o DefaultQosPolicyInterfaceInArrayOutput) ToDefaultQosPolicyInterfaceInArrayOutputWithContext(ctx context.Context) DefaultQosPolicyInterfaceInArrayOutput {
	return o
}

func (o DefaultQosPolicyInterfaceInArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DefaultQosPolicyInterfaceIn] {
	return pulumix.Output[[]*DefaultQosPolicyInterfaceIn]{
		OutputState: o.OutputState,
	}
}

func (o DefaultQosPolicyInterfaceInArrayOutput) Index(i pulumi.IntInput) DefaultQosPolicyInterfaceInOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultQosPolicyInterfaceIn {
		return vs[0].([]*DefaultQosPolicyInterfaceIn)[vs[1].(int)]
	}).(DefaultQosPolicyInterfaceInOutput)
}

type DefaultQosPolicyInterfaceInMapOutput struct{ *pulumi.OutputState }

func (DefaultQosPolicyInterfaceInMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultQosPolicyInterfaceIn)(nil)).Elem()
}

func (o DefaultQosPolicyInterfaceInMapOutput) ToDefaultQosPolicyInterfaceInMapOutput() DefaultQosPolicyInterfaceInMapOutput {
	return o
}

func (o DefaultQosPolicyInterfaceInMapOutput) ToDefaultQosPolicyInterfaceInMapOutputWithContext(ctx context.Context) DefaultQosPolicyInterfaceInMapOutput {
	return o
}

func (o DefaultQosPolicyInterfaceInMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DefaultQosPolicyInterfaceIn] {
	return pulumix.Output[map[string]*DefaultQosPolicyInterfaceIn]{
		OutputState: o.OutputState,
	}
}

func (o DefaultQosPolicyInterfaceInMapOutput) MapIndex(k pulumi.StringInput) DefaultQosPolicyInterfaceInOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultQosPolicyInterfaceIn {
		return vs[0].(map[string]*DefaultQosPolicyInterfaceIn)[vs[1].(string)]
	}).(DefaultQosPolicyInterfaceInOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultQosPolicyInterfaceInInput)(nil)).Elem(), &DefaultQosPolicyInterfaceIn{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultQosPolicyInterfaceInArrayInput)(nil)).Elem(), DefaultQosPolicyInterfaceInArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultQosPolicyInterfaceInMapInput)(nil)).Elem(), DefaultQosPolicyInterfaceInMap{})
	pulumi.RegisterOutputType(DefaultQosPolicyInterfaceInOutput{})
	pulumi.RegisterOutputType(DefaultQosPolicyInterfaceInArrayOutput{})
	pulumi.RegisterOutputType(DefaultQosPolicyInterfaceInMapOutput{})
}
