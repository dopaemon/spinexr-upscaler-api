#!/bin/bash
BASE_URL="http://localhost:8080"

# Test Upload
echo "Testing Upload..."
UPLOAD_RESPONSE=$(curl -s -X POST "$BASE_URL/upload" -F "image=@assets/anh-1.dicom")
echo "Upload Response: $UPLOAD_RESPONSE"
UUID=$(echo $UPLOAD_RESPONSE | jq -r '.uuid')

if [ "$UUID" == "null" ] || [ -z "$UUID" ]; then
    echo "Upload failed or UUID not found"
    exit 1
fi

echo "Extracted UUID: $UUID"

# Test Diagnosis
echo "Testing Diagnosis..."
curl -X GET "$BASE_URL/diagnosis/$UUID"

# Test Download
echo "Testing Download..."
curl -X GET "$BASE_URL/download/$UUID.png" -o downloaded.png
