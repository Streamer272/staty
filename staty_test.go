package staty

import "testing"

func TestStaty(t *testing.T) {
	t.Parallel()
	t.Run("test store", func(t *testing.T) {
		store := NewStore()
		store.Set("str", "value")
		store.Set("int", 1)
		store.Set("bool", true)
		store.Set("struct", struct{ name string }{name: "World"})

		if store.Get("str") != "value" {
			t.Errorf("name is '%v', but should be 'value'", store.Get("str"))
		}
		if store.Get("int") != 1 {
			t.Errorf("name is '%v', but should be 1", store.Get("int"))
		}
		if store.Get("bool") != true {
			t.Errorf("name is '%v', but should be true", store.Get("bool"))
		}
		if store.Get("struct") != struct{ name string }{name: "World"} {
			t.Errorf("name is '%v', but should be struct{name: \"World\"}", store.Get("struct"))
		}
	})
}
