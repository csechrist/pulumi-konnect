// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Represents a role assigned to a user to access a given entity
 *
 * ## Import
 *
 * User roles can be imported using a proper value of `id` as described above
 */
export class UserRole extends pulumi.CustomResource {
    /**
     * Get an existing UserRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserRoleState, opts?: pulumi.CustomResourceOptions): UserRole {
        return new UserRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'konnect:index/userRole:UserRole';

    /**
     * Returns true if the given object is an instance of UserRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserRole.__pulumiType;
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
     * **(Required, ForceNew, String)** The id of the user assigned the role
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a UserRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserRoleArgs | UserRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserRoleState | undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["entityRegion"] = state ? state.entityRegion : undefined;
            resourceInputs["entityTypeDisplayName"] = state ? state.entityTypeDisplayName : undefined;
            resourceInputs["roleDisplayName"] = state ? state.roleDisplayName : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserRoleArgs | undefined;
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
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["entityRegion"] = args ? args.entityRegion : undefined;
            resourceInputs["entityTypeDisplayName"] = args ? args.entityTypeDisplayName : undefined;
            resourceInputs["roleDisplayName"] = args ? args.roleDisplayName : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserRole resources.
 */
export interface UserRoleState {
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
     * **(Required, ForceNew, String)** The id of the user assigned the role
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserRole resource.
 */
export interface UserRoleArgs {
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
     * **(Required, ForceNew, String)** The id of the user assigned the role
     */
    userId: pulumi.Input<string>;
}
