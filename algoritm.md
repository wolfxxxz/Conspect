# Algoritm
Набор последовательных действий которые решают какую то задачу
## Сложность алгоритма
O(Log2n) - хорошая скорость
O(n) - количеству элементов в массиве
O(n*log2n)
O(n*n)
O(n!) - стремится к безконечности
## Линейный поиск (сложность O(n))
func main() {
	arr := []int{8, 2, 9, 4, 7, 6, 5, 1, 3}
	fmt.Println(lookForNumber(3, arr))
}

func lookForNumber(lookNum int, arr []int) (counter int) {
	for _, v := range arr {
		counter++
		if v == lookNum {
			return counter
		}
	}
	fmt.Println("your number doesn't exist")
	return counter
}
## Бинарный поиск O(Log2n)
1. Массив отсортирован по порядку
func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	fmt.Println(binarySearch(arr, 5))
}

func binarySearch(arr []int, target int) int {
	count := 0
	low := 0
	high := len(arr) - 1

	for low <= high {
		count++
		// Средний елемент
		mid := (low + high) / 2
		if arr[mid] == target {
			return count
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Возвращаем -1, если элемент не найден
}
## Сортировка выбором O(n*n)
func main() {
	arr := []int{5, 3, 17, 7, 15, 11, 13, 9, 1}
	fmt.Println(selectionSort(arr))
}

// Сортировка выбором
func selectionSort(arr []int) ([]int, int) {
	var count int
	for i := 0; i < len(arr); i++ {
		var indexMin = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[indexMin] {
				indexMin = j
			}
			count++
		}
		var tmp = arr[i]
		arr[i] = arr[indexMin]
		arr[indexMin] = tmp
	}
	return arr, count
}
## сортировка Пузырьком O(n*n)
func main() {
	arr := []int{8, 7, 6, 5, 10, 9, 4, 3, 2, 1}
	fmt.Println(bubbleSort(arr))
}

// Сортировка пузырьком
func bubbleSort(arr []int) ([]int, int) {
	var count int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				var tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
			count += 1
		}
	}
	return arr, count
}
## рекурсия и Быстрая сортировка или сортировка Хуара O(n)
### Используется рекурсия Факториал и числа фибоначи
#### Факториал
func main() {
	num := 5
	result := factorial(num)
	fmt.Printf("Факториал числа %d равен %d\n", num, result)
}

func factorial(n int) int {
    //условие выхода (функция факториал не будет вызвына)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
#### Сумма Чмсел фибоначи 0, 1, 1, 2, 3, 5, 8
func main() {
	num := 8
	fmt.Println(fibonachi(num))
}

func fibonachi(i int) int {
    // Условие выхода рекурсивная функция не будет вызвана
	if i == 1 || i == 2 {
		return 1
	}
	return fibonachi(i-1) + fibonachi(i-2)
}
### Быстрая сортировка (ХОАРА) O(log2n*n)
var Count int

func main() {
	arr := []int{8, 7, 6, 5, 10, 9, 4, 3, 2, 1}
	fmt.Println(quickSort(arr), Count)
}

// Быстрая сортировка
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	less := []int{}
	greater := []int{}
	for i := 0; i < len(arr); i++ {
		Count += 1
		if i == pivotIndex {
			continue
		}
		if arr[i] < pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
### Бинарный поиск (рекурсия) O(Log2n)
func main() {
	arr := []int{-1, 1, 3, 5, 7, 9, 11, 13, 15, 17}
	fmt.Println(recursiveBinarySearch(arr, 7, 0, len(arr)))
}

// Рекурсивный бинарный поиск
func recursiveBinarySearch(arr []int, item int, start int, end int) int {
	if start > end {
		return -1
	}
	middle := (start + end) / 2
	if item == arr[middle] {
		return middle
	}
	if item < arr[middle] {
		return recursiveBinarySearch(arr, item, 0, middle-1)
	}
	if item > arr[middle] {
		return recursiveBinarySearch(arr, item, middle+1, end)
	}
	return -1

}
## Графы (координаты)
### Поиск в ширину
type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.getVertex(k1)
	v2 := g.getVertex(k2)
	if v1 == nil || v2 == nil {
		return
	}
	if contains(v1.adjacent, k2) {
		return
	}
	v1.adjacent = append(v1.adjacent, v2)
	v2.adjacent = append(v2.adjacent, v1)
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) BreadthFirstSearch(start *Vertex, end *Vertex) bool {
	queue := []*Vertex{start}
	visited := make(map[*Vertex]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return true
		}

		visited[current] = true

		for _, v := range current.adjacent {
			if !visited[v] {
				queue = append(queue, v)
			}
		}
	}

	return false
}

func main() {
	g := Graph{}

	for i := 0; i < 6; i++ {
		g.AddVertex(i)
	}

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	fmt.Println(g.BreadthFirstSearch(g.getVertex(0), g.getVertex(5))) // true
	fmt.Println(g.BreadthFirstSearch(g.getVertex(0), g.getVertex(6))) // false
}
### https://www.youtube.com/watch?v=NErrGZ64OdE&t=2412s

## Кеширование
# Теория
## Очередь
Первый пришёл -> первый ушёл
## Стек
как стопка бумаги - последний листок оказывается верхним и его забирают первым