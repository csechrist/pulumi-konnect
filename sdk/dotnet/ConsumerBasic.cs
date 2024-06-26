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
    /// Represents a basic auth credential for a consumer within a control plane
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
    ///     var example = new Konnect.ConsumerBasic("example", new()
    ///     {
    ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
    ///         ConsumerId = consumer.Apply(getConsumerResult =&gt; getConsumerResult.ConsumerId),
    ///         Username = "my-username",
    ///         Password = "my-password",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Consumer basics can be imported using a proper value of `id` as described above
    /// </summary>
    [KonnectResourceType("konnect:index/consumerBasic:ConsumerBasic")]
    public partial class ConsumerBasic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// **(String)** Id of the consumer basic auth alone
        /// </summary>
        [Output("basicId")]
        public Output<string> BasicId { get; private set; } = null!;

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
        /// **(Required, String)** The password value.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// **(String)** Hash of the password
        /// </summary>
        [Output("passwordHash")]
        public Output<string> PasswordHash { get; private set; } = null!;

        /// <summary>
        /// **(Required, String)** The username value.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerBasic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerBasic(string name, ConsumerBasicArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/consumerBasic:ConsumerBasic", name, args ?? new ConsumerBasicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerBasic(string name, Input<string> id, ConsumerBasicState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/consumerBasic:ConsumerBasic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerBasic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerBasic Get(string name, Input<string> id, ConsumerBasicState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerBasic(name, id, state, options);
        }
    }

    public sealed class ConsumerBasicArgs : global::Pulumi.ResourceArgs
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
        /// **(Required, String)** The password value.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// **(Required, String)** The username value.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public ConsumerBasicArgs()
        {
        }
        public static new ConsumerBasicArgs Empty => new ConsumerBasicArgs();
    }

    public sealed class ConsumerBasicState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(String)** Id of the consumer basic auth alone
        /// </summary>
        [Input("basicId")]
        public Input<string>? BasicId { get; set; }

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
        /// **(Required, String)** The password value.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// **(String)** Hash of the password
        /// </summary>
        [Input("passwordHash")]
        public Input<string>? PasswordHash { get; set; }

        /// <summary>
        /// **(Required, String)** The username value.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ConsumerBasicState()
        {
        }
        public static new ConsumerBasicState Empty => new ConsumerBasicState();
    }
}
