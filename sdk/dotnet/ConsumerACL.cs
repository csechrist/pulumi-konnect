// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    /// <summary>
    /// Represents an ACL credential for a consumer within a control plane
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Konnect = Pulumi.Konnect;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var controlPlane = Konnect.GetControlPlane.Invoke(new()
    ///     {
    ///         Name = "TestControlPlane",
    ///     });
    /// 
    ///     var consumer = Konnect.GetConsumer.Invoke(new()
    ///     {
    ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
    ///         SearchUsername = "Bob",
    ///     });
    /// 
    ///     var example = new Konnect.ConsumerACL("example", new()
    ///     {
    ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
    ///         ConsumerId = consumer.Apply(getConsumerResult =&gt; getConsumerResult.ConsumerId),
    ///         Group = "my-acl-group",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Consumer ACLs can be imported using a proper value of `id` as described above
    /// </summary>
    [KonnectResourceType("konnect:index/consumerACL:ConsumerACL")]
    public partial class ConsumerACL : global::Pulumi.CustomResource
    {
        /// <summary>
        /// **(String)** Id of the consumer ACL alone
        /// </summary>
        [Output("aclId")]
        public Output<string> AclId { get; private set; } = null!;

        /// <summary>
        /// **(Required, String)** The id of the consumer.
        /// </summary>
        [Output("consumerId")]
        public Output<string> ConsumerId { get; private set; } = null!;

        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Output("controlPlaneId")]
        public Output<string> ControlPlaneId { get; private set; } = null!;

        /// <summary>
        /// **(Required, String)** The ACL group value.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerACL resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerACL(string name, ConsumerACLArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/consumerACL:ConsumerACL", name, args ?? new ConsumerACLArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerACL(string name, Input<string> id, ConsumerACLState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/consumerACL:ConsumerACL", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/csechrist/pulumi-konnect",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConsumerACL resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerACL Get(string name, Input<string> id, ConsumerACLState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerACL(name, id, state, options);
        }
    }

    public sealed class ConsumerACLArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(Required, String)** The id of the consumer.
        /// </summary>
        [Input("consumerId", required: true)]
        public Input<string> ConsumerId { get; set; } = null!;

        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        /// <summary>
        /// **(Required, String)** The ACL group value.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        public ConsumerACLArgs()
        {
        }
        public static new ConsumerACLArgs Empty => new ConsumerACLArgs();
    }

    public sealed class ConsumerACLState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(String)** Id of the consumer ACL alone
        /// </summary>
        [Input("aclId")]
        public Input<string>? AclId { get; set; }

        /// <summary>
        /// **(Required, String)** The id of the consumer.
        /// </summary>
        [Input("consumerId")]
        public Input<string>? ConsumerId { get; set; }

        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Input("controlPlaneId")]
        public Input<string>? ControlPlaneId { get; set; }

        /// <summary>
        /// **(Required, String)** The ACL group value.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        public ConsumerACLState()
        {
        }
        public static new ConsumerACLState Empty => new ConsumerACLState();
    }
}
