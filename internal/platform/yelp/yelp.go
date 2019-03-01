package yelp

type Restaurants []string

func (r *Restaurants) append(t string) {
	*r = append(*r, t)
}
