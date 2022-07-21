package event

import (
	"bytes"
	"encoding/gob"
	"meower/schema"

	"github.com/nats-io/nats.go"
)

// struct for Storing subscription and Connection
type NatsEventStore struct {
	nc                      *nats.Conn
	MeowCreatedSubscription *nats.Subscription
	meowCreatedChan         chan MeowCreatedMessage
}

//Create new NAts connection
func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

//Close our NAts connection and unsub
func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.MeowCreatedSubscription != nil {
		e.MeowCreatedSubscription.Unsubscribe()
	}
	close(e.meowCreatedChan)
}

func (e *NatsEventStore)PublishMeowCreated(meow schema.Meow)error{
	m := MeowCreatedMessage{meow.ID,meow.Body,meow.CreatedAt}
	data,err := e.writeMessage(&m)
	if err != nil{
		return err
	}
	return e.nc.Publish(m.Key(),data)
}

func (mq *NatsEventStore)writeMessage(m Message)([]byte,error){
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil{
		return nil,err
	}
	return b.Bytes(),nil
}

func (e *NatsEventStore) OnMeowCreated(f func(MeowCreatedMessage))(err error){
	m := MeowCreatedMessage{}
	e.MeowCreatedSubscription, err := e.nc.Subscribe(m.Key(),func(msg *nats.Msg){
		e.readMessage(msg.Data,&m)
		f(m)
	})
	return
}

func( mq * NatsEventStore)readMessage(data []byte,m interface{})error{
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}