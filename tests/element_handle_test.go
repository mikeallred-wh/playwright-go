package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestElementHandleInnerText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)
	handle, err := page.QuerySelector("#inner")
	require.NoError(t, err)
	t1, err := handle.InnerText()
	require.NoError(t, err)
	require.Equal(t, t1, "Text, more text")
	t2, err := page.InnerText("#inner")
	require.NoError(t, err)
	require.Equal(t, t2, "Text, more text")
}

func TestElementHandleOwnerFrame(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = utils.AttachFrame(page, "iframe1", server.EMPTY_PAGE)
	require.NoError(t, err)
	frame := page.Frames()[1]
	elementHandle, err := frame.EvaluateHandle("document.body")
	require.NoError(t, err)
	ownerFrame, err := elementHandle.(playwright.ElementHandle).OwnerFrame()
	require.NoError(t, err)
	require.Equal(t, ownerFrame, frame)
	require.Equal(t, "iframe1", ownerFrame.Name())
}
func TestElementHandleContentFrame(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = utils.AttachFrame(page, "frame1", server.EMPTY_PAGE)
	require.NoError(t, err)
	elementHandle, err := page.QuerySelector("#frame1")
	require.NoError(t, err)
	frame, err := elementHandle.ContentFrame()
	require.NoError(t, err)
	require.Equal(t, frame, page.Frames()[1])
}
func TestElementHandleGetAttribute(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/dom.html")
	require.NoError(t, err)
	handle, err := page.QuerySelector("#outer")
	require.NoError(t, err)
	a1, err := handle.GetAttribute("name")
	require.NoError(t, err)
	require.Equal(t, "value", a1)
	a2, err := page.GetAttribute("#outer", "name")
	require.NoError(t, err)
	require.Equal(t, "value", a2)
}

func TestElementHandleDispatchEvent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	require.NoError(t, page.DispatchEvent("button", "click"))
	result, err := page.Evaluate("result")
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementHandleHover(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/scrollable.html")
	require.NoError(t, err)
	btn, err := page.QuerySelector("#button-6")
	require.NoError(t, err)
	require.NoError(t, btn.Hover())
	result, err := page.Evaluate(`document.querySelector("button:hover").id`)
	require.NoError(t, err)
	require.Equal(t, "button-6", result)
}

func TestElementHandleClick(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	btn, err := page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, btn.Click())
	result, err := page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementHandleDblclick(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.PREFIX + "/input/button.html")
	require.NoError(t, err)
	_, err = page.Evaluate(`() => {
            window.double = false;
            button = document.querySelector('button');
            button.addEventListener('dblclick', event => {
            window.double = true;
            });
	}`)
	require.NoError(t, err)
	btn, err := page.QuerySelector("button")
	require.NoError(t, err)
	require.NoError(t, btn.Dblclick())
	result, err := page.Evaluate("double")
	require.NoError(t, err)
	require.Equal(t, true, result)

	result, err = page.Evaluate(`result`)
	require.NoError(t, err)
	require.Equal(t, "Clicked", result)
}

func TestElementBoundingBox(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	require.NoError(t, page.SetViewportSize(500, 500))
	_, err := page.Goto(server.PREFIX + "/grid.html")
	require.NoError(t, err)
	element_handle, err := page.QuerySelector(".box:nth-of-type(13)")
	require.NoError(t, err)
	box, err := element_handle.BoundingBox()
	require.NoError(t, err)
	require.Equal(t, 100, box.X)
	require.Equal(t, 50, box.Y)
	require.Equal(t, 50, box.Width)
	require.Equal(t, 50, box.Height)
}

func TestElementHandleTap(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent("<input id='checkbox' type='checkbox'></input>"))
	value, err := page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, false, value)

	elemHandle, err := page.QuerySelector("#checkbox")
	require.NoError(t, err)
	require.NoError(t, elemHandle.Tap())
	value, err = page.EvalOnSelector("input", "el => el.checked")
	require.NoError(t, err)
	require.Equal(t, true, value)
}