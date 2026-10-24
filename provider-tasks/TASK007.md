# TASK007 - Refresh stale default models

## Goal
Refresh stale default models to ensure hygiene and up-to-date baseline for providers.

## Background
Default models can become stale over time due to updates in dependencies, bug fixes, or improvements. This task ensures that the default models are refreshed regularly to maintain quality and correctness.

## Acceptance Criteria
- Default models are refreshed and updated in the primary repo.
- Tests validate that the refreshed models are correctly loaded and used.
- Documentation and CHANGELOG updated if any user-visible changes occur.
- The PR references this task and GitHub issue #3.

## Implementation Notes
- Identify the current default models that are stale.
- Replace or update them with the latest versions.
- Run existing tests to ensure no regressions.
- Add tests if necessary to cover the refresh.
- Keep the change minimal and focused.

## Testing
- Validate that the refreshed models load without errors.
- Confirm that model outputs remain consistent or improve.
- Add snapshot or integration tests if applicable.

## References
- GitHub issue #3
- Provider hygiene guidelines
