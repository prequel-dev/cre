# PowerShell script to test the AutoGPT infinite loop detection rule

# Path to the preq executable
$preqPath = "C:\Users\HP\cre\bin\preq-0.1.31-windows_amd64\preq.exe"

# Test with the sample log file
Write-Host "Testing AutoGPT infinite loop detection rule..."
Get-Content tests/autogpt-infinite-loop-test.log | & $preqPath --rules rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml --action analyze

# Try with debug output
Write-Host "`nTrying with debug output..."
Get-Content tests/autogpt-infinite-loop-test.log | & $preqPath --rules rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml --action analyze --debug
