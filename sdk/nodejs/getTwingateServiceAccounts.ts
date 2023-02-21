// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Service Accounts offer a way to provide programmatic, centrally-controlled, and consistent access controls.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = pulumi.output(twingate.getTwingateServiceAccounts({
 *     name: "<your service account's name>",
 * }));
 * ```
 */
export function getTwingateServiceAccounts(args?: GetTwingateServiceAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateServiceAccountsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("twingate:index/getTwingateServiceAccounts:getTwingateServiceAccounts", {
        "name": args.name,
        "serviceAccounts": args.serviceAccounts,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateServiceAccounts.
 */
export interface GetTwingateServiceAccountsArgs {
    /**
     * Filter results by the name of the Service Account.
     */
    name?: string;
    /**
     * List of Service Accounts
     */
    serviceAccounts?: inputs.GetTwingateServiceAccountsServiceAccount[];
}

/**
 * A collection of values returned by getTwingateServiceAccounts.
 */
export interface GetTwingateServiceAccountsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Filter results by the name of the Service Account.
     */
    readonly name?: string;
    /**
     * List of Service Accounts
     */
    readonly serviceAccounts?: outputs.GetTwingateServiceAccountsServiceAccount[];
}

export function getTwingateServiceAccountsOutput(args?: GetTwingateServiceAccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTwingateServiceAccountsResult> {
    return pulumi.output(args).apply(a => getTwingateServiceAccounts(a, opts))
}

/**
 * A collection of arguments for invoking getTwingateServiceAccounts.
 */
export interface GetTwingateServiceAccountsOutputArgs {
    /**
     * Filter results by the name of the Service Account.
     */
    name?: pulumi.Input<string>;
    /**
     * List of Service Accounts
     */
    serviceAccounts?: pulumi.Input<pulumi.Input<inputs.GetTwingateServiceAccountsServiceAccountArgs>[]>;
}
