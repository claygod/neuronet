package problembasedapproach

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
	CurState    *State
	TargetState *State
	ParentTasks []*Task
	ChildTasks  []*Task
	// TODO: scope - контекст задачи. Возможно скоп, он подобен State (допустим это стартовый стейт при решении задачи).
	// Также возможно, что скоп снаружи, и возможный для использования генератор Chainlet-наборов уже относится к какому-то скопу).
}

type TaskCompletionCheck interface {
	Check(*Task, *Task) float64 // оценка скорей всего от 0.0 до 1.0
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
