
# 5.1 Паралелизм Многозадачность
Свойство среды паралельно обрабатывать задачи
**Вытесняющая многозадачность**
Система сама распределяет приоритеты
windows Linux and Mac
**конкурентная многозадачность Go**
где планировщик горутин самостоятельно управляет выполнением горутин. Каждая горутина выполняется независимо и может быть приостановлена или возобновлена планировщиком горутин в любой момент. Планировщик горутин определяет, когда горутина будет выполнена и на каком процессоре.
**Переключение контекста**
Желательно сохранится при переключении на другую задачу
# 5.2 Процесс и Поток Треды и горутины вступление
## Теория
Процесс - это запуск программы
Поток - запуск части процесса
Goroutine - это легковесный поток который может выполнятся паралельно и независимо от выполнения остального кода в программе
## Горутина на практике
### Без горутины
**main.go**
func workAndPrint(num int) {
	fmt.Printf("Start Job #%v\n", num)
	var calc int
	for i := 0; i < 1000; i++ {
		calc = i * num
	}
	fmt.Printf("End Job #%v: valc = %v\n ", num, calc)
}
func main() {
	for i := 1; i <= 5; i++ {
		workAndPrint(i)
	}
}
**go run .**
Start Job #1
End Job #1: valc = 999
 Start Job #2
End Job #2: valc = 1998
 Start Job #3
End Job #3: valc = 2997
 Start Job #4
End Job #4: valc = 3996
 Start Job #5
End Job #5: valc = 4995
### +go
**main.go**
func workAndPrint(num int) {
	fmt.Printf("Start Job #%v\n", num)
	var calc int
	for i := 0; i < 1000; i++ {
		calc = i * num
	}
	fmt.Printf("End Job #%v: valc = %v\n ", num, calc)
}
func main() {
	for i := 1; i <= 5; i++ {
		go workAndPrint(i)
	}
	time.Sleep(100 * time.Millisecond)
}
**go run .**
Start Job #5
End Job #5: valc = 4995
 Start Job #1
End Job #1: valc = 999
 Start Job #4
End Job #4: valc = 3996
 Start Job #2
Start Job #3
End Job #3: valc = 2997
 End Job #2: valc = 1998
**ОПИСАНИЕ**
1. Цикл в функции мейн отработал полностью и завершился
2. По сути выполнение go отвязало последовательное выполнение func workAndPrint
4. Выполнение go rutine main имеет приоритет по сравнению с остальными го рутинами (потоками)
5. В цикле было запущено 5 горутин и все они выполняются но если не будет time.Sleep - они просто не успеют вывести результат
6. Разбалансированными получаются даже принты - это доказывает что горутины выполняются разное количество времени
7. Горутины как пчёлы без летка - кто куда и как хочет так и летит
### горутины выполняются разное количество времени
**main.go**
func workAndPrint(num int) {
	fmt.Printf("Start Job #%v\n", num)

	for i := 0; i < 10000; i++ {
		_ = i * num
	}
	fmt.Printf("End Job #%v \n ", num)
}
func main() {
	for i := 1; i <= 5; i++ {
		go workAndPrint(i)
	}
	time.Sleep(100 * time.Millisecond)
}
**go run .**
Start Job #5
End Job #5 
 Start Job #1
End Job #1 
 Start Job #2
Start Job #4
Start Job #3
End Job #4 
 End Job #2 
 End Job #3 
## Пакет runtime.Gosched() - типа паузы 
#### runtime.Gosched() 
является простой функцией планировщика горутин и не принимает никаких аргументов. Она выполняет переключение контекста между горутинами, предоставляя возможность другим горутинам запуститься.
#### Пример
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printNumbers()

	// Вызов runtime.Gosched() позволяет переключить контекст выполнения горутины
	// и предоставить шанс другим горутинам на выполнение.
	for i := 1; i <= 3; i++ {
		runtime.Gosched()
		fmt.Println("Main Goroutine")
		time.Sleep(1 * time.Second)
	}
}
#### runtime.GOMAXPROCS(n int):
 Эта функция устанавливает максимальное количество процессоров, которые могут выполнять горутины параллельно. Если аргумент n равен 1, это означает, что горутины будут выполняться последовательно на одном процессоре. Если n больше 1, горутины могут выполняться параллельно на нескольких процессорах. Значение по умолчанию равно количеству доступных процессоров на машине.
#### runtime.NumCPU(): 
Эта функция возвращает количество доступных процессоров на машине.
fmt.Println(runtime.NumCPU())
#### runtime.NumGoroutine(): 
Эта функция возвращает текущее количество активных горутин.
#### runtime.Goexit(): 
Эта функция завершает текущую горутину без завершения остальных горутин.
#### runtime.GC():
 Эта функция запускает сборщик мусора (garbage collector) для освобождения неиспользуемой памяти.
# 5.3 Обработка ошибок и Panic
## 5.3.1 err := errors.New("error happened")
### Теория
Errors vs Exceptions
Err - находится под управлением програмиста (без чего в принципе можно работать приложению)
Panic - бросает система, паника останавливает работу приложения(что то серьёзное)
### Interface error или Как создать ошибку
**ошибка не должна начинатся с заглавной буквы**
type error interface {
    Error() string
}
если есть функция удовлетворяющая интерфейсу error
**Как создать ошибку**
err := errors.New("error happened")
//В строку можно вложить переменную
err := fmt.Errorf("unsupported type: %T", v)
//В строку можно вложить ошибку
err := fmt.Errorf("wrapped err: %w", firsterr)
### Вложенные ошибки и методы их обработки
//Как разделить вложенную ошибку?
// Вернёт вложенную ошибку
prevErr := errors.Unvrap(currentErr)
// Проверяет ошибку на соответствие типу ошибки (создание)
if errors.Is(err, target) {...}
### Пример errors.Unvrap
func newMistake() (string, error) {
	d := "one mistake"
	err := fmt.Errorf("first mistake: %w", errors.New("second mistake"))
	return d, err
}

func main() {
	_, mistake := newMistake()
	fmt.Println(mistake)                //first mistake: second mistake
	fmt.Println(errors.Unwrap(mistake)) //second mistake
}
### Примеры
### Пример if errors.Is(err, target)
**Если ошибка связана с типом...**
type TypeMistake struct {
	err error
}

func NewMistake(message string) *TypeMistake {
	return &TypeMistake{errors.New(message)}
}

var static = NewMistake("i'm a little mistake")

func newMistake() (string, error) {
	d := "one mistake"
	err := fmt.Errorf("first mistake: %T", static)
	return d, err
}

func main() {
	_, mistake := newMistake()
	fmt.Println(mistake) //first mistake: second mistake
	if errors.Is(mistake, static.err) {
		fmt.Println(errors.Unwrap(mistake)) //second mistake
	}
}
### Пример 1
func calc(a, b int) int {
	return a / b
}

func main() {
    // нельзя делить на 0
	i := calc(10, 0)
	fmt.Println(i)
}
**panic: runtime error**
**решение**
func calc(a, b int) (int, error) {
	if b == 0 {
		//Возвращаем дефолт + информацию про ошибку
		return 0, errors.New("division is zero")
	} else {
		return a / b, nil
	}
}

func main() {
	if i, err := calc(10, 0); err != nil {
		fmt.Println(i, err)
	} else {
		fmt.Println(i)
	}
    if i, err := calc(10, 5); err != nil {
		fmt.Println(i, err)
	} else {
		fmt.Println(i * i)
	}
}
**0 Division is zero**
**4**
**Программа завершится хоть и с ошибкой**
### Пример 2 можно вложить ошибку
func calc(a, b int) (int, error) {
	if b == 0 {
		//Возвращаем дефолт + информацию про ошибку
		mistake := errors.New("zero mistake")
		return a, fmt.Errorf("error division is %w", mistake)
	} else {
		return a / b, nil
	}
}

func main() {
	if i, err := calc(10, 0); err != nil {
		fmt.Println(i, err)

	} else {
		fmt.Println(i * i)
	}
	if i, err := calc(10, 5); err != nil {
		fmt.Println(i, err)
	} else {
		fmt.Println(i * i)
	}
}
**10 error division is zero mistake**
**4**
### Пример 3 Проверка на тип ошибки errors.Is
type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func main() {
	err := &CustomError{"custom error message"}

	// Проверка ошибки на соответствие типу CustomError
	if errors.Is(err, &CustomError{}) {
		fmt.Println("Error is of type CustomError")
	} else {
		fmt.Println("Error is not of type CustomError")
	}
}
### Пример 4
func Devote(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("zero isn't a second argument")
		return 0, err
	} else {
		return a / b, nil
	}
}

func main() {
	numbers := []int{50, 25, 0, 10, 5, 1}
	for i, v := range numbers {
		if i == len(numbers)-1 {
			fmt.Println("cancel, last argument is: ", v)
			break
		}
		if d, err := Devote(v, numbers[i+1]); err == nil {
			fmt.Println(d)
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("ok")
}
## 5.3.2 Panic
### Теория
Паника - это функция которая создаёт ошибку рантайма
panic("message")
recErr := recover() // recErr == message
**Runtime** (рантайм) в контексте программирования обычно относится к программной среде, которая обеспечивает выполнение и управление программой во время её работы. В языке Go (Golang) пакет runtime предоставляет функциональность, связанную с выполнением программы, такую как управление горутинами, аллокацией памяти, сборкой мусора и другими операциями, связанными с запущенной программой.
### пакет runtime в Go, включают:
**go** - для запуска новой горутины.
**GOMAXPROCS** - для установки максимального числа используемых процессоров.
**NumCPU** - для получения количества доступных процессоров.
**MemStats** - для получения статистики использования памяти.
**GC** - для явного запуска сборки мусора.
**Gosched** - для переключения выполнения между горутинами.
**Panic и Recover** - для обработки паники (непредвиденного сбоя программы) и восстановления контроля над программой.
Пакет runtime в Go предоставляет различные возможности для управления выполнением программы и управления ресурсами. Он является частью стандартной библиотеки Go и используется для создания эффективных и надежных программ.
### Пример борьбы с паникой
**Паника завершается закрытием программы**
**defer достаёт panic**
**recover() - стопорит и достаёт содержимое panic**
func main() {
	defer func() {
		fmt.Println("why does panic happen")
		err := recover()
		fmt.Println(err) // I'm a panic
	}()

	fmt.Println("Hello this is the begining of Panic  ")

	panic("I'm a panic")

	fmt.Println("everything doing wrong") // эта часть не будет выполнена
}
### Пример 2 panic
#### Шире
func callPanic() {
	panic("panic err")
}

func iteration(a, b int) {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	fmt.Println(a, b)

	callPanic()

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range arr {

		iteration(i, v)
		fmt.Println(i * v)
	}

	fmt.Println("everything doing wrong")
}
//3 4
//panic err
//12
#### Дальше
func callPanic() {
	panic("panic err")
}

func iteration(a, b int) int {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	fmt.Println(a, b)

	callPanic()
	return a * b

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range arr {

		iteration(i, v)

	}

	fmt.Println("everything doing wrong")
}
//6 7
//panic err
//everything doing wrong
### goroutine + panic + recover
func callPanic() {
	panic("panic err")
}

func iteration() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go callPanic() // при таком вызове defer не успеет сработать
	}
}

func main() {
	iteration()
	fmt.Println("everything doing wrong")
}
# 5.4 Каналы + Select
## Теория Синхронизация, Блокировка, Буфер, <-, close(ch)
### Синхронизация и канал без буфера: 
Каналы обеспечивают синхронизацию между горутинами. Запись в канал блокирует горутину, пока другая горутина не прочитает данные из канала, и наоборот, чтение из канала блокирует горутину, пока другая горутина не запишет данные в канал.
### Блокировка: 
Операции чтения и записи с каналами блокируют горутину до тех пор, пока они не будут успешно выполнены. Это позволяет горутинам синхронизировать свою работу и избежать состояния гонки.
### Пример Синхронизация и блокировка
func main() {
	ch := make(chan int)
	quit := make(chan int)

	// Горутина может работать с каналом без буфера
	go func() {
		for i := 1; i < 5; i++ {
			ch <- i
			fmt.Println("ch writed and wait read", i)
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}()

	go func() {
		for j := range ch {
			quit <- j
			fmt.Println("quit writed and wait read", j)
			time.Sleep(time.Second * 1)
		}

		close(quit)
	}()

	// Получается что мейн синхронизирует работу горутин

	for j := range quit {
		fmt.Println("main read", j)
	}
	fmt.Println("All goroutines completed.")

}
#### Пример 2 Канал без буфера Синхронизация
func send(ch chan int) {
	fmt.Println("Send: Начинаем работать")
	for i := 0; i <= 10; i++ {
		ch <- i
		fmt.Printf("Send: Отправляю в канал %v и жду", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Send: Закрываю канал")
	close(ch)
}

func get(ch chan int, b chan bool) {
	fmt.Println("Get: Открыта получать")
	for num := range ch {
		fmt.Printf(" || Get: Принимаю %v и жду следующего\n", num)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Get: Отправляю tru в канал для закрытия main")
	b <- true
}
func main() {
	fmt.Println("main начинаю работать")
	ch := make(chan int)
	bch := make(chan bool)
	go send(ch)

	go get(ch, bch)

	<-bch
	fmt.Println("main получила true и закончила работу")

}
### Пример main wait <- true
func main() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i < 11; i++ {
			fmt.Println("write ", i)
			ch <- i
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}()

	go func() {
		for num := range ch {
			fmt.Println("read ", num)
			time.Sleep(time.Second * 1)
		}
		done <- true

	}()
	<-done
}
### Буферизация:
 Каналы могут быть буферизованными или небуферизованными. Небуферизованные каналы имеют емкость 0, что означает, что запись и чтение происходят синхронно. Буферизованные каналы имеют емкость больше 0, что позволяет записывать данные в канал без блокировки до тех пор, пока канал не заполнится.
### Операторы <-: Оператор <- 
используется для отправки (записи) или получения (чтения) данных из канала. Например, ch <- value используется для отправки значения value в канал ch, а x := <-ch используется для получения значения из канала ch и присваивания его переменной x.
### close(ch), j,ok := <-chan, for i := range ch
**close(ch) сигнализирует о закрытии канала и range ch знает где закончить итерацию**
 Каналы могут быть закрыты с помощью функции close(ch). Закрытие канала указывает, что больше значений не будет отправлено в канал, но его можно по-прежнему использовать для чтения значений, которые уже находятся в канале.
## Небуферезированный канал
//небуферизованный канал создает точку синхронизации между
//горутинами, где запись и чтение происходят синхронно.
// Когда записывающая горутина записывает значение в
//небуферизованный канал, она блокируется до тех пор,
//пока другая горутина не прочитает значение.

func writeChan(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		// и вот тут происходит чудо
		// записав первое значение в канал
		// горутина останавлевается
		// и ждёт пока кто то не прочтёт его и
		// только потом записывает второе значение
		ch <- i
		fmt.Println("Жду пока прочитают", i)
		time.Sleep(time.Second * 1)
	}

	close(ch)
}

func main() {
	fmt.Println("Start First")

	ch := make(chan int)

	go writeChan(ch)

	for i := range ch {
		fmt.Println("Читаю")
		fmt.Println("chan i = ", i)
		time.Sleep(time.Second * 1)
	}

	fmt.Println("END First")
}
## Буферизованный канал
// буферизованный канал принимает все ОБЬЯВЛЕННЫЕ
// значения сразу а потом ведёт себя как небуферезированный

func writeChan(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		ch <- i
		fmt.Println("Пишу: ", i)
	}
	close(ch)
}

func main() {
	fmt.Println("Start First")

	ch := make(chan int, 5)

	go writeChan(ch)

	fmt.Println("time sleep")
	time.Sleep(time.Second * 2)

	for i := range ch {
		fmt.Println("Читаю chan i = ", i)
		time.Sleep(time.Second * 1)
	}

	fmt.Println("END First")
}
### Горутина создаёт буфер для небуфер канала
func main() {
	ch := make(chan int)
	quit := make(chan int, 1)

	// Горутина может работать с каналом без буфера
	go func() {
		ch <- 1
		close(ch)
	}()

	j := <-ch
	// main не может без буффера
	quit <- j
	close(quit)

	// результат не определён

	select {
	case x := <-ch:
		fmt.Println("ch = ", x)
	case <-quit:
		fmt.Println("quit")
	default:
		fmt.Println("default")
	}
}
### Пример 1 канал как входящий аргумент Функции
// Функция является БЛОКИРУЮЩЕЙ так как
// горутина будет ждать пока появятся данные в канале
func readChan(ch chan int) {
	// <- Читаем из канала
	value := <-ch
	fmt.Println("chan value: ", value)
}

func main() {
	fmt.Println("Start Main")

	// Обьявление переменной типа канал
	//var ch chan int

	// Создание канала без буфера
	// Маленький канал - мало данных
	ch := make(chan int)

	// если положить данные в канал до их вызова то будет дедлок
	//ch <- 100

	// Запуск горутины приведёт к ожиданию пока данные появятся 
	// в канале
	go readChan(ch)
	// Появляется вероятность что readChan - не успеет отработать
	time.Sleep(1 * time.Second)

	// Ложим в канал 77
	// бывает чтение из канала и запись в канал - стрелочка имеет значение
	ch <- 77

	fmt.Println("END MAIN")
}
### Пример 2 Утечка горутин
func readChan(ch chan int) {
	// <- Читаем из канала
	value := <-ch
	fmt.Println("chan value: ", value)
}

func main() {
	fmt.Println("Start Main")

	ch := make(chan int)

	go readChan(ch)

	// Утечка горутин
	// Если не послать данные в channel
	// goroutines будут накапливатся и
	// забивать память своим ожиданием
	//ch <- 50

	fmt.Println("END MAIN")
}
### Утечка горутин от gpt (goroutine leak)
package main

import (
	"fmt"
	"time"
)

func leakyFunction() {
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go leakyFunction()
	}

	// Без какого-либо механизма для остановки горутин
	// они будут продолжать работать вечно.
	// Основная горутина не будет завершаться.
	fmt.Println("End main")
	time.Sleep(time.Hour)
}
### канал с буфером ch := make(chan int, 1)
func readChan(ch chan int) {
	// <- Читаем из канала
	value := <-ch
	fmt.Println("chan value: ", value)
}

func main() {
	fmt.Println("Start Main")

	// Канал с буфером даёт возможность
	// хранить в нём значение до вызова
	// если количество отправленых значений в канал
	// будет больше чем может прочитать программа == дедлок
	ch := make(chan int, 2)

	ch <- 50
	ch <- 75
	// ch <- 22 //deadlock!
	go readChan(ch)

	ch <- 100

	go readChan(ch)
	go readChan(ch)

	time.Sleep(1 * time.Second)

	fmt.Println("END MAIN")
}
### Каналы от gpt
func sleep(t time.Duration, ch chan bool) {
	fmt.Println("Sleep ", t)
	time.Sleep(t)
	ch <- true
}

func main() {
	ch := make(chan bool)

	go sleep(1*time.Second, ch)
	<-ch // Ждем завершения первой горутины

	go sleep(5*time.Second, ch)
	<-ch // Ждем завершения второй горутины

	go sleep(7*time.Second, ch)
	<-ch // Ждем завершения третьей горутины

	fmt.Println("End main")
}
## 5.4.2 ch chan <- int || ch <-chan int
// ch chan<- int  - канал доступен только на запись
// ch <-chan int  - канал доступен только на чтение
func writeChan(ch chan<- int) {
    ch <- 42 // отправка значения в канал
}

В объявлении функции writeChan(ch chan<- int), стрелка <- указывает на направление передачи данных по каналу. В данном случае, chan<- int означает, что функция writeChan является отправителем и может только отправлять (писать) значения в канал ch.

То есть, функция writeChan принимает канал ch в качестве параметра, и стрелка <- перед типом канала (chan) указывает на то, что функция будет использовать этот канал только для отправки значений. Внутри функции writeChan можно использовать оператор ch <- value для отправки значения value в канал.

Таким образом, объявление функции с chan<- int позволяет явно указать, что функция может только отправлять значения в указанный канал. Это помогает сделать код более ясным и улучшить безопасность типов при работе с каналами.
## 5.4.3 Канал в цыкле for i := range ch; close(ch)
### for i:= range
func writeChan(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		ch <- i
		// и ждём пока кто то прочитает из канала
	}
	// закрываем канал
	// close(ch) - ставит метку что работа окончена
	// что б range знал что канал закрылся
	close(ch)
}

func main() {
	fmt.Println("Start Main")

	ch := make(chan int)

	go writeChan(ch)

	// Читаем с канала
	for i := range ch {
		fmt.Println("chan i = ", i)
	}

	fmt.Println("END MAIN")
}
### for j,ok := <-chan
**канал возвращает два значения**
func main() {
	ch := make(chan int, 1)
	quit := make(chan int, 1)

	ch <- 1
	close(ch)

	go func() {
		for {
			j, ok := <-ch
			if !ok {
				break
			}
			quit <- j
		}
		close(quit)
	}()

	runtime.Gosched()
	time.Sleep(time.Second * 1)

	select {
	case x := <-ch:
		fmt.Println("ch = ", x)
	case <-quit:
		fmt.Println("quit")
	default:
		fmt.Println("default")
	}
}
## 5.4.5 Select
### Инструкция Select
select {
    case c <- x:
         fmt.Println(c)
    case <- quit:
         fmt.Println("quit)
         return
}
### Ошибка при отправке в канал без буфера и без горутины
func main() {
	ch := make(chan int)

	// Возможен такой способ отправки в канал
	//Внимательно НЕбуферезированный канал так работать не будет
	ch <- 1
	//Когда в главной горутине выполняется ch <- 1,
	//она пытается отправить значение 1 в канал ch,
	// но остановится на этой строке, так как никакая
	// другая горутина не готова принять это значение из канала.
	//Это приводит к зависанию программы.
	runtime.Gosched()
	time.Sleep(time.Second * 1)

	select {
	case x := <-ch:
		fmt.Println("ch = ", x)
	default:
		fmt.Println("default")
	}
}
### Пример 1 Select
func writeChanal(ch chan<- int, b int) {
	time.Sleep(time.Second * 1)
	ch <- b
	close(ch)
}

func main() {
	ch := make(chan int, 1)
	quit := make(chan int, 1)

	go writeChanal(quit, 10)
	//go writeChanal(ch, 5)

	// Возможен такой способ отправки в канал
	//Внимательно НЕбуферезированный канал так работать не будет
	ch <- 1 //ch := make(chan int) == deadlock

	runtime.Gosched()
	time.Sleep(time.Second * 1)

	select {
	case x := <-ch:
		fmt.Println("ch = ", x)
	case <-quit:
		fmt.Println("quit")
	default:
		fmt.Println("default")
	}
}
### Пример 2 Select в цикле
func writeChanal(ch chan<- int, b int) {
	ch <- b
	close(ch)
}

func read(ch, quit <-chan int) {
	for {
		select {
		case x := <-ch:
			fmt.Println("ch = ", x)
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)

	go read(ch, quit)

	go writeChanal(quit, 10)
	go writeChanal(ch, 5)
	runtime.Gosched()
	time.Sleep(time.Second * 1)

}
### Пример 3 Select в цикле
func writeChanal(ch chan<- int, b int) {
	ch <- b
	close(ch)
}

func read(ch, quit <-chan int) {
	for {
		select {
		case x := <-ch:
			if x == 0 {
				continue
			} else {
				fmt.Println("ch = ", x)
			}
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)

	go read(ch, quit)

	go writeChanal(quit, 10)
	go writeChanal(ch, 5)
	runtime.Gosched()
	time.Sleep(time.Second * 1)

}
### Пример 4 select go and for
func send(ch chan int) {
	fmt.Println("Send: Начинаем работать")
	for i := 0; i <= 10; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func main() {
	fmt.Println("main начинаю работать")
	ch := make(chan int)
	ch2 := make(chan bool, 1)
	go send(ch)
	go func() {
		for {
			select {
			case _, ok := <-ch:
				if ok {
					fmt.Println("ok")
					continue
				} else {
					fmt.Println("ch<-true")
					//break
					ch2 <- true

				}
			}
			break
		}
	}()
	<-ch2

	fmt.Println("main получила true и закончила работу")
}
# 5.5 Примитивы синхронизации и гонки данных. Мьютекс, семофор и atomic. WaitGroup, timer, ticker, Context
## 5.5.2 Atomic "sync/atomic"
**операция "атомарно" означает, что операция выполняется неделимо и не может быть прервана другими операциями**
type Client struct {
	Age int64
}
type SonClient struct {
	Age int64
}

func addAge(cl *Client, sc *SonClient, add int64) {
	atomic.AddInt64(&cl.Age, add)
	sc.Age = sc.Age + add
}

func main() {
	cl := &Client{}
	sc := &SonClient{}

	for i := 1; i <= 100; i++ {
		go addAge(cl, sc, int64(i))
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("%#v\n", cl) //&main.Client{Age:5050}
	fmt.Printf("%#v\n", sc) //&main.SonClient{Age:4951}
}
## 5.5.3 Мьютекс блокирует процесс для одного пользователя (одной горутины) mu := sync.Mutex{}
Мьютекс - синхронизирует доступ к данным путём явной блокировки (без каналов)
### Создание mu := &sync.Mutex{}
### Example 1
#### Mistake &main.Client{Age:977}
type Client struct {
	Age int
}

func addAge(cl *Client, add int) {
	cl.Age = cl.Age + add
}

func main() {
	cl := &Client{}

	for i := 1; i <= 1000; i++ {
		go addAge(cl, 1)
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("%#v\n", cl)
}
#### Solution &Mutex mu.Lock() mu.Unlock()
**Использование mutex ограничевает доступ к функции только одной goRutine**
type Client struct {
	Age int
}

// Функция принимает ссылку на mutex
func addAge(cl *Client, add int, mu *sync.Mutex) {
	// Если одна горутина работают то остальным доступ закрыт
	mu.Lock()
	cl.Age = cl.Age + add
	// После выполнения разблокировать доступ
	mu.Unlock()
}

func main() {
	cl := &Client{}
	// Обязательно ссылка на sync.Mutex
	mu := &sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		// Передаём ссылку на mutex в функцию
		go addAge(cl, 1, mu)
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("%#v\n", cl)
}
## 5.5.4 Семофор ограничевает количество одновременно запущенных потоков
### Пример
func worker(ctx context.Context, id int, sema chan struct{}, wg *sync.WaitGroup) {
	sema <- struct{}{} // Захват семафора
	defer func() {
		<-sema // Освобождение семафора
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d aborted\n", id)
		return
	default:
		fmt.Printf("Worker %d starts working\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker %d finishes working\n", id)
	}
}

func main() {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()
	// Максимальное количество одновременно работающих горутин
	concurrency := 3

	sema := make(chan struct{}, concurrency)
	wg := sync.WaitGroup{}

	n := 10
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go worker(ctx, i, sema, &wg)
	}

	wg.Wait()
}
## 5.5.5 WaitGroup and time.Duration
### Обязательно
1. var wg sync.WaitGroup
2. firstMessage(message string, ftimes int, wg *sync.WaitGroup) **функция получает ссылку на WaitGroup**
3. wg.Add(2) - установить количество горутин
4. defer wg.Done() **уменьшает количество ожидаемых горутин**
5. wg.Wait() - ожидает пока wg.Done не закроет нужное количество горутин и не даёт main закончится 
### Example 1 and time.Duration
func sleep(t time.Duration, wg *sync.WaitGroup) {
	fmt.Println("Sleep ", t)
	time.Sleep(t)
	// wg.Done - уменьшает количество процессов
	// в счётчике wg.Wait
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	//wg.Add(1) - Даём счётчику wg.Wait понять,
	// сколько процессов нужно ждать
	wg.Add(1)
	go sleep(1*time.Second, wg)

	wg.Add(1)
	go sleep(2*time.Second, wg)

	wg.Add(1)
	go sleep(3*time.Second, wg)

	// wg.Wait() - ожидает пока значение счётчика не станет == 0
	// якорь который не даёт свернутся main
	wg.Wait()
	fmt.Println("End main")
}
### Example 2
func firstMessage(message string, ftimes int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(ftimes) * time.Second)
	fmt.Println(message)
}

func secondMessage(message string, ftimes int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(ftimes) * time.Second)
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go secondMessage("SECOND", 5, &wg)
	go firstMessage("I'm first message", 2, &wg)
	wg.Wait()

	fmt.Println("End main")
}
### Example 3
func main() {
	ch := make(chan int, 12)
	ch2 := make(chan int)
	//done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			d := rand.Intn(11)
			ch <- d
		}
		close(ch)
		fmt.Println("random numbers generated")

	}()

	go func() {
		defer wg.Done()
		for s := range ch {
			d := s * 2
			fmt.Printf("mult(%v, 2)=%v\n", s, d)
			ch2 <- d
			time.Sleep(time.Second * 1)
		}
		close(ch2)

	}()

	go func() {

		for num := range ch2 {
			fmt.Println("read ", num)
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()

}
## 5.5.7 Таймер timer := time.NewTimer(3*time.Second)
### Теория
Позволяет настраивать таймауты и ограничивать время выполнения функций
### Exampl simple
func jobWithTimeout(t *time.Timer, q chan string) {
	time.Sleep(time.Second * 3)
	q <- "Перемога !"
	fmt.Println("end jobwithTimeout")
}

func main() {
	timer := time.NewTimer(3 * time.Second)
	quit := make(chan string)
	// Ложим таймер и канал в функцию
	go jobWithTimeout(timer, quit)

	// Ждём ответ из канала таймера
	// Или из канала quit
	select {
	case <-timer.C:
		fmt.Println("Time over")
	case x := <-quit:
		fmt.Println("Received value from quit:", x)
	}

	fmt.Println("End main")
}
### Example fmt.scan +
func jobWithTimeout(t *time.Timer, q chan string) {
	fmt.Println("What is your name")
	var word string
	fmt.Scan(&word)
	q <- word
}

func main() {
	timer := time.NewTimer(15 * time.Second)
	quit := make(chan string)
	// Ложим таймер и канал в функцию
	go jobWithTimeout(timer, quit)

	// Ждём ответ из канала таймера
	// Или из канала quit

	select {
	case <-timer.C:
		fmt.Println("Time over")
	case x := <-quit:
		fmt.Println(" are you already?\nif it's correct?")
		time.Sleep(time.Second * 3)
		
			if strings.EqualFold(x, answer) {
				fmt.Println("You win")
			} else {
				fmt.Println("Try again")
			}
	}

	fmt.Println("End main")
}
### Smile exampl
func jobWithTimeout(t *time.Timer, q chan string) {
	fmt.Println("What is your name")
	var word string
	fmt.Scan(&word)
	q <- word
}

func main() {
	timer := time.NewTimer(15 * time.Second)
	quit := make(chan string)

	go jobWithTimeout(timer, quit)

	select {
	case <-timer.C:
		fmt.Println("Time over")
	case x := <-quit:
		fmt.Println("are you ready?")
		time.Sleep(time.Second * 3)
		compareNameIs(x)
	}
}

func compareNameIs(name string) {
	names := make(map[string]string)
	names["Вова"] = "Хозяин жизни"
	names["Даша"] = "Мелкая козявка высокого роста"
	names["Аня"] = "Мелкая козявка сладкоешка"
	names["Юля"] = "Любимая наша мамуля"
	if nik, ok := names[name]; ok { // strings.EqualFold как его сюда засунуть?
		fmt.Println("wait i'm work")
		printDots()
		time.Sleep(time.Second * 3)
		fmt.Println(nik)
	} else {
		fmt.Println("Try again")
	}
}
### Exampl тупой
func main() {
	// если 1 Время таймера timer закончилось
	// если 3 Function successfull
	var first = time.Duration(3) 
	var last = time.Duration(2)
	//timer
	timer := time.NewTimer(first * time.Second)

	// канал без буфера
	quit := make(chan int)

	// функция принимает таймер и канал
	// тоесть если она выпонится за 3 секунды то
	// "Время таймера timer закончилось"
	go jobWithTimeout(timer, quit)

	// ждём 2 секунды пока горутина выполнится
	time.Sleep(last * time.Second)
	// Выполняем работу...
	// quit <- 1 - если отправить 

	fmt.Println("End main")
}

// функция получает таймер (который запускается сам:) и
// канал. Ждёт что произойдёт раньше (закончится таймер или
// что то прийдёт в канал)
func jobWithTimeout(t *time.Timer, q chan int) {
	var middle = time.Duration(1)
	// ждём 1 секунду
	time.Sleep(middle * time.Second)
	// типа роутер для каналов
	select {
	// когда выйдет время таймера == сработает канал <-t.C
	// у таймера тоже есть канал
	case <-t.C:
		fmt.Println("Время таймера timer закончилось")
	// канал таймера для принудительной остановки
	// который невозможно просто спровоцировать в этом коде
	case <-q:
		if !t.Stop() {
			<-t.C
		}
		fmt.Println("Timer stoped")
		// если таймер не успеет сработать
		// и никто принудительно не остановит таймер
	default:
		fmt.Println("Function successfull")
	}
}
## 5.5.8 Time stopwatch (Секундомер)
### result := testing.Benchmark(BenchmarkMyFunction)
func myFunction() {
	i := 0
	for {
		i++
		time.Sleep(time.Second * 1)
		if i == 10 {
			break
		}

	}
}

func BenchmarkMyFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFunction()
	}
}

func main() {
	result := testing.Benchmark(BenchmarkMyFunction)
	fmt.Println("Время выполнения:", result.T)
}
### duration := time.Since(startTime)
func myFunction() {
	i := 0
	for {
		i++
		time.Sleep(time.Second * 1)
		if i == 10 {
			break
		}

	}
}

func main() {
	startTime := time.Now()

	// Вызов функции, время выполнения которой нужно измерить
	myFunction()

	duration := time.Since(startTime)
	fmt.Println("Время выполнения:", duration)
	fmt.Printf("Время выполнения: %.0f seconds\n", duration.Seconds())
}
## 5.5.9 Ticker ticker:=time.NewTicker(1*time.Second)
### Ticker main()
func main() {
	ticker := time.NewTicker(1 * time.Second)

	count := 0
	for tick := range ticker.C {
		count++
		fmt.Println(count, tick)
		if count == 10 {
			ticker.Stop()
			break
		}
	}
}
### Ticker(chan bool, chan string)
func main() {
	done := make(chan bool)
	stringa := make(chan string)
	fmt.Println("i'm ticker")
	go ticker(done, stringa)
	time.Sleep(time.Second * 1)
	fmt.Println("Write something")

	time.Sleep(time.Second * 1)
	var d string
	fmt.Scan(&d)
	stringa <- d
	fmt.Println(d)
	<-done
}

func ticker(done chan<- bool, stringa <-chan string) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	word := "Hello"

	count := 0
	for range ticker.C {
		count++
		fmt.Println(count, word)
		if count == 10 {
			break
		}
		select {
		case c := <-stringa:
			word = c
		default:
			word = "Hello"
		}
	}
	done <- true
}
## 5.5.10 Пакет Context ctx := context.Background()
### Теория
**кнопка отмены** ctx, cancel := context.WithCancel(context.Background())
**таймер** ctx, _ := context.WithTimeout(context.Background()
**расписание** context.WithDeadLine()
### ctx, cancel := context.WithCancel(context.Background())
#### Пример simple
func main() {
	// Создаем контекст с помощью context.Background()
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем горутину, которая будет выполнять работу
	go doWork(ctx)

	// Ждем некоторое время
	time.Sleep(2 * time.Second)

	// Отменяем контекст, чтобы сигнализировать о завершении работы
	cancel()

	// Ждем, пока горутина завершит работу
	time.Sleep(1 * time.Second)
}

func doWork(ctx context.Context) {
	// Создаем контекст с помощью context.WithCancel(), чтобы иметь возможность отменить работу
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // Убеждаемся, что функция cancel() будет вызвана, чтобы освободить ресурсы

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Работа прервана")
			return
		default:
			fmt.Println("Работа выполняется...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
#### Пример Slurm
func main() {
	//Получаем ссылку на контекст и функцию отмены
	ctx, cancel := context.WithCancel(context.Background())

	//Запускаем 10 горутин которые получают инт для 
	//таймера от 0...10
	for i := 0; i <= 10; i++ {
		go sendData(ctx, i)
	}
	//Те горутины что успеют сработать за этот timeSleep
	//будут считатся успешными
	//потому что следом будет cancel который отменит все 
	//ранее запущенные горутины
	time.Sleep(time.Second * 5)
	cancel()
	//этот sleep для того что б успели напечататся 
	//aborted
	time.Sleep(time.Millisecond * 500)
}

func sendData(ctx context.Context, num int) {
	timer := time.NewTimer(time.Duration(num) * time.Second)

	select {
		//Функция cancel передаёт сигнал ctx.Done
	case <-ctx.Done():
		fmt.Printf("Procces #%v aborted \n", num)
		return
	case <-timer.C:
		fmt.Printf("Date procces #%v send successfully\n", num)
	}
}
### ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
#### Пример 1
func main() {
	parent := context.Background()
	//Получаем ссылку на контекст и функцию отмены
	ctx, _ := context.WithTimeout(parent, 1*time.Second)

	for i := 1; i <= 10; i++ {
		go sendData(ctx, i)
	}

	time.Sleep(time.Second)
	//cancel()
	//cancel произойдёт сам через время из ctx
	time.Sleep(time.Millisecond * 500)
}

func sendData(ctx context.Context, num int) {
	timer := time.NewTimer(time.Duration(num) * time.Second)

	select {
	//Функция context.WithTimeout передаст сигнал cancel
	//в заданое время сигнал ctx.Done
	case <-ctx.Done():
		fmt.Printf("Procces #%v aborted \n", num)
		return
	case <-timer.C:
		fmt.Printf("Date procces #%v send successfully\n", num)
	}
}
#### Пример context.WithTimeout + wg 
func main() {
	wg := &sync.WaitGroup{}
	parent := context.Background()
	ctx, _ := context.WithTimeout(parent, 3*time.Second)
	n := 10
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go sendData(ctx, i, wg)
	}
	wg.Wait()
}

func sendData(ctx context.Context, num int, wg *sync.WaitGroup) {
	timer := time.NewTimer(time.Duration(num) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Printf("Procces #%v aborted \n", num)
		wg.Done()
		return
	case <-timer.C:
		fmt.Printf("Date procces #%v send successfully\n", num)
	}
	wg.Done()
}
### context.WithDeadLine()
Позволяет устанавливать время срабатывания
# 6 Go на практике
## 6.1 Структура и управление зависимостями
### Принципы организации кода
1. Single responsibilities - каждая функция отвечает за что то одно
2. Функции и методы класса - отвечают только за работу с классом (type)
### Структуры организации кода
1. Source file (функции, методы...)
2. Package - source files обьеденены одним смыслом (обьектом, назначением)
3. Module - Хранилище packages
### go get, go tidy - качает все пакеты которые не скачаны в проекте
### Vendoring - go mod vendor
Скачать все пакеты с зависимостями (типа сторонние пакеты) и положить их локально на пк
## 6.2 Работа с ОС, флаги
### **Package OS**
1.Работа с системой
2.Абстракция над реальными функциями
3.Заточена под Linux системы
### Package OS can
1. Запуск приложений
2. Работа с файлами и папками
3. Отслеживание и управление процессами
### OS.exec - Запуск команд из go в bash
1. Оборачивает системные вызовы
2. Не использует шелл с паттернами
3. Может не корректно работать с виндовс
### Practice 
#### run firefox
**SLURM/exec/simple.go**

package exec

import (
	"log"
	"os/exec"
)

func RunSimpleApp() {
	cmd := exec.Command("firefox")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
**SLURM/main.go**
package main

import "SLURM/exec"

func main() {
	exec.RunSimpleApp()
}
- go build SLURM
- ./SLURM
**SLURM - будет работать пока я не закрою firefox**
#### tr Run_app_with_simple_args
func Run_app_with_simple_args() {
	// типа bash $echo "Little slurm goes big" | tr 'a-z' 'A-Z'
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("Little slurm goes big")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String())
}
#### echo Run_multiple_args
func Run_multiple_args() {
	prg := "echo"

	arg1 := "there"
	arg2 := "are slurms"
	arg3 := "in SlurmLand"

	cmd := exec.Command(prg, arg1, arg2, arg3)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
#### Input_pipe()
func Input_pipe() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "an old slurm")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	//fmt.Printf("#{out}\n")
}
#### Capture_output
func Capture_output() {
	out, err := exec.Command("ls", "-l").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
#### Output_pipe
func Output_pipe() {
	cmd := exec.Command("echo", "piping slurms")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(data))
}
### Переменные окружения
- Установка специфичных настроек для приложения (пароли, флаги и т.д)
- Управление приложением из вне
- Принадлежать пользователю
### Flag
#### flag bool
func main() {
	// Определение флагов
	aPtr := flag.Bool("a", false, "Флаг 'a' для вывода 'привет'")
	bPtr := flag.Bool("b", false, "Флаг 'b' для вывода 'пока'")

	// Парсинг флагов командной строки
	flag.Parse()

	// Вывод значений флагов
	if *aPtr {
		fmt.Println("привет")
	}
	if *bPtr {
		fmt.Println("пока")
	}
}
#### flag string
func main() {
	var defaultVal = "Привет"
	// Определение флагов с дефолтными значениями
	aPtr := flag.String("a", defaultVal, "Значение флага 'a'")

	// Парсинг флагов командной строки
	flag.Parse()

	// Вывод значений флагов
	fmt.Println("Значение флага 'a':", *aPtr)

}
//makefile
build:
	go build
modinit:
	go mod init flag2
goFlag2:
	./flag2.exe
goFlag2PlusString:
	./flag2.exe -a=hello
#### flag StringVar
func main() {
	var variable string //переменная принимает значение
	var defaultVal = "8080"
	// Определение флагов с дефолтными значениями
	flag.StringVar(&variable, "path", defaultVal, "description")

	// Парсинг флагов командной строки
	flag.Parse()

	if variable == "env" {
		fmt.Println(".env")
	} else if variable == "toml" {
		fmt.Println("toml")
	} else {
		fmt.Println(variable)
	}
}

// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe -path=toml
// toml
// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe -path=abracadabra
// abracadabra
// PS C:\Users\Mvmir\go\src\github.com\Wolfxxxz\flag2> ./flag2.exe                  
// 8080
## 6.3 Работа с файлами в го
### os.ReadFile(string)

func read(filepath string) {
	dat, err := os.ReadFile(filepath)
	check(err)
	fmt.Print(string(dat))
}

func main() {
	read("test.txt")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
### os.Open, os.Read, os.Seek, io.ReadAtLeast, bufio.NewReader, os.Close - всё на куче
func read(filepath string) {
	//проверяет наличие и доступность прав на чтение файла
	//если всё ок возвращает открытый файл
	f, err := os.Open(filepath)
	check(err)

	//os.Read([]byte)
	//Читает байты (с открытого файла) сколько влезет в массив
	b1 := make([]byte, 7)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Выполнение операций с открытым файлом
	// Изменение позиции чтения/записи
	// offset := int64(6)
	// whence := 0
	// newPos, err := f.Seek(offset, whence)
	o3, err := f.Seek(6, 0)
	check(err)

	// Считает только количество байт
	// f - источник (os.Open)
	// b3 - буфер в который будут прочитаны данные
	// 2 - гарантированное количество прочитанных байт
	// n3 - количество байтов, фактически
	//прочитанных из источника данных.
	// err - mistake
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("io.ReadAtLeast:  %d bytes @ %d: %s\n", n3, o3, string(b3))
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(15)
	check(err)
	fmt.Printf("15 bytes: %s\n", string(b4))

	f.Close()
}

func main() {
	read("test.txt")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
### Чтение из файла до разделителя
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	delimiter := byte('\n') // Определяем символ-разделитель
	arrStr := []string{}

	for {
		line, err := reader.ReadString(delimiter)
		if err != nil {
			// Если конец файла
			break
		}

		// Убираем символ-разделитель из строки
		line = strings.TrimSuffix(line, string(delimiter))

		arrStr = append(arrStr, string(line))
		//fmt.Println("Read line:", line)

		// Если достигнут конец файла, прекращаем чтение
		if err == bufio.ErrBufferFull || err == io.EOF {
			break
		}
	}
	fmt.Println(arrStr)
	fmt.Println("the end")
}
### f.Seek(offset, whence)
//После прочтения символов флажок смещается на
//количество прочитанных символов

func read(filepath string) {
	//Получаем файл
	f, err := os.Open(filepath)
	check(err)

	//Вычитываем первые 7 байт
	b2 := make([]byte, 7)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%s\n", string(b2[:n2])) //Hello

	// Читаем Следующие 7 байт
	b4 := make([]byte, 7)
	n4, err := f.Read(b4)
	check(err)
	fmt.Println("n4 ", n4)         //n4  7
	fmt.Printf("%s\n", string(b4)) //how are

	// Смещение позиции чтения
	//сместить на количество байт
	//сместить можно вперёд или назад + -
	//offset := int64(7)
	offset := int64(-7)
	//откуда - (0 - с начала, 1 - с текущей позиции, с конца)
	whence := 2
	_, err = f.Seek(offset, whence)
	check(err)

	// Читаем следующие 7 байт
	b3 := make([]byte, 7)
	n3, err := f.Read(b3)
	check(err)
	fmt.Printf("%s\n", string(b3[:n3])) //Hello

	f.Close()
}

func main() {
	read("test.txt")
	//Hello
	//how are you
	//i'm well
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
## 6.4 Сеть, модель OSI, TCP & UPD, HTTP
### Api не дал?
### Анализ трафика не дал?
### 6.4.1 Как работает интернет или Модель OSI
#### Уровни OSI
1. - Физический - железо + драйверы
     Кодирование декодирование сигнала
2. - Канальный - MAC, физическая адресация устройств
3. - Сетевой - IP, NAT и прочее. адресация и нахождение маршрутов
4. - Транспортный - TCP, UDP
     Запись данных (упаковка байтов)
5. - Сесионный - RPC, PAP 
     Установка сесии между клиентами для обмена сообщениями
6. - Презентационный - HTTP agent, SSL, XDR
     Дешифровка и шифрование
7. - Прикладной - HTTP, Telnet, SMTP
     Работа с прикладными клиентами
### 6.4.3 TCP & UDP - Транспортный уровень (только отправка и получение)
#### Theory 
Transmission Control Protocol
User Datagram Protocol (Unreliable)
- Оба транспортных протокола содержат данные
- Делят данные на пакеты для передачи по сети
- Работают на уровне конкретных машин
- Содержат инфо о портах машин
- Отличия
        TCP                         UDP
 Гарантирует доставку данных || Не гарантирует
 Пакет большого размера      || маленький размер
 Медленный                   || быстрый
 Используют в ЭП, загрузке   || Стриминг видео
 файлов, API                 || VoIP
#### TCP Протокол с подтверждением (use Hercules SETUP utility by HW-group.com & Windows)
##### TCP Server code
**https://www.hw-group.com/software/hercules-setup-utility**
func main() {
	TCPServer()
}

func TCPServer() {
	// Открываем соединение на listen(слушать) или на Write(писать)
	// пакет net сам оформляет (заворачивает пакеты) в зависимости
	// от протокола (tcp - с подтверждением)
	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		os.Exit(1)
	}

	//если не закрыть приложение то придётся искать этот
	//процесс и убивать его принудительно
	//будет висеть в hook (петля) и кушать ресурсы
	defer l.Close()

	//бесконечный цикл который слушает(ждёт) поступление данных
	fmt.Println("Listening on localhost:1234")
	for {
		// метод соединение возвращает net.Conn
		// произошло соединение
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}

		// запускаем хендлеры в мультиплекс режиме (паралельно)
		// что бы наш сервер мог принимать несколько соединений
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	//Read []byte{}
	buf := make([]byte, 1024)
	//Байты в буф, количество байт в reqLen
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(reqLen)
	fmt.Println(string(buf[:reqLen]))

	//Write []byte{}
	bytesWritten, err := conn.Write([]byte("Message received."))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	fmt.Printf("Bytes written %d\n", bytesWritten)

	//close connection
	conn.Close()
}
##### TCP Client (отправка)
func TCPClient() {
	p := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Printf("Some err #{err}")
		return
	}
	fmt.Fprintf(conn, "H1 TCP Server, how are you doing?")
	_, err = bufio.NewReader(conn).Read(p)

	if err == nil {
		fmt.Printf("%s", p)
	} else {
		fmt.Printf("Some error #{err}\n")
	}
	conn.Close()
}
#### UDP Протокол (быстрый) но беззащитный (use Hercules SETUP utility by HW-group.com & Windows)
##### UDP Server
// Если пришло сообщение кинуть назад ответ
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello i got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v\n", err)
	}
}

func UDPServer() {
//Можно структуркой где ip и порт
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)

	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	fmt.Println("Listening on localhost:1234")

	p := make([]byte, 2048)
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)

		if err != nil {
			fmt.Printf("Some error %v\n", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
		// Напечатать сообщение и
		// закрыть сервер:)
		if len(p) != 0 {
			fmt.Println(string(p))
			break
		}
		//go readMessage(ser, remoteaddr)
	}
}
##### UDP Client
func UDPClient() {
	p := make([]byte, 1024)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some err %v\n", err)
		return
	}
	fmt.Fprintf(conn, "H1 UDP Server, how are you doing?")
	_, err = bufio.NewReader(conn).Read(p)

	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some err %v\n", err)
	}
	conn.Close()
}
### 6.4.5 HTTP Прикладной уровень
#### Theory
HTTP - 90 % сайтов используют именно его
Имеет несколько версий(1.1,2,3) - 1.1 и 2 используют за основу tcp, 3 версия имеет свой протокол QUICK
Протокол запрос ответ
Используется в большинстве Api в интернете
- есть статус коды
- get put delete post(update)
- heders 
#### HTTP server (use insomnia & Windows)
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", h, name)
		}
	}
}

func HTTPServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
//insomnia
get http://localhost:8080/hello
answer - hello
get http://localhost:8080/headers
answer - insomnia/2023.2.2: User-Agent
         */*: Accept
get http://localhost:8080/got
answer - 404 page not found
#### HTTPs server (security защищённый)
func http.ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error

- работает на основе приватный и публичный ключи
- дальше сами ищите ключи...
#### HTTP client
func HTTPClientSimpleGet() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func HTTPClientHeadersGet() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/headers", nil)
	req.Header.Add("test", "test")
	resp, err := client.Do(req)
	//resp,err := http.Get("http:/localhost:8080/headers")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	for name, headers := range resp.Header {
		for _, h := range headers {
			fmt.Printf("H: %v : %v\n", name, h)
		}
	}
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
## 6.5 Протоколы обмена
### Json Marshal
#### 6.5.1 Theory Маршалинг
Маршалинг - конвертация данных в передаваемый формат
Анмаршалинг - разворачивание данных после маршалинга
#### 6.5.3 JSON - javaScript object notation
##### https://jsonformatter.curiousconcept.com/
##### Json Unmarshal simple
type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func ParseJson() {
	birdJson := `{"species":"pigeon","description":"likes to perch onrocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	err := json.Unmarshal([]byte(birdJson), &bird)

	if err != nil {
		panic(err)
	}
	fmt.Println(bird)
}
##### Json Marshal simple
func CreateJson() {
	bird := Bird{
		Species:     "Eagle",
		Description: "Cool eagle",
		Dimensions: Dimensions{
			Height: 100,
			Width:  50,
		},
	}
	//data, _ := json.Marshal(bird) //форматировать в строчку
	data, _ := json.MarshalIndent(bird, "", "    ") //форматирует читабельно
	fmt.Println(string(data))
}
##### My simple json
type Person struct {
	Name     string `json:"pogonjalo"`
	LastName string `json:"lastName"`
	Age      int    `json:"years"`
}

func TakeJson(doc string) ([]Person, bool) {
	var person []Person
	jsonData, err := os.ReadFile(doc)
	if err != nil {
		fmt.Println(err)
		return person, false
	}
	json.Unmarshal(jsonData, &person)
	return person, true
}

// Устаревший подход
func TakeJsonOpen(doc string) ([]Person, bool) {
	jsonFile, err := os.Open(doc)
	var person []Person
	if err != nil {
		fmt.Println(err)
		return person, false
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return person, false
	}
	json.Unmarshal(byteValue, &person)
	return person, true
}

// Marshaling

func WriteJson(doc string, db []Person) {
	byteArr, _ := json.MarshalIndent(db, "", "    ")
	err := os.WriteFile(doc, byteArr, 0666) //-rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}
}
#### 6.5.5 
