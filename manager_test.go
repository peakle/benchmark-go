package main

import (
	"fmt"
	"testing"
	"time"

	idgen "github.com/wakeapp/go-id-generator"
	sg "github.com/wakeapp/go-sql-generator"
)

// 11.67s
func BenchmarkInsertWithId(t *testing.B) {
	var err error
	var m *SQLManager

	m, err = InitManager()
	if err != nil {
		t.Fatalf("error on connect %s", err)
	}
	defer CloseManager()

	ti := time.Now()
	var d = &sg.InsertData{
		TableName: "Parsing",
		Fields: []string{
			"id",
			"time",
		},
		WithID: true,
	}

	var id string
	for count := 1; count <= 10000; count++ {
		id = fmt.Sprintf("%d", count)
		d.AddWithID(
			id,
			[]string{
				id,
				ti.Format("2006-01-02 15:04:05"),
			},
		)
	}

	t.ResetTimer()

	t.StartTimer()
	_, err = m.Insert(d)
	if err != nil {
		t.Fatalf("error on insert %s", err)
	}
	t.StopTimer()

	m.Query("TRUNCATE TABLE Parsing")
}

// 13.623s
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
		TableName: "Parsing",
		Fields: []string{
			"id",
			"time",
		},
		WithID: false,
	}

	for count := 1; count <= 10000; count++ {
		d.Add(
			[]string{
				idgen.Id(),
				ti.Format("2006-01-02 15:04:05"),
			},
		)
	}

	t.ResetTimer()

	t.StartTimer()
	_, err = m.Insert(d)
	if err != nil {
		t.Fatalf("error on insert %s", err)
	}
	t.StopTimer()

	m.Query("TRUNCATE TABLE Parsing")
}
