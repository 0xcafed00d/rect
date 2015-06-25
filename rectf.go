// Package rect implements utilities for manipulating axis aligned 2d rectangles
package rect

func minf(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func maxf(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

// Vecf is vector in 2d space with floating point Coordinates
type Vecf struct {
	X, Y float64
}

// Add standalone function implements lhs + rhs
func Addf(lhs, rhs Vecf) Vecf {
	return Vecf{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

// Add method implements c = c + rhs. Modifies receiver value
func (c *Vecf) Add(rhs Vecf) {
	c.X += rhs.X
	c.Y += rhs.Y
}

// Sub standalone function implements lhs - rhs
func Subf(lhs, rhs Vecf) Vecf {
	return Vecf{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

// Sub method implements c = c - rhs. Modifies receiver value
func (c *Vecf) Sub(rhs Vecf) {
	c.X -= rhs.X
	c.Y -= rhs.Y
}

// Rectanglef is an Axis aligned 2d Rectangle
// It comprises 2 Vecf values, The Min Value and The Max Value, that describe
// opposing corners of the Rectangle
// A Rectanglef is considered to be "normal" if Min.X <= Max.X && Min.Y <= Max.Y
// All the functions (apart from Normalise) expect supplied rectangles to be in
// a normal format
type Rectanglef struct {
	Min, Max Vecf
}

// XYWHf constructs a Rectanglef from an x,y position and width and height
func XYWHf(x, y, w, h float64) Rectanglef {
	return Rectanglef{Vecf{x, y}, Vecf{x + w, y + h}}
}

// XYWHf constructs a Rectanglef from an 2 x,y positions - the supplied values
// do not need to be in a normal form as this function normalises the resultant Rectanglef
func XYXYf(x1, y1, x2, y2 float64) Rectanglef {
	return Normalisef(Rectanglef{Vecf{x1, y1}, Vecf{x2, y2}})
}

// FromPosSize constructs a Rectanglef from a Position & a Size Vector
func FromPosSizef(pos, size Vecf) Rectanglef {
	return Rectanglef{pos, Addf(pos, size)}
}

// FromSize constructs a Rectanglef from a Size Vector. In this case The Min value of
// the Rectanglef will be {0,0}
func FromSizef(size Vecf) Rectanglef {
	return Rectanglef{Max: size}
}

// Normalise transforms the Rectanglef float64o its normal form.
func (r *Rectanglef) Normalise() {
	if r.Min.X > r.Max.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}

	if r.Min.Y > r.Max.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
}

// Normalise transforms the Rectanglef float64o its normal form.
// Standalone version which does not modify its paramaters
func Normalisef(r Rectanglef) Rectanglef {
	r.Normalise()
	return r
}

// Width returns the size of the Rectanglef in the X axis
func (r *Rectanglef) Width() float64 {
	return r.Max.X - r.Min.X
}

// Height returns the size of the Rectanglef in the Y axis
func (r *Rectanglef) Height() float64 {
	return r.Max.Y - r.Min.Y
}

// Size returns the size of the Rectanglef as a Vecf type
func (r *Rectanglef) Size() Vecf {
	return Vecf{r.Width(), r.Height()}
}

// IsEmpty returns true if the width and height of the Rectanglef are both zero
func (r *Rectanglef) IsEmpty() bool {
	return (r.Min.X == r.Max.X) || (r.Min.Y == r.Max.Y)
}

// IsNormal returns true if the Rectanglef is normalised
func (r *Rectanglef) IsNormal() bool {
	return (r.Min.X <= r.Max.X) && (r.Min.Y <= r.Max.Y)
}

// Expands the Rectanglef in each direction by the size specified in c
func (r *Rectanglef) Expand(c Vecf) {
	r.Min.Sub(c)
	r.Max.Add(c)
}

// Expands the Rectanglef in each direction by the size specified in c
// Stand alone function version
func Expandf(r Rectanglef, c Vecf) Rectanglef {
	r.Expand(c)
	return r
}

// Translates the Rectanglef position by the offset specified in c
func (r *Rectanglef) Translate(c Vecf) {
	r.Min.Add(c)
	r.Max.Add(c)
}

// Translates the Rectanglef position by the offset specified in c
// Stand alone function version
func Translatef(r Rectanglef, c Vecf) Rectanglef {
	r.Translate(c)
	return r
}

// PointInRectangle tests to see if the point p is inside the Rectanglef r
// returns true if it is
func PointInRectanglef(r Rectanglef, p Vecf) bool {
	return (r.Min.X <= p.X) && (r.Min.Y <= p.Y) && (p.X < r.Max.X) && (p.Y < r.Max.Y)
}

// Intersection returns a Rectanglef that is the Intersection between the
// two supplied rectangles.
// if the rectangles Intersect, returns Intersection Rectanglef, true
// if none Intersecting, returns non-normal Rectanglef, false
func Intersectionf(r1, r2 Rectanglef) (intersect Rectanglef, ok bool) {
	intersect = Rectanglef{
		Min: Vecf{maxf(r1.Min.X, r2.Min.X), maxf(r1.Min.Y, r2.Min.Y)},
		Max: Vecf{minf(r1.Max.X, r2.Max.X), minf(r1.Max.Y, r2.Max.Y)},
	}

	ok = intersect.IsNormal()
	return
}

// Union returns a Rectanglef that is the smallest Rectanglef, containing both
// the supplied rectangles
func Unionf(r1, r2 Rectanglef) Rectanglef {
	return Rectanglef{
		Min: Vecf{minf(r1.Min.X, r2.Min.X), minf(r1.Min.Y, r2.Min.Y)},
		Max: Vecf{maxf(r1.Max.X, r2.Max.X), maxf(r1.Max.Y, r2.Max.Y)},
	}
}

// Contains returns true if rInner, is completly contained within rOuter
func Containsf(rOuter, rInner Rectanglef) bool {
	return PointInRectanglef(rOuter, rInner.Min) && PointInRectanglef(rOuter, rInner.Max)
}
