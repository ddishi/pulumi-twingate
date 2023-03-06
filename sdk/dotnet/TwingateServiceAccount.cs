// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TwingateLabs.Twingate
{
    [TwingateResourceType("twingate:index/twingateServiceAccount:TwingateServiceAccount")]
    public partial class TwingateServiceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Service Account in Twingate
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a TwingateServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TwingateServiceAccount(string name, TwingateServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("twingate:index/twingateServiceAccount:TwingateServiceAccount", name, args ?? new TwingateServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TwingateServiceAccount(string name, Input<string> id, TwingateServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("twingate:index/twingateServiceAccount:TwingateServiceAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/Twingate-Labs/pulumi-twingate/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TwingateServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TwingateServiceAccount Get(string name, Input<string> id, TwingateServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new TwingateServiceAccount(name, id, state, options);
        }
    }

    public sealed class TwingateServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Service Account in Twingate
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public TwingateServiceAccountArgs()
        {
        }
        public static new TwingateServiceAccountArgs Empty => new TwingateServiceAccountArgs();
    }

    public sealed class TwingateServiceAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Service Account in Twingate
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TwingateServiceAccountState()
        {
        }
        public static new TwingateServiceAccountState Empty => new TwingateServiceAccountState();
    }
}