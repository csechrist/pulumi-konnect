// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetConsumer
    {
        /// <summary>
        /// Represents a consumer
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = Konnect.GetConsumer.Invoke(new()
        ///     {
        ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
        ///         SearchUsername = "Bob",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetConsumerResult> InvokeAsync(GetConsumerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConsumerResult>("konnect:index/getConsumer:getConsumer", args ?? new GetConsumerArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a consumer
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = Konnect.GetConsumer.Invoke(new()
        ///     {
        ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
        ///         SearchUsername = "Bob",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetConsumerResult> Invoke(GetConsumerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConsumerResult>("konnect:index/getConsumer:getConsumer", args ?? new GetConsumerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConsumerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// **(Required, String)** The id of the control plane containing consumer
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public string ControlPlaneId { get; set; } = null!;

        /// <summary>
        /// **(Optional, String)** The filter string to apply to the custom_id of the consumer. Uses equality.
        /// </summary>
        [Input("customId")]
        public string? CustomId { get; set; }

        /// <summary>
        /// **(Optional, String)** The search string to apply to the custom_id of the consumer. Uses contains.
        /// </summary>
        [Input("searchCustomId")]
        public string? SearchCustomId { get; set; }

        /// <summary>
        /// **(Optional, String)** The search string to apply to the username of the consumer. Uses contains.
        /// </summary>
        [Input("searchUsername")]
        public string? SearchUsername { get; set; }

        /// <summary>
        /// **(Optional, String)** The filter string to apply to the username of the consumer. Uses equality.
        /// </summary>
        [Input("username")]
        public string? Username { get; set; }

        public GetConsumerArgs()
        {
        }
        public static new GetConsumerArgs Empty => new GetConsumerArgs();
    }

    public sealed class GetConsumerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// **(Required, String)** The id of the control plane containing consumer
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        /// <summary>
        /// **(Optional, String)** The filter string to apply to the custom_id of the consumer. Uses equality.
        /// </summary>
        [Input("customId")]
        public Input<string>? CustomId { get; set; }

        /// <summary>
        /// **(Optional, String)** The search string to apply to the custom_id of the consumer. Uses contains.
        /// </summary>
        [Input("searchCustomId")]
        public Input<string>? SearchCustomId { get; set; }

        /// <summary>
        /// **(Optional, String)** The search string to apply to the username of the consumer. Uses contains.
        /// </summary>
        [Input("searchUsername")]
        public Input<string>? SearchUsername { get; set; }

        /// <summary>
        /// **(Optional, String)** The filter string to apply to the username of the consumer. Uses equality.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GetConsumerInvokeArgs()
        {
        }
        public static new GetConsumerInvokeArgs Empty => new GetConsumerInvokeArgs();
    }


    [OutputType]
    public sealed class GetConsumerResult
    {
        /// <summary>
        /// **(String)** Id of the consumer alone
        /// </summary>
        public readonly string ConsumerId;
        public readonly string ControlPlaneId;
        public readonly string? CustomId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? SearchCustomId;
        public readonly string? SearchUsername;
        public readonly string? Username;

        [OutputConstructor]
        private GetConsumerResult(
            string consumerId,

            string controlPlaneId,

            string? customId,

            string id,

            string? searchCustomId,

            string? searchUsername,

            string? username)
        {
            ConsumerId = consumerId;
            ControlPlaneId = controlPlaneId;
            CustomId = customId;
            Id = id;
            SearchCustomId = searchCustomId;
            SearchUsername = searchUsername;
            Username = username;
        }
    }
}
