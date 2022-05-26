package problembasedapproach

import "sync"

// Problem-based approach
// Chainlet entities
// Copyright © 2021-2022 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type Chainlet struct { // цепочка действий имеющая удовленворяющий результат (смысл)
	//ID uint64 // возможно снаружи
	// Rate float64
	Chain []uint64 // храним идентификаторы а не ссылки чтобы сравнивать цепочки на похожесть
}

func NewChainlet() *Chainlet {
	return &Chainlet{
		Chain: make([]uint64, 0),
	}
}

func (c *Chainlet) Add(chID uint64) {
	c.Chain = append(c.Chain, chID)
}

/*
hainletContainer - контейнер нужен для того, чтобы иметь возможность сравнить
*/
type ChainletContainer struct {
	// ID uint64 // возможно снаружи
	Rate     float64 // исчисляется исходя не только из коэффициэнта сравнения state, но и длины цепочки (количества действий)
	Chainlet *Chainlet
}

type ChainletRepo interface { // репо цепочек
	SetNewChainlet(*Chainlet) (ID uint64)
}

/*
ChainletGenerator - генерирует набор цепочек действий, которые можно провести с текущим состоянием
*/
type ChainletGenerator struct {
	MaxChainletLenght int
	MaxVersionsCount  int
	// TODO: Parallelism
	ChangersRepo AtomicChangerRepo
	Comparer     StateComparer
}

func (c *ChainletGenerator) GenChainlets(maxSimilarity float64, minSimilarity float64, curState *State, targetState *State) []*ChainletContainer {
	wg := sync.WaitGroup{}
	wg.Add(c.MaxVersionsCount)

	out := make([]*ChainletContainer, c.MaxVersionsCount)

	for i := 0; i < c.MaxVersionsCount; i++ {
		num := i

		go func() {
			out[num] = c.GenChainlet(maxSimilarity, curState, targetState)

			wg.Done()
		}()
	}

	wg.Wait()

	// TODO: тут сортируем и обрезаем по minSimilarity

	return out
}

/*
GenChainlet - генерируем цепочку (один из вариантов набора последовательности действий)
*/
func (c *ChainletGenerator) GenChainlet(maxSimilarity float64, curState *State, targetState *State) *ChainletContainer {
	out := &ChainletContainer{
		Rate:     0.0,
		Chainlet: NewChainlet(),
	}

	for i := 0; i < c.MaxChainletLenght; i++ {
		chID, chGer := c.ChangersRepo.GetRandom() // каждый раз берём случайное действие
		out.Chainlet.Add(chID)
		curState = chGer.Change(curState)

		if out.Rate = c.Comparer.Comparison(curState, targetState); out.Rate >= maxSimilarity {
			break
		}
	}

	return out
}
