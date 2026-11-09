package juego

type Personaje struct {
	nombre  string
	vida    int
	vidaMax int
} //recuerdos de vietnam (registros ADA)
// NewPersonaje es el "constructor", segun vi, go no tiene, asi que lo digo por convencion
// si no empieza con Mayuscula, entonces no queda exportado xd
func NewPersonaje(nombre string, vidaMax int) *Personaje {
	if vidaMax <= 0 {
		vidaMax = 1 // min para que nazca vivo
	}
	return &Personaje{nombre: nombre, vida: vidaMax, vidaMax: vidaMax}
}

// Metodos en mayuscula: publicos (se exportan)
func (p *Personaje) Nombre() string { return p.nombre }
func (p *Personaje) Vida() int      { return p.vida }
func (p *Personaje) EstaVivo() bool { return p.vida > 0 }

// RecibirDano es la unica forma de bajar vida: tonces se protege el estado
func (p *Personaje) RecibirDano(cant int) {
	if cant < 0 {
		return
	}
	p.vida -= cant
	if p.vida < 0 {
		p.vida = 0
	}
}
