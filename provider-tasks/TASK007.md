# TASK007 - Refresh stale default models

## Goal
Implement a capability to refresh stale default models in the provider to maintain hygiene and ensure the latest default models are used.

## Background
Default models may become stale over time due to updates or improvements. This task ensures that the system refreshes these models automatically or on-demand to keep them current.

## Requirements
- Detect when default models are stale.
- Refresh or reload default models accordingly.
- Provide tests to validate the refresh mechanism.
- Update documentation and changelog if user-visible changes occur.

## Acceptance Criteria
- [ ] The system refreshes stale default models automatically or via an explicit trigger.
- [ ] Tests cover the detection and refresh of stale default models.
- [ ] Documentation or changelog updated if applicable.
- [ ] PR references this task and GitHub issue #8.

## Notes
- Keep the change minimal and focused.
- This is a good first issue for community contributors.
- Avoid scaffolding or empty PRs.
