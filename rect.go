// Package rect implements utilities for manipulating axis aligned 2d rectangles
package rect

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Vec is vector in 2d space with Integer Coordinates
type Vec struct {
	X, Y int
}

// Add standalone function implements lhs + rhs
func Add(lhs, rhs Vec) Vec {
	return Vec{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

// Add method implements c = c + rhs. Modifies receiver value
func (c *Vec) Add(rhs Vec) {
	c.X += rhs.X
	c.Y += rhs.Y
}

// Sub standalone function implements lhs - rhs
func Sub(lhs, rhs Vec) Vec {
	return Vec{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

// Sub method implements c = c - rhs. Modifies receiver value
func (c *Vec) Sub(rhs Vec) {
	c.X -= rhs.X
	c.Y -= rhs.Y
}

// Rectangle is an Axis aligned 2d Rectangle
// It comprises 2 Vec values, The Min Value and The Max Value, that describe
// apposing corners of the rectangle
// A rectangle is considered to be "normal" if Min.X <= Max.X && Min.Y <= Max.Y
// All the functions (apart from Normalise) expect supplied rectangles to be in
// a normal format
type Rectangle struct {
	Min, Max Vec
}

// XYWH constructs a Rectangle from an x,y position and width and height
func XYWH(x, y, w, h int) Rectangle {
	return Rectangle{Vec{x, y}, Vec{x + w, y + h}}
}

// XYXY constructs a Rectangle from an 2 x,y positions - the supplied values
// do not need to be in a normal form as this function normalises the resultant rectangle
func XYXY(x1, y1, x2, y2 int) Rectangle {
	return Normalise(Rectangle{Vec{x1, y1}, Vec{x2, y2}})
}

// FromPosSize constructs a Rectangle from a Position & a Size Vector
func FromPosSize(pos, size Vec) Rectangle {
	return Rectangle{pos, Add(pos, size)}
}

// FromSize constructs a Rectangle from a Size Vector. In this case The Min value of
// the rectangle will be {0,0}
func FromSize(size Vec) Rectangle {
	return Rectangle{Max: size}
}

// Normalise transforms the rectangle into its normal form.
func (r *Rectangle) Normalise() {
	if r.Min.X > r.Max.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}

	if r.Min.Y > r.Max.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
}

// Normalise transforms the rectangle into its normal form.
// Standalone version which does not modify its paramaters
func Normalise(r Rectangle) Rectangle {
	r.Normalise()
	return r
}

// Width returns the size of the rectangle in the X axis
func (r *Rectangle) Width() int {
	return r.Max.X - r.Min.X
}

// Height returns the size of the rectangle in the Y axis
func (r *Rectangle) Height() int {
	return r.Max.Y - r.Min.Y
}

// Size returns the size of the rectangle as a Vec type
func (r *Rectangle) Size() Vec {
	return Vec{r.Width(), r.Height()}
}

// IsEmpty returns true if the width and height of the Rectangle are both zero
func (r *Rectangle) IsEmpty() bool {
	return (r.Min.X == r.Max.X) || (r.Min.Y == r.Max.Y)
}

// IsNormal returns true if the Rectangle is normalised
func (r *Rectangle) IsNormal() bool {
	return (r.Min.X <= r.Max.X) && (r.Min.Y <= r.Max.Y)
}

// Expands the rectangle in each direction by the size specified in c
func (r *Rectangle) Expand(c Vec) {
	r.Min.Sub(c)
	r.Max.Add(c)
}

// Expand the rectangle in each direction by the size specified in c
// Stand alone function version
func Expand(r Rectangle, c Vec) Rectangle {
	r.Expand(c)
	return r
}

// Translate the rectangle position by the offset specified in c
func (r *Rectangle) Translate(c Vec) {
	r.Min.Add(c)
	r.Max.Add(c)
}

// Translate the rectangle position by the offset specified in c
// Stand alone function version
func Translate(r Rectangle, c Vec) Rectangle {
	r.Translate(c)
	return r
}

// PointInRectangle tests to see if the point p is inside the rectangle r
// returns true if it is
func PointInRectangle(r Rectangle, p Vec) bool {
	return (r.Min.X <= p.X) && (r.Min.Y <= p.Y) && (p.X < r.Max.X) && (p.Y < r.Max.Y)
}

// Intersection returns a rectangle that is the interection between the
// two supplied rectangles.
// if the rectangles Intersect, returns intersection Rectangle, true
// if none Intersecting, returns non-normal rectangle, false
func Intersection(r1, r2 Rectangle) (intersect Rectangle, ok bool) {
	intersect = Rectangle{
		Min: Vec{max(r1.Min.X, r2.Min.X), max(r1.Min.Y, r2.Min.Y)},
		Max: Vec{min(r1.Max.X, r2.Max.X), min(r1.Max.Y, r2.Max.Y)},
	}

	ok = intersect.IsNormal()
	return
}

// Union returns a rectangle that is the smallest rectangle, containing both
// the supplied rectangles
func Union(r1, r2 Rectangle) Rectangle {
	return Rectangle{
		Min: Vec{min(r1.Min.X, r2.Min.X), min(r1.Min.Y, r2.Min.Y)},
		Max: Vec{max(r1.Max.X, r2.Max.X), max(r1.Max.Y, r2.Max.Y)},
	}
}

// Contains returns true if rInner, is completly contained within rOuter
func Contains(rOuter, rInner Rectangle) bool {
	return PointInRectangle(rOuter, rInner.Min) && PointInRectangle(rOuter, rInner.Max)
}
