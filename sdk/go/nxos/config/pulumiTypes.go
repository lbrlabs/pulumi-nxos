// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type Devices struct {
	Name string `pulumi:"name"`
	Url  string `pulumi:"url"`
}

// DevicesInput is an input type that accepts DevicesArgs and DevicesOutput values.
// You can construct a concrete instance of `DevicesInput` via:
//
//	DevicesArgs{...}
type DevicesInput interface {
	pulumi.Input

	ToDevicesOutput() DevicesOutput
	ToDevicesOutputWithContext(context.Context) DevicesOutput
}

type DevicesArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Url  pulumi.StringInput `pulumi:"url"`
}

func (DevicesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Devices)(nil)).Elem()
}

func (i DevicesArgs) ToDevicesOutput() DevicesOutput {
	return i.ToDevicesOutputWithContext(context.Background())
}

func (i DevicesArgs) ToDevicesOutputWithContext(ctx context.Context) DevicesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicesOutput)
}

func (i DevicesArgs) ToOutput(ctx context.Context) pulumix.Output[Devices] {
	return pulumix.Output[Devices]{
		OutputState: i.ToDevicesOutputWithContext(ctx).OutputState,
	}
}

// DevicesArrayInput is an input type that accepts DevicesArray and DevicesArrayOutput values.
// You can construct a concrete instance of `DevicesArrayInput` via:
//
//	DevicesArray{ DevicesArgs{...} }
type DevicesArrayInput interface {
	pulumi.Input

	ToDevicesArrayOutput() DevicesArrayOutput
	ToDevicesArrayOutputWithContext(context.Context) DevicesArrayOutput
}

type DevicesArray []DevicesInput

func (DevicesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Devices)(nil)).Elem()
}

func (i DevicesArray) ToDevicesArrayOutput() DevicesArrayOutput {
	return i.ToDevicesArrayOutputWithContext(context.Background())
}

func (i DevicesArray) ToDevicesArrayOutputWithContext(ctx context.Context) DevicesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DevicesArrayOutput)
}

func (i DevicesArray) ToOutput(ctx context.Context) pulumix.Output[[]Devices] {
	return pulumix.Output[[]Devices]{
		OutputState: i.ToDevicesArrayOutputWithContext(ctx).OutputState,
	}
}

type DevicesOutput struct{ *pulumi.OutputState }

func (DevicesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Devices)(nil)).Elem()
}

func (o DevicesOutput) ToDevicesOutput() DevicesOutput {
	return o
}

func (o DevicesOutput) ToDevicesOutputWithContext(ctx context.Context) DevicesOutput {
	return o
}

func (o DevicesOutput) ToOutput(ctx context.Context) pulumix.Output[Devices] {
	return pulumix.Output[Devices]{
		OutputState: o.OutputState,
	}
}

func (o DevicesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Devices) string { return v.Name }).(pulumi.StringOutput)
}

func (o DevicesOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v Devices) string { return v.Url }).(pulumi.StringOutput)
}

type DevicesArrayOutput struct{ *pulumi.OutputState }

func (DevicesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Devices)(nil)).Elem()
}

func (o DevicesArrayOutput) ToDevicesArrayOutput() DevicesArrayOutput {
	return o
}

func (o DevicesArrayOutput) ToDevicesArrayOutputWithContext(ctx context.Context) DevicesArrayOutput {
	return o
}

func (o DevicesArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]Devices] {
	return pulumix.Output[[]Devices]{
		OutputState: o.OutputState,
	}
}

func (o DevicesArrayOutput) Index(i pulumi.IntInput) DevicesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Devices {
		return vs[0].([]Devices)[vs[1].(int)]
	}).(DevicesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DevicesInput)(nil)).Elem(), DevicesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DevicesArrayInput)(nil)).Elem(), DevicesArray{})
	pulumi.RegisterOutputType(DevicesOutput{})
	pulumi.RegisterOutputType(DevicesArrayOutput{})
}
