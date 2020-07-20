# awssession - Super simple AWS session generator.

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