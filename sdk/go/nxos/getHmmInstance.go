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

// This data source can read the HMM Fabric Forwarding Instance configuration.
//
// - API Documentation: [hmmFwdInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:FwdInst/)
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
//			_, err := nxos.LookupHmmInstance(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupHmmInstance(ctx *pulumi.Context, args *LookupHmmInstanceArgs, opts ...pulumi.InvokeOption) (*LookupHmmInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupHmmInstanceResult
	err := ctx.Invoke("nxos:index/getHmmInstance:getHmmInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHmmInstance.
type LookupHmmInstanceArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// A collection of values returned by getHmmInstance.
type LookupHmmInstanceResult struct {
	// Administrative state.
	AdminState string `pulumi:"adminState"`
	// Anycast Gateway MAC address.
	AnycastMac string `pulumi:"anycastMac"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
}

func LookupHmmInstanceOutput(ctx *pulumi.Context, args LookupHmmInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupHmmInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHmmInstanceResult, error) {
			args := v.(LookupHmmInstanceArgs)
			r, err := LookupHmmInstance(ctx, &args, opts...)
			var s LookupHmmInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHmmInstanceResultOutput)
}

// A collection of arguments for invoking getHmmInstance.
type LookupHmmInstanceOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupHmmInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHmmInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getHmmInstance.
type LookupHmmInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupHmmInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHmmInstanceResult)(nil)).Elem()
}

func (o LookupHmmInstanceResultOutput) ToLookupHmmInstanceResultOutput() LookupHmmInstanceResultOutput {
	return o
}

func (o LookupHmmInstanceResultOutput) ToLookupHmmInstanceResultOutputWithContext(ctx context.Context) LookupHmmInstanceResultOutput {
	return o
}

func (o LookupHmmInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupHmmInstanceResult] {
	return pulumix.Output[LookupHmmInstanceResult]{
		OutputState: o.OutputState,
	}
}

// Administrative state.
func (o LookupHmmInstanceResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHmmInstanceResult) string { return v.AdminState }).(pulumi.StringOutput)
}

// Anycast Gateway MAC address.
func (o LookupHmmInstanceResultOutput) AnycastMac() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHmmInstanceResult) string { return v.AnycastMac }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupHmmInstanceResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHmmInstanceResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupHmmInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHmmInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHmmInstanceResultOutput{})
}
