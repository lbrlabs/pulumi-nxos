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
    public static class GetNtpServer
    {
        /// <summary>
        /// This data source can read an ntp server or peer.
        /// 
        /// - API Documentation: [datetimeNtpProvider](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/datetime:NtpProvider/)
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
        ///     var example = Nxos.GetNtpServer.Invoke(new()
        ///     {
        ///         Name = "1.2.3.4",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNtpServerResult> InvokeAsync(GetNtpServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNtpServerResult>("nxos:index/getNtpServer:getNtpServer", args ?? new GetNtpServerArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read an ntp server or peer.
        /// 
        /// - API Documentation: [datetimeNtpProvider](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/datetime:NtpProvider/)
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
        ///     var example = Nxos.GetNtpServer.Invoke(new()
        ///     {
        ///         Name = "1.2.3.4",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNtpServerResult> Invoke(GetNtpServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNtpServerResult>("nxos:index/getNtpServer:getNtpServer", args ?? new GetNtpServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNtpServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// NTP server.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNtpServerArgs()
        {
        }
        public static new GetNtpServerArgs Empty => new GetNtpServerArgs();
    }

    public sealed class GetNtpServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// NTP server.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNtpServerInvokeArgs()
        {
        }
        public static new GetNtpServerInvokeArgs Empty => new GetNtpServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetNtpServerResult
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// NTP provider key ID. Possible range is from `1` to `65535`.
        /// </summary>
        public readonly int KeyId;
        /// <summary>
        /// NTP maximum interval default in seconds. Possible range is from `4` to `16`.
        /// </summary>
        public readonly int MaxPoll;
        /// <summary>
        /// NTP minimum interval default in seconds. Possible range is from `4` to `16`.
        /// </summary>
        public readonly int MinPoll;
        /// <summary>
        /// NTP server.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// NTP provider type. Possible values are `server` or `peer`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Identifies the VRF for the NTP providers.
        /// </summary>
        public readonly string Vrf;

        [OutputConstructor]
        private GetNtpServerResult(
            string? device,

            string id,

            int keyId,

            int maxPoll,

            int minPoll,

            string name,

            string type,

            string vrf)
        {
            Device = device;
            Id = id;
            KeyId = keyId;
            MaxPoll = maxPoll;
            MinPoll = minPoll;
            Name = name;
            Type = type;
            Vrf = vrf;
        }
    }
}
