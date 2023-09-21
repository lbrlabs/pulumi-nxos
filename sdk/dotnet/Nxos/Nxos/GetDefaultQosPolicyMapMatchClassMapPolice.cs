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
    public static class GetDefaultQosPolicyMapMatchClassMapPolice
    {
        public static Task<GetDefaultQosPolicyMapMatchClassMapPoliceResult> InvokeAsync(GetDefaultQosPolicyMapMatchClassMapPoliceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDefaultQosPolicyMapMatchClassMapPoliceResult>("nxos:nxos/getDefaultQosPolicyMapMatchClassMapPolice:getDefaultQosPolicyMapMatchClassMapPolice", args ?? new GetDefaultQosPolicyMapMatchClassMapPoliceArgs(), options.WithDefaults());

        public static Output<GetDefaultQosPolicyMapMatchClassMapPoliceResult> Invoke(GetDefaultQosPolicyMapMatchClassMapPoliceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDefaultQosPolicyMapMatchClassMapPoliceResult>("nxos:nxos/getDefaultQosPolicyMapMatchClassMapPolice:getDefaultQosPolicyMapMatchClassMapPolice", args ?? new GetDefaultQosPolicyMapMatchClassMapPoliceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDefaultQosPolicyMapMatchClassMapPoliceArgs : global::Pulumi.InvokeArgs
    {
        [Input("classMapName", required: true)]
        public string ClassMapName { get; set; } = null!;

        [Input("device")]
        public string? Device { get; set; }

        [Input("policyMapName", required: true)]
        public string PolicyMapName { get; set; } = null!;

        public GetDefaultQosPolicyMapMatchClassMapPoliceArgs()
        {
        }
        public static new GetDefaultQosPolicyMapMatchClassMapPoliceArgs Empty => new GetDefaultQosPolicyMapMatchClassMapPoliceArgs();
    }

    public sealed class GetDefaultQosPolicyMapMatchClassMapPoliceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("classMapName", required: true)]
        public Input<string> ClassMapName { get; set; } = null!;

        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("policyMapName", required: true)]
        public Input<string> PolicyMapName { get; set; } = null!;

        public GetDefaultQosPolicyMapMatchClassMapPoliceInvokeArgs()
        {
        }
        public static new GetDefaultQosPolicyMapMatchClassMapPoliceInvokeArgs Empty => new GetDefaultQosPolicyMapMatchClassMapPoliceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDefaultQosPolicyMapMatchClassMapPoliceResult
    {
        public readonly int BcRate;
        public readonly string BcUnit;
        public readonly int BeRate;
        public readonly string BeUnit;
        public readonly int CirRate;
        public readonly string CirUnit;
        public readonly string ClassMapName;
        public readonly string ConformAction;
        public readonly int ConformSetCos;
        public readonly int ConformSetDscp;
        public readonly string ConformSetPrecedence;
        public readonly int ConformSetQosGroup;
        public readonly string? Device;
        public readonly string ExceedAction;
        public readonly int ExceedSetCos;
        public readonly int ExceedSetDscp;
        public readonly string ExceedSetPrecedence;
        public readonly int ExceedSetQosGroup;
        public readonly string Id;
        public readonly int PirRate;
        public readonly string PirUnit;
        public readonly string PolicyMapName;
        public readonly string ViolateAction;
        public readonly int ViolateSetCos;
        public readonly int ViolateSetDscp;
        public readonly string ViolateSetPrecedence;
        public readonly int ViolateSetQosGroup;

        [OutputConstructor]
        private GetDefaultQosPolicyMapMatchClassMapPoliceResult(
            int bcRate,

            string bcUnit,

            int beRate,

            string beUnit,

            int cirRate,

            string cirUnit,

            string classMapName,

            string conformAction,

            int conformSetCos,

            int conformSetDscp,

            string conformSetPrecedence,

            int conformSetQosGroup,

            string? device,

            string exceedAction,

            int exceedSetCos,

            int exceedSetDscp,

            string exceedSetPrecedence,

            int exceedSetQosGroup,

            string id,

            int pirRate,

            string pirUnit,

            string policyMapName,

            string violateAction,

            int violateSetCos,

            int violateSetDscp,

            string violateSetPrecedence,

            int violateSetQosGroup)
        {
            BcRate = bcRate;
            BcUnit = bcUnit;
            BeRate = beRate;
            BeUnit = beUnit;
            CirRate = cirRate;
            CirUnit = cirUnit;
            ClassMapName = classMapName;
            ConformAction = conformAction;
            ConformSetCos = conformSetCos;
            ConformSetDscp = conformSetDscp;
            ConformSetPrecedence = conformSetPrecedence;
            ConformSetQosGroup = conformSetQosGroup;
            Device = device;
            ExceedAction = exceedAction;
            ExceedSetCos = exceedSetCos;
            ExceedSetDscp = exceedSetDscp;
            ExceedSetPrecedence = exceedSetPrecedence;
            ExceedSetQosGroup = exceedSetQosGroup;
            Id = id;
            PirRate = pirRate;
            PirUnit = pirUnit;
            PolicyMapName = policyMapName;
            ViolateAction = violateAction;
            ViolateSetCos = violateSetCos;
            ViolateSetDscp = violateSetDscp;
            ViolateSetPrecedence = violateSetPrecedence;
            ViolateSetQosGroup = violateSetQosGroup;
        }
    }
}