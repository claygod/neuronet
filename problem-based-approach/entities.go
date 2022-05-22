package problembasedapproach

import "sync"

// Problem-based approach
// Copyright © 2021-2022 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Need - потребность
*/
type Need struct {
	ID string
}

type TaskGenerator interface {
	GenTaskFromNeed(*Need)
}

/*
Task - задача
*/
type Task struct {
	ID          string
	State       uint64 // состояние, возможно будет каждый раз пересчитываться (получать через метод)
	ParentTasks []*Task
	ChildTasks  []*Task
	// TODO: Target
}

type TaskCompletionCheck interface {
	Check(*Task) float64 // оценка скорей всего от 0.0 до 1.0
}

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

type ChainletContainer struct { // цепочка действий имеющая удовленворяющий результат (смысл)
	// ID uint64 // возможно снаружи
	Rate     float64
	Chainlet *Chainlet
}

type ChainletRepo interface { // репо атомиков
	SetNewChainlet(*Chainlet) (ID uint64)
}

type AtomicChanger interface { // минимальное атомарное изменение
	Change(*State) *State
}

type AtomicChangerRepo interface { // репо атомиков
	GetRandom() (ID uint64, aChanger AtomicChanger)
}

type StateComparer interface { // сравниваем состояния (направление и координаты)
	Comparison(*State, *State) float64
}

type State struct {
	// vector - coord. and direct
}

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
GenChainlet - генерируем цепочку.
*/
func (c *ChainletGenerator) GenChainlet(maxSimilarity float64, curState *State, targetState *State) *ChainletContainer {
	out := &ChainletContainer{
		Rate:     0.0,
		Chainlet: NewChainlet(),
	}

	for i := 0; i < c.MaxChainletLenght; i++ {
		chID, chGer := c.ChangersRepo.GetRandom()
		out.Chainlet.Add(chID)
		curState = chGer.Change(curState)

		if out.Rate = c.Comparer.Comparison(curState, targetState); out.Rate >= maxSimilarity {
			break
		}
	}

	return out
}
