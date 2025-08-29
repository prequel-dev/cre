# Add AutoGPT Infinite Loop Detection Rule (CRE-2025-0071)

## Description
This PR adds a new CRE rule to detect when an AutoGPT agent enters an infinite loop, repeatedly executing the same or similar tasks without making progress.

## Changes
- Added new rule: `rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml`
- Added test log file: `tests/autogpt-infinite-loop-test.log`
- Added test YAML: `tests/autogpt-infinite-loop-test.yaml`

## Testing
1. Rule was tested locally using:
   ```bash
   type tests\autogpt-infinite-loop-test.log | preq --rules rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml
   ```
2. Test YAML file added for automated testing

## Impact
Detects when AutoGPT agents get stuck in repetitive execution patterns, helping to prevent resource wastage and unexpected costs.

## Related Issues
- Fixes #[ISSUE_NUMBER]  <!-- If applicable -->

## Checklist
- [x] Rule follows CRE schema
- [x] Test case provided
- [x] Documentation complete
- [x] Pre-commit checks pass
