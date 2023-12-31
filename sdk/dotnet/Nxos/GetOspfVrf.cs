// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos
{
    public static class GetOspfVrf
    {
        /// <summary>
        /// This data source can read the OSPF VRF configuration.
        /// 
        /// - API Documentation: [ospfDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Dom/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetOspfVrf.Invoke(new()
        ///     {
        ///         InstanceName = "OSPF1",
        ///         Name = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOspfVrfResult> InvokeAsync(GetOspfVrfArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOspfVrfResult>("nxos:index/getOspfVrf:getOspfVrf", args ?? new GetOspfVrfArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the OSPF VRF configuration.
        /// 
        /// - API Documentation: [ospfDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Dom/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetOspfVrf.Invoke(new()
        ///     {
        ///         InstanceName = "OSPF1",
        ///         Name = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOspfVrfResult> Invoke(GetOspfVrfInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOspfVrfResult>("nxos:index/getOspfVrf:getOspfVrf", args ?? new GetOspfVrfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOspfVrfArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// OSPF instance name.
        /// </summary>
        [Input("instanceName", required: true)]
        public string InstanceName { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetOspfVrfArgs()
        {
        }
        public static new GetOspfVrfArgs Empty => new GetOspfVrfArgs();
    }

    public sealed class GetOspfVrfInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// OSPF instance name.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetOspfVrfInvokeArgs()
        {
        }
        public static new GetOspfVrfInvokeArgs Empty => new GetOspfVrfInvokeArgs();
    }


    [OutputType]
    public sealed class GetOspfVrfResult
    {
        /// <summary>
        /// Administrative state.
        /// </summary>
        public readonly string AdminState;
        /// <summary>
        /// Bandwidth reference value.
        /// </summary>
        public readonly int BandwidthReference;
        /// <summary>
        /// Bandwidth reference unit.
        /// </summary>
        public readonly string BanwidthReferenceUnit;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// Administrative distance preference.
        /// </summary>
        public readonly int Distance;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// OSPF instance name.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Router ID.
        /// </summary>
        public readonly string RouterId;

        [OutputConstructor]
        private GetOspfVrfResult(
            string adminState,

            int bandwidthReference,

            string banwidthReferenceUnit,

            string? device,

            int distance,

            string id,

            string instanceName,

            string name,

            string routerId)
        {
            AdminState = adminState;
            BandwidthReference = bandwidthReference;
            BanwidthReferenceUnit = banwidthReferenceUnit;
            Device = device;
            Distance = distance;
            Id = id;
            InstanceName = instanceName;
            Name = name;
            RouterId = routerId;
        }
    }
}
