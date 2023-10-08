# Automate Release Workflow

## Key Tasks

- **Generate Branch Tags**: Automatically create tags for a specified branch (e.g., `certif-prod`) when a commit is made on that branch, signifying changes to the application's release version.

- **Automate Releases**: Generate releases based on tagged commits, simplifying version management and changelog generation.

## Ressource used
- [github-action-get-previous-tag](https://github.com/WyriHaximus/github-action-get-previous-tag)
- [github-action-next-semvers](https://github.com/WyriHaximus/github-action-next-semvers)
- [github-action-create-milestone](https://github.com/WyriHaximus/github-action-create-milestone)
- [action-gh-release](https://github.com/softprops/action-gh-release)

## Permissions

This Action requires the following permissions on the [GitHub token permissions](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token):

permissions:

    contents: write

When used with `discussion_category_name`, additional permission is needed:

permissions:

    contents: write
    discussions: write
