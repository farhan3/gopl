//package main
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// assigns String method to type Celsius
func (c Celsius) String() string {
	return fmt.Sprintf("%g*C", c)
}

func CToF(c Celsius) Fahrenheit {
	//  Fahrenheit(float64) is a type conversion
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Println(AbsoluteZeroC)

}
