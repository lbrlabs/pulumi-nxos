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

type DhcpRelayInterface struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
}

// NewDhcpRelayInterface registers a new resource with the given unique name, arguments, and options.
func NewDhcpRelayInterface(ctx *pulumi.Context,
	name string, args *DhcpRelayInterfaceArgs, opts ...pulumi.ResourceOption) (*DhcpRelayInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DhcpRelayInterface
	err := ctx.RegisterResource("nxos:nxos/dhcpRelayInterface:DhcpRelayInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpRelayInterface gets an existing DhcpRelayInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpRelayInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpRelayInterfaceState, opts ...pulumi.ResourceOption) (*DhcpRelayInterface, error) {
	var resource DhcpRelayInterface
	err := ctx.ReadResource("nxos:nxos/dhcpRelayInterface:DhcpRelayInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpRelayInterface resources.
type dhcpRelayInterfaceState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId *string `pulumi:"interfaceId"`
}

type DhcpRelayInterfaceState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringPtrInput
}

func (DhcpRelayInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpRelayInterfaceState)(nil)).Elem()
}

type dhcpRelayInterfaceArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
}

// The set of arguments for constructing a DhcpRelayInterface resource.
type DhcpRelayInterfaceArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringInput
}

func (DhcpRelayInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpRelayInterfaceArgs)(nil)).Elem()
}

type DhcpRelayInterfaceInput interface {
	pulumi.Input

	ToDhcpRelayInterfaceOutput() DhcpRelayInterfaceOutput
	ToDhcpRelayInterfaceOutputWithContext(ctx context.Context) DhcpRelayInterfaceOutput
}

func (*DhcpRelayInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpRelayInterface)(nil)).Elem()
}

func (i *DhcpRelayInterface) ToDhcpRelayInterfaceOutput() DhcpRelayInterfaceOutput {
	return i.ToDhcpRelayInterfaceOutputWithContext(context.Background())
}

func (i *DhcpRelayInterface) ToDhcpRelayInterfaceOutputWithContext(ctx context.Context) DhcpRelayInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpRelayInterfaceOutput)
}

func (i *DhcpRelayInterface) ToOutput(ctx context.Context) pulumix.Output[*DhcpRelayInterface] {
	return pulumix.Output[*DhcpRelayInterface]{
		OutputState: i.ToDhcpRelayInterfaceOutputWithContext(ctx).OutputState,
	}
}

// DhcpRelayInterfaceArrayInput is an input type that accepts DhcpRelayInterfaceArray and DhcpRelayInterfaceArrayOutput values.
// You can construct a concrete instance of `DhcpRelayInterfaceArrayInput` via:
//
//	DhcpRelayInterfaceArray{ DhcpRelayInterfaceArgs{...} }
type DhcpRelayInterfaceArrayInput interface {
	pulumi.Input

	ToDhcpRelayInterfaceArrayOutput() DhcpRelayInterfaceArrayOutput
	ToDhcpRelayInterfaceArrayOutputWithContext(context.Context) DhcpRelayInterfaceArrayOutput
}

type DhcpRelayInterfaceArray []DhcpRelayInterfaceInput

func (DhcpRelayInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpRelayInterface)(nil)).Elem()
}

func (i DhcpRelayInterfaceArray) ToDhcpRelayInterfaceArrayOutput() DhcpRelayInterfaceArrayOutput {
	return i.ToDhcpRelayInterfaceArrayOutputWithContext(context.Background())
}

func (i DhcpRelayInterfaceArray) ToDhcpRelayInterfaceArrayOutputWithContext(ctx context.Context) DhcpRelayInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpRelayInterfaceArrayOutput)
}

func (i DhcpRelayInterfaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*DhcpRelayInterface] {
	return pulumix.Output[[]*DhcpRelayInterface]{
		OutputState: i.ToDhcpRelayInterfaceArrayOutputWithContext(ctx).OutputState,
	}
}

// DhcpRelayInterfaceMapInput is an input type that accepts DhcpRelayInterfaceMap and DhcpRelayInterfaceMapOutput values.
// You can construct a concrete instance of `DhcpRelayInterfaceMapInput` via:
//
//	DhcpRelayInterfaceMap{ "key": DhcpRelayInterfaceArgs{...} }
type DhcpRelayInterfaceMapInput interface {
	pulumi.Input

	ToDhcpRelayInterfaceMapOutput() DhcpRelayInterfaceMapOutput
	ToDhcpRelayInterfaceMapOutputWithContext(context.Context) DhcpRelayInterfaceMapOutput
}

type DhcpRelayInterfaceMap map[string]DhcpRelayInterfaceInput

func (DhcpRelayInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpRelayInterface)(nil)).Elem()
}

func (i DhcpRelayInterfaceMap) ToDhcpRelayInterfaceMapOutput() DhcpRelayInterfaceMapOutput {
	return i.ToDhcpRelayInterfaceMapOutputWithContext(context.Background())
}

func (i DhcpRelayInterfaceMap) ToDhcpRelayInterfaceMapOutputWithContext(ctx context.Context) DhcpRelayInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpRelayInterfaceMapOutput)
}

func (i DhcpRelayInterfaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DhcpRelayInterface] {
	return pulumix.Output[map[string]*DhcpRelayInterface]{
		OutputState: i.ToDhcpRelayInterfaceMapOutputWithContext(ctx).OutputState,
	}
}

type DhcpRelayInterfaceOutput struct{ *pulumi.OutputState }

func (DhcpRelayInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpRelayInterface)(nil)).Elem()
}

func (o DhcpRelayInterfaceOutput) ToDhcpRelayInterfaceOutput() DhcpRelayInterfaceOutput {
	return o
}

func (o DhcpRelayInterfaceOutput) ToDhcpRelayInterfaceOutputWithContext(ctx context.Context) DhcpRelayInterfaceOutput {
	return o
}

func (o DhcpRelayInterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*DhcpRelayInterface] {
	return pulumix.Output[*DhcpRelayInterface]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o DhcpRelayInterfaceOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DhcpRelayInterface) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
func (o DhcpRelayInterfaceOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpRelayInterface) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

type DhcpRelayInterfaceArrayOutput struct{ *pulumi.OutputState }

func (DhcpRelayInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpRelayInterface)(nil)).Elem()
}

func (o DhcpRelayInterfaceArrayOutput) ToDhcpRelayInterfaceArrayOutput() DhcpRelayInterfaceArrayOutput {
	return o
}

func (o DhcpRelayInterfaceArrayOutput) ToDhcpRelayInterfaceArrayOutputWithContext(ctx context.Context) DhcpRelayInterfaceArrayOutput {
	return o
}

func (o DhcpRelayInterfaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DhcpRelayInterface] {
	return pulumix.Output[[]*DhcpRelayInterface]{
		OutputState: o.OutputState,
	}
}

func (o DhcpRelayInterfaceArrayOutput) Index(i pulumi.IntInput) DhcpRelayInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpRelayInterface {
		return vs[0].([]*DhcpRelayInterface)[vs[1].(int)]
	}).(DhcpRelayInterfaceOutput)
}

type DhcpRelayInterfaceMapOutput struct{ *pulumi.OutputState }

func (DhcpRelayInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpRelayInterface)(nil)).Elem()
}

func (o DhcpRelayInterfaceMapOutput) ToDhcpRelayInterfaceMapOutput() DhcpRelayInterfaceMapOutput {
	return o
}

func (o DhcpRelayInterfaceMapOutput) ToDhcpRelayInterfaceMapOutputWithContext(ctx context.Context) DhcpRelayInterfaceMapOutput {
	return o
}

func (o DhcpRelayInterfaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DhcpRelayInterface] {
	return pulumix.Output[map[string]*DhcpRelayInterface]{
		OutputState: o.OutputState,
	}
}

func (o DhcpRelayInterfaceMapOutput) MapIndex(k pulumi.StringInput) DhcpRelayInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpRelayInterface {
		return vs[0].(map[string]*DhcpRelayInterface)[vs[1].(string)]
	}).(DhcpRelayInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpRelayInterfaceInput)(nil)).Elem(), &DhcpRelayInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpRelayInterfaceArrayInput)(nil)).Elem(), DhcpRelayInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpRelayInterfaceMapInput)(nil)).Elem(), DhcpRelayInterfaceMap{})
	pulumi.RegisterOutputType(DhcpRelayInterfaceOutput{})
	pulumi.RegisterOutputType(DhcpRelayInterfaceArrayOutput{})
	pulumi.RegisterOutputType(DhcpRelayInterfaceMapOutput{})
}
