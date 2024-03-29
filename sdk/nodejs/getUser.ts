// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Represents a user
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as konnect from "@pulumi/konnect";
 *
 * const example = konnect.getUser({
 *     searchEmail: "@example.com",
 * });
 * ```
 */
export function getUser(args?: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("konnect:index/getUser:getUser", {
        "active": args.active,
        "email": args.email,
        "fullName": args.fullName,
        "searchEmail": args.searchEmail,
        "searchFullName": args.searchFullName,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * **(Optional, Boolean)** The filter flag to apply to the active flag of the user. Uses equality. Default: `true`
     */
    active?: boolean;
    /**
     * **(Optional, String)** The filter string to apply to the email of the user. Uses equality.
     */
    email?: string;
    /**
     * **(Optional, String)** The filter string to apply to the full name of the user. Uses equality.
     */
    fullName?: string;
    /**
     * **(Optional, String)** The search string to apply to the email of the user. Uses contains.
     */
    searchEmail?: string;
    /**
     * **(Optional, String)** The search string to apply to the full name of the user. Uses contains.
     */
    searchFullName?: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    readonly active?: boolean;
    readonly email?: string;
    readonly fullName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * **(String)** The preferred name of the user.
     */
    readonly preferredName: string;
    readonly searchEmail?: string;
    readonly searchFullName?: string;
}
/**
 * Represents a user
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as konnect from "@pulumi/konnect";
 *
 * const example = konnect.getUser({
 *     searchEmail: "@example.com",
 * });
 * ```
 */
export function getUserOutput(args?: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply((a: any) => getUser(a, opts))
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    /**
     * **(Optional, Boolean)** The filter flag to apply to the active flag of the user. Uses equality. Default: `true`
     */
    active?: pulumi.Input<boolean>;
    /**
     * **(Optional, String)** The filter string to apply to the email of the user. Uses equality.
     */
    email?: pulumi.Input<string>;
    /**
     * **(Optional, String)** The filter string to apply to the full name of the user. Uses equality.
     */
    fullName?: pulumi.Input<string>;
    /**
     * **(Optional, String)** The search string to apply to the email of the user. Uses contains.
     */
    searchEmail?: pulumi.Input<string>;
    /**
     * **(Optional, String)** The search string to apply to the full name of the user. Uses contains.
     */
    searchFullName?: pulumi.Input<string>;
}
