##GH

This is a project that basically exists for me to familiarize myself with Go.

It's essentially a port of [git-gh](https://github.com/nlf/git-gh), though with a slightly different interface.

Instead of generating several `git-` style binaries, it builds only one binary `gh` that has its own subcommands.

Implemented so far:

* `gh setup`
* `gh issues [--milestone] [--label] [--state] [--assignee] [--creator] [--mentioned] [--sort] [--direction] [--since]`
