# Add AutoGPT Infinite Loop Detection Rule (CRE-2025-0071)

## Description
This PR adds a new CRE rule to detect when an AutoGPT agent enters an infinite loop, repeatedly executing the same or similar tasks without making progress.

## Changes
- Added new rule: `rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml`
  - Uses existing `asynchronous-task-problem` category for better compatibility
  - Uses valid tags: `critical-failure`, `cpu-memory-exhaustion`, `performance`
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

