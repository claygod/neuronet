package domain

// NeuroNet
// Axon
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"math/rand"
)

/*
Axon - аксон должен реализовать стохастический (не детерменистский) подход.
*/
type Axon struct {
	ownerNeuronId uint64
	outList       []AxonOut
	random        *rand.Rand
	neuronsRepo   NeuronsRepo
}

func (a *Axon) signalBroadcast(sigs []*Signal) {
	for _, sig := range sigs {
		neu, err := a.neuronsRepo.Get(sig.to)
		if err != nil {
			//TODO: удалить из списка несуществующий нейрон
			//TODO: лог ошибок return err
		}
		dd, err := neu.GetDendrite(sig.from)
		if err != nil {
			//TODO: удалить из списка несуществующую связь
			//TODO: лог ошибок return err
		}
		dd.TransmitSignal(sig)
	}
}

/*
SendSignalStochasticMode - отправка сигнала одному из подключенных к аксону нейронов,
выбранному случайным стохастическим методом. Это может имитировать любознательность.
Но рассылка будет всем, просто это будут пустышки. Рассылка всем нужна для синхронизации.
*/
func (a *Axon) SendSignalStochastic(sig *Signal) error {
	if len(a.outList) == 0 {
		return fmt.Errorf("List of AxonOut is empty.")
	}
	var total int64 = 0
	for _, ao := range a.outList {
		if ao.weigth > 0 {
			total += ao.weigth
		}
	}
	if total == 0 {
		return fmt.Errorf("The total weight is zero.")
	}
	rnd := a.random.Int63n(total)
	total = 0
	var neuronTo uint64
	for _, ao := range a.outList {
		if ao.weigth > 0 {
			total += ao.weigth
		}
		if total > rnd {
			neuronTo = ao.neuronId
			break
		}
	}
	// подготавливаем список сигналов, среди которых только один будет не пустышкой
	sigList := a.prepareSignalDoubled(sig, 0)
	sigList[neuronTo] = sig
	return nil
}

func (a *Axon) prepareSignalDoubled(sig *Signal, level uint64) map[uint64]*Signal {
	sigList := make(map[uint64]*Signal, len(a.outList))
	for _, ao := range a.outList {
		s := &Signal{
			uid:    sig.uid,
			level:  level,
			from:   sig.from,
			to:     ao.neuronId,
			weight: sig.weight,
		}
		// if weigthCopy {
		// 	s.weigth = sig.weigth
		// }
		sigList[ao.neuronId] = s
	}
	return sigList
}

type AxonOut struct {
	weigth   int64
	neuronId uint64
}
