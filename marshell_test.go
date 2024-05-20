package jsonparcer_go

import "testing"

type TestStruct struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Surname  string `json:"surname"`
	Age      int    `json:"age"`
	Age2     int    `json:"age2"`
}

func TestDecode(t *testing.T) {
	expected := TestStruct{
		Name:     "qwe",
		LastName: "Qwe",
		Surname:  "QWE",
		Age:      123,
		Age2:     23,
	}

	result, err := Unmarshalling[TestStruct]()
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("expected: %#v, got: %#v", expected, result)
	}
}
