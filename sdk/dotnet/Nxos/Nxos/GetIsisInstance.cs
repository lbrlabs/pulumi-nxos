// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos.Nxos
{
    public static class GetIsisInstance
    {
        public static Task<GetIsisInstanceResult> InvokeAsync(GetIsisInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIsisInstanceResult>("nxos:nxos/getIsisInstance:getIsisInstance", args ?? new GetIsisInstanceArgs(), options.WithDefaults());

        public static Output<GetIsisInstanceResult> Invoke(GetIsisInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIsisInstanceResult>("nxos:nxos/getIsisInstance:getIsisInstance", args ?? new GetIsisInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIsisInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIsisInstanceArgs()
        {
        }
        public static new GetIsisInstanceArgs Empty => new GetIsisInstanceArgs();
    }

    public sealed class GetIsisInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetIsisInstanceInvokeArgs()
        {
        }
        public static new GetIsisInstanceInvokeArgs Empty => new GetIsisInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetIsisInstanceResult
    {
        public readonly string AdminState;
        public readonly string? Device;
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetIsisInstanceResult(
            string adminState,

            string? device,

            string id,

            string name)
        {
            AdminState = adminState;
            Device = device;
            Id = id;
            Name = name;
        }
    }
}