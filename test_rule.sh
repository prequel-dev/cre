#!/bin/sh

# Test the retry loop detection rule
echo "Testing retry loop detection rule..."
cat /app/rules/cre-2025-0166/test.log | /go/bin/preq -r /app/rules/cre-2025-0166/retry-loop-detection.yaml -d

echo "\nTesting with false positive log..."
cat /app/rules/cre-2025-0166/test-fp.log | /go/bin/preq -r /app/rules/cre-2025-0166/retry-loop-detection.yaml -d
