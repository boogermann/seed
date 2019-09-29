package style

import "github.com/qlova/seed/style/css"

//SetSize sets the width and height of the element. Takes em, vm, px or percentage values.
func (style Style) SetSize(width, height complex128) {
	style.SetWidth(width)
	style.SetHeight(height)
}

//SetTextSize sets the Text Size, a multiple of the default text size.
func (style Style) SetTextSize(size complex128) {
	style.SetFontSize(css.Decode(size))
}

//SetMaxSize sets the max width and height of the element. Takes em, vm, px or percentage values.
func (style Style) SetMaxSize(width, height complex128) {
	style.Style.SetMaxWidth(css.Decode(width))
	style.Style.SetMaxHeight(css.Decode(height))
}

//SetExpand sets this element to expand and take up all available space.
func (style Style) SetExpand(expand float64) {
	style.SetFlexGrow(css.Number(expand))
}

//Expand sets this element to expand and take up all available space.
func (style Style) Expand() {
	style.SetExpand(1)
}

//SetUnshrinkable means this element should not shrink to make space for other elements.
func (style Style) SetUnshrinkable() {
	style.SetFlexShrink(css.Number(0))
}

//DontShrink means this should not shrink to make space for other elements.
func (style Style) DontShrink() {
	style.SetFlexShrink(css.Number(0))
}

//Shrink means this element should shrink to make space for other elements.
func (style Style) Shrink() {
	style.SetFlexShrink(css.Number(1))
}

//Compress means this element should shrink to make space for other elements.
func (style Style) Compress() {
	style.SetFlexShrink(css.Number(1))
}

//SetZoomable sets the style to be zoomable on touch screens. DOES NOT WORK.
func (style Style) SetZoomable() {
	style.Set("touch-action", "pinch-zoom")
}

//SetMaxWidth sets the max width of this element.
func (style Style) SetMaxWidth(width Unit) {
	style.Style.SetMaxWidth(css.Decode(width))
}

//SetMaxHeight sets the max height of this element.
func (style Style) SetMaxHeight(height Unit) {
	style.Style.SetMaxHeight(css.Decode(height))
}

//SetMinWidth sets the min width of this element.
func (style Style) SetMinWidth(width Unit) {
	style.Style.SetMinWidth(css.Decode(width))
}

//SetMinHeight sets the min height of this element.
func (style Style) SetMinHeight(height Unit) {
	style.Style.SetMinHeight(css.Decode(height))
}

//SetWidth sets the desired width of this element.
func (style Style) SetWidth(width Unit) {
	style.Style.SetWidth(css.Decode(width))
}

//SetHeight sets the desired height of this element.
func (style Style) SetHeight(height Unit) {
	style.Style.SetHeight(css.Decode(height))
}
