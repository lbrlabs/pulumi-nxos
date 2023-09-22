// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos.Inputs
{

    public sealed class RestChildrenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Class name of the child object.
        /// </summary>
        [Input("className", required: true)]
        public Input<string> ClassName { get; set; } = null!;

        [Input("content")]
        private InputMap<string>? _content;

        /// <summary>
        /// Map of key-value pairs which represents the attributes of the child object.
        /// </summary>
        public InputMap<string> Content
        {
            get => _content ?? (_content = new InputMap<string>());
            set => _content = value;
        }

        /// <summary>
        /// The relative name of the child object.
        /// </summary>
        [Input("rn", required: true)]
        public Input<string> Rn { get; set; } = null!;

        public RestChildrenArgs()
        {
        }
        public static new RestChildrenArgs Empty => new RestChildrenArgs();
    }
}