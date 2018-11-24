#!/bin/bash

# chmod +x Comp.sh / CGO_ENABLED=0 static binary

echo "Build the binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Create a ZIP file"
zip deployment.zip main

echo "Cleaning up"
rm main

# $ aws s3 mb s3://msundt.foodie --region us-east-1
# make sure the input yaml has correct s3 bucket

aws cloudformation package \
   --template-file foodie.yaml \
   --output-template-file foodie-deploy.yaml \
   --s3-bucket msundt.foodie \
   --region us-east-1

aws cloudformation deploy \
 --template-file foodie-deploy.yaml \
  --stack-name foodie \
  --capabilities CAPABILITY_IAM \
  --region us-east-1