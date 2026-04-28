Look at open PRs on this GitHub repository (use `gh pr list`). Find all PRs matching either:
- Title contains "Update Konflux references"
- Title is a UBI base image bump (e.g. contains "Update registry.access.redhat.com/ubi")

For each matching PR:

1. Add a `/ok-to-test` comment if the `ok-to-test` label is not present on the PR.
2. Schedule a check in 1 hour. When the check fires:
   a. Look at CI status using `gh pr checks`.
   b. If all checks passed (green):
      - Add a single comment containing all three lines:
        ```
        /retitle NO-JIRA: <original PR title>
        /lgtm
        /approve
        ```
   c. If checks are still running or have failed, report the status to the user. Do NOT add `/lgtm`, `/approve`, or `/retitle`.
