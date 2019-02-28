package main

func main() {
	s := &Stringer{}
	s.Read([]byte("hello"))
}

type Reader interface { // HL
	Read(p []byte) (int, error) // HL
} // HL

type Stringer struct{}

func (s *Stringer) Read(p []byte) (int, error) { // HL
	return 1, nil
}
