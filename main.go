package fraction

type Fraction struct {
	Bottom int
	Top    int
}

func (f *Fraction) Add(other *Fraction)  {
	if f.Bottom != other.Bottom{
		tmp := *f
		r := *other
		f.Top *= r.Bottom
		f.Bottom *= r.Bottom
		r.Top *= tmp.Bottom
		r.Bottom *= tmp.Bottom
		f.Top += r.Top
	}else{
		f.Top += other.Top
	}
}

func (f *Fraction) Subtraction(other *Fraction)  {
	if f.Bottom != other.Bottom{
		tmp := *f
		r := *other
		f.Top *= r.Bottom
		f.Bottom *= r.Bottom
		r.Top *= tmp.Bottom
		r.Bottom *= tmp.Bottom
		f.Top -= r.Top
	}else{
		f.Top -= other.Top
	}
}

func (f *Fraction) MultiplicationWithInt(other int)  {
	f.Top *= other
}

func (f *Fraction) Multiplication(other *Fraction)  {
	//WIP
}

func (f *Fraction) Abbreviation() {
	d := gcd(f.Top, f.Bottom)
	f.Top /= d
	f.Bottom /= d
}

func MixedToFraction(num int, fraction *Fraction) *Fraction {
	return &Fraction{Bottom: fraction.Bottom, Top: num * fraction.Bottom + fraction.Top}
}
