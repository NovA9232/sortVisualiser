package animatedArr

func (a *AnimArr) CocktailShakerSort() {
  swapped := true
  for swapped {
    swapped = false
    for i := 0; i < len(a.Data)-1; i++ {
      a.mainBubble(i, &swapped, &CCT_SLEEP) // pass reference to swap rather than returning it (easier).
    }
    if swapped {
      swapped = false
      for i := len(a.Data)-2; i > 0; i-- {
        a.mainBubble(i, &swapped, &CCT_SLEEP)
      }
    }
  }
}
