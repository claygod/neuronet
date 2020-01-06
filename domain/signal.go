package domain

// NeuroNet
// Signal
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type Signal struct {
	UID    uint64
	Level  uint64 // нулевое значение говорит о том, что это пустышка, в остальных случаях он пробрасываясь растёт
	From   uint64
	To     uint64
	Parent *Signal
	Owner  NeuronInterface
	Weight int64 //TODO: вот тут главная кухня, надо разобраться с допустимыми значениями
}

func NewSignal(uid uint64, level uint64, parent *Signal, owner NeuronInterface, weight int64) *Signal {
	return &Signal{
		UID:    uid,
		Level:  level,
		Parent: parent,
		Owner:  owner,
		Weight: weight,
	}
}

func (s *Signal) CreateChild(owner NeuronInterface, weight int64) *Signal {
	return NewSignal(s.UID, s.Level+1, s, owner, weight) //TODO:
}

func (s *Signal) Clone() *Signal {
	return NewSignal(s.UID, s.Level, s.Parent, s.Owner, s.Weight) //TODO: это для аксона
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
