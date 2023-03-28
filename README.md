# launchdarkly

[LaunchDarkly](https://launchdarkly.com/) is a SaaS platform for developers to manage feature flags. By decoupling feature rollout and code deployment, LaunchDarkly enables developers to test their code live in production, gradually release features to groups of users, and manage flags throughout their entire lifecycle.

Use this integration to tell Okteto to include a LaunchDarkly environment in your Cloud Development Environment.

To use it:
1. Add the following secrets to your Okteto instance: 

    - `LAUNCHDARKLY_ACCESS_TOKEN`: A token with read/write access to your LaunchDarkly project.
    - `LAUNCHDARKLY_PROJECT_KEY`: The key of your LaunchDarkly project.
    - `LAUNCHDARKLY_SOURCE`: The key of the environment you want to clone (optional).

2. Update your manifest with the configuration below:
    ```yaml
    ...
    dependencies:
        launchdarkly:
            - https://github.com/okteto/launchdarkly
    ...
    ```

