done := make(chan struct{})
go func() {
	work()
	close(done)
}()
<-done

// Внутри функции main (функцию объявлять не нужно),
// вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения. 
