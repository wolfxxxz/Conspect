package main

import (
	"fmt"
	"time"
)

/*
	func receiveDispatches(channel <-chan DispatchNotification) {
		for details := range channel {
			fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
				"x", details.Product.Name)
		}
		fmt.Println("Channel has been closed")
	}
*/
func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		channel <- p
		time.Sleep(time.Millisecond * 800)
	}
	close(channel)
}

func main() {
	dispatchChannel := make(chan DispatchNotification, 100)
	//var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	//var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	go DispatchOrders(dispatchChannel)
	//receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
	productChannel := make(chan *Product)
	go enumerateProducts(productChannel)
	openChannels := 2
	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Dispatch channel has been closed")
				dispatchChannel = nil
				openChannels--
			}
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("Product:", product.Name)
			} else {
				fmt.Println("Product channel has been closed")
				productChannel = nil
				openChannels--
			}
		default:
			if openChannels == 0 {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")
}

/*
	for details := range dispatchChannel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
	/*
		go DispatchOrders(dispatchChannel)
		for {
			if details, open := <-dispatchChannel; open {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
					"x", details.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				break
			}
		}
		/*fmt.Println("main function started")
		CalcStoreTotal(Products)
		//time.Sleep(time.Second * 5)
		fmt.Println("main function complete")

}*/
