package main

import "fmt"

type sString []string

type out_of_bounds struct {
	s        sString
	i        int
	elemento string
}

func (o out_of_bounds) Error() string {
	return fmt.Sprintf("Fuera de rango: no se puede colocar el elemento %q en el lugar %v de la slice %q\n", o.elemento, o.i, o.s)
}

func (s sString) cambiar(elemento string, lugar int) error {
	if len(s) < lugar || lugar < 0 {
		return out_of_bounds{s, lugar, elemento}
	}
	s[lugar] = elemento
	return nil
}

func main() {
	s := sString([]string{"Palabra", "Word", "Omelete au fromage"})
	e := s.cambiar("Omelete du fromage", 5)
	if e != nil {
		fmt.Println(e)
	}
}
