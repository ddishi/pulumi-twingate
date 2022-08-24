// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getTwingateGroups(args?: GetTwingateGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("twingate:index/getTwingateGroups:getTwingateGroups", {
        "groups": args.groups,
        "isActive": args.isActive,
        "name": args.name,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateGroups.
 */
export interface GetTwingateGroupsArgs {
    groups?: inputs.GetTwingateGroupsGroup[];
    isActive?: boolean;
    name?: string;
    type?: string;
}

/**
 * A collection of values returned by getTwingateGroups.
 */
export interface GetTwingateGroupsResult {
    readonly groups?: outputs.GetTwingateGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isActive?: boolean;
    readonly name?: string;
    readonly type?: string;
}

export function getTwingateGroupsOutput(args?: GetTwingateGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTwingateGroupsResult> {
    return pulumi.output(args).apply(a => getTwingateGroups(a, opts))
}

/**
 * A collection of arguments for invoking getTwingateGroups.
 */
export interface GetTwingateGroupsOutputArgs {
    groups?: pulumi.Input<pulumi.Input<inputs.GetTwingateGroupsGroupArgs>[]>;
    isActive?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
