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
    /// Represents an API key credential for a consumer within a control plane
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
    ///     var example = new Konnect.ConsumerKey("example", new()
    ///     {
    ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
    ///         ConsumerId = consumer.Apply(getConsumerResult =&gt; getConsumerResult.ConsumerId),
    ///         Key = "my-api-key",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Consumer keys can be imported using a proper value of `id` as described above
    /// </summary>
    [KonnectResourceType("konnect:index/consumerKey:ConsumerKey")]
    public partial class ConsumerKey : global::Pulumi.CustomResource
    {
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
        /// **(Optional/Computed, String)** The API key value.  If left out, a key will be generated for you.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// **(String)** Id of the consumer API key alone
        /// </summary>
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;


        /// <summary>
        /// Create a ConsumerKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConsumerKey(string name, ConsumerKeyArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/consumerKey:ConsumerKey", name, args ?? new ConsumerKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConsumerKey(string name, Input<string> id, ConsumerKeyState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/consumerKey:ConsumerKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConsumerKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConsumerKey Get(string name, Input<string> id, ConsumerKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new ConsumerKey(name, id, state, options);
        }
    }

    public sealed class ConsumerKeyArgs : global::Pulumi.ResourceArgs
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
        /// **(Optional/Computed, String)** The API key value.  If left out, a key will be generated for you.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        public ConsumerKeyArgs()
        {
        }
        public static new ConsumerKeyArgs Empty => new ConsumerKeyArgs();
    }

    public sealed class ConsumerKeyState : global::Pulumi.ResourceArgs
    {
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
        /// **(Optional/Computed, String)** The API key value.  If left out, a key will be generated for you.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// **(String)** Id of the consumer API key alone
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        public ConsumerKeyState()
        {
        }
        public static new ConsumerKeyState Empty => new ConsumerKeyState();
    }
}
