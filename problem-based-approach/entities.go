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
	State       uint64 // состояние, возможно будет каждый раз пересчитываться (получать через метод)
	ParentTasks []*Task
	ChildTasks  []*Task
}
