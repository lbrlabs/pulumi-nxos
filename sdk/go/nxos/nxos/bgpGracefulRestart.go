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

type BgpGracefulRestart struct {
	pulumi.CustomResourceState

	// Autonomous system number.
	Asn pulumi.StringOutput `pulumi:"asn"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
	RestartInterval pulumi.IntOutput `pulumi:"restartInterval"`
	// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
	StaleInterval pulumi.IntOutput `pulumi:"staleInterval"`
	// VRF name.
	Vrf pulumi.StringOutput `pulumi:"vrf"`
}

// NewBgpGracefulRestart registers a new resource with the given unique name, arguments, and options.
func NewBgpGracefulRestart(ctx *pulumi.Context,
	name string, args *BgpGracefulRestartArgs, opts ...pulumi.ResourceOption) (*BgpGracefulRestart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Asn == nil {
		return nil, errors.New("invalid value for required argument 'Asn'")
	}
	if args.Vrf == nil {
		return nil, errors.New("invalid value for required argument 'Vrf'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BgpGracefulRestart
	err := ctx.RegisterResource("nxos:nxos/bgpGracefulRestart:BgpGracefulRestart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpGracefulRestart gets an existing BgpGracefulRestart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpGracefulRestart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpGracefulRestartState, opts ...pulumi.ResourceOption) (*BgpGracefulRestart, error) {
	var resource BgpGracefulRestart
	err := ctx.ReadResource("nxos:nxos/bgpGracefulRestart:BgpGracefulRestart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpGracefulRestart resources.
type bgpGracefulRestartState struct {
	// Autonomous system number.
	Asn *string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
	RestartInterval *int `pulumi:"restartInterval"`
	// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
	StaleInterval *int `pulumi:"staleInterval"`
	// VRF name.
	Vrf *string `pulumi:"vrf"`
}

type BgpGracefulRestartState struct {
	// Autonomous system number.
	Asn pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
	RestartInterval pulumi.IntPtrInput
	// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
	StaleInterval pulumi.IntPtrInput
	// VRF name.
	Vrf pulumi.StringPtrInput
}

func (BgpGracefulRestartState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpGracefulRestartState)(nil)).Elem()
}

type bgpGracefulRestartArgs struct {
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
	RestartInterval *int `pulumi:"restartInterval"`
	// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
	StaleInterval *int `pulumi:"staleInterval"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

// The set of arguments for constructing a BgpGracefulRestart resource.
type BgpGracefulRestartArgs struct {
	// Autonomous system number.
	Asn pulumi.StringInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
	RestartInterval pulumi.IntPtrInput
	// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
	StaleInterval pulumi.IntPtrInput
	// VRF name.
	Vrf pulumi.StringInput
}

func (BgpGracefulRestartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpGracefulRestartArgs)(nil)).Elem()
}

type BgpGracefulRestartInput interface {
	pulumi.Input

	ToBgpGracefulRestartOutput() BgpGracefulRestartOutput
	ToBgpGracefulRestartOutputWithContext(ctx context.Context) BgpGracefulRestartOutput
}

func (*BgpGracefulRestart) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpGracefulRestart)(nil)).Elem()
}

func (i *BgpGracefulRestart) ToBgpGracefulRestartOutput() BgpGracefulRestartOutput {
	return i.ToBgpGracefulRestartOutputWithContext(context.Background())
}

func (i *BgpGracefulRestart) ToBgpGracefulRestartOutputWithContext(ctx context.Context) BgpGracefulRestartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpGracefulRestartOutput)
}

func (i *BgpGracefulRestart) ToOutput(ctx context.Context) pulumix.Output[*BgpGracefulRestart] {
	return pulumix.Output[*BgpGracefulRestart]{
		OutputState: i.ToBgpGracefulRestartOutputWithContext(ctx).OutputState,
	}
}

// BgpGracefulRestartArrayInput is an input type that accepts BgpGracefulRestartArray and BgpGracefulRestartArrayOutput values.
// You can construct a concrete instance of `BgpGracefulRestartArrayInput` via:
//
//	BgpGracefulRestartArray{ BgpGracefulRestartArgs{...} }
type BgpGracefulRestartArrayInput interface {
	pulumi.Input

	ToBgpGracefulRestartArrayOutput() BgpGracefulRestartArrayOutput
	ToBgpGracefulRestartArrayOutputWithContext(context.Context) BgpGracefulRestartArrayOutput
}

type BgpGracefulRestartArray []BgpGracefulRestartInput

func (BgpGracefulRestartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpGracefulRestart)(nil)).Elem()
}

func (i BgpGracefulRestartArray) ToBgpGracefulRestartArrayOutput() BgpGracefulRestartArrayOutput {
	return i.ToBgpGracefulRestartArrayOutputWithContext(context.Background())
}

func (i BgpGracefulRestartArray) ToBgpGracefulRestartArrayOutputWithContext(ctx context.Context) BgpGracefulRestartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpGracefulRestartArrayOutput)
}

func (i BgpGracefulRestartArray) ToOutput(ctx context.Context) pulumix.Output[[]*BgpGracefulRestart] {
	return pulumix.Output[[]*BgpGracefulRestart]{
		OutputState: i.ToBgpGracefulRestartArrayOutputWithContext(ctx).OutputState,
	}
}

// BgpGracefulRestartMapInput is an input type that accepts BgpGracefulRestartMap and BgpGracefulRestartMapOutput values.
// You can construct a concrete instance of `BgpGracefulRestartMapInput` via:
//
//	BgpGracefulRestartMap{ "key": BgpGracefulRestartArgs{...} }
type BgpGracefulRestartMapInput interface {
	pulumi.Input

	ToBgpGracefulRestartMapOutput() BgpGracefulRestartMapOutput
	ToBgpGracefulRestartMapOutputWithContext(context.Context) BgpGracefulRestartMapOutput
}

type BgpGracefulRestartMap map[string]BgpGracefulRestartInput

func (BgpGracefulRestartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpGracefulRestart)(nil)).Elem()
}

func (i BgpGracefulRestartMap) ToBgpGracefulRestartMapOutput() BgpGracefulRestartMapOutput {
	return i.ToBgpGracefulRestartMapOutputWithContext(context.Background())
}

func (i BgpGracefulRestartMap) ToBgpGracefulRestartMapOutputWithContext(ctx context.Context) BgpGracefulRestartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpGracefulRestartMapOutput)
}

func (i BgpGracefulRestartMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpGracefulRestart] {
	return pulumix.Output[map[string]*BgpGracefulRestart]{
		OutputState: i.ToBgpGracefulRestartMapOutputWithContext(ctx).OutputState,
	}
}

type BgpGracefulRestartOutput struct{ *pulumi.OutputState }

func (BgpGracefulRestartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpGracefulRestart)(nil)).Elem()
}

func (o BgpGracefulRestartOutput) ToBgpGracefulRestartOutput() BgpGracefulRestartOutput {
	return o
}

func (o BgpGracefulRestartOutput) ToBgpGracefulRestartOutputWithContext(ctx context.Context) BgpGracefulRestartOutput {
	return o
}

func (o BgpGracefulRestartOutput) ToOutput(ctx context.Context) pulumix.Output[*BgpGracefulRestart] {
	return pulumix.Output[*BgpGracefulRestart]{
		OutputState: o.OutputState,
	}
}

// Autonomous system number.
func (o BgpGracefulRestartOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpGracefulRestart) pulumi.StringOutput { return v.Asn }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o BgpGracefulRestartOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpGracefulRestart) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
func (o BgpGracefulRestartOutput) RestartInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *BgpGracefulRestart) pulumi.IntOutput { return v.RestartInterval }).(pulumi.IntOutput)
}

// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
func (o BgpGracefulRestartOutput) StaleInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *BgpGracefulRestart) pulumi.IntOutput { return v.StaleInterval }).(pulumi.IntOutput)
}

// VRF name.
func (o BgpGracefulRestartOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpGracefulRestart) pulumi.StringOutput { return v.Vrf }).(pulumi.StringOutput)
}

type BgpGracefulRestartArrayOutput struct{ *pulumi.OutputState }

func (BgpGracefulRestartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpGracefulRestart)(nil)).Elem()
}

func (o BgpGracefulRestartArrayOutput) ToBgpGracefulRestartArrayOutput() BgpGracefulRestartArrayOutput {
	return o
}

func (o BgpGracefulRestartArrayOutput) ToBgpGracefulRestartArrayOutputWithContext(ctx context.Context) BgpGracefulRestartArrayOutput {
	return o
}

func (o BgpGracefulRestartArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BgpGracefulRestart] {
	return pulumix.Output[[]*BgpGracefulRestart]{
		OutputState: o.OutputState,
	}
}

func (o BgpGracefulRestartArrayOutput) Index(i pulumi.IntInput) BgpGracefulRestartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpGracefulRestart {
		return vs[0].([]*BgpGracefulRestart)[vs[1].(int)]
	}).(BgpGracefulRestartOutput)
}

type BgpGracefulRestartMapOutput struct{ *pulumi.OutputState }

func (BgpGracefulRestartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpGracefulRestart)(nil)).Elem()
}

func (o BgpGracefulRestartMapOutput) ToBgpGracefulRestartMapOutput() BgpGracefulRestartMapOutput {
	return o
}

func (o BgpGracefulRestartMapOutput) ToBgpGracefulRestartMapOutputWithContext(ctx context.Context) BgpGracefulRestartMapOutput {
	return o
}

func (o BgpGracefulRestartMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpGracefulRestart] {
	return pulumix.Output[map[string]*BgpGracefulRestart]{
		OutputState: o.OutputState,
	}
}

func (o BgpGracefulRestartMapOutput) MapIndex(k pulumi.StringInput) BgpGracefulRestartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpGracefulRestart {
		return vs[0].(map[string]*BgpGracefulRestart)[vs[1].(string)]
	}).(BgpGracefulRestartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpGracefulRestartInput)(nil)).Elem(), &BgpGracefulRestart{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpGracefulRestartArrayInput)(nil)).Elem(), BgpGracefulRestartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpGracefulRestartMapInput)(nil)).Elem(), BgpGracefulRestartMap{})
	pulumi.RegisterOutputType(BgpGracefulRestartOutput{})
	pulumi.RegisterOutputType(BgpGracefulRestartArrayOutput{})
	pulumi.RegisterOutputType(BgpGracefulRestartMapOutput{})
}