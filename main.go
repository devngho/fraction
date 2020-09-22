package fraction

type Fraction struct {
	Bottom int
	Top int
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
