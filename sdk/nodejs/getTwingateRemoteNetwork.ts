// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getTwingateRemoteNetwork(args?: GetTwingateRemoteNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateRemoteNetworkResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("twingate:index/getTwingateRemoteNetwork:getTwingateRemoteNetwork", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateRemoteNetwork.
 */
export interface GetTwingateRemoteNetworkArgs {
    id?: string;
    name?: string;
}

/**
 * A collection of values returned by getTwingateRemoteNetwork.
 */
export interface GetTwingateRemoteNetworkResult {
    readonly id?: string;
    readonly name?: string;
}

export function getTwingateRemoteNetworkOutput(args?: GetTwingateRemoteNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTwingateRemoteNetworkResult> {
    return pulumi.output(args).apply(a => getTwingateRemoteNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getTwingateRemoteNetwork.
 */
export interface GetTwingateRemoteNetworkOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}
