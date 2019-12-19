package domain

// NeuroNet
// interfaces
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// import (
// 	"unsafe"
// )

/*
NeuronInterface - интерфейс нейрона. Он должен уметь:
- получать сигнал и рассылать результаты
- изменяться после оценки прогноза
*/
type NeuronInterface interface {
	AppendSignal(sig *Signal)                   // в нейрон добавляем новый сигнал
	ForecastEstimate(weigth int64, sig *Signal) // оценка совпадения прогноза и результата
	NextStep()                                  // переход к следующему шагу (списывание здоровья на шаг и т.п.)
	HealthCheck() int64                         // показатель здоровья
}

/*
DendriteInterface - интерфейс дендрита. Он должен уметь:
- получать сигнал и анализировать сигнал
- отдавать результат анализа сигнала
- изменяться после оценки прогноза
*/
type DendriteInterface interface {
	AppendSignal(sig *Signal)      // в дендрит добавляем новый сигнал
	ForecastEstimate(weigth int64) // оценка совпадения прогноза и результата
	HealthCheck() int64            // показатель здоровья
}

/*
AxonInterface - интерфейс аксона. Он должен уметь:
- подключать и отключать нейроны
- рассылать сигнал по подключенным нейронам
*/
type AxonInterface interface {
	AddNeuron(id uint64)                   // подключаем к рассылке новый нейрон
	RemoveNeuron(id uint64)                // отключаем один из нейронов
	BroadcastTotal(sig *Signal) int64      // рассылка сигнала всем подключенным нейронам
	BroadcastStochastic(sig *Signal) int64 // рассылка пустышки всем и одного случайному нейрону (из подключенных)
}
