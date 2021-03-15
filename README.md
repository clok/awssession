# awssession - Super simple AWS session generator.

[![Go Report Card](https://goreportcard.com/badge/clok/awssession)](https://goreportcard.com/report/clok/awssession) [![Coverage Status](https://coveralls.io/repos/github/clok/awssession/badge.svg?branch=master)](https://coveralls.io/github/clok/awssession?branch=master) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/clok/awssession?tab=overview)

### Example

Simple example of an S3 Get.

```go
package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/clok/awssession"
    "os"
)

func main() {
    var dstFile *os.File
    dstFile, _ = os.Create(dest)
    
    defer dstFile.Close()

    sess, _ := awssession.New()
    downloader := s3manager.NewDownloader(sess)
    
    _, _ = downloader.Download(dstFile, &s3.GetObjectInput{
        Bucket: aws.String("s3://aBucket"),
        Key:    aws.String("an/object.json"),
    })

    return
}
```

## Development

1. Fork the [clok/awssession](https://github.com/clok/awssession) repo
1. Use `go >= 1.16`
1. Branch & Code
1. Run linters :broom: `golangci-lint run`
    - The project uses [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
1. Commit with a Conventional Commit
1. Open a PR