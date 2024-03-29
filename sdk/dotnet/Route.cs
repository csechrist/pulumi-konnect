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
    /// Represents a route within a control plane
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
    ///     var service = new Konnect.Service("service", new()
    ///     {
    ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
    ///         Host = "mockbin.org",
    ///     });
    /// 
    ///     var example = new Konnect.Route("example", new()
    ///     {
    ///         ControlPlaneId = controlPlane.Apply(getControlPlaneResult =&gt; getControlPlaneResult.Id),
    ///         ServiceId = service.ServiceId,
    ///         Protocols = new[]
    ///         {
    ///             "http",
    ///         },
    ///         Paths = new[]
    ///         {
    ///             "/example",
    ///         },
    ///         Headers = new[]
    ///         {
    ///             new Konnect.Inputs.RouteHeaderArgs
    ///             {
    ///                 Name = "required-header",
    ///                 Values = new[]
    ///                 {
    ///                     "required-header-values",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Routes can be imported using a proper value of `id` as described above
    /// </summary>
    [KonnectResourceType("konnect:index/route:Route")]
    public partial class Route : global::Pulumi.CustomResource
    {
        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Output("controlPlaneId")]
        public Output<string> ControlPlaneId { get; private set; } = null!;

        /// <summary>
        /// **(Optional, set{header})** Configuration block for a header.  Can be specified multiple times for each header.  Each block supports the fields documented below.
        /// </summary>
        [Output("headers")]
        public Output<ImmutableArray<Outputs.RouteHeader>> Headers { get; private set; } = null!;

        /// <summary>
        /// **(Optional, List of String)** The hosts this route should allow.
        /// </summary>
        [Output("hosts")]
        public Output<ImmutableArray<string>> Hosts { get; private set; } = null!;

        /// <summary>
        /// **(Optional, Integer)** The status code Kong responds with when all properties of a Route match except the protocol. Allowed values: `426`, `301`, `302`, `307`, `308`. Default: `426`
        /// </summary>
        [Output("httpsRedirectStatusCode")]
        public Output<int?> HttpsRedirectStatusCode { get; private set; } = null!;

        /// <summary>
        /// **(Optional, List of String)** The methods this route should allow. Allowed values: `GET`, `PUT`, `POST`, `PATCH`, `DELETE`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`
        /// </summary>
        [Output("methods")]
        public Output<ImmutableArray<string>> Methods { get; private set; } = null!;

        /// <summary>
        /// **(Optional, String)** The name of the route.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// **(Optional, String)** Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. Allowed values: `v0`, `v1`. Default: `v0`
        /// </summary>
        [Output("pathHandling")]
        public Output<string?> PathHandling { get; private set; } = null!;

        /// <summary>
        /// **(Optional, List of String)** The paths this route should allow.
        /// </summary>
        [Output("paths")]
        public Output<ImmutableArray<string>> Paths { get; private set; } = null!;

        /// <summary>
        /// **(Optional, Boolean)** Whether to use the request `Host` header during the Service request. Default: `false`
        /// </summary>
        [Output("preserveHost")]
        public Output<bool?> PreserveHost { get; private set; } = null!;

        /// <summary>
        /// **(Optional, List of String)** The protocols this route should allow. Allowed values: `http`, `https`
        /// </summary>
        [Output("protocols")]
        public Output<ImmutableArray<string>> Protocols { get; private set; } = null!;

        /// <summary>
        /// **(Optional, Integer)** A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. Default: `0`
        /// </summary>
        [Output("regexPriority")]
        public Output<int?> RegexPriority { get; private set; } = null!;

        /// <summary>
        /// **(Optional, Boolean)** Whether to enable request body buffering. Default: `true`
        /// </summary>
        [Output("requestBuffering")]
        public Output<bool?> RequestBuffering { get; private set; } = null!;

        /// <summary>
        /// **(Optional, Boolean)** Whether to enable response body buffering. Default: `true`
        /// </summary>
        [Output("responseBuffering")]
        public Output<bool?> ResponseBuffering { get; private set; } = null!;

        /// <summary>
        /// **(String)** Id of the route alone
        /// </summary>
        [Output("routeId")]
        public Output<string> RouteId { get; private set; } = null!;

        /// <summary>
        /// **(Optional, String)** The id of the service to forward traffic to.
        /// </summary>
        [Output("serviceId")]
        public Output<string?> ServiceId { get; private set; } = null!;

        /// <summary>
        /// **(Optional, Boolean)** Whether to strip the matching prefix from the Service request. Default: `true`
        /// </summary>
        [Output("stripPath")]
        public Output<bool?> StripPath { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/route:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/route:Route", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
        {
            return new Route(name, id, state, options);
        }
    }

    public sealed class RouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Input("controlPlaneId", required: true)]
        public Input<string> ControlPlaneId { get; set; } = null!;

        [Input("headers")]
        private InputList<Inputs.RouteHeaderArgs>? _headers;

        /// <summary>
        /// **(Optional, set{header})** Configuration block for a header.  Can be specified multiple times for each header.  Each block supports the fields documented below.
        /// </summary>
        public InputList<Inputs.RouteHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.RouteHeaderArgs>());
            set => _headers = value;
        }

        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// **(Optional, List of String)** The hosts this route should allow.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// **(Optional, Integer)** The status code Kong responds with when all properties of a Route match except the protocol. Allowed values: `426`, `301`, `302`, `307`, `308`. Default: `426`
        /// </summary>
        [Input("httpsRedirectStatusCode")]
        public Input<int>? HttpsRedirectStatusCode { get; set; }

        [Input("methods")]
        private InputList<string>? _methods;

        /// <summary>
        /// **(Optional, List of String)** The methods this route should allow. Allowed values: `GET`, `PUT`, `POST`, `PATCH`, `DELETE`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`
        /// </summary>
        public InputList<string> Methods
        {
            get => _methods ?? (_methods = new InputList<string>());
            set => _methods = value;
        }

        /// <summary>
        /// **(Optional, String)** The name of the route.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// **(Optional, String)** Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. Allowed values: `v0`, `v1`. Default: `v0`
        /// </summary>
        [Input("pathHandling")]
        public Input<string>? PathHandling { get; set; }

        [Input("paths")]
        private InputList<string>? _paths;

        /// <summary>
        /// **(Optional, List of String)** The paths this route should allow.
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        /// <summary>
        /// **(Optional, Boolean)** Whether to use the request `Host` header during the Service request. Default: `false`
        /// </summary>
        [Input("preserveHost")]
        public Input<bool>? PreserveHost { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// **(Optional, List of String)** The protocols this route should allow. Allowed values: `http`, `https`
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// **(Optional, Integer)** A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. Default: `0`
        /// </summary>
        [Input("regexPriority")]
        public Input<int>? RegexPriority { get; set; }

        /// <summary>
        /// **(Optional, Boolean)** Whether to enable request body buffering. Default: `true`
        /// </summary>
        [Input("requestBuffering")]
        public Input<bool>? RequestBuffering { get; set; }

        /// <summary>
        /// **(Optional, Boolean)** Whether to enable response body buffering. Default: `true`
        /// </summary>
        [Input("responseBuffering")]
        public Input<bool>? ResponseBuffering { get; set; }

        /// <summary>
        /// **(Optional, String)** The id of the service to forward traffic to.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// **(Optional, Boolean)** Whether to strip the matching prefix from the Service request. Default: `true`
        /// </summary>
        [Input("stripPath")]
        public Input<bool>? StripPath { get; set; }

        public RouteArgs()
        {
        }
        public static new RouteArgs Empty => new RouteArgs();
    }

    public sealed class RouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(Required, String)** The id of the control plane.
        /// </summary>
        [Input("controlPlaneId")]
        public Input<string>? ControlPlaneId { get; set; }

        [Input("headers")]
        private InputList<Inputs.RouteHeaderGetArgs>? _headers;

        /// <summary>
        /// **(Optional, set{header})** Configuration block for a header.  Can be specified multiple times for each header.  Each block supports the fields documented below.
        /// </summary>
        public InputList<Inputs.RouteHeaderGetArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.RouteHeaderGetArgs>());
            set => _headers = value;
        }

        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// **(Optional, List of String)** The hosts this route should allow.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// **(Optional, Integer)** The status code Kong responds with when all properties of a Route match except the protocol. Allowed values: `426`, `301`, `302`, `307`, `308`. Default: `426`
        /// </summary>
        [Input("httpsRedirectStatusCode")]
        public Input<int>? HttpsRedirectStatusCode { get; set; }

        [Input("methods")]
        private InputList<string>? _methods;

        /// <summary>
        /// **(Optional, List of String)** The methods this route should allow. Allowed values: `GET`, `PUT`, `POST`, `PATCH`, `DELETE`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`
        /// </summary>
        public InputList<string> Methods
        {
            get => _methods ?? (_methods = new InputList<string>());
            set => _methods = value;
        }

        /// <summary>
        /// **(Optional, String)** The name of the route.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// **(Optional, String)** Controls how the Service path, Route path and requested path are combined when sending a request to the upstream. Allowed values: `v0`, `v1`. Default: `v0`
        /// </summary>
        [Input("pathHandling")]
        public Input<string>? PathHandling { get; set; }

        [Input("paths")]
        private InputList<string>? _paths;

        /// <summary>
        /// **(Optional, List of String)** The paths this route should allow.
        /// </summary>
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        /// <summary>
        /// **(Optional, Boolean)** Whether to use the request `Host` header during the Service request. Default: `false`
        /// </summary>
        [Input("preserveHost")]
        public Input<bool>? PreserveHost { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// **(Optional, List of String)** The protocols this route should allow. Allowed values: `http`, `https`
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        /// <summary>
        /// **(Optional, Integer)** A number used to choose which route resolves a given request when several routes match it using regexes simultaneously. Default: `0`
        /// </summary>
        [Input("regexPriority")]
        public Input<int>? RegexPriority { get; set; }

        /// <summary>
        /// **(Optional, Boolean)** Whether to enable request body buffering. Default: `true`
        /// </summary>
        [Input("requestBuffering")]
        public Input<bool>? RequestBuffering { get; set; }

        /// <summary>
        /// **(Optional, Boolean)** Whether to enable response body buffering. Default: `true`
        /// </summary>
        [Input("responseBuffering")]
        public Input<bool>? ResponseBuffering { get; set; }

        /// <summary>
        /// **(String)** Id of the route alone
        /// </summary>
        [Input("routeId")]
        public Input<string>? RouteId { get; set; }

        /// <summary>
        /// **(Optional, String)** The id of the service to forward traffic to.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// **(Optional, Boolean)** Whether to strip the matching prefix from the Service request. Default: `true`
        /// </summary>
        [Input("stripPath")]
        public Input<bool>? StripPath { get; set; }

        public RouteState()
        {
        }
        public static new RouteState Empty => new RouteState();
    }
}
