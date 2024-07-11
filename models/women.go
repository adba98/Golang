package models

type Women struct {
	Man
}

func (w *Women) Breathe()    { w.breathing = true }
func (w *Women) Think()      { w.thinking = true }
func (w *Women) Eat()        { w.eating = false }
func (w *Women) Sex() string { return "Female" }
