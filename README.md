# github-check-prod-releases

This repo is a small tool that gives insight into how far behind our production deployments are getting.

It simply grabs a list of all of our repositories from GitHub, then asks for a diff between the `prod` tag and master. GitHub will return a number of commits that "prod" is behind, which is the number of releases (as we do squash merges). It will ignore any repos that don't have a `prod` tag.

Even if you're not going squash merges, this will give the number of commits unreleased, which if above 0 you probably need to deploy!

## Prerequisites

1. Have installed GoLang
2. Tag your releases automatically in your Continous Integration environment with "prod" when deployed.

## Run me

Ensure you've setup the GITHUB_TOKEN, and GITHUB_ORG in your environment, then go for it:

```
export GITHUB_TOKEN=xxx
export GITHUB_ORG=xxx
go build && ./github-check-prod-releases
```

* GITHUB_TOKEN is a token that has read access to your private repos.
* GITHUB_ORG is the name of your organisation that you want to scan. It's the name that appears before the slash in a repo name e.g. `OrgName/repo-name`, you would use `OrgName`

## Success

Successful output looks like this, but you might want to go and release some stuff into production...

```
repo-name-1 is behind by 1 releases
repo-name-2 is behind by 1 releases
repo-name-3 is behind by 2 releases
repo-name-4 is behind by 6 releases
```