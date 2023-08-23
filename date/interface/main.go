package main

/*
Интерфейс это контракт, который говорит, что у типа, который его реализует,
	должны быть методы с такими-то именами и сигнатурами.
 	Тип, который реализует интерфейс, автоматически удовлетворяет контракту интерфейса.
*/

type Storage interface {
	Add(a int) error
	Get(id int) (int, error)
}

type storage struct {
}

func (s *storage) Add(_ int) error {
	return nil
}

func (s *storage) Get(id int) (int, error) {
	return 0, nil
}

func main() {

}
