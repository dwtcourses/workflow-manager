package dynamodb

import (
	"bytes"
	"context"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/Clever/workflow-manager/gen-go/server/db"
	"github.com/Clever/workflow-manager/gen-go/server/db/tests"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestDynamoDBStore(t *testing.T) {
	// spin up dynamodb local
	testCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(testCtx, "./dynamodb-local.sh")
	var ddbLocalOutput bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &ddbLocalOutput)
	cmd.Stderr = io.MultiWriter(os.Stderr, &ddbLocalOutput)
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// wait for dynamodb local to output the correct startup log
	for !strings.Contains(ddbLocalOutput.String(), "Initializing DynamoDB Local with the following configuration") {
		time.Sleep(1 * time.Second)
	}

	dynamoDBAPI := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String("doesntmatter"),
			Endpoint:    aws.String("http://localhost:8002" /* default dynamodb-local port */),
			Credentials: credentials.NewStaticCredentials("id", "secret", "token"),
		},
	})))

	tests.RunDBTests(t, func() db.Interface {
		prefix := "automated-testing"
		listTablesOutput, err := dynamoDBAPI.ListTablesWithContext(testCtx, &dynamodb.ListTablesInput{})
		if err != nil {
			t.Fatal(err)
		}
		for _, tableName := range listTablesOutput.TableNames {
			if strings.HasPrefix(*tableName, prefix) {
				dynamoDBAPI.DeleteTableWithContext(testCtx, &dynamodb.DeleteTableInput{
					TableName: tableName,
				})
			}
		}
		d, err := New(Config{
			DynamoDBAPI:               dynamoDBAPI,
			DefaultPrefix:             prefix,
			DefaultReadCapacityUnits:  10,
			DefaultWriteCapacityUnits: 10,
		})
		if err != nil {
			t.Fatal(err)
		}
		if err := d.CreateTables(testCtx); err != nil {
			t.Fatal(err)
		}
		return d
	})
}
