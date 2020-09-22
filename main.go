package fraction

type Fraction struct {
	bottom int
	top int
}

func (f *Fraction) Add(other *Fraction)  {
	if f.bottom != other.bottom{
		tmp := *f
		f.top *= other.bottom
		f.bottom *= other.bottom
		other.top *= tmp.bottom
		other.bottom *= tmp.bottom
		f.top += other.top
	}else{
		f.top += other.top
	}
}
