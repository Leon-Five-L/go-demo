package main

/*
	map 的 value 是 struct 时，如何更新 struct 中的值
*/

func modifyMap() {
	// value 为非指针的修改方式
	m := make(map[string]Person)
	m["first"] = Person{Name: "first", Age: 40}
	m["secord"] = Person{Name: "secord", Age: 40}

	// 循环 m 并修改 Age
	for k, v := range m {
		v.Age = 100
		m[k] = v
	}

	// 如果 value 是指针，则可以直接修改
	m2 := make(map[string]*Person)
	m2["first"] = &Person{Name: "first", Age: 40}
	m2["secord"] = &Person{Name: "secord", Age: 40}
	for k := range m2 {
		m2[k].Age = 100
	}

}
