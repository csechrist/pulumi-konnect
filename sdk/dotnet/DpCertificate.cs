// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    [KonnectResourceType("konnect:index/dpCertificate:DpCertificate")]
    public partial class DpCertificate : global::Pulumi.CustomResource
    {
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        [Output("controlPlaneId")]
        public Output<string> ControlPlaneId { get; private set; } = null!;


        /// <summary>
        /// Create a DpCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DpCertificate(string name, DpCertificateArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/dpCertificate:DpCertificate", name, args ?? new DpCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DpCertificate(string name, Input<string> id, DpCertificateState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/dpCertificate:DpCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DpCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DpCertificate Get(string name, Input<string> id, DpCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new DpCertificate(name, id, state, options);
        }
    }

    public sealed class DpCertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("cert", required: true)]
        public Input<string> Cert { get; set; } = null!;

        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        public DpCertificateArgs()
        {
        }
        public static new DpCertificateArgs Empty => new DpCertificateArgs();
    }

    public sealed class DpCertificateState : global::Pulumi.ResourceArgs
    {
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        [Input("controlPlaneId")]
        public Input<string>? ControlPlaneId { get; set; }

        public DpCertificateState()
        {
        }
        public static new DpCertificateState Empty => new DpCertificateState();
    }
}