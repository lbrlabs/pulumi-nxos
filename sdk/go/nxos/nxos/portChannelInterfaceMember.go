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

type PortChannelInterfaceMember struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn pulumi.StringOutput `pulumi:"interfaceDn"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
}

// NewPortChannelInterfaceMember registers a new resource with the given unique name, arguments, and options.
func NewPortChannelInterfaceMember(ctx *pulumi.Context,
	name string, args *PortChannelInterfaceMemberArgs, opts ...pulumi.ResourceOption) (*PortChannelInterfaceMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceDn == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceDn'")
	}
	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PortChannelInterfaceMember
	err := ctx.RegisterResource("nxos:nxos/portChannelInterfaceMember:PortChannelInterfaceMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortChannelInterfaceMember gets an existing PortChannelInterfaceMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortChannelInterfaceMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortChannelInterfaceMemberState, opts ...pulumi.ResourceOption) (*PortChannelInterfaceMember, error) {
	var resource PortChannelInterfaceMember
	err := ctx.ReadResource("nxos:nxos/portChannelInterfaceMember:PortChannelInterfaceMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortChannelInterfaceMember resources.
type portChannelInterfaceMemberState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn *string `pulumi:"interfaceDn"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId *string `pulumi:"interfaceId"`
}

type PortChannelInterfaceMemberState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId pulumi.StringPtrInput
}

func (PortChannelInterfaceMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*portChannelInterfaceMemberState)(nil)).Elem()
}

type portChannelInterfaceMemberArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn string `pulumi:"interfaceDn"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId string `pulumi:"interfaceId"`
}

// The set of arguments for constructing a PortChannelInterfaceMember resource.
type PortChannelInterfaceMemberArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn pulumi.StringInput
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId pulumi.StringInput
}

func (PortChannelInterfaceMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portChannelInterfaceMemberArgs)(nil)).Elem()
}

type PortChannelInterfaceMemberInput interface {
	pulumi.Input

	ToPortChannelInterfaceMemberOutput() PortChannelInterfaceMemberOutput
	ToPortChannelInterfaceMemberOutputWithContext(ctx context.Context) PortChannelInterfaceMemberOutput
}

func (*PortChannelInterfaceMember) ElementType() reflect.Type {
	return reflect.TypeOf((**PortChannelInterfaceMember)(nil)).Elem()
}

func (i *PortChannelInterfaceMember) ToPortChannelInterfaceMemberOutput() PortChannelInterfaceMemberOutput {
	return i.ToPortChannelInterfaceMemberOutputWithContext(context.Background())
}

func (i *PortChannelInterfaceMember) ToPortChannelInterfaceMemberOutputWithContext(ctx context.Context) PortChannelInterfaceMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortChannelInterfaceMemberOutput)
}

func (i *PortChannelInterfaceMember) ToOutput(ctx context.Context) pulumix.Output[*PortChannelInterfaceMember] {
	return pulumix.Output[*PortChannelInterfaceMember]{
		OutputState: i.ToPortChannelInterfaceMemberOutputWithContext(ctx).OutputState,
	}
}

// PortChannelInterfaceMemberArrayInput is an input type that accepts PortChannelInterfaceMemberArray and PortChannelInterfaceMemberArrayOutput values.
// You can construct a concrete instance of `PortChannelInterfaceMemberArrayInput` via:
//
//	PortChannelInterfaceMemberArray{ PortChannelInterfaceMemberArgs{...} }
type PortChannelInterfaceMemberArrayInput interface {
	pulumi.Input

	ToPortChannelInterfaceMemberArrayOutput() PortChannelInterfaceMemberArrayOutput
	ToPortChannelInterfaceMemberArrayOutputWithContext(context.Context) PortChannelInterfaceMemberArrayOutput
}

type PortChannelInterfaceMemberArray []PortChannelInterfaceMemberInput

func (PortChannelInterfaceMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortChannelInterfaceMember)(nil)).Elem()
}

func (i PortChannelInterfaceMemberArray) ToPortChannelInterfaceMemberArrayOutput() PortChannelInterfaceMemberArrayOutput {
	return i.ToPortChannelInterfaceMemberArrayOutputWithContext(context.Background())
}

func (i PortChannelInterfaceMemberArray) ToPortChannelInterfaceMemberArrayOutputWithContext(ctx context.Context) PortChannelInterfaceMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortChannelInterfaceMemberArrayOutput)
}

func (i PortChannelInterfaceMemberArray) ToOutput(ctx context.Context) pulumix.Output[[]*PortChannelInterfaceMember] {
	return pulumix.Output[[]*PortChannelInterfaceMember]{
		OutputState: i.ToPortChannelInterfaceMemberArrayOutputWithContext(ctx).OutputState,
	}
}

// PortChannelInterfaceMemberMapInput is an input type that accepts PortChannelInterfaceMemberMap and PortChannelInterfaceMemberMapOutput values.
// You can construct a concrete instance of `PortChannelInterfaceMemberMapInput` via:
//
//	PortChannelInterfaceMemberMap{ "key": PortChannelInterfaceMemberArgs{...} }
type PortChannelInterfaceMemberMapInput interface {
	pulumi.Input

	ToPortChannelInterfaceMemberMapOutput() PortChannelInterfaceMemberMapOutput
	ToPortChannelInterfaceMemberMapOutputWithContext(context.Context) PortChannelInterfaceMemberMapOutput
}

type PortChannelInterfaceMemberMap map[string]PortChannelInterfaceMemberInput

func (PortChannelInterfaceMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortChannelInterfaceMember)(nil)).Elem()
}

func (i PortChannelInterfaceMemberMap) ToPortChannelInterfaceMemberMapOutput() PortChannelInterfaceMemberMapOutput {
	return i.ToPortChannelInterfaceMemberMapOutputWithContext(context.Background())
}

func (i PortChannelInterfaceMemberMap) ToPortChannelInterfaceMemberMapOutputWithContext(ctx context.Context) PortChannelInterfaceMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortChannelInterfaceMemberMapOutput)
}

func (i PortChannelInterfaceMemberMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PortChannelInterfaceMember] {
	return pulumix.Output[map[string]*PortChannelInterfaceMember]{
		OutputState: i.ToPortChannelInterfaceMemberMapOutputWithContext(ctx).OutputState,
	}
}

type PortChannelInterfaceMemberOutput struct{ *pulumi.OutputState }

func (PortChannelInterfaceMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortChannelInterfaceMember)(nil)).Elem()
}

func (o PortChannelInterfaceMemberOutput) ToPortChannelInterfaceMemberOutput() PortChannelInterfaceMemberOutput {
	return o
}

func (o PortChannelInterfaceMemberOutput) ToPortChannelInterfaceMemberOutputWithContext(ctx context.Context) PortChannelInterfaceMemberOutput {
	return o
}

func (o PortChannelInterfaceMemberOutput) ToOutput(ctx context.Context) pulumix.Output[*PortChannelInterfaceMember] {
	return pulumix.Output[*PortChannelInterfaceMember]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o PortChannelInterfaceMemberOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PortChannelInterfaceMember) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
func (o PortChannelInterfaceMemberOutput) InterfaceDn() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterfaceMember) pulumi.StringOutput { return v.InterfaceDn }).(pulumi.StringOutput)
}

// Must match first field in the output of `show intf brief`. Example: `po1`.
func (o PortChannelInterfaceMemberOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortChannelInterfaceMember) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

type PortChannelInterfaceMemberArrayOutput struct{ *pulumi.OutputState }

func (PortChannelInterfaceMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortChannelInterfaceMember)(nil)).Elem()
}

func (o PortChannelInterfaceMemberArrayOutput) ToPortChannelInterfaceMemberArrayOutput() PortChannelInterfaceMemberArrayOutput {
	return o
}

func (o PortChannelInterfaceMemberArrayOutput) ToPortChannelInterfaceMemberArrayOutputWithContext(ctx context.Context) PortChannelInterfaceMemberArrayOutput {
	return o
}

func (o PortChannelInterfaceMemberArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PortChannelInterfaceMember] {
	return pulumix.Output[[]*PortChannelInterfaceMember]{
		OutputState: o.OutputState,
	}
}

func (o PortChannelInterfaceMemberArrayOutput) Index(i pulumi.IntInput) PortChannelInterfaceMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PortChannelInterfaceMember {
		return vs[0].([]*PortChannelInterfaceMember)[vs[1].(int)]
	}).(PortChannelInterfaceMemberOutput)
}

type PortChannelInterfaceMemberMapOutput struct{ *pulumi.OutputState }

func (PortChannelInterfaceMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortChannelInterfaceMember)(nil)).Elem()
}

func (o PortChannelInterfaceMemberMapOutput) ToPortChannelInterfaceMemberMapOutput() PortChannelInterfaceMemberMapOutput {
	return o
}

func (o PortChannelInterfaceMemberMapOutput) ToPortChannelInterfaceMemberMapOutputWithContext(ctx context.Context) PortChannelInterfaceMemberMapOutput {
	return o
}

func (o PortChannelInterfaceMemberMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PortChannelInterfaceMember] {
	return pulumix.Output[map[string]*PortChannelInterfaceMember]{
		OutputState: o.OutputState,
	}
}

func (o PortChannelInterfaceMemberMapOutput) MapIndex(k pulumi.StringInput) PortChannelInterfaceMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PortChannelInterfaceMember {
		return vs[0].(map[string]*PortChannelInterfaceMember)[vs[1].(string)]
	}).(PortChannelInterfaceMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortChannelInterfaceMemberInput)(nil)).Elem(), &PortChannelInterfaceMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortChannelInterfaceMemberArrayInput)(nil)).Elem(), PortChannelInterfaceMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortChannelInterfaceMemberMapInput)(nil)).Elem(), PortChannelInterfaceMemberMap{})
	pulumi.RegisterOutputType(PortChannelInterfaceMemberOutput{})
	pulumi.RegisterOutputType(PortChannelInterfaceMemberArrayOutput{})
	pulumi.RegisterOutputType(PortChannelInterfaceMemberMapOutput{})
}
