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
    /// Represents a member of a team
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
    ///     var team = new Konnect.Team("team", new()
    ///     {
    ///         Description = "testing",
    ///     });
    /// 
    ///     var user = Konnect.GetUser.Invoke(new()
    ///     {
    ///         SearchFullName = "Joe",
    ///     });
    /// 
    ///     var example = new Konnect.TeamUser("example", new()
    ///     {
    ///         TeamId = team.Id,
    ///         UserId = user.Apply(getUserResult =&gt; getUserResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Team users can be imported using a proper value of `id` as described above
    /// </summary>
    [KonnectResourceType("konnect:index/teamUser:TeamUser")]
    public partial class TeamUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// **(Required, ForceNew, String)** The id of the team.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;

        /// <summary>
        /// **(Required, ForceNew, String)** The id of the user.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a TeamUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TeamUser(string name, TeamUserArgs args, CustomResourceOptions? options = null)
            : base("konnect:index/teamUser:TeamUser", name, args ?? new TeamUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TeamUser(string name, Input<string> id, TeamUserState? state = null, CustomResourceOptions? options = null)
            : base("konnect:index/teamUser:TeamUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TeamUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TeamUser Get(string name, Input<string> id, TeamUserState? state = null, CustomResourceOptions? options = null)
        {
            return new TeamUser(name, id, state, options);
        }
    }

    public sealed class TeamUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(Required, ForceNew, String)** The id of the team.
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        /// <summary>
        /// **(Required, ForceNew, String)** The id of the user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public TeamUserArgs()
        {
        }
        public static new TeamUserArgs Empty => new TeamUserArgs();
    }

    public sealed class TeamUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **(Required, ForceNew, String)** The id of the team.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// **(Required, ForceNew, String)** The id of the user.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public TeamUserState()
        {
        }
        public static new TeamUserState Empty => new TeamUserState();
    }
}