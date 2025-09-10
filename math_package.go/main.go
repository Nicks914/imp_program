package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
)

func main() {

	log.Println(runtime.GOOS)
	log.Println(runtime.GOARCH)
	log.Println(runtime.NumCPU())
	log.Println(runtime.NumGoroutine())
	log.Println(runtime.Version())

	// ---------- Constants ----------
	// Predefined values for mathematical calculations
	fmt.Println("🔹 Constants")
	fmt.Println("Pi:", math.Pi)                     // π constant (circle, trigonometry)
	fmt.Println("E:", math.E)                       // Euler’s number (exponential, growth)
	fmt.Println("Positive Infinity:", math.Inf(1))  // +∞
	fmt.Println("Negative Infinity:", math.Inf(-1)) // -∞
	fmt.Println("NaN:", math.NaN())                 // Not a Number (invalid result like 0/0)
	fmt.Println()

	// ---------- Basic Arithmetic ----------
	// General math operations
	fmt.Println("🔹 Arithmetic")
	fmt.Println("Abs(-42):", math.Abs(-42))       // Absolute value → always positive
	fmt.Println("Pow(2, 5):", math.Pow(2, 5))     // Power (2^5 = 32)
	fmt.Println("Sqrt(49):", math.Sqrt(49))       // Square root (√49 = 7)
	fmt.Println("Cbrt(27):", math.Cbrt(27))       // Cube root (∛27 = 3)
	fmt.Println("Hypot(3, 4):", math.Hypot(3, 4)) // Pythagoras theorem √(3²+4²) = 5
	fmt.Println()

	// ---------- Rounding ----------
	// Rounding numbers in different styles
	fmt.Println("🔹 Rounding")
	fmt.Println("Ceil(4.1):", math.Ceil(4.1))   // Always round UP → 5
	fmt.Println("Floor(4.9):", math.Floor(4.9)) // Always round DOWN → 4
	fmt.Println("Round(4.5):", math.Round(4.5)) // Normal rounding → 5
	fmt.Println("Trunc(4.9):", math.Trunc(4.9)) // Cut off decimal → 4
	fmt.Println("Mod(10, 3):", math.Mod(10, 3)) // Remainder of division → 1
	fmt.Println()

	// ---------- Logarithms & Exponential ----------
	// Growth, decay, and logarithmic functions
	fmt.Println("🔹 Logarithms & Exponential")
	fmt.Println("Exp(2):", math.Exp(2))           // e^2 = 7.389
	fmt.Println("Log(10):", math.Log(10))         // Natural log (ln 10)
	fmt.Println("Log10(1000):", math.Log10(1000)) // Base-10 log → 3
	fmt.Println("Log2(8):", math.Log2(8))         // Base-2 log → 3
	fmt.Println()

	// ---------- Trigonometry ----------
	// Used in geometry, physics, and angles
	fmt.Println("🔹 Trigonometry")
	fmt.Println("Sin(π/2):", math.Sin(math.Pi/2)) // sin(90°) = 1
	fmt.Println("Cos(0):", math.Cos(0))           // cos(0°) = 1
	fmt.Println("Tan(π/4):", math.Tan(math.Pi/4)) // tan(45°) = 1

	fmt.Println("Asin(1):", math.Asin(1))         // arcsin(1) = π/2
	fmt.Println("Acos(0):", math.Acos(0))         // arccos(0) = π/2
	fmt.Println("Atan(1):", math.Atan(1))         // arctan(1) = π/4
	fmt.Println("Atan2(3, 4):", math.Atan2(3, 4)) // slope angle (y=3, x=4)
	fmt.Println()

	// ---------- Min / Max / Special ----------
	// Useful for comparing and checking values
	fmt.Println("🔹 Min / Max / Special")
	fmt.Println("Max(10, 20):", math.Max(10, 20))                     // Bigger of two numbers
	fmt.Println("Min(10, 20):", math.Min(10, 20))                     // Smaller of two numbers
	fmt.Println("Copysign(2, -3):", math.Copysign(2, -3))             // Copy sign from -3 → -2
	fmt.Println("IsNaN(math.NaN()):", math.IsNaN(math.NaN()))         // Check for NaN → true
	fmt.Println("IsInf(math.Inf(1), 1):", math.IsInf(math.Inf(1), 1)) // Check +∞ → true
}
