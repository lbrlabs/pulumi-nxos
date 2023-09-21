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

type OspfArea struct {
	pulumi.CustomResourceState

	// Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
	AreaId pulumi.StringOutput `pulumi:"areaId"`
	// Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
	AuthenticationType pulumi.StringOutput `pulumi:"authenticationType"`
	// Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
	// value: `1`
	Cost pulumi.IntOutput `pulumi:"cost"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// OSPF instance name.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
	Type pulumi.StringOutput `pulumi:"type"`
	// VRF name.
	VrfName pulumi.StringOutput `pulumi:"vrfName"`
}

// NewOspfArea registers a new resource with the given unique name, arguments, and options.
func NewOspfArea(ctx *pulumi.Context,
	name string, args *OspfAreaArgs, opts ...pulumi.ResourceOption) (*OspfArea, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AreaId == nil {
		return nil, errors.New("invalid value for required argument 'AreaId'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.VrfName == nil {
		return nil, errors.New("invalid value for required argument 'VrfName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OspfArea
	err := ctx.RegisterResource("nxos:nxos/ospfArea:OspfArea", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOspfArea gets an existing OspfArea resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOspfArea(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OspfAreaState, opts ...pulumi.ResourceOption) (*OspfArea, error) {
	var resource OspfArea
	err := ctx.ReadResource("nxos:nxos/ospfArea:OspfArea", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OspfArea resources.
type ospfAreaState struct {
	// Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
	AreaId *string `pulumi:"areaId"`
	// Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
	AuthenticationType *string `pulumi:"authenticationType"`
	// Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
	// value: `1`
	Cost *int `pulumi:"cost"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// OSPF instance name.
	InstanceName *string `pulumi:"instanceName"`
	// Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
	Type *string `pulumi:"type"`
	// VRF name.
	VrfName *string `pulumi:"vrfName"`
}

type OspfAreaState struct {
	// Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
	AreaId pulumi.StringPtrInput
	// Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
	AuthenticationType pulumi.StringPtrInput
	// Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
	// value: `1`
	Cost pulumi.IntPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// OSPF instance name.
	InstanceName pulumi.StringPtrInput
	// Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
	Type pulumi.StringPtrInput
	// VRF name.
	VrfName pulumi.StringPtrInput
}

func (OspfAreaState) ElementType() reflect.Type {
	return reflect.TypeOf((*ospfAreaState)(nil)).Elem()
}

type ospfAreaArgs struct {
	// Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
	AreaId string `pulumi:"areaId"`
	// Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
	AuthenticationType *string `pulumi:"authenticationType"`
	// Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
	// value: `1`
	Cost *int `pulumi:"cost"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// OSPF instance name.
	InstanceName string `pulumi:"instanceName"`
	// Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
	Type *string `pulumi:"type"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

// The set of arguments for constructing a OspfArea resource.
type OspfAreaArgs struct {
	// Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
	AreaId pulumi.StringInput
	// Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
	AuthenticationType pulumi.StringPtrInput
	// Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
	// value: `1`
	Cost pulumi.IntPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// OSPF instance name.
	InstanceName pulumi.StringInput
	// Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
	Type pulumi.StringPtrInput
	// VRF name.
	VrfName pulumi.StringInput
}

func (OspfAreaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ospfAreaArgs)(nil)).Elem()
}

type OspfAreaInput interface {
	pulumi.Input

	ToOspfAreaOutput() OspfAreaOutput
	ToOspfAreaOutputWithContext(ctx context.Context) OspfAreaOutput
}

func (*OspfArea) ElementType() reflect.Type {
	return reflect.TypeOf((**OspfArea)(nil)).Elem()
}

func (i *OspfArea) ToOspfAreaOutput() OspfAreaOutput {
	return i.ToOspfAreaOutputWithContext(context.Background())
}

func (i *OspfArea) ToOspfAreaOutputWithContext(ctx context.Context) OspfAreaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OspfAreaOutput)
}

func (i *OspfArea) ToOutput(ctx context.Context) pulumix.Output[*OspfArea] {
	return pulumix.Output[*OspfArea]{
		OutputState: i.ToOspfAreaOutputWithContext(ctx).OutputState,
	}
}

// OspfAreaArrayInput is an input type that accepts OspfAreaArray and OspfAreaArrayOutput values.
// You can construct a concrete instance of `OspfAreaArrayInput` via:
//
//	OspfAreaArray{ OspfAreaArgs{...} }
type OspfAreaArrayInput interface {
	pulumi.Input

	ToOspfAreaArrayOutput() OspfAreaArrayOutput
	ToOspfAreaArrayOutputWithContext(context.Context) OspfAreaArrayOutput
}

type OspfAreaArray []OspfAreaInput

func (OspfAreaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OspfArea)(nil)).Elem()
}

func (i OspfAreaArray) ToOspfAreaArrayOutput() OspfAreaArrayOutput {
	return i.ToOspfAreaArrayOutputWithContext(context.Background())
}

func (i OspfAreaArray) ToOspfAreaArrayOutputWithContext(ctx context.Context) OspfAreaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OspfAreaArrayOutput)
}

func (i OspfAreaArray) ToOutput(ctx context.Context) pulumix.Output[[]*OspfArea] {
	return pulumix.Output[[]*OspfArea]{
		OutputState: i.ToOspfAreaArrayOutputWithContext(ctx).OutputState,
	}
}

// OspfAreaMapInput is an input type that accepts OspfAreaMap and OspfAreaMapOutput values.
// You can construct a concrete instance of `OspfAreaMapInput` via:
//
//	OspfAreaMap{ "key": OspfAreaArgs{...} }
type OspfAreaMapInput interface {
	pulumi.Input

	ToOspfAreaMapOutput() OspfAreaMapOutput
	ToOspfAreaMapOutputWithContext(context.Context) OspfAreaMapOutput
}

type OspfAreaMap map[string]OspfAreaInput

func (OspfAreaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OspfArea)(nil)).Elem()
}

func (i OspfAreaMap) ToOspfAreaMapOutput() OspfAreaMapOutput {
	return i.ToOspfAreaMapOutputWithContext(context.Background())
}

func (i OspfAreaMap) ToOspfAreaMapOutputWithContext(ctx context.Context) OspfAreaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OspfAreaMapOutput)
}

func (i OspfAreaMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OspfArea] {
	return pulumix.Output[map[string]*OspfArea]{
		OutputState: i.ToOspfAreaMapOutputWithContext(ctx).OutputState,
	}
}

type OspfAreaOutput struct{ *pulumi.OutputState }

func (OspfAreaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OspfArea)(nil)).Elem()
}

func (o OspfAreaOutput) ToOspfAreaOutput() OspfAreaOutput {
	return o
}

func (o OspfAreaOutput) ToOspfAreaOutputWithContext(ctx context.Context) OspfAreaOutput {
	return o
}

func (o OspfAreaOutput) ToOutput(ctx context.Context) pulumix.Output[*OspfArea] {
	return pulumix.Output[*OspfArea]{
		OutputState: o.OutputState,
	}
}

// Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
func (o OspfAreaOutput) AreaId() pulumi.StringOutput {
	return o.ApplyT(func(v *OspfArea) pulumi.StringOutput { return v.AreaId }).(pulumi.StringOutput)
}

// Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
func (o OspfAreaOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v *OspfArea) pulumi.StringOutput { return v.AuthenticationType }).(pulumi.StringOutput)
}

// Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
// value: `1`
func (o OspfAreaOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v *OspfArea) pulumi.IntOutput { return v.Cost }).(pulumi.IntOutput)
}

// A device name from the provider configuration.
func (o OspfAreaOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OspfArea) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// OSPF instance name.
func (o OspfAreaOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *OspfArea) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
func (o OspfAreaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OspfArea) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VRF name.
func (o OspfAreaOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v *OspfArea) pulumi.StringOutput { return v.VrfName }).(pulumi.StringOutput)
}

type OspfAreaArrayOutput struct{ *pulumi.OutputState }

func (OspfAreaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OspfArea)(nil)).Elem()
}

func (o OspfAreaArrayOutput) ToOspfAreaArrayOutput() OspfAreaArrayOutput {
	return o
}

func (o OspfAreaArrayOutput) ToOspfAreaArrayOutputWithContext(ctx context.Context) OspfAreaArrayOutput {
	return o
}

func (o OspfAreaArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OspfArea] {
	return pulumix.Output[[]*OspfArea]{
		OutputState: o.OutputState,
	}
}

func (o OspfAreaArrayOutput) Index(i pulumi.IntInput) OspfAreaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OspfArea {
		return vs[0].([]*OspfArea)[vs[1].(int)]
	}).(OspfAreaOutput)
}

type OspfAreaMapOutput struct{ *pulumi.OutputState }

func (OspfAreaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OspfArea)(nil)).Elem()
}

func (o OspfAreaMapOutput) ToOspfAreaMapOutput() OspfAreaMapOutput {
	return o
}

func (o OspfAreaMapOutput) ToOspfAreaMapOutputWithContext(ctx context.Context) OspfAreaMapOutput {
	return o
}

func (o OspfAreaMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OspfArea] {
	return pulumix.Output[map[string]*OspfArea]{
		OutputState: o.OutputState,
	}
}

func (o OspfAreaMapOutput) MapIndex(k pulumi.StringInput) OspfAreaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OspfArea {
		return vs[0].(map[string]*OspfArea)[vs[1].(string)]
	}).(OspfAreaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OspfAreaInput)(nil)).Elem(), &OspfArea{})
	pulumi.RegisterInputType(reflect.TypeOf((*OspfAreaArrayInput)(nil)).Elem(), OspfAreaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OspfAreaMapInput)(nil)).Elem(), OspfAreaMap{})
	pulumi.RegisterOutputType(OspfAreaOutput{})
	pulumi.RegisterOutputType(OspfAreaArrayOutput{})
	pulumi.RegisterOutputType(OspfAreaMapOutput{})
}
