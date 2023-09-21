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

type BgpPeerTemplate struct {
	pulumi.CustomResourceState

	// Autonomous system number.
	Asn pulumi.StringOutput `pulumi:"asn"`
	// Peer template description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrOutput `pulumi:"device"`
	// Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
	// `fabric-internal`
	PeerType pulumi.StringOutput `pulumi:"peerType"`
	// Peer template autonomous system number.
	RemoteAsn pulumi.StringPtrOutput `pulumi:"remoteAsn"`
	// Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
	SourceInterface pulumi.StringOutput `pulumi:"sourceInterface"`
	// Peer template name.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewBgpPeerTemplate registers a new resource with the given unique name, arguments, and options.
func NewBgpPeerTemplate(ctx *pulumi.Context,
	name string, args *BgpPeerTemplateArgs, opts ...pulumi.ResourceOption) (*BgpPeerTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Asn == nil {
		return nil, errors.New("invalid value for required argument 'Asn'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BgpPeerTemplate
	err := ctx.RegisterResource("nxos:nxos/bgpPeerTemplate:BgpPeerTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpPeerTemplate gets an existing BgpPeerTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpPeerTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpPeerTemplateState, opts ...pulumi.ResourceOption) (*BgpPeerTemplate, error) {
	var resource BgpPeerTemplate
	err := ctx.ReadResource("nxos:nxos/bgpPeerTemplate:BgpPeerTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpPeerTemplate resources.
type bgpPeerTemplateState struct {
	// Autonomous system number.
	Asn *string `pulumi:"asn"`
	// Peer template description.
	Description *string `pulumi:"description"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
	// `fabric-internal`
	PeerType *string `pulumi:"peerType"`
	// Peer template autonomous system number.
	RemoteAsn *string `pulumi:"remoteAsn"`
	// Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
	SourceInterface *string `pulumi:"sourceInterface"`
	// Peer template name.
	TemplateName *string `pulumi:"templateName"`
}

type BgpPeerTemplateState struct {
	// Autonomous system number.
	Asn pulumi.StringPtrInput
	// Peer template description.
	Description pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
	// `fabric-internal`
	PeerType pulumi.StringPtrInput
	// Peer template autonomous system number.
	RemoteAsn pulumi.StringPtrInput
	// Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
	SourceInterface pulumi.StringPtrInput
	// Peer template name.
	TemplateName pulumi.StringPtrInput
}

func (BgpPeerTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerTemplateState)(nil)).Elem()
}

type bgpPeerTemplateArgs struct {
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// Peer template description.
	Description *string `pulumi:"description"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
	// `fabric-internal`
	PeerType *string `pulumi:"peerType"`
	// Peer template autonomous system number.
	RemoteAsn *string `pulumi:"remoteAsn"`
	// Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
	SourceInterface *string `pulumi:"sourceInterface"`
	// Peer template name.
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a BgpPeerTemplate resource.
type BgpPeerTemplateArgs struct {
	// Autonomous system number.
	Asn pulumi.StringInput
	// Peer template description.
	Description pulumi.StringPtrInput
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput
	// Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
	// `fabric-internal`
	PeerType pulumi.StringPtrInput
	// Peer template autonomous system number.
	RemoteAsn pulumi.StringPtrInput
	// Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
	SourceInterface pulumi.StringPtrInput
	// Peer template name.
	TemplateName pulumi.StringInput
}

func (BgpPeerTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerTemplateArgs)(nil)).Elem()
}

type BgpPeerTemplateInput interface {
	pulumi.Input

	ToBgpPeerTemplateOutput() BgpPeerTemplateOutput
	ToBgpPeerTemplateOutputWithContext(ctx context.Context) BgpPeerTemplateOutput
}

func (*BgpPeerTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpPeerTemplate)(nil)).Elem()
}

func (i *BgpPeerTemplate) ToBgpPeerTemplateOutput() BgpPeerTemplateOutput {
	return i.ToBgpPeerTemplateOutputWithContext(context.Background())
}

func (i *BgpPeerTemplate) ToBgpPeerTemplateOutputWithContext(ctx context.Context) BgpPeerTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerTemplateOutput)
}

func (i *BgpPeerTemplate) ToOutput(ctx context.Context) pulumix.Output[*BgpPeerTemplate] {
	return pulumix.Output[*BgpPeerTemplate]{
		OutputState: i.ToBgpPeerTemplateOutputWithContext(ctx).OutputState,
	}
}

// BgpPeerTemplateArrayInput is an input type that accepts BgpPeerTemplateArray and BgpPeerTemplateArrayOutput values.
// You can construct a concrete instance of `BgpPeerTemplateArrayInput` via:
//
//	BgpPeerTemplateArray{ BgpPeerTemplateArgs{...} }
type BgpPeerTemplateArrayInput interface {
	pulumi.Input

	ToBgpPeerTemplateArrayOutput() BgpPeerTemplateArrayOutput
	ToBgpPeerTemplateArrayOutputWithContext(context.Context) BgpPeerTemplateArrayOutput
}

type BgpPeerTemplateArray []BgpPeerTemplateInput

func (BgpPeerTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpPeerTemplate)(nil)).Elem()
}

func (i BgpPeerTemplateArray) ToBgpPeerTemplateArrayOutput() BgpPeerTemplateArrayOutput {
	return i.ToBgpPeerTemplateArrayOutputWithContext(context.Background())
}

func (i BgpPeerTemplateArray) ToBgpPeerTemplateArrayOutputWithContext(ctx context.Context) BgpPeerTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerTemplateArrayOutput)
}

func (i BgpPeerTemplateArray) ToOutput(ctx context.Context) pulumix.Output[[]*BgpPeerTemplate] {
	return pulumix.Output[[]*BgpPeerTemplate]{
		OutputState: i.ToBgpPeerTemplateArrayOutputWithContext(ctx).OutputState,
	}
}

// BgpPeerTemplateMapInput is an input type that accepts BgpPeerTemplateMap and BgpPeerTemplateMapOutput values.
// You can construct a concrete instance of `BgpPeerTemplateMapInput` via:
//
//	BgpPeerTemplateMap{ "key": BgpPeerTemplateArgs{...} }
type BgpPeerTemplateMapInput interface {
	pulumi.Input

	ToBgpPeerTemplateMapOutput() BgpPeerTemplateMapOutput
	ToBgpPeerTemplateMapOutputWithContext(context.Context) BgpPeerTemplateMapOutput
}

type BgpPeerTemplateMap map[string]BgpPeerTemplateInput

func (BgpPeerTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpPeerTemplate)(nil)).Elem()
}

func (i BgpPeerTemplateMap) ToBgpPeerTemplateMapOutput() BgpPeerTemplateMapOutput {
	return i.ToBgpPeerTemplateMapOutputWithContext(context.Background())
}

func (i BgpPeerTemplateMap) ToBgpPeerTemplateMapOutputWithContext(ctx context.Context) BgpPeerTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpPeerTemplateMapOutput)
}

func (i BgpPeerTemplateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpPeerTemplate] {
	return pulumix.Output[map[string]*BgpPeerTemplate]{
		OutputState: i.ToBgpPeerTemplateMapOutputWithContext(ctx).OutputState,
	}
}

type BgpPeerTemplateOutput struct{ *pulumi.OutputState }

func (BgpPeerTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpPeerTemplate)(nil)).Elem()
}

func (o BgpPeerTemplateOutput) ToBgpPeerTemplateOutput() BgpPeerTemplateOutput {
	return o
}

func (o BgpPeerTemplateOutput) ToBgpPeerTemplateOutputWithContext(ctx context.Context) BgpPeerTemplateOutput {
	return o
}

func (o BgpPeerTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*BgpPeerTemplate] {
	return pulumix.Output[*BgpPeerTemplate]{
		OutputState: o.OutputState,
	}
}

// Autonomous system number.
func (o BgpPeerTemplateOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplate) pulumi.StringOutput { return v.Asn }).(pulumi.StringOutput)
}

// Peer template description.
func (o BgpPeerTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpPeerTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A device name from the provider configuration.
func (o BgpPeerTemplateOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpPeerTemplate) pulumi.StringPtrOutput { return v.Device }).(pulumi.StringPtrOutput)
}

// Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
// `fabric-internal`
func (o BgpPeerTemplateOutput) PeerType() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplate) pulumi.StringOutput { return v.PeerType }).(pulumi.StringOutput)
}

// Peer template autonomous system number.
func (o BgpPeerTemplateOutput) RemoteAsn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpPeerTemplate) pulumi.StringPtrOutput { return v.RemoteAsn }).(pulumi.StringPtrOutput)
}

// Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
func (o BgpPeerTemplateOutput) SourceInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplate) pulumi.StringOutput { return v.SourceInterface }).(pulumi.StringOutput)
}

// Peer template name.
func (o BgpPeerTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *BgpPeerTemplate) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

type BgpPeerTemplateArrayOutput struct{ *pulumi.OutputState }

func (BgpPeerTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BgpPeerTemplate)(nil)).Elem()
}

func (o BgpPeerTemplateArrayOutput) ToBgpPeerTemplateArrayOutput() BgpPeerTemplateArrayOutput {
	return o
}

func (o BgpPeerTemplateArrayOutput) ToBgpPeerTemplateArrayOutputWithContext(ctx context.Context) BgpPeerTemplateArrayOutput {
	return o
}

func (o BgpPeerTemplateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BgpPeerTemplate] {
	return pulumix.Output[[]*BgpPeerTemplate]{
		OutputState: o.OutputState,
	}
}

func (o BgpPeerTemplateArrayOutput) Index(i pulumi.IntInput) BgpPeerTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BgpPeerTemplate {
		return vs[0].([]*BgpPeerTemplate)[vs[1].(int)]
	}).(BgpPeerTemplateOutput)
}

type BgpPeerTemplateMapOutput struct{ *pulumi.OutputState }

func (BgpPeerTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BgpPeerTemplate)(nil)).Elem()
}

func (o BgpPeerTemplateMapOutput) ToBgpPeerTemplateMapOutput() BgpPeerTemplateMapOutput {
	return o
}

func (o BgpPeerTemplateMapOutput) ToBgpPeerTemplateMapOutputWithContext(ctx context.Context) BgpPeerTemplateMapOutput {
	return o
}

func (o BgpPeerTemplateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BgpPeerTemplate] {
	return pulumix.Output[map[string]*BgpPeerTemplate]{
		OutputState: o.OutputState,
	}
}

func (o BgpPeerTemplateMapOutput) MapIndex(k pulumi.StringInput) BgpPeerTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BgpPeerTemplate {
		return vs[0].(map[string]*BgpPeerTemplate)[vs[1].(string)]
	}).(BgpPeerTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerTemplateInput)(nil)).Elem(), &BgpPeerTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerTemplateArrayInput)(nil)).Elem(), BgpPeerTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BgpPeerTemplateMapInput)(nil)).Elem(), BgpPeerTemplateMap{})
	pulumi.RegisterOutputType(BgpPeerTemplateOutput{})
	pulumi.RegisterOutputType(BgpPeerTemplateArrayOutput{})
	pulumi.RegisterOutputType(BgpPeerTemplateMapOutput{})
}
