// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class TwingateResource extends pulumi.CustomResource {
    /**
     * Get an existing TwingateResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TwingateResourceState, opts?: pulumi.CustomResourceOptions): TwingateResource {
        return new TwingateResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'twingate:index/twingateResource:TwingateResource';

    /**
     * Returns true if the given object is an instance of TwingateResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TwingateResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TwingateResource.__pulumiType;
    }

    /**
     * The Resource's IP/CIDR or FQDN/DNS zone
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * List of Group IDs that have permission to access the Resource, cannot be generated by Terraform and must be retrieved
     * from the Twingate Admin Console or API
     */
    public readonly groupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the Resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Restrict access to certain protocols and ports. By default or when this argument is not defined, there is no
     * restriction, and all protocols and ports are allowed.
     */
    public readonly protocols!: pulumi.Output<outputs.TwingateResourceProtocols | undefined>;
    /**
     * Remote Network ID where the Resource lives
     */
    public readonly remoteNetworkId!: pulumi.Output<string>;

    /**
     * Create a TwingateResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TwingateResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TwingateResourceArgs | TwingateResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TwingateResourceState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["groupIds"] = state ? state.groupIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocols"] = state ? state.protocols : undefined;
            resourceInputs["remoteNetworkId"] = state ? state.remoteNetworkId : undefined;
        } else {
            const args = argsOrState as TwingateResourceArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.remoteNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteNetworkId'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["groupIds"] = args ? args.groupIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocols"] = args ? args.protocols : undefined;
            resourceInputs["remoteNetworkId"] = args ? args.remoteNetworkId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TwingateResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TwingateResource resources.
 */
export interface TwingateResourceState {
    /**
     * The Resource's IP/CIDR or FQDN/DNS zone
     */
    address?: pulumi.Input<string>;
    /**
     * List of Group IDs that have permission to access the Resource, cannot be generated by Terraform and must be retrieved
     * from the Twingate Admin Console or API
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Resource
     */
    name?: pulumi.Input<string>;
    /**
     * Restrict access to certain protocols and ports. By default or when this argument is not defined, there is no
     * restriction, and all protocols and ports are allowed.
     */
    protocols?: pulumi.Input<inputs.TwingateResourceProtocols>;
    /**
     * Remote Network ID where the Resource lives
     */
    remoteNetworkId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TwingateResource resource.
 */
export interface TwingateResourceArgs {
    /**
     * The Resource's IP/CIDR or FQDN/DNS zone
     */
    address: pulumi.Input<string>;
    /**
     * List of Group IDs that have permission to access the Resource, cannot be generated by Terraform and must be retrieved
     * from the Twingate Admin Console or API
     */
    groupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Resource
     */
    name: pulumi.Input<string>;
    /**
     * Restrict access to certain protocols and ports. By default or when this argument is not defined, there is no
     * restriction, and all protocols and ports are allowed.
     */
    protocols?: pulumi.Input<inputs.TwingateResourceProtocols>;
    /**
     * Remote Network ID where the Resource lives
     */
    remoteNetworkId: pulumi.Input<string>;
}
