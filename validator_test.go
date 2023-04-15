package validator

import (
	"fmt"
	"testing"
)

func TestNewValidator(t *testing.T) {
	v := New()
	if v.Errors == nil {
		t.Error("expected errors map not nil, got nil")
	}
}

func TestNoErrorsIsValid(t *testing.T) {
	v := New()
	valid := v.Valid()
	if !valid {
		t.Errorf("expected valid to be true, got %v", valid)
	}
}

func TestAddError(t *testing.T) {
	v := New()
	v.AddError("test", "this is the test message")
	valid := v.Valid()
	if valid {
		t.Errorf("expected valid to be false, got %v", valid)
	}
}

func TestCheck(t *testing.T) {
	type test struct {
		subject  any
		target   any
		test     string
		expected any
	}
	cases := []test{
		{"", "", "equals", true},
		{"Test", "Test", "equals", true},
		{"Test", "", "not equal", true},
		{"test", "", "length", 4},
	}
	v := New()
	for _, test := range cases {
		switch test.test {
		case "equals":
			v.Check(test.subject == test.target, "subject", fmt.Sprintf("'%s' should equal '%s'", test.subject, test.target))
		case "not equal":
			v.Check(test.subject != test.target, "subject", fmt.Sprintf("'%s' should not equal '%s'", test.subject, test.target))
		case "length":
			v.Check(len(test.subject.(string)) == test.expected, "subject", fmt.Sprintf("'%s' expected length '%d'", test.subject, test.expected))
		}
	}
}

func TestPermittedValue(t *testing.T) {
	type test struct {
		permitted []string
		value     string
		expected  bool
	}
	cases := []test{
		{[]string{"A", "B", "C"}, "A", true},
		{[]string{"A", "B", "C"}, "D", false},
	}
	for _, test := range cases {
		permitted := PermittedValue(test.value, test.permitted...)
		if permitted != test.expected {
			t.Errorf("expected '%v', got '%v", test.expected, permitted)
		}
	}
}

func TestUnique(t *testing.T) {
	type test struct {
		values   []string
		expected bool
	}
	cases := []test{
		{[]string{"A", "B", "C", "D"}, true},
		{[]string{"A", "B", "C", "C"}, false},
	}
	for _, test := range cases {
		unique := Unique(test.values)
		if unique != test.expected {
			t.Errorf("expected '%v', got '%v", test.expected, unique)
		}
	}
}
