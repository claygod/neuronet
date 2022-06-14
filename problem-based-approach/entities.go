package problembasedapproach

// Problem-based approach
// Entities
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
	ID            string
	maxSimilarity float64
	minSimilarity float64

	recursionLevel   int
	findDirectsCount int

	beginState  *State
	curState    *State
	targetState *State

	chlGen    *ChainletGenerator
	rStateGen IntermediateRandomStateGenerator
	sComparer StateComparer

	// Шаги, ведущие к цели
	Steps []*ChainletContainer

	ParentTasks []*Task
	ChildTasks  []*Task
	// TODO: scope - контекст задачи. Возможно скоп, он подобен State (допустим это стартовый стейт при решении задачи).
	// Также возможно, что скоп снаружи, и возможный для использования генератор Chainlet-наборов уже относится к какому-то скопу).
}

func (t *Task) FindChainlets() []*ChainletContainer { // тут мы ищем оптимальный путь
	decisions := t.chlGen.GenChainlets(t.maxSimilarity, t.minSimilarity, t.curState, t.targetState)

	if len(decisions) == 0 && t.recursionLevel > 0 { // не найдено подходящих решений и ещё можно создавать промежуточные шаги
		// TODO:
		// генерируем новые (промежуточные) цели, которых можем добиться
		// и уже в каждой точке промежуточных целей пробуем заново добиться основной цели
		// (действуем рекурсивно)

		for i := 0; i < t.findDirectsCount; i++ {
			newState := t.rStateGen.GenTask(t.curState, t.targetState, t.sComparer)

			newTask := &Task{ // TODO: сделать через NewTask с заполнением всех полей
				recursionLevel: t.recursionLevel - 1,

				beginState:  t.beginState,
				curState:    t.curState,
				targetState: newState,
			}

			for _, dt := range newTask.FindChainlets() { // это получаем результаты к промежуточной цели

				newTask2 := &Task{ // TODO: сделать через NewTask с заполнением всех полей
					recursionLevel: t.recursionLevel - 2,

					beginState:  t.beginState,
					curState:    newState,
					targetState: t.targetState,
				}

				for _, dt2 := range newTask2.FindChainlets() { // теперь из промежуточной точки пытамся добраться до основной цели
					decisions = append(decisions, MergeChainletContainers(dt, dt2))
				}
			}

		}
	}

	return decisions
}

type TaskCompletionCheck interface {
	CompletionCheck(*Task, *Task) float64 // оценка скорей всего от 0.0 до 1.0 CompletionCheck
}

type IntermediateRandomStateGenerator interface {
	GenTask(*State, *State, StateComparer) *State // генерирование некоторого состояния, находящегося где-то между начальной и конечной задачей
}

type AtomicChanger interface { // минимальное атомарное изменение
	Change(*State) *State
}

type AtomicChangerRepo interface { // репо атомиков
	/*
	   GetRandom - берём случайную, это удобно для генерации случайного Chainlet-набора
	*/
	GetRandom() (ID uint64, aChanger AtomicChanger)

	/*
		SetRandom - сначала добавляем действительно базовые возможности, а потом можно добавлять
		Chainlet-наборы, которые используются часто или которыекороткие но эффективные
	*/
	SetRandom(aChanger AtomicChanger) (ID uint64)
}

type StateComparer interface { // сравниваем состояния (направление и координаты)
	Comparison(*State, *State) float64
}

type State struct {
	// vector - coord. and direct
}

/*
TaskResource -  учёт ресурсов, выделенных для решения задачи, обычно ресурсы только тратятся,
но при каких-то определённых обстоятельствах ресурсы могут и повышаться
(например найден Chainlet, достойный добавление в репо Changer-атомиков.
*/
type TaskResource interface { // NOTE: возможно тут потребуются float
	Add(int64) int64
	Cut(int64) int64
	Total() int64
	ResetToZero() // напоминание о том, что у задачи может оказаться ситуация, когда точно надо остановить поиски путей (Chainlet) её выполнения
}
