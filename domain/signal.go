package domain

// NeuroNet
// Signal
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type Signal struct {
	uid   uint64
	level uint64 // нулевое значение говорит о том, что это пустышка, в остальных случаях он пробрасываясь растёт
	from  uint64
	to    uint64
	// parent *Signal
	// owner  *Neuron
	weight int64 //TODO: вот тут главная кухня, надо разобраться с допустимыми значениями
}

type SignalAggregator struct {
	//TODO: эта структура может быть встроена в ownerNeuron *Neuron
	claimCount uint64
	counter    uint64
	signals    []*Signal //количество сигналов должно быть увязано с их местами
}

func (s *SignalAggregator) AddSignal(sig *Signal) {
	//TODO: по накоплению сигналов они передаются в нейрон
}
