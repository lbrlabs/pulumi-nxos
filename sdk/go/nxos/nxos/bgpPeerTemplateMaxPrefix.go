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

type BgpPeerTemplateMaxPrefix struct {
	pulumi.CustomResourceState

	// Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
	Action pulumi.StringOutput `pulumi:"action"`
	// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
	// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
	// `ipv4-mdt` - Default value: `ipv4-ucast`
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// Autonomous system number.
	Asn pulumi.StringOutput `pulumi:"asn"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
	MaximumPrefix pulumi.IntPtrOutput `pulumi:"maximumPrefix"`
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
	RestartTime pulumi.IntOutput `pulumi:"restartTime"`
	// Peer template name.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
	// value: `0`
	Threshold pulumi.IntOutput `pulumi:"threshold"`
}

// NewBgpPeerTemplateMaxPrefix registers a new resource with the given unique name, arguments, and options.
func NewBgpPeerTemplateMaxPrefix(ctx *pulumi.Context,
	name string, args *BgpPeerTemplateMaxPrefixArgs, opts ...pulumi.ResourceOption) (*BgpPeerTemplateMaxPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.Asn == nil {
		return nil, errors.New("invalid value for required argument 'Asn'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BgpPeerTemplateMaxPrefix
	err := ctx.RegisterResource("nxos:nxos/bgpPeerTemplateMaxPrefix:BgpPeerTemplateMaxPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpPeerTemplateMaxPrefix gets an existing BgpPeerTemplateMaxPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpPeerTemplateMaxPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpPeerTemplateMaxPrefixState, opts ...pulumi.ResourceOption) (*BgpPeerTemplateMaxPrefix, error) {
	var resource BgpPeerTemplateMaxPrefix
	err := ctx.ReadResource("nxos:nxos/bgpPeerTemplateMaxPrefix:BgpPeerTemplateMaxPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpPeerTemplateMaxPrefix resources.
type bgpPeerTemplateMaxPrefixState struct {
	// Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
	Action *string `pulumi:"action"`
	// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
	// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
	// `ipv4-mdt` - Default value: `ipv4-ucast`
	AddressFamily *string `pulumi:"addressFamily"`
	// Autonomous system number.
	Asn *string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
	MaximumPrefix *int `pulumi:"maximumPrefix"`
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
	RestartTime *int `pulumi:"restartTime"`
	// Peer template name.
	TemplateName *string `pulumi:"templateName"`
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
	// value: `0`
	Threshold *int `pulumi:"threshold"`
}

type BgpPeerTemplateMaxPrefixState struct {
	// Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
	Action pulumi.StringPtrInput
	// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
	// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
	// `ipv4-mdt` - Default value: `ipv4-ucast`
	AddressFamily pulumi.StringPtrInput
	// Autonomous system number.
	Asn pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
	MaximumPrefix pulumi.IntPtrInput
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
	RestartTime pulumi.IntPtrInput
	// Peer template name.
	TemplateName pulumi.StringPtrInput
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
	// value: `0`
	Threshold pulumi.IntPtrInput
}

func (BgpPeerTemplateMaxPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerTemplateMaxPrefixState)(nil)).Elem()
}

type bgpPeerTemplateMaxPrefixArgs struct {
	// Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
	Action *string `pulumi:"action"`
	// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
	// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
	// `ipv4-mdt` - Default value: `ipv4-ucast`
	AddressFamily string `pulumi:"addressFamily"`
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
	MaximumPrefix *int `pulumi:"maximumPrefix"`
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
	RestartTime *int `pulumi:"restartTime"`
	// Peer template name.
	TemplateName string `pulumi:"templateName"`
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
	// value: `0`
	Threshold *int `pulumi:"threshold"`
}

// The set of arguments for constructing a BgpPeerTemplateMaxPrefix resource.
type BgpPeerTemplateMaxPrefixArgs struct {
	// Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
	Action pulumi.StringPtrInput
	// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
	// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
	// `ipv4-mdt` - Default value: `ipv4-ucast`
	AddressFamily pulumi.StringInput
	// Autonomous system number.
	Asn pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
	MaximumPrefix pulumi.IntPtrInput
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
	RestartTime pulumi.IntPtrInput
	// Peer template name.
	TemplateName pulumi.StringInput
	// The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
	// value: `0`
	Threshold pulumi.IntPtrInput
}

func (BgpPeerTemplateMaxPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerTemplateMaxPrefixArgs)(nil)).Elem()
}

type BgpPeerTemplateMaxPrefixInput interface {
	pulumi.Input

	ToBgpPeerTemplateMaxPrefixOutput() BgpPeerTemplateMaxPrefixOutput
	ToBgpPeerTemplateMaxPrefixOutputWithContext(ctx context.Context) BgpPeerTemplateMaxPrefixOutput
}

func (*BgpPeerTemplateMaxPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpPeerTemplateMaxPrefix)(nil)).Elem()
}

func (i *BgpPeerTemplateMaxPrefix) ToBgpPeerTemplateMaxPrefixOutput() BgpPeerTemplateMaxPrefixOutput {
	return i.ToBgpPeerTemplateMaxPrefixOutputWithContext(context.Background())
}

func (i *BgpPeerTemplateMaxPrefix) ToBgpPeerTemplateMaxPrefixOutputWithContext(ctx context.Context) BgpPeerTemplateMaxPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerTemplateMaxPrefixOutput)
}

func (i *BgpPeerTemplateMaxPrefix) ToOutput(ctx context.Context) pulumix.Output[*BgpPeerTemplateMaxPrefix] {
	return pulumix.Output[*BgpPeerTemplateMaxPrefix]{
		OutputState: i.ToBgpPeerTemplateMaxPrefixOutputWithContext(ctx).OutputState,
	}
}

// BgpPeerTemplateMaxPrefixArrayInput is an input type that accepts BgpPeerTemplateMaxPrefixArray and BgpPeerTemplateMaxPrefixArrayOutput values.
// You can construct a concrete instance of `BgpPeerTemplateMaxPrefixArrayInput` via:
//
//	BgpPeerTemplateMaxPrefixArray{ BgpPeerTemplateMaxPrefixArgs{...} }
type BgpPeerTemplateMaxPrefixArrayInput interface {
	pulumi.Input

	ToBgpPeerTemplateMaxPrefixArrayOutput() BgpPeerTemplateMaxPrefixArrayOutput
	ToBgpPeerTemplateMaxPrefixArrayOutputWithContext(context.Context) BgpPeerTemplateMaxPrefixArrayOutput
}

type BgpPeerTemplateMaxPrefixArray []BgpPeerTemplateMaxPrefixInput

func (BgpPeerTemplateMaxPrefixArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpPeerTemplateMaxPrefix)(nil)).Elem()
}

func (i BgpPeerTemplateMaxPrefixArray) ToBgpPeerTemplateMaxPrefixArrayOutput() BgpPeerTemplateMaxPrefixArrayOutput {
	return i.ToBgpPeerTemplateMaxPrefixArrayOutputWithContext(context.Background())
}

func (i BgpPeerTemplateMaxPrefixArray) ToBgpPeerTemplateMaxPrefixArrayOutputWithContext(ctx context.Context) BgpPeerTemplateMaxPrefixArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerTemplateMaxPrefixArrayOutput)
}

func (i BgpPeerTemplateMaxPrefixArray) ToOutput(ctx context.Context) pulumix.Output[[]*BgpPeerTemplateMaxPrefix] {
	return pulumix.Output[[]*BgpPeerTemplateMaxPrefix]{
		OutputState: i.ToBgpPeerTemplateMaxPrefixArrayOutputWithContext(ctx).OutputState,
	}
}

// BgpPeerTemplateMaxPrefixMapInput is an input type that accepts BgpPeerTemplateMaxPrefixMap and BgpPeerTemplateMaxPrefixMapOutput values.
// You can construct a concrete instance of `BgpPeerTemplateMaxPrefixMapInput` via:
//
//	BgpPeerTemplateMaxPrefixMap{ "key": BgpPeerTemplateMaxPrefixArgs{...} }
type BgpPeerTemplateMaxPrefixMapInput interface {
	pulumi.Input

	ToBgpPeerTemplateMaxPrefixMapOutput() BgpPeerTemplateMaxPrefixMapOutput
	ToBgpPeerTemplateMaxPrefixMapOutputWithContext(context.Context) BgpPeerTemplateMaxPrefixMapOutput
}

type BgpPeerTemplateMaxPrefixMap map[string]BgpPeerTemplateMaxPrefixInput

func (BgpPeerTemplateMaxPrefixMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpPeerTemplateMaxPrefix)(nil)).Elem()
}

func (i BgpPeerTemplateMaxPrefixMap) ToBgpPeerTemplateMaxPrefixMapOutput() BgpPeerTemplateMaxPrefixMapOutput {
	return i.ToBgpPeerTemplateMaxPrefixMapOutputWithContext(context.Background())
}

func (i BgpPeerTemplateMaxPrefixMap) ToBgpPeerTemplateMaxPrefixMapOutputWithContext(ctx context.Context) BgpPeerTemplateMaxPrefixMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerTemplateMaxPrefixMapOutput)
}

func (i BgpPeerTemplateMaxPrefixMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpPeerTemplateMaxPrefix] {
	return pulumix.Output[map[string]*BgpPeerTemplateMaxPrefix]{
		OutputState: i.ToBgpPeerTemplateMaxPrefixMapOutputWithContext(ctx).OutputState,
	}
}

type BgpPeerTemplateMaxPrefixOutput struct{ *pulumi.OutputState }

func (BgpPeerTemplateMaxPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpPeerTemplateMaxPrefix)(nil)).Elem()
}

func (o BgpPeerTemplateMaxPrefixOutput) ToBgpPeerTemplateMaxPrefixOutput() BgpPeerTemplateMaxPrefixOutput {
	return o
}

func (o BgpPeerTemplateMaxPrefixOutput) ToBgpPeerTemplateMaxPrefixOutputWithContext(ctx context.Context) BgpPeerTemplateMaxPrefixOutput {
	return o
}

func (o BgpPeerTemplateMaxPrefixOutput) ToOutput(ctx context.Context) pulumix.Output[*BgpPeerTemplateMaxPrefix] {
	return pulumix.Output[*BgpPeerTemplateMaxPrefix]{
		OutputState: o.OutputState,
	}
}

// Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
func (o BgpPeerTemplateMaxPrefixOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
// `ipv4-mdt` - Default value: `ipv4-ucast`
func (o BgpPeerTemplateMaxPrefixOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// Autonomous system number.
func (o BgpPeerTemplateMaxPrefixOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.StringOutput { return v.Asn }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o BgpPeerTemplateMaxPrefixOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
func (o BgpPeerTemplateMaxPrefixOutput) MaximumPrefix() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.IntPtrOutput { return v.MaximumPrefix }).(pulumi.IntPtrOutput)
}

// The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
func (o BgpPeerTemplateMaxPrefixOutput) RestartTime() pulumi.IntOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.IntOutput { return v.RestartTime }).(pulumi.IntOutput)
}

// Peer template name.
func (o BgpPeerTemplateMaxPrefixOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

// The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
// value: `0`
func (o BgpPeerTemplateMaxPrefixOutput) Threshold() pulumi.IntOutput {
	return o.ApplyT(func(v *BgpPeerTemplateMaxPrefix) pulumi.IntOutput { return v.Threshold }).(pulumi.IntOutput)
}

type BgpPeerTemplateMaxPrefixArrayOutput struct{ *pulumi.OutputState }

func (BgpPeerTemplateMaxPrefixArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpPeerTemplateMaxPrefix)(nil)).Elem()
}

func (o BgpPeerTemplateMaxPrefixArrayOutput) ToBgpPeerTemplateMaxPrefixArrayOutput() BgpPeerTemplateMaxPrefixArrayOutput {
	return o
}

func (o BgpPeerTemplateMaxPrefixArrayOutput) ToBgpPeerTemplateMaxPrefixArrayOutputWithContext(ctx context.Context) BgpPeerTemplateMaxPrefixArrayOutput {
	return o
}

func (o BgpPeerTemplateMaxPrefixArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BgpPeerTemplateMaxPrefix] {
	return pulumix.Output[[]*BgpPeerTemplateMaxPrefix]{
		OutputState: o.OutputState,
	}
}

func (o BgpPeerTemplateMaxPrefixArrayOutput) Index(i pulumi.IntInput) BgpPeerTemplateMaxPrefixOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpPeerTemplateMaxPrefix {
		return vs[0].([]*BgpPeerTemplateMaxPrefix)[vs[1].(int)]
	}).(BgpPeerTemplateMaxPrefixOutput)
}

type BgpPeerTemplateMaxPrefixMapOutput struct{ *pulumi.OutputState }

func (BgpPeerTemplateMaxPrefixMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpPeerTemplateMaxPrefix)(nil)).Elem()
}

func (o BgpPeerTemplateMaxPrefixMapOutput) ToBgpPeerTemplateMaxPrefixMapOutput() BgpPeerTemplateMaxPrefixMapOutput {
	return o
}

func (o BgpPeerTemplateMaxPrefixMapOutput) ToBgpPeerTemplateMaxPrefixMapOutputWithContext(ctx context.Context) BgpPeerTemplateMaxPrefixMapOutput {
	return o
}

func (o BgpPeerTemplateMaxPrefixMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpPeerTemplateMaxPrefix] {
	return pulumix.Output[map[string]*BgpPeerTemplateMaxPrefix]{
		OutputState: o.OutputState,
	}
}

func (o BgpPeerTemplateMaxPrefixMapOutput) MapIndex(k pulumi.StringInput) BgpPeerTemplateMaxPrefixOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpPeerTemplateMaxPrefix {
		return vs[0].(map[string]*BgpPeerTemplateMaxPrefix)[vs[1].(string)]
	}).(BgpPeerTemplateMaxPrefixOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerTemplateMaxPrefixInput)(nil)).Elem(), &BgpPeerTemplateMaxPrefix{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerTemplateMaxPrefixArrayInput)(nil)).Elem(), BgpPeerTemplateMaxPrefixArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerTemplateMaxPrefixMapInput)(nil)).Elem(), BgpPeerTemplateMaxPrefixMap{})
	pulumi.RegisterOutputType(BgpPeerTemplateMaxPrefixOutput{})
	pulumi.RegisterOutputType(BgpPeerTemplateMaxPrefixArrayOutput{})
	pulumi.RegisterOutputType(BgpPeerTemplateMaxPrefixMapOutput{})
}