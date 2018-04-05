package json

import (
	"encoding/json"
	"fmt"
	"reflect"

	eventsource "github.com/skf/go-eventsource"
)

// JSONSerializer ...
type serializer struct {
	eventTypes map[string]reflect.Type
}

func NewSerializer(events ...eventsource.Event) eventsource.Serializer {
	eventTypes := map[string]reflect.Type{}
	for _, event := range events {
		eventType := getTypeOfValue(event)
		eventTypes[eventType.Name()] = eventType
	}
	return &serializer{eventTypes: eventTypes}
}

func getTypeOfValue(input interface{}) reflect.Type {
	value := reflect.TypeOf(input)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value
}

// UnmarshalEvent ...
func (s *serializer) Unmarshal(data []byte, eventType string) (out eventsource.Event, err error) {
	recordType, ok := s.eventTypes[eventType]
	if !ok {
		err = fmt.Errorf("Unmarshal error, unbound event type, %v", eventType)
		return
	}

	event := reflect.New(recordType).Interface()
	if err = json.Unmarshal(data, event); err != nil {
		return
	}

	if out, ok = event.(eventsource.Event); !ok {
		err = fmt.Errorf("Event doesnt implement Event")
		return
	}

	return
}

func (s *serializer) Marshal(event eventsource.Event) (data []byte, err error) {
	return json.Marshal(event)
}
