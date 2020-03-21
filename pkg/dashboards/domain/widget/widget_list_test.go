package widget_test

import (
	"github.com/fpetkovski/newrelic-alert-manager/pkg/dashboards/domain/widget"
	"testing"
)

type TestCase struct {
	first  widget.WidgetList
	second widget.WidgetList
	equals bool
}

func invalidTestCases() []TestCase {
	return []TestCase{
		{
			first: widget.WidgetList{
				newWidget(1),
			},
			second: widget.WidgetList{
				newWidget(2),
			},
			equals: false,
		},
		{
			first: widget.WidgetList{
				newWidget(2),
				newWidget(1),
				newWidget(2),
			},
			second: widget.WidgetList{
				newWidget(1),
				newWidget(1),
				newWidget(2),
			},
			equals: false,
		},
	}
}

func TestWidgetList_Equals_InvalidTestCases(t *testing.T) {
	for i, testCase := range invalidTestCases() {
		equals := testCase.first.Equals(testCase.second)
		if equals != testCase.equals {
			t.Errorf("Test case %d should fail", i)
		}
	}
}

func validTestCases() []TestCase {
	return []TestCase{
		{
			first: widget.WidgetList{
				newWidget(1),
			},
			second: widget.WidgetList{
				newWidget(1),
			},
			equals: true,
		},
		{
			first: widget.WidgetList{
				newWidget(2),
				newWidget(1),
				newWidget(2),
			},
			second: widget.WidgetList{
				newWidget(1),
				newWidget(2),
				newWidget(2),
			},
			equals: true,
		},
	}
}

func TestWidgetList_Equals_ValidTestCases(t *testing.T) {
	for i, testCase := range validTestCases() {
		equals := testCase.first.Equals(testCase.second)
		if equals != testCase.equals {
			t.Errorf("Test case %d should succeed", i)
		}
	}
}

func newWidget(id int) widget.Widget {
	return widget.Widget{
		Visualization: "v" + string(id),
		Data: widget.DataList{
			{
				Nrql: "query " + string(id),
			},
		},
		Layout: widget.Layout{
			Width:  id,
			Height: id,
			Row:    id,
			Column: id,
		},
	}
}