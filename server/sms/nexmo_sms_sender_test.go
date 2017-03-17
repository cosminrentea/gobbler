package sms

import (
	"testing"
	"time"

	"github.com/cosminrentea/gobbler/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_HttpClientRecreation(t *testing.T) {
	defer testutil.EnableDebugForMethod()()
	a := assert.New(t)

	port := createRandomPort(7000, 8000)
	URL = "http://127.0.0.1" + port

	sender := createNexmoSender(t)
	msg := encodeProtocolMessage(t, 2)

	err := sender.Send(&msg)
	time.Sleep(3 * timeInterval)

	a.Equal(ErrRetryFailed, err)
}

func TestNexmoSender_SendWithError(t *testing.T) {
	defer testutil.EnableDebugForMethod()
	RequestTimeout = time.Second
	a := assert.New(t)
	sender, err := NewNexmoSender(KEY, SECRET)
	a.NoError(err)

	msg := encodeProtocolMessage(t, 0)

	err = sender.Send(&msg)
	time.Sleep(3 * timeInterval)
	a.Error(err)
	a.Equal(ErrRetryFailed, err)
}