<a href="https://www.nxpod.khulnasoft.com">
  <img width="1500" alt="Nxpod banner image" src="https://github.com/nxpkg/nxpod/assets/55068936/ff393a65-a4f3-4997-a066-10337562cc34">
</a>

<br>
<br>



<div align="center" style="flex-direction: row;">
  <a href="https://www.nxpod.khulnasoft.com#gh-light-mode-only">
    <img src="https://github.com/nxpkg/nxpod/assets/55068936/01a00b23-e1f5-4650-a629-89db8e300708" style="width: 256px">
  </a>
  <a href="https://www.nxpod.khulnasoft.com#gh-dark-mode-only">
    <img src="https://github.com/nxpkg/nxpod/assets/55068936/ff437ec6-adda-4814-9e92-fff44cfd00ad" style="width: 256px">
  </a>
</div>


[Nxpod](https://www.nxpod.khulnasoft.com)’s developer platform provides on-demand, pre-configured environments that automatically integrate into any tool, library, or dependency required for creating software. Nxpod workspaces are the fastest and most secure way to ship software and are as easy as adding a .nxpod.yml file to the root of any repository.

📄 [Read Cloud Development Environment white paper](https://www.nxpod.khulnasoft.com/whitepaper/cde)


<div align="center" style="flex-direction: row;">
  <a href="https://www.nxpod.khulnasoft.com#gh-light-mode-only">
    <img src="https://github.com/nxpkg/nxpod/assets/55068936/84beb5da-fa5b-4966-a454-8f5b2607a8ed">
  </a>
  <a href="https://www.nxpod.khulnasoft.com#gh-dark-mode-only">
    <img src="https://github.com/nxpkg/nxpod/assets/55068936/3912a60a-65a9-45f2-b151-93621ac8bf53">
  </a>
</div>



## Features

- **Dev environments as code** - Nxpod is like infrastructure-as-code, but for your development environment. Nxpod defines your editor extensions and requires dependencies in a declarative [`.nxpod.yml` configuration](https://www.nxpod.khulnasoft.com/docs/introduction/nxpod-tutorial/2-configure-your-nxpod-yml). Spinning up dev environments is easily repeatable and reproducible empowering you to automate, version-control, and share dev environments across your team.
- [**Prebuilt dev environments**](https://www.nxpod.khulnasoft.com/docs/configure/projects/prebuilds) - Nxpod continuously prebuilds all your git branches similar to a CI server. Control how Nxpod pre-configures and initializes environments before you even start a workspace through tasks commands in your .nxpod.yml. No more watching apt-get or npm install again. 
- [**Secure**](https://www.nxpod.khulnasoft.com/security) - Each Nxpod workspace or prebuild runs on a secured single-use container providing fast startup times without compromising on security. Nxpod generates SLSA level 1 compliant provenance. Nxpod is also GDPR and SOC2 compliant. And, of course, Nxpod is open-source and available for review by everyone.
- **Workspaces based on Docker** - Nxpod instantly starts a container in the cloud based on an (optional) [Docker image](https://www.nxpod.khulnasoft.com/docs/config-docker/). If you’re already using Docker, you can easily re-use your Docker file. 
- **GitLab, GitHub, and Bitbucket integration** - Nxpod seamlessly [integrates](https://www.nxpod.khulnasoft.com/docs/configure/authentication) into your workflow and works with all major Git hosting platforms including GitHub, GitLab, and Bitbucket.
- **Integrated code reviews** - with Nxpod you can do native code reviews on any PR/MR. No need to switch contexts anymore and clutter your local machine with your colleagues' PR/MR.
- **Collaboration** - invite team members to your dev environment or snapshot of any state of your dev environment to share it with your team asynchronously.
**Professional & customizable developer experience** - a Nxpod workspace gives you the same capabilities as your Linux machine - pre-configured and optimized for your development workflow. Install any [VS Code extension](https://www.nxpod.khulnasoft.com/docs/references/ides-and-editors/vscode-extensions) with one click on a user and/or team level. You can also bring your [dotfiles](https://www.nxpod.khulnasoft.com/docs/configure/user-settings/dotfiles#dotfiles) and customize your dev environment as you like.


## Getting Started

- **Browser**: 
    - Using Nxpod dashboard [nxpod.khulnasoft.com/new](https://nxpod.khulnasoft.com/new).
    - Add `nxpod.khulnasoft.com/# `as a prefix to any of your GitHub/ GitLab/ Bitbucket repository, like [this](https://nxpod.khulnasoft.com/#https://github.com/nxpkg/template-typescript-react)
- **CLI**: You can also [install Nxpod CLI](https://www.nxpod.khulnasoft.com/docs/references/nxpod-cli#installation) and create your first workspace directly from your terminal :)


## Documentation

All documentation can be found on [www.nxpod.khulnasoft.com/docs](https://www.nxpod.khulnasoft.com/docs). For example, see [Nxpod tutorial](https://www.nxpod.khulnasoft.com/docs/introduction/nxpod-tutorial) and check the following helpful resources:
  - [Workspace Lifecycle](https://www.nxpod.khulnasoft.com/docs/configure/workspaces/workspace-lifecycle)
  - [Configure repositories](https://www.nxpod.khulnasoft.com/docs/configure/repositories)
  - [Organizations](https://www.nxpod.khulnasoft.com/docs/configure/orgs)
  - [IDE & Editors support](https://www.nxpod.khulnasoft.com/docs/references/ides-and-editors)
  - [Video screencasts](https://www.nxpod.khulnasoft.com/screencasts)
  - [Nxpod samples](https://github.com/nxpod-samples)


## Questions

For questions and support please use [Nxpod community Discord](https://www.nxpod.khulnasoft.com/chat). Join the conversation, and connect with other community members. 💬
You can also follow [@nxpod](https://twitter.com/nxpod) for announcements and updates from our team.

For enterprise deployment and customized solutions, please explore our [**Enterprise offerings**](https://www.nxpod.khulnasoft.com/contact/enterprise-self-serve) to get started with a setup that meets your organization's needs.

## Issues

The issue tracker is used for tracking bug reports and feature requests for the Nxpod open source project as well as planning current and future development efforts. 🗺️

You can upvote popular feature requests or create a new one.


## Related Projects

During the development of Nxpod, we also developed some of our infrastructure toolings to make development easier and more efficient. To this end, we've developed many open-source projects including:
- [Workspace images](https://github.com/nxpkg/workspace-images): Ready to use docker images for Nxpod workspaces
- [OpenVSCode Server](https://github.com/nxpkg/openvscode-server): Run the latest VS Code on a remote machine accessed through a browser
- [Nxpod browser extension](https://github.com/nxpkg/browser-extension): It adds a Nxpod button to the configured GitLab, GitHub and Bitbucket installations
- [Leeway](https://github.com/nxpkg/leeway) - A heavily caching build system
- [Dazzle](https://github.com/nxpkg/dazzle) - An experimental Docker image builder
- [Werft](https://github.com/csweichel/werft) - A Kubernetes native CI system

## Code of Conduct

We want to create a welcoming environment for everyone interested in contributing to Nxpod or participating in discussions with the Nxpod community.
This project has adopted the [Contributor Covenant Code of Conduct](https://github.com/nxpkg/.github/blob/main/CODE_OF_CONDUCT.md), [version 2.0](https://www.contributor-covenant.org/version/2/0/code_of_conduct/).
