package main

import "fmt"

type DB struct {
	driver string
}

func (db *DB) getData(key string) (*Data, error) {
	// simulate something to make it fail if DB's field(s) are not allocated
	_ = db.driver
	_ = key
	return &Data{Name: "BFG", Num: 9000}, nil
}

// ================================================================================

type Data struct {
	Name string
	Num  int
}

// ================================================================================

type Service struct {
	db *DB
}

func (s *Service) Retrieve(key string) (*Data, error) {
	data, err := s.db.getData(key)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ================================================================================

func main() {

	var service Service

	data, err := service.Retrieve("96aa80d5-466c-4612-a45a-91920f3b3788")
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

}
