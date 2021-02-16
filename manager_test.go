package main

import (
	"testing"
	"time"

	idgen "github.com/wakeapp/go-id-generator"
	sg "github.com/wakeapp/go-sql-generator"
)

func BenchmarkInsertWithOptimize(t *testing.B) {
	var err error
	var m *SQLManager

	m, err = InitManager()
	if err != nil {
		t.Fatalf("error on connect %s", err)
	}
	defer CloseManager()

	ti := time.Now()
	var d = &sg.InsertData{
		TableName: "TestTable",
		Fields: []string{
			"id",
			"time",
		},
	}

	d.SetOptimize(true)

	var id string
	for count := 1; count <= 10000; count++ {
		id = idgen.Id()

		d.Add(
			[]string{
				id,
				ti.Format("2006-01-02 15:04:05"),
			},
		)
	}

	t.ResetTimer()
	_, err = m.Insert(d)
	if err != nil {
		t.Fatalf("error on insert %s", err)
	}
	t.StopTimer()

	_, err = m.Query("TRUNCATE TABLE TestTable")
	if err != nil {
		t.Fatalf("error on clean %s", err)
	}
}

func BenchmarkInsert(t *testing.B) {
	var err error
	var m *SQLManager

	m, err = InitManager()
	if err != nil {
		t.Fatalf("error on connect %s", err)
	}
	defer CloseManager()

	ti := time.Now()
	var d = &sg.InsertData{
		TableName: "TestTable",
		Fields: []string{
			"id",
			"time",
		},
	}

	rev := func(s string) string {
		runes := []rune(s)
		size := len(runes)
		for i := 0; i < size/2; i++ {
			runes[size-i-1], runes[i] = runes[i], runes[size-i-1]
		}
		return string(runes)
	}

	var id string
	for count := 1; count <= 10000; count++ {
		id = idgen.Id()

		id = rev(id)

		d.Add(
			[]string{
				id,
				ti.Format("2006-01-02 15:04:05"),
			},
		)
	}

	t.ResetTimer()
	_, err = m.Insert(d)
	if err != nil {
		t.Fatalf("error on insert %s", err)
	}
	t.StopTimer()

	_, err = m.Query("TRUNCATE TABLE TestTable")
	if err != nil {
		t.Fatalf("error on clean %s", err)
	}
}
