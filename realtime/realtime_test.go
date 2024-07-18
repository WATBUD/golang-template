package realtime

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/centrifugal/gocent/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mai.today/realtime/internal/centrifugo"
)

type testStruct struct {
	Name  string
	Value int
}

func Test_Instance(t *testing.T) {
	// Call Instance() multiple times
	i1 := Instance()
	i2 := Instance()

	assert.Equal(t, i1, i2)
}

func TestRealtime_Broadcast(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := centrifugo.NewMockClient(ctrl)

	ctx := context.Background()
	channels := []string{"channel1", "channel2"}
	data := testStruct{
		Name:  "test",
		Value: 42,
	}
	expectedResult := gocent.BroadcastResult{}
	encodedData, err := encodeToBytes(data)
	assert.NoError(t, err)

	// Set up expected calls and return values
	mockClient.EXPECT().Broadcast(ctx, channels, encodedData).Return(expectedResult, nil)

	// Call the method under test
	rt := &Realtime{mockClient}
	result, err := rt.Broadcast(ctx, channels, data)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestRealtime_BroadcastToUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := centrifugo.NewMockClient(ctrl)

	ctx := context.Background()
	userIDs := []string{"id1", "id2"}
	channels := make([]string, len(userIDs))
	for i, userID := range userIDs {
		channels[i] = addUserPrefix(userID)
	}
	data := testStruct{
		Name:  "test",
		Value: 42,
	}
	expectedResult := gocent.BroadcastResult{}
	encodedData, err := encodeToBytes(data)
	assert.NoError(t, err)

	// Set up expected calls and return values
	mockClient.EXPECT().Broadcast(ctx, channels, encodedData).Return(expectedResult, nil)

	// Call the method under test
	rt := &Realtime{mockClient}
	result, err := rt.BroadcastToUsers(ctx, userIDs, data)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestRealtime_Publish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := centrifugo.NewMockClient(ctrl)

	ctx := context.Background()
	channel := "test_channel"
	data := testStruct{
		Name:  "test",
		Value: 42,
	}
	expectedResult := gocent.PublishResult{}
	encodedData, err := encodeToBytes(data)
	assert.NoError(t, err)

	// Set up expected calls and return values
	mockClient.EXPECT().Publish(ctx, channel, encodedData).Return(expectedResult, nil)

	// Call the method you want to test, injecting the mock client
	rt := &Realtime{mockClient}
	result, err := rt.Publish(ctx, channel, data)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestRealtime_PublishToUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := centrifugo.NewMockClient(ctrl)

	ctx := context.Background()
	userID := "test_user_id"
	channel := addUserPrefix(userID)
	data := testStruct{
		Name:  "test",
		Value: 42,
	}
	expectedResult := gocent.PublishResult{}
	encodedData, err := encodeToBytes(data)
	assert.NoError(t, err)

	// Set up expected calls and return values
	mockClient.EXPECT().Publish(ctx, channel, encodedData).Return(expectedResult, nil)

	// Call the method you want to test, injecting the mock client
	rt := &Realtime{mockClient}
	result, err := rt.PublishToUser(ctx, userID, data)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_addUserPrefix(t *testing.T) {
	tests := []struct {
		userID   string
		expected string
	}{
		{"user123", "#user123"},
		{"john_doe", "#john_doe"},
		{"", "#"},
	}

	for _, test := range tests {
		result := addUserPrefix(test.userID)
		if result != test.expected {
			t.Errorf("addUserPrefix(%s) returned %s, expected %s", test.userID, result, test.expected)
		}
	}
}

func Test_encodeToBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   testStruct
		wantErr bool
	}{
		{
			name:    "Encode struct",
			input:   testStruct{Name: "test", Value: 42},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeToBytes(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				data := testStruct{}
				err = json.Unmarshal(got, &data)
				assert.NoError(t, err)

				assert.NoError(t, err)
				assert.Equal(t, data.Name, tt.input.Name)
			}
		})
	}
}

func Test_newRealtime(t *testing.T) {
	const addr = "http://localhost:8888/api"
	const key = "testkey"

	rt := newRealtime(addr, key)
	c := gocent.New(gocent.Config{Addr: addr, Key: key})
	assert.Equal(t, c, rt.centrifugo)
}
