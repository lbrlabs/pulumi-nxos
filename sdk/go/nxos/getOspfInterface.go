// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source can read the OSPF interface configuration.
//
// - API Documentation: [ospfIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:If/)
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
//			_, err := nxos.LookupOspfInterface(ctx, &nxos.LookupOspfInterfaceArgs{
//				InstanceName: "OSPF1",
//				InterfaceId:  "eth1/10",
//				VrfName:      "VRF1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOspfInterface(ctx *pulumi.Context, args *LookupOspfInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupOspfInterfaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOspfInterfaceResult
	err := ctx.Invoke("nxos:index/getOspfInterface:getOspfInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOspfInterface.
type LookupOspfInterfaceArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// OSPF instance name.
	InstanceName string `pulumi:"instanceName"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

// A collection of values returned by getOspfInterface.
type LookupOspfInterfaceResult struct {
	// Advertise secondary IP addresses.
	AdvertiseSecondaries bool `pulumi:"advertiseSecondaries"`
	// Area identifier to which a network or interface belongs in IPv4 address format.
	Area string `pulumi:"area"`
	// Bidirectional Forwarding Detection (BFD).
	Bfd string `pulumi:"bfd"`
	// Specifies the cost of interface.
	Cost int `pulumi:"cost"`
	// Dead interval, interval after which router declares that neighbor as down.
	DeadInterval int `pulumi:"deadInterval"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Hello interval, interval between hello packets that OSPF sends on the interface.
	HelloInterval int `pulumi:"helloInterval"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// OSPF instance name.
	InstanceName string `pulumi:"instanceName"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
	// Network type.
	NetworkType string `pulumi:"networkType"`
	// Passive interface control. Interface can be configured as passive or non-passive.
	Passive string `pulumi:"passive"`
	// Priority, used in determining the designated router on this network.
	Priority int `pulumi:"priority"`
	// VRF name.
	VrfName string `pulumi:"vrfName"`
}

func LookupOspfInterfaceOutput(ctx *pulumi.Context, args LookupOspfInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupOspfInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOspfInterfaceResult, error) {
			args := v.(LookupOspfInterfaceArgs)
			r, err := LookupOspfInterface(ctx, &args, opts...)
			var s LookupOspfInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOspfInterfaceResultOutput)
}

// A collection of arguments for invoking getOspfInterface.
type LookupOspfInterfaceOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// OSPF instance name.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringInput `pulumi:"interfaceId"`
	// VRF name.
	VrfName pulumi.StringInput `pulumi:"vrfName"`
}

func (LookupOspfInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfInterfaceArgs)(nil)).Elem()
}

// A collection of values returned by getOspfInterface.
type LookupOspfInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupOspfInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfInterfaceResult)(nil)).Elem()
}

func (o LookupOspfInterfaceResultOutput) ToLookupOspfInterfaceResultOutput() LookupOspfInterfaceResultOutput {
	return o
}

func (o LookupOspfInterfaceResultOutput) ToLookupOspfInterfaceResultOutputWithContext(ctx context.Context) LookupOspfInterfaceResultOutput {
	return o
}

func (o LookupOspfInterfaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOspfInterfaceResult] {
	return pulumix.Output[LookupOspfInterfaceResult]{
		OutputState: o.OutputState,
	}
}

// Advertise secondary IP addresses.
func (o LookupOspfInterfaceResultOutput) AdvertiseSecondaries() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) bool { return v.AdvertiseSecondaries }).(pulumi.BoolOutput)
}

// Area identifier to which a network or interface belongs in IPv4 address format.
func (o LookupOspfInterfaceResultOutput) Area() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.Area }).(pulumi.StringOutput)
}

// Bidirectional Forwarding Detection (BFD).
func (o LookupOspfInterfaceResultOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.Bfd }).(pulumi.StringOutput)
}

// Specifies the cost of interface.
func (o LookupOspfInterfaceResultOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) int { return v.Cost }).(pulumi.IntOutput)
}

// Dead interval, interval after which router declares that neighbor as down.
func (o LookupOspfInterfaceResultOutput) DeadInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) int { return v.DeadInterval }).(pulumi.IntOutput)
}

// A device name from the provider configuration.
func (o LookupOspfInterfaceResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// Hello interval, interval between hello packets that OSPF sends on the interface.
func (o LookupOspfInterfaceResultOutput) HelloInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) int { return v.HelloInterval }).(pulumi.IntOutput)
}

// The distinguished name of the object.
func (o LookupOspfInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// OSPF instance name.
func (o LookupOspfInterfaceResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
func (o LookupOspfInterfaceResultOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.InterfaceId }).(pulumi.StringOutput)
}

// Network type.
func (o LookupOspfInterfaceResultOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.NetworkType }).(pulumi.StringOutput)
}

// Passive interface control. Interface can be configured as passive or non-passive.
func (o LookupOspfInterfaceResultOutput) Passive() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.Passive }).(pulumi.StringOutput)
}

// Priority, used in determining the designated router on this network.
func (o LookupOspfInterfaceResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) int { return v.Priority }).(pulumi.IntOutput)
}

// VRF name.
func (o LookupOspfInterfaceResultOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfInterfaceResult) string { return v.VrfName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOspfInterfaceResultOutput{})
}
