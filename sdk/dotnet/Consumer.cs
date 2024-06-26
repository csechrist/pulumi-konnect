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
    /// Represents a consumer within a control plane
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
    ///     var example = new Konnect.Consumer("example", new()
    ///     {
    ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
    ///         Username = "testuser",
    ///         CustomId = "testcustom",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Consumers can be imported using a proper value of `id` as described above
    /// </summary>
    [KonnectResourceType("konnect:index/consumer:Consumer")]
    public partial class Consumer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// **(String)** Id of the consumer alone
        /// </summary>
        [Output("consumerId")]
        public Output<string> ConsumerId { get; private set; } = null!;

        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Output("controlPlaneId")]
        public Output<string> ControlPlaneId { get; private set; } = null!;

        /// <summary>
        /// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
        /// </summary>
        [Output("customId")]
        public Output<string?> CustomId { get; private set; } = null!;

        /// <summary>
        /// **(Optional, String)** The unique username of the Consumer.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Consumer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Consumer(string name, ConsumerArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/consumer:Consumer", name, args ?? new ConsumerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Consumer(string name, Input<string> id, ConsumerState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/consumer:Consumer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Consumer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Consumer Get(string name, Input<string> id, ConsumerState? state = null, CustomResourceOptions? options = null)
        {
            return new Consumer(name, id, state, options);
        }
    }

    public sealed class ConsumerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        /// <summary>
        /// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
        /// </summary>
        [Input("customId")]
        public Input<string>? CustomId { get; set; }

        /// <summary>
        /// **(Optional, String)** The unique username of the Consumer.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ConsumerArgs()
        {
        }
        public static new ConsumerArgs Empty => new ConsumerArgs();
    }

    public sealed class ConsumerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(String)** Id of the consumer alone
        /// </summary>
        [Input("consumerId")]
        public Input<string>? ConsumerId { get; set; }

        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Input("controlPlaneId")]
        public Input<string>? ControlPlaneId { get; set; }

        /// <summary>
        /// **(Optional, String)** Field for storing an existing unique ID for the Consumer.
        /// </summary>
        [Input("customId")]
        public Input<string>? CustomId { get; set; }

        /// <summary>
        /// **(Optional, String)** The unique username of the Consumer.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public ConsumerState()
        {
        }
        public static new ConsumerState Empty => new ConsumerState();
    }
}
