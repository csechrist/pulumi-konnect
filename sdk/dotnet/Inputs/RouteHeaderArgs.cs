// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect.Inputs
{

    public sealed class RouteHeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(Required, String)** Name of header this route should require.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// **(Required, List of String)** Allowed values this header should equal.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public RouteHeaderArgs()
        {
        }
        public static new RouteHeaderArgs Empty => new RouteHeaderArgs();
    }
}
