package awssession

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/clok/kemba"
)

var (
	k = kemba.New("awssession:new")
)

// New establishes a Session to interact with AWS.
//
// Credit to https://github.com/cyberark/summon-aws-secrets/blob/master/main.go for initial implementation.
//
// All clients require a Session. The Session provides the client with
// shared configuration such as region, endpoint, and credentials. A
// Session should be shared where possible to take advantage of
// configuration and credential caching. See the session package for
// more information.
//
// AWS Go SDK does not currently support automatic fetching of region from ec2metadata.
// If the region could not be found in an environment variable or a shared config file,
// create metaSession to fetch the ec2 instance region and pass to the regular Session.
//
// If running on an EC2 instance, the metaClient will be available and we can set the region to match the instance
// If not on an EC2 instance, the region will remain blank and AWS returns a "MissingRegion: ..." error
func New() (*session.Session, error) {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		k.Printf("error during session creation with options: %e", err)
		return nil, err
	}

	// AWS Go SDK does not currently support automatic fetching of region from ec2metadata.
	// If the region could not be found in an environment variable or a shared config file,
	// create metaSession to fetch the ec2 instance region and pass to the regular Session.
	if *sess.Config.Region == "" {
		metaSession, err := session.NewSession()
		if err != nil {
			k.Printf("error during session creation with ec2metadata: %e", err)
			return nil, err
		}

		metaClient := ec2metadata.New(metaSession)
		// If running on an EC2 instance, the metaClient will be available and we can set the region to match the instance
		// If not on an EC2 instance, the region will remain blank and AWS returns a "MissingRegion: ..." error
		if metaClient.Available() {
			if region, err := metaClient.Region(); err == nil {
				k.Printf("region found: %s", region)
				sess.Config.Region = aws.String(region)
			} else {
				k.Printf("error during session creation with ec2metadata: %e", err)
				return nil, err
			}
		}
	}
	k.Println("successfully create AWS session")
	return sess, nil
}
