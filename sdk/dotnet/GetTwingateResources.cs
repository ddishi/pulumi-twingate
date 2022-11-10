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
    public static class GetTwingateResources
    {
        public static Task<GetTwingateResourcesResult> InvokeAsync(GetTwingateResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTwingateResourcesResult>("twingate:index/getTwingateResources:getTwingateResources", args ?? new GetTwingateResourcesArgs(), options.WithDefaults());

        public static Output<GetTwingateResourcesResult> Invoke(GetTwingateResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTwingateResourcesResult>("twingate:index/getTwingateResources:getTwingateResources", args ?? new GetTwingateResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTwingateResourcesArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("resources")]
        private List<Inputs.GetTwingateResourcesResourceArgs>? _resources;
        public List<Inputs.GetTwingateResourcesResourceArgs> Resources
        {
            get => _resources ?? (_resources = new List<Inputs.GetTwingateResourcesResourceArgs>());
            set => _resources = value;
        }

        public GetTwingateResourcesArgs()
        {
        }
        public static new GetTwingateResourcesArgs Empty => new GetTwingateResourcesArgs();
    }

    public sealed class GetTwingateResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("resources")]
        private InputList<Inputs.GetTwingateResourcesResourceInputArgs>? _resources;
        public InputList<Inputs.GetTwingateResourcesResourceInputArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.GetTwingateResourcesResourceInputArgs>());
            set => _resources = value;
        }

        public GetTwingateResourcesInvokeArgs()
        {
        }
        public static new GetTwingateResourcesInvokeArgs Empty => new GetTwingateResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class GetTwingateResourcesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetTwingateResourcesResourceResult> Resources;

        [OutputConstructor]
        private GetTwingateResourcesResult(
            string id,

            string name,

            ImmutableArray<Outputs.GetTwingateResourcesResourceResult> resources)
        {
            Id = id;
            Name = name;
            Resources = resources;
        }
    }
}
