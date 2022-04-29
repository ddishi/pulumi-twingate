// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./provider";
export * from "./twingateConnector";
export * from "./twingateConnectorTokens";
export * from "./twingateRemoteNetwork";
export * from "./twingateResource";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { TwingateConnector } from "./twingateConnector";
import { TwingateConnectorTokens } from "./twingateConnectorTokens";
import { TwingateRemoteNetwork } from "./twingateRemoteNetwork";
import { TwingateResource } from "./twingateResource";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "twingate:index/twingateConnector:TwingateConnector":
                return new TwingateConnector(name, <any>undefined, { urn })
            case "twingate:index/twingateConnectorTokens:TwingateConnectorTokens":
                return new TwingateConnectorTokens(name, <any>undefined, { urn })
            case "twingate:index/twingateRemoteNetwork:TwingateRemoteNetwork":
                return new TwingateRemoteNetwork(name, <any>undefined, { urn })
            case "twingate:index/twingateResource:TwingateResource":
                return new TwingateResource(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("twingate", "index/twingateConnector", _module)
pulumi.runtime.registerResourceModule("twingate", "index/twingateConnectorTokens", _module)
pulumi.runtime.registerResourceModule("twingate", "index/twingateRemoteNetwork", _module)
pulumi.runtime.registerResourceModule("twingate", "index/twingateResource", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("twingate", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:twingate") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});