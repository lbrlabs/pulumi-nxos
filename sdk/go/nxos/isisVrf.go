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

// This resource can manage the IS-IS VRF configuration.
//
// - API Documentation: [isisDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Dom/)
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
//			_, err := nxos.NewIsisVrf(ctx, "example", &nxos.IsisVrfArgs{
//				AdminState:            pulumi.String("enabled"),
//				AuthenticationCheckL1: pulumi.Bool(false),
//				AuthenticationCheckL2: pulumi.Bool(false),
//				AuthenticationKeyL1:   pulumi.String(""),
//				AuthenticationKeyL2:   pulumi.String(""),
//				AuthenticationTypeL1:  pulumi.String("unknown"),
//				AuthenticationTypeL2:  pulumi.String("unknown"),
//				BandwidthReference:    pulumi.Int(400000),
//				BanwidthReferenceUnit: pulumi.String("mbps"),
//				InstanceName:          pulumi.String("ISIS1"),
//				IsType:                pulumi.String("l2"),
//				MetricType:            pulumi.String("wide"),
//				Mtu:                   pulumi.Int(2000),
//				Net:                   pulumi.String("49.0001.0000.0000.3333.00"),
//				PassiveDefault:        pulumi.String("l12"),
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
//	$ pulumi import nxos:index/isisVrf:IsisVrf example "sys/isis/inst-[ISIS1]/dom-[default]"
//
// ```
type IsisVrf struct {
	pulumi.CustomResourceState

	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringOutput `pulumi:"adminState"`
	// Authentication Check for ISIS on Level-1. - Default value: `true`
	AuthenticationCheckL1 pulumi.BoolOutput `pulumi:"authenticationCheckL1"`
	// Authentication Check for ISIS on Level-2. - Default value: `true`
	AuthenticationCheckL2 pulumi.BoolOutput `pulumi:"authenticationCheckL2"`
	// Authentication Key for IS-IS on Level-1.
	AuthenticationKeyL1 pulumi.StringPtrOutput `pulumi:"authenticationKeyL1"`
	// Authentication Key for IS-IS on Level-2.
	AuthenticationKeyL2 pulumi.StringPtrOutput `pulumi:"authenticationKeyL2"`
	// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL1 pulumi.StringOutput `pulumi:"authenticationTypeL1"`
	// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL2 pulumi.StringOutput `pulumi:"authenticationTypeL2"`
	// The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
	// metric. - Range: `0`-`4294967295` - Default value: `40000`
	BandwidthReference pulumi.IntOutput `pulumi:"bandwidthReference"`
	// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
	BanwidthReferenceUnit pulumi.StringOutput `pulumi:"banwidthReferenceUnit"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// IS-IS instance name.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
	IsType pulumi.StringOutput `pulumi:"isType"`
	// IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
	MetricType pulumi.StringOutput `pulumi:"metricType"`
	// The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
	// bytes. - Range: `256`-`4352` - Default value: `1492`
	Mtu pulumi.IntOutput `pulumi:"mtu"`
	// VRF name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Holds IS-IS domain NET (address) value.
	Net pulumi.StringPtrOutput `pulumi:"net"`
	// IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
	PassiveDefault pulumi.StringOutput `pulumi:"passiveDefault"`
}

// NewIsisVrf registers a new resource with the given unique name, arguments, and options.
func NewIsisVrf(ctx *pulumi.Context,
	name string, args *IsisVrfArgs, opts ...pulumi.ResourceOption) (*IsisVrf, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IsisVrf
	err := ctx.RegisterResource("nxos:index/isisVrf:IsisVrf", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIsisVrf gets an existing IsisVrf resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIsisVrf(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IsisVrfState, opts ...pulumi.ResourceOption) (*IsisVrf, error) {
	var resource IsisVrf
	err := ctx.ReadResource("nxos:index/isisVrf:IsisVrf", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IsisVrf resources.
type isisVrfState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState *string `pulumi:"adminState"`
	// Authentication Check for ISIS on Level-1. - Default value: `true`
	AuthenticationCheckL1 *bool `pulumi:"authenticationCheckL1"`
	// Authentication Check for ISIS on Level-2. - Default value: `true`
	AuthenticationCheckL2 *bool `pulumi:"authenticationCheckL2"`
	// Authentication Key for IS-IS on Level-1.
	AuthenticationKeyL1 *string `pulumi:"authenticationKeyL1"`
	// Authentication Key for IS-IS on Level-2.
	AuthenticationKeyL2 *string `pulumi:"authenticationKeyL2"`
	// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL1 *string `pulumi:"authenticationTypeL1"`
	// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL2 *string `pulumi:"authenticationTypeL2"`
	// The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
	// metric. - Range: `0`-`4294967295` - Default value: `40000`
	BandwidthReference *int `pulumi:"bandwidthReference"`
	// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
	BanwidthReferenceUnit *string `pulumi:"banwidthReferenceUnit"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// IS-IS instance name.
	InstanceName *string `pulumi:"instanceName"`
	// IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
	IsType *string `pulumi:"isType"`
	// IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
	MetricType *string `pulumi:"metricType"`
	// The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
	// bytes. - Range: `256`-`4352` - Default value: `1492`
	Mtu *int `pulumi:"mtu"`
	// VRF name.
	Name *string `pulumi:"name"`
	// Holds IS-IS domain NET (address) value.
	Net *string `pulumi:"net"`
	// IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
	PassiveDefault *string `pulumi:"passiveDefault"`
}

type IsisVrfState struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringPtrInput
	// Authentication Check for ISIS on Level-1. - Default value: `true`
	AuthenticationCheckL1 pulumi.BoolPtrInput
	// Authentication Check for ISIS on Level-2. - Default value: `true`
	AuthenticationCheckL2 pulumi.BoolPtrInput
	// Authentication Key for IS-IS on Level-1.
	AuthenticationKeyL1 pulumi.StringPtrInput
	// Authentication Key for IS-IS on Level-2.
	AuthenticationKeyL2 pulumi.StringPtrInput
	// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL1 pulumi.StringPtrInput
	// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL2 pulumi.StringPtrInput
	// The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
	// metric. - Range: `0`-`4294967295` - Default value: `40000`
	BandwidthReference pulumi.IntPtrInput
	// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
	BanwidthReferenceUnit pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// IS-IS instance name.
	InstanceName pulumi.StringPtrInput
	// IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
	IsType pulumi.StringPtrInput
	// IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
	MetricType pulumi.StringPtrInput
	// The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
	// bytes. - Range: `256`-`4352` - Default value: `1492`
	Mtu pulumi.IntPtrInput
	// VRF name.
	Name pulumi.StringPtrInput
	// Holds IS-IS domain NET (address) value.
	Net pulumi.StringPtrInput
	// IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
	PassiveDefault pulumi.StringPtrInput
}

func (IsisVrfState) ElementType() reflect.Type {
	return reflect.TypeOf((*isisVrfState)(nil)).Elem()
}

type isisVrfArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState *string `pulumi:"adminState"`
	// Authentication Check for ISIS on Level-1. - Default value: `true`
	AuthenticationCheckL1 *bool `pulumi:"authenticationCheckL1"`
	// Authentication Check for ISIS on Level-2. - Default value: `true`
	AuthenticationCheckL2 *bool `pulumi:"authenticationCheckL2"`
	// Authentication Key for IS-IS on Level-1.
	AuthenticationKeyL1 *string `pulumi:"authenticationKeyL1"`
	// Authentication Key for IS-IS on Level-2.
	AuthenticationKeyL2 *string `pulumi:"authenticationKeyL2"`
	// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL1 *string `pulumi:"authenticationTypeL1"`
	// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL2 *string `pulumi:"authenticationTypeL2"`
	// The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
	// metric. - Range: `0`-`4294967295` - Default value: `40000`
	BandwidthReference *int `pulumi:"bandwidthReference"`
	// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
	BanwidthReferenceUnit *string `pulumi:"banwidthReferenceUnit"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// IS-IS instance name.
	InstanceName string `pulumi:"instanceName"`
	// IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
	IsType *string `pulumi:"isType"`
	// IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
	MetricType *string `pulumi:"metricType"`
	// The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
	// bytes. - Range: `256`-`4352` - Default value: `1492`
	Mtu *int `pulumi:"mtu"`
	// VRF name.
	Name *string `pulumi:"name"`
	// Holds IS-IS domain NET (address) value.
	Net *string `pulumi:"net"`
	// IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
	PassiveDefault *string `pulumi:"passiveDefault"`
}

// The set of arguments for constructing a IsisVrf resource.
type IsisVrfArgs struct {
	// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
	AdminState pulumi.StringPtrInput
	// Authentication Check for ISIS on Level-1. - Default value: `true`
	AuthenticationCheckL1 pulumi.BoolPtrInput
	// Authentication Check for ISIS on Level-2. - Default value: `true`
	AuthenticationCheckL2 pulumi.BoolPtrInput
	// Authentication Key for IS-IS on Level-1.
	AuthenticationKeyL1 pulumi.StringPtrInput
	// Authentication Key for IS-IS on Level-2.
	AuthenticationKeyL2 pulumi.StringPtrInput
	// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL1 pulumi.StringPtrInput
	// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
	AuthenticationTypeL2 pulumi.StringPtrInput
	// The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
	// metric. - Range: `0`-`4294967295` - Default value: `40000`
	BandwidthReference pulumi.IntPtrInput
	// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
	BanwidthReferenceUnit pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// IS-IS instance name.
	InstanceName pulumi.StringInput
	// IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
	IsType pulumi.StringPtrInput
	// IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
	MetricType pulumi.StringPtrInput
	// The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
	// bytes. - Range: `256`-`4352` - Default value: `1492`
	Mtu pulumi.IntPtrInput
	// VRF name.
	Name pulumi.StringPtrInput
	// Holds IS-IS domain NET (address) value.
	Net pulumi.StringPtrInput
	// IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
	PassiveDefault pulumi.StringPtrInput
}

func (IsisVrfArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*isisVrfArgs)(nil)).Elem()
}

type IsisVrfInput interface {
	pulumi.Input

	ToIsisVrfOutput() IsisVrfOutput
	ToIsisVrfOutputWithContext(ctx context.Context) IsisVrfOutput
}

func (*IsisVrf) ElementType() reflect.Type {
	return reflect.TypeOf((**IsisVrf)(nil)).Elem()
}

func (i *IsisVrf) ToIsisVrfOutput() IsisVrfOutput {
	return i.ToIsisVrfOutputWithContext(context.Background())
}

func (i *IsisVrf) ToIsisVrfOutputWithContext(ctx context.Context) IsisVrfOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsisVrfOutput)
}

func (i *IsisVrf) ToOutput(ctx context.Context) pulumix.Output[*IsisVrf] {
	return pulumix.Output[*IsisVrf]{
		OutputState: i.ToIsisVrfOutputWithContext(ctx).OutputState,
	}
}

// IsisVrfArrayInput is an input type that accepts IsisVrfArray and IsisVrfArrayOutput values.
// You can construct a concrete instance of `IsisVrfArrayInput` via:
//
//	IsisVrfArray{ IsisVrfArgs{...} }
type IsisVrfArrayInput interface {
	pulumi.Input

	ToIsisVrfArrayOutput() IsisVrfArrayOutput
	ToIsisVrfArrayOutputWithContext(context.Context) IsisVrfArrayOutput
}

type IsisVrfArray []IsisVrfInput

func (IsisVrfArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IsisVrf)(nil)).Elem()
}

func (i IsisVrfArray) ToIsisVrfArrayOutput() IsisVrfArrayOutput {
	return i.ToIsisVrfArrayOutputWithContext(context.Background())
}

func (i IsisVrfArray) ToIsisVrfArrayOutputWithContext(ctx context.Context) IsisVrfArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsisVrfArrayOutput)
}

func (i IsisVrfArray) ToOutput(ctx context.Context) pulumix.Output[[]*IsisVrf] {
	return pulumix.Output[[]*IsisVrf]{
		OutputState: i.ToIsisVrfArrayOutputWithContext(ctx).OutputState,
	}
}

// IsisVrfMapInput is an input type that accepts IsisVrfMap and IsisVrfMapOutput values.
// You can construct a concrete instance of `IsisVrfMapInput` via:
//
//	IsisVrfMap{ "key": IsisVrfArgs{...} }
type IsisVrfMapInput interface {
	pulumi.Input

	ToIsisVrfMapOutput() IsisVrfMapOutput
	ToIsisVrfMapOutputWithContext(context.Context) IsisVrfMapOutput
}

type IsisVrfMap map[string]IsisVrfInput

func (IsisVrfMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IsisVrf)(nil)).Elem()
}

func (i IsisVrfMap) ToIsisVrfMapOutput() IsisVrfMapOutput {
	return i.ToIsisVrfMapOutputWithContext(context.Background())
}

func (i IsisVrfMap) ToIsisVrfMapOutputWithContext(ctx context.Context) IsisVrfMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IsisVrfMapOutput)
}

func (i IsisVrfMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IsisVrf] {
	return pulumix.Output[map[string]*IsisVrf]{
		OutputState: i.ToIsisVrfMapOutputWithContext(ctx).OutputState,
	}
}

type IsisVrfOutput struct{ *pulumi.OutputState }

func (IsisVrfOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IsisVrf)(nil)).Elem()
}

func (o IsisVrfOutput) ToIsisVrfOutput() IsisVrfOutput {
	return o
}

func (o IsisVrfOutput) ToIsisVrfOutputWithContext(ctx context.Context) IsisVrfOutput {
	return o
}

func (o IsisVrfOutput) ToOutput(ctx context.Context) pulumix.Output[*IsisVrf] {
	return pulumix.Output[*IsisVrf]{
		OutputState: o.OutputState,
	}
}

// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
func (o IsisVrfOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.AdminState }).(pulumi.StringOutput)
}

// Authentication Check for ISIS on Level-1. - Default value: `true`
func (o IsisVrfOutput) AuthenticationCheckL1() pulumi.BoolOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.BoolOutput { return v.AuthenticationCheckL1 }).(pulumi.BoolOutput)
}

// Authentication Check for ISIS on Level-2. - Default value: `true`
func (o IsisVrfOutput) AuthenticationCheckL2() pulumi.BoolOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.BoolOutput { return v.AuthenticationCheckL2 }).(pulumi.BoolOutput)
}

// Authentication Key for IS-IS on Level-1.
func (o IsisVrfOutput) AuthenticationKeyL1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringPtrOutput { return v.AuthenticationKeyL1 }).(pulumi.StringPtrOutput)
}

// Authentication Key for IS-IS on Level-2.
func (o IsisVrfOutput) AuthenticationKeyL2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringPtrOutput { return v.AuthenticationKeyL2 }).(pulumi.StringPtrOutput)
}

// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
func (o IsisVrfOutput) AuthenticationTypeL1() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.AuthenticationTypeL1 }).(pulumi.StringOutput)
}

// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
func (o IsisVrfOutput) AuthenticationTypeL2() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.AuthenticationTypeL2 }).(pulumi.StringOutput)
}

// The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
// metric. - Range: `0`-`4294967295` - Default value: `40000`
func (o IsisVrfOutput) BandwidthReference() pulumi.IntOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.IntOutput { return v.BandwidthReference }).(pulumi.IntOutput)
}

// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
func (o IsisVrfOutput) BanwidthReferenceUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.BanwidthReferenceUnit }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o IsisVrfOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// IS-IS instance name.
func (o IsisVrfOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
func (o IsisVrfOutput) IsType() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.IsType }).(pulumi.StringOutput)
}

// IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
func (o IsisVrfOutput) MetricType() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.MetricType }).(pulumi.StringOutput)
}

// The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
// bytes. - Range: `256`-`4352` - Default value: `1492`
func (o IsisVrfOutput) Mtu() pulumi.IntOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.IntOutput { return v.Mtu }).(pulumi.IntOutput)
}

// VRF name.
func (o IsisVrfOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Holds IS-IS domain NET (address) value.
func (o IsisVrfOutput) Net() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringPtrOutput { return v.Net }).(pulumi.StringPtrOutput)
}

// IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
func (o IsisVrfOutput) PassiveDefault() pulumi.StringOutput {
	return o.ApplyT(func(v *IsisVrf) pulumi.StringOutput { return v.PassiveDefault }).(pulumi.StringOutput)
}

type IsisVrfArrayOutput struct{ *pulumi.OutputState }

func (IsisVrfArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IsisVrf)(nil)).Elem()
}

func (o IsisVrfArrayOutput) ToIsisVrfArrayOutput() IsisVrfArrayOutput {
	return o
}

func (o IsisVrfArrayOutput) ToIsisVrfArrayOutputWithContext(ctx context.Context) IsisVrfArrayOutput {
	return o
}

func (o IsisVrfArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IsisVrf] {
	return pulumix.Output[[]*IsisVrf]{
		OutputState: o.OutputState,
	}
}

func (o IsisVrfArrayOutput) Index(i pulumi.IntInput) IsisVrfOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IsisVrf {
		return vs[0].([]*IsisVrf)[vs[1].(int)]
	}).(IsisVrfOutput)
}

type IsisVrfMapOutput struct{ *pulumi.OutputState }

func (IsisVrfMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IsisVrf)(nil)).Elem()
}

func (o IsisVrfMapOutput) ToIsisVrfMapOutput() IsisVrfMapOutput {
	return o
}

func (o IsisVrfMapOutput) ToIsisVrfMapOutputWithContext(ctx context.Context) IsisVrfMapOutput {
	return o
}

func (o IsisVrfMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IsisVrf] {
	return pulumix.Output[map[string]*IsisVrf]{
		OutputState: o.OutputState,
	}
}

func (o IsisVrfMapOutput) MapIndex(k pulumi.StringInput) IsisVrfOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IsisVrf {
		return vs[0].(map[string]*IsisVrf)[vs[1].(string)]
	}).(IsisVrfOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IsisVrfInput)(nil)).Elem(), &IsisVrf{})
	pulumi.RegisterInputType(reflect.TypeOf((*IsisVrfArrayInput)(nil)).Elem(), IsisVrfArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IsisVrfMapInput)(nil)).Elem(), IsisVrfMap{})
	pulumi.RegisterOutputType(IsisVrfOutput{})
	pulumi.RegisterOutputType(IsisVrfArrayOutput{})
	pulumi.RegisterOutputType(IsisVrfMapOutput{})
}
