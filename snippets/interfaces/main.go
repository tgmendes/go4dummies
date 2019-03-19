package main

type Reader interface { // HL
	Read(p []byte) (int, error) // HL
} // HL

type Stringer struct{}

func (s *Stringer) Read(p []byte) (int, error) { // HL
	return 1, nil
}

func main() {
	s := &Stringer{}
	s.Read([]byte("hello"))
}
