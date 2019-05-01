
# KubeDevice

KubeDevice is an extended plugin framework for Kubernetes and offers more flexibility
than using the default device plugins and scheduler extender.
KubeDevice does not fork the upstream Kubernetes core, but does fork off the Kubernetes scheduler
codebase to build a custom scheduler.

In particular, KubeDevice still allows for 
the allocation of devices to return a list of devices, volume mounts, and environment
variables as device plugins do, but in addition gives the device the flexibility to examine the entire pod specification.
This allows a custom scheduler to write in specific allocation instructions into the pod annotation which can be
examined by the device plugin.
For example, the list of exact devices to use can be written by the scheduler into the pod annotations.

# Architecture



# Usage

# Writing your own plugins

# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.microsoft.com.

When you submit a pull request, a CLA-bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., label, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
