package utils

import rl "github.com/gen2brain/raylib-go/raylib"

func MakeLighter(color rl.Color, factor float32) rl.Color {
	r := uint8(clamp(float32(color.R)+(255.0-float32(color.R))*factor, 0, 255))
	g := uint8(clamp(float32(color.G)+(255.0-float32(color.G))*factor, 0, 255))
	b := uint8(clamp(float32(color.B)+(255.0-float32(color.B))*factor, 0, 255))
	a := color.A
	return rl.Color{R: r, G: g, B: b, A: a}
}

func MakeDarker(color rl.Color, factor float32) rl.Color {
	r := uint8(clamp(float32(color.R)*(1.0-factor), 0, 255))
	g := uint8(clamp(float32(color.G)*(1.0-factor), 0, 255))
	b := uint8(clamp(float32(color.B)*(1.0-factor), 0, 255))
	a := color.A
	return rl.Color{R: r, G: g, B: b, A: a}
}

func clamp(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
