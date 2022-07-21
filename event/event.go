package event

import "meower/schema"

type EvenStore interface{
	CLose()
	PublishMeowCreated(meow schema.Meow)error
	SubscibeMeowCreated()(<-chan MeowCreatedMessage,error)
	OnMeowCreated(f func(MeowCreatedMessage))error
}

var impl EvenStore

func SetEvenStore(es EvenStore){
	impl = es
}

func Close(){
	impl.CLose()
}

func PublishMeowCreated(meow schema.Meow) error{
	return impl.PublishMeowCreated(meow)
}

func SubscibeMeowCreated()(<-chan MeowCreatedMessage,error){
	return impl.SubscibeMeowCreated()
}

func OnMeowCreated(f func(MeowCreatedMessage))error{
	return impl.OnMeowCreated(f)
}