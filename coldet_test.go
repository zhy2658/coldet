package coldet

import (
	"math"
	"testing"
)

func floatEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) < 1e-6
}

func vec3Equal(a, b [3]float32) bool {
	return floatEqual(a[0], b[0]) && floatEqual(a[1], b[1]) && floatEqual(a[2], b[2])
}

func TestNewBoundingPoint(t *testing.T) {
	pos := [3]float32{1, 2, 3}
	p := NewBoundingPoint(pos)
	if p == nil {
		t.Fatal("NewBoundingPoint returned nil")
	}
	if !vec3Equal(p.position, pos) {
		t.Errorf("expected position %v, got %v", pos, p.position)
	}
}

func TestNewBoundingSphere(t *testing.T) {
	pos := [3]float32{0, 0, 0}
	radius := float32(5)
	s := NewBoundingSphere(pos, radius)
	if s == nil {
		t.Fatal("NewBoundingSphere returned nil")
	}
	if !vec3Equal(s.position, pos) || s.radius != radius {
		t.Errorf("expected pos=%v radius=%v, got pos=%v radius=%v", pos, radius, s.position, s.radius)
	}
}

func TestNewBoundingBox(t *testing.T) {
	pos := [3]float32{0, 0, 0}
	w, l, h := float32(2), float32(4), float32(6)
	b := NewBoundingBox(pos, w, l, h)
	if b == nil {
		t.Fatal("NewBoundingBox returned nil")
	}
	if b.width != w || b.length != l || b.height != h {
		t.Errorf("expected w=%v l=%v h=%v, got w=%v l=%v h=%v", w, l, h, b.width, b.length, b.height)
	}
}

func TestPointDistance(t *testing.T) {
	p := NewBoundingPoint([3]float32{0, 0, 0})
	to := [3]float32{3, 4, 0}
	dist := p.Distance(to)
	if !floatEqual(dist, 5) {
		t.Errorf("expected distance 5, got %v", dist)
	}
}

func TestPointClosestPoint(t *testing.T) {
	pos := [3]float32{1, 2, 3}
	p := NewBoundingPoint(pos)
	to := [3]float32{10, 20, 30}
	closest := p.ClosestPoint(to)
	if !vec3Equal(closest, pos) {
		t.Errorf("expected closest point %v, got %v", pos, closest)
	}
}

func TestSphereDistance(t *testing.T) {
	s := NewBoundingSphere([3]float32{0, 0, 0}, 2)
	to := [3]float32{5, 0, 0}
	dist := s.Distance(to)
	if !floatEqual(dist, 3) {
		t.Errorf("expected distance 3 (5-2), got %v", dist)
	}
}

func TestCheckAabbVsAabb(t *testing.T) {
	b1 := NewBoundingBox([3]float32{0, 0, 0}, 2, 2, 2)
	b2 := NewBoundingBox([3]float32{1, 1, 1}, 2, 2, 2)
	if !CheckAabbVsAabb(*b1, *b2) {
		t.Error("overlapping AABBs should collide")
	}

	b3 := NewBoundingBox([3]float32{10, 10, 10}, 2, 2, 2)
	if CheckAabbVsAabb(*b1, *b3) {
		t.Error("non-overlapping AABBs should not collide")
	}
}

func TestCheckPointInAabb(t *testing.T) {
	b := NewBoundingBox([3]float32{0, 0, 0}, 2, 2, 2)
	pInside := NewBoundingPoint([3]float32{0, 0, 0})
	pOutside := NewBoundingPoint([3]float32{5, 5, 5})

	if !CheckPointInAabb(*pInside, *b) {
		t.Error("point inside AABB should be detected")
	}
	if CheckPointInAabb(*pOutside, *b) {
		t.Error("point outside AABB should not be detected")
	}
}

func TestCheckPointInSphere(t *testing.T) {
	s := NewBoundingSphere([3]float32{0, 0, 0}, 5)
	pInside := NewBoundingPoint([3]float32{1, 1, 1})
	pOutside := NewBoundingPoint([3]float32{10, 10, 10})

	if !CheckPointInSphere(*pInside, *s) {
		t.Error("point inside sphere should be detected")
	}
	if CheckPointInSphere(*pOutside, *s) {
		t.Error("point outside sphere should not be detected")
	}
}

func TestCheckSphereVsSphere(t *testing.T) {
	s1 := NewBoundingSphere([3]float32{0, 0, 0}, 2)
	s2 := NewBoundingSphere([3]float32{3, 0, 0}, 2)
	s3 := NewBoundingSphere([3]float32{10, 0, 0}, 2)

	if !CheckSphereVsSphere(*s1, *s2) {
		t.Error("overlapping spheres should collide (distance 3 < 4)")
	}
	if CheckSphereVsSphere(*s1, *s3) {
		t.Error("non-overlapping spheres should not collide")
	}
}

func TestCheckSphereVsAabb(t *testing.T) {
	s := NewBoundingSphere([3]float32{0, 0, 0}, 2)
	bOverlap := NewBoundingBox([3]float32{2, 0, 0}, 2, 2, 2)
	bNoOverlap := NewBoundingBox([3]float32{10, 10, 10}, 2, 2, 2)

	if !CheckSphereVsAabb(*s, *bOverlap) {
		t.Error("sphere overlapping AABB should collide")
	}
	if CheckSphereVsAabb(*s, *bNoOverlap) {
		t.Error("sphere not overlapping AABB should not collide")
	}
}

func TestAABBClosestPoint(t *testing.T) {
	b := NewBoundingBox([3]float32{0, 0, 0}, 2, 2, 2)
	to := [3]float32{5, 5, 5}
	closest := b.ClosestPoint(to)
	expected := [3]float32{1, 1, 1}
	if !vec3Equal(closest, expected) {
		t.Errorf("expected closest point %v, got %v", expected, closest)
	}
}
