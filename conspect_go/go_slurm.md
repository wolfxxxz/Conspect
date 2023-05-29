
## 5.1 Паралелизм Многозадачность
Свойство среды паралельно обрабатывать задачи
### Вытесняющая многозадачность
Система сама распределяет приоритеты
windows Linux and Mac
### конкурентная многозадачность Go
где планировщик горутин самостоятельно управляет выполнением горутин. Каждая горутина выполняется независимо и может быть приостановлена или возобновлена планировщиком горутин в любой момент. Планировщик горутин определяет, когда горутина будет выполнена и на каком процессоре.
### Переключение контекста
Желательно сохранится при переключении на другую задачу
## 5.2.1 Процесс и Поток Треды и горутины
Процесс - это запуск программы
Поток - запуск части процесса
Goroutine - это легковесный поток который может выполнятся паралельно и независимо от выполнения остального кода в программе
### Горутина на практике
#### Без горутины
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
#### +go
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
#### горутины выполняются разное количество времени
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
### Пакет runtime.Gosched() - типа паузы 
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
## 5.3.1 Обработка ошибок
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
Паника - это функция которая создаёт ошибку рантайма
panic("message")
recErr := recover() // recErr == message
### Чуть теории
Runtime (рантайм) в контексте программирования обычно относится к программной среде, которая обеспечивает выполнение и управление программой во время её работы. В языке Go (Golang) пакет runtime предоставляет функциональность, связанную с выполнением программы, такую как управление горутинами, аллокацией памяти, сборкой мусора и другими операциями, связанными с запущенной программой.

Некоторые из функций и возможностей, предоставляемых пакетом runtime в Go, включают:

go - для запуска новой горутины.
GOMAXPROCS - для установки максимального числа используемых процессоров.
NumCPU - для получения количества доступных процессоров.
MemStats - для получения статистики использования памяти.
GC - для явного запуска сборки мусора.
Gosched - для переключения выполнения между горутинами.
Panic и Recover - для обработки паники (непредвиденного сбоя программы) и восстановления контроля над программой.
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
## 5.4.1 Каналы
### Каналы нужны для взаимодействия между горутинами
### канал как входящий аргумент Функции
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
	// будет больше чем созданых в буфере == дедлок
	ch := make(chan int, 2)

	ch <- 50
	ch <- 75

	go readChan(ch)

	ch <- 100

	go readChan(ch)
	go readChan(ch)

	time.Sleep(1 * time.Second)

	fmt.Println("END MAIN")
}
## 5.4.3 Канал в цыкле for i := range ch; close(ch)
// ch chan<- int  - канал доступен только на запись
// ch <-chan int  - канал доступен только на чтение
func writeChan(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}
	// закрываем канал
	// он типа ставит метку что работа закончина
	// если его не закрыть то он останется открытым на приём и
	// при чтении из него всё будет ждать от него новых данных
	close(ch)
	// ждём пока кто то решит прочитать этот канал
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
## 5.4.5 Select
### Инструкция Select
select {
    case c <- x:
         fmt.Println(c)
    case <- quit:
         fmt.Println("quit)
         return
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
	//Внимательно НЕбуферезированный канал так работать не будет
	ch <- 1

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
## 5.5.1 Базовые принципы синхронизации
## 5.5.1 Примитивы синхронизации Мьютекс и семофор
## 5.5.3 Мьютекс блокирует процесс для одного пользователя
### Семофор ограничевает количество потоков
## WaitGroup, Каналы, Select
## Пакет Context

## 