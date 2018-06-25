package que
import  (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"time"
	"net/http"
)


const (
	sqsPollingFrequencySecs            = 5
	maxNumMessagesToFetch              = 10
	longPollTimeSeconds                = 20
	defaultVisibilityTimeout           = 120
	numSQSFailuresBeforeReregistration = 10
)

func init() {
	// Agent id and signature are mandatory attributes in every SQS message that agent processes.
}

func createSQSClient(opsGenieToken *OpsGenieToken) (*sqs.SQS) {
	creds := credentials.NewStaticCredentials(opsGenieToken.akid, opsGenieToken.secret, opsGenieToken.token)
	awsConfig := aws.NewConfig().WithCredentials(creds).
		WithRegion(REGION).
		WithHTTPClient(http.DefaultClient).
		WithMaxRetries(aws.UseServiceDefaultRetries).
		WithLogger(aws.NewDefaultLogger()).
		WithLogLevel(aws.LogOff).
		WithSleepDelay(time.Sleep)
	return sqs.New(session.New(awsConfig))
}


func readMessage(svc *sqs.SQS, queUrl *string) (*sqs.ReceiveMessageOutput, error) {
	resp, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:queUrl,
		MaxNumberOfMessages:   aws.Int64(maxNumMessagesToFetch),
		VisibilityTimeout:     aws.Int64(defaultVisibilityTimeout),
		WaitTimeSeconds:       aws.Int64(longPollTimeSeconds),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func deleteMessage(svc *sqs.SQS, queUrl *string)  {
	svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:queUrl,
	})
}