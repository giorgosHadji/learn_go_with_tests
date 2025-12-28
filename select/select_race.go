package select_test

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	/*you can wait for values to be sent to a channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.

	select allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.*/
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	/*Always make channels

	  Notice how we have to use make when creating a channel; rather than say var ch chan struct{}. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.

	  For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels
	*/
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
