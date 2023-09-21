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

type Ipv4Interface struct {
	pulumi.CustomResourceState

	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	DropGlean pulumi.StringOutput `pulumi:"dropGlean"`
	// ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	Forward pulumi.StringOutput `pulumi:"forward"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringOutput `pulumi:"interfaceId"`
	// IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
	// Default value: `unspecified`
	Unnumbered pulumi.StringOutput `pulumi:"unnumbered"`
	// URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
	// `strict-allow-vni-hosts` - Default value: `disabled`
	Urpf pulumi.StringOutput `pulumi:"urpf"`
	// VRF name.
	Vrf pulumi.StringOutput `pulumi:"vrf"`
}

// NewIpv4Interface registers a new resource with the given unique name, arguments, and options.
func NewIpv4Interface(ctx *pulumi.Context,
	name string, args *Ipv4InterfaceArgs, opts ...pulumi.ResourceOption) (*Ipv4Interface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'InterfaceId'")
	}
	if args.Vrf == nil {
		return nil, errors.New("invalid value for required argument 'Vrf'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ipv4Interface
	err := ctx.RegisterResource("nxos:nxos/ipv4Interface:Ipv4Interface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv4Interface gets an existing Ipv4Interface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv4Interface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv4InterfaceState, opts ...pulumi.ResourceOption) (*Ipv4Interface, error) {
	var resource Ipv4Interface
	err := ctx.ReadResource("nxos:nxos/ipv4Interface:Ipv4Interface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv4Interface resources.
type ipv4InterfaceState struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	DropGlean *string `pulumi:"dropGlean"`
	// ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	Forward *string `pulumi:"forward"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId *string `pulumi:"interfaceId"`
	// IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
	// Default value: `unspecified`
	Unnumbered *string `pulumi:"unnumbered"`
	// URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
	// `strict-allow-vni-hosts` - Default value: `disabled`
	Urpf *string `pulumi:"urpf"`
	// VRF name.
	Vrf *string `pulumi:"vrf"`
}

type Ipv4InterfaceState struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	DropGlean pulumi.StringPtrInput
	// ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	Forward pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringPtrInput
	// IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
	// Default value: `unspecified`
	Unnumbered pulumi.StringPtrInput
	// URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
	// `strict-allow-vni-hosts` - Default value: `disabled`
	Urpf pulumi.StringPtrInput
	// VRF name.
	Vrf pulumi.StringPtrInput
}

func (Ipv4InterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4InterfaceState)(nil)).Elem()
}

type ipv4InterfaceArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	DropGlean *string `pulumi:"dropGlean"`
	// ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	Forward *string `pulumi:"forward"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
	// IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
	// Default value: `unspecified`
	Unnumbered *string `pulumi:"unnumbered"`
	// URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
	// `strict-allow-vni-hosts` - Default value: `disabled`
	Urpf *string `pulumi:"urpf"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

// The set of arguments for constructing a Ipv4Interface resource.
type Ipv4InterfaceArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	DropGlean pulumi.StringPtrInput
	// ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
	Forward pulumi.StringPtrInput
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringInput
	// IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
	// Default value: `unspecified`
	Unnumbered pulumi.StringPtrInput
	// URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
	// `strict-allow-vni-hosts` - Default value: `disabled`
	Urpf pulumi.StringPtrInput
	// VRF name.
	Vrf pulumi.StringInput
}

func (Ipv4InterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4InterfaceArgs)(nil)).Elem()
}

type Ipv4InterfaceInput interface {
	pulumi.Input

	ToIpv4InterfaceOutput() Ipv4InterfaceOutput
	ToIpv4InterfaceOutputWithContext(ctx context.Context) Ipv4InterfaceOutput
}

func (*Ipv4Interface) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Interface)(nil)).Elem()
}

func (i *Ipv4Interface) ToIpv4InterfaceOutput() Ipv4InterfaceOutput {
	return i.ToIpv4InterfaceOutputWithContext(context.Background())
}

func (i *Ipv4Interface) ToIpv4InterfaceOutputWithContext(ctx context.Context) Ipv4InterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4InterfaceOutput)
}

func (i *Ipv4Interface) ToOutput(ctx context.Context) pulumix.Output[*Ipv4Interface] {
	return pulumix.Output[*Ipv4Interface]{
		OutputState: i.ToIpv4InterfaceOutputWithContext(ctx).OutputState,
	}
}

// Ipv4InterfaceArrayInput is an input type that accepts Ipv4InterfaceArray and Ipv4InterfaceArrayOutput values.
// You can construct a concrete instance of `Ipv4InterfaceArrayInput` via:
//
//	Ipv4InterfaceArray{ Ipv4InterfaceArgs{...} }
type Ipv4InterfaceArrayInput interface {
	pulumi.Input

	ToIpv4InterfaceArrayOutput() Ipv4InterfaceArrayOutput
	ToIpv4InterfaceArrayOutputWithContext(context.Context) Ipv4InterfaceArrayOutput
}

type Ipv4InterfaceArray []Ipv4InterfaceInput

func (Ipv4InterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Interface)(nil)).Elem()
}

func (i Ipv4InterfaceArray) ToIpv4InterfaceArrayOutput() Ipv4InterfaceArrayOutput {
	return i.ToIpv4InterfaceArrayOutputWithContext(context.Background())
}

func (i Ipv4InterfaceArray) ToIpv4InterfaceArrayOutputWithContext(ctx context.Context) Ipv4InterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4InterfaceArrayOutput)
}

func (i Ipv4InterfaceArray) ToOutput(ctx context.Context) pulumix.Output[[]*Ipv4Interface] {
	return pulumix.Output[[]*Ipv4Interface]{
		OutputState: i.ToIpv4InterfaceArrayOutputWithContext(ctx).OutputState,
	}
}

// Ipv4InterfaceMapInput is an input type that accepts Ipv4InterfaceMap and Ipv4InterfaceMapOutput values.
// You can construct a concrete instance of `Ipv4InterfaceMapInput` via:
//
//	Ipv4InterfaceMap{ "key": Ipv4InterfaceArgs{...} }
type Ipv4InterfaceMapInput interface {
	pulumi.Input

	ToIpv4InterfaceMapOutput() Ipv4InterfaceMapOutput
	ToIpv4InterfaceMapOutputWithContext(context.Context) Ipv4InterfaceMapOutput
}

type Ipv4InterfaceMap map[string]Ipv4InterfaceInput

func (Ipv4InterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Interface)(nil)).Elem()
}

func (i Ipv4InterfaceMap) ToIpv4InterfaceMapOutput() Ipv4InterfaceMapOutput {
	return i.ToIpv4InterfaceMapOutputWithContext(context.Background())
}

func (i Ipv4InterfaceMap) ToIpv4InterfaceMapOutputWithContext(ctx context.Context) Ipv4InterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4InterfaceMapOutput)
}

func (i Ipv4InterfaceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ipv4Interface] {
	return pulumix.Output[map[string]*Ipv4Interface]{
		OutputState: i.ToIpv4InterfaceMapOutputWithContext(ctx).OutputState,
	}
}

type Ipv4InterfaceOutput struct{ *pulumi.OutputState }

func (Ipv4InterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Interface)(nil)).Elem()
}

func (o Ipv4InterfaceOutput) ToIpv4InterfaceOutput() Ipv4InterfaceOutput {
	return o
}

func (o Ipv4InterfaceOutput) ToIpv4InterfaceOutputWithContext(ctx context.Context) Ipv4InterfaceOutput {
	return o
}

func (o Ipv4InterfaceOutput) ToOutput(ctx context.Context) pulumix.Output[*Ipv4Interface] {
	return pulumix.Output[*Ipv4Interface]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o Ipv4InterfaceOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv4Interface) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
func (o Ipv4InterfaceOutput) DropGlean() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Interface) pulumi.StringOutput { return v.DropGlean }).(pulumi.StringOutput)
}

// ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
func (o Ipv4InterfaceOutput) Forward() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Interface) pulumi.StringOutput { return v.Forward }).(pulumi.StringOutput)
}

// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
func (o Ipv4InterfaceOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Interface) pulumi.StringOutput { return v.InterfaceId }).(pulumi.StringOutput)
}

// IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
// Default value: `unspecified`
func (o Ipv4InterfaceOutput) Unnumbered() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Interface) pulumi.StringOutput { return v.Unnumbered }).(pulumi.StringOutput)
}

// URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
// `strict-allow-vni-hosts` - Default value: `disabled`
func (o Ipv4InterfaceOutput) Urpf() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Interface) pulumi.StringOutput { return v.Urpf }).(pulumi.StringOutput)
}

// VRF name.
func (o Ipv4InterfaceOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Interface) pulumi.StringOutput { return v.Vrf }).(pulumi.StringOutput)
}

type Ipv4InterfaceArrayOutput struct{ *pulumi.OutputState }

func (Ipv4InterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Interface)(nil)).Elem()
}

func (o Ipv4InterfaceArrayOutput) ToIpv4InterfaceArrayOutput() Ipv4InterfaceArrayOutput {
	return o
}

func (o Ipv4InterfaceArrayOutput) ToIpv4InterfaceArrayOutputWithContext(ctx context.Context) Ipv4InterfaceArrayOutput {
	return o
}

func (o Ipv4InterfaceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Ipv4Interface] {
	return pulumix.Output[[]*Ipv4Interface]{
		OutputState: o.OutputState,
	}
}

func (o Ipv4InterfaceArrayOutput) Index(i pulumi.IntInput) Ipv4InterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv4Interface {
		return vs[0].([]*Ipv4Interface)[vs[1].(int)]
	}).(Ipv4InterfaceOutput)
}

type Ipv4InterfaceMapOutput struct{ *pulumi.OutputState }

func (Ipv4InterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Interface)(nil)).Elem()
}

func (o Ipv4InterfaceMapOutput) ToIpv4InterfaceMapOutput() Ipv4InterfaceMapOutput {
	return o
}

func (o Ipv4InterfaceMapOutput) ToIpv4InterfaceMapOutputWithContext(ctx context.Context) Ipv4InterfaceMapOutput {
	return o
}

func (o Ipv4InterfaceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ipv4Interface] {
	return pulumix.Output[map[string]*Ipv4Interface]{
		OutputState: o.OutputState,
	}
}

func (o Ipv4InterfaceMapOutput) MapIndex(k pulumi.StringInput) Ipv4InterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv4Interface {
		return vs[0].(map[string]*Ipv4Interface)[vs[1].(string)]
	}).(Ipv4InterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4InterfaceInput)(nil)).Elem(), &Ipv4Interface{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4InterfaceArrayInput)(nil)).Elem(), Ipv4InterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4InterfaceMapInput)(nil)).Elem(), Ipv4InterfaceMap{})
	pulumi.RegisterOutputType(Ipv4InterfaceOutput{})
	pulumi.RegisterOutputType(Ipv4InterfaceArrayOutput{})
	pulumi.RegisterOutputType(Ipv4InterfaceMapOutput{})
}