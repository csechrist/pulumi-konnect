// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Konnect
{
    public static class GetTeam
    {
        /// <summary>
        /// Represents a team
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
        ///     var example = Konnect.GetTeam.Invoke(new()
        ///     {
        ///         SearchName = "Panther",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTeamResult> InvokeAsync(GetTeamArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTeamResult>("konnect:index/getTeam:getTeam", args ?? new GetTeamArgs(), options.WithDefaults());

        /// <summary>
        /// Represents a team
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
        ///     var example = Konnect.GetTeam.Invoke(new()
        ///     {
        ///         SearchName = "Panther",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTeamResult> Invoke(GetTeamInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTeamResult>("konnect:index/getTeam:getTeam", args ?? new GetTeamInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTeamArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// **(Optional, String)** The filter string to apply to the name of the team. Uses equality.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// **(Optional, String)** The search string to apply to the name of the team. Uses contains.
        /// </summary>
        [Input("searchName")]
        public string? SearchName { get; set; }

        public GetTeamArgs()
        {
        }
        public static new GetTeamArgs Empty => new GetTeamArgs();
    }

    public sealed class GetTeamInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// **(Optional, String)** The filter string to apply to the name of the team. Uses equality.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// **(Optional, String)** The search string to apply to the name of the team. Uses contains.
        /// </summary>
        [Input("searchName")]
        public Input<string>? SearchName { get; set; }

        public GetTeamInvokeArgs()
        {
        }
        public static new GetTeamInvokeArgs Empty => new GetTeamInvokeArgs();
    }


    [OutputType]
    public sealed class GetTeamResult
    {
        /// <summary>
        /// **(String)** The description of the team.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// **(Boolean)** Whether the team is predefined.
        /// </summary>
        public readonly bool IsPredefined;
        public readonly string? Name;
        public readonly string? SearchName;

        [OutputConstructor]
        private GetTeamResult(
            string description,

            string id,

            bool isPredefined,

            string? name,

            string? searchName)
        {
            Description = description;
            Id = id;
            IsPredefined = isPredefined;
            Name = name;
            SearchName = searchName;
        }
    }
}
