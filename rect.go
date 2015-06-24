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

// Coord is the a basic cordinate in 2d space
type Coord struct {
	X, Y int
}

// Add standalone function implements lhs + rhs
func Add(lhs, rhs Coord) Coord {
	return Coord{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

// Add method implements c = c + rhs. Modifies receiver value
func (c *Coord) Add(rhs Coord) {
	c.X += rhs.X
	c.Y += rhs.Y
}

// Sub standalone function implements lhs - rhs
func Sub(lhs, rhs Coord) Coord {
	return Coord{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

// Sub method implements c = c - rhs. Modifies receiver value
func (c *Coord) Sub(rhs Coord) {
	c.X -= rhs.X
	c.Y -= rhs.Y
}

// Rectangle is an Axis aligned 2d Rectangle
// It comprises 2 Coord values, The Min Value and The Max Value.
// A rectangle is considered to be "normal" if Min.X <= Max.X && Min.Y <= Max.Y
// All the functions (apart from Normalise) expect supplied rectangles to be in
// a normal format
type Rectangle struct {
	Min, Max Coord
}

// XYWH constructs a Rectangle from an x,y position and width and height
func XYWH(x, y, w, h int) Rectangle {
	return Rectangle{Coord{x, y}, Coord{x + w, y + h}}
}

// XYWH constructs a Rectangle from an 2 x,y positions - the supplied values
// do not need to be in a normal form as this function normalises the resultant rectangle
func XYXY(x1, y1, x2, y2 int) Rectangle {
	return Normalise(Rectangle{Coord{x1, y1}, Coord{x2, y2}})
}

func FromPosSize(pos, size Coord) Rectangle {
	return Rectangle{pos, Add(pos, size)}
}

func FromSize(size Coord) Rectangle {
	return Rectangle{Max: size}
}

func (r *Rectangle) Normalise() {
	if r.Min.X > r.Max.X {
		r.Min.X, r.Max.X = r.Max.X, r.Min.X
	}

	if r.Min.Y > r.Max.Y {
		r.Min.Y, r.Max.Y = r.Max.Y, r.Min.Y
	}
}

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

// Size returns the size of the rectangle as a Coord type
func (r *Rectangle) Size() Coord {
	return Coord{r.Width(), r.Height()}
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
func (r *Rectangle) Expand(c Coord) {
	r.Min.Sub(c)
	r.Max.Add(c)
}

// Expands the rectangle in each direction by the size specified in c
// Stand alone function version
func Expand(r Rectangle, c Coord) Rectangle {
	r.Expand(c)
	return r
}

// Translates the rectangle position by the offset specified in c
func (r *Rectangle) Translate(c Coord) {
	r.Min.Add(c)
	r.Max.Add(c)
}

// Translates the rectangle position by the offset specified in c
// Stand alone function version
func Translate(r Rectangle, c Coord) Rectangle {
	r.Translate(c)
	return r
}

// PointInRectangle tests to see if the point p is inside the rectangle r
// returns true if it is
func PointInRectangle(r Rectangle, p Coord) bool {
	return (r.Min.X <= p.X) && (r.Min.Y <= p.Y) && (p.X < r.Max.X) && (p.Y < r.Max.Y)
}

// Intersection returns a rectangle that is the interection between the
// two supplied rectangles.
func Intersection(r1, r2 Rectangle) (intersect Rectangle, ok bool) {
	intersect = Rectangle{
		Min: Coord{max(r1.Min.X, r2.Min.X), max(r1.Min.Y, r2.Min.Y)},
		Max: Coord{min(r1.Max.X, r2.Max.X), min(r1.Max.Y, r2.Max.Y)},
	}

	ok = intersect.IsNormal()
	return
}

func Union(r1, r2 Rectangle) Rectangle {
	return Rectangle{
		Min: Coord{min(r1.Min.X, r2.Min.X), min(r1.Min.Y, r2.Min.Y)},
		Max: Coord{max(r1.Max.X, r2.Max.X), max(r1.Max.Y, r2.Max.Y)},
	}
}

func Contains(rOuter, rInner Rectangle) bool {
	return PointInRectangle(rOuter, rInner.Min) && PointInRectangle(rOuter, rInner.Max)
}
