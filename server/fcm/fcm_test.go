package fcm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Bogh/gcm"
	"github.com/cosminrentea/gobbler/protocol"
	"github.com/cosminrentea/gobbler/server/connector"
	"github.com/cosminrentea/gobbler/server/kafka"
	"github.com/cosminrentea/gobbler/server/kvstore"
	"github.com/cosminrentea/gobbler/server/router"
	"github.com/cosminrentea/gobbler/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var fullFCMMessage = `{
	"notification": {
		"title": "TEST",
		"body": "notification body",
		"icon": "ic_notification_test_icon",
		"click_action": "estimated_arrival"
	},
	"data": {"field1": "value1", "field2": "value2"}
}`

type mocks struct {
	router    *MockRouter
	store     *MockMessageStore
	gcmSender *MockSender
}

func TestConnector_GetErrorMessageFromFCM(t *testing.T) {
	_, finish := testutil.NewMockCtrl(t)
	defer finish()

	a := assert.New(t)
	fcm, mocks := testFCM(t, true, nil)

	err := fcm.Start()
	a.NoError(err)

	var route *router.Route

	mocks.router.EXPECT().Subscribe(gomock.Any()).Do(func(r *router.Route) (*router.Route, error) {
		a.Equal("/topic", string(r.Path))
		a.Equal("user01", r.Get("user_id"))
		a.Equal("device01", r.Get(deviceTokenKey))
		route = r
		return r, nil
	})

	// put a dummy FCM message with minimum information
	postSubscription(t, fcm, "user01", "device01", "topic")
	time.Sleep(100 * time.Millisecond)
	a.NoError(err)
	a.NotNil(route)

	// expect the route unsubscribed
	mocks.router.EXPECT().Unsubscribe(gomock.Any()).Do(func(route *router.Route) {
		a.Equal("/topic", string(route.Path))
		a.Equal("device01", route.Get(deviceTokenKey))
	})

	// expect the route subscribe with the new canonicalID from replaceSubscriptionWithCanonicalID
	mocks.router.EXPECT().Subscribe(gomock.Any()).Do(func(route *router.Route) {
		a.Equal("/topic", string(route.Path))
		a.Equal("user01", route.Get("user_id"))
		appid := route.Get(deviceTokenKey)
		a.Equal("fcmCanonicalID", appid)
	})
	// mocks.store.EXPECT().MaxMessageID(gomock.Any()).Return(uint64(4), nil)

	response := new(gcm.Response)
	err = json.Unmarshal([]byte(ErrorFCMResponse), response)
	a.NoError(err)
	mocks.gcmSender.EXPECT().Send(gomock.Any()).Return(response, nil)

	// send the message into the subscription route channel
	route.Deliver(&protocol.Message{
		ID:         uint64(4),
		Path:       "/topic",
		Body:       []byte("{id:id}"),
		HeaderJSON: `{"Correlation-Id": "7sdks723ksgqn"}`,
	}, true)

	// wait before closing the FCM connector
	time.Sleep(100 * time.Millisecond)

	err = fcm.Stop()
	a.NoError(err)
}

func TestFCMFormatMessage(t *testing.T) {
	_, finish := testutil.NewMockCtrl(t)
	defer finish()

	a := assert.New(t)

	var subRoute *router.Route

	fcm, mocks := testFCM(t, false, nil)
	fcm.Start()
	defer fcm.Stop()
	time.Sleep(50 * time.Millisecond)

	mocks.router.EXPECT().Subscribe(gomock.Any()).Do(func(route *router.Route) (*router.Route, error) {
		subRoute = route
		return route, nil
	})

	postSubscription(t, fcm, "user01", "device01", "topic")
	time.Sleep(100 * time.Millisecond)

	// send a fully formated GCM message
	m := &protocol.Message{
		Path:       "/topic",
		ID:         1,
		Body:       []byte(fullFCMMessage),
		HeaderJSON: `{"Correlation-Id": "7sdks723ksgqn"}`,
	}

	if !a.NotNil(subRoute) {
		return
	}

	doneC := make(chan bool)

	mocks.gcmSender.EXPECT().Send(gomock.Any()).Do(func(m *gcm.Message) (*gcm.Response, error) {
		a.NotNil(m.Notification)
		a.Equal("TEST", m.Notification.Title)
		a.Equal("notification body", m.Notification.Body)
		a.Equal("ic_notification_test_icon", m.Notification.Icon)
		a.Equal("estimated_arrival", m.Notification.ClickAction)

		a.NotNil(m.Data)
		if a.Contains(m.Data, "field1") {
			a.Equal("value1", m.Data["field1"])
		}
		if a.Contains(m.Data, "field2") {
			a.Equal("value2", m.Data["field2"])
		}

		doneC <- true
		return nil, nil
	}).Return(&gcm.Response{}, nil)

	subRoute.Deliver(m, true)
	select {
	case <-doneC:
	case <-time.After(100 * time.Millisecond):
		a.Fail("Message not received by FCM")
	}

	m = &protocol.Message{
		Path:       "/topic",
		ID:         1,
		Body:       []byte(`plain body`),
		HeaderJSON: `{"Correlation-Id": "7sdks723ksgqn"}`,
	}

	mocks.gcmSender.EXPECT().Send(gomock.Any()).Do(func(m *gcm.Message) (*gcm.Response, error) {
		a.Nil(m.Notification)

		a.NotNil(m.Data)
		a.Contains(m.Data, "message")

		doneC <- true
		return nil, nil
	}).Return(&gcm.Response{}, nil)

	subRoute.Deliver(m, true)
	select {
	case <-doneC:
	case <-time.After(100 * time.Millisecond):
		a.Fail("Message not received by FCM")
	}
}

func TestConn_HandleResponseReporting(t *testing.T) {
	ctrl, finish := testutil.NewMockCtrl(t)
	defer finish()
	mockProducer := NewMockProducer(ctrl)
	a := assert.New(t)
	fcm, mocks := testFCM(t, true, mockProducer)

	err := fcm.Start()
	a.NoError(err)

	var route *router.Route

	mocks.router.EXPECT().Subscribe(gomock.Any()).Do(func(r *router.Route) (*router.Route, error) {
		a.Equal("/topic", string(r.Path))
		a.Equal("user_id", r.Get("user_id"))
		a.Equal("device_id", r.Get(deviceTokenKey))
		route = r
		return r, nil
	})


	mockProducer.EXPECT().Report(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(topic string, bytes []byte, key string) {
		a.Equal("sub_topic", topic)
	})

	// put a dummy FCM message with minimum information
	postSubscription(t, fcm, "user_id", "device_id", "topic")
	time.Sleep(100 * time.Millisecond)
	a.NoError(err)
	a.NotNil(route)



	// expect the route unsubscribed
	mocks.router.EXPECT().Unsubscribe(gomock.Any()).Do(func(route *router.Route) {
		a.Equal("/topic", string(route.Path))
		a.Equal("device_id", route.Get(deviceTokenKey))
	})

	// expect the route subscribe with the new canonicalID from replaceSubscriptionWithCanonicalID
	mocks.router.EXPECT().Subscribe(gomock.Any()).Do(func(route *router.Route) {
		a.Equal("/topic", string(route.Path))
		a.Equal("user_id", route.Get("user_id"))
		appid := route.Get(deviceTokenKey)
		a.Equal("fcmCanonicalID", appid)
	})
	// mocks.store.EXPECT().MaxMessageID(gomock.Any()).Return(uint64(4), nil)

	message := &protocol.Message{
		UserID:     "user_id",
		ID:         42,
		HeaderJSON: `{"Content-Type": "text/plain", "Correlation-Id": "7sdks723ksgqn"}`,
		Body: []byte(`{
		"to":"",
		"data":{
			"time":"2016-09-08T08:25:13+02:00",
			"type":"general",
			"notification_title":"Valid Title",
			"notification_body":"Die größte Sonderangebot!",
			"deep_link":"rewe://angebote"
		}
	}`),
	}

	mockProducer.EXPECT().Report(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(topic string, bytes []byte, key string) {
		a.Equal("fcm_topic", topic)

		var event FcmEvent
		err := json.Unmarshal(bytes, &event)
		a.NoError(err)
		a.Equal("pn_reporting_fcm", event.Type)
		a.Equal("Fail", event.Payload.Status)
		a.Equal("5", event.Payload.CanonicalID)
		a.Equal("7sdks723ksgqn", event.Payload.CorrelationID)
		a.Equal("device_id", event.Payload.DeviceID)
		a.Equal("user_id", event.Payload.UserID)
		a.Equal("Valid Title", event.Payload.NotificationTitle)
		a.Equal("Die größte Sonderangebot!", event.Payload.NotificationBody)
		a.Equal("rewe://angebote", event.Payload.DeepLink)
		a.Equal("general", event.Payload.Topic)
		a.Equal("InvalidRegistration", event.Payload.ErrorText)
	})

	response := new(gcm.Response)
	err = json.Unmarshal([]byte(ErrorFCMResponse), response)
	a.NoError(err)
	mocks.gcmSender.EXPECT().Send(gomock.Any()).Return(response, nil)

	// send the message into the subscription route channel
	route.Deliver(message, true)

	// wait before closing the FCM connector
	time.Sleep(100 * time.Millisecond)

	err = fcm.Stop()
	a.NoError(err)

	//then
	a.NoError(err)
}

func testFCM(t *testing.T, mockStore bool, producer kafka.Producer) (connector.ResponsiveConnector, *mocks) {
	mcks := new(mocks)

	mcks.router = NewMockRouter(testutil.MockCtrl)
	mcks.router.EXPECT().Cluster().Return(nil).AnyTimes()

	kvs := kvstore.NewMemoryKVStore()
	mcks.router.EXPECT().KVStore().Return(kvs, nil).AnyTimes()

	key := "TEST-API-KEY"
	nWorkers := 1
	endpoint := ""
	prefix := "/fcm/"
	intervalMetrics := false

	mcks.gcmSender = NewMockSender(testutil.MockCtrl)
	sender := NewSender(key)
	sender.gcmSender = mcks.gcmSender

	conn, err := New(mcks.router, sender, Config{
		APIKey:          &key,
		Workers:         &nWorkers,
		Endpoint:        &endpoint,
		Prefix:          &prefix,
		IntervalMetrics: &intervalMetrics,
	},
		producer,
		"sub_topic",
		"fcm_topic",
	)
	assert.NoError(t, err)
	if mockStore {
		mcks.store = NewMockMessageStore(testutil.MockCtrl)
		mcks.router.EXPECT().MessageStore().Return(mcks.store, nil).AnyTimes()
	}

	return conn, mcks
}

func postSubscription(t *testing.T, fcmConn connector.ResponsiveConnector, userID, gcmID, topic string) {
	a := assert.New(t)
	u := fmt.Sprintf("http://localhost/fcm/%s/%s/%s", gcmID, userID, topic)
	req, err := http.NewRequest(http.MethodPost, u, nil)
	a.NoError(err)
	w := httptest.NewRecorder()

	fcmConn.ServeHTTP(w, req)

	a.Equal(fmt.Sprintf(`{"subscribed":"/%s"}`, topic), string(w.Body.Bytes()))
}

func deleteSubscription(t *testing.T, fcmConn connector.ResponsiveConnector, userID, gcmID, topic string) {
	a := assert.New(t)
	u := fmt.Sprintf("http://localhost/fcm/%s/%s/%s", gcmID, userID, topic)
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	a.NoError(err)
	w := httptest.NewRecorder()

	fcmConn.ServeHTTP(w, req)

	a.Equal(fmt.Sprintf(`{"unsubscribed":"/%s"}`, topic), string(w.Body.Bytes()))
}

func removeTrailingSlash(path string) string {
	if len(path) > 1 && path[len(path)-1] == '/' {
		return path[:len(path)-1]
	}
	return path
}
