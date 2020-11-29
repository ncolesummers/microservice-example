package sqs

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/ncolesummers/microservice-example/lib/msgqueue"
)

type SQSEmitter struct {
	sqsSvc   *sqs.SQS
	QueueUrl *string
}

func NewSQSEventEmitter(s *session.Session, queueName string) (emitter msgqueue.EventEmitter, err error) {
	if s == nil {
		s, err = session.NewSession()
		if err != nil {
			return nil, err
		}
	}

	svc := sqs.New(s)

	QUResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		return
	}

	emitter = &SQSEmitter{
		sqsSvc:   svc,
		QueueUrl: QUResult.QueueUrl,
	}
	return emitter, err
}

func (sqsEmit *SQSEmitter) Emit(event msgqueue.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	_, err = sqsEmit.sqsSvc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"event_name": &sqs.MessageAttributeValue{
				DataType:    aws.String("string"),
				StringValue: aws.String(event.EventName()),
			},
		},
		MessageBody: aws.String(string(data)),
		QueueUrl:    sqsEmit.QueueUrl,
	})
	return err
}
