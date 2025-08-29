#!/bin/bash
# Test the AutoGPT infinite loop detection rule

# Set the path to the preq executable
PREQ="./bin/preq-0.1.31-windows_amd64/preq.exe"

# Test with the sample log file
echo "Testing AutoGPT infinite loop detection rule..."
cat tests/autogpt-infinite-loop-test.log | "$PREQ" --rules rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml --action analyze

echo -e "\nIf you don't see any output, try running with --debug flag:"
echo "cat tests/autogpt-infinite-loop-test.log | "$PREQ" --rules rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml --action analyze --debug"
