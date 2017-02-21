package query

import "testing"

func TestMarshal(t *testing.T) {

	type TestType struct {
		GUID      string
		Name      string `query:"nom"`
		lowercase string `query:"lc"`
		Number    int    `query:"mispar"`
	}

	obj := TestType{"abc", "def", "hello", 31}
	v := Marshal(obj)
	obj2 := TestType{}
	err := Unmarshal(v, &obj2)
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}
	expected := map[string]string{"nom": "def", "GUID": "abc", "mispar": "31"}
	for k, val := range expected {
		if val2 := v.Get(k); val2 != val {
			t.Fatalf("Expected %s for key %s, got %s", val, k, val2)
		}
	}
	for k, val := range v {
		if val2 := expected[k]; val2 != val[0] {
			t.Fatalf("Expected %s for key %s, got %s", val, k, val2)
		}
	}
	if obj.GUID != obj2.GUID {
		t.Fatal("GUID mismatch")
	}
	if obj.Name != obj2.Name {
		t.Fatal("Name mismatch")
	}
	if obj.Number != obj2.Number {
		t.Fatal("Number mismatch")
	}
	if obj2.lowercase != "" {
		t.Fatal("Unexpected value in lowercase")
	}

}
