// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DpCertificate extends pulumi.CustomResource {
    /**
     * Get an existing DpCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DpCertificateState, opts?: pulumi.CustomResourceOptions): DpCertificate {
        return new DpCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'konnect:index/dpCertificate:DpCertificate';

    /**
     * Returns true if the given object is an instance of DpCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DpCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DpCertificate.__pulumiType;
    }

    public readonly cert!: pulumi.Output<string>;
    public readonly controlPlaneId!: pulumi.Output<string>;

    /**
     * Create a DpCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DpCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DpCertificateArgs | DpCertificateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DpCertificateState | undefined;
            resourceInputs["cert"] = state ? state.cert : undefined;
            resourceInputs["controlPlaneId"] = state ? state.controlPlaneId : undefined;
        } else {
            const args = argsOrState as DpCertificateArgs | undefined;
            if ((!args || args.cert === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cert'");
            }
            if ((!args || args.controlPlaneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'controlPlaneId'");
            }
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["controlPlaneId"] = args ? args.controlPlaneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DpCertificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DpCertificate resources.
 */
export interface DpCertificateState {
    cert?: pulumi.Input<string>;
    controlPlaneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DpCertificate resource.
 */
export interface DpCertificateArgs {
    cert: pulumi.Input<string>;
    controlPlaneId: pulumi.Input<string>;
}