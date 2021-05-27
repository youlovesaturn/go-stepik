toDelete := append(groupCity[10], groupCity[1000]...)
for _, k := range toDelete {
	delete(cityPopulation, k)
}

// В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:
// groupCity := map[int][]string{
// 	10:   []string{...}, // города с населением 10-99 тыс. человек
// 	100:  []string{...}, // города с населением 100-999 тыс. человек
// 	1000: []string{...}, // города с населением 1000 тыс. человек и более
// }
// При подготовке важного отчета о городах с населением 100-999 тыс. человек был подготовлен другой объект типа map:
// cityPopulation := map[string]int{
// 	город: население города в тысячах человек,
// }
// Однако база данных с информацией о точной численности населения содержала ошибки,
// поэтому в cityPopulation в т.ч. была сохранена информация о городах, которые входят в другие группы из groupCity.
// Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation,
// чтобы в ней была сохранена информация только о городах из группы groupCity[100].
