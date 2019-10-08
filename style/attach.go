package style

import "github.com/qlova/seed/style/css"

//Attacher is used to specify an attachpoint.
type Attacher struct {
	style Style
}

//Top attaches this to the top of the context.
func (a Attacher) Top() Attacher {
	a.style.SetTop(css.Zero)
	return a
}

//Bottom attaches this to the bottom of the context.
func (a Attacher) Bottom() Attacher {
	a.style.SetBottom(css.Zero)
	return a
}

//Left attaches this to the left of the context.
func (a Attacher) Left() Attacher {
	a.style.SetLeft(css.Zero)
	return a
}

//Right attaches this to the right of the context.
func (a Attacher) Right() Attacher {
	a.style.SetRight(css.Zero)
	return a
}

//AttachToScreen attaches this element to the screen, returns an attacher to specify where.
func (style Style) AttachToScreen() Attacher {
	style.SetPosition(css.Fixed)
	return Attacher{style}
}

//AttachToParent attaches this element to its parent, returns an attacher to specify where.
func (style Style) AttachToParent() Attacher {
	style.SetPosition(css.Absolute)
	return Attacher{style}
}

//StickyToScreen sticks this element to the screen, returns an attacher to specify where.
func (style Style) StickyToScreen() Attacher {
	style.SetPosition(css.Sticky)
	return Attacher{style}
}
