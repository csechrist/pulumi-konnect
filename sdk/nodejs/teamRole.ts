// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Represents a role assigned to a team to access a given entity
 *
 * ## Import
 *
 * Team roles can be imported using a proper value of `id` as described above
 */
export class TeamRole extends pulumi.CustomResource {
    /**
     * Get an existing TeamRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamRoleState, opts?: pulumi.CustomResourceOptions): TeamRole {
        return new TeamRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'konnect:index/teamRole:TeamRole';

    /**
     * Returns true if the given object is an instance of TeamRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TeamRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TeamRole.__pulumiType;
    }

    /**
     * **(Required, ForceNew, String)** The id of the entity for which the role applies.
     */
    public readonly entityId!: pulumi.Output<string>;
    /**
     * **(Required, ForceNew, String)** The region of the entity for which the role applies.
     */
    public readonly entityRegion!: pulumi.Output<string>;
    /**
     * **(Required, ForceNew, String)** The display name of the entity type, like `Control Planes` or `Services`.
     */
    public readonly entityTypeDisplayName!: pulumi.Output<string>;
    /**
     * **(Required, ForceNew, String)** The display name of the role.
     */
    public readonly roleDisplayName!: pulumi.Output<string>;
    /**
     * **(Required, ForceNew, String)** The id of the team assigned the role
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a TeamRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TeamRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamRoleArgs | TeamRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamRoleState | undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["entityRegion"] = state ? state.entityRegion : undefined;
            resourceInputs["entityTypeDisplayName"] = state ? state.entityTypeDisplayName : undefined;
            resourceInputs["roleDisplayName"] = state ? state.roleDisplayName : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as TeamRoleArgs | undefined;
            if ((!args || args.entityId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityId'");
            }
            if ((!args || args.entityRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityRegion'");
            }
            if ((!args || args.entityTypeDisplayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityTypeDisplayName'");
            }
            if ((!args || args.roleDisplayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleDisplayName'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["entityRegion"] = args ? args.entityRegion : undefined;
            resourceInputs["entityTypeDisplayName"] = args ? args.entityTypeDisplayName : undefined;
            resourceInputs["roleDisplayName"] = args ? args.roleDisplayName : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TeamRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TeamRole resources.
 */
export interface TeamRoleState {
    /**
     * **(Required, ForceNew, String)** The id of the entity for which the role applies.
     */
    entityId?: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The region of the entity for which the role applies.
     */
    entityRegion?: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The display name of the entity type, like `Control Planes` or `Services`.
     */
    entityTypeDisplayName?: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The display name of the role.
     */
    roleDisplayName?: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The id of the team assigned the role
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TeamRole resource.
 */
export interface TeamRoleArgs {
    /**
     * **(Required, ForceNew, String)** The id of the entity for which the role applies.
     */
    entityId: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The region of the entity for which the role applies.
     */
    entityRegion: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The display name of the entity type, like `Control Planes` or `Services`.
     */
    entityTypeDisplayName: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The display name of the role.
     */
    roleDisplayName: pulumi.Input<string>;
    /**
     * **(Required, ForceNew, String)** The id of the team assigned the role
     */
    teamId: pulumi.Input<string>;
}
