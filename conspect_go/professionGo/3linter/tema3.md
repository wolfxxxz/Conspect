## линтер - проверка + коменты
go install github.com/mgechev/revive@latest
// Ставить в каждый проэкт сразу после go mod init иначе тупит жёстко
*** revive** - показывает ошибки
*** go doc -all** - показывает коменты
***// revive:disable:exported*** - отключить эту функцию от комментов
func PrintHello() {
	fmt.Println("hello")
}
146